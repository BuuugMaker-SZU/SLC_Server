package dto

import (
	"SmartLinkProject/app/SLC/models"
	common "SmartLinkProject/common/models"
)

// SysLocGetPageReq 列表或者搜索使用结构体
type SysLocGetPageReq struct {
	LocId    int    `form:"locId" search:"type:exact;column:loc_id;table:sys_loc" comment:"id"`         //id
	ParentId int    `form:"parentId" search:"type:exact;column:parent_id;table:sys_loc" comment:"上级区域"` //上级区域
	LocPath  string `form:"locPath" search:"type:exact;column:loc_path;table:sys_loc" comment:""`       //路径
	LocName  string `form:"locName" search:"type:exact;column:loc_name;table:sys_loc" comment:"区域名称"`   //区域名称
	Sort     int    `form:"sort" search:"type:exact;column:sort;table:sys_loc" comment:"排序"`            //排序
	Leader   string `form:"leader" search:"type:exact;column:leader;table:sys_loc" comment:"负责人"`       //负责人
	Phone    string `form:"phone" search:"type:exact;column:phone;table:sys_loc" comment:"手机"`          //手机
	Email    string `form:"email" search:"type:exact;column:email;table:sys_loc" comment:"邮箱"`          //邮箱
	Status   string `form:"status" search:"type:exact;column:status;table:sys_loc" comment:"状态"`        //状态
}

func (m *SysLocGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type SysLocInsertReq struct {
	LocId    int    `uri:"id" comment:"编码"`                                         // 编码
	ParentId int    `json:"parentId" comment:"上级区域" vd:"?"`                         //上级区域
	LocPath  string `json:"locPath" comment:""`                                     //路径
	LocName  string `json:"locName" comment:"区域名称" vd:"len($)>0"`                   //区域名称
	Sort     int    `json:"sort" comment:"排序" vd:"?"`                               //排序
	Leader   string `json:"leader" comment:"负责人" vd:"@:len($)>0; msg:'leader不能为空'"` //负责人
	Phone    string `json:"phone" comment:"手机" vd:"?"`                              //手机
	Email    string `json:"email" comment:"邮箱" vd:"?"`                              //邮箱
	Status   int    `json:"status" comment:"状态" vd:"$>0"`                           //状态
	common.ControlBy
}

func (s *SysLocInsertReq) Generate(model *models.SysLoc) {
	if s.LocId != 0 {
		model.LocId = s.LocId
	}
	model.LocName = s.LocName
	model.ParentId = s.ParentId
	model.LocPath = s.LocPath
	model.Sort = s.Sort
	model.Leader = s.Leader
	model.Phone = s.Phone
	model.Email = s.Email
	model.Status = s.Status
}

// GetId 获取数据对应的ID
func (s *SysLocInsertReq) GetId() interface{} {
	return s.LocId
}

type SysLocUpdateReq struct {
	LocId    int    `uri:"id" comment:"编码"`                                         // 编码
	ParentId int    `json:"parentId" comment:"上级区域" vd:"?"`                         //上级部门
	LocPath  string `json:"locPath" comment:""`                                     //路径
	LocName  string `json:"locName" comment:"区域名称" vd:"len($)>0"`                   //部门名称
	Sort     int    `json:"sort" comment:"排序" vd:"?"`                               //排序
	Leader   string `json:"leader" comment:"负责人" vd:"@:len($)>0; msg:'leader不能为空'"` //负责人
	Phone    string `json:"phone" comment:"手机" vd:"?"`                              //手机
	Email    string `json:"email" comment:"邮箱" vd:"?"`                              //邮箱
	Status   int    `json:"status" comment:"状态" vd:"$>0"`                           //状态
	common.ControlBy
}

// Generate 结构体数据转化 从 SysLocControl 至 SysLoc 对应的模型
func (s *SysLocUpdateReq) Generate(model *models.SysLoc) {
	if s.LocId != 0 {
		model.LocId = s.LocId
	}
	model.LocName = s.LocName
	model.ParentId = s.ParentId
	model.LocPath = s.LocPath
	model.Sort = s.Sort
	model.Leader = s.Leader
	model.Phone = s.Phone
	model.Email = s.Email
	model.Status = s.Status
}

// GetId 获取数据对应的ID
func (s *SysLocUpdateReq) GetId() interface{} {
	return s.LocId
}

type SysLocGetReq struct {
	Id int `uri:"id"`
}

func (s *SysLocGetReq) GetId() interface{} {
	return s.Id
}

type SysLocDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *SysLocDeleteReq) GetId() interface{} {
	return s.Ids
}

type DeptLabel struct {
	Id       int         `gorm:"-" json:"id"`
	Label    string      `gorm:"-" json:"label"`
	Children []DeptLabel `gorm:"-" json:"children"`
}
