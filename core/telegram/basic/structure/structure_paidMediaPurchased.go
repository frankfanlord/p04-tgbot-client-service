package structure

// PaidMediaPurchased 包含付费媒体购买的信息
type PaidMediaPurchased struct {
	From             *User  `json:"from"`               // 购买媒体的用户
	PaidMediaPayload string `json:"paid_media_payload"` // 由机器人指定的付费媒体负载
}
