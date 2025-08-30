package structure

// Animation 表示一个动画文件（无声的 GIF 或 H.264/MPEG-4 AVC 视频）。
type Animation struct {
	FileID       string     `json:"file_id"`        // 文件标识符，可用于下载或复用该文件
	FileUniqueID string     `json:"file_unique_id"` // 文件的唯一标识符，在不同机器人之间应保持一致，不能用于下载或复用
	Width        int        `json:"width"`          // 视频宽度，由发送者定义
	Height       int        `json:"height"`         // 视频高度，由发送者定义
	Duration     int        `json:"duration"`       // 视频时长（单位：秒），由发送者定义
	Thumbnail    *PhotoSize `json:"thumbnail"`      // （可选）动画缩略图，由发送者定义
	FileName     string     `json:"file_name"`      // （可选）原始动画文件名，由发送者定义
	MimeType     string     `json:"mime_type"`      // （可选）文件的 MIME 类型，由发送者定义
	FileSize     int        `json:"file_size"`      // （可选）文件大小（字节数），可超过2^31，建议使用 int64 存储
}
