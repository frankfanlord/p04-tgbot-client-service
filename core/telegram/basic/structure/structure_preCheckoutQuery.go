package structure

// PreCheckoutQuery 包含有关传入的预结账查询的信息
type PreCheckoutQuery struct {
	ID               string     `json:"id"`                 // 查询唯一标识符
	From             *User      `json:"from"`               // 发送此查询的用户
	Currency         string     `json:"currency"`           // 三位 ISO 4217 货币代码，或“XTR”代表 Telegram Stars 支付
	TotalAmount      int        `json:"total_amount"`       // 金额（以最小单位表示的总价，非浮点数）
	InvoicePayload   string     `json:"invoice_payload"`    // 由 Bot 指定的发票载荷
	ShippingOptionID string     `json:"shipping_option_id"` // （可选）用户选择的配送选项 ID
	OrderInfo        *OrderInfo `json:"order_info"`         // （可选）用户提供的订单信息
}
