package structure

// VideoChatScheduled 表示群聊中计划的视频通话的服务消息对象
type VideoChatScheduled struct {
	// start_date 表示由聊天管理员计划开始视频通话的时间点（Unix 时间戳）
	StartDate int `json:"start_date"`
}
