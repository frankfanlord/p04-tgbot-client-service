package structure

// AcceptedGiftTypes 描述可赠送给用户或群组的礼物类型
type AcceptedGiftTypes struct {
	UnlimitedGifts      bool `json:"unlimited_gifts"`      // 如果为 true，表示接受无限常规礼物
	LimitedGifts        bool `json:"limited_gifts"`        // 如果为 true，表示接受有限常规礼物
	UniqueGifts         bool `json:"unique_gifts"`         // 如果为 true，表示接受独特礼物或可免费升级为独特的礼物
	PremiumSubscription bool `json:"premium_subscription"` // 如果为 true，表示接受 Telegram Premium 订阅作为礼物
}
