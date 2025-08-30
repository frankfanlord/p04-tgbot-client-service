package structure

// Poll 该对象包含有关民意调查的信息。
type Poll struct {
	ID                    string           `json:"id"`                      // 投票的唯一标识符
	Question              string           `json:"question"`                // 投票问题，长度为1-300个字符
	QuestionEntities      []*MessageEntity `json:"question_entities"`       // 可选。出现在问题中的特殊实体。目前仅支持自定义表情符号实体
	Options               []*PollOption    `json:"options"`                 // 投票选项列表
	TotalVoterCount       int              `json:"total_voter_count"`       // 参与投票的总用户数
	IsClosed              bool             `json:"is_closed"`               // 是否已经关闭投票
	IsAnonymous           bool             `json:"is_anonymous"`            // 是否是匿名投票
	Type                  string           `json:"type"`                    // 投票类型，当前为 “regular” 或 “quiz”
	AllowsMultipleAnswers bool             `json:"allows_multiple_answers"` // 是否允许多个选项
	CorrectOptionID       int              `json:"correct_option_id"`       // 可选。正确答案的选项编号，仅适用于 quiz 类型且已关闭或发送至私聊的投票
	Explanation           string           `json:"explanation"`             // 可选。用户选择错误答案或点击灯泡图标时显示的文字，0-200个字符
	ExplanationEntities   []*MessageEntity `json:"explanation_entities"`    // 可选。解释说明中出现的特殊实体，如用户名、URL、命令等
	OpenPeriod            int              `json:"open_period"`             // 可选。投票创建后保持开放的时间（以秒为单位）
	CloseDate             int              `json:"close_date"`              // 可选。投票自动关闭的时间点（Unix 时间戳）
}
