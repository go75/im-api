package handler

import (
	"context"
	"im-api/im-ws/entity"
	"im-api/im-ws/global"
	"im-api/im-ws/log"
	"im-api/im-ws/pb"
	"im-api/im-ws/utils"
)

// 获取好友会话
func GetFriendSession(r *entity.Request) {
	res, err := global.RPC.GetFriendSession(context.Background(), &pb.Id2{
		SenderId: r.SenderId,
		ProcessId: r.ProcessId,
	})
	if err != nil {
		log.Warn.Println("get friend session error:", err)
		return
	}

	err = utils.Send(r.SenderId, r.ID, entity.Text, 0, res.Msgs)
	if err != nil {
		log.Warn.Println("发送websocket消息失败: ", err)
		return
	}
}

// 获取新好友信息
func GetNewFriend(r *entity.Request) {
	res, err := global.RPC.GetNewFriend(context.Background(), &pb.ProcessId{
		Id: r.SenderId,
	})

	if err != nil {
		log.Warn.Println("get new friend error:", err)		
		return
	}

	err = utils.Send(r.SenderId, r.ID, entity.Text, 0, res.Users)
	if err != nil {
		log.Warn.Println("发送websocket消息失败: ", err)
		return
	}
}

// 发送好友消息
func SendFriendMsg(r *entity.Request) {
	// 用户消息持久化
	msg := &pb.UserMessage{
		SenderId: r.SenderId,
		ReceiverId: r.ProcessId,
		Type: uint32(r.Type),
		Content: r.Payload,
	}

	ok, err := global.RPC.SaveUserMessage(context.Background(), msg)
	if !ok.Ok || err != nil {
		log.Warn.Println("send friend msg error:", err)		
		return
	}

	err = utils.Send(r.ProcessId, r.ID, r.Type, r.SenderId, msg)
	if err != nil {
		log.Warn.Println("发送websocket消息失败: ", err)
		return
	}
}

// 获取好友列表
func GetFriendList(r *entity.Request) {
	// 查询用户会话
	res, err := global.RPC.GetFriendList(context.Background(), &pb.ProcessId{Id: r.SenderId})
	if err != nil {
		log.Warn.Println("get friend list error:", err)		
		return
	}

	err = utils.Send(r.SenderId, r.ID, r.Type, r.ProcessId, res.Users)
	if err != nil {
		log.Warn.Println("发送websocket消息失败: ", err)
		return
	}
}

// 添加好友
func AddFriend(r *entity.Request) {
	// 新好友消息持久化
	res, err := global.RPC.AddFriend(context.Background(), &pb.Id2{
		SenderId: r.SenderId,
		ProcessId: r.ProcessId,
	})
	if err != nil {
		log.Warn.Println("add friend error:", err)
		return
	}

	err = utils.Send(r.ProcessId, r.ID, entity.Text, r.ProcessId, res)
	if err != nil {
		log.Warn.Println("发送websocket消息失败: ", err)
		return
	}
}

// 通过用户名称模糊查询用户
func GetFuzzyUserByUserName(r *entity.Request) {
	res, err := global.RPC.GetFuzzyUserByUserName(context.Background(), &pb.Name{
		Name: r.Payload,
	})
	if err != nil {
		log.Warn.Println("get fuzzy user by user name error:", err)		
		return
	}
	err = utils.Send(r.SenderId, r.ID, entity.Text, 0, res.Users)
	if err != nil {
		log.Warn.Println("发送websocket消息失败: ", err)
		return
	}
}

// 同意新好友请求
func AgreeNewFriend(r *entity.Request) {
	res, err := global.RPC.AgreeNewFriend(context.Background(), &pb.Id2{
		SenderId: r.SenderId,
		ProcessId: r.ProcessId,
	})

	if err != nil {
		log.Warn.Println("agree new friend error:", err)	
		return
	}

	// 返回新好友的信息
	err = utils.Send(r.ProcessId, r.ID, entity.Text, 0, res)
	if err != nil {
		log.Warn.Println("发送websocket消息失败: ", err)
		return
	}
}

// 拒绝新好友请求
func RefuseNewFriend(r *entity.Request) {
	ok, err := global.RPC.RefuseNewFriend(context.Background(), &pb.Id2{
		SenderId: r.SenderId,
		ProcessId: r.ProcessId,
	})
	if !ok.Ok || err != nil {
		log.Warn.Println("refuse new friend error: ", err)
		return
	}
}