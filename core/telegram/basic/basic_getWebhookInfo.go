package basic

import (
	"client-service/core/telegram/basic/structure"
	"errors"
	"fmt"
	"jarvis/logger"
	"net/http"

	"github.com/bytedance/sonic"
)

func GetWebhookInfo() (*structure.WebHookInfo, error) {
	rawURL := GenerateURL(MethodGetWebhookInfo)

	info := new(structure.WebHookInfo)

	baseResponse := new(BaseResponse)
	if code, status, err := exec(rawURL, "", nil, baseResponse); err != nil {
		logger.App().Errorf("exec [%s] error : %s", MethodGetWebhookInfo.String(), err.Error())
		return nil, err
	} else {
		if code != http.StatusOK {
			return nil, errors.New(fmt.Sprintf("request %s failed : %d - %s", rawURL, code, status))
		}
	}

	if !baseResponse.OK {
		logger.App().Warnf("base response failed : %+v", *baseResponse)
	} else {
		if data, mErr := sonic.Marshal(baseResponse.Result); mErr != nil {
			return nil, mErr
		} else {
			if err := sonic.Unmarshal(data, info); err != nil {
				return nil, err
			}
		}
	}

	logger.App().Infof("[%s] : %+v", MethodGetWebhookInfo.String(), *info)

	return info, nil
}
