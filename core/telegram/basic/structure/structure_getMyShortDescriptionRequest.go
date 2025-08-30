package structure

// GetMyShortDescriptionRequest 获取当前机器人的简短描述，请求参数
type GetMyShortDescriptionRequest struct {
	LanguageCode string `json:"language_code"` // 两位 ISO 639-1 语言代码，或为空字符串
}
