package structure

// PollAnswer 表示一个用户在非匿名投票中的回答
type PollAnswer struct {
	PollID    string `json:"poll_id"`    // 投票的唯一标识符
	VoterChat *Chat  `json:"voter_chat"` // （可选）如果投票者是匿名的，表示更改答案的聊天
	User      *User  `json:"user"`       // （可选）如果投票者不是匿名的，表示更改答案的用户
	OptionIDs []int  `json:"option_ids"` // 所选答案选项的编号（从0开始），如果撤回了投票，则可能为空
}
