package entity

type Request struct {
	// 发送方id
	SenderId uint32
	// 发送方名称
	SenderName string
	// 请求对象
	Obj
}