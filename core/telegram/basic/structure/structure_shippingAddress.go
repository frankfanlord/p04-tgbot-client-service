package structure

// ShippingAddress 表示一个收货地址
type ShippingAddress struct {
	CountryCode string `json:"country_code"` // 国家代码（ISO 3166-1 alpha-2 的两位字母）
	State       string `json:"state"`        // 州、省（如适用）
	City        string `json:"city"`         // 城市
	StreetLine1 string `json:"street_line1"` // 地址第一行
	StreetLine2 string `json:"street_line2"` // 地址第二行
	PostCode    string `json:"post_code"`    // 邮政编码
}
