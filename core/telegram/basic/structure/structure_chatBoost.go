package structure

// ChatBoost 表示一个聊天加速（Boost）的信息
type ChatBoost struct {
	BoostID        string           `json:"boost_id"`        // Boost 的唯一标识符
	AddDate        int              `json:"add_date"`        // 加速添加的时间点（Unix 时间戳）
	ExpirationDate int              `json:"expiration_date"` // 加速自动过期的时间点（Unix 时间戳），除非续费了 Telegram Premium
	Source         *ChatBoostSource `json:"source"`          // 加速的来源
}
