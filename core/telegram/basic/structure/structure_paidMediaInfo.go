package structure

type PaidMediaInfo struct {
	StarCount int          `json:"star_count"` // 需要支付的 Telegram 星星数量
	PaidMedia []*PaidMedia `json:"paid_media"` // 有关付费媒体的信息
}
