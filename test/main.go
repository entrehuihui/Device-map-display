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
		time.Sleep(time.Second * 2)
		a1 := rand.Float64()/100 - 0.005
		a2 := rand.Float64()/100 - 0.005
		s := rand.Int() % 10
		fmt.Println(a1, a2)
		if a1 < -0.005 || a2 < -0.005 {
			continue
		}
		b1 += a1
		b2 += a2
		func() {
			data := map[string]interface{}{
				"devEUI":    "1231231231234123",
				"latitude":  b1,
				"longitude": b2,
				"name":      "大蛇王19",
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
			req.Header.Add("Authorization", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NjQxMzQ2NDgsImlkIjoiMiIsInlpIjoiMTAwMTAzMjM1MTE3In0.lDP_RoLkDlc0odvcAR5gW7QGDuHu6QZNxD9eXLNQAcQ")
			_, err = http.DefaultClient.Do(req)
			if err != nil {
				log.Println(err.Error())
				return
			}
		}()
	}
}
