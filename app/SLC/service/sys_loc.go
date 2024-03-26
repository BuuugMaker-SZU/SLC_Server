package service

import (
	"SmartLinkProject/app/SLC/models"
	"errors"

	log "github.com/go-admin-team/go-admin-core/logger"
	"github.com/go-admin-team/go-admin-core/sdk/pkg"

	"gorm.io/gorm"

	"SmartLinkProject/app/SLC/service/dto"
	cDto "SmartLinkProject/common/dto"

	"github.com/go-admin-team/go-admin-core/sdk/service"
)

type SysLoc struct {
	service.Service
}

// GetPage 获取SysDept列表
//func (e *SysDept) GetPage(c *dto.SysDeptGetPageReq, list *[]models.SysDept) error {
//	var err error
//	var data models.SysDept
//
//	err = e.Orm.Model(&data).
//		Scopes(
//			cDto.MakeCondition(c.GetNeedSearch()),
//		).
//		Find(list).Error
//	if err != nil {
//		e.Log.Errorf("db error:%s", err)
//		return err
//	}
//	return nil
//}

// Get 获取SysLoc对象
func (e *SysLoc) Get(d *dto.SysLocGetReq, model *models.SysLoc) error {
	var err error
	var data models.SysLoc

	db := e.Orm.Model(&data).
		First(model, d.GetId())
	err = db.Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("db error:%s", err)
		return err
	}
	if err = db.Error; err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建SysLoc对象
