package structure

// ChatMember 表示一个聊天成员的通用信息，兼容六种子类型
type ChatMember struct {
	Status string `json:"status"` // 成员状态，如 creator、administrator、member、restricted、left、kicked
	User   *User  `json:"user"`   // 用户信息

	// 所有角色可能存在的通用字段
	IsAnonymous bool   `json:"is_anonymous"` // （可选）是否匿名
	CustomTitle string `json:"custom_title"` // （可选）自定义头衔
	UntilDate   int    `json:"until_date"`   // （可选）封禁/订阅/限制到期时间（Unix 时间戳）

	// 管理员权限字段
	CanBeEdited         bool `json:"can_be_edited"`          // 是否允许编辑此管理员
	CanManageChat       bool `json:"can_manage_chat"`        // 是否可以管理聊天
	CanDeleteMessages   bool `json:"can_delete_messages"`    // 是否可以删除消息
	CanManageVideoChats bool `json:"can_manage_video_chats"` // 是否可以管理视频聊天
	CanRestrictMembers  bool `json:"can_restrict_members"`   // 是否可以封禁成员
	CanPromoteMembers   bool `json:"can_promote_members"`    // 是否可以提升成员
	CanChangeInfo       bool `json:"can_change_info"`        // 是否可以更改群信息
	CanInviteUsers      bool `json:"can_invite_users"`       // 是否可以邀请用户
	CanPostStories      bool `json:"can_post_stories"`       // 是否可以发布故事
	CanEditStories      bool `json:"can_edit_stories"`       // 是否可以编辑故事
	CanDeleteStories    bool `json:"can_delete_stories"`     // 是否可以删除故事
	CanPostMessages     bool `json:"can_post_messages"`      // （可选）是否可以发消息，仅限频道
	CanEditMessages     bool `json:"can_edit_messages"`      // （可选）是否可以编辑消息，仅限频道
	CanPinMessages      bool `json:"can_pin_messages"`       // （可选）是否可以置顶消息
	CanManageTopics     bool `json:"can_manage_topics"`      // （可选）是否可以管理论坛话题

	// 受限成员的权限控制字段
	IsMember              bool `json:"is_member"`                 // 当前是否为成员
	CanSendMessages       bool `json:"can_send_messages"`         // 是否可以发送普通消息
	CanSendAudios         bool `json:"can_send_audios"`           // 是否可以发送音频
	CanSendDocuments      bool `json:"can_send_documents"`        // 是否可以发送文档
	CanSendPhotos         bool `json:"can_send_photos"`           // 是否可以发送照片
	CanSendVideos         bool `json:"can_send_videos"`           // 是否可以发送视频
	CanSendVideoNotes     bool `json:"can_send_video_notes"`      // 是否可以发送视频便笺
	CanSendVoiceNotes     bool `json:"can_send_voice_notes"`      // 是否可以发送语音便笺
	CanSendPolls          bool `json:"can_send_polls"`            // 是否可以发送投票
	CanSendOtherMessages  bool `json:"can_send_other_messages"`   // 是否可以发送动画、游戏、贴纸等
	CanAddWebPagePreviews bool `json:"can_add_web_page_previews"` // 是否可以添加网页预览
}
