package basic

import (
	"client-service/core/telegram/basic/structure"
	"errors"
	"fmt"
	"jarvis/logger"
	"net/http"

	"github.com/bytedance/sonic"
)

func PinChatMessage[T structure.NS](messageID int, chatID T) error {
	rawURL := GenerateURL(MethodPinChatMessage)

	request := &structure.PinChatMessage[T]{
		BusinessConnectionID: "",        // 所代表业务连接的唯一标识符（可选）
		ChatID:               chatID,    // 目标聊天的唯一标识符或频道用户名（必须）
		MessageID:            messageID, // 要回复的消息ID（当前聊天中，或指定的 chat_id 中）
		DisableNotification:  false,     // 静默发送消息，用户收到无声通知（可选）
	}

	data, err := sonic.Marshal(request)
	if err != nil {
		return err
	}

	baseResponse := new(BaseResponse)
	if code, status, err := exec(rawURL, "", data, baseResponse); err != nil {
		logger.App().Errorf("exec [%s] error : %s", MethodSendMessage.String(), err.Error())
		return err
	} else {
		if code != http.StatusOK {
			return errors.New(fmt.Sprintf("request %s failed : %d - %s", rawURL, code, status))
		}
	}

	logger.App().Infof("base response success : %+v", *baseResponse)

	return nil
}
