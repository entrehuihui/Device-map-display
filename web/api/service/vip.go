package service

import (
	"log"
	"strconv"

	"github.com/jinzhu/gorm"

	"../../../db"
)

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

}

// 保存进redis
func setVipRedis(permissons []db.Permisson) {
	conn := db.GetRedis()
	defer conn.Close()
	conn.Send("MULTI")
	for _, permisson := range permissons {
		// 缓存协议类型
		conn.Send("SET", VIPKey(permisson.ID), setVipPermisson(permisson))
	}
}

// user 最多1000 Devices最多5000
func setVipPermisson(permisson db.Permisson) int {
	model := permisson.Users
	model = model<<13 | permisson.Devices
	model = model<<1 | permisson.Orbit
	model = model<<1 | permisson.Fence
	model = model<<1 | permisson.FenceAlarm
	model = model<<1 | permisson.Real
	model = model<<1 | permisson.Logo
	model = model<<1 | permisson.State
	model = model << 1
	return model
}

func getVipPermisson(model int) db.Permisson {
	permisson := db.Permisson{}
	permisson.State = model >> 1 & 0x01
	permisson.Logo = model >> 1 & 0x01
	permisson.Real = model >> 1 & 0x01
	permisson.FenceAlarm = model >> 1 & 0x01
	permisson.Fence = model >> 1 & 0x01
	permisson.Orbit = model >> 1 & 0x01
	permisson.Devices = model >> 1 & 0x01fff
	permisson.Users = model >> 13 & 0x07f
	return permisson
}

// VIPKey 返回VIP redis key
func VIPKey(vip int) string {
	return "vipkey:" + strconv.Itoa(vip)
}

// 1100100 1001110001000

func createVIP() []db.Permisson {
	permissons := []db.Permisson{
		db.Permisson{
			ID:         1,
			Users:      0,
			Devices:    1,
			Orbit:      2,
			Fence:      2,
			FenceAlarm: 2,
			Real:       2,
			Logo:       2,
		},
		db.Permisson{
			ID:         2,
			Users:      1,
			Devices:    2,
			Orbit:      1,
			Fence:      2,
			FenceAlarm: 2,
			Real:       2,
			Logo:       2,
		},
		db.Permisson{
			ID:         3,
			Users:      2,
			Devices:    3,
			Orbit:      1,
			Fence:      1,
			FenceAlarm: 2,
			Real:       2,
			Logo:       2,
		},
		db.Permisson{
			ID:         4,
			Users:      3,
			Devices:    4,
			Orbit:      1,
			Fence:      1,
			FenceAlarm: 1,
			Real:       2,
			Logo:       2,
		},
		db.Permisson{
			ID:         5,
			Users:      4,
			Devices:    5,
			Orbit:      1,
			Fence:      1,
			FenceAlarm: 1,
			Real:       1,
			Logo:       2,
		},
		db.Permisson{
			ID:         6,
			Users:      5,
			Devices:    6,
			Orbit:      1,
			Fence:      1,
			FenceAlarm: 1,
			Real:       1,
			Logo:       1,
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
