package structure

// MessageEntity 表示文本消息中的特殊实体，例如 hashtags、用户名、URL 等
type MessageEntity struct {
	Type          string `json:"type"`                      // 实体类型，例如 mention、hashtag、bot_command 等
	Offset        int    `json:"offset"`                    // 实体起始位置的 UTF-16 代码单位偏移量
	Length        int    `json:"length"`                    // 实体长度（以 UTF-16 代码单位计）
	URL           string `json:"url,omitempty"`             // （可选）仅对 "text_link" 类型有效，表示点击文本后打开的 URL
	User          *User  `json:"user,omitempty"`            // （可选）仅对 "text_mention" 类型有效，表示被提及的用户
	Language      string `json:"language,omitempty"`        // （可选）仅对 "pre" 类型有效，表示代码块的编程语言
	CustomEmojiID string `json:"custom_emoji_id,omitempty"` // （可选）仅对 "custom_emoji" 类型有效，自定义表情的唯一 ID
}
