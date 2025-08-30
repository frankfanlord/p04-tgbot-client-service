package structure

type NS interface {
	~int | ~string
}

// SendMessageRequest 表示发送消息的请求结构体
type SendMessageRequest[T NS] struct {
	BusinessConnectionID string              `json:"business_connection_id"` // 所代表业务连接的唯一标识符（可选）
	ChatID               T                   `json:"chat_id"`                // 目标聊天的唯一标识符或频道用户名（必须）
	MessageThreadID      int                 `json:"message_thread_id"`      // 论坛中目标消息线程（话题）的唯一标识符，仅用于超级群组论坛（可选）
	Text                 string              `json:"text"`                   // 要发送的消息文本，实体解析后长度为 1-4096 字符（必须）
	ParseMode            string              `json:"parse_mode"`             // 消息文本实体解析模式（可选）
	Entities             []*MessageEntity    `json:"entities"`               // 消息中出现的特殊实体列表（可选）
	LinkPreviewOptions   *LinkPreviewOptions `json:"link_preview_options"`   // 链接预览的生成选项（可选）
	DisableNotification  bool                `json:"disable_notification"`   // 静默发送消息，用户收到无声通知（可选）
	ProtectContent       bool                `json:"protect_content"`        // 保护已发送消息内容不被转发和保存（可选）
	AllowPaidBroadcast   bool                `json:"allow_paid_broadcast"`   // 开启付费广播模式（最多每秒1000条）（可选）
	MessageEffectID      string              `json:"message_effect_id"`      // 要添加到消息的消息特效ID，仅限私聊使用（可选）
	ReplyParameters      *ReplyParameters[T] `json:"reply_parameters"`       // 回复的消息相关描述（可选）
	ReplyMarkup          any                 `json:"reply_markup"`           // 附加交互界面选项，如键盘、强制回复等（可选）||| InlineKeyboardMarkup or ReplyKeyboardMarkup or ReplyKeyboardRemove or ForceReply
}
