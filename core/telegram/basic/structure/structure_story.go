package structure

// Story 表示一个故事对象
type Story struct {
	Chat *Chat `json:"chat"` // 发布该故事的聊天对象
	ID   int   `json:"id"`   // 在该聊天中的故事唯一标识符
}
