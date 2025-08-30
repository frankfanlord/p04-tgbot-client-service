package structure

// Message 该对象代表一条消息。
type Message struct {
	MessageID                     int                            `json:"message_id"`                        // 聊天中的唯一消息标识符
	MessageThreadID               int                            `json:"message_thread_id"`                 // （可选）消息所属线程的唯一标识符，仅适用于超级群组
	From                          *User                          `json:"from"`                              // （可选）消息发送者（频道中可能为空）
	SenderChat                    *Chat                          `json:"sender_chat"`                       // （可选）代表聊天发送的消息的发送者
	SenderBoostCount              int                            `json:"sender_boost_count"`                // （可选）如果发送者为聊天进行了加成，加成数量
	SenderBusinessBot             *User                          `json:"sender_business_bot"`               // （可选）代表商业账户发送消息的机器人
	Date                          int                            `json:"date"`                              // 消息发送的 Unix 时间戳
	BusinessConnectionID          string                         `json:"business_connection_id"`            // （可选）接收消息的商业连接 ID
	Chat                          *Chat                          `json:"chat"`                              // 消息所属的聊天
	ForwardOrigin                 *MessageOrigin                 `json:"forward_origin"`                    // （可选）转发消息的原始信息
	IsTopicMessage                bool                           `json:"is_topic_message"`                  // （可选）是否为发送到论坛主题的消息
	IsAutomaticForward            bool                           `json:"is_automatic_forward"`              // （可选）是否为自动转发的频道消息
	ReplyToMessage                *Message                       `json:"reply_to_message"`                  // （可选）回复的原始消息（不含嵌套 reply_to_message）
	ExternalReply                 *ExternalReplyInfo             `json:"external_reply"`                    // （可选）引用来自其他聊天或主题的消息信息
	Quote                         *TextQuote                     `json:"quote"`                             // （可选）引用消息的一部分文本
	ReplyToStory                  *Story                         `json:"reply_to_story"`                    // （可选）回复的原始 Story
	ViaBot                        *User                          `json:"via_bot"`                           // （可选）通过哪个机器人发送的消息
	EditDate                      int                            `json:"edit_date"`                         // （可选）最后编辑时间
	HasProtectedContent           bool                           `json:"has_protected_content"`             // （可选）消息是否不可转发
	IsFromOffline                 bool                           `json:"is_from_offline"`                   // （可选）是否为隐式操作发送的离线消息
	MediaGroupID                  string                         `json:"media_group_id"`                    // （可选）媒体消息组的唯一标识符
	AuthorSignature               string                         `json:"author_signature"`                  // （可选）作者签名或匿名管理员自定义头衔
	PaidStarCount                 int                            `json:"paid_star_count"`                   // （可选）发送者为发送此消息支付的 Telegram Stars 数量
	Text                          string                         `json:"text"`                              // （可选）文本消息内容
	Entities                      []*MessageEntity               `json:"entities"`                          // （可选）消息中的特殊实体，如用户名、链接、命令等
	LinkPreviewOptions            *LinkPreviewOptions            `json:"link_preview_options"`              // （可选）链接预览设置
	EffectID                      string                         `json:"effect_id"`                         // （可选）消息使用的特效 ID
	Animation                     *Animation                     `json:"animation"`                         // （可选）动画消息内容
	Audio                         *Audio                         `json:"audio"`                             // （可选）音频消息内容
	Document                      *Document                      `json:"document"`                          // （可选）文件消息内容
	PaidMedia                     *PaidMediaInfo                 `json:"paid_media"`                        // （可选）付费媒体信息
	Photo                         []*PhotoSize                   `json:"photo"`                             // （可选）图片消息内容（多种尺寸）
	Sticker                       *Sticker                       `json:"sticker"`                           // （可选）贴纸消息内容
	Story                         *Story                         `json:"story"`                             // （可选）转发的 Story
	Video                         *Video                         `json:"video"`                             // （可选）视频消息内容
	VideoNote                     *VideoNote                     `json:"video_note"`                        // （可选）视频笔记
	Voice                         *Voice                         `json:"voice"`                             // （可选）语音消息内容
	Caption                       string                         `json:"caption"`                           // （可选）多媒体消息的说明
	CaptionEntities               []*MessageEntity               `json:"caption_entities"`                  // （可选）说明中的特殊实体
	ShowCaptionAboveMedia         bool                           `json:"show_caption_above_media"`          // （可选）说明是否显示在媒体上方
	HasMediaSpoiler               bool                           `json:"has_media_spoiler"`                 // （可选）媒体是否包含剧透动画
	Contact                       *Contact                       `json:"contact"`                           // （可选）共享的联系人信息
	Dice                          *Dice                          `json:"dice"`                              // （可选）掷骰子结果
	Game                          *Game                          `json:"game"`                              // （可选）游戏消息
	Poll                          *Poll                          `json:"poll"`                              // （可选）投票消息
	Venue                         *Venue                         `json:"venue"`                             // （可选）地点信息
	Location                      *Location                      `json:"location"`                          // （可选）地理位置共享
	NewChatMembers                []*User                        `json:"new_chat_members"`                  // （可选）新加入的成员列表
	LeftChatMember                *User                          `json:"left_chat_member"`                  // （可选）被移除的成员
	NewChatTitle                  string                         `json:"new_chat_title"`                    // （可选）新聊天标题
	NewChatPhoto                  []*PhotoSize                   `json:"new_chat_photo"`                    // （可选）新聊天头像
	DeleteChatPhoto               bool                           `json:"delete_chat_photo"`                 // （可选）聊天头像已删除
	GroupChatCreated              bool                           `json:"group_chat_created"`                // （可选）群组已创建
	SupergroupChatCreated         bool                           `json:"supergroup_chat_created"`           // （可选）超级群组已创建
	ChannelChatCreated            bool                           `json:"channel_chat_created"`              // （可选）频道已创建
	MessageAutoDeleteTimerChanged *MessageAutoDeleteTimerChanged `json:"message_auto_delete_timer_changed"` // （可选）自动删除定时器设置变更
	MigrateToChatID               int                            `json:"migrate_to_chat_id"`                // （可选）群组迁移至超级群组的 ID
	MigrateFromChatID             int                            `json:"migrate_from_chat_id"`              // （可选）超级群组迁移前的原群组 ID
	PinnedMessage                 *MaybeInaccessibleMessage      `json:"pinned_message"`                    // （可选）被置顶的消息
	Invoice                       *Invoice                       `json:"invoice"`                           // （可选）支付发票
	SuccessfulPayment             *SuccessfulPayment             `json:"successful_payment"`                // （可选）支付成功信息
	RefundedPayment               *RefundedPayment               `json:"refunded_payment"`                  // （可选）支付退款信息
	UsersShared                   *UsersShared                   `json:"users_shared"`                      // （可选）共享的用户信息
	ChatShared                    *ChatShared                    `json:"chat_shared"`                       // （可选）共享的聊天信息
	Gift                          *GiftInfo                      `json:"gift"`                              // （可选）普通礼物信息
	UniqueGift                    *UniqueGiftInfo                `json:"unique_gift"`                       // （可选）独特礼物信息
	ConnectedWebsite              string                         `json:"connected_website"`                 // （可选）用户登录的网站域名
	WriteAccessAllowed            *WriteAccessAllowed            `json:"write_access_allowed"`              // （可选）用户允许机器人发消息
	PassportData                  *PassportData                  `json:"passport_data"`                     // （可选）Telegram Passport 数据
	ProximityAlertTriggered       *ProximityAlertTriggered       `json:"proximity_alert_triggered"`         // （可选）附近提醒信息
	BoostAdded                    *ChatBoostAdded                `json:"boost_added"`                       // （可选）用户为聊天增加了加成
	ChatBackgroundSet             *ChatBackground                `json:"chat_background_set"`               // （可选）聊天背景设置信息
	ForumTopicCreated             *ForumTopicCreated             `json:"forum_topic_created"`               // （可选）创建论坛主题
	ForumTopicEdited              *ForumTopicEdited              `json:"forum_topic_edited"`                // （可选）编辑论坛主题
	ForumTopicClosed              *ForumTopicClosed              `json:"forum_topic_closed"`                // （可选）关闭论坛主题
	ForumTopicReopened            *ForumTopicReopened            `json:"forum_topic_reopened"`              // （可选）重新打开论坛主题
	GeneralForumTopicHidden       *GeneralForumTopicHidden       `json:"general_forum_topic_hidden"`        // （可选）“General” 论坛主题已隐藏
	GeneralForumTopicUnhidden     *GeneralForumTopicUnhidden     `json:"general_forum_topic_unhidden"`      // （可选）“General” 论坛主题已取消隐藏
	GiveawayCreated               *GiveawayCreated               `json:"giveaway_created"`                  // （可选）计划抽奖已创建
	Giveaway                      *Giveaway                      `json:"giveaway"`                          // （可选）计划抽奖消息
	GiveawayWinners               *GiveawayWinners               `json:"giveaway_winners"`                  // （可选）有公开获奖者的抽奖已完成
	GiveawayCompleted             *GiveawayCompleted             `json:"giveaway_completed"`                // （可选）无公开获奖者的抽奖已完成
	PaidMessagePriceChanged       *PaidMessagePriceChanged       `json:"paid_message_price_changed"`        // （可选）付费消息价格已变更
	VideoChatScheduled            *VideoChatScheduled            `json:"video_chat_scheduled"`              // （可选）视频聊天计划
	VideoChatStarted              *VideoChatStarted              `json:"video_chat_started"`                // （可选）视频聊天已开始
	VideoChatEnded                *VideoChatEnded                `json:"video_chat_ended"`                  // （可选）视频聊天已结束
	VideoChatParticipantsInvited  *VideoChatParticipantsInvited  `json:"video_chat_participants_invited"`   // （可选）被邀请加入视频聊天的用户
	WebAppData                    *WebAppData                    `json:"web_app_data"`                      // （可选）Web 应用发送的数据
	ReplyMarkup                   *InlineKeyboardMarkup          `json:"reply_markup"`                      // （可选）内联键盘标记
}
