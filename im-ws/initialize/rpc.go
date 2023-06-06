package initialize

import (
	"im-api/im-ws/global"
	"im-api/im-ws/pb"
	"im-api/etcd"
)

func RPC() {
	// 创建客户端
	grpcConn := etcd.EtcdPull(global.Config.Etcd.Endpoints, global.Config.Etcd.ServiceName)
	global.RPC = pb.NewWebSocketClient(grpcConn)
}
