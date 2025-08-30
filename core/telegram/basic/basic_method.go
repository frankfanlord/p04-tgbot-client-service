package basic

// Method method of action
type Method string

func (m Method) String() string { return string(m) }

const (
	MethodSetWebhook            Method = "setWebhook"
	MethodDeleteWebhook         Method = "deleteWebhook"
	MethodGetWebhookInfo        Method = "getWebhookInfo"
	MethodSetMyCommands         Method = "setMyCommands"
	MethodDeleteMyCommands      Method = "deleteMyCommands"
	MethodSetMyName             Method = "setMyName"
	MethodSetMyDescription      Method = "setMyDescription"
	MethodSetMyShortDescription Method = "setMyShortDescription"
	MethodGetMyName             Method = "getMyName"
	MethodGetMyDescription      Method = "getMyDescription"
	MethodGetMyShortDescription Method = "getMyShortDescription"
	MethodGetMe                 Method = "getMe"
	MethodSendMessage           Method = "sendMessage"
	MethodEditMessageText       Method = "editMessageText"
	MethodPinChatMessage        Method = "pinChatMessage"
	MethodDeleteMessage         Method = "deleteMessage"
	MethodSendVideo             Method = "sendVideo"
)
