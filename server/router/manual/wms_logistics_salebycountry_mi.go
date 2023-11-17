package manual

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type WmsLogisticsSalebycountryMiRouter struct {
}

// InitWmsLogisticsSalebycountryMiRouter 初始化 费用by供应商 路由信息
func (s *WmsLogisticsSalebycountryMiRouter) InitWmsLogisticsSalebycountryMiRouter(Router *gin.RouterGroup) {
	wmsLogisticsSalebycountryMiRouter := Router.Group("wmsLogisticsSalebycountryMi").Use(middleware.OperationRecord())
	wmsLogisticsSalebycountryMiRouterWithoutRecord := Router.Group("wmsLogisticsSalebycountryMi")
	var wmsLogisticsSalebycountryMiApi = v1.ApiGroupApp.ManualApiGroup.WmsLogisticsSalebycountryMiApi
	{
		wmsLogisticsSalebycountryMiRouter.POST("createWmsLogisticsSalebycountryMi", wmsLogisticsSalebycountryMiApi.CreateWmsLogisticsSalebycountryMi)   // 新建费用by供应商
		wmsLogisticsSalebycountryMiRouter.DELETE("deleteWmsLogisticsSalebycountryMi", wmsLogisticsSalebycountryMiApi.DeleteWmsLogisticsSalebycountryMi) // 删除费用by供应商
		wmsLogisticsSalebycountryMiRouter.DELETE("deleteWmsLogisticsSalebycountryMiAll", wmsLogisticsSalebycountryMiApi.DeleteWmsLogisticsSalebycountryMiAll) // 全部删除费用by供应商
		wmsLogisticsSalebycountryMiRouter.DELETE("deleteWmsLogisticsSalebycountryMiByIds", wmsLogisticsSalebycountryMiApi.DeleteWmsLogisticsSalebycountryMiByIds) // 批量删除费用by供应商
		wmsLogisticsSalebycountryMiRouter.PUT("updateWmsLogisticsSalebycountryMi", wmsLogisticsSalebycountryMiApi.UpdateWmsLogisticsSalebycountryMi)    // 更新费用by供应商
	}
	{
		wmsLogisticsSalebycountryMiRouterWithoutRecord.GET("findWmsLogisticsSalebycountryMi", wmsLogisticsSalebycountryMiApi.FindWmsLogisticsSalebycountryMi)        // 根据ID获取费用by供应商
		wmsLogisticsSalebycountryMiRouterWithoutRecord.GET("getWmsLogisticsSalebycountryMiList", wmsLogisticsSalebycountryMiApi.GetWmsLogisticsSalebycountryMiList)  // 获取费用by供应商列表
	}
}
