package api

import (
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
// @Param	starttime	query	string	flase	"用户创建起始时间"
// @Param	endtime		query	string	flase	"用户创建终止时间"
// @Param	name		query	string	flase	"用户名"
// @Param	ownid		query	int	flase	"群组id"
// @Param	id			query	int	flase	"用户id"
// @Param	permisson 	query	int	flase	"用户级别"
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
		if status != 0 {
			result = result.Where("status = ?", status)
		}
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
		err = result.Limit(limit).Offset(offset).Find(&userinfo).Error
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
		result = result.Where("expiredtime == 0 or expiredtime > ?", now)
	}
	err = result.Update("", userStatus.Status).Error
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
		uid := c.GetInt("id")
		updates["classid"] = uid
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
	ID int
}

// UserDel .删除用户
// @Tags devices
// @Summary .删除用户(仅管理员以上)
// @Description .删除用户
// @Accept  json
// @Produce  json
// @Param 	Authorization 	header 	string 	true "With the bearer started JWT"
// @param userDelID body api.UserDelID  true "用户ID"
// @Success 200 {string} json "{"Error":"Success","Data": object}"
// @Failure  500 {string} json "{"Error":"error","Data": null}"
// @Failure  301 {string} json "{"Error":"Re-login","Data": object}"
// @Router /devices [delete]
func UserDel(c *gin.Context) {
	userDelID := UserDelID{}
	err := c.ShouldBind(&userDelID)
	if err != nil {
		retError(c, 1, err)
		return
	}
	if userDelID.ID == 0 {
		retError(c, 12, nil)
		return
	}
	s := service.GetServer()
	user, err := s.CheckUserID(userDelID.ID)
	if err != nil {
		retError(c, 7, err)
		return
	}
	if c.GetInt("permisson") != 3 {
		if user.Ownid != c.GetInt("id") {
			retError(c, 18, err)
			return
		}
	}

	updates := map[string]interface{}{
		"del": 1,
	}
	_, err = s.UpdateUser(userDelID.ID, updates)
	if err != nil {
		retError(c, 7, err)
		return
	}
	retSuccess(c, "Success")
}
