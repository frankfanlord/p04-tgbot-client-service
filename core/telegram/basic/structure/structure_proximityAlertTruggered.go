package structure

// ProximityAlertTriggered 此对象表示服务消息的内容，每当聊天中的用户触发另一个用户设置的接近警报时发送。
type ProximityAlertTriggered struct {
	Traveler *User `json:"traveler"` // 触发警报的用户
	Watcher  *User `json:"watcher"`  // 设置警报的用户
	Distance int   `json:"distance"` // 用户之间的距离
}
