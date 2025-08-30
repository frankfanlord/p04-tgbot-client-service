package structure

// ChatJoinRequest 表示用户发送到某个聊天的加入请求
type ChatJoinRequest struct {
	Chat       *Chat           `json:"chat"`                  // 加入请求发送到的聊天
	From       *User           `json:"from"`                  // 发送加入请求的用户
	UserChatID int             `json:"user_chat_id"`          // 与发送请求用户的私聊ID，最多52位，可安全用int64存储
	Date       int             `json:"date"`                  // 请求发送的时间，Unix 时间戳
	Bio        string          `json:"bio,omitempty"`         // （可选）用户的简介
	InviteLink *ChatInviteLink `json:"invite_link,omitempty"` // （可选）用户用于发送请求的邀请链接
}
