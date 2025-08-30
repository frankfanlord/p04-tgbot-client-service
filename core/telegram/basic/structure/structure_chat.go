package structure

// Chat 该对象代表聊天。
type Chat struct {
	ID        int    `json:"id"`         // 本次聊天的唯一标识符。此数字的有效位可能超过 32 位，某些编程语言在解释它时可能存在困难/隐含缺陷。但它最多只有 52 位有效位，因此使用有符号的 64 位整数或双精度浮点类型存储此标识符是安全的。
	Type      string `json:"type"`       // 聊天类型，可以是“私聊”、“群组”、“超级群组”或“频道”。 (“private”, “group”, “supergroup” or “channel”)
	Title     string `json:"title"`      // 可选。标题，用于超级群组、频道和群聊。
	Username  string `json:"username"`   // 可选。用户名，用于私聊、超级群组和频道（如有）。
	FirstName string `json:"first_name"` // 可选。私聊对方的名字。
	LastName  string `json:"last_name"`  // 可选。私聊对方的姓氏。
	IsForum   bool   `json:"is_forum"`   // 可选。如果超级群组聊天是论坛（已启用主题），则为 True。
}
