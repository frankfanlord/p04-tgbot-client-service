package structure

// ChatFullInfo 该对象包含有关聊天的完整信息。
type ChatFullInfo struct {
	ID                             int                   `json:"id"`                                      // 聊天的唯一标识符，最大 52 位，建议使用 int 存储
	Type                           string                `json:"type"`                                    // 聊天类型，可能为 “private”, “group”, “supergroup” 或 “channel”
	Title                          string                `json:"title"`                                   // 可选，标题，仅适用于群组、超级群组和频道
	Username                       string                `json:"username"`                                // 可选，用户名，适用于私人聊天、超级群组和频道
	FirstName                      string                `json:"first_name"`                              // 可选，私人聊天中对方的名字
	LastName                       string                `json:"last_name"`                               // 可选，私人聊天中对方的姓氏
	IsForum                        bool                  `json:"is_forum"`                                // 可选，是否是启用话题的论坛型超级群组
	AccentColorID                  int                   `json:"accent_color_id"`                         // 聊天名称和背景的主题颜色 ID
	MaxReactionCount               int                   `json:"max_reaction_count"`                      // 聊天中每条消息可设置的最大反应数
	Photo                          *ChatPhoto            `json:"photo"`                                   // 可选，聊天照片
	ActiveUsernames                []string              `json:"active_usernames"`                        // 可选，所有活跃用户名（非空时）
	Birthdate                      *Birthdate            `json:"birthdate"`                               // 可选，私人聊天中对方的出生日期
	BusinessIntro                  *BusinessIntro        `json:"business_intro"`                          // 可选，私人聊天中企业账户的简介
	BusinessLocation               *BusinessLocation     `json:"business_location"`                       // 可选，企业账户的地理位置
	BusinessOpeningHours           *BusinessOpeningHours `json:"business_opening_hours"`                  // 可选，企业账户的营业时间
	PersonalChat                   *Chat                 `json:"personal_chat"`                           // 可选，用户的个人频道
	AvailableReactions             []*ReactionType       `json:"available_reactions"`                     // 可选，允许的表情反应类型
	BackgroundCustomEmojiID        string                `json:"background_custom_emoji_id"`              // 可选，用于回复标题和预览背景的自定义 emoji ID
	ProfileAccentColorID           int                   `json:"profile_accent_color_id"`                 // 可选，聊天个人资料背景的主题颜色 ID
	ProfileBackgroundCustomEmojiID string                `json:"profile_background_custom_emoji_id"`      // 可选，聊天背景所选 emoji 的自定义 ID
	EmojiStatusCustomEmojiID       string                `json:"emoji_status_custom_emoji_id"`            // 可选，聊天状态的自定义 emoji ID
	EmojiStatusExpirationDate      int                   `json:"emoji_status_expiration_date"`            // 可选，emoji 状态过期时间（Unix 时间戳）
	Bio                            string                `json:"bio"`                                     // 可选，私人聊天中对方的个人简介
	HasPrivateForwards             bool                  `json:"has_private_forwards"`                    // 可选，是否只能在和用户对话中使用 tg://user?id=<user_id> 链接
	HasRestrictedVoiceAndVideo     bool                  `json:"has_restricted_voice_and_video_messages"` // 可选，是否禁止发送语音和视频消息
	JoinToSendMessages             bool                  `json:"join_to_send_messages"`                   // 可选，是否必须加入超级群组才能发送消息
	JoinByRequest                  bool                  `json:"join_by_request"`                         // 可选，是否需要管理员审批才能加入群组
	Description                    string                `json:"description"`                             // 可选，群组或频道的描述
	InviteLink                     string                `json:"invite_link"`                             // 可选，主邀请链接
	PinnedMessage                  *Message              `json:"pinned_message"`                          // 可选，最近置顶的消息
	Permissions                    *ChatPermissions      `json:"permissions"`                             // 可选，群组默认权限
	AcceptedGiftTypes              *AcceptedGiftTypes    `json:"accepted_gift_types"`                     // 接受的礼物类型信息
	CanSendPaidMedia               bool                  `json:"can_send_paid_media"`                     // 可选，是否允许发送或转发付费媒体（仅频道）
	SlowModeDelay                  int                   `json:"slow_mode_delay"`                         // 可选，慢速模式下允许发送的最小时间间隔（秒）
	UnrestrictBoostCount           int                   `json:"unrestrict_boost_count"`                  // 可选，绕过限制所需的最少 boost 数
	MessageAutoDeleteTime          int                   `json:"message_auto_delete_time"`                // 可选，聊天消息自动删除的时间（秒）
	HasAggressiveAntiSpamEnabled   bool                  `json:"has_aggressive_anti_spam_enabled"`        // 可选，是否启用了强力反垃圾检查（仅管理员可见）
	HasHiddenMembers               bool                  `json:"has_hidden_members"`                      // 可选，是否隐藏非管理员成员
	HasProtectedContent            bool                  `json:"has_protected_content"`                   // 可选，是否禁止消息被转发
	HasVisibleHistory              bool                  `json:"has_visible_history"`                     // 可选，新成员是否能查看历史消息（仅管理员可见）
	StickerSetName                 string                `json:"sticker_set_name"`                        // 可选，群组贴纸集名称
	CanSetStickerSet               bool                  `json:"can_set_sticker_set"`                     // 可选，机器人是否可以设置贴纸集
	CustomEmojiStickerSetName      string                `json:"custom_emoji_sticker_set_name"`           // 可选，群组自定义 emoji 贴纸集名称
	LinkedChatID                   int                   `json:"linked_chat_id"`                          // 可选，关联聊天的唯一标识符（52 位以内）
	Location                       *ChatLocation         `json:"location"`                                // 可选，超级群组所关联的位置
}
