package structure

type Video struct {
	FileID         string       `json:"file_id"`         // 文件标识符，可用于下载或重用该文件
	FileUniqueID   string       `json:"file_unique_id"`  // 文件的唯一标识符，对不同机器人和时间应保持一致，不能用于下载或重用
	Width          int          `json:"width"`           // 视频宽度，由发送方定义
	Height         int          `json:"height"`          // 视频高度，由发送方定义
	Duration       int          `json:"duration"`        // 视频时长（以秒为单位），由发送方定义
	Thumbnail      *PhotoSize   `json:"thumbnail"`       // 可选，视频缩略图
	Cover          []*PhotoSize `json:"cover"`           // 可选，消息中视频封面的多种尺寸
	StartTimestamp int          `json:"start_timestamp"` // 可选，视频将在消息中播放的起始秒数
	FileName       string       `json:"file_name"`       // 可选，发送方定义的原始文件名
	MimeType       string       `json:"mime_type"`       // 可选，发送方定义的文件 MIME 类型
	FileSize       int          `json:"file_size"`       // 可选，文件大小（字节），最大可安全使用 64 位整数存储
}
