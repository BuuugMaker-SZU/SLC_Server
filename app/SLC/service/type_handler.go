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

type Typehandler struct {
	service.Service
}

func (e *Typehandler) GetPage(c *dto.TypehandlerGetPageReq, p *actions.DataPermission, list *[]models.Typehandler, count *int64) error {
	var err error
	var data models.Typehandler

	err = e.Orm.Debug().
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

func (e *Typehandler) Get(d *dto.TypehandlerById, p *actions.DataPermission, model *models.Typehandler) error {
	var data models.Typehandler

	err := e.Orm.Model(&data).Debug().
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

func (e *Typehandler) Insert(c *dto.TypehandlerInsertReq) error {
	var err error
	var data models.Typehandler
	c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("db error: %s", err)
		return err
	}
	return nil
}

func (e *Typehandler) Update(c *dto.TypehandlerUpdateReq, p *actions.DataPermission) error {
	var err error
	var model models.Typehandler
	db := e.Orm.First(&model, c.GetId())
	if err = db.Error; err != nil {
		e.Log.Errorf("Service Update type_handler error: %s", err)
		return err
	}

	c.Generate(&model)
	update := e.Orm.Model(&model).Where("type_id = ?", &model.TypeId).Updates(&model)
	if err = update.Error; err != nil {
		e.Log.Errorf("db error: %s", err)
		return err
	}
	if update.RowsAffected == 0 {
		err = errors.New("update Typehandler error")
		log.Warnf("db update error")
		return err
	}
	return nil
}

func (e *Typehandler) Remove(c *dto.TypehandlerById, p *actions.DataPermission) error {
	var err error
	var data models.Typehandler

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, c.GetId())
	if err = db.Error; err != nil {
		e.Log.Errorf("Error found in  RemoveTypehandler : %s", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权删除该数据")
	}
	return nil
}
