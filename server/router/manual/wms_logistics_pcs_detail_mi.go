package manual

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type WmsLogisticsPcsDetailMiRouter struct {
}

// InitWmsLogisticsPcsDetailMiRouter 初始化 主要产品成本明细表 路由信息
func (s *WmsLogisticsPcsDetailMiRouter) InitWmsLogisticsPcsDetailMiRouter(Router *gin.RouterGroup) {
	wmsLogisticsPcsDetailMiRouter := Router.Group("wmsLogisticsPcsDetailMi").Use(middleware.OperationRecord())
	wmsLogisticsPcsDetailMiRouterWithoutRecord := Router.Group("wmsLogisticsPcsDetailMi")
	var wmsLogisticsPcsDetailMiApi = v1.ApiGroupApp.ManualApiGroup.WmsLogisticsPcsDetailMiApi
	{
		wmsLogisticsPcsDetailMiRouter.POST("createWmsLogisticsPcsDetailMi", wmsLogisticsPcsDetailMiApi.CreateWmsLogisticsPcsDetailMi)   // 新建主要产品成本明细表
		wmsLogisticsPcsDetailMiRouter.DELETE("deleteWmsLogisticsPcsDetailMi", wmsLogisticsPcsDetailMiApi.DeleteWmsLogisticsPcsDetailMi) // 删除主要产品成本明细表
		wmsLogisticsPcsDetailMiRouter.DELETE("deleteWmsLogisticsPcsDetailMiByIds", wmsLogisticsPcsDetailMiApi.DeleteWmsLogisticsPcsDetailMiByIds) // 批量删除主要产品成本明细表
		wmsLogisticsPcsDetailMiRouter.PUT("updateWmsLogisticsPcsDetailMi", wmsLogisticsPcsDetailMiApi.UpdateWmsLogisticsPcsDetailMi)    // 更新主要产品成本明细表
	}
	{
		wmsLogisticsPcsDetailMiRouterWithoutRecord.GET("findWmsLogisticsPcsDetailMi", wmsLogisticsPcsDetailMiApi.FindWmsLogisticsPcsDetailMi)        // 根据ID获取主要产品成本明细表
		wmsLogisticsPcsDetailMiRouterWithoutRecord.GET("getWmsLogisticsPcsDetailMiList", wmsLogisticsPcsDetailMiApi.GetWmsLogisticsPcsDetailMiList)  // 获取主要产品成本明细表列表
	}
}
