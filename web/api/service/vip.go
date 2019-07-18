package service

import (
	"log"

	"github.com/jinzhu/gorm"

	"../../../db"
)

// InitVIP 加载vip列表
func InitVIP() {
	permisson := db.Permisson{}
	err := GetServer().First(&permisson).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		log.Fatal(err)
	}
	if err == gorm.ErrRecordNotFound {
		createVIP()
	}
}

func createVIP() {
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
}
