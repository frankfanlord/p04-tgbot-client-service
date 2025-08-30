package structure

// ChatShared 表示通过 KeyboardButtonRequestChat 按钮与机器人共享的聊天信息。
type ChatShared struct {
	RequestID int          `json:"request_id"`         // 请求的标识符
	ChatID    int64        `json:"chat_id"`            // 共享聊天的标识符，可能超过32位，但不超过52位
	Title     string       `json:"title,omitempty"`    // 聊天标题（如果机器人请求了标题）
	Username  string       `json:"username,omitempty"` // 聊天的用户名（如果请求并可用）
	Photo     []*PhotoSize `json:"photo,omitempty"`    // 聊天头像的可用尺寸（如果请求了头像）
}
