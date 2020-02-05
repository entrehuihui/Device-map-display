package api

import (
	"strconv"
	"strings"
	"time"

	"github.com/jinzhu/gorm"

	"mymap/db"
	"mymap/web/api/service"
	"github.com/gin-gonic/gin"
)

// GetLogin 获取背景图列表....
// @Tags logo
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
	err := s.Raw("SELECT * FROM logoinfos WHERE id IN (SELECT max(id) FROM logoinfos WHERE uid = 1 and del = 0 GROUP BY types)").Scan(&logoinfos).Error
	if err != nil {
		retError(c, 7, err)
		return
	}
	//
	maps := make(map[int]string, 0)
	maps[1] = "/images/logo/beijing.jpg"
	maps[2] = "/images/logo/beijinglogo.jpg"
	for _, v := range logoinfos {
		maps[v.Types] = v.URL
	}
	retSuccess(c, maps)
}

// GetLogo 获取logo....
// @Tags logo
// @Summary 获取logo
// @Description 获取logo
// @Accept  json
// @Produce  json
// @Param 	Authorization 	header 	string 	true "With the bearer started JWT"
// @Success 200 {string} json "{"Error":"Success","Data": object}"
// @Failure  500 {string} json "{"Error":"error","Data": null}"
// @Failure  301 {string} json "{"Error":"Re-login","Data": object}"
// @Router /login/logo/3 [get]
func GetLogo(c *gin.Context) {
	logoinfo := db.Logoinfo{}
	s := service.GetServer()
	uid := c.GetInt("id")
	err := service.VipLogo(c.GetInt("vip"))
	if err == nil {
		if c.GetInt("permisson") == 1 {
			info, err := s.CheckUserID(uid)
			if err != nil {
				retError(c, 7, err)
				return
			}
			uid = info.Ownid
		}
		err = s.Raw("select * from logoinfos where uid = ? and types = 3 and del = 0 order by id desc limit 1", uid).Scan(&logoinfo).Error
		if err != nil && err != gorm.ErrRecordNotFound {
			retError(c, 7, err)
			return
		}
	}
	if err != nil {
		err = s.Raw("select * from logoinfos where uid = 1 and types = 3 and del = 0 order by id desc limit 1").Scan(&logoinfo).Error
		if err != nil && err != gorm.ErrRecordNotFound {
			retError(c, 7, err)
			return
		}
		if err == gorm.ErrRecordNotFound {
			logoinfo.URL = "/images/logo/shanbiaologo.jpg"
		}
	}
	retSuccess(c, logoinfo.URL)
}

// Logo .图片信息
type Logo struct {
	// 图片类型 1登陆背景 2登陆界面 3logo
	Types int
	// 图片地址
	Photo string
}

// PutLogin 修改背景图列表....
// @Tags logo
// @Summary 修改背景图列表
// @Description 修改背景图列表
// @Accept  json
// @Produce  json
// @Param 	Authorization 	header 	string 	true "With the bearer started JWT"
// @param Logo body api.Logo  true "图片信息"
// @Success 200 {string} json "{"Error":"Success","Data": object}"
// @Failure  500 {string} json "{"Error":"error","Data": null}"
// @Failure  301 {string} json "{"Error":"Re-login","Data": object}"
// @Router /login/logo [put]
func PutLogin(c *gin.Context) {
	if c.GetInt("permisson") != 3 {
		err := service.VipLogo(c.GetInt("vip"))
		if err != nil {
			retError(c, 45, err)
			return
		}
	}
	logo := Logo{}
	err := c.ShouldBind(&logo)
	if err != nil {
		retError(c, 1, err)
		return
	}
	if logo.Types > 3 || logo.Types < 1 {
		retError(c, 1, err)
		return
	}
	if logo.Photo == "" {
		retError(c, 1, err)
		return
	}
	if c.GetInt("permisson") != 3 && logo.Types != 3 {
		retError(c, 1, err)
		return
	}
	uid := c.GetInt("id")
	newname := ""
	if logo.Photo != "" {
		filename := strings.Split(logo.Photo, ".")
		if len(filename) != 2 {
			retError(c, 31, nil)
			return
		}
		if filename[1] != "jpg" && filename[1] != "png" {
			retError(c, 30, nil)
			return
		}
		newname = "/images/logo/" + strconv.FormatInt(time.Now().Unix(), 10) + getRandomString(4) + strconv.Itoa(uid) + "." + filename[1]
		types := logo.Types + 1
		if err := prictureSize(logo.Photo, newname, types); err != nil {
			retError(c, 31, err)
			return
		}
	}
	s := service.GetServer()
	userinfo, err := s.CheckUserID(uid)
	if err != nil {
		removefile(newname, 1)
		retError(c, 7, err)
		return
	}
	logoinfo := db.Logoinfo{
		UID:        uid,
		Uname:      userinfo.Name,
		Types:      logo.Types,
		URL:        newname,
		Createtime: time.Now().Unix(),
		Del:        0,
	}
	tx := s.Begin()
	// 删除就图片
	err = tx.Table("logoinfos").Where("uid = ? and types = ?", uid, logo.Types).Update("del", 1).Error
	if err != nil {
		tx.Rollback()
		removefile(newname, 1)
		retError(c, 7, err)
		return
	}
	// 创建新图片
	err = tx.Create(&logoinfo).Error
	if err != nil {
		tx.Rollback()
		removefile(newname, 1)
		retError(c, 7, err)
		return
	}
	err = tx.Commit().Error
	if err != nil {
		removefile(newname, 1)
		retError(c, 7, err)
		return
	}
	retSuccess(c, "Success")
}

