package structure

// UniqueGiftSymbol 描述出现在独特礼物图案上的符号
type UniqueGiftSymbol struct {
	Name           string   `json:"name"`             // 符号名称
	Sticker        *Sticker `json:"sticker"`          // 代表该礼物的贴纸
	RarityPerMille int      `json:"rarity_per_mille"` // 每升级1000个礼物中出现该模型的数量
}
