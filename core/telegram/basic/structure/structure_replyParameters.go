package structure

// ReplyParameters 描述发送消息时的回复参数
type ReplyParameters[T NS] struct {
	MessageID                int              `json:"message_id"`                            // 要回复的消息ID（当前聊天中，或指定的 chat_id 中）
	ChatID                   T                `json:"chat_id,omitempty"`                     // （可选）如果要回复的消息来自其他聊天，填入唯一聊天ID或频道用户名（格式为 @channelusername）
	AllowSendingWithoutReply bool             `json:"allow_sending_without_reply,omitempty"` // （可选）即使指定的消息未找到也发送；业务账号发消息时总为 true，其他聊天中回复总为 false
	Quote                    string           `json:"quote,omitempty"`                       // （可选）消息中的引用部分，必须是原始消息中的精确子串（支持富文本格式）
	QuoteParseMode           string           `json:"quote_parse_mode,omitempty"`            // （可选）引用内容的解析模式，参考格式化选项
	QuoteEntities            []*MessageEntity `json:"quote_entities,omitempty"`              // （可选）引用部分中的特殊实体列表，替代 quote_parse_mode 使用
	QuotePosition            int              `json:"quote_position,omitempty"`              // （可选）引用内容在原始消息中的位置（UTF-16 编码单位）
}
