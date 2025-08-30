package structure

type PinChatMessage[T NS] struct {
	BusinessConnectionID string `json:"business_connection_id"` // 代表该业务连接固定消息的唯一标识符
	ChatID               T      `json:"chat_id"`                // 目标聊天或目标频道用户名的唯一标识符（格式为@channelusername）
	MessageID            int    `json:"message_id"`             // 要固定的消息的标识符
	DisableNotification  bool   `json:"disable_notification"`   // 如果不需要向所有聊天成员发送有关新置顶消息的通知，则传递 True。在频道和私人聊天中，通知始终处于禁用状态。
}
