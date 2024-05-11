package apis

import (
	"SmartLinkProject/app/SLC/service"
	"SmartLinkProject/app/SLC/service/dto"
	"github.com/gin-gonic/gin"
	"github.com/go-admin-team/go-admin-core/sdk/api"
	"net/http"
)

type Home struct {
	api.Api
}

func (e Home) GetHomeStats(c *gin.Context) {
	// 创建SysHome服务实例
	s := service.SysHome{}
	// 从请求中解析HomeStatsReq结构体
	var req dto.HomeStatsReq
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

	// 调用服务层的方法获取统计信息
	stats, err := s.GetHomeStats(&req)
	if err != nil {
		// 如果发生错误，返回错误信息
		c.JSON(http.StatusInternalServerError, gin.H{"code": 500, "msg": "获取统计信息失败", "data": nil})
		return
	}
	// 成功返回统计结果
	c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "查询成功", "data": stats})
}
