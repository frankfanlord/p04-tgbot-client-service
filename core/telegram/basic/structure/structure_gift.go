package structure

// Gift 表示一个可以由机器人发送的礼物对象。
type Gift struct {
	ID               string   `json:"id"`                 // 礼物的唯一标识符
	Sticker          *Sticker `json:"sticker"`            // 代表该礼物的贴纸
	StarCount        int      `json:"star_count"`         // 发送该贴纸所需的 Telegram 星星数量
	UpgradeStarCount int      `json:"upgrade_star_count"` // 可选。将礼物升级为唯一礼物所需的星星数量
	TotalCount       int      `json:"total_count"`        // 可选。此类礼物可发送的总数，仅适用于限量礼物
	RemainingCount   int      `json:"remaining_count"`    // 可选。此类礼物剩余可发送数量，仅适用于限量礼物
}
