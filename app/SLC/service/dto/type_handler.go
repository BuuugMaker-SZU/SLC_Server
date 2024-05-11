package dto

import (
	"SmartLinkProject/app/SLC/models"

	"SmartLinkProject/common/dto"
	common "SmartLinkProject/common/models"
)

type TypehandlerGetPageReq struct {
	dto.Pagination `search:"-"`
	TypeId         int    `form:"typeId" search:"type:exact;column:type_id;table:type_handler" comment:"类型ID"`
	HandlerName    string `form:"handlerName" search:"type:contains;column:handler_name;table:type_handler"  comment:"执行器英文名"`
	HandlerCNName  string `form:"handlerCNName" search:"type:contains;column:handler_cn_name;table:type_handler"  comment:"执行器中文名"`
	TypehandlerOrder
}

// 排序参数
type TypehandlerOrder struct {
	TypeIdOrder        string `search:"type:order;column:type_id;table:type_handler" form:"typeIdOrder"`
	HandlerNameOrder   string `search:"type:order;column:handler_Name;table:type_handler" form:"handlerNameOrder"`
	HandlerCNNameOrder string `search:"type:order;column:handler_CN_Name;table:type_handler" form:"handlerCNNameOrder"`
	CreatedAtOrder     string `search:"type:order;column:created_at;table:type_handler" form:"createdAtOrder"`
}

func (m *TypehandlerGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type TypehandlerInsertReq struct {
	TypeId        int    `json:"typeId" comment:"类型ID" `
	HandlerName   string `json:"handlerName" comment:"执行器英文名"`
	HandlerCNName string `json:"handlerCNName" comment:"执行器中文名"`
}

func (d *TypehandlerInsertReq) Generate(model *models.Typehandler) {
	if d.TypeId != 0 {
		model.TypeId = d.TypeId
	}
	model.HandlerName = d.HandlerName
	model.HandlerCNName = d.HandlerCNName
}

func (s *TypehandlerInsertReq) GetId() interface{} {
	return s.TypeId
}

type TypehandlerUpdateReq struct {
	TypeId        int    `json:"typeId" comment:"类型ID" `
	HandlerName   string `json:"handlerName" comment:"执行器英文名"`
	HandlerCNName string `json:"handlerCNName" comment:"执行器中文名"`
}

func (d *TypehandlerUpdateReq) Generate(model *models.Typehandler) {
	if d.TypeId != 0 {
		model.TypeId = d.TypeId
	}
	model.HandlerCNName = d.HandlerCNName
	model.HandlerName = d.HandlerName
}

func (d *TypehandlerUpdateReq) GetId() interface{} {
	return d.TypeId
}

type TypehandlerById struct {
	dto.ObjectById
}

func (s *TypehandlerById) GetId() interface{} {
	if len(s.Ids) > 0 {
		s.Ids = append(s.Ids, s.Id)
		return s.Ids
	}
	return s.Id
}

func (s *TypehandlerById) GenerateM() (common.ActiveRecord, error) {
	return &models.Typehandler{}, nil
}
