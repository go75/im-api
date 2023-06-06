package handler

import (
	"context"
	"im-api/im-ws/common/dto"
	"im-api/im-ws/entity"
	"im-api/im-ws/global"
	"im-api/im-ws/log"
	"im-api/im-ws/pb"
	"im-api/im-ws/utils"
)

// 获取群聊会话
func GetGroupSession(r *entity.Request) {
	res, err := global.RPC.GetGroupSession(context.Background(), &pb.ProcessId{
		Id: r.ProcessId,
	})

	if err != nil {
		log.Warn.Println("get group session error:", err)
		return
	}

	err = utils.Send(r.SenderId, r.ID, entity.Text, r.ProcessId, res.Msgs)
	if err != nil {
		log.Warn.Println("发送websocket消息失败: ", err)
		return
	}
}

// 获取新群聊信息
func GetNewGroup(r *entity.Request) {
	res, err := global.RPC.GetNewGroup(context.Background(), &pb.ProcessId{
		Id: r.SenderId,
	})
	
	if err != nil {
		log.Warn.Println("get new group error:", err)
		return
	}

	utils.Send(r.SenderId, r.ID, entity.Text, 0, res.Msgs)
	if err != nil {
		log.Warn.Println("发送websocket消息失败: ", err)
		return
	}
}

// 发送群聊消息
func SendGroupMsg(r *entity.Request) {
	res, err := global.RPC.SendGroupMsg(context.Background(), &pb.GroupMessage{
		GroupId: r.ProcessId,
		SenderId: r.SenderId,
		Type: uint32(r.Type),
		Content: r.Payload,
	})

	if err != nil {
		log.Warn.Println("send group error:", err)
		return
	}

	for _, id := range res.UserIds.Ids {
		utils.Send(id, r.ID, entity.Text, r.ProcessId, res.Message)
	}
}

// 获取群聊列表
func GetGroupList(r *entity.Request) {
	res, err := global.RPC.GetGroupList(context.Background(), &pb.ProcessId{
		Id: r.SenderId,
	})

	if err != nil {
		log.Warn.Println("get group list error:", err)
		return
	}

	err = utils.Send(r.SenderId, r.ID, entity.Text, 0, res.Groups)
	if err != nil {
		log.Warn.Println("发送websocket消息失败: ", err)
		return
	}
}

// 添加群聊
func AddGroup(r *entity.Request) {
	res, err := global.RPC.AddGroup(context.Background(), &pb.AddGroupReq{
		Id2: &pb.Id2{
			SenderId: r.SenderId,
			ProcessId: r.ProcessId,
		},
		SenderName: r.SenderName,
	})

	if err != nil {
		log.Warn.Println("add group error: ", err)
		return
	}

	data := &dto.AddGroup{
		GroupId: res.GroupId,
		Username: r.SenderName,
		GroupName: res.GroupName,
	}

	err = utils.Send(res.MasterId, r.ID, r.Type, r.ProcessId, data)
	if err != nil {
		log.Warn.Println("发送websocket消息失败: ", err)
		return
	}
}

// 通过群聊名称模糊查询群聊
func GetFuzzyGroupByGroupName(r *entity.Request) {
	res, err := global.RPC.GetFuzzyGroupByGroupName(context.Background(), &pb.Name{
		Name: r.Payload,
	})

	if err != nil {
		log.Warn.Println("get fuzzy group by group name error: ", err)
		return
	}

	err = utils.Send(r.SenderId, r.ID, entity.Text, 0, res.Groups)
	if err != nil {
		log.Warn.Println("发送websocket消息失败: ", err)
		return
	}
}

// 同意新群友请求
func AgreeNewGroup(r *entity.Request) {
	res, err := global.RPC.AgreeNewGroup(context.Background(), &pb.AgreeNewGroupReq{
		SenderName: r.Payload,
		GroupId: r.ProcessId,
	})

	if err != nil {
		log.Warn.Println("agree new group error: ", err)
		return
	}

	// 返回新群聊的信息
	err = utils.Send(res.ReceiverId, r.ID, entity.Text, 0, res.Group)
	if err != nil {
		log.Warn.Println("发送websocket消息失败: ", err)
		return
	}
}

// 拒绝新群友请求
func RefuseNewGroup(r *entity.Request) {
	res, err := global.RPC.RefuseNewGroup(context.Background(), &pb.Id2{
		SenderId: r.SenderId,
		ProcessId: r.ProcessId,
	})
	if !res.Ok || err != nil {
		log.Warn.Println("refuse new group error: ", err)
		return
	}
}
