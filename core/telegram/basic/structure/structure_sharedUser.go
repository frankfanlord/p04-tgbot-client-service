package structure

// SharedUser 包含通过 KeyboardButtonRequestUsers 按钮与机器人共享的用户信息。
type SharedUser struct {
	UserID    int64        `json:"user_id"`    // 用户 ID。可能大于 32 位，应使用 64 位整数或 float64 存储。
	FirstName string       `json:"first_name"` // （可选）用户的名字，如果机器人请求了名字。
	LastName  string       `json:"last_name"`  // （可选）用户的姓氏，如果机器人请求了名字。
	Username  string       `json:"username"`   // （可选）用户的用户名，如果机器人请求了用户名。
	Photo     []*PhotoSize `json:"photo"`      // （可选）用户头像的不同尺寸集合，如果机器人请求了头像。
}
