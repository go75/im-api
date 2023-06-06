package main

import (
	"im-api/im-http/global"
	"im-api/im-http/initialize"
)

// @title im系统http网关接口
// @version 1.0 版本
// @description http网关功能: 登录, 注册, 头像
// @contact.name 陈序缘
// @contact.url https://github.com/go75?tab=repositories
func main() {
	initialize.Config()
	initialize.RPC()
	r := initialize.Router()
	err := r.Run(global.Config.Server.Addr)
	if err != nil {
		panic(err)
	}
}
