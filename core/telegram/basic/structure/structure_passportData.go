package structure

// PassportData Describes Telegram Passport data shared with the bot by the user.
type PassportData struct {
	Data        []*EncryptedPassportElement `json:"data"`        // 共享给机器人的文档和其他Telegram护照元素的信息数组
	Credentials *EncryptedCredentials       `json:"credentials"` // 解密数据所需的加密凭据
}
