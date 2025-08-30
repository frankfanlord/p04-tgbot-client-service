package structure

// ChatPhoto 此对象表示一张聊天照片。
type ChatPhoto struct {
	SmallFileID       string `json:"small_file_id"`        // 小型 (160x160) 聊天照片的文件标识符。此 file_id 仅可用于下载照片，且仅在照片未更改时有效。
	SmallFileUniqueID string `json:"small_file_unique_id"` // 小型 (160x160) 聊天照片的唯一文件标识符，该标识符应随时间和不同机器人而变化。不可用于下载或重复使用该文件。
	BigFileID         string `json:"big_file_id"`          // 大型 (640x640) 聊天照片的文件标识符。此 file_id 仅可用于下载照片，且仅在照片未更改时有效。
	BigFileUniqueID   string `json:"big_file_unique_id"`   // 大型 (640x640) 聊天照片的唯一文件标识符，该标识符应随时间和不同机器人而变化。不可用于下载或重复使用该文件。
}
