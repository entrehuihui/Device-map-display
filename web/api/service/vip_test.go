package service

import (
	"fmt"
	"testing"

	"mymap/db"
)

func Test_vipSetGet(t *testing.T) {
	// initConfig()
	fmt.Println("test Test_vipSetGet")
	permisson := db.Permisson{
		State:   1,
		Logo:    2,
		Real:    1,
		Fence:   2,
		Orbit:   1,
		Devices: 3,
		Users:   1,
	}
	vip := setVipPermisson(permisson)
	permisson1 := getVipPermisson(vip)
	fmt.Println(permisson.State, permisson1.State)
	fmt.Println(permisson.Logo, permisson1.Logo)
	fmt.Println(permisson.Real, permisson1.Real)
	fmt.Println(permisson.Fence, permisson1.Fence)
	fmt.Println(permisson.Orbit, permisson1.Orbit)
	fmt.Println(permisson.Devices, permisson1.Devices)
	fmt.Println(permisson.Users, permisson1.Users)
}
