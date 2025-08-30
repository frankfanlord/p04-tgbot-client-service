package structure

// InlineKeyboardButton 表示内联键盘上的一个按钮，必须指定其中一个可选字段来定义按钮的行为。
type InlineKeyboardButton struct {
	Text                         string                       `json:"text"`                                       // 按钮上的标签文本
	Url                          string                       `json:"url,omitempty"`                              // 可选。点击按钮时打开的 HTTP 或 tg:// 链接
	CallbackData                 string                       `json:"callback_data,omitempty"`                    // 可选。按钮点击时发送给机器人的回调数据，1-64 字节
	WebApp                       *WebAppInfo                  `json:"web_app,omitempty"`                          // 可选。描述点击按钮时要启动的 Web 应用，仅支持私聊
	LoginUrl                     *LoginUrl                    `json:"login_url,omitempty"`                        // 可选。用于自动授权用户的 HTTPS 链接，可替代 Telegram 登录部件
	SwitchInlineQuery            string                       `json:"switch_inline_query,omitempty"`              // 可选。点击按钮后在输入框插入指定内联查询，可为空
	SwitchInlineQueryCurrentChat string                       `json:"switch_inline_query_current_chat,omitempty"` // 可选。在当前聊天中插入内联查询
	SwitchInlineQueryChosenChat  *SwitchInlineQueryChosenChat `json:"switch_inline_query_chosen_chat,omitempty"`  // 可选。选择指定类型聊天后插入内联查询
	CopyText                     *CopyTextButton              `json:"copy_text,omitempty"`                        // 可选。按钮点击后复制指定文本到剪贴板
	CallbackGame                 *CallbackGame                `json:"callback_game,omitempty"`                    // 可选。点击按钮时启动的游戏（必须是第一行的第一个按钮）
	Pay                          bool                         `json:"pay,omitempty"`                              // 可选。设置为 True 以发送付款按钮（只能用于发票消息）
}
