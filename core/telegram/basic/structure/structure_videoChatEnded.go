package structure

// VideoChatEnded 表示群组中视频通话结束的服务消息
type VideoChatEnded struct {
	Duration int `json:"duration"` // 视频通话持续时间（单位：秒）
}
