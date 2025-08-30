package structure

// BusinessMessagesDeleted 表示连接的企业账户中消息被删除的通知
type BusinessMessagesDeleted struct {
	BusinessConnectionID string `json:"business_connection_id"` // 企业连接的唯一标识
	Chat                 *Chat  `json:"chat"`                   // 企业账户中某个聊天的信息，可能无法访问
	MessageIDs           []int  `json:"message_ids"`            // 被删除的消息 ID 列表
}
