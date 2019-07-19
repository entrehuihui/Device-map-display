package api

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"strings"
	"time"

	"github.com/jinzhu/gorm"

	"../../db"
	"./service"
	"github.com/gin-gonic/gin"
)

// Registerinfo 申请的用户信息
type Registerinfo struct {
	// 账号
	Name string
	// 邮箱
	Email string
	// 密码 MD5加密(小写)
	Password string
	// 手机号码
	Mobile string
	// 验证码
	Code string
	// 验证码标记
	CodeJWT string
}

// Register .
// @Tags login
// @Summary 账号申请
// @Description 账号申请
// @Accept  json
// @Produce  json
// @param userInfo body api.Registerinfo  true "申请的用户信息"
// @Success 200 {string} json "{"Result": object}"
// @Failure  500 {string} json "{"Result": null}"
// @Failure  301 {string} json "{"Error":"Re-login","Data": object}"
// @Router /login/register [put]
func Register(c *gin.Context) {
	registerinfo := Registerinfo{}
	err := c.ShouldBind(&registerinfo)
	if err != nil {
		retError(c, 1, err)
		return
	}
	// 检测名称
	if registerinfo.Name == "" || len(registerinfo.Name) > 20 {
		retError(c, 2, nil)
		return
	}
	// 检测密码
	if registerinfo.Password == "" || len(registerinfo.Password) != 32 {
		retError(c, 3, nil)
		return
	}
	// 检测email
	if registerinfo.Email == "" || !verifyEmailFormat(registerinfo.Email) {
		retError(c, 5, nil)
		return
	}
	// 检测验证码 --- 待添加
	//电话验证
	if registerinfo.Mobile != "" && !verifyMobile(registerinfo.Mobile) {
		retError(c, 6, nil)
		return
	}
	s := service.GetServer()
	_, err = s.CheckUserName(registerinfo.Name)
	if err == nil {
		retError(c, 8, err)
		return
	}
	if err != gorm.ErrRecordNotFound {
		retError(c, 7, err)
		return
	}
	now := time.Now()
	// 创建账号
	userInfo := db.Userinfo{
		Ownid:       1,
		Name:        registerinfo.Name,
		Password:    registerinfo.Password,
		Photo:       "/images/user/defaultuser.png",
		Permisson:   2,
		VIP:         1,
		Status:      1,
		Createtime:  now.Unix(),
		Updatetime:  now.Unix(),
		Expiredtime: now.AddDate(0, 1, 0).Unix(),
		Mobile:      registerinfo.Mobile,
		Email:       registerinfo.Email,
		Address:     "",
		Details:     "",
	}
	err = s.CreateUser(userInfo)
	if err != nil {
		retError(c, 7, err)
		return
	}
	retSuccess(c, "Success")
}

// CheckUserName .
// @Tags login
// @Summary 检测账号是否存在
// @Description 检测账号是否存在
// @Accept  json
// @Produce  json
// @Param	name	query	string	flase	"用户名"
// @Success 200 {string} json "{"Error":"Success","Data": object}"
// @Failure  500 {string} json "{"Error":"error","Data": null}"
// @Failure  301 {string} json "{"Error":"Re-login","Data": object}"
// @Router /login/register/name [get]
func CheckUserName(c *gin.Context) {
	name := c.Query("name")
	if len(name) > 20 || name == "" {
		retError(c, 2, nil)
		return
	}
	s := service.GetServer()
	_, err := s.CheckUserName(name)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			retSuccess(c, false)
			return
		}
		retError(c, 7, err)
		return
	}
	retSuccess(c, true)
}

// LoginInfos ..
type LoginInfos struct {
	// 登陆名
	Name string
	// 密码 MD5加密(小写)
	Password string
	// 验证码
	Code string
}

