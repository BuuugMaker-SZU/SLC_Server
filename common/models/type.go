package models

import "gorm.io/gorm/schema"

type ActiveRecord interface {
	schema.Tabler
	Generate() ActiveRecord
	GetId() interface{}
}
