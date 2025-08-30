package basic

import (
	"client-service/core/telegram/basic/structure"
	"errors"
	"fmt"
	"jarvis/logger"
	"net/http"

	"github.com/bytedance/sonic"
	"github.com/spf13/cast"
)

func SendMessage[T structure.NS](messageID int, chatID T, text, parseMode string, replyMarkup any) (int, error) {
	rawURL := GenerateURL(MethodSendMessage)

	request := &structure.SendMessageRequest[T]{
		BusinessConnectionID: "",                                              // 所代表业务连接的唯一标识符（可选）
		ChatID:               chatID,                                          // 目标聊天的唯一标识符或频道用户名（必须）
		MessageThreadID:      0,                                               // 论坛中目标消息线程（话题）的唯一标识符，仅用于超级群组论坛（可选）
		Text:                 text,                                            // 要发送的消息文本，实体解析后长度为 1-4096 字符（必须）
		ParseMode:            parseMode,                                       // 消息文本实体解析模式（可选）
		Entities:             []*structure.MessageEntity{},                    // 消息中出现的特殊实体列表（可选）
		LinkPreviewOptions:   &structure.LinkPreviewOptions{IsDisabled: true}, // 链接预览的生成选项（可选）
		DisableNotification:  false,                                           // 静默发送消息，用户收到无声通知（可选）
		ProtectContent:       false,                                           // 保护已发送消息内容不被转发和保存（可选）
		AllowPaidBroadcast:   false,                                           // 开启付费广播模式（最多每秒1000条）（可选）
		MessageEffectID:      "",                                              // 要添加到消息的消息特效ID，仅限私聊使用（可选）
		ReplyParameters: &structure.ReplyParameters[T]{
			MessageID: messageID, //  要回复的消息ID（当前聊天中，或指定的 chat_id 中）
			ChatID:    chatID,    //（可选）如果要回复的消息来自其他聊天，填入唯一聊天ID或频道用户名（格式为 @channelusername）
		}, // 回复的消息相关描述（可选）
		ReplyMarkup: replyMarkup, // 附加交互界面选项，如键盘、强制回复等（可选）||| InlineKeyboardMarkup or ReplyKeyboardMarkup or ReplyKeyboardRemove or ForceReply
	}

	data, err := sonic.Marshal(request)
	if err != nil {
		return 0, err
	}

	mid := 0

	baseResponse := new(BaseResponse)
	if code, status, err := exec(rawURL, "", data, baseResponse); err != nil {
		logger.App().Errorf("exec [%s] error : %s", MethodSendMessage.String(), err.Error())
		return 0, err
	} else {
		if code != http.StatusOK {
			return 0, errors.New(fmt.Sprintf("request %s failed : %d - %s", rawURL, code, status))
		} else {
			if baseResponse.Result != nil {
				if m, tErr := cast.ToStringMapE(baseResponse.Result); tErr != nil {
					logger.App().Errorf("baseResponse.Result to string map error : %s", err.Error())
				} else {
					if v, e := m["message_id"]; e {
						mid = cast.ToInt(v)
					}
				}
			}
		}
	}

	logger.App().Infof("base response success : %+v", *baseResponse)

	return mid, nil
}
