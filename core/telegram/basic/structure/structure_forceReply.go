package structure

// ForceReply 表示强制客户端显示回复界面（频道和商业帐号不支持）
type ForceReply struct {
	ForceReply            bool   `json:"force_reply"`             // 请求客户端显示回复界面（就像用户手动点击“回复”）
	InputFieldPlaceholder string `json:"input_field_placeholder"` // （可选）激活回复时输入框显示的占位符文本，1-64 个字符
	Selective             bool   `json:"selective"`               // （可选）是否只对特定用户强制回复
}
