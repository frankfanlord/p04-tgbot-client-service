package structure

// Update 表示传入的更新对象
type Update struct {
	UpdateID                int                          `json:"update_id"`                 // 更新的唯一标识符
	Message                 *Message                     `json:"message"`                   // 新的传入消息（可能是文本、照片、贴纸等）
	EditedMessage           *Message                     `json:"edited_message"`            // 已知消息的新版本，已被编辑
	ChannelPost             *Message                     `json:"channel_post"`              // 新的传入频道消息（可能是文本、照片、贴纸等）
	EditedChannelPost       *Message                     `json:"edited_channel_post"`       // 已知频道消息的新版本，已被编辑
	BusinessConnection      *BusinessConnection          `json:"business_connection"`       // 与企业账户连接或断开连接，或用户编辑了与机器人的连接
	BusinessMessage         *Message                     `json:"business_message"`          // 来自企业账户的新消息
	EditedBusinessMessage   *Message                     `json:"edited_business_message"`   // 来自企业账户的消息被编辑
	DeletedBusinessMessages *BusinessMessagesDeleted     `json:"deleted_business_messages"` // 企业账户中的消息被删除
	MessageReaction         *MessageReactionUpdated      `json:"message_reaction"`          // 用户更改了对消息的反应
	MessageReactionCount    *MessageReactionCountUpdated `json:"message_reaction_count"`    // 匿名反应的数量发生变化
	InlineQuery             *InlineQuery                 `json:"inline_query"`              // 新的内联查询
	ChosenInlineResult      *ChosenInlineResult          `json:"chosen_inline_result"`      // 用户选择并发送了某个内联查询的结果
	CallbackQuery           *CallbackQuery               `json:"callback_query"`            // 新的回调查询
	ShippingQuery           *ShippingQuery               `json:"shipping_query"`            // 新的送货查询（用于灵活定价的发票）
	PreCheckoutQuery        *PreCheckoutQuery            `json:"pre_checkout_query"`        // 新的预结账查询，包含完整结账信息
	PurchasedPaidMedia      *PaidMediaPurchased          `json:"purchased_paid_media"`      // 用户购买了带有效载荷的付费媒体内容
	Poll                    *Poll                        `json:"poll"`                      // 新的投票状态（仅手动结束或机器人发起的投票）
	PollAnswer              *PollAnswer                  `json:"poll_answer"`               // 用户在非匿名投票中更改了答案
	MyChatMember            *ChatMemberUpdated           `json:"my_chat_member"`            // 机器人的聊天成员状态发生变化
	ChatMember              *ChatMemberUpdated           `json:"chat_member"`               // 聊天成员状态发生变化
	ChatJoinRequest         *ChatJoinRequest             `json:"chat_join_request"`         // 有人申请加入聊天
	ChatBoost               *ChatBoostUpdated            `json:"chat_boost"`                // 聊天加成被添加或更改
	RemovedChatBoost        *ChatBoostRemoved            `json:"removed_chat_boost"`        // 聊天加成被移除
}
