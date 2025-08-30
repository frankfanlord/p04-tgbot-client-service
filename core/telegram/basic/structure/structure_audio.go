package structure

// Audio 表示一个音频文件，Telegram 客户端将其视为音乐
type Audio struct {
	FileID       string     `json:"file_id"`        // 文件的唯一标识，可用于下载或重用该文件
	FileUniqueID string     `json:"file_unique_id"` // 文件的唯一标识，跨时间和不同机器人保持不变，不能用于下载或重用
	Duration     int        `json:"duration"`       // 音频的持续时间（单位：秒），由发送者定义
	Performer    string     `json:"performer"`      // （可选）演奏者，由发送者或音频标签定义
	Title        string     `json:"title"`          // （可选）标题，由发送者或音频标签定义
	FileName     string     `json:"file_name"`      // （可选）原始文件名，由发送者定义
	MimeType     string     `json:"mime_type"`      // （可选）MIME 类型，由发送者定义
	FileSize     int64      `json:"file_size"`      // （可选）文件大小（单位：字节）
	Thumbnail    *PhotoSize `json:"thumbnail"`      // （可选）音乐文件所属专辑封面的缩略图
}
