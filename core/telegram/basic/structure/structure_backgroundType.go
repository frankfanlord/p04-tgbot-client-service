package structure

// BackgroundType 描述聊天背景的类型（填充、壁纸、图案或聊天主题）。
type BackgroundType struct {
	Type      string `json:"type"`       // 背景类型: "fill"、"wallpaper"、"pattern"、"chat_theme"
	ThemeName string `json:"theme_name"` // 聊天主题的名称（仅当 type 为 "chat_theme" 时有效）

	// 通用字段（根据 type 判断是否使用）
	Fill     *BackgroundFill `json:"fill"`     // 背景填充（type 为 fill 或 pattern 时使用）
	Document *Document       `json:"document"` // 壁纸或图案文档（type 为 wallpaper 或 pattern 时使用）

	// 可选字段（取决于类型）
	DarkThemeDimming int  `json:"dark_theme_dimming"` // 暗色主题下的背景变暗程度（0-100）（fill, wallpaper类型可用）
	Intensity        int  `json:"intensity"`          // 图案强度（type 为 pattern 时使用）
	IsInverted       bool `json:"is_inverted"`        // 是否反转（type 为 pattern 且暗色主题时使用）
	IsBlurred        bool `json:"is_blurred"`         // 是否模糊处理（type 为 wallpaper 时使用）
	IsMoving         bool `json:"is_moving"`          // 是否随设备倾斜而移动（wallpaper, pattern 类型可用）
}
