package structure

// SwitchInlineQueryChosenChat 表示一个内联按钮，用于将当前用户切换到指定聊天的内联模式，并带有可选的默认内联查询。
type SwitchInlineQueryChosenChat struct {
	Query             string `json:"query"`               // 可选。将插入输入字段的默认内联查询。如果为空，则只插入机器人的用户名。
	AllowUserChats    bool   `json:"allow_user_chats"`    // 可选。如果为 true，则可以选择与用户的私人聊天。
	AllowBotChats     bool   `json:"allow_bot_chats"`     // 可选。如果为 true，则可以选择与机器人的私人聊天。
	AllowGroupChats   bool   `json:"allow_group_chats"`   // 可选。如果为 true，则可以选择群组和超级群组聊天。
	AllowChannelChats bool   `json:"allow_channel_chats"` // 可选。如果为 true，则可以选择频道聊天。
}
