package dto

import (
	"SmartLinkProject/app/SLC/models"
)

// DeviceTypeGetPageReq 列表或搜索使用结构体
type DeviceTypeGetPageReq struct {
	TypeId   int    `form:"typeId" search:"type:exact;column:type_id;table:device_type" comment:"类型ID"`        // 类型ID
	ParentId int    `form:"parentId" search:"type:exact;column:parent_id;table:device_type" comment:"上级类型"`    // 上级类型
	TypeName string `form:"typeName" search:"type:contains;column:type_name;table:device_type" comment:"类型名称"` // 类型名称
	TypePath string `form:"typePath" search:"type:exact;column:type_path;table:device_type" comment:""`        //路径
	Config   string `form:"config" search:"type:exact;column:config;table:device_type" comment:"配置类型"`         // 配置类型
	Protocol string `form:"protocol" search:"type:exact;column:protocol;table:device_type" comment:"协议类型"`     // 协议类型
	LOGO     string `form:"logo" search:"type:exact;column:logo;table:device_type" comment:"图标URL"`            // 图标URL
}

func (m *DeviceTypeGetPageReq) GetNeedSearch() interface{} {
	return *m
}

// DeviceTypeInsertReq 结构体插入请求
type DeviceTypeInsertReq struct {
	TypeId   int    `json:"typeId" comment:"类型ID"`   // 类型ID
	ParentId int    `json:"parentId" comment:"上级类型"` // 上级类型
	TypePath string `json:"typePath" comment:""`
	TypeName string `json:"typeName" comment:"类型名称"` // 类型名称
	Config   string `json:"config" comment:"配置类型"`   // 配置类型
	Protocol string `json:"protocol" comment:"协议类型"` // 协议类型
	LOGO     string `json:"logo" comment:"图标URL"`    // 图标URL
}

// DeviceTypeUpdateReq 结构体更新请求
type DeviceTypeUpdateReq struct {
	TypeId   int    `json:"typeId" comment:"类型ID"`   // 类型ID
	ParentId int    `json:"parentId" comment:"上级类型"` // 上级类型
	TypePath string `json:"typePath" comment:""`
	TypeName string `json:"typeName" comment:"类型名称"` // 类型名称
	Config   string `json:"config" comment:"配置类型"`   // 配置类型
	Protocol string `json:"protocol" comment:"协议类型"` // 协议类型
	LOGO     string `json:"logo" comment:"图标URL"`    // 图标URL
}

// DeviceTypeReq 通用设备类型请求结构体
type DeviceTypeReq struct {
	TypeId   int    `json:"typeId" comment:"类型ID"` // 类型ID
	ParentId int    `json:"parentId"`              // 上级类型
	TypeName string `json:"typeName"`              // 类型名称
	TypePath string `json:"typePath" comment:""`
	Config   string `json:"config"`   // 配置类型
	Protocol string `json:"protocol"` // 协议类型
	LOGO     string `json:"logo"`     // 图标URL
}

// Generate 方法用于生成设备类型模型
func (d *DeviceTypeReq) Generate(model *models.Devicetype) {
	if d.TypeId != 0 {
		model.TypeId = d.TypeId
	}
	model.ParentId = d.ParentId
	model.TypeName = d.TypeName
	model.TypePath = d.TypePath
	model.Config = d.Config
	model.Protocol = d.Protocol
	model.LOGO = d.LOGO
}

func (d *DeviceTypeUpdateReq) Generate(model *models.Devicetype) {
	if d.TypeId != 0 {
		model.TypeId = d.TypeId
	}
	model.ParentId = d.ParentId
	model.TypeName = d.TypeName
	model.TypePath = d.TypePath
	model.Config = d.Config
	model.Protocol = d.Protocol
	model.LOGO = d.LOGO
}

func (d *DeviceTypeInsertReq) Generate(model *models.Devicetype) {
	if d.TypeId != 0 {
		model.TypeId = d.TypeId
	}
	model.ParentId = d.ParentId
	model.TypeName = d.TypeName
	model.TypePath = d.TypePath
	model.Config = d.Config
	model.Protocol = d.Protocol
	model.LOGO = d.LOGO
}

// GetId 方法用于获取设备类型的ID
func (d *DeviceTypeReq) GetId() interface{} {
	return d.TypeId
}

// GetId 获取数据对应的ID
func (d *DeviceTypeUpdateReq) GetId() interface{} {
	return d.TypeId
}

// GetId 获取数据对应的ID
func (d *DeviceTypeInsertReq) GetId() interface{} {
	return d.TypeId
}

type DeviceTypeGetReq struct {
	Id int `uri:"id"`
}

func (d *DeviceTypeGetReq) GetId() interface{} {
	return d.Id
}

type DeviceTypeDeleteReq struct {
	Ids []int `json:"ids"`
}

func (d *DeviceTypeDeleteReq) GetId() interface{} {
	return d.Ids
}

type DeviceTypeLabel struct {
	Id       int               `gorm:"-" json:"id"`
	Label    string            `gorm:"-" json:"label"`
	Children []DeviceTypeLabel `gorm:"-" json:"children"`
}
