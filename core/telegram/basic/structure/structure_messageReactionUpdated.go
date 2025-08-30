package structure

// MessageReactionUpdated 表示用户对消息的反应发生变化的对象
type MessageReactionUpdated struct {
	Chat        *Chat           `json:"chat"`         // 聊天对象，包含用户做出反应的消息
	MessageID   int             `json:"message_id"`   // 消息在聊天中的唯一标识符
	User        *User           `json:"user"`         // 可选项。如果用户不是匿名的，则为更改反应的用户
	ActorChat   *Chat           `json:"actor_chat"`   // 可选项。如果用户是匿名的，则为代表其更改反应的聊天
	Date        int             `json:"date"`         // 更改的时间（Unix 时间戳）
	OldReaction []*ReactionType `json:"old_reaction"` // 用户之前设置的反应类型列表
	NewReaction []*ReactionType `json:"new_reaction"` // 用户当前设置的反应类型列表
}
