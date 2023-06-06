package initialize

import (
	"im-api/im-http/global"
	"im-api/im-http/pb"
	"im-api/etcd"
)

func RPC() {
	// 创建客户端
	grpcConn := etcd.EtcdPull(global.Config.Etcd.Endpoints, global.Config.Etcd.ServiceName)
	global.RPC = pb.NewHTTPClient(grpcConn)
}