func (e *SysLoc) Insert(c *dto.SysLocInsertReq) error {
	var err error
	var data models.SysLoc
	c.Generate(&data)
	tx := e.Orm.Debug().Begin()
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()
	err = tx.Create(&data).Error
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	deptPath := pkg.IntToString(data.LocId) + "/"
	if data.ParentId != 0 {
		var deptP models.SysLoc
		tx.First(&deptP, data.ParentId)
		deptPath = deptP.LocPath + deptPath
	} else {
		deptPath = "/0/" + deptPath
	}
	var mp = map[string]string{}
	mp["dept_path"] = deptPath
	if err := tx.Model(&data).Update("dept_path", deptPath).Error; err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Update 修改SysLoc对象
func (e *SysLoc) Update(c *dto.SysLocUpdateReq) error {
	var err error
	var model = models.SysLoc{}
	tx := e.Orm.Debug().Begin()
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()
	tx.First(&model, c.GetId())
	c.Generate(&model)

	deptPath := pkg.IntToString(model.LocId) + "/"
	if model.ParentId != 0 {
		var deptP models.SysLoc
		tx.First(&deptP, model.ParentId)
		deptPath = deptP.LocPath + deptPath
	} else {
		deptPath = "/0/" + deptPath
	}
	model.LocPath = deptPath
	db := tx.Save(&model)
	if err = db.Error; err != nil {
		e.Log.Errorf("UpdateSysLoc error:%s", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权更新该数据")
	}
	return nil
}

// Remove 删除SysLoc
func (e *SysLoc) Remove(d *dto.SysLocDeleteReq) error {
	var err error
	var data models.SysLoc

	db := e.Orm.Model(&data).Delete(&data, d.GetId())
	if err = db.Error; err != nil {
		err = db.Error
		e.Log.Errorf("Delete error: %s", err)
		return err
	}
	if db.RowsAffected == 0 {
		err = errors.New("无权删除该数据")
		return err
	}
	return nil
}

// GetSysLocList 获取组织数据
func (e *SysLoc) getList(c *dto.SysLocGetPageReq, list *[]models.SysLoc) error {
	var err error
	var data models.SysLoc

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
		).
		Find(list).Error
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// SetLocTree 设置组织数据
func (e *SysLoc) SetLocTree(c *dto.SysLocGetPageReq) (m []dto.DeptLabel, err error) {
	var list []models.SysLoc
	err = e.getList(c, &list)

	m = make([]dto.DeptLabel, 0)
	for i := 0; i < len(list); i++ {
		if list[i].ParentId != 0 {
			continue
		}
		e := dto.DeptLabel{}
		e.Id = list[i].LocId
		e.Label = list[i].LocName
		deptsInfo := deptTreeCall(&list, e)

		m = append(m, deptsInfo)
	}
	return
}

// Call 递归构造组织数据
func deptTreeCall(deptList *[]models.SysLoc, dept dto.DeptLabel) dto.DeptLabel {
	list := *deptList
	min := make([]dto.DeptLabel, 0)
	for j := 0; j < len(list); j++ {
		if dept.Id != list[j].ParentId {
			continue
		}
		mi := dto.DeptLabel{Id: list[j].LocId, Label: list[j].LocName, Children: []dto.DeptLabel{}}
		ms := deptTreeCall(deptList, mi)
		min = append(min, ms)
	}
	dept.Children = min
	return dept
}

// SetLocPage 设置dept页面数据
func (e *SysLoc) SetLocPage(c *dto.SysLocGetPageReq) (m []models.SysLoc, err error) {
	var list []models.SysLoc
	err = e.getList(c, &list)
	for i := 0; i < len(list); i++ {
		if list[i].ParentId != 0 {
			continue
		}
		info := e.LocPageCall(&list, list[i])
		m = append(m, info)
	}
	return
}

func (e *SysLoc) LocPageCall(deptlist *[]models.SysLoc, menu models.SysLoc) models.SysLoc {
	list := *deptlist
	min := make([]models.SysLoc, 0)
	for j := 0; j < len(list); j++ {
		if menu.LocId != list[j].ParentId {
			continue
		}
		mi := models.SysLoc{}	
		mi.LocId = list[j].LocId
		mi.ParentId = list[j].ParentId
		mi.LocPath = list[j].LocPath
		mi.LocName = list[j].LocName
		mi.Sort = list[j].Sort
		mi.Leader = list[j].Leader
		mi.Phone = list[j].Phone
		mi.Email = list[j].Email
		mi.Status = list[j].Status
		mi.CreatedAt = list[j].CreatedAt
		mi.Children = []models.SysLoc{}
		ms := e.LocPageCall(deptlist, mi)
		min = append(min, ms)
	}
	menu.Children = min
	return menu
}

// GetWithRoleId 获取角色的区域ID集合
func (e *SysLoc) GetWithRoleId(roleId int) ([]int, error) {
	deptIds := make([]int, 0)
	deptList := make([]dto.LocIdList, 0)
	if err := e.Orm.Table("sys_role_dept").
		Select("sys_role_dept.dept_id").
		Joins("LEFT JOIN sys_dept on sys_dept.dept_id=sys_role_dept.dept_id").
		Where("role_id = ? ", roleId).
		Where(" sys_role_dept.dept_id not in(select sys_dept.parent_id from sys_role_dept LEFT JOIN sys_dept on sys_dept.dept_id=sys_role_dept.dept_id where role_id =? )", roleId).
		Find(&deptList).Error; err != nil {
		return nil, err
	}
	for i := 0; i < len(deptList); i++ {
		deptIds = append(deptIds, deptList[i].LocId)
	}
	return deptIds, nil
}

func (e *SysLoc) SetDeptLabel() (m []dto.DeptLabel, err error) {
	list := make([]models.SysLoc, 0)
	err = e.Orm.Find(&list).Error
	if err != nil {
		log.Error("find dept list error, %s", err.Error())
		return
	}
	m = make([]dto.DeptLabel, 0)
	var item dto.DeptLabel
	for i := range list {
		if list[i].ParentId != 0 {
			continue
		}
		item = dto.DeptLabel{}
		item.Id = list[i].LocId
		item.Label = list[i].LocName
		deptInfo := deptLabelCall(&list, item)
		m = append(m, deptInfo)
	}
	return
}

// deptLabelCall
func deptLabelCall(deptList *[]models.SysLoc, dept dto.DeptLabel) dto.DeptLabel {
	list := *deptList
	var mi dto.DeptLabel
	min := make([]dto.DeptLabel, 0)
	for j := 0; j < len(list); j++ {
		if dept.Id != list[j].ParentId {
			continue
		}
		mi = dto.DeptLabel{Id: list[j].LocId, Label: list[j].LocName, Children: []dto.DeptLabel{}}
		ms := deptLabelCall(deptList, mi)
		min = append(min, ms)
	}
	dept.Children = min
	return dept
}
