package apis

import (
	"SmartLinkProject/app/SLC/service"
	"SmartLinkProject/app/SLC/service/dto"
	"github.com/gin-gonic/gin"
	"github.com/go-admin-team/go-admin-core/sdk/api"
	"log"
	"net/http"
)

type Home struct {
	api.Api
}

func (e Home) GetHomeStats(c *gin.Context) {
	log.Print("进入注册的地方")
	// 从请求中解析HomeStatsReq结构体
	var req dto.HomeStatsReq
	if err := c.ShouldBind(&req); err != nil {
		// 如果绑定失败，返回错误信息
		c.JSON(http.StatusBadRequest, gin.H{"message": "请求参数绑定失败"})
		return
	}

	// 创建SysHome服务实例
	s := service.SysHome{}

	// 调用服务层的方法获取统计信息
	stats, err := s.GetHomeStats(&req)
	if err != nil {
		// 如果发生错误，返回错误信息
		c.JSON(http.StatusInternalServerError, gin.H{"message": "获取统计信息失败"})
		return
	}
	// 成功返回统计结果
	c.JSON(http.StatusOK, stats)
}
