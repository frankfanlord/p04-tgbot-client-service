package structure

// ChatBoostSource 表示聊天加速的来源
type ChatBoostSource struct {
	Source            string `json:"source"`              // 加速来源，可为 "premium"、"gift_code" 或 "giveaway"
	User              *User  `json:"user"`                // 加速该聊天的用户，适用于 premium 和 gift_code，或 giveaway 的中奖用户
	GiveawayMessageID int    `json:"giveaway_message_id"` // 抽奖消息的 ID，适用于 giveaway，若消息尚未发送则为 0
	PrizeStarCount    int    `json:"prize_star_count"`    // Telegram Star 抽奖中的星星数，仅适用于 star 类型 giveaway
	IsUnclaimed       bool   `json:"is_unclaimed"`        // 抽奖已完成但无人中奖时为 true，仅适用于 giveaway
}
