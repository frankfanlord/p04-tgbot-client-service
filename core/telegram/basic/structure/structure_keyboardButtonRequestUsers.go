package structure

// KeyboardButtonRequestUsers 用户请求条件对象。
// 当按钮被点击时，所选用户的信息将会共享给 Bot。
type KeyboardButtonRequestUsers struct {
	RequestID       int  `json:"request_id"`       // 请求 ID，32 位有符号整数，将在 UsersShared 对象中返回。消息中必须唯一。
	UserIsBot       bool `json:"user_is_bot"`      // （可选）是否请求机器人用户；为 true 表示请求机器人，为 false 表示请求普通用户，未指定则不限制。
	UserIsPremium   bool `json:"user_is_premium"`  // （可选）是否请求高级用户；为 true 表示请求高级用户，为 false 表示请求非高级用户，未指定则不限制。
	MaxQuantity     int  `json:"max_quantity"`     // （可选）最多选择的用户数量，范围 1-10，默认为 1。
	RequestName     bool `json:"request_name"`     // （可选）是否请求用户的名字（名和姓）。
	RequestUsername bool `json:"request_username"` // （可选）是否请求用户的用户名。
	RequestPhoto    bool `json:"request_photo"`    // （可选）是否请求用户的头像。
}
