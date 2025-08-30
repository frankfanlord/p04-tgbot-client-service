package structure

// UniqueGiftInfo 描述一个关于唯一礼物的服务消息，可能是发送或接收的礼物。
type UniqueGiftInfo struct {
	Gift              *UniqueGift `json:"gift"`                // 礼物信息
	Origin            string      `json:"origin"`              // 礼物的来源，目前是 "upgrade" 或 "transfer"
	OwnedGiftID       string      `json:"owned_gift_id"`       // （可选）机器人的收到礼物的唯一标识，仅在代表商业账户收到礼物时存在
	TransferStarCount int         `json:"transfer_star_count"` // （可选）转移该礼物需要支付的 Telegram 星星数；如果无法转移则不会返回该字段
}
