package structure

// Dice 表示一个带有随机值动画效果的 emoji 对象。
type Dice struct {
	Emoji string `json:"emoji"` // 骰子动画使用的 emoji 表情
	Value int    `json:"value"` // 骰子的点数，对于“🎲”、“🎯”、“🎳”是 1-6，对于“🏀”和“⚽”是 1-5，对于“🎰”是 1-64
}
