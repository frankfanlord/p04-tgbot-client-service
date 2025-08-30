package structure

// LoginUrl 表示用于自动授权用户的内联键盘按钮参数
type LoginUrl struct {
	Url                string `json:"url"`                  // 一个 HTTPS URL，在点击按钮时将附加用户授权数据的查询参数。如果用户拒绝授权，则会打开原始 URL。
	ForwardText        string `json:"forward_text"`         // （可选）当消息被转发时按钮的新文本。
	BotUsername        string `json:"bot_username"`         // （可选）用于用户授权的机器人用户名；如果未指定，则默认为当前机器人。
	RequestWriteAccess bool   `json:"request_write_access"` // （可选）传递 True 以请求允许机器人向用户发送消息的权限。
}
