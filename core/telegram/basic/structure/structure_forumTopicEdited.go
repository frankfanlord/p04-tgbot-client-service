package structure

// ForumTopicEdited 表示一个关于论坛话题被编辑的服务消息。
type ForumTopicEdited struct {
	Name              string `json:"name"`                 // （可选）如果话题名称被编辑，则为新的话题名称
	IconCustomEmojiID string `json:"icon_custom_emoji_id"` // （可选）如果话题图标被编辑，则为新的自定义表情符号 ID；如果图标被移除则为空字符串
}
