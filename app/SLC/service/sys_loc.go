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

// GetPage 获取Sysloc列表
//func (e *Sysloc) GetPage(c *dto.SyslocGetPageReq, list *[]models.Sysloc) error {
//	var err error
//	var data models.Sysloc
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
	locPath := pkg.IntToString(data.LocId) + "/"
	if data.ParentId != 0 {
		var locP models.SysLoc
		tx.First(&locP, data.ParentId)
		locPath = locP.LocPath + locPath
	} else {
		locPath = "/0/" + locPath
	}
	var mp = map[string]string{}
	mp["loc_path"] = locPath
	if err := tx.Model(&data).Update("loc_path", locPath).Error; err != nil {
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

	locPath := pkg.IntToString(model.LocId) + "/"
	if model.ParentId != 0 {
		var locP models.SysLoc
		tx.First(&locP, model.ParentId)
		locPath = locP.LocPath + locPath
	} else {
		locPath = "/0/" + locPath
	}
	model.LocPath = locPath
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
func (e *SysLoc) SetLocTree(c *dto.SysLocGetPageReq) (m []dto.LocLabel, err error) {
	var list []models.SysLoc
	err = e.getList(c, &list)

	m = make([]dto.LocLabel, 0)
	for i := 0; i < len(list); i++ {
		if list[i].ParentId != 0 {
			continue
		}
		e := dto.LocLabel{}
		e.Id = list[i].LocId
		e.Label = list[i].LocName
		locsInfo := locTreeCall(&list, e)

		m = append(m, locsInfo)
	}
	return
}

// Call 递归构造组织数据
func locTreeCall(locList *[]models.SysLoc, loc dto.LocLabel) dto.LocLabel {
	list := *locList
	min := make([]dto.LocLabel, 0)
	for j := 0; j < len(list); j++ {
		if loc.Id != list[j].ParentId {
			continue
		}
		mi := dto.LocLabel{Id: list[j].LocId, Label: list[j].LocName, Children: []dto.LocLabel{}}
		ms := locTreeCall(locList, mi)
		min = append(min, ms)
	}
	loc.Children = min
	return loc
}

// SetLocPage 设置loc页面数据
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

func (e *SysLoc) LocPageCall(loclist *[]models.SysLoc, menu models.SysLoc) models.SysLoc {
	list := *loclist
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
		ms := e.LocPageCall(loclist, mi)
		min = append(min, ms)
	}
	menu.Children = min
	return menu
}

// GetWithRoleId 获取角色的区域ID集合
func (e *SysLoc) GetWithRoleId(roleId int) ([]int, error) {
	locIds := make([]int, 0)
	locList := make([]dto.LocIdList, 0)
	if err := e.Orm.Table("sys_role_loc").
		Select("sys_role_loc.loc_id").
		Joins("LEFT JOIN sys_loc on sys_loc.loc_id=sys_role_loc.loc_id").
		Where("role_id = ? ", roleId).
		Where(" sys_role_loc.loc_id not in(select sys_loc.parent_id from sys_role_loc LEFT JOIN sys_loc on sys_loc.loc_id=sys_role_loc.loc_id where role_id =? )", roleId).
		Find(&locList).Error; err != nil {
		return nil, err
	}
	for i := 0; i < len(locList); i++ {
		locIds = append(locIds, locList[i].LocId)
	}
	return locIds, nil
}

func (e *SysLoc) SetLocLabel() (m []dto.LocLabel, err error) {
	list := make([]models.SysLoc, 0)
	err = e.Orm.Find(&list).Error
	if err != nil {
		log.Error("find loc list error, %s", err.Error())
		return
	}
	m = make([]dto.LocLabel, 0)
	var item dto.LocLabel
	for i := range list {
		if list[i].ParentId != 0 {
			continue
		}
		item = dto.LocLabel{}
		item.Id = list[i].LocId
		item.Label = list[i].LocName
		locInfo := LocLabelCall(&list, item)
		m = append(m, locInfo)
	}
	return
}

// LocLabelCall
func LocLabelCall(locList *[]models.SysLoc, loc dto.LocLabel) dto.LocLabel {
	list := *locList
	var mi dto.LocLabel
	min := make([]dto.LocLabel, 0)
	for j := 0; j < len(list); j++ {
		if loc.Id != list[j].ParentId {
			continue
		}
		mi = dto.LocLabel{Id: list[j].LocId, Label: list[j].LocName, Children: []dto.LocLabel{}}
		ms := LocLabelCall(locList, mi)
		min = append(min, ms)
	}
	loc.Children = min
	return loc
}
