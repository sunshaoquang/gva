package manual

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type WmsKpi2023SummaryMiRouter struct {
}

// InitWmsKpi2023SummaryMiRouter 初始化 wmsKpi2023SummaryMi表 路由信息
func (s *WmsKpi2023SummaryMiRouter) InitWmsKpi2023SummaryMiRouter(Router *gin.RouterGroup) {
	wmsKpi2023SummaryMiRouter := Router.Group("wmsKpi2023SummaryMi").Use(middleware.OperationRecord())
	wmsKpi2023SummaryMiRouterWithoutRecord := Router.Group("wmsKpi2023SummaryMi")
	var wmsKpi2023SummaryMiApi = v1.ApiGroupApp.ManualApiGroup.WmsKpi2023SummaryMiApi
	{
		wmsKpi2023SummaryMiRouter.POST("createWmsKpi2023SummaryMi", wmsKpi2023SummaryMiApi.CreateWmsKpi2023SummaryMi)   // 新建wmsKpi2023SummaryMi表
		wmsKpi2023SummaryMiRouter.DELETE("deleteWmsKpi2023SummaryMi", wmsKpi2023SummaryMiApi.DeleteWmsKpi2023SummaryMi) // 删除wmsKpi2023SummaryMi表
		wmsKpi2023SummaryMiRouter.DELETE("deleteWmsKpi2023SummaryMiAll", wmsKpi2023SummaryMiApi.DeleteWmsKpi2023SummaryMiAll) // 删除wmsKpi2023SummaryMiAll表
		wmsKpi2023SummaryMiRouter.DELETE("deleteWmsKpi2023SummaryMiByIds", wmsKpi2023SummaryMiApi.DeleteWmsKpi2023SummaryMiByIds) // 批量删除wmsKpi2023SummaryMi表
		wmsKpi2023SummaryMiRouter.PUT("updateWmsKpi2023SummaryMi", wmsKpi2023SummaryMiApi.UpdateWmsKpi2023SummaryMi)    // 更新wmsKpi2023SummaryMi表
	}
	{
		wmsKpi2023SummaryMiRouterWithoutRecord.GET("findWmsKpi2023SummaryMi", wmsKpi2023SummaryMiApi.FindWmsKpi2023SummaryMi)        // 根据ID获取wmsKpi2023SummaryMi表
		wmsKpi2023SummaryMiRouterWithoutRecord.GET("getWmsKpi2023SummaryMiList", wmsKpi2023SummaryMiApi.GetWmsKpi2023SummaryMiList)  // 获取wmsKpi2023SummaryMi表列表
	}
}
