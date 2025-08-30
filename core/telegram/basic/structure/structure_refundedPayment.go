package structure

// RefundedPayment 包含有关已退款支付的基本信息。
type RefundedPayment struct {
	Currency                string `json:"currency"`                   // 三位 ISO 4217 货币代码，或“XTR”代表 Telegram Stars 支付，目前恒为 "XTR"
	TotalAmount             int    `json:"total_amount"`               // 按最小货币单位表示的退款总额（整数）。例如 1.45 美元应为 145
	InvoicePayload          string `json:"invoice_payload"`            // 由机器人指定的账单载荷
	TelegramPaymentChargeID string `json:"telegram_payment_charge_id"` // Telegram 支付标识符
	ProviderPaymentChargeID string `json:"provider_payment_charge_id"` // 可选。支付提供商的支付标识符
}
