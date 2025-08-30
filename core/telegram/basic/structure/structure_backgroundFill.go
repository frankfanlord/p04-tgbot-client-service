package structure

// BackgroundFill 背景填充统一结构体，支持 solid、gradient、freeform_gradient 三种类型
type BackgroundFill struct {
	Type string `json:"type"` // 背景填充类型，可为 "solid"、"gradient"、"freeform_gradient"

	// ===== solid 类型专用 =====
	Color int `json:"color,omitempty"` // 实色填充颜色，RGB24 格式，仅当 type 为 "solid" 时有效

	// ===== gradient 类型专用 =====
	TopColor      int `json:"top_color,omitempty"`      // 渐变顶部颜色，RGB24 格式，仅当 type 为 "gradient" 时有效
	BottomColor   int `json:"bottom_color,omitempty"`   // 渐变底部颜色，RGB24 格式，仅当 type 为 "gradient" 时有效
	RotationAngle int `json:"rotation_angle,omitempty"` // 渐变顺时针旋转角度（0-359），仅当 type 为 "gradient" 时有效

	// ===== freeform_gradient 类型专用 =====
	Colors []int `json:"colors,omitempty"` // 自由形式渐变的基础颜色（3-4个），RGB24 格式，仅当 type 为 "freeform_gradient" 时有效
}
