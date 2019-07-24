package db

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq" //
	"github.com/spf13/viper"
)

// DB 连接DB
var db *gorm.DB

//SetSQL 初始化postgresql连接
func SetSQL(f io.Writer) (err error) {
	// dbparameter := "host=120.78.76.139 port=5432 user=mapPlatform dbname=mp2 password=mapPlatform sslmode=disable"
	dbparameter := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", viper.GetString("Postgresql.host"), viper.GetString("Postgresql.port"), viper.GetString("Postgresql.username"), viper.GetString("Postgresql.dbname"), viper.GetString("Postgresql.password"))

loop:
	db, err = gorm.Open("postgres", dbparameter)
	if err != nil {
		log.Println("postgresql db isn't connect error :", err)
		time.Sleep(3 * time.Second)
		goto loop
	}
	db.LogMode(true)
	db.SetLogger(log.New(f, "\r\n", 0))
	// 设置连接池为100
	db.DB().SetMaxIdleConns(100)
	// 迁移数据
	err = db.AutoMigrate(
		&Userinfo{},
		&Logininfo{},
		&Deviceinfo{},
		&Devicedata{},
		&Logoinfo{},
		&Configuration{},
		&Configstates{},
		&Fencelists{},
		&Permisson{},
	).Error
	if err != nil {
		log.Fatal(err)
	}
	// 初始化redis
	initPoll()
	createSuperAdmin()
	return
}

// GetDB 获取DB
func GetDB() *gorm.DB {
	return db
}

// MD5 密码加密
func MD5(password string) string {
	//密码MD5加密
	p5 := md5.New()
	p5.Write([]byte(password))
	return hex.EncodeToString(p5.Sum(nil))
}

//设置超级用户
func createSuperAdmin() {
	num := 0
	err := GetDB().Model(&Userinfo{}).Where("id = 1 and permisson = 3").Count(&num).Error
	if err != nil {
		log.Fatal(err)
	}
	if num != 0 {
		return
	}
	err = GetDB().Create(&Userinfo{
		ID:          1,
		Ownid:       0,
		Name:        viper.GetString("Admin.name"),
		Password:    MD5(viper.GetString("Admin.password")),
		Photo:       "/images/user/defaultuser.png",
		Permisson:   3,
		VIP:         6,
		Status:      1,
		Createtime:  0,
		Updatetime:  0,
		Expiredtime: 0,
		Mobile:      "",
		Email:       "",
		Address:     "",
		Details:     "",
		Del:         0,
	}).Error
	if err != nil {
		log.Fatal(err)
	}
}
