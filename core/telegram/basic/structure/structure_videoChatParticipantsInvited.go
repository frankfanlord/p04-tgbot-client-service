package structure

// VideoChatParticipantsInvited 表示关于被邀请加入视频聊天的新成员的服务消息
type VideoChatParticipantsInvited struct {
	Users []*User `json:"users"` // 被邀请加入视频聊天的新成员
}
