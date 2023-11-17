package manual

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type WmsLogisticsProviderMiRouter struct {
}

// InitWmsLogisticsProviderMiRouter 初始化 费用by供应商 路由信息
func (s *WmsLogisticsProviderMiRouter) InitWmsLogisticsProviderMiRouter(Router *gin.RouterGroup) {
	wmsLogisticsProviderMiRouter := Router.Group("wmsLogisticsProviderMi").Use(middleware.OperationRecord())
	wmsLogisticsProviderMiRouterWithoutRecord := Router.Group("wmsLogisticsProviderMi")
	var wmsLogisticsProviderMiApi = v1.ApiGroupApp.ManualApiGroup.WmsLogisticsProviderMiApi
	{
		wmsLogisticsProviderMiRouter.POST("createWmsLogisticsProviderMi", wmsLogisticsProviderMiApi.CreateWmsLogisticsProviderMi)   // 新建费用by供应商
		wmsLogisticsProviderMiRouter.DELETE("deleteWmsLogisticsProviderMi", wmsLogisticsProviderMiApi.DeleteWmsLogisticsProviderMi) // 删除费用by供应商
		wmsLogisticsProviderMiRouter.DELETE("deleteWmsLogisticsProviderMiAll", wmsLogisticsProviderMiApi.DeleteWmsLogisticsProviderMiAll) // 全部删除费用by供应商
		wmsLogisticsProviderMiRouter.DELETE("deleteWmsLogisticsProviderMiByIds", wmsLogisticsProviderMiApi.DeleteWmsLogisticsProviderMiByIds) // 批量删除费用by供应商
		wmsLogisticsProviderMiRouter.PUT("updateWmsLogisticsProviderMi", wmsLogisticsProviderMiApi.UpdateWmsLogisticsProviderMi)    // 更新费用by供应商
	}
	{
		wmsLogisticsProviderMiRouterWithoutRecord.GET("findWmsLogisticsProviderMi", wmsLogisticsProviderMiApi.FindWmsLogisticsProviderMi)        // 根据ID获取费用by供应商
		wmsLogisticsProviderMiRouterWithoutRecord.GET("getWmsLogisticsProviderMiList", wmsLogisticsProviderMiApi.GetWmsLogisticsProviderMiList)  // 获取费用by供应商列表
	}
}
