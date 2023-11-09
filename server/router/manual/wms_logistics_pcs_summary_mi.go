package manual

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type WmsLogisticsPcsSummaryMiRouter struct {
}

// InitWmsLogisticsPcsSummaryMiRouter 初始化 主要产品成本降本汇总表 路由信息
func (s *WmsLogisticsPcsSummaryMiRouter) InitWmsLogisticsPcsSummaryMiRouter(Router *gin.RouterGroup) {
	wmsLogisticsPcsSummaryMiRouter := Router.Group("wmsLogisticsPcsSummaryMi").Use(middleware.OperationRecord())
	wmsLogisticsPcsSummaryMiRouterWithoutRecord := Router.Group("wmsLogisticsPcsSummaryMi")
	var wmsLogisticsPcsSummaryMiApi = v1.ApiGroupApp.ManualApiGroup.WmsLogisticsPcsSummaryMiApi
	{
		wmsLogisticsPcsSummaryMiRouter.POST("createWmsLogisticsPcsSummaryMi", wmsLogisticsPcsSummaryMiApi.CreateWmsLogisticsPcsSummaryMi)   // 新建主要产品成本降本汇总表
		wmsLogisticsPcsSummaryMiRouter.DELETE("deleteWmsLogisticsPcsSummaryMi", wmsLogisticsPcsSummaryMiApi.DeleteWmsLogisticsPcsSummaryMi) // 删除主要产品成本降本汇总表
		wmsLogisticsPcsSummaryMiRouter.DELETE("deleteWmsLogisticsPcsSummaryMiByIds", wmsLogisticsPcsSummaryMiApi.DeleteWmsLogisticsPcsSummaryMiByIds) // 批量删除主要产品成本降本汇总表
		wmsLogisticsPcsSummaryMiRouter.PUT("updateWmsLogisticsPcsSummaryMi", wmsLogisticsPcsSummaryMiApi.UpdateWmsLogisticsPcsSummaryMi)    // 更新主要产品成本降本汇总表
	}
	{
		wmsLogisticsPcsSummaryMiRouterWithoutRecord.GET("findWmsLogisticsPcsSummaryMi", wmsLogisticsPcsSummaryMiApi.FindWmsLogisticsPcsSummaryMi)        // 根据ID获取主要产品成本降本汇总表
		wmsLogisticsPcsSummaryMiRouterWithoutRecord.GET("getWmsLogisticsPcsSummaryMiList", wmsLogisticsPcsSummaryMiApi.GetWmsLogisticsPcsSummaryMiList)  // 获取主要产品成本降本汇总表列表
	}
}
