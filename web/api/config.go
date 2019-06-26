package api

import (
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
		retSuccess(c, db.Configuration{
			ID:       uid,
			Message:  1,
			Sound:    1,
			Language: 1,
		})
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
