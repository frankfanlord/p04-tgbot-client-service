package structure

// ChatLocation 表示一个聊天关联的地理位置。
type ChatLocation struct {
	Location *Location `json:"location"` // 超级群组关联的位置，不能是实时位置
	Address  string    `json:"address"`  // 位置地址；由聊天所有者定义，长度为1-64字符
}
