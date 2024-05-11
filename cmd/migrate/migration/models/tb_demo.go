package models

type TbDemo struct {
	Model
	Name string `json:"name" gorm:"type:varchar(128);comment:名称"`
	ModelTime
}

func (TbDemo) TableName() string {
	return "tb_demo"
}
