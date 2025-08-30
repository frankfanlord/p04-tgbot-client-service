package structure

// ChosenInlineResult 表示一个由用户选中的 inline 查询结果，并发送给聊天对象
type ChosenInlineResult struct {
	ResultID        string    `json:"result_id"`                   // 被选中的结果的唯一标识符
	From            *User     `json:"from"`                        // 选中该结果的用户
	Location        *Location `json:"location,omitempty"`          // 可选。发送者的位置，仅对要求用户位置的机器人可用
	InlineMessageID string    `json:"inline_message_id,omitempty"` // 可选。已发送的 inline 消息标识符，仅当消息附带 inline keyboard 时才可用
	Query           string    `json:"query"`                       // 用户用于获取该结果的查询字符串
}
