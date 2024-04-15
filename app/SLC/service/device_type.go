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

type DeviceType struct {
	service.Service
}

// Get 获取DeviceType对象
func (e *DeviceType) Get(d *dto.DeviceTypeGetReq, model *models.Devicetype) error {
	var err error
	var data models.Devicetype

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

// Insert 创建DeviceType对象
func (e *DeviceType) Insert(c *dto.DeviceTypeInsertReq) error {
	var err error
	var data models.Devicetype
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
	typePath := pkg.IntToString(data.TypeId) + "/"
	if data.ParentId != 0 {
		var typeP models.Devicetype
		tx.First(&typeP, data.ParentId)
		typePath = typeP.TypePath + typePath
	} else {
		typePath = "/0/" + typePath
	}
	var mp = map[string]string{}
	mp["type_Path"] = typePath
	if err := tx.Model(&data).Update("type_Path", typePath).Error; err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Update 修改DeviceType对象
func (e *DeviceType) Update(c *dto.DeviceTypeUpdateReq) error {
	var err error
	var model = models.Devicetype{}
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

	typePath := pkg.IntToString(model.TypeId) + "/"
	if model.ParentId != 0 {
		var typeP models.Devicetype
		tx.First(&typeP, model.ParentId)
		typePath = typeP.TypePath + typePath
	} else {
		typePath = "/0/" + typePath
	}
	model.TypePath = typePath

	db := tx.Save(&model)
	if err = db.Error; err != nil {
		e.Log.Errorf("UpdateDeviceType error:%s", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权更新该数据")
	}
	return nil
}

// Remove 删除DeviceType
func (e *DeviceType) Remove(d *dto.DeviceTypeDeleteReq) error {
	var err error
	var data models.Devicetype

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

// GetDeviceTypeList 获取设备类型列表
func (e *DeviceType) GetList(c *dto.DeviceTypeGetPageReq, list *[]models.Devicetype) error {
	var err error
	var data models.Devicetype
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

// SetDeviceTypeTree 设置设备类型树形结构
func (e *DeviceType) SetDeviceTypeTree(c *dto.DeviceTypeGetPageReq) (m []dto.DeviceTypeLabel, err error) {
	var list []models.Devicetype
	err = e.GetList(c, &list)
	m = make([]dto.DeviceTypeLabel, 0)
	for i := 0; i < len(list); i++ {
		if list[i].ParentId != 0 {
			continue
		}
		e := dto.DeviceTypeLabel{Id: list[i].TypeId, Label: list[i].TypeName}
		typesInfo := deviceTypeTreeCall(&list, e)

		m = append(m, typesInfo)
	}
	return
}

// DeviceTypeTreeCall 递归构造设备类型树形结构
func deviceTypeTreeCall(TypeList *[]models.Devicetype, deviceType dto.DeviceTypeLabel) dto.DeviceTypeLabel {
	list := *TypeList
	min := make([]dto.DeviceTypeLabel, 0)
	for j := 0; j < len(list); j++ {
		if deviceType.Id != list[j].ParentId {
			continue
		}
		mi := dto.DeviceTypeLabel{Id: list[j].TypeId, Label: list[j].TypeName, Children: []dto.DeviceTypeLabel{}}
		ms := deviceTypeTreeCall(TypeList, mi)
		min = append(min, ms)
	}
	deviceType.Children = min
	return deviceType
}

// SetDeviceTypePage 设置设备类型页面数据
func (e *DeviceType) SetDeviceTypePage(c *dto.DeviceTypeGetPageReq) (m []models.Devicetype, err error) {
	var list []models.Devicetype
	err = e.GetList(c, &list)
	for i := 0; i < len(list); i++ {
		if list[i].ParentId != 0 {
			continue
		}
		info := e.TypePageCall(&list, list[i])
		m = append(m, info)
	}
	return list, err
}

func (e *DeviceType) TypePageCall(deviceTypeList *[]models.Devicetype, menu models.Devicetype) models.Devicetype {
	list := *deviceTypeList
	min := make([]models.Devicetype, 0)
	for j := 0; j < len(list); j++ {
		if menu.TypeId != list[j].ParentId {
			continue
		}
		mt := models.Devicetype{}
		mt.TypeId = list[j].TypeId
		mt.ParentId = list[j].ParentId
		mt.TypeName = list[j].TypeName
		mt.Config = list[j].Config
		mt.Protocol = list[j].Protocol
		mt.LOGO = list[j].LOGO
		mt.CreatedAt = list[j].CreatedAt
		min = append(min, mt)
	}
	menu.Children = min
	return menu
}

// SetDeviceTypeLabel 设置设备类型标签
func (e *DeviceType) SetDeviceTypeLabel() (m []dto.DeviceTypeLabel, err error) {
	var list []models.Devicetype
	err = e.Orm.Find(&list).Error
	if err != nil {
		log.Error("find device type list error, %s", err.Error())
		return
	}
	m = make([]dto.DeviceTypeLabel, 0)
	var item dto.DeviceTypeLabel
	for i := range list {
		if list[i].ParentId != 0 {
			continue
		}
		item = dto.DeviceTypeLabel{Id: list[i].TypeId, Label: list[i].TypeName}
		info := DeviceTypeLabelCall(&list, item)
		m = append(m, info)
	}
	return
}

// DeviceTypeLabelCall
func DeviceTypeLabelCall(typeList *[]models.Devicetype, deviceType dto.DeviceTypeLabel) dto.DeviceTypeLabel {
	list := *typeList
	var mi dto.DeviceTypeLabel
	min := make([]dto.DeviceTypeLabel, 0)
	for j := 0; j < len(list); j++ {
		if deviceType.Id != list[j].ParentId {
			continue
		}
		mi = dto.DeviceTypeLabel{Id: list[j].TypeId, Label: list[j].TypeName, Children: []dto.DeviceTypeLabel{}}
		ms := DeviceTypeLabelCall(typeList, mi)
		min = append(min, ms)
	}
	deviceType.Children = min
	return deviceType
}
