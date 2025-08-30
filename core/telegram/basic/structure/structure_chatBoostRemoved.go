package structure

// ChatBoostRemoved 表示从聊天中移除的加成信息
type ChatBoostRemoved struct {
	Chat       *Chat            `json:"chat"`        // 被加成的聊天
	BoostID    string           `json:"boost_id"`    // 加成的唯一标识符
	RemoveDate int              `json:"remove_date"` // 移除加成的时间（Unix 时间戳）
	Source     *ChatBoostSource `json:"source"`      // 被移除加成的来源
}
