package structure

// ReplyKeyboardMarkup 表示一个带回复选项的自定义键盘（频道和商业帐号不支持）
type ReplyKeyboardMarkup struct {
	Keyboard              [][]*KeyboardButton `json:"keyboard"`                // 自定义按钮的二维数组，每一行由若干 KeyboardButton 组成
	IsPersistent          bool                `json:"is_persistent"`           // （可选）请求客户端在隐藏系统键盘时仍显示自定义键盘，默认为 false
	ResizeKeyboard        bool                `json:"resize_keyboard"`         // （可选）请求客户端根据内容调整键盘高度，默认为 false
	OneTimeKeyboard       bool                `json:"one_time_keyboard"`       // （可选）使用后隐藏键盘，默认为 false
	InputFieldPlaceholder string              `json:"input_field_placeholder"` // （可选）激活键盘时输入框显示的占位符文本，1-64 个字符
	Selective             bool                `json:"selective"`               // （可选）是否只对特定用户显示键盘，如被提及用户或原消息发送者
}
