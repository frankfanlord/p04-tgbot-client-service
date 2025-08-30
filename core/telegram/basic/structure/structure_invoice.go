package structure

// Invoice 包含有关发票的基本信息
type Invoice struct {
	Title          string `json:"title"`           // 商品名称
	Description    string `json:"description"`     // 商品描述
	StartParameter string `json:"start_parameter"` // 用于生成此发票的唯一机器人深链接参数
	Currency       string `json:"currency"`        // 三位 ISO 4217 货币代码，或 “XTR” 表示 Telegram Stars 支付
	TotalAmount    int    `json:"total_amount"`    // 使用货币最小单位表示的总价（整数）
}
