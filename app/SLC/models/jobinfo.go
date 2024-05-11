package models

import (
	"SmartLinkProject/common/models"
	// "gorm.io/gorm"
)

type Jobinfo struct {
	JobId    int         `gorm:"primaryKey;autoIncrement;comment:任务ID" json:"jobId"`
	DeviceId int         `json:"deviceId" gorm:"comment:设备ID" `
	TypeId   int         `json:"typeId" gorm:"comment:类型ID"`
	LocId    int         `json:"locId" gorm:"comment:位置ID"`
	Loc      *SysLoc     `json:"loc"`
	Type     *Devicetype `json:"type"`
	Device   *Device     `json:"device"`

	models.ModelTime
}

func (*Jobinfo) TableName() string {
	return "jobinfo"
}

func (e *Jobinfo) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *Jobinfo) GetId() interface{} {
	return e.JobId
}
