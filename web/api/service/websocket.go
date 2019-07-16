package service

import (
	"fmt"
	"log"

	"golang.org/x/net/websocket"
)

var socketList = make(map[int]map[string]*websocket.Conn, 0)

// GroupWrite 群发信息到websocket客户端上
func GroupWrite(uid int, data string) {
	defer func() {
		if err := recover(); err != nil {
			log.Printf("websocket  GroupWrite error! [error]%s  [data]%s", err, data)
		}
	}()
	conns := socketList[uid]
	for _, conn := range conns {
		if err := websocket.Message.Send(conn, data); err != nil {
			// 发送失败
			log.Printf("websocket send failed: %s\n data:%s", err, data)
		}
	}
}

// Echo ..
func Echo(ws *websocket.Conn) {
	query := ws.Request().URL.Query()
	jwtList := query["Authorization"]
	if jwtList == nil || len(jwtList) < 1 {
		return
	}
	jwt := jwtList[0]
	if jwt == "" {
		return
	}
	// 鉴权
	info, err := ParseJWT(jwt)
	if err != nil {
		log.Println("WebSocket ParseJWT Error:", err)
		return
	}
	//鉴权成功
	addr := ws.Request().RemoteAddr
	if socketList[info[0]] == nil {
		socketList[info[0]] = make(map[string]*websocket.Conn, 0)
	}
	socketList[info[0]][addr] = ws
	defer func() {
		fmt.Println("closed")
		delete(socketList[info[0]], addr)
	}()
	for {
		var reply []byte
		//websocket接受信息
		if err = websocket.Message.Receive(ws, &reply); err != nil {
			log.Printf("receive failed:%s", err)
			break
		}
		log.Println("Receive from client:" + string(reply))
		//处理接受到的请求 -- 未完成
		err = websocket.Message.Send(ws, string(reply))
		if err != nil {
			log.Println(err)
		}
	}
}
