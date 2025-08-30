package structure

// CopyTextButton 表示一个“复制文字到剪贴板”的内联键盘按钮。
type CopyTextButton struct {
	// Text 要复制到剪贴板的文本；长度为 1 到 256 个字符
	Text string `json:"text"`
}
