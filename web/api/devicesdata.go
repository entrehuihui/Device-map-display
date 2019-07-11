package api

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"../../db"
	"./service"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"golang.org/x/net/websocket"
)

// DeviceInfos 设备信息
type DeviceInfos struct {
	//设备UID
	DevEUI string
	//设备名称 -- 第一次有效
	Name string
	// 经度
	Longitude float64
	// 纬度
	Latitude float64
	// 时间 时间戳(秒)
	Times int64
	// 状态 --自定义
	State int
	// 展示在地图上的信息 json : {speend: 100, azimuth:20...} --
	Info map[string]interface{}
}

// SaveDeviceInfos .
// @Tags devicedata
// @Summary 设备数据推送
// @Description 设备数据推送 -- 只接受管理员账号数据推送
// @Accept  json
// @Produce  json
// @Param 	Authorization 	header 	string 	true "With the bearer started JWT"
// @param userInfo body api.DeviceInfos  true "设备数据信息"
// @Success 200 {string} json "{"Result": object}"
// @Failure  500 {string} json "{"Result": null}"
// @Failure  301 {string} json "{"Error":"Re-login","Data": object}"
// @Router /devicedata [post]
func SaveDeviceInfos(c *gin.Context) {
	permisson := c.GetInt("permisson")
	if permisson != 2 {
		retError(c, 18, nil)
		return
	}
	deviceInfos := DeviceInfos{}
	err := c.ShouldBind(&deviceInfos)
	if err != nil {
		retError(c, 1, err)
		return
	}
	if len(deviceInfos.DevEUI) != 16 {
		retError(c, 22, nil)
		return
	}
	// 判断经纬度
	if deviceInfos.Latitude > 90 || deviceInfos.Latitude < -90 {
		retError(c, 36, nil)
		return
	}
	if deviceInfos.Longitude > 180 || deviceInfos.Longitude < -180 {
		retError(c, 37, nil)
		return
	}
	uid := c.GetInt("id")
	s := service.GetServer()
	// 查询设备是否保存
	infos, err := s.CheckDevicesInfo(uid, deviceInfos.DevEUI)
	if err != nil && err != gorm.ErrRecordNotFound {
		retError(c, 7, err)
		return
	}
	// 设备未添加则创建设备
	if err == gorm.ErrRecordNotFound {
		userInfo, err := s.CheckUserID(uid)
		if err != nil {
			retError(c, 7, err)
			return
		}
		now := time.Now().Unix()
		deviceinfo := db.Deviceinfo{
			Name:       deviceInfos.Name,
			UID:        uid,
			Uname:      userInfo.Name,
			Classid:    uid,
			Calssname:  userInfo.Name,
			DevEUI:     deviceInfos.DevEUI,
			Status:     1,
			Createtime: now,
			Updatetime: 0,
			Expiretime: 0,
		}
		err = s.CreateDevice(&deviceinfo)
		if err != nil {
			retError(c, 7, err)
			return
		}
		infos = [2]int{deviceinfo.ID, uid}
	}
	// 检测是否超出围栏
	if service.ChenckOutFence(infos[1], deviceInfos.Latitude, deviceInfos.Longitude) {
		deviceInfos.Info["fence"] = "overstep"
		deviceInfos.State = 7
	}
	// 保存数据
	jsons, err := json.Marshal(deviceInfos.Info)
	if err != nil {
		retError(c, 1, err)
		return
	}
	if jsons == nil {
		jsons = []byte("{}")
	}
	devicedata := &db.Devicedata{
		Did:        infos[0],
		UID:        infos[1],
		Classid:    uid,
		Longitude:  deviceInfos.Longitude,
		Latitude:   deviceInfos.Latitude,
		Createtime: time.Now().Unix(),
		State:      deviceInfos.State,
		Uptime:     deviceInfos.Times,
		Infos:      string(jsons),
	}
	err = s.SaveDeviceData(devicedata)
	if err != nil {
		retError(c, 7, err)
		return
	}
	datas, err := json.Marshal(devicedata)
	if err != nil {
		log.Printf("SaveDeviceInfos  Marshal error! [error]%s  [data]%v", err, devicedata)
		retSuccess(c, "Success")
		return
	}
	data := string(datas)
	go service.GroupWrite(uid, data)
	if uid != infos[1] {
		go service.GroupWrite(infos[1], data)
	}
	retSuccess(c, "Success")
}

//GRPC接口  --带完成

// GetDevicesDatas 获取数据列表
// @Tags devicedata
// @Summary 获取数据列表
// @Description 获取数据列表
// @Accept  json
// @Produce  json
// @Param 	Authorization 	header 	string 	true "With the bearer started JWT"
// @Param	offset	query	string	flase	"起始数"
// @Param	limit	query	string	flase	"获取数"
// @Param	status	query	string	flase	"数据状态(自定义状态)"
// @Param	starttime	query	string	flase	"数据起始时间"
// @Param	endtime	query	string	flase	"数据终止时间"
// @Param	id		query	int	flase	"设备id"
// @Success 200 {string} json "{"Error":"Success","Data": object}"
// @Failure  500 {string} json "{"Error":"error","Data": null}"
// @Failure  301 {string} json "{"Error":"Re-login","Data": object}"
// @Router /devicedata [get]
func GetDevicesDatas(c *gin.Context) {
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
		result = result.Where("did  = ?", id)
	}
	endtime, err := getEndtime(c)
	if err != nil {
		retError(c, 15, err)
		return
	}
	if endtime != 0 {
		result = result.Where("uptime < ?", endtime)
	}
	//
	starttime, err := getStarttime(c)
	if err != nil {
		retError(c, 14, err)
		return
	}
	if starttime != 0 {
		result = result.Where("uptime > ?", starttime)
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
	// 查数量
	devicedata := make([]db.Devicedata, 0)
	all := 0
	err = result.Find(&devicedata).Count(&all).Error
	if err != nil {
		retError(c, 7, err)
		return
	}
	if all != 0 {
		//查数据
		err = result.Limit(limit).Offset(offset).Order("id desc").Find(&devicedata).Error
		if err != nil {
			retError(c, 7, err)
			return
		}
	}
	retSuccess(c, map[string]interface{}{
		"all":  all,
		"data": devicedata,
	})
}

// WebsocketListen .
// @Tags  websocket
// @Summary websocket连接 /ws?maxiiot_user=token&roomID=roomID
// @Description websocket连接 /ws?maxiiot_user=token&roomID=roomID
// @Accept  json
// @Produce  json
// @Param   Authorization   query  string   true  "jwt"
// @Router /ws?Authorization=jwt [get]
func WebsocketListen(c *gin.Context) {
	// CHECK AUTH HEADER HERE
	Upgrade := c.GetHeader("Upgrade")
	Origin := c.GetHeader("Origin")
	if Upgrade == "" || Origin == "" {
		c.JSON(http.StatusInternalServerError, 7)
		return
	}
	handler := websocket.Handler(service.Echo)
	handler.ServeHTTP(c.Writer, c.Request)
}
