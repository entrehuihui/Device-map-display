package main

import (
	"log"
	"os"

	"mymap/db"
	_ "mymap/docs"
	"mymap/web"
	"mymap/web/api/service"

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
