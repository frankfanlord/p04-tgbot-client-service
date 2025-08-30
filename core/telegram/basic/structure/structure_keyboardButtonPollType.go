package structure

// KeyboardButtonPollType 表示在按下对应按钮时允许创建和发送的投票类型。
type KeyboardButtonPollType struct {
	Type string `json:"type"` // 可选。如果为 "quiz"，用户只能创建测验模式的投票；如果为 "regular"，则只能创建常规投票；否则，允许创建任意类型的投票。
}
