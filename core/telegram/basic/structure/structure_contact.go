package structure

// Contact 表示一个电话联系人对象
type Contact struct {
	PhoneNumber string `json:"phone_number"` // 联系人的电话号码
	FirstName   string `json:"first_name"`   // 联系人的名字
	LastName    string `json:"last_name"`    // 联系人的姓氏（可选）
	UserID      int    `json:"user_id"`      // Telegram 中联系人的用户 ID（可选，最大 52 位有效位）
	VCard       string `json:"vcard"`        // 以 vCard 格式提供的联系人附加信息（可选）
}
