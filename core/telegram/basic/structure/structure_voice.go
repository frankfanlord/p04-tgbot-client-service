package structure

// Voice 表示一个语音消息对象。
type Voice struct {
	FileID       string `json:"file_id"`        // 文件的标识符，可用于下载或复用该文件
	FileUniqueID string `json:"file_unique_id"` // 文件的唯一标识符，跨时间和机器人保持一致，不能用于下载或复用
	Duration     int    `json:"duration"`       // 音频的持续时间（秒），由发送方定义
	MimeType     string `json:"mime_type"`      // （可选）文件的 MIME 类型，由发送方定义
	FileSize     int    `json:"file_size"`      // （可选）文件大小（字节），可能超过 2^31，建议使用 int64 类型
}
