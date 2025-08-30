package structure

// LinkPreviewOptions 描述用于链接预览生成的选项
type LinkPreviewOptions struct {
	IsDisabled       bool   `json:"is_disabled"`        // 可选。如果为 true，则禁用链接预览
	URL              string `json:"url"`                // 可选。用于链接预览的 URL。如果为空，则使用消息文本中找到的第一个 URL
	PreferSmallMedia bool   `json:"prefer_small_media"` // 可选。如果为 true，则缩小链接预览中的媒体；如果未明确指定 URL 或不支持更改媒体大小，则忽略
	PreferLargeMedia bool   `json:"prefer_large_media"` // 可选。如果为 true，则放大链接预览中的媒体；如果未明确指定 URL 或不支持更改媒体大小，则忽略
	ShowAboveText    bool   `json:"show_above_text"`    // 可选。如果为 true，链接预览显示在消息文本上方；否则显示在下方
}
