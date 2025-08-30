package structure

type DeleteMessage[T NS] struct {
	ChatID    T   `json:"chat_id"`    // 整数或字符串 是 目标聊天或目标频道用户名的唯一标识符（格式为@channelusername）
	MessageID int `json:"message_id"` // 待删除消息的标识符
}
