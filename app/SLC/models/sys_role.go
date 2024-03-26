package models

import "SmartLinkProject/common/models"

type SysRole struct {
	RoleId    int        `json:"roleId" gorm:"primaryKey;autoIncrement"` // 角色编码
	RoleName  string     `json:"roleName" gorm:"size:128;"`              // 角色名称
	Status    string     `json:"status" gorm:"size:4;"`                  // 状态 1禁用 2正常
	RoleKey   string     `json:"roleKey" gorm:"size:128;"`               //角色代码
	RoleSort  int        `json:"roleSort" gorm:""`                       //角色排序
	Flag      string     `json:"flag" gorm:"size:128;"`                  //
	Remark    string     `json:"remark" gorm:"size:255;"`                //备注
	Admin     bool       `json:"admin" gorm:"size:4;"`
	DataScope string     `json:"dataScope" gorm:"size:128;"`
	Params    string     `json:"params" gorm:"-"`
	MenuIds   []int      `json:"menuIds" gorm:"-"`
	LocIds    []int      `json:"locIds" gorm:"-"`
	SysLoc    []SysLoc   `json:"sysLoc" gorm:"many2many:sys_role_loc;foreignKey:RoleId;joinForeignKey:role_id;references:LocId;joinReferences:loc_id;"`
	SysMenu   *[]SysMenu `json:"sysMenu" gorm:"many2many:sys_role_menu;foreignKey:RoleId;joinForeignKey:role_id;references:MenuId;joinReferences:menu_id;"`
	models.ControlBy
	models.ModelTime
}

func (*SysRole) TableName() string {
	return "sys_role"
}

func (e *SysRole) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *SysRole) GetId() interface{} {
	return e.RoleId
}
