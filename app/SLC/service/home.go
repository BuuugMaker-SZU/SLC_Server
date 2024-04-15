package service

import (
	"SmartLinkProject/app/SLC/models"
	"SmartLinkProject/app/SLC/service/dto"
	"github.com/go-admin-team/go-admin-core/sdk/service"
	"log"
)

type SysHome struct {
	service.Service
}

func (s *SysHome) GetHomeStats(req *dto.HomeStatsReq) (*dto.HomeStatsRes, error) {
	// 记录请求到达
	log.Println("GetHomeStats method is called")
	// 初始化统计结果
	stats := &dto.HomeStatsRes{
		DeviceCount:     0,
		TypeCount:       0,
		AreaCount:       0,
		UserCount:       0,
		TypeDeviceMap:   make(map[string]int),
		StatusDeviceMap: make(map[string]int),
	}

	// 查询设备总数
	var deviceCount int64
	db := s.Orm.Model(&models.Device{})
	if err := db.Count(&deviceCount).Error; err != nil {
		s.Log.Errorf("db error: %s", err)
		return nil, err
	}
	stats.DeviceCount = int(deviceCount)

	// 查询类型总数
	var typeCount int64
	db = s.Orm.Model(&models.Devicetype{})
	if err := db.Count(&typeCount).Error; err != nil {
		s.Log.Errorf("db error: %s", err)
		return nil, err
	}
	stats.TypeCount = int(typeCount)

	// 查询区域总数
	var areaCount int64
	db = s.Orm.Model(&models.SysLoc{})
	if err := db.Count(&areaCount).Error; err != nil {
		s.Log.Errorf("db error: %s", err)
		return nil, err
	}
	stats.AreaCount = int(areaCount)

	// 查询用户总数
	var userCount int64
	db = s.Orm.Model(&models.SysUser{})
	if err := db.Count(&userCount).Error; err != nil {
		s.Log.Errorf("db error: %s", err)
		return nil, err
	}
	stats.UserCount = int(userCount)

	// 查询每种类型的设备数量
	var deviceTypes []struct {
		Type  string
		Count int64
	}
	// 定义一个关联查询，通过Device表的TypeID字段与DeviceType表的ID字段关联
	db = s.Orm.Model(&models.Device{})                                        // 指定Device模型
	db = db.Joins("LEFT JOIN device_type ON device.type_id = device_type.id") // 定义JOIN操作

	// 然后选择DeviceType的Type字段和计数(*)作为count，并按Type字段分组
	db.Select("device_type.type_id, COUNT(*) as count").Group("device_type.type_name")

	// 执行查询并映射结果到deviceTypes切片
	if err := db.Scan(&deviceTypes).Error; err != nil {
		s.Log.Errorf("db error: %s", err)
		return nil, err
	}

	// 遍历查询结果，将每种类型的设备数量存储到stats.TypeDeviceMap中
	for _, deviceType := range deviceTypes {
		stats.TypeDeviceMap[deviceType.Type] = int(deviceType.Count)
	}

	// 查询每种状态的设备数量
	var deviceStatuses []struct {
		Status string
		Count  int64
	}
	db = s.Orm.Model(&models.Device{})
	db.Select("status, COUNT(*) as count").Group("status").Scan(&deviceStatuses)
	for _, deviceStatus := range deviceStatuses {
		stats.StatusDeviceMap[deviceStatus.Status] = int(deviceStatus.Count)
	}

	return stats, nil
}
