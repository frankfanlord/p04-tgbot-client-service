package structure

// SetMyDescriptionRequest 设置机器人的描述请求结构体
type SetMyDescriptionRequest struct {
	Description  string `json:"description"`   // 新的机器人描述；0-512个字符，传入空字符串表示删除当前语言的描述
	LanguageCode string `json:"language_code"` // 两个字母的 ISO 639-1 语言代码。如果为空，将应用于所有无专属描述的用户语言
}

// GetMyNameResponse 用于获取机器人的描述。
// 成功时返回 true。
type GetMyDescriptionResponse struct {
	Description string `json:"description"` // 新的机器人描述；0-512个字符，传入空字符串表示删除当前语言的描述
}
