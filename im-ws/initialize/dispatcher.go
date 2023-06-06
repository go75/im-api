package initialize

import (
	"encoding/json"
	"im-api/im-ws/api"
	"im-api/im-ws/entity"
	"im-api/im-ws/global"
	"log"
)

func Dispatcher() {
	for i := 0; i < global.Config.Dispatcher.Size; i++ {
		go func() {
			var req *entity.Request
			//订阅消息, 不断阻塞等待请求队列的请求
			for req = range global.Channel {
				bytes, err := json.Marshal(*req)
				if err != nil {
					log.Println("请求出错")
				}
				log.Println("请求: ", string(bytes))
				api.Do(req)
			}
		}()
	}
}