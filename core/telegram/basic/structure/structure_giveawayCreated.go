package structure

// GiveawayCreated 表示一个关于创建定时抽奖的服务消息
type GiveawayCreated struct {
	PrizeStarCount int `json:"prize_star_count"` // Telegram 星星数量（仅用于 Telegram 星星抽奖，可选）
}
