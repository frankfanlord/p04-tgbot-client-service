package structure

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
)

// WebHookInfo 描述 webhook 的当前状态
type WebHookInfo struct {
	URL                          string   `json:"url"`                             // Webhook URL，如果未设置 Webhook，则可能为空
	HasCustomCertificate         bool     `json:"has_custom_certificate"`          // 如果为 Webhook 证书检查提供了自定义证书，则为 True
	PendingUpdateCount           int      `json:"pending_update_count"`            // 等待交付的更新数量
	IPAddress                    string   `json:"ip_address"`                      // 可选。当前使用的 Webhook IP 地址
	LastErrorDate                int      `json:"last_error_date"`                 // 可选。尝试通过 Webhook 交付更新时发生的最新错误的 Unix 时间
	LastErrorMessage             string   `json:"last_error_message"`              // 可选。尝试通过 Webhook 交付更新时发生的最新错误的可读格式的错误消息
	LastSynchronizationErrorDate int      `json:"last_synchronization_error_date"` // 可选。尝试与 Telegram 数据中心同步可用更新时发生的最近一次错误的 Unix 时间
	MaxConnections               int      `json:"max_connections"`                 // 可选。允许同时与 webhook 建立 HTTPS 连接以进行更新传输的最大数量
	AllowedUpdates               []string `json:"allowed_updates"`                 // 可选。机器人订阅的更新类型列表。默认为除 chat_member 之外的所有更新类型
}

// SetWebhookRequest 使用此方法指定 URL，并通过传出 webhook 接收传入更新。每当机器人有更新时，我们将向指定的 URL 发送 HTTPS POST 请求，
// 其中包含 JSON 序列化的更新。如果请求失败（请求的响应 HTTP 状态码不是 2XY），我们将重复该请求，并在合理的尝试次数后放弃。成功时返回 True。
// 如果您想确保 webhook 是由您设置的，您可以在 secret_token 参数中指定机密数据。如果指定，请求将包含一个标头“X-Telegram-Bot-Api-Secret-Token”，其中包含机密令牌作为内容。
type SetWebhookRequest struct {
	URL                string   `json:"url"`                  // 发送更新的 HTTPS URL。使用空字符串可移除 webhook 集成
	Certificate        string   `json:"certificate"`          // 可选 上传您的公钥证书，以便检查正在使用的根证书。详情请参阅我们的自签名指南。
	IPAddress          string   `json:"ip_address"`           // 用于发送 webhook 请求的固定 IP 地址，而不是通过 DNS 解析的 IP 地址。
	MaxConnections     int      `json:"max_connections"`      // 可选 允许同时连接到 webhook 进行更新传递的最大 HTTPS 连接数，范围为 1-100。默认为 40。使用较低的值可限制机器人服务器的负载，使用较高的值可提高机器人的吞吐量。
	AllowedUpdates     []string `json:"allowed_updates"`      // 可选 您希望机器人接收的更新类型的 JSON 序列化列表。例如，指定 ["message", "edited_channel_post", "callback_query"] 表示仅接收这些类型的更新。有关可用更新类型的完整列表，请参阅更新。指定一个空列表，用于接收除 chat_member、message_reaction 和 message_reaction_count（默认）之外的所有更新类型。如果未指定，则使用之前的设置。请注意，此参数不会影响调用 setWebhook 之前创建的更新，因此可能会在短时间内收到不需要的更新。
	DropPendingUpdates bool     `json:"drop_pending_updates"` // 可选 传递 True 以丢弃所有待处理的更新
	SecretToken        string   `json:"secret_token"`         // 可选 在每个 webhook 请求中，通过标头“X-Telegram-Bot-Api-Secret-Token”发送的秘密令牌，长度为 1-256 个字符。仅允许使用 A-Z、a-z、0-9、_ 和 - 字符。此标头有助于确保请求来自您设置的 webhook。
}

func (swr *SetWebhookRequest) Marshal() ([]byte, string, error) {
	var buf bytes.Buffer
	writer := multipart.NewWriter(&buf)

	writer.WriteField("url", swr.URL)

	if swr.Certificate != "" {
		file, err := os.Open(swr.Certificate)
		if err != nil {
			return nil, "", err
		}
		defer file.Close()

		part, err := writer.CreateFormFile("certificate", filepath.Base(swr.Certificate))
		if err != nil {
			return nil, "", err
		}
		_, err = io.Copy(part, file)
		if err != nil {
			return nil, "", err
		}
	}

	if swr.IPAddress != "" {
		writer.WriteField("ip_address", swr.IPAddress)
	}

	if swr.MaxConnections != 0 {
		writer.WriteField("max_connections", fmt.Sprintf("%d", swr.MaxConnections))
	}

	if len(swr.AllowedUpdates) > 0 {
		updatesJSON, _ := json.Marshal(swr.AllowedUpdates)
		writer.WriteField("allowed_updates", string(updatesJSON))
	}

	if swr.DropPendingUpdates {
		writer.WriteField("drop_pending_updates", "true")
	}

	if swr.SecretToken != "" {
		writer.WriteField("secret_token", swr.SecretToken)
	}

	if err := writer.Close(); err != nil {
		return nil, "", err
	}

	return buf.Bytes(), writer.FormDataContentType(), nil
}

// DeleteWebhookRequest 用于取消 webhook 集成，如果你决定改为使用 getUpdates 获取更新。
type DeleteWebhookRequest struct {
	DropPendingUpdates bool `json:"drop_pending_updates"` // 是否丢弃所有待处理的更新
}
