package dto

type HomeStatsReq struct {
	Token string `form:"token" json:"token"` // 添加token字段
}

type HomeStatsRes struct {
	DeviceCount     int            `json:"deviceCount"`     // 设备总数
	TypeCount       int            `json:"typeCount"`       // 类型总数
	AreaCount       int            `json:"areaCount"`       // 区域总数
	UserCount       int            `json:"userCount"`       // 用户总数
	TypeDeviceMap   map[string]int `json:"typeDeviceMap"`   // 每种类型的设备数量
	StatusDeviceMap map[string]int `json:"statusDeviceMap"` // 每种状态的设备数量
}
