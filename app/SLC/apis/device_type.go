package apis

import (
	"SmartLinkProject/app/SLC/models"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-admin-team/go-admin-core/sdk/api"
	_ "github.com/go-admin-team/go-admin-core/sdk/pkg/response"

	"SmartLinkProject/app/SLC/service"
	"SmartLinkProject/app/SLC/service/dto"
)

type DeviceType struct {
	api.Api
}

func (e DeviceType) GetPage(c *gin.Context) {
	s := service.DeviceType{}
	req := dto.DeviceTypeGetPageReq{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	list := make([]models.Devicetype, 0)
	list, err = s.SetDeviceTypePage(&req)
	if err != nil {
		e.Error(500, err, "查询失败")
		return
	}
	e.OK(list, "查询成功")
}

func (e DeviceType) Get(c *gin.Context) {
	s := service.DeviceType{}
	req := dto.DeviceTypeGetReq{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req, binding.JSON, nil).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	var object models.Devicetype

	err = s.Get(&req, &object)
	if err != nil {
		e.Error(500, err, "查询失败")
		return
	}

	e.OK(object, "查询成功")
}

func (e DeviceType) Insert(c *gin.Context) {
	s := service.DeviceType{}
	req := dto.DeviceTypeInsertReq{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req, binding.JSON).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	err = s.Insert(&req)
	if err != nil {
		e.Error(500, err, "创建失败")
		return
	}
	e.OK(req.GetId(), "创建成功")
}

func (e DeviceType) Update(c *gin.Context) {
	s := service.DeviceType{}
	req := dto.DeviceTypeUpdateReq{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	err = s.Update(&req)
	if err != nil {
		e.Error(500, err, err.Error())
		return
	}
	e.OK(req.GetId(), "更新成功")
}

func (e DeviceType) Delete(c *gin.Context) {
	s := service.DeviceType{}
	req := dto.DeviceTypeDeleteReq{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req, binding.JSON, nil).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	err = s.Remove(&req)
	if err != nil {
		e.Error(500, err, "删除失败")
		return
	}
	e.OK(req.GetId(), "删除成功")
}

// Get2Tree 设备管理 左侧类型树
func (e DeviceType) Get2Tree(c *gin.Context) {
	s := service.DeviceType{}
	req := dto.DeviceTypeGetPageReq{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req, binding.Form).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	list := make([]dto.DeviceTypeLabel, 0)
	list, err = s.SetDeviceTypeTree(&req)
	if err != nil {
		e.Error(500, err, "查询失败")
		return
	}
	e.OK(list, "")
}
