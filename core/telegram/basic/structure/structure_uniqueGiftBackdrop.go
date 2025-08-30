package structure

// UniqueGiftBackdrop 描述独特礼物的背景
type UniqueGiftBackdrop struct {
	Name           string                    `json:"name"`             // 背景名称
	Colors         *UniqueGiftBackdropColors `json:"colors"`           // 背景颜色
	RarityPerMille int                       `json:"rarity_per_mille"` // 每升级1000个礼物中出现该背景的数量
}
