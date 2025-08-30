package structure

// Location 表示地图上的一个点
type Location struct {
	Latitude             float64 `json:"latitude"`               // 纬度，由发送者定义
	Longitude            float64 `json:"longitude"`              // 经度，由发送者定义
	HorizontalAccuracy   float64 `json:"horizontal_accuracy"`    // （可选）位置的不确定性半径，单位为米，范围 0-1500
	LivePeriod           int     `json:"live_period"`            // （可选）从消息发送时间起，位置可被更新的时间（秒）；仅用于实时位置
	Heading              int     `json:"heading"`                // （可选）用户移动的方向，单位为度，范围 1-360；仅用于实时位置
	ProximityAlertRadius int     `json:"proximity_alert_radius"` // （可选）用于靠近其他聊天成员的距离提醒的最大距离，仅用于发送的实时位置
}
