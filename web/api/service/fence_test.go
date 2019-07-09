package service

import (
	"fmt"
	"testing"
)

func Test_twoPointsDistance(t *testing.T) {
	s := twoPointsDistance(22.87152917188582, 120.36482691764833, 22.9471, 113.68)
	fmt.Printf("%.6f\n", s)
	fmt.Println(s)
}

func Test_polyOut(t *testing.T) {
	fenceinfo2 := []Fenceinfo2{
		Fenceinfo2{
			Lat: 22.586752469922832,
			Lng: 113.91582012176515,
		},
		Fenceinfo2{
			Lat: 22.578708624700443,
			Lng: 113.91582012176515,
		},
		Fenceinfo2{
			Lat: 2.578708624700443,
			Lng: 113.92654895782472,
		},
		Fenceinfo2{
			Lat: 22.586752469922832,
			Lng: 113.92654895782472,
		},
	}
	// true
	s := polyOut(22.58279004328858, 113.92611980438234, fenceinfo2)
	fmt.Printf("Expectation:true  Resule:%v\n", s)
	// // true
	// s = polyOut(22.561153184238872, 113.90642166137697, fenceinfo2)
	// fmt.Printf("Expectation:true  Resule:%v\n", s)
	// // false
	// s = polyOut(22.583285352851178, 113.91906023025514, fenceinfo2)
	// fmt.Printf("Expectation:false  Resule:%v\n", s)
}

// [{"Lat":22.579243574798344,"Lng":113.91264438629152},{"Lat":22.57934263940337,"Lng":113.92343759536745},{"Lat":22.586712846220713,"Lng":113.92337322235107},{"Lat":22.58702983551847,"Lng":113.91723632812501}]
