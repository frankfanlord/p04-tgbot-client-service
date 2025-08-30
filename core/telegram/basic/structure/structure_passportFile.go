package structure

// PassportFile This object represents a file uploaded to Telegram Passport. Currently all Telegram Passport files are in JPEG format when decrypted and don't exceed 10MB.
type PassportFile struct {
	FileID       string `json:"file_id"`        // 此文件的标识符，可用于下载或重用此文件
	FileUniqueID string `json:"file_unique_id"` // 此文件的唯一标识符，应在不同时间和不同机器人之间保持相同。不能用于下载或重用此文件。
	FileSize     int    `json:"file_size"`      // 文件大小，单位为字节
	FileDate     int    `json:"file_date"`      // 文件上传的Unix时间
}
