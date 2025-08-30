package structure

// PaidMedia 描述付费媒体对象，目前可能是以下几种类型：
// - PaidMediaPreview：付费前的预览
// - PaidMediaPhoto：付费图片
// - PaidMediaVideo：付费视频
type PaidMedia struct {
	Type     string       `json:"type"`     // 媒体类型（如 preview, photo, video）
	Width    int          `json:"width"`    // 媒体宽度（可选，由发送者定义）
	Height   int          `json:"height"`   // 媒体高度（可选，由发送者定义）
	Duration int          `json:"duration"` // 媒体时长（秒）（可选，由发送者定义） 仅在 type = preview 时有以下字段
	Photo    []*PhotoSize `json:"photo"`    // 图片数组 仅在 type = photo 时有以下字段
	Video    *Video       `json:"video"`    // 视频对象 仅在 type = video 时有以下字段
}
