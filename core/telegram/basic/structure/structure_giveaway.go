package structure

// Giveaway 表示关于预定抽奖活动的信息
type Giveaway struct {
	Chats                         []*Chat  `json:"chats"`                                      // 用户必须加入的聊天列表，才可参与抽奖
	WinnersSelectionDate          int      `json:"winners_selection_date"`                     // 抽奖中奖者选择的时间点（Unix 时间戳）
	WinnerCount                   int      `json:"winner_count"`                               // 要选出的中奖用户数量
	OnlyNewMembers                bool     `json:"only_new_members,omitempty"`                 // 可选，仅加入聊天后才有资格中奖的用户
	HasPublicWinners              bool     `json:"has_public_winners,omitempty"`               // 可选，中奖用户名单是否对所有人可见
	PrizeDescription              string   `json:"prize_description,omitempty"`                // 可选，额外奖品的描述
	CountryCodes                  []string `json:"country_codes,omitempty"`                    // 可选，符合资格的用户所属国家的两位国家码；为空表示所有用户都可参与
	PrizeStarCount                int      `json:"prize_star_count,omitempty"`                 // 可选，Telegram Stars 奖励分配给中奖者的数量，仅限 Telegram Star 抽奖
	PremiumSubscriptionMonthCount int      `json:"premium_subscription_month_count,omitempty"` // 可选，Telegram Premium 订阅奖品的有效月份，仅限 Premium 抽奖
}
