package api

import (
	"../../db"
	"./service"
	"github.com/gin-gonic/gin"
)

// GetVIP 获取vip列表(超级用户权限)
// @Tags vip
// @Summary 获取vip列表
// @Description 获取vip列表
// @Accept  json
// @Produce  json
// @Param 	Authorization 	header 	string 	true "With the bearer started JWT"
// @Success 200 {string} json "{"Error":"Success","Data": object}"
// @Failure  500 {string} json "{"Error":"error","Data": null}"
// @Failure  301 {string} json "{"Error":"Re-login","Data": object}"
// @Router /vip [get]
func GetVIP(c *gin.Context) {
	permissons := make([]db.Permisson, 0)
	err := service.GetServer().Order("id").Find(&permissons).Error
	if err != nil {
		retError(c, 7, err)
	}
	retSuccess(c, permissons)
}

// Permisson 权限
type Permisson struct {
	// vipID
	ID int
	// 子用户数 最多1000
	Users int
	// 设备数 最多5000
	Devices int
	// 轨迹权限 1启用 2禁止
	Orbit int
	// 围栏权限 1启用 2禁止
	Fence int
	// 围栏报警权限 1启用 2禁止
	FenceAlarm int
	// 实时数据权限 1启用 2禁止
	Real int
	// 定制LOGO权限 1启用 2禁止
	Logo int
	// 定制状态权限 1启用 2禁止
	State int
}

// PutVIP 获取vip列表(超级用户权限)
// @Tags vip
// @Summary 获取vip列表
// @Description 获取vip列表
// @Accept  json
// @Produce  json
// @Param 	Authorization 	header 	string 	true "With the bearer started JWT"
// @param Permisson body api.Permisson  true "vip信息"
// @Success 200 {string} json "{"Error":"Success","Data": object}"
// @Failure  500 {string} json "{"Error":"error","Data": null}"
// @Failure  301 {string} json "{"Error":"Re-login","Data": object}"
// @Router /vip [put]
func PutVIP(c *gin.Context) {
	permisson := Permisson{}
	err := c.ShouldBind(&permisson)
	if err != nil {
		retError(c, 1, err)
		return
	}
	//
	if permisson.Users < 0 || permisson.Users > 1000 {
		retError(c, 1, nil)
		return
	}
	//
	if permisson.Devices < 1 || permisson.Devices > 5000 {
		retError(c, 1, nil)
		return
	}
	//
	if permisson.Orbit < 1 || permisson.Orbit > 2 {
		retError(c, 1, nil)
		return
	}
	//
	if permisson.Fence < 1 || permisson.Fence > 2 {
		retError(c, 1, nil)
		return
	}
	//
	if permisson.FenceAlarm < 1 || permisson.FenceAlarm > 2 {
		retError(c, 1, nil)
		return
	}
	//
	if permisson.Real < 1 || permisson.Real > 2 {
		retError(c, 1, nil)
		return
	}
	//
	if permisson.Logo < 1 || permisson.Logo > 2 {
		retError(c, 1, nil)
		return
	}
	//
	if permisson.State < 1 || permisson.State > 2 {
		retError(c, 1, nil)
		return
	}
	//
	if permisson.ID > 6 || permisson.ID < 1 {
		retError(c, 1, nil)
		return
	}
	//
	err = service.GetServer().Save(&permisson).Error
	if err != nil {
		retError(c, 7, err)
		return
	}
	retSuccess(c, "Success")
}
