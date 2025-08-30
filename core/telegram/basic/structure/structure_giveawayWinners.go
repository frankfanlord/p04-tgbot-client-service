package structure

// GiveawayWinners 表示一个关于抽奖完成及公开获奖者的消息
type GiveawayWinners struct {
	Chat                          *Chat   `json:"chat"`                             // 发起抽奖的聊天对象
	GiveawayMessageID             int     `json:"giveaway_message_id"`              // 聊天中用于抽奖的消息ID
	WinnersSelectionDate          int     `json:"winners_selection_date"`           // 抽奖结果公布的时间（Unix时间戳）
	WinnerCount                   int     `json:"winner_count"`                     // 抽奖中的获奖者总数
	Winners                       []*User `json:"winners"`                          // 最多100个获奖用户的列表
	AdditionalChatCount           int     `json:"additional_chat_count"`            // 可选。用户需加入的其他聊天数量以符合资格
	PrizeStarCount                int     `json:"prize_star_count"`                 // 可选。被分配给获奖者的 Telegram 星星数（仅适用于星星抽奖）
	PremiumSubscriptionMonthCount int     `json:"premium_subscription_month_count"` // 可选。抽中 Telegram 高级订阅奖项的可用月数（仅适用于高级订阅抽奖）
	UnclaimedPrizeCount           int     `json:"unclaimed_prize_count"`            // 可选。未发出的奖品数量
	OnlyNewMembers                bool    `json:"only_new_members"`                 // 可选。是否仅新加入聊天的用户有资格获奖
	WasRefunded                   bool    `json:"was_refunded"`                     // 可选。是否因付款退款而取消抽奖
	PrizeDescription              string  `json:"prize_description"`                // 可选。额外奖品的描述
}
