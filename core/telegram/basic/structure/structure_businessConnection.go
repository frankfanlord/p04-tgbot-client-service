package structure

// BusinessConnection 描述机器人与商业账户的连接信息
type BusinessConnection struct {
	ID         string             `json:"id"`           // 唯一标识符
	User       *User              `json:"user"`         // 创建该连接的商业账户用户
	UserChatID int64              `json:"user_chat_id"` // 与创建该连接用户的私人聊天标识符，建议使用 int64 存储
	Date       int                `json:"date"`         // 连接建立的 Unix 时间戳
	Rights     *BusinessBotRights `json:"rights"`       // 商业机器人的权限（可选）
	IsEnabled  bool               `json:"is_enabled"`   // 连接是否处于激活状态
}