// Login .
// @Tags login
// @Summary 账号登陆
// @Description 账号登陆
// @Accept  json
// @Produce  json
// @param userInfo body api.LoginInfos  true "登陆信息"
// @Success 200 {string} json "{"Error":"Success","Data": object}"
// @Failure  500 {string} json "{"Error":"error","Data": null}"
// @Failure  301 {string} json "{"Error":"Re-login","Data": object}"
// @Router /login [post]
func Login(c *gin.Context) {
	loginInfos := LoginInfos{}
	if err := c.ShouldBind(&loginInfos); err != nil {
		retError(c, 1, err)
		return
	}
	// 账号
	if loginInfos.Name == "" || len(loginInfos.Name) > 20 {
		retError(c, 2, nil)
		return
	}
	// 密码
	if len(loginInfos.Password) != 32 {
		retError(c, 10, nil)
		return
	}
	// 验证码 -- 待定

	// 获取账户信息
	s := service.GetServer()
	userInfo, err := s.CheckUserName(loginInfos.Name)
	if err != nil {
		retError(c, 7, err)
		return
	}
	// 判断删除
	if userInfo.Del != 0 {
		retError(c, 10, nil)
		return
	}
	// 判断密码
	if userInfo.Password != loginInfos.Password {
		retError(c, 10, nil)
		return
	}
	// 判断状态
	if userInfo.Status != 1 {
		code := 9
		if userInfo.Status == 2 {
			code = 23
		}
		retError(c, code, nil)
		return
	}
	// 判断有效期
	nows := time.Now().Unix()
	if nows > userInfo.Expiredtime && userInfo.Expiredtime != 0 {
		// 更新用户状态已过期
		_, err = s.UpdateUser(userInfo.ID, map[string]interface{}{
			"status": 3,
		})
		retError(c, 9, err)
		return
	}
	//如果是子账户 -- 判断父账户是否已经过期或者禁用
	if userInfo.Permisson == 1 {
		ownUserinfo, err := s.CheckUserID(userInfo.Ownid)
		if err != nil {
			retError(c, 9, err)
			return
		}
		// 判断父账户
		if ownUserinfo.Status != 1 {
			code := 9
			if ownUserinfo.Status == 2 {
				code = 23
			}
			retError(c, code, nil)
			return
		}
		// 判断父账户是否过期
		if nows > ownUserinfo.Expiredtime && ownUserinfo.Expiredtime != 0 {
			retError(c, 9, nil)
			return
		}
	}

	jwt := service.NewJWT(userInfo.ID, userInfo.Permisson)
	if jwt == "" {
		retError(c, 7, nil)
		return
	}
	// 创建登陆信息
	logininfo := db.Logininfo{
		UID:        userInfo.ID,
		Createtime: nows,
		IP:         c.ClientIP(),
	}
	err = s.CreateLoginfo(logininfo)
	if err != nil {
		retError(c, 7, err)
		return
	}
	// 对手机号码进行隐藏
	if userInfo.Mobile != "" && len(userInfo.Mobile) == 11 {
		userInfo.Mobile = userInfo.Mobile[:3] + "****" + userInfo.Mobile[7:]
	}
	// 对email进行隐藏
	if userInfo.Email != "" {
		emails := strings.Split(userInfo.Email, "@")
		if len(emails) < 2 {
			userInfo.Email = ""
		} else if len(emails[0]) > 3 {
			userInfo.Email = emails[0][:2] + "******" + userInfo.Email[len(emails[0]):]
		} else {
			userInfo.Email = "******" + userInfo.Email[len(emails[0]):]
		}
	}
	// c.Header("Authorization", jwt)
	retSuccess(c, map[string]interface{}{
		"id":        userInfo.ID,
		"jwt":       jwt,
		"name":      userInfo.Name,
		"mobile":    userInfo.Mobile,
		"permisson": userInfo.Permisson,
		"vip":       userInfo.VIP,
		"email":     userInfo.Email,
		"address":   userInfo.Address,
		"photo":     userInfo.Photo,
	})
}

