package basic

import (
	"client-service/core/telegram/basic/structure"
	"errors"
	"fmt"
	"jarvis/logger"
	"net/http"

	"github.com/bytedance/sonic"
)

func EditMessageText[T structure.NS](messageID int, chatID T, text, parseMode string, replyMarkup any) error {
	rawURL := GenerateURL(MethodEditMessageText)

	request := &structure.EditMessageTextRequest[T]{
		BusinessConnectionID: "",                                              // 商业连接唯一标识符（可选）
		ChatID:               chatID,                                          // 目标聊天唯一标识符或频道用户名（可选，若未指定 inline_message_id 则必需）
		MessageID:            messageID,                                       // 要编辑的消息标识符（可选，若未指定 inline_message_id 则必需）
		InlineMessageID:      "",                                              // 内联消息的标识符（可选，若未指定 chat_id 和 message_id 则必需）
		Text:                 text,                                            // 新的消息文本，解析后字符数为 1-4096（必填）
		ParseMode:            parseMode,                                       // 文本解析模式，如 Markdown 或 HTML（可选）
		Entities:             []*structure.MessageEntity{},                    // 特殊实体的 JSON 序列化列表，可替代 parse_mode（可选）
		LinkPreviewOptions:   &structure.LinkPreviewOptions{IsDisabled: true}, // 链接预览设置（可选）
		ReplyMarkup:          replyMarkup,                                     // 内联键盘的 JSON 序列化对象（可选）
	}

	data, err := sonic.Marshal(request)
	if err != nil {
		return err
	}

	baseResponse := new(BaseResponse)
	if code, status, err := exec(rawURL, "", data, baseResponse); err != nil {
		logger.App().Errorf("exec [%s] error : %s", MethodEditMessageText.String(), err.Error())
		return err
	} else {
		if code != http.StatusOK {
			return errors.New(fmt.Sprintf("request %s failed : %d - %s", rawURL, code, status))
		}
	}

	logger.App().Infof("base response success : %+v", *baseResponse)

	return nil
}
