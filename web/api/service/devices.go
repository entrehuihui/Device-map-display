package service

import (
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/gomodule/redigo/redis"

	"../../../db"
)

// CheckDevicesInfo .
func (s *Server) CheckDevicesInfo(classid int, deveui string) ([2]int, error) {
	conn := db.GetRedis()
	defer conn.Close()
	valeus, err := redis.String(conn.Do("GET", deviceMarkKey(classid, deveui)))
	if err != nil {
		log.Println("CheckDevicesInfo Redis Error : ", err)
	} else {
		valueList := strings.Split(valeus, ":")
		if len(valueList) == 2 {
			id, _ := strconv.Atoi(valueList[0])
			uid, _ := strconv.Atoi(valueList[1])
			if id != 0 && uid != 0 {
				return [2]int{id, uid}, nil
			}
		}
	}
	deviceinfo, err := s.GetDeviceinfos(classid, deveui)
	return AddDevicesInfo(&deviceinfo), err
}

func deviceMarkKey(classid int, deveui string) string {
	var deviceMark = "deviceMark:"
	return deviceMark + strconv.Itoa(classid) + ":" + deveui
}

// DelDeviceListInfo 批量删除redis key
func (s *Server) DelDeviceListInfo(ids []int) {
	deviceinfos, err := s.GetDeviceinfosByIDs(ids)
	if err != nil {
		log.Println("DelDeviceListInfo Error : ", err)
	}
	for _, deviceinfo := range deviceinfos {
		DelDevicesInfo(deviceinfo.ID, deviceinfo.DevEUI)
	}
}

// DelDevicesInfo .
func DelDevicesInfo(classid int, deveui string) {
	conn := db.GetRedis()
	defer conn.Close()
	_, err := conn.Do("DEL", deviceMarkKey(classid, deveui))
	if err != nil {
		log.Println("DelDevicesInfo Error : ", err)
	}
}

// AddDevicesInfo .
func AddDevicesInfo(deviceinfo *db.Deviceinfo) [2]int {
	conn := db.GetRedis()
	defer conn.Close()
	now := time.Now().Unix()
	uid := deviceinfo.UID
	if deviceinfo.Expiredtime != 0 && deviceinfo.Expiredtime < now || deviceinfo.Status != 1 {
		uid = deviceinfo.Classid
	}
	_, err := conn.Do("SET", deviceMarkKey(deviceinfo.Classid, deviceinfo.DevEUI), strconv.Itoa(deviceinfo.ID)+":"+strconv.Itoa(uid), "EX", 86400)
	if err != nil {
		log.Println("AddDevicesInfo Error : ", err)
	}
	return [2]int{deviceinfo.ID, uid}
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
		AddDevicesInfo(deviceinfo)
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

// GetDeviceinfosByIDs .按id列表获取设备信息
func (s *Server) GetDeviceinfosByIDs(ids []int) ([]db.Deviceinfo, error) {
	deviceinfo := make([]db.Deviceinfo, 0)
	if len(ids) == 0 {
		return deviceinfo, nil
	}
	err := s.Where("id in (?)", ids).Find(&deviceinfo).Error
	return deviceinfo, err
}
