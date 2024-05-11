package models

import (
	"SmartLinkProject/common/models"
	// "gorm.io/gorm"
)

type Device struct {
	DeviceId   int         `gorm:"primaryKey;autoIncrement;comment:设备ID" json:"deviceId"`
	TypeId     int         `json:"typeId" gorm:"comment:类型ID"`              // 类型ID
	LocId      int         `json:"locId" gorm:"comment:位置ID"`               // 位置ID
	DeviceName string      `json:"deviceName" gorm:"size:255;comment:设备名称"` // 设备名称
	DeviceData string      `json:"deviceData" gorm:"size:255;comment:设备数据"` // 设备数据
	SNCode     string      `json:"sNCode" gorm:"size:255;comment:序列号"`      // 序列号
	Key        string      `json:"key" gorm:"size:255;comment:密钥"`          // 密钥
	Status     string      `json:"status" gorm:"size:255;comment:状态"`       // 状态
	QRCode     string      `json:"QRCode" gorm:"size:255;comment:二维码"`      // 二维码
	Loc        *SysLoc     `json:"loc"`
	Type       *Devicetype `json:"type"`

	models.ModelTime
}

func (*Device) TableName() string {
	return "device"
}

func (e *Device) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *Device) GetId() interface{} {
	return e.DeviceId
}
