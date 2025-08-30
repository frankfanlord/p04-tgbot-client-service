package structure

// GetMyNameRequest 获取当前机器人的名称，请求参数
type GetMyNameRequest struct {
	LanguageCode string `json:"language_code"` // 两位 ISO 639-1 语言代码，或为空字符串
}
