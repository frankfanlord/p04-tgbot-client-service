package structure

// SetMyNameRequest 用于设置机器人的名称。
// 成功时返回 true。
type SetMyNameRequest struct {
	Name         string `json:"name"`          // 新的机器人名称；0-64 个字符。传空字符串表示删除指定语言的专用名称。
	LanguageCode string `json:"language_code"` // 两位的 ISO 639-1 语言代码。如果为空，将对所有没有专用名称的用户显示该名称。
}

// GetMyNameResponse 用于获取机器人的名称。
// 成功时返回 true。
type GetMyNameResponse struct {
	Name string `json:"name"` // 新的机器人名称；0-64 个字符。传空字符串表示删除指定语言的专用名称。
}
