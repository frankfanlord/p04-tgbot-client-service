package structure

// DeleteMyCommandsRequest 用于删除机器人命令的请求结构体
type DeleteMyCommandsRequest struct {
	Scope        *BotCommandScope `json:"scope"`         // 指定命令作用范围的结构体，默认为 BotCommandScopeDefault
	LanguageCode string           `json:"language_code"` // 两位语言代码，如果为空则适用于指定作用范围内没有专属命令语言的所有用户
}
