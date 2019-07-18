package main

import (
	"log"
	"os"

	"./db"
	_ "./docs"
	"./web"
	"./web/api/service"
	"github.com/spf13/viper"
)

func main() {

	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("init config fail", err)
		return
	}
	db.SetSQL(os.Stdout)
	// 加载VIP列表
	service.InitVIP()
	web.Run()
}
