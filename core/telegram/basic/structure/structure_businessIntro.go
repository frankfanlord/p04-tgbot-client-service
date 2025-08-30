package structure

// BusinessIntro 包含 Telegram Business 帐户首页设置信息。
type BusinessIntro struct {
	Title   string   `json:"title"`   // 可选。业务简介的标题文本
	Message string   `json:"message"` // 可选。业务简介的消息文本
	Sticker *Sticker `json:"sticker"` // 可选。业务简介贴纸
}
