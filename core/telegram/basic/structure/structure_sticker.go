package structure

// Sticker 表示一个贴纸对象。
type Sticker struct {
	FileID           string        `json:"file_id"`           // 文件的标识符，可用于下载或重用该文件
	FileUniqueID     string        `json:"file_unique_id"`    // 文件的唯一标识符，在不同机器人和时间中保持不变，不能用于下载或重用
	Type             string        `json:"type"`              // 贴纸的类型，目前为 "regular"、"mask"、"custom_emoji"
	Width            int           `json:"width"`             // 贴纸的宽度
	Height           int           `json:"height"`            // 贴纸的高度
	IsAnimated       bool          `json:"is_animated"`       // 如果贴纸是动画的，则为 true
	IsVideo          bool          `json:"is_video"`          // 如果贴纸是视频贴纸，则为 true
	Thumbnail        *PhotoSize    `json:"thumbnail"`         // （可选）贴纸的缩略图，格式为 .WEBP 或 .JPG
	Emoji            string        `json:"emoji"`             // （可选）与贴纸关联的表情符号
	SetName          string        `json:"set_name"`          // （可选）贴纸所属贴纸集的名称
	PremiumAnimation *File         `json:"premium_animation"` // （可选）对于高级贴纸，附带的动画文件
	MaskPosition     *MaskPosition `json:"mask_position"`     // （可选）如果是面具贴纸，表示面具应该放置的位置
	CustomEmojiID    string        `json:"custom_emoji_id"`   // （可选）自定义表情贴纸的唯一标识符
	NeedsRepainting  bool          `json:"needs_repainting"`  // （可选）如果贴纸需要根据环境重新上色为 Telegram Premium 颜色或其他合适颜色，则为 true
	FileSize         int           `json:"file_size"`         // （可选）文件大小（字节）
}
