package structure

// BotCommandScope 表示融合的 Bot 指令作用域结构体
type BotCommandScope struct {
	Type   string `json:"type"`    // 指令作用域类型，可为 default, all_private_chats, all_group_chats, all_chat_administrators, chat, chat_administrators, chat_member
	ChatID string `json:"chat_id"` // 聊天唯一标识符或超级群组用户名（如 @supergroupusername），适用于 chat/chat_administrators/chat_member
	UserID int    `json:"user_id"` // 用户唯一标识符，仅适用于 chat_member
}
