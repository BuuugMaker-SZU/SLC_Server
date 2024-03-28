package models

import "SmartLinkProject/common/models"

type SysLoc struct {
	LocId    int    `json:"locId" gorm:"primaryKey;autoIncrement;"` //部门编码
	ParentId int    `json:"parentId" gorm:""`                       //上级部门
	LocPath  string `json:"locPath" gorm:"size:255;"`               //
	LocName  string `json:"locName"  gorm:"size:128;"`              //部门名称
	Sort     int    `json:"sort" gorm:"size:4;"`                    //排序
	Leader   string `json:"leader" gorm:"size:128;"`                //负责人
	Phone    string `json:"phone" gorm:"size:11;"`                  //手机
	Email    string `json:"email" gorm:"size:64;"`                  //邮箱
	Status   int    `json:"status" gorm:"size:4;"`                  //状态
	models.ControlBy
	models.ModelTime
	DataScope string   `json:"dataScope" gorm:"-"`
	Params    string   `json:"params" gorm:"-"`
	Children  []SysLoc `json:"children" gorm:"-"`
}

func (*SysLoc) TableName() string {
	return "sys_loc"
}

func (e *SysLoc) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *SysLoc) GetId() interface{} {
	return e.LocId
}
