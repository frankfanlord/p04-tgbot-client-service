package structure

// ForumTopicCreated 表示聊天中创建新话题的服务消息对象
type ForumTopicCreated struct {
	Name              string `json:"name"`                 // 话题名称
	IconColor         int    `json:"icon_color"`           // 话题图标的 RGB 颜色
	IconCustomEmojiID string `json:"icon_custom_emoji_id"` // 可选。用作话题图标的自定义表情符唯一标识符
}
