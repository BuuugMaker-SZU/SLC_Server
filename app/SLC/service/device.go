package service

import (
	"SmartLinkProject/app/SLC/models"
	"SmartLinkProject/app/SLC/service/dto"
	"errors"

	log "github.com/go-admin-team/go-admin-core/logger"
	"github.com/go-admin-team/go-admin-core/sdk/service"
	"gorm.io/gorm"

	"SmartLinkProject/common/actions"
	cDto "SmartLinkProject/common/dto"
)

type Device struct {
	service.Service
}

// GetPage 获取Device列表
func (e *Device) GetPage(c *dto.DeviceGetPageReq, p *actions.DataPermission, list *[]models.Device, count *int64) error {
	var err error
	var data models.Device

	err = e.Orm.Debug().
		Preload("Loc").
		Preload("Type").
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("db error: %s", err)
		return err
	}
	return nil
}

// Get 获取Device对象
func (e *Device) Get(d *dto.DeviceById, p *actions.DataPermission, model *models.Device) error {
	var data models.Device

	err := e.Orm.Model(&data).Debug().Preload("Loc").Preload("Type").
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("db error: %s", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error: %s", err)
		return err
	}
	return nil
}

// Insert 创建Device对象
func (e *Device) Insert(c *dto.DeviceInsertReq) error {
	var err error
	var data models.Device
	var i int64
	err = e.Orm.Model(&data).Where("device_name = ?", c.DeviceName).Count(&i).Error
	if err != nil {
		e.Log.Errorf("db error: %s", err)
		return err
	}
	if i > 0 {
		err := errors.New("设备名已存在！")
		e.Log.Errorf("db error: %s", err)
		return err
	}
	c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("db error: %s", err)
		return err
	}
	return nil
}

// Update 修改Device对象
func (e *Device) Update(c *dto.DeviceUpdateReq, p *actions.DataPermission) error {
	var err error
	var model models.Device
	db := e.Orm.First(&model, c.GetId())
	if err = db.Error; err != nil {
		e.Log.Errorf("Service UpdateDevice error: %s", err)
		return err
	}

	c.Generate(&model)
	update := e.Orm.Model(&model).Where("device_id = ?", &model.DeviceId).Updates(&model)
	if err = update.Error; err != nil {
		e.Log.Errorf("db error: %s", err)
		return err
	}
	if update.RowsAffected == 0 {
		err = errors.New("update device error")
		log.Warnf("db update error")
		return err
	}
	return nil
}

// UpdateStatus 更新设备状态
func (e *Device) UpdateStatus(c *dto.UpdateDeviceStatusReq, p *actions.DataPermission) error {
	var err error
	var model models.Device
	db := e.Orm.First(&model, c.GetId())
	if err = db.Error; err != nil {
		e.Log.Errorf("Service UpdateDevice error: %s", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权更新该数据")

	}
	err = e.Orm.Table(model.TableName()).Where("device_id =? ", c.DeviceId).Updates(c).Error
	if err != nil {
		e.Log.Errorf("Service UpdateDevice error: %s", err)
		return err
	}
	return nil
}

// Remove 删除Device对象
func (e *Device) Remove(c *dto.DeviceById, p *actions.DataPermission) error {
	var err error
	var data models.Device

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, c.GetId())
	if err = db.Error; err != nil {
		e.Log.Errorf("Error found in  RemoveDevice : %s", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权删除该数据")
	}
	return nil
}

func (e *Device) GetProfile(c *dto.DeviceById, device *models.Device) error {
	err := e.Orm.Preload("Loc").Preload("Type").First(device, c.GetId()).Error
	if err != nil {
		return err
	}
	return nil
}
