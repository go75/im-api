package entity

// 群聊消息数据
type GroupMessage struct {
	ID uint64 `gorm:"primary_key"`
	// 群聊id
	GroupId uint32
	// 发送方id
	SenderId uint32
	// 发送方名称
	SenderName string
	// 消息类型 文字为0, 图片为1, 音频为2
	Type uint8
	// 消息内容
	Content string
}
