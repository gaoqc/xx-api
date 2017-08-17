package models

//省份
type Province struct {
	TimeModel
	Code string `json:"code"`
	Name string `json:"name"`
}

//地市
type City struct {
	Province
	ParentCode string `json:"parent_code"`
}

//区域
type Area struct {
	City
}

//街道
type Street struct {
	City
}
