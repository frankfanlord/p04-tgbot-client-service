package structure

// MaybeInaccessibleMessage 此对象描述了机器人无法访问的消息
type MaybeInaccessibleMessage struct {
	Message
}

func (mim *MaybeInaccessibleMessage) IsInaccessibleMessage() bool {
	return mim.Date == 0
}

func (mim *MaybeInaccessibleMessage) ThenMessage() Message {
	return mim.Message
}

func (mim *MaybeInaccessibleMessage) ThenInaccessibleMessage() InaccessibleMessage {
	return InaccessibleMessage{
		MessageID: mim.Message.MessageID,
		Date:      mim.Message.Date,
		Chat:      mim.Message.Chat,
	}
}

type InaccessibleMessage struct {
	MessageID int   `json:"message_id"` // 聊天中的唯一消息标识符
	Date      int   `json:"date"`       // 消息发送的 Unix 时间戳
	Chat      *Chat `json:"chat"`       // 消息所属的聊天
}
