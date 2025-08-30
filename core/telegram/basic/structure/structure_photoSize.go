package structure

// PhotoSize 表示照片或文件/贴纸缩略图的一种尺寸。
type PhotoSize struct {
	FileID       string `json:"file_id"`        // 文件标识符，可用于下载或复用该文件
	FileUniqueID string `json:"file_unique_id"` // 文件的唯一标识符，跨时间和不同机器人保持不变。不可用于下载或复用
	Width        int    `json:"width"`          // 照片宽度
	Height       int    `json:"height"`         // 照片高度
	FileSize     int    `json:"file_size"`      // （可选）文件大小（字节）
}
