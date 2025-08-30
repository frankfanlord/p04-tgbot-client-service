package structure

// InlineKeyboardMarkup 表示一个内联键盘，会显示在所属消息的右侧。
type InlineKeyboardMarkup struct {
	// inline_keyboard 是一个按钮行数组，每一行由多个 InlineKeyboardButton 组成。
	InlineKeyboard [][]*InlineKeyboardButton `json:"inline_keyboard"` // 按钮行数组，每行包含多个按钮
}
