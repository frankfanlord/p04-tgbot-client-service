package structure

// KeyboardButtonRequestChat 定义请求合适聊天的条件
type KeyboardButtonRequestChat struct {
	RequestID               int                      `json:"request_id"`                          // 请求的唯一标识符，32位有符号整数，在 ChatShared 对象中返回
	ChatIsChannel           bool                     `json:"chat_is_channel"`                     // 是否请求频道聊天，true 表示频道，false 表示群组或超级群组
	ChatIsForum             bool                     `json:"chat_is_forum,omitempty"`             // （可选）是否请求论坛超级群组，true 表示论坛，false 表示非论坛
	ChatHasUsername         bool                     `json:"chat_has_username,omitempty"`         // （可选）是否请求带有用户名的超级群组或频道
	ChatIsCreated           bool                     `json:"chat_is_created,omitempty"`           // （可选）是否请求由用户创建的聊天
	UserAdministratorRights *ChatAdministratorRights `json:"user_administrator_rights,omitempty"` // （可选）用户在聊天中需要的管理员权限，必须是 bot_administrator_rights 的超集
	BotAdministratorRights  *ChatAdministratorRights `json:"bot_administrator_rights,omitempty"`  // （可选）机器人在聊天中需要的管理员权限，必须是 user_administrator_rights 的子集
	BotIsMember             bool                     `json:"bot_is_member,omitempty"`             // （可选）是否请求机器人已是该聊天成员
	RequestTitle            bool                     `json:"request_title,omitempty"`             // （可选）是否请求聊天标题
	RequestUsername         bool                     `json:"request_username,omitempty"`          // （可选）是否请求聊天用户名
	RequestPhoto            bool                     `json:"request_photo,omitempty"`             // （可选）是否请求聊天头像
}
