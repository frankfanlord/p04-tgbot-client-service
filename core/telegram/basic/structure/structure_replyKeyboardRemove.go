package structure

// ReplyKeyboardRemove 表示请求客户端移除当前自定义键盘，恢复默认键盘（频道和商业帐号不支持）
type ReplyKeyboardRemove struct {
	RemoveKeyboard bool `json:"remove_keyboard"` // 请求移除键盘（完全禁用此键盘，不能再次召唤）
	Selective      bool `json:"selective"`       // （可选）是否只对特定用户移除键盘
}
