package structure

// TextQuote 表示消息中引用部分的信息
type TextQuote struct {
	Text     string           `json:"text"`      // 被引用消息中的文本内容
	Entities []*MessageEntity `json:"entities"`  // 可选。出现在引用中的特殊格式实体（如加粗、斜体等）
	Position int              `json:"position"`  // 被引用文本在原始消息中的大致位置（UTF-16 单位）
	IsManual bool             `json:"is_manual"` // 可选。如果为 true，表示引用是用户手动选择的；否则是服务器自动添加的
}
