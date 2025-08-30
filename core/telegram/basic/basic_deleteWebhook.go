package basic

import (
	"client-service/core/telegram/basic/structure"
	"errors"
	"fmt"
	"jarvis/logger"
	"net/http"

	"github.com/bytedance/sonic"
)

func DeleteWebhook() error {
	rawURL := GenerateURL(MethodDeleteWebhook)

	request := &structure.DeleteWebhookRequest{DropPendingUpdates: false}

	data, err := sonic.Marshal(request)
	if err != nil {
		return err
	}

	baseResponse := new(BaseResponse)
	if code, status, err := exec(rawURL, "", data, baseResponse); err != nil {
		logger.App().Errorf("exec [%s] error : %s", MethodGetWebhookInfo.String(), err.Error())
		return err
	} else {
		if code != http.StatusOK {
			return errors.New(fmt.Sprintf("request %s failed : %d - %s", rawURL, code, status))
		}
	}

	logger.App().Infof("base response success : %+v", *baseResponse)

	return nil
}
