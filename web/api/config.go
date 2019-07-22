package api

import (
	"time"

	"../../db"
	"./service"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// GetConfiguration 获取用户配置
// @Tags config
// @Summary 获取用户配置
// @Description 获取用户配置
// @Accept  json
// @Produce  json
// @Param 	Authorization 	header 	string 	true "With the bearer started JWT"
// @Success 200 {string} json "{"Error":"Success","Data": object}"
// @Failure  500 {string} json "{"Error":"error","Data": null}"
// @Failure  301 {string} json "{"Error":"Re-login","Data": object}"
// @Router /config [get]
func GetConfiguration(c *gin.Context) {
	uid := c.GetInt("id")
	s := service.GetServer()
	donfiguration := db.Configuration{
		ID: uid,
	}
	err := s.First(&donfiguration).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		retError(c, 7, err)
		return
	}
	if err == gorm.ErrRecordNotFound {
		config := db.Configuration{
			ID:       uid,
			Message:  1,
			Sound:    1,
			Language: 1,
		}
		s.Create(&config)
		retSuccess(c, config)
		return
	}
	retSuccess(c, donfiguration)
}

// Configuration 用户配置
type Configuration struct {
	//用户ID --不填
	ID int
	//消息推送开关 1开2关 不更改不填或填0
	Message int
	//声音开关 1开2关 不更改不填或填0
	Sound int
	//语言选择 1简体中文 2繁体中文 3英文 不更改不填或填0
	Language int
}

// UpdateConfiguration .更改用户配置
// @Tags config
// @Summary .更改用户配置
// @Description .更改用户配置
// @Accept  json
// @Produce  json
// @Param 	Authorization 	header 	string 	true "With the bearer started JWT"
// @param configuration body api.Configuration  true "用户配置信息"
// @Success 200 {string} json "{"Error":"Success","Data": object}"
// @Failure  500 {string} json "{"Error":"error","Data": null}"
// @Failure  301 {string} json "{"Error":"Re-login","Data": object}"
// @Router /config [put]
func UpdateConfiguration(c *gin.Context) {
	configuration := Configuration{}
	err := c.ShouldBind(&configuration)
	if err != nil {
		retError(c, 1, err)
		return
	}
	configuration.ID = c.GetInt("id")
	s := service.GetServer()
	err = s.Model(&configuration).Updates(configuration).Error
	if err != nil {
		retError(c, 7, err)
		return
	}
	retSuccess(c, "Success")
}

// GetConfigState .获取状态配置
// @Tags config
// @Summary .获取状态配置
// @Description .获取状态配置
// @Accept  json
// @Produce  json
// @Param 	Authorization 	header 	string 	true "With the bearer started JWT"
// @Success 200 {string} json "{"Error":"Success","Data": object}"
// @Failure  500 {string} json "{"Error":"error","Data": null}"
// @Failure  301 {string} json "{"Error":"Re-login","Data": object}"
// @Router /configState [get]
func GetConfigState(c *gin.Context) {
	uid := c.GetInt("id")
	s := service.GetServer()
	permisson := c.GetInt("permisson")
	if permisson < 2 {
		userinfo, err := s.CheckUserID(uid)
		if err != nil {
			retError(c, 7, err)
			return
		}
		uid = userinfo.Ownid
	}
	var err error
	configstates := make([]db.Configstates, 0)
	err = s.Where("uid = ?", uid).Order("id").Find(&configstates).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		retError(c, 7, err)
		return
	}
	if len(configstates) == 0 {
		err = s.CreateState(uid)
		if err != nil && err != gorm.ErrRecordNotFound {
			retError(c, 7, err)
			return
		}
		err = s.Where("uid = ?", uid).Order("id").Find(&configstates).Error
		if err != nil {
			retError(c, 7, err)
			return
		}
	}
	retSuccess(c, configstates)
}

// ConfigState 状态配置信息
type ConfigState struct {
	ID     int
	States string
}

// UpdateConfigState .更改状态配置
// @Tags config
// @Summary .更改状态配置
// @Description .更改状态配置
// @Accept  json
// @Produce  json
// @Param 	Authorization 	header 	string 	true "With the bearer started JWT"
// @param configState body api.ConfigState  true "状态配置信息"
// @Success 200 {string} json "{"Error":"Success","Data": object}"
// @Failure  500 {string} json "{"Error":"error","Data": null}"
// @Failure  301 {string} json "{"Error":"Re-login","Data": object}"
// @Router /configState [put]
func UpdateConfigState(c *gin.Context) {
	configState := ConfigState{}
	err := c.ShouldBind(&configState)
	if err != nil {
		retError(c, 1, err)
		return
	}
	if configState.ID == 0 {
		retError(c, 12, nil)
		return
	}
	if len(configState.States) > 20 || configState.States == "" {
		retError(c, 35, nil)
		return
	}
	configstates := db.Configstates{
		ID:         configState.ID,
		States:     configState.States,
		Updatetime: time.Now().Unix(),
	}
	s := service.GetServer()
	err = s.Model(&configstates).Where("uid = ?", c.GetInt("id")).Updates(configstates).Error
	if err != nil {
		retError(c, 7, err)
		return
	}
	retSuccess(c, "Success")
}
