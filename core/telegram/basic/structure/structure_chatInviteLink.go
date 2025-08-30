package structure

// ChatInviteLink 表示一个聊天邀请链接
type ChatInviteLink struct {
	InviteLink              string `json:"invite_link"`                // 邀请链接。如果链接由其他管理员创建，则链接的后半部分将被替换为“...”
	Creator                 *User  `json:"creator"`                    // 链接的创建者
	CreatesJoinRequest      bool   `json:"creates_join_request"`       // 如果通过该链接加入聊天的用户需要管理员批准，则为 true
	IsPrimary               bool   `json:"is_primary"`                 // 如果该链接是主链接，则为 true
	IsRevoked               bool   `json:"is_revoked"`                 // 如果该链接已被撤销，则为 true
	Name                    string `json:"name"`                       // 可选。邀请链接名称
	ExpireDate              int    `json:"expire_date"`                // 可选。链接过期的时间点（Unix 时间戳）
	MemberLimit             int    `json:"member_limit"`               // 可选。通过此链接加入后，聊天的最大成员数量（1-99999）
	PendingJoinRequestCount int    `json:"pending_join_request_count"` // 可选。通过此链接创建的待处理加入请求数量
	SubscriptionPeriod      int    `json:"subscription_period"`        // 可选。订阅在下一次付款前持续的秒数
	SubscriptionPrice       int    `json:"subscription_price"`         // 可选。用户最初和每次续订时必须支付的 Telegram Stars 数量
}
