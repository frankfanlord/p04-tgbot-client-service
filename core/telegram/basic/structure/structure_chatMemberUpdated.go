package structure

// ChatMemberUpdated 表示聊天成员状态的变更信息
type ChatMemberUpdated struct {
	Chat                    *Chat           `json:"chat"`                                  // 聊天对象，该用户属于此聊天
	From                    *User           `json:"from"`                                  // 执行动作的用户（导致该变更）
	Date                    int             `json:"date"`                                  // 变更发生的时间（Unix 时间戳）
	OldChatMember           *ChatMember     `json:"old_chat_member"`                       // 聊天成员变更前的信息
	NewChatMember           *ChatMember     `json:"new_chat_member"`                       // 聊天成员变更后的信息
	InviteLink              *ChatInviteLink `json:"invite_link,omitempty"`                 // 可选。用户加入聊天使用的邀请码链接，仅用于通过邀请链接加入的事件
	ViaJoinRequest          bool            `json:"via_join_request,omitempty"`            // 可选。若为 true，表示用户是通过发送加入请求并被管理员批准后加入的，而非使用邀请链接
	ViaChatFolderInviteLink bool            `json:"via_chat_folder_invite_link,omitempty"` // 可选。若为 true，表示用户是通过聊天文件夹的邀请链接加入的
}
