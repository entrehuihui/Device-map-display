package service

import (
	"log"
	"time"

	"../../../db"
	"github.com/jinzhu/gorm"
)

// 设备ID 设备State
var devices map[int]map[string][2]int

// InitDevices 初始化设备列表
func InitDevices() {
	s := db.GetDB()
	now := time.Now().AddDate(0, 0, -2).Unix()
	deviceinfo := make([]db.Deviceinfo, 0)
	err := s.Raw("SELECT * FROM deviceinfos WHERE id IN ( SELECT id FROM devicedata WHERE uptime > ? ) AND del = 0", now).Scan(&deviceinfo).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		log.Fatal(err)
	}

	devices = make(map[int]map[string][2]int, 0)
	for _, v := range deviceinfo {
		if devices[v.Classid] == nil {
			devices[v.Classid] = make(map[string][2]int, 0)
		}
		devices[v.Classid][v.DevEUI] = [2]int{v.ID, v.UID}
	}
}

// CheckDevicesInfo .
func (s *Server) CheckDevicesInfo(classid int, deveui string) ([2]int, error) {
	if devices[classid][deveui][0] != 0 {
		return devices[classid][deveui], nil
	}
	deviceinfo, err := s.GetDeviceinfos(classid, deveui)
	return [2]int{deviceinfo.ID, deviceinfo.UID}, err
}

// DelDevicesInfo .
func DelDevicesInfo(classid int, deveui string) {
	delete(devices[classid], deveui)
}

// AddDevicesInfo .
func AddDevicesInfo(deveui string, infos [2]int) {
	if devices[infos[0]] == nil {
		devices[infos[0]] = make(map[string][2]int, 0)
	}
	devices[infos[0]][deveui] = infos
}

// GetDeviceinfos .按classid和deveui获取设备信息
func (s *Server) GetDeviceinfos(classid int, deveui string) (db.Deviceinfo, error) {
	deviceinfo := db.Deviceinfo{}
	err := s.Where("classid = ? and dev_e_ui = ? and del != 1", classid, deveui).Find(&deviceinfo).Error
	return deviceinfo, err
}

// CreateDevice 创建设备
func (s *Server) CreateDevice(deviceinfo *db.Deviceinfo) error {
	var err error
	err = s.Create(deviceinfo).Error
	if err == nil {
		AddDevicesInfo(deviceinfo.DevEUI, [2]int{deviceinfo.ID, deviceinfo.UID})
	}
	return err
}

// UpdateDevice 更新设备
func (s *Server) UpdateDevice(id int, updates map[string]interface{}) (db.Deviceinfo, error) {
	deviceinfo := db.Deviceinfo{
		ID: id,
	}
	err := s.Model(&deviceinfo).Updates(updates).Error
	return deviceinfo, err
}

// SaveDeviceData 保存数据
func (s *Server) SaveDeviceData(devicedata *db.Devicedata) error {
	var err error
	err = s.Create(devicedata).Error
	return err
}

// GetDeviceinfosByID .按id获取设备信息
func (s *Server) GetDeviceinfosByID(id int) (db.Deviceinfo, error) {
	deviceinfo := db.Deviceinfo{
		ID: id,
	}
	err := s.Where("del != 1").Find(&deviceinfo).Error
	return deviceinfo, err
}
