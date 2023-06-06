package main

import (
	"encoding/json"
	"im-api/im-ws/entity"
	"im-api/im-ws/global"
	"im-api/im-ws/initialize"
	"im-api/im-ws/log"
	"im-api/im-ws/manager"
	"im-api/im-ws/utils"
	"net/http"
	"net/url"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func main() {
	initialize.Config()
	initialize.Dispatcher()
	initialize.RPC()

	initialize.Router()
	http.HandleFunc("/", upgrade)
	println(global.Config.Server.Addr)
	err := http.ListenAndServe(global.Config.Server.Addr, nil)
	panic(err)
}

func upgrade(w http.ResponseWriter, r *http.Request) {
	log.Info.Println()
	cookie, err := r.Cookie("im-token")
	if err != nil {
		log.Info.Println("not have im-token")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	token, _ := url.QueryUnescape(cookie.Value)
	claims, err := utils.ParseToken(token)
	if err != nil {
		log.Info.Println("parse token failed, token: ", token)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	id := claims.ID
	name := claims.Name

	socket, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Info.Println("websocket upgrade failed")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	defer socket.Close()
	defer manager.Remove(id)
	log.Info.Println(id, "websocket open")
	err = manager.Put(id, socket)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var p []byte
	var obj = entity.Obj{}
	for {
		// 接收上游生产的数据
		_, p, err = socket.ReadMessage()
		if err != nil {
			log.Info.Println("websocket read json err: ",err)
			manager.Remove(id)
			return
		}
		err := json.Unmarshal(p, &obj)
		if err != nil {
			log.Info.Println(err)
			continue
		}

		// 填充字段
		req := &entity.Request {
			SenderId: id,
			SenderName: name,
			Obj: obj,
		}
		// 发布消息
		global.Channel <- req
	}
}