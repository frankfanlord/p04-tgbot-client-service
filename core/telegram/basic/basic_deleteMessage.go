package basic

import (
	"client-service/core/telegram/basic/structure"
	"errors"
	"fmt"
	"jarvis/logger"
	"net/http"

	"github.com/bytedance/sonic"
)

func DeleteMessage[T structure.NS](chatID T, messageID int) error {
	rawURL := GenerateURL(MethodDeleteMessage)

	request := &structure.DeleteMessage[T]{
		ChatID:    chatID,    // 整数或字符串 是 目标聊天或目标频道用户名的唯一标识符（格式为@channelusername）
		MessageID: messageID, // 待删除消息的标识符
	}

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
