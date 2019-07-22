package service

import (
	"fmt"
	"testing"

	"../../../db"
)

func Test_vipSetGet(t *testing.T) {
	// initConfig()
	fmt.Println("test Test_vipSetGet")
	permisson := db.Permisson{
		State:      1,
		Logo:       0,
		Real:       1,
		FenceAlarm: 0,
		Fence:      1,
		Orbit:      1,
		Devices:    5000,
		Users:      1000,
	}
	vip := setVipPermisson(permisson)
	permisson1 := getVipPermisson(vip)
	fmt.Println(permisson.State, permisson1.State)
	fmt.Println(permisson.Logo, permisson1.Logo)
	fmt.Println(permisson.Real, permisson1.Real)
	fmt.Println(permisson.FenceAlarm, permisson1.FenceAlarm)
	fmt.Println(permisson.Fence, permisson1.Fence)
	fmt.Println(permisson.Orbit, permisson1.Orbit)
	fmt.Println(permisson.Devices, permisson1.Devices)
	fmt.Println(permisson.Users, permisson1.Users)
}
