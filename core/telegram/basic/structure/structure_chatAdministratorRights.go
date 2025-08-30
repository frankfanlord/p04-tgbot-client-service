package structure

// ChatAdministratorRights 表示聊天中管理员的权限
type ChatAdministratorRights struct {
	IsAnonymous         bool `json:"is_anonymous"`           // 如果为 true，表示该用户在聊天中的存在是匿名的
	CanManageChat       bool `json:"can_manage_chat"`        // 如果为 true，表示管理员可以访问聊天事件日志、获取加成列表、查看隐藏成员、举报垃圾消息和忽略慢速模式
	CanDeleteMessages   bool `json:"can_delete_messages"`    // 如果为 true，表示管理员可以删除其他用户的消息
	CanManageVideoChats bool `json:"can_manage_video_chats"` // 如果为 true，表示管理员可以管理视频聊天
	CanRestrictMembers  bool `json:"can_restrict_members"`   // 如果为 true，表示管理员可以限制、封禁、解封成员，或访问超级群统计信息
	CanPromoteMembers   bool `json:"can_promote_members"`    // 如果为 true，表示管理员可以任命新管理员或降职其任命的管理员
	CanChangeInfo       bool `json:"can_change_info"`        // 如果为 true，表示用户可以更改聊天标题、头像和其他设置
	CanInviteUsers      bool `json:"can_invite_users"`       // 如果为 true，表示用户可以邀请新成员加入聊天
	CanPostStories      bool `json:"can_post_stories"`       // 如果为 true，表示管理员可以发布聊天故事
	CanEditStories      bool `json:"can_edit_stories"`       // 如果为 true，表示管理员可以编辑他人发布的故事、发布到聊天主页、置顶故事、访问故事归档
	CanDeleteStories    bool `json:"can_delete_stories"`     // 如果为 true，表示管理员可以删除他人发布的故事
	CanPostMessages     bool `json:"can_post_messages"`      // 可选字段：如果为 true，表示管理员可以发布频道消息或访问频道统计（仅频道）
	CanEditMessages     bool `json:"can_edit_messages"`      // 可选字段：如果为 true，表示管理员可以编辑他人消息和置顶消息（仅频道）
	CanPinMessages      bool `json:"can_pin_messages"`       // 可选字段：如果为 true，表示用户可以置顶消息（仅群组或超级群）
	CanManageTopics     bool `json:"can_manage_topics"`      // 可选字段：如果为 true，表示用户可以创建、重命名、关闭和重新开启论坛话题（仅超级群）
}
