package service

import "math"
import "../../../db"

var mapFence map[int]fenceinfo

type fenceinfo struct {
	types  int //1 圆形 2多边形 3不存在
	raduis int
	latlng [][2]float64
}

func rad(lang float64) float64 {
	return lang * math.Pi / 180.0
}

// 计算经纬度两点之间距离
func twoPointsDistance(lat1, lng1, lat2, lng2 float64) float64 {
	lat1 = rad(lat1)
	lat2 = rad(lat2)
	a := lat1 - lat2
	b := rad(lng1) - rad(lng2)
	s := 2 * math.Asin(math.Sqrt(math.Pow(math.Sin(a/2), 2)+math.Cos(lat1)*math.Cos(lat2)*math.Pow(math.Sin(b/2), 2)))
	s = s * 6378137
	s = math.Round(s*10000) / 10000
	return s
}

// ChenckOutFence 计算是否超出围栏 return bool: false不越界 true越界
func (s *Server) ChenckOutFence(uid int, lat, log float64) bool {
	if mapFence[uid].types == 3 {
		return false
	} else if mapFence[uid].types == 1 {
		// 计算是否超出圆外
		if float64(mapFence[uid].raduis) >= twoPointsDistance(lat, log, mapFence[uid].latlng[0][0], mapFence[uid].latlng[0][1]) {
			return false
		}
		return true
	} else if mapFence[uid].types == 2 {

	} else {

	}
	return true
}

//  数据库查找围栏数据 --- 此处小概率会出现BUG ---就是懒得加锁耗性能
func (s *Server) selectFence(uid int) {
	fencelists := db.Fencelists{}
	err := s.Where("uid = ? and del = 0", uid).Last(&fencelists).Error
	if err != nil {

	}
}
