package structure

// MessageOrigin 表示消息的来源对象，可以是用户、隐藏用户、聊天、频道等
type MessageOrigin struct {
	Type            string `json:"type"`                       // 消息来源类型，固定值: "user"、"hidden_user"、"chat"、"channel"
	Date            int    `json:"date"`                       // 原始消息发送的 Unix 时间戳
	SenderUser      *User  `json:"sender_user,omitempty"`      // 原始发送消息的用户（仅 type = "user" 时）
	SenderUserName  string `json:"sender_user_name,omitempty"` // 原始发送者的用户名（仅 type = "hidden_user" 时）
	SenderChat      *Chat  `json:"sender_chat,omitempty"`      // 原始发送消息的聊天（仅 type = "chat" 时）
	AuthorSignature string `json:"author_signature,omitempty"` // 原始匿名管理员的签名或频道作者签名（仅 type = "chat" 或 "channel" 时）
	Chat            *Chat  `json:"chat,omitempty"`             // 原始发送的频道聊天（仅 type = "channel" 时）
	MessageID       int    `json:"message_id,omitempty"`       // 消息在频道内的唯一标识符（仅 type = "channel" 时）
}
