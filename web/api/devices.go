package api

import (
	"time"

	"mymap/db"
	"mymap/web/api/service"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// GetDevices 获取设备列表
// @Tags devices
// @Summary 获取设备列表
// @Description 获取设备列表
// @Accept  json
// @Produce  json
// @Param 	Authorization 	header 	string 	true "With the bearer started JWT"
// @Param	offset	query	string	flase	"起始数"
// @Param	limit	query	string	flase	"获取数"
// @Param	status	query	string	flase	"设备状态 1启用 2禁用 3已过期"
// @Param	starttime	query	string	flase	"设备到期起始时间"
// @Param	endtime	query	string	flase	"设备到期终止时间"
// @Param	name	query	string	flase	"设备名"
// @Param	classid	query	int	flase	"群组id"
// @Param	uid		query	int	flase	"用户id"
// @Param	id		query	int	flase	"设备id"
// @Success 200 {string} json "{"Error":"Success","Data": object}"
// @Failure  500 {string} json "{"Error":"error","Data": null}"
// @Failure  301 {string} json "{"Error":"Re-login","Data": object}"
// @Router /devices [get]
func GetDevices(c *gin.Context) {
	var err error
	s := service.GetServer()
	permisson := c.GetInt("permisson")
	uid := c.GetInt("id")
	result := &gorm.DB{}
	limit := 30
	offset := 0
	if permisson == 3 {
		result = s.DB
	} else if permisson == 2 {
		result = s.Where("classid = ?", uid)
	} else {
		result = s.Where("uid = ?", uid)
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
		ownid, err := getUID(c)
		if err != nil {
			retError(c, 27, err)
			return
		}
		if ownid != 0 {
			result = result.Where("uid = ?", ownid)
		}
		//
		classid, err := getClassid(c)
		if err != nil {
			retError(c, 24, err)
			return
		}
		if classid != 0 {
			result = result.Where("classid = ?", classid)
		}
		//
		endtime, err := getEndtime(c)
		if err != nil {
			retError(c, 15, err)
			return
		}
		if endtime != 0 {
			result = result.Where("Expiredtime < ?", endtime)
		}
		//
		starttime, err := getStarttime(c)
		if err != nil {
			retError(c, 14, err)
			return
		}
		if starttime != 0 {
			result = result.Where("Expiredtime > ?", starttime)
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
		//
		name := getName(c)
		if name != "" {
			result = result.Where("name like ?", "%"+name+"%")
			result = result.Where("dev_e_ui like ?", "%"+name+"%")
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
	deviceinfo := make([]db.Deviceinfo, 0)
	all := 0
	err = result.Find(&deviceinfo).Count(&all).Error
	if err != nil {
		retError(c, 7, err)
		return
	}
	if all != 0 {
		//查数据
		err = result.Where("del != 1").Limit(limit).Offset(offset).Order("id").Find(&deviceinfo).Error
		if err != nil {
			retError(c, 7, err)
			return
		}
	}
	retSuccess(c, map[string]interface{}{
		"all":  all,
		"data": deviceinfo,
	})
}

// DeviceStatus .
type DeviceStatus struct {
	// 设备ID
	IDs []int
	// 状态
	Status int
}

// UpdateDevicesStatus .更改设备状态
// @Tags devices
// @Summary 更改设备状态(仅管理员以上)
// @Description 更改设备状态
// @Accept  json
// @Produce  json
// @Param 	Authorization 	header 	string 	true "With the bearer started JWT"
// @param deviceStatus body api.DeviceStatus  true "设备状态信息"
// @Success 200 {string} json "{"Error":"Success","Data": object}"
// @Failure  500 {string} json "{"Error":"error","Data": null}"
// @Failure  301 {string} json "{"Error":"Re-login","Data": object}"
// @Router /devices/status [put]
func UpdateDevicesStatus(c *gin.Context) {
	permisson := c.GetInt("permisson")
	uid := c.GetInt("id")
	deviceStatus := DeviceStatus{}
	err := c.ShouldBind(&deviceStatus)
	if err != nil {
		retError(c, 1, err)
		return
	}
	if len(deviceStatus.IDs) == 0 {
		retSuccess(c, "Success")
		return
	}
	if deviceStatus.Status < 1 || deviceStatus.Status > 2 {
		retError(c, 13, nil)
		return
	}
	s := service.GetServer()
	result := s.Table("deviceinfos").Where("id in (?)", deviceStatus.IDs)
	if permisson != 3 {
		result = result.Where("classid = ?", uid)
	}
	result = result.Where("status != 3")
	// 如果是启用 -- 判断是否过期
	if deviceStatus.Status == 1 {
		now := time.Now().Unix()
		result = result.Where("expiredtime = 0 or expiredtime > ?", now)
	}
	err = result.Update("status", deviceStatus.Status).Error
	if err != nil {
		retError(c, 7, err)
		return
	}
	// 删除redis-缓存
	s.DelDeviceListInfo(deviceStatus.IDs)
	retSuccess(c, "Success")
}

// DevicesUser 设备所属权信息
type DevicesUser struct {
	// 设备ID
	IDs []int
	// 子用户ID
	UID int
}

// UpdateDevicesUser .更改设备所属权
// @Tags devices
// @Summary .更改设备所属权(仅管理员以上)
// @Description .更改设备所属权
// @Accept  json
// @Produce  json
// @Param 	Authorization 	header 	string 	true "With the bearer started JWT"
// @param devicesUser body api.DevicesUser  true "设备所属权信息"
// @Success 200 {string} json "{"Error":"Success","Data": object}"
// @Failure  500 {string} json "{"Error":"error","Data": null}"
// @Failure  301 {string} json "{"Error":"Re-login","Data": object}"
// @Router /devices/user [put]
func UpdateDevicesUser(c *gin.Context) {
	permisson := c.GetInt("permisson")
	uid := c.GetInt("id")
	devicesUser := DevicesUser{}
	err := c.ShouldBind(&devicesUser)
	if err != nil {
		retError(c, 1, err)
		return
	}
	if len(devicesUser.IDs) == 0 {
		retSuccess(c, "Success")
		return
	}
	if devicesUser.UID == 0 {
		retError(c, 12, nil)
		return
	}
	s := service.GetServer()
	userinfo, err := s.CheckUserID(devicesUser.UID)
	if err != nil {
		retError(c, 12, err)
		return
	}
	if userinfo.Permisson > 2 {
		retError(c, 12, nil)
		return
	}
	result := s.Table("deviceinfos").Where("id in (?)", devicesUser.IDs)
	if permisson != 3 {
		if userinfo.Ownid != uid {
			retError(c, 20, nil)
			return
		}
		result = result.Where("classid = ?", uid)
	}
	err = result.Update(map[string]interface{}{
		"uid": userinfo.ID, "uname": userinfo.Name,
	}).Error
	if err != nil {
		retError(c, 7, err)
		return
	}
	// 删除redis-缓存
	s.DelDeviceListInfo(devicesUser.IDs)
	retSuccess(c, "Success")
}

// DeviceExpire 设备过期时间信息
type DeviceExpire struct {
	// 设备ID
	IDs []int
	// 过期时间
	Expiredtime int64
}

// UpdateDevicesExpire .更改设备过期时间
// @Tags devices
// @Summary .更改设备过期时间(仅管理员以上)
// @Description .更改设备过期时间
// @Accept  json
// @Produce  json
// @Param 	Authorization 	header 	string 	true "With the bearer started JWT"
// @param deviceExpire body api.DeviceExpire  true "设备过期时间信息"
// @Success 200 {string} json "{"Error":"Success","Data": object}"
// @Failure  500 {string} json "{"Error":"error","Data": null}"
// @Failure  301 {string} json "{"Error":"Re-login","Data": object}"
// @Router /devices/expire [put]
func UpdateDevicesExpire(c *gin.Context) {
	deviceExpire := DeviceExpire{}
	err := c.ShouldBind(&deviceExpire)
	if err != nil {
		retError(c, 1, err)
		return
	}
	if len(deviceExpire.IDs) == 0 {
		retError(c, 12, nil)
		return
	}
	status := 3
	if deviceExpire.Expiredtime == 0 {
		status = 1
	} else {
		now := time.Now().Unix()
		if deviceExpire.Expiredtime > now {
			status = 1
		} else {
			status = 3
		}
	}
	s := service.GetServer()
	db := s.Table("deviceinfos").Where("id in (?)", deviceExpire.IDs)
	if c.GetInt("permisson") != 3 {
		db = db.Where("classid = ?", c.GetInt("id"))
	}
	err = db.Updates(map[string]interface{}{
		"expiredtime": deviceExpire.Expiredtime, "status": status,
	}).Error
	if err != nil {
		retError(c, 7, err)
		return
	}
	// 删除redis-缓存
	s.DelDeviceListInfo(deviceExpire.IDs)
	retSuccess(c, "Success")
}

// DevicesID 设备ID
type DevicesID struct {
	// 设备ID
	IDs []int
}

// DevicesDel .删除设备
// @Tags devices
// @Summary .删除设备(仅管理员以上)
// @Description .删除设备
// @Accept  json
// @Produce  json
// @Param 	Authorization 	header 	string 	true "With the bearer started JWT"
// @param devicesID body api.DevicesID  true "设备ID"
// @Success 200 {string} json "{"Error":"Success","Data": object}"
// @Failure  500 {string} json "{"Error":"error","Data": null}"
// @Failure  301 {string} json "{"Error":"Re-login","Data": object}"
// @Router /devices [delete]
func DevicesDel(c *gin.Context) {
	devicesID := DevicesID{}
	err := c.ShouldBind(&devicesID)
	if err != nil {
		retError(c, 1, err)
		return
	}
	if len(devicesID.IDs) == 0 {
		retError(c, 12, nil)
		return
	}
	s := service.GetServer()
	db := s.Table("deviceinfos").Where("id in (?)", devicesID.IDs)
	if c.GetInt("permisson") != 3 {
		db = db.Where("classid = ?", c.GetInt("id"))
	}
	err = db.Update("del", 1).Error
	if err != nil {
		retError(c, 7, err)
		return
	}
	// 删除redis-缓存
	s.DelDeviceListInfo(devicesID.IDs)
	retSuccess(c, "Success")
}
