package models

import "SmartLinkProject/common/models"

type Devicetype struct {
	TypeId   int    `json:"typeId" gorm:"primaryKey;autoIncrement;"` // 类型编码
	ParentId int    `json:"parentId" gorm:""`                        // 上级类型
	TypeName string `json:"typeName"  gorm:"size:128;"`              // 类型名称
	TypePath string `json:"typePath" gorm:"size:255;"`
	Config   string `json:"config" gorm:"size:255;"`   // 配置类型
	Protocol string `json:"protocol" gorm:"size:128;"` // 协议类型
	LOGO     string `json:"logo" gorm:"size:255;"`     // 图标URL

	models.ModelTime
	Params   string       `json:"params" gorm:"-"`
	Children []Devicetype `json:"children" gorm:"-"`
}

func (*Devicetype) TableName() string {
	return "device_type"
}

func (e *Devicetype) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *Devicetype) GetId() interface{} {
	return e.TypeId
}
