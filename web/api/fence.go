package api

import (
	"encoding/json"
	"time"

	"github.com/jinzhu/gorm"

	"../../db"
	"./service"
	"github.com/gin-gonic/gin"
)

// GetFence 获取用户围栏
// @Tags fence
// @Summary 获取用户围栏
// @Description 获取用户围栏
// @Accept  json
// @Produce  json
// @Param 	Authorization 	header 	string 	true "With the bearer started JWT"
// @Param	id	query	int	flase	"用户ID"
// @Success 200 {string} json "{"Error":"Success","Data": object}"
// @Failure  500 {string} json "{"Error":"error","Data": null}"
// @Failure  301 {string} json "{"Error":"Re-login","Data": object}"
// @Router /fence [get]
func GetFence(c *gin.Context) {
	id, err := getID(c)
	if err != nil {
		retError(c, 12, err)
		return
	}
	uid := c.GetInt("id")
	if id == 0 {
		id = uid
	}
	s := service.GetServer()
	permisson := c.GetInt("permisson")
	if permisson != 3 {
		vip := c.GetInt("vip")
		if id != uid {
			// 检测是否是子账户
			user, err := s.CheckUserID(id)
			if err != nil {
				retError(c, 12, err)
				return
			}
			if user.Ownid != uid {
				retError(c, 12, nil)
				return
			}
			vip = user.VIP
		}
		// 检测是否有围栏权限
		err = service.VipFence(vip)
		if err != nil {
			retError(c, 43, err)
			return
		}
	}

	fencelists := db.Fencelists{}
	err = s.Where("uid = ? and del = 0", id).Last(&fencelists).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		retError(c, 7, err)
		return
	}
	retSuccess(c, fencelists)
}

// Fenceinfo 用户围栏信息
type Fenceinfo struct {
	// 1 圆 2方
	Types int
	// 围栏信息json  1: {Lat:21.32, Lng:113.21, radius: 123}  2:[{Lat:21.32, Lng:113.21}, {Lat:21.32, Lng:113.21}....]
	Info interface{}
}

// Fenceinfo1 圆围栏
type Fenceinfo1 struct {
	Lat    float64
	Lng    float64
	Radius int
}

// Fenceinfo2 方围栏
type Fenceinfo2 struct {
	Lat float64
	Lng float64
}

// PostFence .设置用户围栏
// @Tags fence
// @Summary .设置用户围栏
// @Description .设置用户围栏
// @Accept  json
// @Produce  json
// @Param 	Authorization 	header 	string 	true "With the bearer started JWT"
// @param 	fenceinfo body api.Fenceinfo  true "用户围栏信息"
// @Success 200 {string} json "{"Error":"Success","Data": object}"
// @Failure  500 {string} json "{"Error":"error","Data": null}"
// @Failure  301 {string} json "{"Error":"Re-login","Data": object}"
// @Router /fence [post]
func PostFence(c *gin.Context) {
	// 检测是否有围栏权限
	err := service.VipFence(c.GetInt("vip"))
	if err != nil {
		retError(c, 43, err)
		return
	}
	fenceinfo := Fenceinfo{}
	err = c.ShouldBind(&fenceinfo)
	if err != nil {
		retError(c, 7, err)
		return
	}
	buf, err := json.Marshal(fenceinfo.Info)
	if err != nil {
		retError(c, 7, err)
		return
	}
	uid := c.GetInt("id")
	fencelists := db.Fencelists{
		UID:        uid,
		Types:      fenceinfo.Types,
		Createtime: time.Now().Unix(),
		Del:        0,
	}
	if fenceinfo.Types == 1 {
		fenceinfo1 := Fenceinfo1{}
		err = json.Unmarshal(buf, &fenceinfo1)
		if err != nil {
			retError(c, 7, err)
			return
		}
		if fenceinfo1.Radius <= 0 {
			retError(c, 1, nil)
			return
		}
		if fenceinfo1.Lat > 90 || fenceinfo1.Lat < -90 {
			retError(c, 36, nil)
			return
		}
		if fenceinfo1.Lng > 180 || fenceinfo1.Lng < -180 {
			retError(c, 37, nil)
			return
		}
		buf, err = json.Marshal(fenceinfo1)
		if err != nil {
			retError(c, 7, err)
			return
		}
		fencelists.Info = string(buf)
	} else if fenceinfo.Types == 2 {
		fenceinfo2 := make([]Fenceinfo2, 0)
		err = json.Unmarshal(buf, &fenceinfo2)
		if err != nil {
			retError(c, 7, err)
			return
		}
		if len(fenceinfo2) < 3 {
			retError(c, 1, nil)
			return
		}
		for _, v := range fenceinfo2 {
			if v.Lat > 90 || v.Lat < -90 {
				retError(c, 36, nil)
				return
			}
			if v.Lng > 180 || v.Lng < -180 {
				retError(c, 37, nil)
				return
			}
		}
		buf, err = json.Marshal(fenceinfo2)
		if err != nil {
			retError(c, 7, err)
			return
		}
		fencelists.Info = string(buf)
	} else {
		retError(c, 1, err)
		return
	}

	tx := service.GetServer().Begin()
	err = tx.Model(&db.Fencelists{}).Where("uid = ? and del = 0", uid).Update("del", 1).Error
	if err != nil {
		tx.Rollback()
		retError(c, 7, err)
		return
	}
	err = tx.Create(&fencelists).Error
	if err != nil {
		tx.Rollback()
		retError(c, 7, err)
		return
	}
	err = service.DelFence(uid)
	if err != nil {
		tx.Rollback()
		retError(c, 7, err)
		return
	}
	service.AddFence(&fencelists)
	tx.Commit()
	retSuccess(c, fencelists)
}

// FencetID 围栏ID
type FencetID struct {
	// 围栏ID  --传0则为删除自己的围栏
	ID int
}

// DelFence .删除用户围栏
// @Tags fence
// @Summary .删除用户围栏
// @Description .删除用户围栏
// @Accept  json
// @Produce  json
// @Param 	Authorization 	header 	string 	true "With the bearer started JWT"
// @param fencetID body api.FencetID  true "用户围栏信息"
// @Success 200 {string} json "{"Error":"Success","Data": object}"
// @Failure  500 {string} json "{"Error":"error","Data": null}"
// @Failure  301 {string} json "{"Error":"Re-login","Data": object}"
// @Router /fence [delete]
func DelFence(c *gin.Context) {
	fenceID := FencetID{}
	err := c.ShouldBind(&fenceID)
	if err != nil {
		retError(c, 7, err)
		return
	}
	tx := service.GetServer().Begin()
	uid := c.GetInt("id")
	if fenceID.ID == 0 {
		err = tx.Exec("update fencelists set del = 1 where uid = ? and del = 0", uid).Error
		if err != nil {
			tx.Rollback()
			retError(c, 7, err)
			return
		}
		retSuccess(c, "Success")
		return
	}
	if c.GetInt("permisson") != 3 {
		s := service.GetServer()
		user, err := s.CheckUserID(fenceID.ID)
		if err != nil {
			retError(c, 7, nil)
			return
		}
		if user.Ownid != uid {
			retError(c, 17, nil)
			return
		}
	}
	err = tx.Exec("update fencelists set del = 1 where id = ? ", fenceID.ID).Error
	if err != nil {
		tx.Rollback()
		retError(c, 7, err)
		return
	}
	err = service.DelFence(uid)
	if err != nil {
		tx.Rollback()
		retError(c, 7, err)
		return
	}
	tx.Commit()
	retSuccess(c, "Success")
}
