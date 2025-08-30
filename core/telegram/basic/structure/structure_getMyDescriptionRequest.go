package structure

// GetMyDescriptionRequest 获取当前机器人的描述，请求参数
type GetMyDescriptionRequest struct {
	LanguageCode string `json:"language_code"` // 两位 ISO 639-1 语言代码，或为空字符串
}
