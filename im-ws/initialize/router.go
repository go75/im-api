package initialize

import (
	"im-api/im-ws/api"
	"im-api/im-ws/entity"
	"im-api/im-ws/handler"
)

func Router() {
	registFiveChess()
	registFriend()
	registGroup()
}

// 注册friend模块路由
func registFriend() {
	// 获取好友会话
	api.Regist(entity.GetFriendSession, handler.GetFriendSession)
	// 获取新好友信息
	api.Regist(entity.GetNewFriend, handler.GetNewFriend)
	// 获取好友列表
	api.Regist(entity.GetFriendList, handler.GetFriendList)
	// 添加好友
	api.Regist(entity.AddFriend, handler.AddFriend)
	// 发送好友消息
	api.Regist(entity.SendFriendMsg, handler.SendFriendMsg)
	// 通过用户名称模糊查询用户
	api.Regist(entity.GetFuzzyUserByUserName, handler.GetFuzzyUserByUserName)
	// 同意新好友请求	
	api.Regist(entity.AgreeNewFriend, handler.AgreeNewFriend)
	// 拒绝新好友请求
	api.Regist(entity.RefuseNewFriend, handler.RefuseNewFriend)
}

// 注册group模块路由
func registGroup() {
	// 获取群聊会话
	api.Regist(entity.GetGroupSession, handler.GetGroupSession)
	// 获取新群友信息
	api.Regist(entity.GetNewGroup, handler.GetNewGroup)
	// 获取群聊列表
	api.Regist(entity.GetGroupList, handler.GetGroupList)
	// 添加群聊
	api.Regist(entity.AddGroup, handler.AddGroup)
	// 发送群聊消息
	api.Regist(entity.SendGroupMsg, handler.SendGroupMsg)
	// 通过群聊名称模糊查询群聊
	api.Regist(entity.GetFuzzyGroupByGroupName, handler.GetFuzzyGroupByGroupName)
	// 同意新群友请求
	api.Regist(entity.AgreeNewGroup, handler.AgreeNewGroup)
	// 拒绝新群友请求
	api.Regist(entity.RefuseNewGroup, handler.RefuseNewGroup)
}

// 注册五子棋模块路由
func registFiveChess() {
	// 请求五子棋对战
	api.Regist(entity.RequestFiveChess, handler.FiveChess)
	// 同意五子棋对战
	api.Regist(entity.AgreeFiveChess, handler.AgreeFiveChess)
	// 拒绝五子棋对战
	api.Regist(entity.RefuseFiveChess, handler.RefuseFiveChess)	
	// 发送下棋位置
	api.Regist(entity.DownChessPosition, handler.DownChessPosition)
}