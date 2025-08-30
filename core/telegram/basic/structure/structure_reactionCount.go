package structure

// ReactionCount 表示消息中的某个表情反应及其被添加的次数。
type ReactionCount struct {
	Type       *ReactionType `json:"type"`        // 表情反应的类型
	TotalCount int           `json:"total_count"` // 此表情被添加的次数
}
