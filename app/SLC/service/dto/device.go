package dto

import (
	"SmartLinkProject/app/SLC/models"

	"SmartLinkProject/common/dto"
	common "SmartLinkProject/common/models"
)

type DeviceGetPageReq struct {
	dto.Pagination `search:"-"`
	DeviceId       int    `form:"deviceId" search:"type:exact;column:device_id;table:device" comment:"设备ID"`
	TypeId         int    `form:"typeId" search:"type:exact;column:type_id;table:device" comment:"类型ID"`
	LocId          int    `form:"locId" search:"type:exact;column:loc_id;table:device" comment:"位置ID"`
	DeviceName     string `form:"deviceName" search:"type:contains;column:device_name;table:device" comment:"设备名称"`
	Status         string `form:"status" search:"type:exact;column:status;table:device" comment:"状态"`
	QRCode         string `form:"qRCode" search:"type:contains;column:qr_code;table:device" comment:"二维码"`
	SNCode         string `form:"sNCode" search:"type:contains;column:sn_code;table:device" comment:"序列号"`

	DLocJoin `search:"type:left;on:loc_id:loc_id;table:device;join:sys_loc"`
	TypeJoin `search:"type:left;on:type_id:type_id;table:device;join:device_type"`
	DeviceOrder
}

// DeviceOrder 设备排序参数
type DeviceOrder struct {
	DeviceIdOrder   string `search:"type:order;column:device_id;table:device" form:"deviceIdOrder"`
	DeviceNameOrder string `search:"type:order;column:device_name;table:device" form:"deviceNameOrder"`
	StatusOrder     string `search:"type:order;column:status;table:device" form:"statusOrder"`
	CreatedAtOrder  string `search:"type:order;column:created_at;table:device" form:"createdAtOrder"`
}

type DLocJoin struct {
	LocId string `search:"type:contains;column:loc_path;table:sys_loc" form:"locId"`
}

type TypeJoin struct {
	TypeId string `search:"type:contains;column:type_path;table:device_type" form:"typeId"`
}

func (m *DeviceGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type UpdateDeviceStatusReq struct {
	DeviceId int    `json:"device" comment:"设备ID" vd:"$>0"`
	Status   string `json:"status" comment:"状态" vd:"len($)>0"`
}

func (s *UpdateDeviceStatusReq) GetId() interface{} {
	return s.DeviceId
}

func (s *UpdateDeviceStatusReq) Generate(model *models.Device) {
	if s.DeviceId != 0 {
		model.DeviceId = s.DeviceId
	}
	model.Status = s.Status
}

type DeviceInsertReq struct {
	DeviceId   int    `json:"deviceId" comment:"设备ID" `
	TypeId     int    `json:"typeId" comment:"类型ID" vd:"$>0"`                // 设备类型ID
	LocId      int    `json:"locId" comment:"位置ID" vd:"$>0"`                 // 设备位置ID
	DeviceName string `json:"deviceName" comment:"设备名称" vd:"len($)>0"`       // 设备名称
	DeviceData string `json:"deviceData" comment:"设备数据"`                     // 设备数据
	SNCode     string `json:"sNCode" comment:"序列号"`                          // 序列号
	Key        string `json:"key" comment:"密钥"`                              // 密钥
	Status     string `json:"status" comment:"状态" vd:"len($)>0" default:"1"` // 状态
	QRCode     string `json:"QRCode" comment:"二维码"`                          // 二维码
}

func (d *DeviceInsertReq) Generate(model *models.Device) {
	model.TypeId = d.TypeId
	model.LocId = d.LocId
	model.DeviceName = d.DeviceName
	model.DeviceData = d.DeviceData
	model.SNCode = d.SNCode
	model.Key = d.Key
	model.Status = d.Status
	model.QRCode = d.QRCode
}

func (s *DeviceInsertReq) GetId() interface{} {
	return s.DeviceId
}

type DeviceUpdateReq struct {
	DeviceId   int    `json:"deviceId" comment:"设备ID"`        // 设备ID
	TypeId     int    `json:"typeId" comment:"类型ID" vd:"$>0"` // 设备类型ID
	LocId      int    `json:"locId" comment:"位置ID" vd:"$>0"`  // 设备位置ID
	DeviceName string `json:"deviceName" comment:"设备名称"`      // 设备名称
	DeviceData string `json:"deviceData" comment:"设备数据"`      // 设备数据
	SNCode     string `json:"sNCode" comment:"序列号"`           // 序列号
	Key        string `json:"key" comment:"密钥"`               // 密钥
	Status     string `json:"status" comment:"状态"`            // 状态
	QRCode     string `json:"QRCode" comment:"二维码"`           // 二维码
	Remark     string `json:"remark" comment:"备注"`            // 备注
	// 可以添加更多的字段，如控制权限等
}

func (d *DeviceUpdateReq) Generate(model *models.Device) {
	if d.DeviceId != 0 {
		model.DeviceId = d.DeviceId
	}
	model.TypeId = d.TypeId
	model.LocId = d.LocId
	model.DeviceName = d.DeviceName
	model.DeviceData = d.DeviceData
	model.SNCode = d.SNCode
	model.Key = d.Key
	model.Status = d.Status
	model.QRCode = d.QRCode
}

func (d *DeviceUpdateReq) GetId() interface{} {
	return d.DeviceId
}

type DeviceById struct {
	dto.ObjectById
}

func (s *DeviceById) GetId() interface{} {
	if len(s.Ids) > 0 {
		s.Ids = append(s.Ids, s.Id)
		return s.Ids
	}
	return s.Id
}

func (s *DeviceById) GenerateM() (common.ActiveRecord, error) {
	return &models.Device{}, nil
}
