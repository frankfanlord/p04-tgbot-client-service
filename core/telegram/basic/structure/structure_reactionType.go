package structure

// ReactionType 表示一个通用的反应类型结构体接口占位符。
type ReactionType struct {
	Type          string `json:"type"`            // 反应类型，固定为 "emoji"/"custom_emoji"/"paid"
	Emoji         string `json:"emoji"`           // 表情符号，例如 "❤", "👍" 等 --- "emoji" only
	CustomEmojiID string `json:"custom_emoji_id"` // 自定义表情的唯一标识符 --- "custom_emoji" only
}
