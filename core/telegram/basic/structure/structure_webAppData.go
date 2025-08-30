package structure

// WebAppData 表示从 Web 应用发送到 Bot 的数据。
type WebAppData struct {
	// data 是数据内容。注意，恶意客户端可能会在此字段中发送任意数据。
	Data string `json:"data"`

	// button_text 是打开 Web 应用所使用的 web_app 键盘按钮的文本。
	// 注意，恶意客户端可能会在此字段中发送任意数据。
	ButtonText string `json:"button_text"`
}
