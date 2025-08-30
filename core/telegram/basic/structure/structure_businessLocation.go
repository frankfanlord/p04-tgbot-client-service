package structure

// BusinessLocation 包含 Telegram 商业账户的位置信息
type BusinessLocation struct {
	Address  string    `json:"address"`  // 商业地址
	Location *Location `json:"location"` // 可选，商业位置
}
