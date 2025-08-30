package structure

// WebAppInfo 描述一个 Web 应用
type WebAppInfo struct {
	URL string `json:"url"` // 一个 HTTPS 的 Web 应用地址，将使用指定的初始化数据打开
}
