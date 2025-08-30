package structure

// PollOption 表示投票中的一个选项信息
type PollOption struct {
	Text         string           `json:"text"`          // 选项文本，长度为1-100个字符
	TextEntities []*MessageEntity `json:"text_entities"` // 可选。出现在选项文本中的特殊实体，目前仅支持自定义表情符号实体
	VoterCount   int              `json:"voter_count"`   // 投票该选项的用户数量
}
