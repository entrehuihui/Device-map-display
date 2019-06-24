package main

import (
	"log"
	"os"

	"./db"
	_ "./docs"
	"./web"
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

	web.Run()
}
