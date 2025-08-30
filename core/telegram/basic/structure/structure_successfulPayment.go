package structure

// SuccessfulPayment 表示一次成功的支付交易的基本信息。
type SuccessfulPayment struct {
	Currency                string     `json:"currency"`                               // 三位 ISO 4217 货币代码，或者“XTR”表示 Telegram Stars 支付
	TotalAmount             int        `json:"total_amount"`                           // 以货币最小单位表示的总金额（整数，例如 $1.45 表示为 145）
	InvoicePayload          string     `json:"invoice_payload"`                        // 由机器人指定的发票负载数据
	SubscriptionExpireDate  int        `json:"subscription_expiration_date,omitempty"` // 订阅的过期时间（Unix 时间戳，仅用于周期性支付）
	IsRecurring             bool       `json:"is_recurring,omitempty"`                 // 如果为周期性支付则为 true
	IsFirstRecurring        bool       `json:"is_first_recurring,omitempty"`           // 如果是订阅的首次支付则为 true
	ShippingOptionID        string     `json:"shipping_option_id,omitempty"`           // 用户选择的配送选项标识
	OrderInfo               *OrderInfo `json:"order_info,omitempty"`                   // 用户提供的订单信息
	TelegramPaymentChargeID string     `json:"telegram_payment_charge_id"`             // Telegram 支付标识符
	ProviderPaymentChargeID string     `json:"provider_payment_charge_id"`             // 支付提供商的支付标识符
}
