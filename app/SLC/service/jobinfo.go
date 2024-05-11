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

type Jobinfo struct {
	service.Service
}

func (e *Jobinfo) GetPage(c *dto.JobinfoGetPageReq, p *actions.DataPermission, list *[]models.Jobinfo, count *int64) error {
	var err error
	var data models.Jobinfo

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

func (e *Jobinfo) Get(d *dto.JobinfoById, p *actions.DataPermission, model *models.Jobinfo) error {
	var data models.Jobinfo

	err := e.Orm.Model(&data).Debug().Preload("Loc").Preload("Type").Preload("Device").
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

func (e *Jobinfo) Insert(c *dto.JobinfoInsertReq) error {
	var err error
	var data models.Jobinfo
	c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("db error: %s", err)
		return err
	}
	return nil
}

func (e *Jobinfo) Update(c *dto.JobinfoUpdateReq, p *actions.DataPermission) error {
	var err error
	var model models.Jobinfo
	db := e.Orm.First(&model, c.GetId())
	if err = db.Error; err != nil {
		e.Log.Errorf("Service Update Job error: %s", err)
		return err
	}

	c.Generate(&model)
	update := e.Orm.Model(&model).Where("job_id = ?", &model.JobId).Updates(&model)
	if err = update.Error; err != nil {
		e.Log.Errorf("db error: %s", err)
		return err
	}
	if update.RowsAffected == 0 {
		err = errors.New("update jobinfo error")
		log.Warnf("db update error")
		return err
	}
	return nil
}

func (e *Jobinfo) Remove(c *dto.JobinfoById, p *actions.DataPermission) error {
	var err error
	var data models.Jobinfo

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, c.GetId())
	if err = db.Error; err != nil {
		e.Log.Errorf("Error found in  RemoveJobinfo : %s", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权删除该数据")
	}
	return nil
}
