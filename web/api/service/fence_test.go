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
