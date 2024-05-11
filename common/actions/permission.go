package actions

import (
	"errors"

	"github.com/gin-gonic/gin"
	log "github.com/go-admin-team/go-admin-core/logger"
	"github.com/go-admin-team/go-admin-core/sdk/config"
	"github.com/go-admin-team/go-admin-core/sdk/pkg"
	"github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth/user"
	"github.com/go-admin-team/go-admin-core/sdk/pkg/response"
	"gorm.io/gorm"
)

type DataPermission struct {
	DataScope string
	UserId    int
	LocId     int
	RoleId    int
}

func PermissionAction() gin.HandlerFunc {
	return func(c *gin.Context) {
		db, err := pkg.GetOrm(c)
		if err != nil {
			log.Error(err)
			return
		}

		msgID := pkg.GenerateMsgIDFromContext(c)
		var p = new(DataPermission)
		if userId := user.GetUserIdStr(c); userId != "" {
			p, err = newDataPermission(db, userId)
			if err != nil {
				log.Errorf("MsgID[%s] PermissionAction error: %s", msgID, err)
				response.Error(c, 500, err, "权限范围鉴定错误")
				c.Abort()
				return
			}
		}
		c.Set(PermissionKey, p)
		c.Next()
	}
}

func newDataPermission(tx *gorm.DB, userId interface{}) (*DataPermission, error) {
	var err error
	p := &DataPermission{}

	err = tx.Table("sys_user").
		Select("sys_user.user_id", "sys_role.role_id", "sys_user.loc_id", "sys_role.data_scope").
		Joins("left join sys_role on sys_role.role_id = sys_user.role_id").
		Where("sys_user.user_id = ?", userId).
		Scan(p).Error
	if err != nil {
		err = errors.New("获取用户数据出错 msg:" + err.Error())
		return nil, err
	}
	return p, nil
}

// 定义一个函数Permission，它接收一个表名tableName，一个指向DataPermission类型实例的指针p
func Permission(tableName string, p *DataPermission) func(db *gorm.DB) *gorm.DB {
	// 返回一个闭包函数，这个闭包函数接收一个指向gorm.DB类型实例的指针db
	return func(db *gorm.DB) *gorm.DB {
		// 检查是否启用了数据权限（EnableDP）
		if !config.ApplicationConfig.EnableDP {
			// 如果没有启用数据权限，则直接返回原始的db
			return db
		}
		// 根据DataPermission实例中的DataScope字段的值，来决定如何添加查询条件
		switch p.DataScope {
		case "2":
			// 当DataScope为2时，添加查询条件，筛选出create_by字段在特定角色的用户列表中的记录
			// 这里使用了SQL注入的预防措施，通过参数化查询来传递RoleId
			return db.Where(tableName+".create_by in (select sys_user.user_id from sys_role_loc left join sys_user on sys_user.loc_id=sys_role_loc.loc_id where sys_role_loc.role_id = ?)", p.RoleId)
		case "3":
			// 当DataScope为3时，添加查询条件，筛选出create_by字段在特定区域的用户列表中的记录
			// 同样使用参数化查询来传递DeptId
			return db.Where(tableName+".create_by in (SELECT user_id from sys_user where loc_id = ? )", p.LocId)
		case "4":
			// 当DataScope为4时，添加查询条件，筛选出create_by字段在特定地点路径下的用户列表中的记录
			// 使用LIKE操作符和通配符来匹配地点路径，并且使用pkg.IntToString函数将p.LocId转换为字符串
			return db.Where(tableName+".create_by in (SELECT user_id from sys_user where sys_user.loc_id in(select loc_id from sys_loc where loc_path like ? ))", "%/"+pkg.IntToString(p.LocId)+"/%")
		case "5":
			// 当DataScope为5时，添加查询条件，筛选出create_by字段等于特定用户ID的记录
			// 使用参数化查询来传递UserId
			return db.Where(tableName+".create_by = ?", p.UserId)
		default:
			// 如果DataScope的值不是上述任何一种，则不添加任何查询条件，直接返回原始的db
			return db
		}
	}
}

func getPermissionFromContext(c *gin.Context) *DataPermission {
	p := new(DataPermission)
	if pm, ok := c.Get(PermissionKey); ok {
		switch pm.(type) {
		case *DataPermission:
			p = pm.(*DataPermission)
		}
	}
	return p
}

// GetPermissionFromContext 提供非action写法数据范围约束
func GetPermissionFromContext(c *gin.Context) *DataPermission {
	return getPermissionFromContext(c)
}
