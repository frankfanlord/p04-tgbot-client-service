package basic

import (
	"client-service/core/telegram/basic/structure"
	"errors"
	"fmt"
	"jarvis/logger"
	"net/http"

	"github.com/bytedance/sonic"
)

func SetMyShortDescription(desc string) error {
	rawURL := GenerateURL(MethodSetMyShortDescription)

	request := &structure.SetMyShortDescriptionRequest{
		ShortDescription: desc,
		LanguageCode:     "",
	}

	data, err := sonic.Marshal(request)
	if err != nil {
		return err
	}

	baseResponse := new(BaseResponse)
	if code, status, err := exec(rawURL, "", data, baseResponse); err != nil {
		logger.App().Errorf("exec [%s] error : %s", MethodSetMyShortDescription.String(), err.Error())
		return err
	} else {
		if code != http.StatusOK {
			return errors.New(fmt.Sprintf("request %s failed : %d - %s", rawURL, code, status))
		}
	}

	logger.App().Infof("base response success : %+v", *baseResponse)

	return nil
}

func GetMyShortDescription() (*structure.GetMyShortDescriptionResponse, error) {
	rawURL := GenerateURL(MethodGetMyShortDescription)

	request := &structure.GetMyDescriptionRequest{LanguageCode: ""}

	data, err := sonic.Marshal(request)
	if err != nil {
		return nil, err
	}

	baseResponse := new(BaseResponse)
	if code, status, err := exec(rawURL, "", data, baseResponse); err != nil {
		logger.App().Errorf("exec [%s] error : %s", MethodGetMyShortDescription.String(), err.Error())
		return nil, err
	} else {
		if code != http.StatusOK {
			return nil, errors.New(fmt.Sprintf("request %s failed : %d - %s", rawURL, code, status))
		}
	}

	info := new(structure.GetMyShortDescriptionResponse)

	if !baseResponse.OK {
		logger.App().Warnf("base response failed : %+v", *baseResponse)
		return nil, errors.New(baseResponse.Description)
	} else {
		if data, mErr := sonic.Marshal(baseResponse.Result); mErr != nil {
			return nil, mErr
		} else {
			if err := sonic.Unmarshal(data, info); err != nil {
				return nil, err
			}
		}
	}

	return info, nil
}
