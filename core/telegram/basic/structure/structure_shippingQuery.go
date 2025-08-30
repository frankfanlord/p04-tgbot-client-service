package structure

// ShippingQuery 表示一个传入的运费查询信息。
type ShippingQuery struct {
	ID              string           `json:"id"`               // 唯一的查询标识符
	From            *User            `json:"from"`             // 发送查询的用户
	InvoicePayload  string           `json:"invoice_payload"`  // Bot 指定的发票负载
	ShippingAddress *ShippingAddress `json:"shipping_address"` // 用户填写的收货地址
}
