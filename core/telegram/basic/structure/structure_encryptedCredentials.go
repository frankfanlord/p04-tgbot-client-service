package structure

// EncryptedCredentials 描述了解密和认证 EncryptedPassportElement 所需的数据。
// 有关数据解密和认证过程的完整描述，请参阅 Telegram Passport 文档。
type EncryptedCredentials struct {
	Data   string `json:"data"`   // Base64 编码的加密 JSON 序列化数据，包含唯一用户有效负载、数据哈希和 EncryptedPassportElement 解密和认证所需的密钥
	Hash   string `json:"hash"`   // 用于数据认证的 Base64 编码数据哈希
	Secret string `json:"secret"` // 使用机器人公共 RSA 密钥加密的 Base64 编码密钥，数据解密所需
}
