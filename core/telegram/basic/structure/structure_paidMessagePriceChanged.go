package structure

// PaidMessagePriceChanged 描述超级群组中付费消息价格变更的服务消息
type PaidMessagePriceChanged struct {
	// PaidMessageStarCount 非管理员用户每发送一条消息所需支付的新 Telegram 星星数
	PaidMessageStarCount int `json:"paid_message_star_count"`
}
