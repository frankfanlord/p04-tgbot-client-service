package structure

// VideoNote 表示一个视频消息对象（从 Telegram v4.0 开始支持）
type VideoNote struct {
	FileID       string     `json:"file_id"`        // 文件的标识符，可用于下载或重用该文件
	FileUniqueID string     `json:"file_unique_id"` // 文件的唯一标识符，跨时间和机器人保持一致，但不能用于下载或重用
	Length       int        `json:"length"`         // 视频的宽度和高度（即视频消息的直径），由发送方定义
	Duration     int        `json:"duration"`       // 视频时长（以秒为单位），由发送方定义
	Thumbnail    *PhotoSize `json:"thumbnail"`      // 可选，视频缩略图
	FileSize     int        `json:"file_size"`      // 可选，文件大小（以字节为单位）
}
