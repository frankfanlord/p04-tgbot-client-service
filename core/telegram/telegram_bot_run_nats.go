package telegram

import (
	"client-service/core/telegram/basic"
	"errors"
	"jarvis/logger"
	"jarvis/middleware/mq/nats"
	"time"

	"github.com/bytedance/sonic"
	"github.com/google/uuid"
	ONats "github.com/nats-io/nats.go"
)

type MissionRequest struct {
	TraceID   string `json:"trace_id"`
	UserID    int    `json:"user_id"`
	MessageID int    `json:"message_id"`
	ChatID    int    `json:"chat_id"`
	Usernmae  string `json:"username"`
	Behavior  string `json:"behavior"`
	Content   string `json:"content"`
}

type MissionResponse struct {
	Content   string         `json:"content"`
	ParseMode string         `json:"parse_mode"`
	Markup    map[string]any `json:"markup"`
	Error     string         `json:"error"`
	TraceID   string         `json:"trace_id"`
}

type PinRequest struct {
	UserID    int            `json:"user_id"`
	MessageID int            `json:"message_id"`
	ChatID    int            `json:"chat_id"`
	Content   string         `json:"content"`
	ParseMode string         `json:"parse_mode"`
	Markup    map[string]any `json:"markup"`
}

type DeleteRequest struct {
	MessageID int `json:"message_id"`
	ChatID    int `json:"chat_id"`
}

const (
	JSMissionSubject = "Search.Mission"
	JSPinSubject     = "Search.Pin"
	JSDeleteSubject  = "Search.Delete"
	JSClientQueue    = "ClientQueue"
)

var (
	_subscriptions = make([]*ONats.Subscription, 0)
)

func requestMission(userID, messageID, chatID int, username, behavior, content string) (string, string, map[string]any) {
	request := &MissionRequest{
		TraceID:   uuid.New().String(),
		UserID:    userID,
		MessageID: messageID,
		ChatID:    chatID,
		Usernmae:  username,
		Behavior:  behavior,
		Content:   content,
	}

	data, mErr := sonic.Marshal(request)
	if mErr != nil {
		logger.App().Errorf("[%s] sonic.Marshal error : %s", request.TraceID, mErr.Error())
		return "Error:" + request.TraceID, "", map[string]any{}
	}

	logger.App().Infof("=================================== request : %s", string(data))

	reply, err := nats.Instance().Request(JSMissionSubject, data, time.Second*time.Duration(3))
	if err != nil {
		logger.App().Errorf("[%s] nats.Request error : %s", request.TraceID, err.Error())
		return "Error:" + request.TraceID, "", map[string]any{}
	}

	logger.App().Infof("=================================== request : %s", string(reply.Data))

	response := new(MissionResponse)

	if err = sonic.Unmarshal(reply.Data, response); err != nil {
		logger.App().Errorf("[%s] sonic.Unmarshal error : %s", request.TraceID, err.Error())
		return "Error:" + request.TraceID, "", map[string]any{}
	}

	if response.Error != "" {
		logger.App().Errorf("[%s] response error : %s", request.TraceID, response.Error)
		return "Error:" + request.TraceID, "", map[string]any{}
	}

	return response.Content, response.ParseMode, response.Markup
}

func publishBehavior(userID, messageID, chatID int, username, behavior, content string) (string, string, map[string]any) {
	request := &MissionRequest{
		TraceID:   uuid.New().String(),
		UserID:    userID,
		MessageID: messageID,
		ChatID:    chatID,
		Usernmae:  username,
		Behavior:  behavior,
		Content:   content,
	}

	data, mErr := sonic.Marshal(request)
	if mErr != nil {
		logger.App().Errorf("[%s] sonic.Marshal error : %s", request.TraceID, mErr.Error())
		return "Error:" + request.TraceID, "", map[string]any{}
	}

	logger.App().Infof("=================================== request : %s", string(data))

	reply, err := nats.Instance().Request(JSMissionSubject, data, time.Second*time.Duration(3))
	if err != nil {
		logger.App().Errorf("[%s] nats.Request error : %s", request.TraceID, err.Error())
		return "Error:" + request.TraceID, "", map[string]any{}
	}

	logger.App().Infof("=================================== request : %s", string(reply.Data))

	response := new(MissionResponse)

	if err = sonic.Unmarshal(reply.Data, response); err != nil {
		logger.App().Errorf("[%s] sonic.Unmarshal error : %s", request.TraceID, err.Error())
		return "Error:" + request.TraceID, "", map[string]any{}
	}

	if response.Error != "" {
		logger.App().Errorf("[%s] response error : %s", request.TraceID, response.Error)
		return "Error:" + request.TraceID, "", map[string]any{}
	}

	return response.Content, response.ParseMode, response.Markup
}

func doPin(msg *ONats.Msg) {
	request := new(PinRequest)

	logger.App().Infof("=================================== request : %s", string(msg.Data))

	if err := sonic.Unmarshal(msg.Data, request); err != nil {
		logger.App().Errorf("parse error : %s", err.Error())
		return
	}

	if mid, err := basic.SendMessage(0, request.ChatID, request.Content, request.ParseMode, request.Markup); err != nil {
		logger.App().Errorf("send pin message to [%d][%d][%d] error : %s", request.UserID, request.MessageID, request.ChatID, err.Error())
		return
	} else {
		if err := basic.PinChatMessage(mid, request.ChatID); err != nil {
			logger.App().Errorf("send pin message to [%d][%d][%d] error : %s", request.UserID, request.MessageID, request.ChatID, err.Error())
			return
		}
	}
}

