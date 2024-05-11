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

func (e *SysHome) getAreaCount(parentID int64) (int64, error) {
	var count int64
	// 查询当前节点的未删除子节点数量
	if err := e.Orm.Model(&models.SysLoc{}).Where("parent_id = ? AND deleted_at IS NULL", parentID).Count(&count).Error; err != nil {
		e.Log.Errorf("db error: %s", err)
		return 0, err
	}

	// 递归查询子节点的子节点数量
	var subCount int64
	subLocations := make([]models.SysLoc, 0)
	if err := e.Orm.Model(&models.SysLoc{}).Where("parent_id = ?", parentID).Find(&subLocations).Error; err != nil {
		e.Log.Errorf("db error: %s", err)
		return 0, err
	}
	for _, loc := range subLocations {
		if loc.DeletedAt.Valid {
			continue // 跳过已删除的节点及其子孙节点
		}
		childCount, err := e.getAreaCount(int64(loc.LocId))
		if err != nil {
			return 0, err
		}
		subCount += childCount
	}

	return count + subCount, nil
}
func (e *SysHome) getTypeCount(parentID int64) (int64, []int64, error) {
	var count int64
	var typeIDs []int64

	// 查询当前节点的未删除子节点数量
	if err := e.Orm.Model(&models.Devicetype{}).Where("parent_id = ? AND deleted_at IS NULL", parentID).Count(&count).Error; err != nil {
		e.Log.Errorf("db error: %s", err)
		return 0, nil, err
	}

	// 递归查询子节点的子节点数量
	subTypes := make([]models.Devicetype, 0)
	if err := e.Orm.Model(&models.Devicetype{}).Where("parent_id = ?", parentID).Find(&subTypes).Error; err != nil {
		e.Log.Errorf("db error: %s", err)
		return 0, nil, err
	}
	for _, tp := range subTypes {
		if tp.DeletedAt.Valid {
			continue // 跳过已删除的节点及其子孙节点
		}
		childCount, childTypeIDs, err := e.getTypeCount(int64(tp.TypeId))
		if err != nil {
			return 0, nil, err
		}
		count += childCount
		typeIDs = append(typeIDs, int64(tp.TypeId)) // 将当前节点的 type_id 添加到 typeIDs 中
		typeIDs = append(typeIDs, childTypeIDs...)  // 将子节点的 typeIDs 合并到 typeIDs 中
	}

	// 返回类型数和存在的类型 type_id
	return count, typeIDs, nil
}
func isInTypeIDs(typeID int64, typeIDs []int64) bool {
	for _, id := range typeIDs {
		if typeID == id {
			return true
		}
	}
	return false
}
func (s *SysHome) GetHomeStats(req *dto.HomeStatsReq) (*dto.HomeStatsRes, error) {
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
	if err := s.Orm.Model(&models.Device{}).Count(&deviceCount).Error; err != nil {
		// 如果发生错误，记录错误并返回
		s.Log.Errorf("db error: %s", err)
		return nil, err
	}
	stats.DeviceCount = int(deviceCount)

	// 查询类型总数
	typeCount, typeIDs, err := s.getTypeCount(0) // 根节点的父ID通常为0
	if err != nil {
		s.Log.Errorf("db error: %s", err)
		return nil, err
	}
	stats.TypeCount = int(typeCount)
	log.Println("typesIDs:", typeIDs)

	// 查询区域总数
	areaCount, err := s.getAreaCount(0) // 根节点的父ID通常为0
	if err != nil {
		s.Log.Errorf("db error: %s", err)
		return nil, err
	}
	stats.AreaCount = int(areaCount)

	// 查询用户总数
	var userCount int64
	if err := s.Orm.Model(&models.SysUser{}).Count(&userCount).Error; err != nil {
		// 如果发生错误，记录错误并返回
		s.Log.Errorf("db error: %s", err)
		return nil, err
	}
	stats.UserCount = int(userCount)

	// 查询每种类型的设备数量
	var deviceTypes []struct {
		TypeID int64
		Type   string
		Count  int64
	}

	// 执行关联查询，并将结果映射到deviceTypes切片
	if err := s.Orm.Model(&models.Device{}).
		Joins("LEFT JOIN device_type ON device.type_id = device_type.type_id").
		Select("device_type.type_id as type_id, device_type.type_name as type, COUNT(*) as count").
		Group("device_type.type_id, device_type.type_name").
		Scan(&deviceTypes).Error; err != nil {
		s.Log.Errorf("db error: %s", err)
		return nil, err
	}

	// 遍历查询结果，将每种类型的设备数量存储到stats.TypeDeviceMap中
	for _, deviceType := range deviceTypes {
		log.Println(deviceType.TypeID)
		var groupName string
		if isInTypeIDs(deviceType.TypeID, typeIDs) {
			groupName = deviceType.Type
		} else {
			groupName = "null"
		}
		stats.TypeDeviceMap[groupName] += int(deviceType.Count)
	}

	// 查询每种状态的设备数量
	var deviceStatuses []struct {
		Status string
		Count  int64
	}

	// 执行查询，并将结果映射到deviceStatuses切片
	if err := s.Orm.Model(&models.Device{}).
		Select("status, COUNT(*) as count").
		Group("status").
		Scan(&deviceStatuses).Error; err != nil {
		s.Log.Errorf("db error: %s", err)
		return nil, err
	}

	// 将每种状态的设备数量存储到stats.StatusDeviceMap中
	for _, deviceStatus := range deviceStatuses {
		stats.StatusDeviceMap[deviceStatus.Status] = int(deviceStatus.Count)
	}
	return stats, nil
}
