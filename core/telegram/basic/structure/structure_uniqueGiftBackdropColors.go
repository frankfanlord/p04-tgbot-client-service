package structure

// UniqueGiftBackdropColors 表示一个独特礼物背景的颜色设置
type UniqueGiftBackdropColors struct {
	CenterColor int `json:"center_color"` // 背景中心的颜色（RGB 格式）
	EdgeColor   int `json:"edge_color"`   // 背景边缘的颜色（RGB 格式）
	SymbolColor int `json:"symbol_color"` // 应用于符号的颜色（RGB 格式）
	TextColor   int `json:"text_color"`   // 背景上文字的颜色（RGB 格式）
}
