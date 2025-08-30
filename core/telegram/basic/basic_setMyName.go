package basic

import (
	"client-service/core/telegram/basic/structure"
	"errors"
	"fmt"
	"jarvis/logger"
	"net/http"

	"github.com/bytedance/sonic"
)

func SetMyName(name string) error {
	rawURL := GenerateURL(MethodSetMyName)

	request := &structure.SetMyNameRequest{
		Name:         name,
		LanguageCode: "",
	}

	data, err := sonic.Marshal(request)
	if err != nil {
		return err
	}

	baseResponse := new(BaseResponse)
	if code, status, err := exec(rawURL, "", data, baseResponse); err != nil {
		logger.App().Errorf("exec [%s] error : %s", MethodSetMyName.String(), err.Error())
		return err
	} else {
		if code != http.StatusOK {
			return errors.New(fmt.Sprintf("request %s failed : %d - %s", rawURL, code, status))
		}
	}

	logger.App().Infof("base response success : %+v", *baseResponse)

	return nil
}

func GetMyName() (*structure.GetMyNameResponse, error) {
	rawURL := GenerateURL(MethodGetMyName)

	request := &structure.GetMyNameRequest{LanguageCode: ""}

	data, err := sonic.Marshal(request)
	if err != nil {
		return nil, err
	}

	baseResponse := new(BaseResponse)
	if code, status, err := exec(rawURL, "", data, baseResponse); err != nil {
		logger.App().Errorf("exec [%s] error : %s", MethodGetMyName.String(), err.Error())
		return nil, err
	} else {
		if code != http.StatusOK {
			return nil, errors.New(fmt.Sprintf("request %s failed : %d - %s", rawURL, code, status))
		}
	}

	info := new(structure.GetMyNameResponse)

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
