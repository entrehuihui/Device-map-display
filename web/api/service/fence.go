package service

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"

	"github.com/gomodule/redigo/redis"

	"../../../db"
	"github.com/jinzhu/gorm"
)

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

// 计算点是否在多边形外
func polyOut(lat, lng float64, fenceinfo2 []Fenceinfo2) bool {
	if len(fenceinfo2) < 3 {
		return false
	}
	motheds := true
	l := len(fenceinfo2)
	j := l - 1
	for i := 0; i < l; i++ {
		sx := fenceinfo2[i].Lng
		sy := fenceinfo2[i].Lat
		tx := fenceinfo2[j].Lng
		ty := fenceinfo2[j].Lat
		// 点与多边形顶点重合
		if (sx == lng && sy == lat) || (tx == lat && lng == lat) {
			return false
		}
		// 判断线段两端点是否在射线两侧
		if (sy < lat && ty >= lat) || (sy >= lat && ty < lat) {
			// 线段上与射线 Y 坐标相同的点的 X 坐标
			var x = sx + (lat-sy)*(tx-sx)/(ty-sy)
			// 点在多边形的边上
			if x == lng {
				return false
			}
			// 射线穿过多边形的边界
			if x > lng {
				motheds = !motheds
			}
			// fmt.Println("=============", fenceinfo2[i].Lng, fenceinfo2[j].Lng, x, lng)
		}
		j = i
	}
	return motheds
}

// Fenceinfo1 圆围栏
type Fenceinfo1 struct {
	Lat    float64
	Lng    float64
	Radius float64
}

// Fenceinfo2 方围栏
type Fenceinfo2 struct {
	Lat float64
	Lng float64
}

// ChenckOutFence 计算是否超出围栏 return bool: false不越界 true越界
func ChenckOutFence(uid int, lat, lng float64) bool {
	defer func() {
		if err := recover(); err != nil {
			log.Println("painc Error:", err)
		}
	}()
	values := getFence(uid)
	valueList := strings.Split(values, "||")
	if len(valueList) < 2 {
		return false
	}
	buf := []byte(valueList[1])
	if valueList[0] == "1" {
		fenceinfo1 := Fenceinfo1{}
		err := json.Unmarshal(buf, &fenceinfo1)
		if err != nil {
			log.Println("ChenckOutFence Unmarshal Error:", err, valueList[1])
			return false
		}
		length := twoPointsDistance(lat, lng, fenceinfo1.Lat, fenceinfo1.Lng)
		if fenceinfo1.Radius < length {
			return true
		}
		return false
	} else if valueList[0] == "2" {
		fenceinfo2 := make([]Fenceinfo2, 0)
		err := json.Unmarshal(buf, &fenceinfo2)
		if err != nil {
			log.Println("ChenckOutFence Unmarshal Error:", err, valueList[1])
			return false
		}
		return polyOut(lat, lng, fenceinfo2)
	}
	return false
}

func getFence(uid int) string {
	fmt.Println(getFenceMark(uid))
	conn := db.GetRedis()
	defer conn.Close()
	value, err := redis.String(conn.Do("GET", getFenceMark(uid)))
	if err != nil {
		log.Println("addFence Error : ", err)
		return selectFence(uid)
	}
	return value
}

//  数据库查找围栏数据 -
func selectFence(uid int) string {
	s := db.GetDB()
	fencelists := &db.Fencelists{}
	err := s.Where("uid = ? and del = 0", uid).Last(fencelists).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		fmt.Println("selectFence Error : ", err)
	}
	return AddFence(fencelists)
}

func getFenceMark(uid int) string {
	fenceMark := "fenceMark:"
	return fenceMark + strconv.Itoa(uid)
}

// AddFence .	添加围栏
func AddFence(fencelists *db.Fencelists) string {
	var value string
	if fencelists.Types == 1 || fencelists.Types == 2 {
		value = strconv.Itoa(fencelists.Types) + "||" + fencelists.Info
	} else {
		value = "0"
	}
	conn := db.GetRedis()
	defer conn.Close()
	_, err := conn.Do("SET", getFenceMark(fencelists.UID), value, "EX", 86400)
	if err != nil {
		log.Println("addFence Error : ", err)
	}
	return value
}

// DelFence 删除围栏缓存
func DelFence(uid int) error {
	conn := db.GetRedis()
	defer conn.Close()
	_, err := conn.Do("DEL", getFenceMark(uid))
	if err != nil {
		log.Println("addFence Error : ", err)
	}
	return err
}
