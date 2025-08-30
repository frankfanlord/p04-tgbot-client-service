package structure

type ExternalReplyInfo struct {
	Origin             *MessageOrigin      `json:"origin"`                         // 被回复消息的来源信息
	Chat               *Chat               `json:"chat,omitempty"`                 // 原始消息所属的聊天，仅当为超级群组或频道时可用
	MessageID          int                 `json:"message_id,omitempty"`           // 原始聊天中的唯一消息标识符，仅当为超级群组或频道时可用
	LinkPreviewOptions *LinkPreviewOptions `json:"link_preview_options,omitempty"` // 如果是文本消息，则为链接预览的生成选项
	Animation          *Animation          `json:"animation,omitempty"`            // 动画消息，包含动画相关信息
	Audio              *Audio              `json:"audio,omitempty"`                // 音频消息，包含音频文件信息
	Document           *Document           `json:"document,omitempty"`             // 普通文件消息，包含文件信息
	PaidMedia          *PaidMediaInfo      `json:"paid_media,omitempty"`           // 包含付费媒体的消息，付费媒体信息
	Photo              []*PhotoSize        `json:"photo,omitempty"`                // 图片消息，包含图片的多种尺寸信息
	Sticker            *Sticker            `json:"sticker,omitempty"`              // 表情贴纸消息，包含贴纸信息
	Story              *Story              `json:"story,omitempty"`                // 转发的限时动态消息
	Video              *Video              `json:"video,omitempty"`                // 视频消息，包含视频相关信息
	VideoNote          *VideoNote          `json:"video_note,omitempty"`           // 视频便笺消息，包含视频便笺信息
	Voice              *Voice              `json:"voice,omitempty"`                // 语音消息，包含语音文件信息
	HasMediaSpoiler    bool                `json:"has_media_spoiler,omitempty"`    // 如果媒体被剧透动画覆盖，则为 true
	Contact            *Contact            `json:"contact,omitempty"`              // 联系人共享消息，包含联系人信息
	Dice               *Dice               `json:"dice,omitempty"`                 // 掷骰子消息，包含随机值
	Game               *Game               `json:"game,omitempty"`                 // 游戏消息，包含游戏信息
	Giveaway           *Giveaway           `json:"giveaway,omitempty"`             // 预定的赠品活动消息
	GiveawayWinners    *GiveawayWinners    `json:"giveaway_winners,omitempty"`     // 公布中奖者的赠品活动消息
	Invoice            *Invoice            `json:"invoice,omitempty"`              // 支付发票消息，包含发票信息
	Location           *Location           `json:"location,omitempty"`             // 位置共享消息，包含位置信息
	Poll               *Poll               `json:"poll,omitempty"`                 // 原生投票消息，包含投票信息
	Venue              *Venue              `json:"venue,omitempty"`                // 场所消息，包含场所信息
}
