package structure

// Document 表示一个通用文件（区别于照片、语音消息和音频文件）
type Document struct {
	FileID       string     `json:"file_id"`        // 文件标识符，可用于下载或重用文件
	FileUniqueID string     `json:"file_unique_id"` // 文件的唯一标识符，在不同机器人和时间中保持不变，不能用于下载或重用
	Thumbnail    *PhotoSize `json:"thumbnail"`      // 可选。由发送者定义的文档缩略图
	FileName     string     `json:"file_name"`      // 可选。由发送者定义的原始文件名
	MimeType     string     `json:"mime_type"`      // 可选。由发送者定义的文件 MIME 类型
	FileSize     int        `json:"file_size"`      // 可选。文件大小（字节），可大于 2^31，最多有 52 位有效数字
}
