package structure

// InlineQuery 表示一个传入的内联查询对象。当用户发送空查询时，Bot 可以返回默认或热门结果。
type InlineQuery struct {
	ID       string    `json:"id"`        // 查询的唯一标识符
	From     *User     `json:"from"`      // 发起查询的用户
	Query    string    `json:"query"`     // 查询文本（最多256个字符）
	Offset   string    `json:"offset"`    // 要返回的结果偏移量，Bot可控制
	ChatType string    `json:"chat_type"` // 可选。查询发起的聊天类型：可能是 "sender"（与发送者的私聊）、"private"、"group"、"supergroup" 或 "channel"
	Location *Location `json:"location"`  // 可选。发送者的位置，仅当Bot请求了用户位置时存在
}
