package structure

// Game 表示一个游戏对象。使用 BotFather 创建和编辑游戏，其短名称将作为唯一标识符。
type Game struct {
	Title        string           `json:"title"`         // 游戏标题
	Description  string           `json:"description"`   // 游戏描述
	Photo        []*PhotoSize     `json:"photo"`         // 在聊天中游戏消息中显示的照片
	Text         string           `json:"text"`          // （可选）游戏的简要描述或包含的高分信息，可自动或手动编辑
	TextEntities []*MessageEntity `json:"text_entities"` // （可选）文本中出现的特殊实体，如用户名、URL、机器人命令等
	Animation    *Animation       `json:"animation"`     // （可选）在聊天中游戏消息中显示的动画
}
