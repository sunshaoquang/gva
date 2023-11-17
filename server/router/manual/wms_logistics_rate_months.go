package manual

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type WmsLogisticsRateMonthsRouter struct {
}

// InitWmsLogisticsRateMonthsRouter 初始化 wmsLogisticsRateMonths表 路由信息
func (s *WmsLogisticsRateMonthsRouter) InitWmsLogisticsRateMonthsRouter(Router *gin.RouterGroup) {
	wmsLogisticsRateMonthsRouter := Router.Group("wmsLogisticsRateMonths").Use(middleware.OperationRecord())
	wmsLogisticsRateMonthsRouterWithoutRecord := Router.Group("wmsLogisticsRateMonths")
	var wmsLogisticsRateMonthsApi = v1.ApiGroupApp.ManualApiGroup.WmsLogisticsRateMonthsApi
	{
		wmsLogisticsRateMonthsRouter.POST("createWmsLogisticsRateMonths", wmsLogisticsRateMonthsApi.CreateWmsLogisticsRateMonths)   // 新建wmsLogisticsRateMonths表
		wmsLogisticsRateMonthsRouter.DELETE("deleteWmsLogisticsRateMonths", wmsLogisticsRateMonthsApi.DeleteWmsLogisticsRateMonths) // 删除wmsLogisticsRateMonths表
		wmsLogisticsRateMonthsRouter.DELETE("deleteWmsLogisticsRateMonthsAll", wmsLogisticsRateMonthsApi.DeleteWmsLogisticsRateMonthsAll) // 全部删除wmsLogisticsRateMonths表
		wmsLogisticsRateMonthsRouter.DELETE("deleteWmsLogisticsRateMonthsByIds", wmsLogisticsRateMonthsApi.DeleteWmsLogisticsRateMonthsByIds) // 批量删除wmsLogisticsRateMonths表
		wmsLogisticsRateMonthsRouter.PUT("updateWmsLogisticsRateMonths", wmsLogisticsRateMonthsApi.UpdateWmsLogisticsRateMonths)    // 更新wmsLogisticsRateMonths表
	}
	{
		wmsLogisticsRateMonthsRouterWithoutRecord.GET("findWmsLogisticsRateMonths", wmsLogisticsRateMonthsApi.FindWmsLogisticsRateMonths)        // 根据ID获取wmsLogisticsRateMonths表
		wmsLogisticsRateMonthsRouterWithoutRecord.GET("getWmsLogisticsRateMonthsList", wmsLogisticsRateMonthsApi.GetWmsLogisticsRateMonthsList)  // 获取wmsLogisticsRateMonths表列表
	}
}
