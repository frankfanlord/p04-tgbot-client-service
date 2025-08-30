package structure

// GiveawayCompleted 表示一个关于完成抽奖的服务消息（无公开获奖者）。
type GiveawayCompleted struct {
	WinnerCount         int      `json:"winner_count"`          // 抽奖中的获奖人数
	UnclaimedPrizeCount int      `json:"unclaimed_prize_count"` // 未分发奖品数量（可选）
	GiveawayMessage     *Message `json:"giveaway_message"`      // 与已完成抽奖相关的消息（可选，若未被删除）
	IsStarGiveaway      bool     `json:"is_star_giveaway"`      // 是否为 Telegram Star 抽奖（否则为 Telegram Premium 抽奖）（可选）
}
