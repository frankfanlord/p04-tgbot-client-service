package structure

// BusinessBotRights 表示企业机器人所拥有的权限
type BusinessBotRights struct {
	CanReply                   bool `json:"can_reply"`                      // 如果为 true，表示机器人可以在最近 24 小时内有传入消息的私聊中发送和编辑消息
	CanReadMessages            bool `json:"can_read_messages"`              // 如果为 true，表示机器人可以将收到的私聊消息标记为已读
	CanDeleteSentMessages      bool `json:"can_delete_sent_messages"`       // 如果为 true，表示机器人可以删除它自己发送的消息
	CanDeleteAllMessages       bool `json:"can_delete_all_messages"`        // 如果为 true，表示机器人可以删除受管聊天中的所有私聊消息
	CanEditName                bool `json:"can_edit_name"`                  // 如果为 true，表示机器人可以编辑企业账户的名字
	CanEditBio                 bool `json:"can_edit_bio"`                   // 如果为 true，表示机器人可以编辑企业账户的简介
	CanEditProfilePhoto        bool `json:"can_edit_profile_photo"`         // 如果为 true，表示机器人可以编辑企业账户的头像
	CanEditUsername            bool `json:"can_edit_username"`              // 如果为 true，表示机器人可以编辑企业账户的用户名
	CanChangeGiftSettings      bool `json:"can_change_gift_settings"`       // 如果为 true，表示机器人可以修改与礼物相关的隐私设置
	CanViewGiftsAndStars       bool `json:"can_view_gifts_and_stars"`       // 如果为 true，表示机器人可以查看礼物和企业账户拥有的 Telegram 星星数量
	CanConvertGiftsToStars     bool `json:"can_convert_gifts_to_stars"`     // 如果为 true，表示机器人可以将礼物转换为 Telegram 星星
	CanTransferAndUpgradeGifts bool `json:"can_transfer_and_upgrade_gifts"` // 如果为 true，表示机器人可以转移和升级企业账户所拥有的礼物
	CanTransferStars           bool `json:"can_transfer_stars"`             // 如果为 true，表示机器人可以转移 Telegram 星星或用来升级/转移礼物
	CanManageStories           bool `json:"can_manage_stories"`             // 如果为 true，表示机器人可以代表企业账户发布、编辑和删除故事
}
