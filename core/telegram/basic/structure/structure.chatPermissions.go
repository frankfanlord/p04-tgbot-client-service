package structure

// ChatPermissions 描述非管理员用户在群组中允许执行的操作
type ChatPermissions struct {
	CanSendMessages       bool  `json:"can_send_messages"`         // 如果为 true，表示允许发送文本消息、联系人、赠品、赠品获奖者、发票、位置和场地
	CanSendAudios         bool  `json:"can_send_audios"`           // 如果为 true，表示允许发送音频
	CanSendDocuments      bool  `json:"can_send_documents"`        // 如果为 true，表示允许发送文档
	CanSendPhotos         bool  `json:"can_send_photos"`           // 如果为 true，表示允许发送照片
	CanSendVideos         bool  `json:"can_send_videos"`           // 如果为 true，表示允许发送视频
	CanSendVideoNotes     bool  `json:"can_send_video_notes"`      // 如果为 true，表示允许发送视频便笺
	CanSendVoiceNotes     bool  `json:"can_send_voice_notes"`      // 如果为 true，表示允许发送语音便笺
	CanSendPolls          bool  `json:"can_send_polls"`            // 如果为 true，表示允许发送投票
	CanSendOtherMessages  bool  `json:"can_send_other_messages"`   // 如果为 true，表示允许发送动画、游戏、贴纸以及使用内联机器人
	CanAddWebPagePreviews bool  `json:"can_add_web_page_previews"` // 如果为 true，表示允许为消息添加网页预览
	CanChangeInfo         bool  `json:"can_change_info"`           // 如果为 true，表示允许更改聊天标题、照片和其他设置（在公开超级群组中会被忽略）
	CanInviteUsers        bool  `json:"can_invite_users"`          // 如果为 true，表示允许邀请新用户加入聊天
	CanPinMessages        bool  `json:"can_pin_messages"`          // 如果为 true，表示允许置顶消息（在公开超级群组中会被忽略）
	CanManageTopics       *bool `json:"can_manage_topics"`         // 如果为 true，表示允许创建论坛主题，省略时默认等于 CanPinMessages
}
