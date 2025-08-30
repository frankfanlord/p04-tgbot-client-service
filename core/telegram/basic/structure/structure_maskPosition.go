package structure

// MaskPosition 描述了默认情况下口罩应放置在脸部的哪个位置
type MaskPosition struct {
	Point  string  `json:"point"`   // 面部位置，可选值为 "forehead"（额头）、"eyes"（眼睛）、"mouth"（嘴巴）、"chin"（下巴）
	XShift float64 `json:"x_shift"` // X 轴偏移，单位是口罩宽度，负数表示向左偏移
	YShift float64 `json:"y_shift"` // Y 轴偏移，单位是口罩高度，正数表示向下偏移
	Scale  float64 `json:"scale"`   // 缩放系数，例如 2.0 表示放大两倍
}
