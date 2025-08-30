package structure

// UniqueGift 表示从普通礼物升级而来的独特礼物
type UniqueGift struct {
	BaseName string              `json:"base_name"` // 原始普通礼物的可读名称
	Name     string              `json:"name"`      // 独特礼物的名称，可用于链接和展示
	Number   int                 `json:"number"`    // 相同类型普通礼物中升级得到的唯一编号
	Model    *UniqueGiftModel    `json:"model"`     // 礼物的模型
	Symbol   *UniqueGiftSymbol   `json:"symbol"`    // 礼物的符号
	Backdrop *UniqueGiftBackdrop `json:"backdrop"`  // 礼物的背景
}
