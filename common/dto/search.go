package dto

import (
	"SmartLinkProject/common/global"

	"github.com/go-admin-team/go-admin-core/tools/search"
	"gorm.io/gorm"
)

type GeneralDelDto struct {
	Id  int   `uri:"id" json:"id" validate:"required"`
	Ids []int `json:"ids"`
}

func (g GeneralDelDto) GetIds() []int {
	ids := make([]int, 0)
	if g.Id != 0 {
		ids = append(ids, g.Id)
	}
	if len(g.Ids) > 0 {
		for _, id := range g.Ids {
			if id > 0 {
				ids = append(ids, id)
			}
		}
	} else {
		if g.Id > 0 {
			ids = append(ids, g.Id)
		}
	}
	if len(ids) <= 0 {
		//方式全部删除
		ids = append(ids, 0)
	}
	return ids
}

type GeneralGetDto struct {
	Id int `uri:"id" json:"id" validate:"required"`
}

// MakeCondition 函数接受一个查询条件接口 q，并返回一个函数，
// 该返回的函数接受一个 *gorm.DB 类型的 db 参数，并返回一个更新后的 *gorm.DB。
func MakeCondition(q interface{}) func(db *gorm.DB) *gorm.DB {
	// 返回一个闭包函数，这个闭包可以访问外部的 q 参数。
	return func(db *gorm.DB) *gorm.DB {
		// 创建一个用于存储解析后的查询条件的实例。
		condition := &search.GormCondition{
			GormPublic: search.GormPublic{},         // 初始化 GormPublic 字段。
			Join:       make([]*search.GormJoin, 0), // 初始化 Join 字段，用于存储连接条件。
		}
		// 使用全局的数据库驱动和传入的查询条件接口解析查询，并填充到 condition 结构中。
		search.ResolveSearchQuery(global.Driver, q, condition)

		// 遍历解析出的连接条件，并将它们应用到 db 上。
		for _, join := range condition.Join {
			if join == nil {
				continue // 如果 join 是 nil，则跳过当前循环。
			}
			db = db.Joins(join.JoinOn) // 添加 Join 语句到 db 查询构建器。

			// 遍历 join 中的 Where 条件并添加到 db。
			for k, v := range join.Where {
				db = db.Where(k, v...)
			}
			// 遍历 join 中的 Or 条件并添加到 db。
			for k, v := range join.Or {
				db = db.Or(k, v...)
			}
			// 遍历 join 中的 Order 条件并添加到 db。
			for _, o := range join.Order {
				db = db.Order(o)
			}
		}

		// 遍历 condition 中的 Where 条件并添加到 db。
		for k, v := range condition.Where {
			db = db.Where(k, v...)
		}
		// 遍历 condition 中的 Or 条件并添加到 db。
		for k, v := range condition.Or {
			db = db.Or(k, v...)
		}
		// 遍历 condition 中的 Order 条件并添加到 db。
		for _, o := range condition.Order {
			db = db.Order(o)
		}

		// 返回最终构建的查询构建器 db。
		return db
	}
}

func Paginate(pageSize, pageIndex int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		offset := (pageIndex - 1) * pageSize
		if offset < 0 {
			offset = 0
		}
		return db.Offset(offset).Limit(pageSize)
	}
}
