package telegram

import (
	"client-service/core/telegram/basic"
	"client-service/core/telegram/basic/structure"
	"jarvis/logger"
	"time"

	"github.com/bytedance/sonic"
	"github.com/google/uuid"
)

func (b *bot) Run() {
	logger.App().Infof("======================================================================= TGBot start run =======================================================================")
	defer logger.App().Infof("======================================================================= TGBot stop run =======================================================================")

	defer close(b.done)

	ticker := time.NewTicker(time.Second * time.Duration(5))
	defer ticker.Stop()

	for {
		done := false

		select {
		case <-b.close:
			{
				done = true
			}
		case data, ok := <-b.channel:
			{
				if !ok {
					done = true
					break
				}

				go parse(data[:])
			}
		case <-ticker.C:
			{
				// keep connection alive
				if _, err := basic.GetMe(); err != nil {
					logger.App().Errorf("keep alive getMe error : %s", err.Error())
				}
			}
		}

		if done {
			break
		}
	}
}

func parse(data []byte) {
	update := new(structure.Update)
	if err := sonic.Unmarshal(data, update); err != nil {
		logger.App().Errorf("unmarshal to structure.Update error : %s - %s", err.Error(), string(data))
		return
	}

	logger.App().Infof("Update : %+v", *update)

	// 1.if it is just a normal message
	if update.Message != nil && update.Message.From.IsBot != true {

		if err := doPublish(&SSMRequestMsg{
			TraceID:  uuid.New().String(),
			UserID:   update.Message.From.ID,
			Username: update.Message.From.Username,
			FLName: func(fn, ln string) string {
				if ln != "" {
					return ln
				}
				return fn
			}(update.Message.From.FirstName, update.Message.From.LastName),
			ChatID:   update.Message.Chat.ID,
			InMsgID:  update.Message.MessageID,
			OutMsgID: 0,
			Behavior: "",
			Content:  update.Message.Text,
		}); err != nil {
			logger.App().Errorf("publish request message : %s ", err.Error())
			return
		}

		return
	}

	// 2.if it is a edited message
	if update.EditedMessage != nil {
		return
	}

	// 3.if it is a callback query
	if update.CallbackQuery != nil && update.CallbackQuery.From.IsBot != true {

		inMessageID := 0
		content := ""
		if update.CallbackQuery.Message.ReplyToMessage != nil {
			inMessageID = update.CallbackQuery.Message.ReplyToMessage.MessageID
			content = update.CallbackQuery.Message.ReplyToMessage.Text
		}

		if err := doPublish(&SSMRequestMsg{
			TraceID:  uuid.New().String(),
			UserID:   update.CallbackQuery.From.ID,
			Username: update.CallbackQuery.From.Username,
			FLName: func(fn, ln string) string {
				if ln != "" {
					return ln
				}
				return fn
			}(update.CallbackQuery.From.FirstName, update.CallbackQuery.From.LastName),
			ChatID:   update.CallbackQuery.Message.Chat.ID,
			InMsgID:  inMessageID,
			OutMsgID: update.CallbackQuery.Message.MessageID,
			Behavior: update.CallbackQuery.Data,
			Content:  content,
		}); err != nil {
			logger.App().Errorf("publish request message : %s ", err.Error())
			return
		}
	}
}
