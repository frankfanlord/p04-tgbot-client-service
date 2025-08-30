package structure

// BotCommand 表示一个机器人命令。
type BotCommand struct {
	Command     string `json:"command"`     // 命令的文本；1-32 个字符，仅包含小写英文字母、数字和下划线
	Description string `json:"description"` // 命令的描述；1-256 个字符
}
