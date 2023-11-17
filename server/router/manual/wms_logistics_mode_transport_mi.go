package manual

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type WmsLogisticsModeTransportMiRouter struct {
}

// InitWmsLogisticsModeTransportMiRouter 初始化 费用by运输方式 路由信息
func (s *WmsLogisticsModeTransportMiRouter) InitWmsLogisticsModeTransportMiRouter(Router *gin.RouterGroup) {
	wmsLogisticsModeTransportMiRouter := Router.Group("wmsLogisticsModeTransportMi").Use(middleware.OperationRecord())
	wmsLogisticsModeTransportMiRouterWithoutRecord := Router.Group("wmsLogisticsModeTransportMi")
	var wmsLogisticsModeTransportMiApi = v1.ApiGroupApp.ManualApiGroup.WmsLogisticsModeTransportMiApi
	{
		wmsLogisticsModeTransportMiRouter.POST("createWmsLogisticsModeTransportMi", wmsLogisticsModeTransportMiApi.CreateWmsLogisticsModeTransportMi)   // 新建费用by运输方式
		wmsLogisticsModeTransportMiRouter.DELETE("deleteWmsLogisticsModeTransportMi", wmsLogisticsModeTransportMiApi.DeleteWmsLogisticsModeTransportMi) // 删除费用by运输方式
		wmsLogisticsModeTransportMiRouter.DELETE("deleteWmsLogisticsModeTransportMiAll", wmsLogisticsModeTransportMiApi.DeleteWmsLogisticsModeTransportMiAll) // 全部删除费用by运输方式
		wmsLogisticsModeTransportMiRouter.DELETE("deleteWmsLogisticsModeTransportMiByIds", wmsLogisticsModeTransportMiApi.DeleteWmsLogisticsModeTransportMiByIds) // 批量删除费用by运输方式
		wmsLogisticsModeTransportMiRouter.PUT("updateWmsLogisticsModeTransportMi", wmsLogisticsModeTransportMiApi.UpdateWmsLogisticsModeTransportMi)    // 更新费用by运输方式
	}
	{
		wmsLogisticsModeTransportMiRouterWithoutRecord.GET("findWmsLogisticsModeTransportMi", wmsLogisticsModeTransportMiApi.FindWmsLogisticsModeTransportMi)        // 根据ID获取费用by运输方式
		wmsLogisticsModeTransportMiRouterWithoutRecord.GET("getWmsLogisticsModeTransportMiList", wmsLogisticsModeTransportMiApi.GetWmsLogisticsModeTransportMiList)  // 获取费用by运输方式列表
	}
}
