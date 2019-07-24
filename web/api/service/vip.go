package service

import (
	"errors"
	"fmt"
	"log"
	"strconv"

	"github.com/gomodule/redigo/redis"
	"github.com/jinzhu/gorm"

	"../../../db"
)

// VipUsers 检测用户是否达到上限
func VipUsers(uid, vip int) error {
	permissons, err := GetVip(vip)
	if err != nil {
		return err
	}
	num := 0
	err = GetServer().Table("userinfos").Where("ownid = ? and del = 0", uid).Count(&num).Error
	if err != nil {
		return err
	}
	if num >= permissons.Users {
		return errors.New("The number of users has reached the maximum")
	}
	return nil
}

// VipDevices 检测设备是否达到上限
func VipDevices(uid, vip int) error {
	permissons, err := GetVip(vip)
	if err != nil {
		return err
	}
	num := 0
	err = GetServer().Table("deviceinfos").Where("classid = ? and del = 0", uid).Count(&num).Error
	if err != nil {
		return err
	}
	if num >= permissons.Devices {
		return errors.New("The number of devices has reached the maximum")
	}
	return nil
}

// VipOrbit 检测轨迹权限
func VipOrbit(vip int) error {
	permissons, err := GetVip(vip)
	fmt.Println(permissons)
	if err != nil {
		return err
	}
	if permissons.Orbit != 1 {
		return errors.New("No Orbit permissions")
	}
	return nil
}

// VipFence 检测围栏权限
func VipFence(vip int) error {
	permissons, err := GetVip(vip)
	if err != nil {
		return err
	}
	if permissons.Fence != 1 {
		return errors.New("No fence permissions")
	}
	return nil
}

// VipReal 检测实时数据权限
func VipReal(vip int) error {
	permissons, err := GetVip(vip)
	if err != nil {
		return err
	}
	if permissons.Real != 1 {
		return errors.New("No Real permissions")
	}
	return nil
}

// VipLogo 检测logo权限
func VipLogo(vip int) error {
	permissons, err := GetVip(vip)
	if err != nil {
		return err
	}
	if permissons.Logo != 1 {
		return errors.New("No logo permissions")
	}
	return nil
}

// VipState 检测状态权限
func VipState(vip int) error {
	permissons, err := GetVip(vip)
	if err != nil {
		return err
	}
	if permissons.State != 1 {
		return errors.New("No state permissions")
	}
	return nil
}

// InitVIP 加载vip列表
func InitVIP() {
	permissons := make([]db.Permisson, 0)
	err := GetServer().Find(&permissons).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		log.Fatal(err)
	}
	if err == gorm.ErrRecordNotFound || len(permissons) != 6 {
		permissons = createVIP()
	}
	setVipRedis(permissons)
}

// 保存进redis
func setVipRedis(permissons []db.Permisson) {
	conn := db.GetRedis()
	defer conn.Close()
	conn.Send("MULTI")
	for _, permisson := range permissons {
		// 缓存协议类型
		conn.Send("SET", GetVIPKey(permisson.ID), setVipPermisson(permisson))
		fmt.Println(setVipPermisson(permisson))
	}
	if _, err := conn.Do("EXEC"); err != nil {
		log.Fatal(err)
	}
}

// SetVipRedis 保存进redis
func SetVipRedis(permisson db.Permisson) error {
	conn := db.GetRedis()
	defer conn.Close()
	_, err := conn.Do("SET", GetVIPKey(permisson.ID), setVipPermisson(permisson))
	return err
}

// GetVip ..
func GetVip(vip int) (db.Permisson, error) {
	conn := db.GetRedis()
	defer conn.Close()
	permissons, err := redis.Int(conn.Do("GET", GetVIPKey(vip)))
	if err != nil {
		return db.Permisson{}, err
	}
	return getVipPermisson(permissons), nil
}

// user 最多1000 Devices最多5000
func setVipPermisson(permisson db.Permisson) int {
	model := permisson.Users
	model = model<<10 | permisson.Devices
	model = model<<2 | permisson.Orbit
	model = model<<2 | permisson.Fence
	model = model<<2 | permisson.Real
	model = model<<2 | permisson.Logo
	model = model<<2 | permisson.State
	return model
}

func getVipPermisson(model int) db.Permisson {
	permisson := db.Permisson{}
	permisson.State = model & 0x03
	model = model >> 2
	permisson.Logo = model & 0x03
	model = model >> 2
	permisson.Real = model & 0x03
	model = model >> 2
	permisson.Fence = model & 0x03
	model = model >> 2
	permisson.Orbit = model & 0x03
	model = model >> 2
	permisson.Devices = model & 0x03ff
	model = model >> 10
	permisson.Users = model
	return permisson
}

// GetVIPKey 返回VIP redis key
func GetVIPKey(vip int) string {
	return "vipkey:" + strconv.Itoa(vip)
}

// 1100100 1001110001000

func createVIP() []db.Permisson {
	permissons := []db.Permisson{
		db.Permisson{
			ID:      1,
			Users:   0,
			Devices: 1,
			Orbit:   2,
			Fence:   2,
			Real:    2,
			Logo:    2,
		},
		db.Permisson{
			ID:      2,
			Users:   1,
			Devices: 2,
			Orbit:   1,
			Fence:   2,
			Real:    2,
			Logo:    2,
		},
		db.Permisson{
			ID:      3,
			Users:   2,
			Devices: 3,
			Orbit:   1,
			Fence:   1,
			Real:    2,
			Logo:    2,
		},
		db.Permisson{
			ID:      4,
			Users:   3,
			Devices: 4,
			Orbit:   1,
			Fence:   1,
			Real:    2,
			Logo:    2,
		},
		db.Permisson{
			ID:      5,
			Users:   4,
			Devices: 5,
			Orbit:   1,
			Fence:   1,
			Real:    1,
			Logo:    2,
		},
		db.Permisson{
			ID:      6,
			Users:   5,
			Devices: 6,
			Orbit:   1,
			Fence:   1,
			Real:    1,
			Logo:    1,
		},
	}
	tx := GetServer().Begin()
	err := tx.Error
	if err != nil {
		log.Fatal(err)
	}
	for _, permisson := range permissons {
		err := tx.Create(&permisson).Error
		if err != nil {
			tx.Rollback()
			log.Fatal(err)
		}
	}
	err = tx.Commit().Error
	if err != nil {
		log.Fatal(err)
	}
	return permissons
}
