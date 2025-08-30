package basic

import (
	"bytes"
	"client-service/core/telegram/basic/structure"
	"errors"
	"fmt"
	"io"
	"jarvis/logger"
	"mime/multipart"
	"net/http"
	"os"

	"github.com/bytedance/sonic"
	"github.com/spf13/cast"
)

func UploadVideo(chatID any, filename, filepath string) (string, error) {
	rawURL := GenerateURL(MethodSendVideo)

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	writer.WriteField("chat_id", fmt.Sprintf("%v", chatID))
	fileWriter, _ := writer.CreateFormFile("video", filename)

	file, _ := os.Open(filepath)
	io.Copy(fileWriter, file)
	writer.Close()

	fileID := ""

	baseResponse := new(BaseResponse)
	if code, status, err := exec(rawURL, writer.FormDataContentType(), body.Bytes(), baseResponse); err != nil {
		logger.App().Errorf("exec [%s] error : %s", MethodSendMessage.String(), err.Error())
		return "", err
	} else {
		if code != http.StatusOK {
			return "", errors.New(fmt.Sprintf("request %s failed : %d - %s", rawURL, code, status))
		} else {
			if data, mErr := sonic.Marshal(baseResponse.Result); mErr != nil {
				return "", mErr
			} else {
				message := new(structure.Message)
				if err = sonic.Unmarshal(data, message); err != nil {
					return "", err
				}
				if message.Video != nil {
					fileID = message.Video.FileID
				}
			}
		}
	}

	logger.App().Infof("base response success : %+v", *baseResponse)

	return fileID, nil
}

func SendVideo[T structure.NS](chatID T, text, parseMode, videoFileID string, replyMarkup any) (int, error) {
	rawURL := GenerateURL(MethodSendVideo)

	request := &structure.SendVideoRequest[T]{
		ChatID:            chatID,      // 聊天ID或用户名（@channelusername）
		Video:             videoFileID, // 视频，可以是 file_id、HTTP URL 或 InputFile
		Caption:           text,        // 视频说明文字，支持格式化
		ParseMode:         parseMode,   // 解析模式：Markdown 或 HTML
		SupportsStreaming: true,        // 是否支持流式播放
		ReplyMarkup:       replyMarkup, // 自定义回复键盘（InlineKeyboardMarkup、ReplyKeyboardMarkup 等）
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
