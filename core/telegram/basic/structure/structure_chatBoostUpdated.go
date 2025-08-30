package structure

// ChatBoostUpdated 表示对某个聊天新增或变更的加速信息
type ChatBoostUpdated struct {
	Chat  *Chat      `json:"chat"`  // 被加速的聊天
	Boost *ChatBoost `json:"boost"` // 聊天加速的相关信息
}