// DelLogo 图片ID
type DelLogo struct {
	ID int
}

// DelLogin 删除背景图列表....
// @Tags logo
// @Summary 删除背景图列表
// @Description 删除背景图列表
// @Accept  json
// @Produce  json
// @Param 	Authorization 	header 	string 	true "With the bearer started JWT"
// @param DelLogo body api.DelLogo  true "图片ID"
// @Success 200 {string} json "{"Error":"Success","Data": object}"
// @Failure  500 {string} json "{"Error":"error","Data": null}"
// @Failure  301 {string} json "{"Error":"Re-login","Data": object}"
// @Router /login/logo [delete]
func DelLogin(c *gin.Context) {
	delLogo := DelLogo{}
	err := c.ShouldBind(&delLogo)
	if err != nil || delLogo.ID == 0 {
		retError(c, 1, err)
		return
	}
	conn := service.GetServer().Table("logoinfos").Where("id = ?", delLogo.ID)
	if c.GetInt("permisson") != 3 {
		conn = conn.Where("uid = ?", c.GetInt("id"))
	}
	err = conn.Update("del", 1).Error
	if err != nil {
		retError(c, 7, err)
		return
	}
	retSuccess(c, "Success")
}

// GetUserLogo 获取用户LOGO图列表....
// @Tags logo
// @Summary 获取用户LOGO图列表
// @Description 获取用户LOGO图列表
// @Accept  json
// @Produce  json
// @Param 	Authorization 	header 	string 	true "With the bearer started JWT"
// @Param	id	query	int	flase	"用户ID"
// @Param	offset	query	string	flase	"起始数"
// @Param	limit	query	string	flase	"获取数"
// @Success 200 {string} json "{"Error":"Success","Data": object}"
// @Failure  500 {string} json "{"Error":"error","Data": null}"
// @Failure  301 {string} json "{"Error":"Re-login","Data": object}"
// @Router /login/logo/user [get]
func GetUserLogo(c *gin.Context) {
	id, err := getID(c)
	if err != nil {
		retError(c, 12, err)
		return
	}
	limit, err := getLimit(c)
	if err != nil {
		retError(c, 17, err)
		return
	}
	offset, err := getOffset(c)
	if err != nil {
		retError(c, 17, err)
		return
	}
	conn := service.GetServer().Table("logoinfos").Where("id in (SELECT MAX(id) from logoinfos WHERE del = 0 AND uid !=1 GROUP BY uid)")
	if id != 0 {
		conn = conn.Where("uid = ?", id)
	}

	//  查询数量
	all := 0
	err = conn.Count(&all).Error
	if err != nil {
		retError(c, 7, err)
		return
	}
	if all == 0 {
		retSuccess(c, map[string]interface{}{
			"all":  all,
			"data": []int{},
		})
	}

	logoinfos := make([]db.Logoinfo, 0)
	err = conn.Count(&all).Order("id").Limit(limit).Offset(offset).Find(&logoinfos).Error
	if err != nil {
		retError(c, 7, err)
		return
	}
	retSuccess(c, map[string]interface{}{
		"all":  all,
		"data": logoinfos,
	})
}
