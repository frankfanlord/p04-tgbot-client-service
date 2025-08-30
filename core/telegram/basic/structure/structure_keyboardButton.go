package structure

// KeyboardButton 表示回复键盘上的一个按钮。
// 最多只能使用一个可选字段来指定按钮的类型。
// 对于简单的文本按钮，可以直接使用字符串代替该对象。
type KeyboardButton struct {
	Text            string                      `json:"text"`                       // 按钮的文本。如果未使用任何可选字段，则按下按钮时将其作为消息发送。
	RequestUsers    *KeyboardButtonRequestUsers `json:"request_users,omitempty"`    // 可选。按下按钮时打开一个合适用户列表。选中的用户标识会通过 “users_shared” 服务消息发送给机器人，仅在私聊中可用。
	RequestChat     *KeyboardButtonRequestChat  `json:"request_chat,omitempty"`     // 可选。按下按钮时打开一个合适聊天列表。选中的聊天 ID 会通过 “chat_shared” 服务消息发送给机器人，仅在私聊中可用。
	RequestContact  bool                        `json:"request_contact,omitempty"`  // 可选。如果为 true，按下按钮时会将用户的电话号码作为联系人发送，仅在私聊中可用。
	RequestLocation bool                        `json:"request_location,omitempty"` // 可选。如果为 true，按下按钮时会发送用户的当前位置，仅在私聊中可用。
	RequestPoll     *KeyboardButtonPollType     `json:"request_poll,omitempty"`     // 可选。按下按钮时，用户会被提示创建一个投票并发送给机器人，仅在私聊中可用。
	WebApp          *WebAppInfo                 `json:"web_app,omitempty"`          // 可选。按下按钮时会启动指定的 Web App，该应用可以发送 “web_app_data” 服务消息，仅在私聊中可用。
}
