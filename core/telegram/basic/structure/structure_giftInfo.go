package structure

// GiftInfo 描述一个关于发送或接收普通礼物的服务消息。
type GiftInfo struct {
	Gift                    *Gift            `json:"gift"`                       // 礼物的相关信息
	OwnedGiftID             string           `json:"owned_gift_id"`              // 可选。机器人的已接收礼物的唯一标识符，仅在代表商业账号接收礼物时存在
	ConvertStarCount        int              `json:"convert_star_count"`         // 可选。收礼者通过转换礼物可获得的 Telegram 星星数量；如果无法转换则省略
	PrepaidUpgradeStarCount int              `json:"prepaid_upgrade_star_count"` // 可选。送礼者为升级礼物预付的 Telegram 星星数量
	CanBeUpgraded           bool             `json:"can_be_upgraded"`            // 可选。如果该礼物可以升级为唯一礼物，则为 true
	Text                    string           `json:"text"`                       // 可选。附加在礼物上的消息文本
	Entities                []*MessageEntity `json:"entities"`                   // 可选。文本中出现的特殊实体列表
	IsPrivate               bool             `json:"is_private"`                 // 可选。如果只有收礼者能看到送礼者和礼物文本，则为 true，否则所有人都可以看到
}
