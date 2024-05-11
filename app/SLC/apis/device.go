// apis包定义了SmartLinkProject的API接口
package apis

import (
	"SmartLinkProject/app/SLC/models" // 导入项目中的设备模型
	"net/http"                        // 导入net包，用于处理HTTP相关功能

	"github.com/gin-gonic/gin"         // 导入Gin框架，用于处理HTTP请求
	"github.com/gin-gonic/gin/binding" // 导入Gin框架的binding功能，用于数据绑定

	"github.com/go-admin-team/go-admin-core/sdk/api" // 导入go-admin-core的api模块
	// 导入go-admin-core的response模块，虽然未使用，但可能用于响应处理
	_ "github.com/go-admin-team/go-admin-core/sdk/pkg/response"

	"SmartLinkProject/app/SLC/service"     // 导入项目中的设备服务层
	"SmartLinkProject/app/SLC/service/dto" // 导入项目中的设备数据传输对象（DTO）

	"SmartLinkProject/common/actions" // 导入项目中的权限操作
)

type Device struct {
	api.Api
}

func (e Device) GetPage(c *gin.Context) {
	// 实例化设备服务
	s := service.Device{}
	// 初始化设备分页请求数据传输对象（DTO）
	req := dto.DeviceGetPageReq{}
	// 绑定请求数据到req DTO，创建ORM和上下文，并将服务实例赋值给s
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		// 记录错误日志并返回500错误响应
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	// 从Gin上下文中获取当前用户的权限信息
	p := actions.GetPermissionFromContext(c)

	// 准备用于接收查询结果的切片和计数器变量
	list := make([]models.Device, 0)
	var count int64

	// 调用服务层的GetPage方法来获取设备分页列表和总记录数
	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		// 如果服务层返回错误，返回500错误响应
		e.Error(500, err, "查询失败")
		return
	}

	// 返回分页成功的响应，包括设备列表、当前页码、每页大小和总记录数
	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// 其他方法（如Get、Insert、Update、Delete、UpdateStatus、GetProfile）遵循类似的模式：
// 1. 实例化服务层对象
// 2. 初始化请求数据传输对象（DTO）
// 3. 绑定请求数据到DTO并创建服务层上下文
// 4. 执行服务层方法，并处理错误
// 5. 根据服务层执行结果返回相应的HTTP响应

func (e Device) Get(c *gin.Context) {
	s := service.Device{}
	req := dto.DeviceById{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req, nil).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	var object models.Device
	//数据权限检查
	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		e.Error(http.StatusUnprocessableEntity, err, "查询失败")
		return
	}
	e.OK(object, "查询成功")
}

func (e Device) Insert(c *gin.Context) {
	s := service.Device{}
	req := dto.DeviceInsertReq{}
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
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	e.OK(req.GetId(), "创建成功")
}

func (e Device) Update(c *gin.Context) {
	s := service.Device{}
	req := dto.DeviceUpdateReq{}
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

	//数据权限检查
	p := actions.GetPermissionFromContext(c)

	err = s.Update(&req, p)
	if err != nil {
		e.Logger.Error(err)
		return
	}
	e.OK(req.GetId(), "更新成功")
}

func (e Device) Delete(c *gin.Context) {
	s := service.Device{}
	req := dto.DeviceById{}
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

	// 数据权限检查
	p := actions.GetPermissionFromContext(c)

	err = s.Remove(&req, p)
	if err != nil {
		e.Logger.Error(err)
		return
	}
	e.OK(req.GetId(), "删除成功")
}

func (e Device) UpdateStatus(c *gin.Context) {
	s := service.Device{}
	req := dto.UpdateDeviceStatusReq{}
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

	//数据权限检查
	p := actions.GetPermissionFromContext(c)

	err = s.UpdateStatus(&req, p)
	if err != nil {
		e.Logger.Error(err)
		return
	}
	e.OK(req.GetId(), "更新成功")
}

func (e Device) GetProfile(c *gin.Context) {
	s := service.Device{}
	req := dto.DeviceById{}

	err := e.MakeContext(c).
		MakeOrm().
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	device := models.Device{}
	err = s.GetProfile(&req, &device)
	if err != nil {
		e.Logger.Errorf("get device profile error, %s", err.Error())
		e.Error(500, err, "获取设备信息失败")
		return
	}
	e.OK(gin.H{
		"device": device,
	}, "查询成功")
}
