// 每秒发送一个坐标
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	b1 := 22.57089004328858
	b2 := 113.9545858520508
	for {
		sleepnum := rand.Intn(5) + 2
		time.Sleep(time.Second * time.Duration(sleepnum))
		fmt.Println("休眠:", sleepnum)
		a1 := rand.Float64()/100 - 0.005
		a2 := rand.Float64()/100 - 0.005
		s := rand.Intn(10)
		fmt.Println(a1, a2)
		if a1 < -0.005 || a2 < -0.005 {
			continue
		}
		b1 += a1
		b2 += a2
		func() {
			data := map[string]interface{}{
				"devEUI":    "1231231231234124",
				"latitude":  b1,
				"longitude": b2,
				"name":      "大蛇王24",
				"state":     s,
				"times":     time.Now().Unix(),
				"types":     2,
				"info": map[string]interface{}{
					"大狗蛋": 123,
					"dw":  21,
				},
			}

			buf, err := json.Marshal(data)
			if err != nil {
				log.Println(err)
				return
			}
			req, err := http.NewRequest("POST", "http://120.78.76.139:8800/devicedata", bytes.NewReader(buf))
			if err != nil {
				log.Println(err.Error())
				return
			}
			req.Header.Add("Content-Type", "application/json")
			req.Header.Add("Authorization", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1ODA5Mjk3MjcsImlkIjoiMiIsInlpIjoiMTAxMTc4MTIwMTczIn0.He2N-QtU0HwdeiAchjiSLYiYIpfkz3QxZ0xhZ7QijI0")
			_, err = http.DefaultClient.Do(req)
			if err != nil {
				log.Println(err.Error())
				return
			}
		}()
	}
}
