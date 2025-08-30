package structure

// OrderInfo 表示一个订单信息
type OrderInfo struct {
	Name            string           `json:"name"`             // （可选）用户姓名
	PhoneNumber     string           `json:"phone_number"`     // （可选）用户手机号
	Email           string           `json:"email"`            // （可选）用户邮箱
	ShippingAddress *ShippingAddress `json:"shipping_address"` // （可选）用户收货地址
}
