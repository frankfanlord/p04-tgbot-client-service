package structure

// Venue 表示一个场地对象
type Venue struct {
	Location        *Location `json:"location"`          // 场地位置，不能是实时位置
	Title           string    `json:"title"`             // 场地名称
	Address         string    `json:"address"`           // 场地地址
	FoursquareID    string    `json:"foursquare_id"`     // 可选。Foursquare 场地标识符
	FoursquareType  string    `json:"foursquare_type"`   // 可选。Foursquare 场地类型，例如 "arts_entertainment/default"
	GooglePlaceID   string    `json:"google_place_id"`   // 可选。Google 地点标识符
	GooglePlaceType string    `json:"google_place_type"` // 可选。Google 地点类型（支持的类型见文档）
}
