package structure

// CallbackQuery 表示从内联键盘中的回调按钮接收到的回调查询。
type CallbackQuery struct {
	ID              string                    `json:"id"`                // 查询的唯一标识符
	From            *User                     `json:"from"`              // 发送该回调的用户
	Message         *MaybeInaccessibleMessage `json:"message"`           // 可选. 由机器人发送的消息，包含产生此查询的回调按钮
	InlineMessageID string                    `json:"inline_message_id"` // 可选. 通过机器人以 inline 模式发送的消息的标识符
	ChatInstance    string                    `json:"chat_instance"`     // 全局唯一标识符，对应包含回调按钮的消息所在的聊天
	Data            string                    `json:"data"`              // 可选. 与回调按钮关联的数据
	GameShortName   string                    `json:"game_short_name"`   // 可选. 游戏的短名称，作为游戏的唯一标识符
}
