package structure

// EncryptedPassportElement 描述了用户与机器人共享的文档或其他 Telegram Passport 元素。
type EncryptedPassportElement struct {
	Type        string          `json:"type"`                   // 元素类型。可以是 “personal_details”、“passport”、“driver_license”、“identity_card”、“internal_passport”、“address”、“utility_bill”、“bank_statement”、“rental_agreement”、“passport_registration”、“temporary_registration”、“phone_number”、“email” 中的一个。
	Data        string          `json:"data,omitempty"`         // 可选。用户提供的 Base64 编码的加密 Telegram Passport 元素数据；仅适用于 “personal_details”、“passport”、“driver_license”、“identity_card”、“internal_passport” 和 “address” 类型。可以使用附带的 EncryptedCredentials 进行解密和验证。
	PhoneNumber string          `json:"phone_number,omitempty"` // 可选。用户验证过的电话号码；仅适用于 “phone_number” 类型。
	Email       string          `json:"email,omitempty"`        // 可选。用户验证过的电子邮件地址；仅适用于 “email” 类型。
	Files       []*PassportFile `json:"files,omitempty"`        // 可选。用户提供的加密文档文件数组；仅适用于 “utility_bill”、“bank_statement”、“rental_agreement”、“passport_registration” 和 “temporary_registration” 类型。文件可以使用附带的 EncryptedCredentials 进行解密和验证。
	FrontSide   *PassportFile   `json:"front_side,omitempty"`   // 可选。用户提供的文档正面加密文件；仅适用于 “passport”、“driver_license”、“identity_card” 和 “internal_passport”。文件可以使用附带的 EncryptedCredentials 进行解密和验证。
	ReverseSide *PassportFile   `json:"reverse_side,omitempty"` // 可选。用户提供的文档反面加密文件；仅适用于 “driver_license” 和 “identity_card”。文件可以使用附带的 EncryptedCredentials 进行解密和验证。
	Selfie      *PassportFile   `json:"selfie,omitempty"`       // 可选。用户手持文档的自拍加密文件；如果请求 “passport”、“driver_license”、“identity_card” 和 “internal_passport”，则可用。文件可以使用附带的 EncryptedCredentials 进行解密和验证。
	Translation []*PassportFile `json:"translation,omitempty"`  // 可选。用户提供的文档翻译版本加密文件数组；如果请求 “passport”、“driver_license”、“identity_card”、“internal_passport”、“utility_bill”、“bank_statement”、“rental_agreement”、“passport_registration” 和 “temporary_registration” 类型，则可用。文件可以使用附带的 EncryptedCredentials 进行解密和验证。
	Hash        string          `json:"hash"`                   // 用于 PassportElementErrorUnspecified 的 Base64 编码的元素哈希。
}