// GetLogin 获取背景图列表....
// @Tags login
// @Summary 获取背景图列表
// @Description 获取背景图列表
// @Accept  json
// @Produce  json
// @Success 200 {string} json "{"Error":"Success","Data": object}"
// @Failure  500 {string} json "{"Error":"error","Data": null}"
// @Failure  301 {string} json "{"Error":"Re-login","Data": object}"
// @Router /login/logo [get]
func GetLogin(c *gin.Context) {
	s := service.GetServer()
	logoinfos := make([]db.Logoinfo, 0)
	err := s.Raw("SELECT * FROM logoinfos WHERE id IN (SELECT max(id) FROM logoinfos WHERE del = 0  GROUP BY types)").Scan(&logoinfos).Error
	if err != nil {
		retError(c, 7, err)
		return
	}
	//
	maps := make(map[int]string, 0)
	maps[1] = "/images/logo/beijing.jpg"
	maps[2] = "/images/logo/beijinglogo.jpg"
	maps[3] = "/images/logo/shanbiaologo.jpg"
	for _, v := range logoinfos {
		maps[v.ID] = v.URL
	}
	retSuccess(c, maps)
}

// GetCode .
// @Tags login
// @Summary 获取错误代码
// @Description 获取错误代码
// @Accept  json
// @Produce  json
// @Success 200 {string} json "{"Error":"Success","Data": object}"
// @Failure  500 {string} json "{"Error":"error","Data": null}"
// @Failure  301 {string} json "{"Error":"Re-login","Data": object}"
// @Router /login/code [get]
func GetCode(c *gin.Context) {
	retSuccess(c, code)
}

// var code = make([]Code, 0)
var code = make([]map[string]interface{}, 0)

// SetCode 读取错误代码
func SetCode() {
	body, err := ioutil.ReadFile("./code.json")
	if err != nil {
		log.Fatal("SetCode error! [error]", err)
	}
	err = json.Unmarshal(body, &code)
	if err != nil {
		log.Fatal("SetCode error! [error]", err)
	}
}

// GetLoginInfo .
// @Tags login
// @Summary 获取登陆信息
// @Description 获取登陆信息
// @Accept  json
// @Produce  json
// @Param 	Authorization 	header 	string 	true "With the bearer started JWT"
// @Success 200 {string} json "{"Error":"Success","Data": object}"
// @Failure  500 {string} json "{"Error":"error","Data": null}"
// @Failure  301 {string} json "{"Error":"Re-login","Data": object}"
// @Router /login/info [get]
func GetLoginInfo(c *gin.Context) {
	// 获取账户信息
	s := service.GetServer()
	userInfo, err := s.CheckUserID(c.GetInt("id"))
	if err != nil {
		c.JSON(301, 19)
		return
	}
	// 判断删除
	if userInfo.Del != 0 {
		c.JSON(301, 19)
		return
	}
	// 判断状态
	if userInfo.Status != 1 {
		c.JSON(301, 19)
		return
	}
	// 判断有效期
	nows := time.Now().Unix()
	if nows > userInfo.Expiredtime && userInfo.Expiredtime != 0 {
		// 更新用户状态未已过期
		s.UpdateUser(userInfo.ID, map[string]interface{}{
			"status": 3,
		})
		c.JSON(301, 19)
		return
	}
	// 创建登陆信息
	logininfo := db.Logininfo{
		UID:        userInfo.ID,
		Createtime: nows,
		IP:         c.ClientIP(),
	}
	err = s.CreateLoginfo(logininfo)
	if err != nil {
		retError(c, 7, err)
		return
	}
	// 对手机号码进行隐藏
	if userInfo.Mobile != "" && len(userInfo.Mobile) == 11 {
		userInfo.Mobile = userInfo.Mobile[:3] + "****" + userInfo.Mobile[7:]
	}
	// 对email进行隐藏
	if userInfo.Email != "" {
		emails := strings.Split(userInfo.Email, "@")
		if len(emails) < 2 {
			userInfo.Email = ""
		} else if len(emails[0]) > 3 {
			userInfo.Email = emails[0][:2] + "******" + userInfo.Email[len(emails[0]):]
		} else {
			userInfo.Email = "******" + userInfo.Email[len(emails[0]):]
		}
	}
	jwt := c.GetHeader("Authorization")
	c.Header("Authorization", jwt)
	retSuccess(c, map[string]interface{}{
		"id":        userInfo.ID,
		"jwt":       jwt,
		"name":      userInfo.Name,
		"mobile":    userInfo.Mobile,
		"permisson": userInfo.Permisson,
		"vip":       userInfo.VIP,
		"email":     userInfo.Email,
		"address":   userInfo.Address,
		"photo":     userInfo.Photo,
	})
}
