package structure

// MessageReactionCountUpdated 表示一个带有匿名反应的消息的反应变化。
type MessageReactionCountUpdated struct {
	Chat      *Chat            `json:"chat"`       // 包含该消息的聊天
	MessageID int              `json:"message_id"` // 聊天中唯一的消息标识符
	Date      int              `json:"date"`       // 变化的时间，Unix 时间戳格式
	Reactions []*ReactionCount `json:"reactions"`  // 当前消息上存在的反应列表
}
