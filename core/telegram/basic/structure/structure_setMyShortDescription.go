package structure

// SetMyShortDescriptionRequest 设置机器人简短描述请求结构体
type SetMyShortDescriptionRequest struct {
	ShortDescription string `json:"short_description"` // 新的简短描述；0-120个字符，传入空字符串表示删除当前语言的简短描述
	LanguageCode     string `json:"language_code"`     // 两个字母的 ISO 639-1 语言代码。如果为空，将应用于所有无专属简短描述的用户语言
}

// GetMyShortDescriptionResponse 用于获取机器人的简短描述。
// 成功时返回 true。
type GetMyShortDescriptionResponse struct {
	Description string `json:"description"` // 新的机器人简短描述；0-512个字符，传入空字符串表示删除当前语言的描述
}
