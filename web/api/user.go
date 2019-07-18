package api

import (
	"strconv"
	"strings"
	"time"

	"../../db"
	"./service"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// GetUsers 获取用户列表
// @Tags user
// @Summary 获取用户列表
// @Description 获取用户列表
// @Accept  json
// @Produce  json
// @Param 	Authorization 	header 	string 	true "With the bearer started JWT"
// @Param	offset	query	string	flase	"起始数"
// @Param	limit	query	string	flase	"获取数"
// @Param	status	query	string	flase	"用户状态 1启用 2禁用 3已过期"
// @Param	starttime	query	string	flase	"用户到期起始时间"
// @Param	endtime		query	string	flase	"用户到期终止时间"
// @Param	name		query	string	flase	"用户名"
// @Param	ownid		query	int	flase	"群组id"
// @Param	id			query	int	flase	"用户id"
// @Param	permisson 	query	int	flase	"用户级别"
// @Param	vip 		query	int	flase	"VIP等级"
// @Success 200 {string} json "{"Error":"Success","Data": object}"
// @Failure  500 {string} json "{"Error":"error","Data": null}"
// @Failure  301 {string} json "{"Error":"Re-login","Data": object}"
// @Router /users [get]
func GetUsers(c *gin.Context) {
	var err error
	s := service.GetServer()
	permisson := c.GetInt("permisson")
	uid := c.GetInt("id")
	result := &gorm.DB{}
	limit := 30
	offset := 0
	if permisson == 3 {
		result = s.Where(" id != ?", uid)
	} else {
		result = s.Where(" ownid = ?", uid)
	}

	id, err := getID(c)
	if err != nil {
		retError(c, 12, err)
		return
	}
	if id != 0 {
		result = result.Where("id = ?", id)
		limit = 1
	} else {
		//
		ownid, err := getOwnid(c)
		if err != nil {
			retError(c, 24, err)
			return
		}
		if ownid != 0 {
			result = result.Where("ownid = ?", ownid)
		}
		//
		endtime, err := getEndtime(c)
		if err != nil {
			retError(c, 15, err)
			return
		}
		if endtime != 0 {
			result = result.Where("createtime < ?", endtime)
		}
		//
		starttime, err := getStarttime(c)
		if err != nil {
			retError(c, 14, err)
			return
		}
		if starttime != 0 {
			result = result.Where("createtime > ?", starttime)
		}
		//
		status, err := getStatus(c)
		if err != nil {
			retError(c, 13, err)
			return
		}
		if status > 0 && status < 4 {
			if status == 3 {
				result = result.Where("expiredtime <= ? and expiredtime != 0", time.Now().Unix())
			} else {
				result = result.Where("status = ?", status)
			}
		}
		//
		vip, err := getVIP(c)
		if err != nil {
			retError(c, 13, err)
			return
		}
		if vip != 0 {
			result = result.Where("v_ip = ?", vip)
		}
		//
		permissonquery, err := getPermisson(c)
		if err != nil {
			retError(c, 18, err)
			return
		}
		if permissonquery != 0 {
			result = result.Where("permisson = ?", permissonquery)
		}
		//
		name := getName(c)
		if name != "" {
			result = result.Where("name like ?", "%"+name+"%")
		}
		//
		result = result.Where("del = 0")
		//
		limit, err = getLimit(c)
		if err != nil {
			retError(c, 17, err)
			return
		}
		offset, err = getOffset(c)
		if err != nil {
			retError(c, 17, err)
			return
		}
	}
	// 查数量
	userinfo := make([]db.Userinfo, 0)
	all := 0
	err = result.Find(&userinfo).Count(&all).Error
	if err != nil {
		retError(c, 7, err)
		return
	}
	if all != 0 {
		//查数据
		err = result.Limit(limit).Offset(offset).Order("id").Find(&userinfo).Error
		if err != nil {
			retError(c, 7, err)
			return
		}
	}
	retSuccess(c, map[string]interface{}{
		"all":  all,
		"data": userinfo,
	})
}

// UserStatus .用户状态信息
type UserStatus struct {
	// 用户ID
	IDs []int
	// 状态
	Status int
}

// UpdateUserStatus .更改用户状态
// @Tags user
// @Summary 更改用户状态(仅管理员以上)
// @Description 更改用户状态
// @Accept  json
// @Produce  json
// @Param 	Authorization 	header 	string 	true "With the bearer started JWT"
// @param userStatus body api.UserStatus  true "用户状态信息"
// @Success 200 {string} json "{"Error":"Success","Data": object}"
// @Failure  500 {string} json "{"Error":"error","Data": null}"
// @Failure  301 {string} json "{"Error":"Re-login","Data": object}"
// @Router /users/status [put]
func UpdateUserStatus(c *gin.Context) {
	permisson := c.GetInt("permisson")
	uid := c.GetInt("id")
	userStatus := UserStatus{}
	err := c.ShouldBind(&userStatus)
	if err != nil {
		retError(c, 1, err)
		return
	}
	if len(userStatus.IDs) == 0 {
		retSuccess(c, "Success")
		return
	}
	if userStatus.Status < 1 || userStatus.Status > 2 {
		retError(c, 13, nil)
		return
	}
	s := service.GetServer()
	result := s.Model(&db.Userinfo{}).Where("id in (?)", userStatus.IDs)
	if permisson != 3 {
		result = result.Where("ownid = ?", uid)
	}
	result = result.Where("status != 3")
	// 如果是启用 -- 判断是否过期
	if userStatus.Status == 1 {
		now := time.Now().Unix()
		result = result.Where("expiredtime = 0 or expiredtime > ?", now)
	}
	err = result.Update("status", userStatus.Status).Error
	if err != nil {
		retError(c, 7, err)
		return
	}
	retSuccess(c, "Success")
}

// UserExpire .用户过期时间信息
type UserExpire struct {
	// 用户ID
	ID int
	// 过期时间
	Expiredtime int64
}

// UpdateUserExpire .更改用户过期时间
// @Tags user
// @Summary .更改用户过期时间(仅管理员以上)
// @Description .更改用户过期时间
// @Accept  json
// @Produce  json
// @Param 	Authorization 	header 	string 	true "With the bearer started JWT"
// @param userExpire body api.UserExpire  true "用户过期时间信息"
// @Success 200 {string} json "{"Error":"Success","Data": object}"
// @Failure  500 {string} json "{"Error":"error","Data": null}"
// @Failure  301 {string} json "{"Error":"Re-login","Data": object}"
// @Router /users/expire [put]
func UpdateUserExpire(c *gin.Context) {
	userExpire := UserExpire{}
	err := c.ShouldBind(&userExpire)
	if err != nil {
		retError(c, 1, err)
		return
	}
	if userExpire.ID == 0 {
		retError(c, 12, nil)
		return
	}
	status := 3
	if userExpire.Expiredtime == 0 {
		status = 1
	} else {
		now := time.Now().Unix()
		if userExpire.Expiredtime > now {
			status = 1
		} else {
			status = 3
		}
	}
	updates := make(map[string]interface{}, 0)
	updates["status"] = status
	updates["expiredtime"] = userExpire.Expiredtime
	permisson := c.GetInt("permisson")
	if permisson != 3 {
		updates["ownid"] = c.GetInt("id")
	}
	s := service.GetServer()
	_, err = s.UpdateUser(userExpire.ID, updates)
	if err != nil {
		retError(c, 7, err)
		return
	}
	retSuccess(c, "Success")
}

// UserDelID 设备ID
type UserDelID struct {
	// 用户ID
	ID []int
}

// UserDel .删除用户
// @Tags user
// @Summary .删除用户(仅管理员以上)
// @Description .删除用户
// @Accept  json
// @Produce  json
// @Param 	Authorization 	header 	string 	true "With the bearer started JWT"
// @param userDelID body api.UserDelID  true "用户ID"
// @Success 200 {string} json "{"Error":"Success","Data": object}"
// @Failure  500 {string} json "{"Error":"error","Data": null}"
// @Failure  301 {string} json "{"Error":"Re-login","Data": object}"
// @Router /users [delete]
func UserDel(c *gin.Context) {
	userDelID := UserDelID{}
	err := c.ShouldBind(&userDelID)
	if err != nil {
		retError(c, 1, err)
		return
	}
	if len(userDelID.ID) == 0 {
		retError(c, 12, nil)
		return
	}
	db := service.GetServer().Table("userinfos").Where("id in (?) or ownid in (?)", userDelID.ID, userDelID.ID)
	if c.GetInt("permisson") != 3 {
		db = db.Where("ownid = ?", c.GetInt("id"))
	}

	err = db.Update("del", 1).Error
	if err != nil {
		retError(c, 7, err)
		return
	}
	retSuccess(c, "Success")
}

// Info ..信息
type Info struct {
	// 用户ID 修改自身不传
	ID int
	// 手机 不修改不传
	Mobile string
	// 邮箱 不修改不传
	Email string
	// 地址 不修改不传
	Address string
	// 头像  不修改不传
	Photo string
}

// UpdateUserInfo .更改用户信息
// @Tags user
// @Summary .更改用户信息
// @Description .更改用户信息
// @Accept  json
// @Produce  json
// @Param 	Authorization 	header 	string 	true "With the bearer started JWT"
// @param Info body api.Info  true "用户信息"
// @Success 200 {string} json "{"Error":"Success","Data": object}"
// @Failure  500 {string} json "{"Error":"error","Data": null}"
// @Failure  301 {string} json "{"Error":"Re-login","Data": object}"
// @Router /users/info [put]
func UpdateUserInfo(c *gin.Context) {
	userinfo := Info{}
	err := c.ShouldBind(&userinfo)
	if err != nil {
		retError(c, 7, err)
		return
	}
	update := make(map[string]string)
	if userinfo.Mobile != "" {
		if len(userinfo.Mobile) != 11 {
			retError(c, 6, nil)
			return
		}
		if !verifyMobile(userinfo.Mobile) {
			retError(c, 6, nil)
			return
		}
		update["mobile"] = userinfo.Mobile
	}
	if userinfo.Email != "" {
		if len(userinfo.Email) > 100 {
			retError(c, 5, nil)
			return
		}
		if !verifyEmailFormat(userinfo.Email) {
			retError(c, 5, nil)
			return
		}
		update["email"] = userinfo.Email
	}
	if userinfo.Address != "" {
		if len(userinfo.Address) > 255 {
			retError(c, 28, nil)
			return
		}
		update["address"] = userinfo.Address
	}
	newname := ""
	if userinfo.Photo != "" {
		filename := strings.Split(userinfo.Photo, ".")
		if len(filename) != 2 {
			retError(c, 31, nil)
			return
		}
		newname = "/images/user/" + strconv.FormatInt(time.Now().Unix(), 10) + getRandomString(4) + strconv.Itoa(c.GetInt("id")) + "." + filename[1]
		if err := prictureSize(userinfo.Photo, newname, 200, 200); err != nil {
			retError(c, 31, err)
			return
		}
		update["photo"] = newname
	}
	uid := c.GetInt("id")
	if userinfo.ID == 0 {
		userinfo.ID = uid
	}
	s := service.GetServer()
	oldPhoto := Info{}
	if userinfo.Photo != "" {
		err = s.Raw("select photo from userinfos where id = ?", userinfo.ID).Scan(&oldPhoto).Error
		if err != nil {
			retError(c, 7, err)
			return
		}
	}
	db := s.Table("userinfos")
	db = db.Where("id = ?", uid)
	if c.GetInt("permisson") == 2 {
		db = db.Where("ownid = ?", uid)
	}
	err = db.Updates(update).Error
	if err != nil {
		removefile(newname, 1)
		retError(c, 7, err)
		return
	}
	// 成功  删除就图片
	if userinfo.Photo != "" && oldPhoto.Photo != "/images/user/defaultuser.png" {
		removefile(oldPhoto.Photo, 1)
	}
	retSuccess(c, "Success")
}

// PasswordInfo .
type PasswordInfo struct {
	// 用户ID 修改自身不传
	ID int
	// 旧密码 MD5加密 修改下级用户不传
	Oldpassword string
	// 新密码 MD5加密
	Password string
}

// UpdatePassword .更改用户密码
// @Tags user
// @Summary .更改用户密码
// @Description .更改用户密码
// @Accept  json
// @Produce  json
// @Param 	Authorization 	header 	string 	true "With the bearer started JWT"
// @param PasswordInfo body api.PasswordInfo  true "用户密码"
// @Success 200 {string} json "{"Error":"Success","Data": object}"
// @Failure  500 {string} json "{"Error":"error","Data": null}"
// @Failure  301 {string} json "{"Error":"Re-login","Data": object}"
// @Router /users/password [put]
func UpdatePassword(c *gin.Context) {
	password := PasswordInfo{}
	err := c.ShouldBind(&password)
	if err != nil {
		retError(c, 1, err)
		return
	}
	if len(password.Password) != 32 {
		retError(c, 11, nil)
		return
	}
	uid := c.GetInt("id")
	s := service.GetServer()
	userinfo := db.Userinfo{}
	if password.ID != 0 && uid != password.ID {
		userinfo, err = s.CheckUserID(2)
		if err != nil {
			retError(c, 7, nil)
			return
		}
		if c.GetInt("permisson") != 3 && userinfo.Ownid != uid {
			retJump(c, 20)
			return
		}
	} else {
		userinfo, err = s.CheckUserID(uid)
		if err != nil {
			retError(c, 7, nil)
			return
		}
		if userinfo.Password != password.Oldpassword {
			retError(c, 3, nil)
			return
		}
	}

	result := s.Table("userinfos").Where("id = ?", userinfo.ID).Update("password", password.Password)
	if err = result.Error; err != nil || result.RowsAffected != 1 {
		retError(c, 7, err)
		return
	}
	retSuccess(c, "Success")
}
