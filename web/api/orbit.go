package api

import (
	"encoding/json"
	"strconv"
	"time"

	"github.com/jinzhu/gorm"

	"../../db"
	"./service"
	"github.com/gin-gonic/gin"
)

// GetOrbit 获取用户围栏
// @Tags orbit
// @Summary 获取用户围栏
// @Description 获取用户围栏
// @Accept  json
// @Produce  json
// @Param 	Authorization 	header 	string 	true "With the bearer started JWT"
// @Param	id	query	int	flase	"用户ID"
// @Success 200 {string} json "{"Error":"Success","Data": object}"
// @Failure  500 {string} json "{"Error":"error","Data": null}"
// @Failure  301 {string} json "{"Error":"Re-login","Data": object}"
// @Router /orbit [get]
func GetOrbit(c *gin.Context) {
	ids := c.Query("id")
	id := 0
	uid := c.GetInt("id")
	var err error
	if ids == "" {
		id = uid
	} else {
		id, err = strconv.Atoi(ids)
		if err != nil {
			retError(c, 12, err)
			return
		}
		if id == 0 {
			retError(c, 12, nil)
			return
		}
	}
	s := service.GetServer()
	permisson := c.GetInt("permisson")
	if id != uid && permisson != 3 {
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
	}
	orbitlists := db.Orbitlists{}
	err = s.Where("uid = ? and del = 0", id).Last(&orbitlists).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		retError(c, 7, err)
		return
	}
	retSuccess(c, orbitlists)
}

// Orbitinfo 用户围栏信息
type Orbitinfo struct {
	// 1 圆 2方
	Types int
	// 围栏信息json  1: {Lat:21.32, Lng:113.21, radius: 123}  2:[{Lat:21.32, Lng:113.21}, {Lat:21.32, Lng:113.21}....]
	Info interface{}
}

// Orbitinfo1 圆围栏
type Orbitinfo1 struct {
	Lat    float64
	Lng    float64
	Radius int
}

// Orbitinfo2 方围栏
type Orbitinfo2 struct {
	Lat float64
	Lng float64
}

// PostOrbit .设置用户围栏
// @Tags orbit
// @Summary .设置用户围栏
// @Description .设置用户围栏
// @Accept  json
// @Produce  json
// @Param 	Authorization 	header 	string 	true "With the bearer started JWT"
// @param configuration body api.Orbitinfo  true "用户围栏信息"
// @Success 200 {string} json "{"Error":"Success","Data": object}"
// @Failure  500 {string} json "{"Error":"error","Data": null}"
// @Failure  301 {string} json "{"Error":"Re-login","Data": object}"
// @Router /orbit [post]
func PostOrbit(c *gin.Context) {
	orbitinfo := Orbitinfo{}
	err := c.ShouldBind(&orbitinfo)
	if err != nil {
		retError(c, 7, err)
		return
	}
	buf, err := json.Marshal(orbitinfo.Info)
	if err != nil {
		retError(c, 7, err)
		return
	}
	uid := c.GetInt("id")
	orbitlists := db.Orbitlists{
		UID:        uid,
		Types:      orbitinfo.Types,
		Createtime: time.Now().Unix(),
		Del:        0,
	}
	if orbitinfo.Types == 1 {
		orbitinfo1 := Orbitinfo1{}
		err = json.Unmarshal(buf, &orbitinfo1)
		if err != nil {
			retError(c, 7, err)
			return
		}
		if orbitinfo1.Radius <= 0 {
			retError(c, 1, nil)
			return
		}
		if orbitinfo1.Lat > 90 || orbitinfo1.Lat < -90 {
			retError(c, 36, nil)
			return
		}
		if orbitinfo1.Lng > 180 || orbitinfo1.Lng < -180 {
			retError(c, 37, nil)
			return
		}
		buf, err = json.Marshal(orbitinfo1)
		if err != nil {
			retError(c, 7, err)
			return
		}
		orbitlists.Info = string(buf)
	} else if orbitinfo.Types == 2 {
		orbitinfo2 := make([]Orbitinfo2, 0)
		err = json.Unmarshal(buf, &orbitinfo2)
		if err != nil {
			retError(c, 7, err)
			return
		}
		if len(orbitinfo2) < 3 {
			retError(c, 1, nil)
			return
		}
		for _, v := range orbitinfo2 {
			if v.Lat > 90 || v.Lat < -90 {
				retError(c, 36, nil)
				return
			}
			if v.Lng > 180 || v.Lng < -180 {
				retError(c, 37, nil)
				return
			}
		}
		buf, err = json.Marshal(orbitinfo2)
		if err != nil {
			retError(c, 7, err)
			return
		}
		orbitlists.Info = string(buf)
	} else {
		retError(c, 1, err)
		return
	}

	tx := service.GetServer().Begin()
	err = tx.Model(&db.Orbitlists{}).Where("uid = ? and del = 0", uid).Update("del", 1).Error
	if err != nil {
		tx.Rollback()
		retError(c, 7, err)
		return
	}
	err = tx.Create(&orbitlists).Error
	if err != nil {
		tx.Rollback()
		retError(c, 7, err)
		return
	}
	tx.Commit()
	retSuccess(c, orbitlists)
}

// OrbitID 围栏ID
type OrbitID struct {
	// 围栏ID
	ID int
}

// DelOrbit .删除用户围栏
// @Tags orbit
// @Summary .删除用户围栏
// @Description .删除用户围栏
// @Accept  json
// @Produce  json
// @Param 	Authorization 	header 	string 	true "With the bearer started JWT"
// @param configuration body api.Orbitinfo  true "用户围栏信息"
// @Success 200 {string} json "{"Error":"Success","Data": object}"
// @Failure  500 {string} json "{"Error":"error","Data": null}"
// @Failure  301 {string} json "{"Error":"Re-login","Data": object}"
// @Router /orbit [delete]
func DelOrbit(c *gin.Context) {
	orbitID := OrbitID{}
	err := c.ShouldBind(&orbitID)
	if err != nil {
		retError(c, 7, err)
		return
	}
	if orbitID.ID == 0 {
		retError(c, 12, nil)
		return
	}
	s := service.GetServer()
	err = s.Exec("update orbitlists set del = 1 where id = ? and uid = ?", orbitID.ID, c.GetInt("id")).Error
	if err != nil {
		retError(c, 7, err)
		return
	}
	retSuccess(c, "Success")
}
