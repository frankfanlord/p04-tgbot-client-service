package structure

// UniqueGiftModel 描述一个独特礼物的模型信息
type UniqueGiftModel struct {
	Name           string   `json:"name"`             // 模型名称
	Sticker        *Sticker `json:"sticker"`          // 代表该礼物的贴纸
	RarityPerMille int      `json:"rarity_per_mille"` // 每升级1000个礼物中出现该模型的数量
}
