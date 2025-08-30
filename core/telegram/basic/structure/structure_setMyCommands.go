package structure

// SetMyCommandsRequest 设置机器人命令列表的请求结构体
type SetMyCommandsRequest struct {
	Commands     []*BotCommand    `json:"commands"`      // 命令列表（最多可设置100个）
	Scope        *BotCommandScope `json:"scope"`         // 命令作用域（默认值为 BotCommandScopeDefault）
	LanguageCode string           `json:"language_code"` // 两位 ISO 639-1 语言代码。如果为空，则应用于所有用户
}
