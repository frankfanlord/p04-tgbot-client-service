package basic

import (
	"client-service/core/telegram/basic/structure"
	"errors"
	"fmt"
	"jarvis/logger"
	"net/http"

	"github.com/bytedance/sonic"
)

func SetWebhook(webhook, secretToken string) error {
	rawURL := GenerateURL(MethodSetWebhook)

	request := &structure.SetWebhookRequest{
		URL:            webhook,
		MaxConnections: 100,
		SecretToken:    secretToken,
	}

	data, err := sonic.Marshal(request)
	if err != nil {
		return err
	}

	baseResponse := new(BaseResponse)
	if code, status, err := exec(rawURL, "", data, baseResponse); err != nil {
		logger.App().Errorf("exec [%s] error : %s", MethodSetWebhook.String(), err.Error())
		return err
	} else {
		if code != http.StatusOK {
			return errors.New(fmt.Sprintf("request %s failed : %d - %s", rawURL, code, status))
		}
	}

	logger.App().Infof("base response success : %+v", *baseResponse)

	return nil
}
