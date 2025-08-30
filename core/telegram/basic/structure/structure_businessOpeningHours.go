package structure

// BusinessOpeningHours 描述商家的营业时间
type BusinessOpeningHours struct {
	TimeZoneName string                          `json:"time_zone_name"` // 时区的唯一名称
	OpeningHours []*BusinessOpeningHoursInterval `json:"opening_hours"`  // 描述营业时间段的列表
}
