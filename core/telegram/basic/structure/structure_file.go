package structure

// File 表示一个准备好可以下载的文件。
// 可以通过链接 https://api.telegram.org/file/bot<token>/<file_path> 下载该文件。
// 链接至少在1小时内有效，若失效可通过调用 getFile 获取新链接。
// 最大可下载文件大小为 20 MB。
type File struct {
	FileID       string `json:"file_id"`        // 文件的标识符，可用于下载或复用该文件
	FileUniqueID string `json:"file_unique_id"` // 文件的唯一标识符，在不同机器人和时间中应保持不变，但不可用于下载或复用
	FileSize     int    `json:"file_size"`      // （可选）文件大小（字节），最大支持 52 位有效数字
	FilePath     string `json:"file_path"`      // （可选）文件路径，可用链接 https://api.telegram.org/file/bot<token>/<file_path> 获取
}
