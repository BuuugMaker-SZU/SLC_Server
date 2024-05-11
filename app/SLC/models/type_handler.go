package models

import (
	"SmartLinkProject/common/models"
)

type Typehandler struct {
	TypeId        int    `json:"typeId" gorm:"primaryKey;comment:类型ID"`
	HandlerName   string `json:"handlerName" gorm:"size:255;comment:处理器英文名称"`
	HandlerCNName string `json:"handlerCNName" gorm:"size:255;comment:处理器中文名称"`

	models.ModelTime
}

func (*Typehandler) TableName() string {
	return "type_handler"
}

func (e *Typehandler) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *Typehandler) GetId() interface{} {
	return e.TypeId
}
