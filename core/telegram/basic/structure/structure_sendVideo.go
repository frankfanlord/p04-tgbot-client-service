package structure

// SendVideoRequest 代表发送视频的请求结构体
type SendVideoRequest[T NS] struct {
	BusinessConnectionID  string              `json:"business_connection_id,omitempty"`   // 代表业务连接的唯一标识
	ChatID                T                   `json:"chat_id"`                            // 聊天ID或用户名（@channelusername）
	MessageThreadID       int                 `json:"message_thread_id,omitempty"`        // 论坛话题ID，仅适用于超级群
	Video                 string              `json:"video"`                              // 视频，可以是 file_id、HTTP URL 或 InputFile
	Duration              int                 `json:"duration,omitempty"`                 // 视频时长（秒）
	Width                 int                 `json:"width,omitempty"`                    // 视频宽度
	Height                int                 `json:"height,omitempty"`                   // 视频高度
	Thumbnail             string              `json:"thumbnail,omitempty"`                // 缩略图（JPEG，小于200KB，可通过 attach:// 方式上传）
	Cover                 string              `json:"cover,omitempty"`                    // 视频封面，支持 attach:// 或 URL
	StartTimestamp        int                 `json:"start_timestamp,omitempty"`          // 视频起始时间戳
	Caption               string              `json:"caption,omitempty"`                  // 视频说明文字，支持格式化
	ParseMode             string              `json:"parse_mode,omitempty"`               // 解析模式：Markdown 或 HTML
	CaptionEntities       []*MessageEntity    `json:"caption_entities,omitempty"`         // caption 中使用的实体
	ShowCaptionAboveMedia bool                `json:"show_caption_above_media,omitempty"` // 是否将 caption 显示在媒体上方
	HasSpoiler            bool                `json:"has_spoiler,omitempty"`              // 是否启用剧透动画
	SupportsStreaming     bool                `json:"supports_streaming,omitempty"`       // 是否支持流式播放
	DisableNotification   bool                `json:"disable_notification,omitempty"`     // 是否静默发送
	ProtectContent        bool                `json:"protect_content,omitempty"`          // 是否保护内容，防止转发/保存
	AllowPaidBroadcast    bool                `json:"allow_paid_broadcast,omitempty"`     // 是否启用付费广播（最多每秒1000条）
	MessageEffectID       string              `json:"message_effect_id,omitempty"`        // 消息特效ID，仅用于私聊
	ReplyParameters       *ReplyParameters[T] `json:"reply_parameters,omitempty"`         // 回复的消息信息
	ReplyMarkup           any                 `json:"reply_markup,omitempty"`             // 自定义回复键盘（InlineKeyboardMarkup、ReplyKeyboardMarkup 等）
}
