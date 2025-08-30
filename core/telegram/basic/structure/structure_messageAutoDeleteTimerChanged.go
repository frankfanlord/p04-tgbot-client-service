package structure

// MessageAutoDeleteTimerChanged 此对象表示有关自动删除计时器设置更改的服务消息。
type MessageAutoDeleteTimerChanged struct {
	MessageAutoDeleteTime int `json:"message_auto_delete_time"` // 聊天消息自动删除时间新增：以秒为单位
}
