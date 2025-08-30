package structure

// UsersShared 包含使用 KeyboardButtonRequestUsers 按钮与机器人共享的用户信息
type UsersShared struct {
	RequestID int           `json:"request_id"` // 请求的标识符
	Users     []*SharedUser `json:"users"`      // 与机器人共享的用户信息列表
}