func doDelete(msg *ONats.Msg) {
	request := new(DeleteRequest)

	logger.App().Infof("=================================== request : %s", string(msg.Data))

	if err := sonic.Unmarshal(msg.Data, request); err != nil {
		logger.App().Errorf("parse error : %s", err.Error())
		return
	}

	// if mid, err := basic.SendMessage(0, request.ChatID, request.Content, request.ParseMode, request.Markup); err != nil {
	// 	logger.App().Errorf("send pin message to [%d][%d][%d] error : %s", request.UserID, request.MessageID, request.ChatID, err.Error())
	// 	return
	// } else {
	// 	if err := basic.PinChatMessage(mid, request.ChatID); err != nil {
	// 		logger.App().Errorf("send pin message to [%d][%d][%d] error : %s", request.UserID, request.MessageID, request.ChatID, err.Error())
	// 		return
	// 	}
	// }
}

// ====================================================================================================================

type SSMResponseType uint16

const (
	RTEdit SSMResponseType = iota // reply to edit
	RTSend
	RTPin
	RTDelete
	RTVideo
)

const (
	CSQueue                  = "ClientQueue"
	SSMissionRequestSubject  = "Search.Mission.Request"
	SSMissionResponseSubject = "Search.Mission.Response"
)

type SSMRequestMsg struct {
	TraceID  string `json:"trace_id"`
	UserID   int    `json:"user_id"`
	Username string `json:"username"`
	FLName   string `json:"fl_name"`
	ChatID   int    `json:"chat_id"`
	InMsgID  int    `json:"in_msg_id"`
	OutMsgID int    `json:"out_msg_id"`
	Behavior string `json:"behavior"`
	Content  string `json:"content"`
}

type SSMResponseMsg struct {
	Type        SSMResponseType `json:"type"`
	TraceID     string          `json:"trace_id"`
	UserID      int             `json:"user_id"`
	Username    string          `json:"username"`
	ChatID      int             `json:"chat_id"`
	InMsgID     int             `json:"in_msg_id"`
	OutMsgID    int             `json:"out_msg_id"`
	Content     string          `json:"content"`
	ParseMode   string          `json:"parse_mode"`
	Markup      map[string]any  `json:"markup"`
	VideoFileID string          `json:"video_file_id"`
	Error       string          `json:"error"`
}

func doPublish(msg *SSMRequestMsg) error {
	if msg == nil {
		return errors.New("msg is nil")
	}

	data, err := sonic.Marshal(msg)
	if err != nil {
		return err
	}

	return nats.Instance().Publish(SSMissionRequestSubject, data)
}

func doResponse(msg *ONats.Msg) {
	response := new(SSMResponseMsg)

	logger.App().Infof("=================================== request : %s", string(msg.Data))

	if err := sonic.Unmarshal(msg.Data, response); err != nil {
		logger.App().Errorf("unmarshal error : %s - %s", err.Error())
		return
	}

	switch response.Type {
	case RTEdit:
		{
			content, parseMode, markup := response.Content, response.ParseMode, response.Markup
			if response.Error != "" {
				content, parseMode, markup = "Error:"+response.TraceID, "", nil
			}

			if smErr := basic.EditMessageText(response.OutMsgID, response.ChatID, content, parseMode, markup); smErr != nil {
				logger.App().Errorf("edit message error : %d - %s - %s", response.OutMsgID, content, smErr.Error())
				return
			}
		}
	case RTSend:
		{
			content, parseMode, markup := response.Content, response.ParseMode, response.Markup
			if response.Error != "" {
				content, parseMode, markup = "Error:"+response.TraceID, "", nil
			}

			if _, err := basic.SendMessage(response.InMsgID, response.ChatID, content, parseMode, markup); err != nil {
				logger.App().Errorf("send message to reply [%d-%d-%d] error : %s - %s", response.ChatID, response.UserID, response.InMsgID, content, err.Error())
				return
			}
		}
	case RTPin:
		{
			mid, err := basic.SendMessage(0, response.ChatID, response.Content, response.ParseMode, response.Markup)
			if err != nil {
				logger.App().Errorf("send pin message to [%d][%d][%d] error : %s", response.UserID, response.InMsgID, response.ChatID, err.Error())
				return
			}

			if err = basic.PinChatMessage(mid, response.ChatID); err != nil {
				logger.App().Errorf("pin message from [%d][%d][%d] error : %s", response.UserID, response.InMsgID, response.ChatID, err.Error())
				return
			}
		}
	case RTDelete:
		{
			if err := basic.DeleteMessage(response.ChatID, response.OutMsgID); err != nil {
				logger.App().Errorf("delete message from [%d][%d][%d] error : %s", response.UserID, response.OutMsgID, response.ChatID, err.Error())
				return
			}
		}
	case RTVideo:
		{
			if _, err := basic.SendVideo(response.ChatID, response.Content, response.ParseMode, response.VideoFileID, response.Markup); err != nil {
				logger.App().Errorf("send video message to reply [%d-%d-%d] error : %s - %s - %s", response.ChatID, response.UserID, response.InMsgID, response.Content, response.VideoFileID, err.Error())
				return
			}
		}
	default:
		{
		}
	}
}
