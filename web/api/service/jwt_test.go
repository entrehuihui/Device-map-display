package service_test

import (
	"fmt"
	"log"
	"testing"

	"mymap/web/api/service"
	"github.com/spf13/viper"
)

func initConfig() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("init config fail", err)
		return
	}
}

func Test_NewJWTParseJWT(t *testing.T) {
	// initConfig()
	fmt.Println("test Test_NewJWT and ParseJWT!")
	id := 1
	p := 1
	for index := 0; index < 1000000; index++ {
		id = index
		p = index%3 + 1
		tokenString := service.NewJWT(id, p, 1)
		if tokenString == "" {
			log.Fatal("tokenString is error")
		}
		ids, err := service.ParseJWT(tokenString)
		if err != nil {
			log.Fatal(err)
		}
		if ids[0] != id || ids[1] != p {
			panic(fmt.Sprintf("Error: id =%d, p=%d, ret: id =%d, p=%d", id, p, ids[0], ids[1]))
		}
	}
	fmt.Println("Success")
}

func Test_NewJWTParseJWTExp(t *testing.T) {
	// initConfig()
	fmt.Println("test Test_NewJWT and ParseJWT is expired!")
	service.SetExp(2, "")
	id := 1
	permisson := 1
	for index := 1; index < 7; index++ {
		id = index
		permisson = index%3 + 1
		tokenString := service.NewJWT(id, permisson, index)
		if tokenString == "" {
			log.Fatal("tokenString is error")
		}
		// time.Sleep(time.Second * 3)
		ids, err := service.ParseJWT(tokenString)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(ids)
	}
}
