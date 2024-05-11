package apis

import (
	"SmartLinkProject/app/SLC/models" // 导入项目中的设备模型
	// "encoding/json"
	"net/http" // 导入net包，用于处理HTTP相关功能
	// "strconv"

	"github.com/gin-gonic/gin"         // 导入Gin框架，用于处理HTTP请求
	"github.com/gin-gonic/gin/binding" // 导入Gin框架的binding功能，用于数据绑定

	"github.com/go-admin-team/go-admin-core/sdk/api" // 导入go-admin-core的api模块
	// 导入go-admin-core的response模块，虽然未使用，但可能用于响应处理
	_ "github.com/go-admin-team/go-admin-core/sdk/pkg/response"

	"SmartLinkProject/app/SLC/service"     // 导入项目中的设备服务层
	"SmartLinkProject/app/SLC/service/dto" // 导入项目中的设备数据传输对象（DTO）

	"SmartLinkProject/common/actions" // 导入项目中的权限操作
	// "net/url"
)

type Jobinfo struct {
	api.Api
}

func (e Jobinfo) GetPage(c *gin.Context) {
	// 实例化设备服务
	s := service.Jobinfo{}
	// 初始化设备分页请求数据传输对象（DTO）
	req := dto.JobinfoGetPageReq{}
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
	list := make([]models.Jobinfo, 0)
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

	// queryParams := url.Values{}
	// queryParams.Add("start", strconv.Itoa(req.Start))
	// queryParams.Add("length", strconv.Itoa(req.Length))
	// queryParams.Add("author", req.Author)
	// queryParams.Add("triggerStatus", req.TriggerStatus)
	// queryParams.Add("jobDesc", req.JobDesc)
	// queryParams.Add("executorHandler", req.ExecutorHandler)
	// queryParams.Add("jobGroup", "1")

	// anotherServerResponse, err := http.PostForm("http://localhost:8080/xxl-job-admin/jobinfo/pageList", queryParams)
	// if err != nil {
	// 	e.Logger.Error(err)
	// 	e.Error(500, err, "调用另一个服务器接口失败")
	// 	return
	// }
	// defer anotherServerResponse.Body.Close()

	// // 解析另一个服务器返回的数据
	// var responseData map[string]interface{}
	// err = json.NewDecoder(anotherServerResponse.Body).Decode(&responseData)
	// if err != nil {
	// 	e.Logger.Error(err)
	// 	e.Error(500, err, "解析另一个服务器响应失败")
	// 	return
	// }

	// // 返回另一个服务器接口返回的数据给前端
	// e.OK(responseData, "查询成功")
}

// 其他方法（如Get、Insert、Update、Delete、UpdateStatus、GetProfile）遵循类似的模式：
// 1. 实例化服务层对象
// 2. 初始化请求数据传输对象（DTO）
// 3. 绑定请求数据到DTO并创建服务层上下文
// 4. 执行服务层方法，并处理错误
// 5. 根据服务层执行结果返回相应的HTTP响应

func (e Jobinfo) Get(c *gin.Context) {
	s := service.Jobinfo{}
	req := dto.JobinfoById{}
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
	var object models.Jobinfo
	//数据权限检查
	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		e.Error(http.StatusUnprocessableEntity, err, "查询失败")
		return
	}
	e.OK(object, "查询成功")
}

func (e Jobinfo) Insert(c *gin.Context) {
	s := service.Jobinfo{}
	req := dto.JobinfoInsertReq{}
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

// func (e Jobinfo) InsertRemote(c *gin.Context) {
// 	s := service.Jobinfo{}
// 	req1 := dto.JobinfoInsertRemoteReq{}

// 	err := e.MakeContext(c).
// 		MakeOrm().
// 		Bind(&req1, binding.JSON).
// 		MakeService(&s.Service).
// 		Errors
// 	if err != nil {
// 		e.Logger.Error(err)
// 		e.Error(http.StatusInternalServerError, err, err.Error())
// 		return
// 	}

// 	queryParams := url.Values{}

// 	queryParams.Add("author", req1.Author)
// 	queryParams.Add("alarmEmail", req1.AlarmEmail)
// 	queryParams.Add("triggerStatus", req1.TriggerStatus)
// 	queryParams.Add("jobDesc", req1.JobDesc)
// 	queryParams.Add("scheduleConf", req1.ScheduleConf)
// 	queryParams.Add("cronGen_display", req1.CronGen_display)
// 	queryParams.Add("schedule_conf_CRON", req1.Schedule_conf_CRON)
// 	queryParams.Add("schedule_conf_FIX_RATE", req1.Schedule_conf_FIX_RATE)
// 	queryParams.Add("schedule_conf_FIX_DELAY", req1.Schedule_conf_FIX_DELAY)
// 	queryParams.Add("glueType", req1.GlueType)
// 	queryParams.Add("executorHandler", req1.ExecutorHandler)
// 	queryParams.Add("executorParam", req1.ExecutorParam)
// 	queryParams.Add("executorRouteStrategy", req1.ExecutorRouteStrategy)
// 	queryParams.Add("childJobId", req1.ChildJobId)
// 	queryParams.Add("misfireStrategy", req1.MisfireStrategy)
// 	queryParams.Add("executorBlockStrategy", req1.ExecutorBlockStrategy)
// 	queryParams.Add("executorTimeout", req1.ExecutorTimeout)
// 	queryParams.Add("executorFailRetryCount", req1.ExecutorFailRetryCount)
// 	queryParams.Add("glueRemark", req1.GlueRemark)
// 	queryParams.Add("glueSource", req1.GlueSource)
// 	queryParams.Add("jobGroup", "1")

// 	// 调用另一个服务器的接口
// 	anotherServerResponse, err := http.PostForm("http://localhost:8080/xxl-job-admin/jobinfo/add", queryParams)
// 	if err != nil {
// 		e.Logger.Error(err)
// 		e.Error(http.StatusInternalServerError, err, "调用另一个服务器接口失败")
// 		return
// 	}
// 	defer anotherServerResponse.Body.Close()

// 	// 检查HTTP响应状态码
// 	if anotherServerResponse.StatusCode != http.StatusOK {
// 		e.Logger.Errorf("服务器返回错误状态码: %d", anotherServerResponse.StatusCode)
// 		e.Error(http.StatusInternalServerError, err, "另一个服务器接口返回非200状态码")
// 		return
// 	}

// 	// 解析另一个服务器返回的数据
// 	var responseData map[string]interface{}
// 	err = json.NewDecoder(anotherServerResponse.Body).Decode(&responseData)
// 	if err != nil {
// 		e.Logger.Error(err)
// 		e.Error(http.StatusInternalServerError, err, "解析另一个服务器响应失败")
// 		return
// 	}

// 	// 返回另一个服务器接口返回的数据给前端
// 	e.OK(responseData, "创建成功")
// }

func (e Jobinfo) Update(c *gin.Context) {
	s := service.Jobinfo{}
	req := dto.JobinfoUpdateReq{}
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

func (e Jobinfo) Delete(c *gin.Context) {
	s := service.Jobinfo{}
	req := dto.JobinfoById{}
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
