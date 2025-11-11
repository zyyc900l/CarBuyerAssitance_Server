package model

// Car 车辆信息
type Car struct {
	ImageUrl          string `json:"image_url"`          // 车外观图 请你在网上搜索
	CarName           string `json:"car_name"`           // 推荐汽车名字
	FuelConsumption   string `json:"fuel_consumption"`   // 油耗
	Power             string `json:"power"`              // 动力
	Seat              string `json:"seat"`               // 座位
	Drive             string `json:"drive"`              // 驱动
	RecommendedReason string `json:"recommended_reason"` // 推荐理由
}

// ConsultResult 购车咨询结果
type ConsultResult struct {
	Analysis string `json:"analysis"` // 分析
	Result   []Car  `json:"result"`   // 推荐结果
	Proposal string `json:"proposal"` // 总的购车建议
}

// Consult 购车咨询信息
type Consult struct {
	BudgetRange     string // 预算范围
	PreferredType   string // 偏好车型
	UseCase         string // 主要使用场景
	FuelType        string // 燃料材料偏好
	BrandPreference string // 品牌偏好
}
