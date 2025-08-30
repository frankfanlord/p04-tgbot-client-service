package structure

// User 此对象代表 Telegram 用户或机器人
type User struct {
	ID                      int    `json:"id"`                          // 此用户或机器人的唯一标识符。此数字的有效位可能超过 32 位，某些编程语言在解释它时可能存在困难/隐蔽缺陷。但它最多有 52 位有效位，因此使用 64 位整数或双精度浮点类型存储此标识符是安全的。
	IsBot                   bool   `json:"is_bot"`                      // 如果此用户是机器人，则为 True
	FirstName               string `json:"first_name"`                  // 用户或机器人的名字
	LastName                string `json:"last_name"`                   // 可选。用户或机器人的姓氏
	Username                string `json:"username"`                    // 可选。用户或机器人的用户名
	LanguageCode            string `json:"language_code"`               // 可选。用户语言的 IETF 语言标签
	IsPremium               bool   `json:"is_premium"`                  // 可选。如果此用户是 Telegram Premium 用户，则为 True
	AddedToAttachmentMenu   bool   `json:"added_to_attachment_menu"`    // 可选。如果此用户已将机器人添加到附件菜单，则返回 True
	CanJoinGroups           bool   `json:"can_join_groups"`             // 可选。如果机器人可以被邀请加入群组，则返回 True。仅在 getMe 中返回。
	CanReadAllGroupMessages bool   `json:"can_read_all_group_messages"` // 可选。如果机器人的隐私模式已禁用，则返回 True。仅在 getMe 中返回。
	SupportsInlineQueries   bool   `json:"supports_inline_queries"`     // 可选。如果机器人支持内联查询，则返回 True。仅在 getMe 中返回。
	CanConnectToBusiness    bool   `json:"can_connect_to_business"`     // 可选。如果机器人可以连接到 Telegram Business 帐户接收其消息，则返回 True。仅在 getMe 中返回。
	HasMainWebApp           bool   `json:"has_main_web_app"`            // 可选。如果机器人拥有主 Web 应用，则返回 True。仅在 getMe 中返回。
}
