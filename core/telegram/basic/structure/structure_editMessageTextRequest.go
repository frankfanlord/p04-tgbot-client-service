package structure

// EditMessageTextRequest 用于编辑文本消息或游戏消息
type EditMessageTextRequest[T NS] struct {
	BusinessConnectionID string                                              `json:"business_connection_id"` // 商业连接唯一标识符（可选）
	ChatID               T                                                   `json:"chat_id"`                // 目标聊天唯一标识符或频道用户名（可选，若未指定 inline_message_id 则必需）
	MessageID            int                                                 `json:"message_id"`             // 要编辑的消息标识符（可选，若未指定 inline_message_id 则必需）
	InlineMessageID      string                                              `json:"inline_message_id"`      // 内联消息的标识符（可选，若未指定 chat_id 和 message_id 则必需）
	Text                 string                                              `json:"text"`                   // 新的消息文本，解析后字符数为 1-4096（必填）
	ParseMode            string                                              `json:"parse_mode"`             // 文本解析模式，如 Markdown 或 HTML（可选）
	Entities             []*MessageEntity                                    `json:"entities"`               // 特殊实体的 JSON 序列化列表，可替代 parse_mode（可选）
	LinkPreviewOptions   *LinkPreviewOptions                                 `json:"link_preview_options"`   // 链接预览设置（可选）
	ReplyMarkup          any/**InlineKeyboardMarkup */ `json:"reply_markup"` // 内联键盘的 JSON 序列化对象（可选）
}
