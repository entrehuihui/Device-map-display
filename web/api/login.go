package api

import (
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
// @Router /login [put]
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
	now := time.Now()
	// 创建账号
	userInfo := db.Userinfo{
		Ownid:       1,
		Name:        registerinfo.Name,
		Password:    registerinfo.Password,
		Photo:       "",
		Permisson:   2,
		VIP:         0,
		Status:      1,
		Createtime:  now.Unix(),
		Updatetime:  now.Unix(),
		Expiredtime: now.AddDate(0, 1, 0).Unix(),
		Mobile:      registerinfo.Mobile,
		Email:       registerinfo.Email,
		Address:     "",
		Details:     "",
	}
	s := service.GetServer()
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
// @Router /login [get]
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
		retError(c, 3, nil)
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
		// 更新用户状态未已过期
		s.UpdateUser(userInfo.ID, map[string]interface{}{
			"status": 3,
		})
		retError(c, 9, nil)
		return
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
