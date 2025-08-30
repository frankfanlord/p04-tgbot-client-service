package basic

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/bytedance/sonic"
)

// BaseResponse 基础响应
type BaseResponse struct {
	OK          bool                `json:"ok"`                    // OK 表示請求是否成功
	Result      any                 `json:"result,omitempty"`      // Result 包含請求成功時的結果數據
	ErrorCode   int                 `json:"error_code,omitempty"`  // ErrorCode 錯誤代碼，只有當 OK 為 false 時才會有值
	Description string              `json:"description,omitempty"` // Description 錯誤描述，只有當 OK 為 false 時才會有值
	Parameters  *ResponseParameters `json:"parameters,omitempty"`  // Parameters 可選的參數，提供關於錯誤的額外信息
}

// ResponseParameters 包含响应的附加参数信息
type ResponseParameters struct {
	MigrateToChatID int64 `json:"migrate_to_chat_id,omitempty"` // MigrateToChatID 當群組被升級為超級群組時，包含新的聊天 ID
	RetryAfter      int   `json:"retry_after,omitempty"`        // RetryAfter 當遇到速率限制時，建議的重試間隔（秒）
}

var (
	_client = &http.Client{
		Timeout: 10 * time.Second,
		Transport: &http.Transport{
			MaxIdleConns:        100,                              // 整个进程最多保留 100 个空闲（KeepAlive）TCP 连接
			IdleConnTimeout:     time.Second * time.Duration(120), // 一个空闲连接最长可以闲置 120 秒，超时就关闭
			DisableKeepAlives:   false,                            // 启用 KeepAlive
			ForceAttemptHTTP2:   true,                             // 强制尝试使用 HTTP/2
			MaxIdleConnsPerHost: 100,                              // 每个目标主机最多同时保持多少个空闲连接
		},
	}
	_domain = "https://api.telegram.org"
	_token  = ""
)

func Init(domain, token string) error {
	if domain == "" || token == "" {
		return errors.New("domain and token both can't be empty")
	}

	_domain = domain
	_token = token
	return nil
}

// GenerateURL 生成 URL
func GenerateURL(method Method) string {
	return fmt.Sprintf("%s/bot%s/%s", _domain, _token, method.String())
}

func exec(rawURL, ct string, raw []byte, receiver any) (int, string, error) {
	URL, err := url.Parse(rawURL)
	if err != nil {
		return 0, "", err
	}

	var reader io.ReadCloser
	if raw != nil && len(raw) > 0 {
		reader = io.NopCloser(bytes.NewReader(raw))
	}

	if ct == "" {
		ct = "application/json"
	}

	request := &http.Request{
		Method: http.MethodPost,
		URL:    URL,
		Header: map[string][]string{"Content-Type": {ct}},
		Body:   reader,
	}

	response, doErr := _client.Do(request)
	if doErr != nil {
		return 0, "", doErr
	}
	defer func() { _ = response.Body.Close() }()

	data, raErr := io.ReadAll(response.Body)
	if raErr != nil {
		return 0, "", raErr
	}

	return response.StatusCode, response.Status, sonic.Unmarshal(data, receiver)
}
