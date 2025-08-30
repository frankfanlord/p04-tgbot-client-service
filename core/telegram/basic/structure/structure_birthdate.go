package structure

// Birthdate 描述用户的出生日期。
type Birthdate struct {
	Day   int `json:"day"`   //  用户出生日期；1-31
	Month int `json:"month"` // 用户出生月份；1-12
	Year  int `json:"year"`  // 可选。用户出生年份
}
