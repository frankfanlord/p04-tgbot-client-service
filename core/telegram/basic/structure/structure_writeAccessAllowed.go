package structure

// WriteAccessAllowed 表示一条服务消息，说明用户允许机器人发送消息
// 这种情况通常发生在：
// - 用户将机器人添加到附件菜单后
// - 通过链接启动 Web App
// - 或者通过 requestWriteAccess 方法接受了 Web App 的显式请求
type WriteAccessAllowed struct {
	FromRequest        bool   `json:"from_request"`         // 如果用户通过 Web App 的显式请求授予访问权限，则为 true
	WebAppName         string `json:"web_app_name"`         // 如果通过链接启动 Web App 时授予访问权限，表示 Web App 的名称
	FromAttachmentMenu bool   `json:"from_attachment_menu"` // 如果用户将机器人添加到附件菜单或侧边菜单时授予访问权限，则为 true
}
