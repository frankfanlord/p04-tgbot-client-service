package structure

// BusinessOpeningHoursInterval 描述商家在一周中营业的时间段
type BusinessOpeningHoursInterval struct {
	OpeningMinute int `json:"opening_minute"` // 商家营业时间段的开始时间（以周一为起点的分钟数，范围：0 - 10080）
	ClosingMinute int `json:"closing_minute"` // 商家营业时间段的结束时间（以周一为起点的分钟数，范围：0 - 11520）
}
