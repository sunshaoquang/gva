package manual

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type WmsLogisticsPcs2023MiRouter struct {
}

// InitWmsLogisticsPcs2023MiRouter 初始化 2023年物流成本明细表 路由信息
func (s *WmsLogisticsPcs2023MiRouter) InitWmsLogisticsPcs2023MiRouter(Router *gin.RouterGroup) {
	wmsLogisticsPcs2023MiRouter := Router.Group("wmsLogisticsPcs2023Mi").Use(middleware.OperationRecord())
	wmsLogisticsPcs2023MiRouterWithoutRecord := Router.Group("wmsLogisticsPcs2023Mi")
	var wmsLogisticsPcs2023MiApi = v1.ApiGroupApp.ManualApiGroup.WmsLogisticsPcs2023MiApi
	{
		wmsLogisticsPcs2023MiRouter.POST("createWmsLogisticsPcs2023Mi", wmsLogisticsPcs2023MiApi.CreateWmsLogisticsPcs2023Mi)   // 新建2023年物流成本明细表
		wmsLogisticsPcs2023MiRouter.DELETE("deleteWmsLogisticsPcs2023Mi", wmsLogisticsPcs2023MiApi.DeleteWmsLogisticsPcs2023Mi) // 删除2023年物流成本明细表
		wmsLogisticsPcs2023MiRouter.DELETE("deleteWmsLogisticsPcs2023MiByIds", wmsLogisticsPcs2023MiApi.DeleteWmsLogisticsPcs2023MiByIds) // 批量删除2023年物流成本明细表
		wmsLogisticsPcs2023MiRouter.PUT("updateWmsLogisticsPcs2023Mi", wmsLogisticsPcs2023MiApi.UpdateWmsLogisticsPcs2023Mi)    // 更新2023年物流成本明细表
	}
	{
		wmsLogisticsPcs2023MiRouterWithoutRecord.GET("findWmsLogisticsPcs2023Mi", wmsLogisticsPcs2023MiApi.FindWmsLogisticsPcs2023Mi)        // 根据ID获取2023年物流成本明细表
		wmsLogisticsPcs2023MiRouterWithoutRecord.GET("getWmsLogisticsPcs2023MiList", wmsLogisticsPcs2023MiApi.GetWmsLogisticsPcs2023MiList)  // 获取2023年物流成本明细表列表
	}
}
