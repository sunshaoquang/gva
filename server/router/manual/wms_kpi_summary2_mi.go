package manual

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type WmsKpiSummary2MiRouter struct {
}

// InitWmsKpiSummary2MiRouter 初始化 wmsKpiSummary2Mi表 路由信息
func (s *WmsKpiSummary2MiRouter) InitWmsKpiSummary2MiRouter(Router *gin.RouterGroup) {
	wmsKpiSummary2MiRouter := Router.Group("wmsKpiSummary2Mi").Use(middleware.OperationRecord())
	wmsKpiSummary2MiRouterWithoutRecord := Router.Group("wmsKpiSummary2Mi")
	var wmsKpiSummary2MiApi = v1.ApiGroupApp.ManualApiGroup.WmsKpiSummary2MiApi
	{
		wmsKpiSummary2MiRouter.POST("createWmsKpiSummary2Mi", wmsKpiSummary2MiApi.CreateWmsKpiSummary2Mi)   // 新建wmsKpiSummary2Mi表
		wmsKpiSummary2MiRouter.DELETE("deleteWmsKpiSummary2Mi", wmsKpiSummary2MiApi.DeleteWmsKpiSummary2Mi) // 删除wmsKpiSummary2Mi表
		wmsKpiSummary2MiRouter.DELETE("deleteWmsKpiSummary2MiAll", wmsKpiSummary2MiApi.DeleteWmsKpiSummary2MiAll) // 全部删除wmsKpiSummary2MiAll表
		wmsKpiSummary2MiRouter.DELETE("deleteWmsKpiSummary2MiByIds", wmsKpiSummary2MiApi.DeleteWmsKpiSummary2MiByIds) // 批量删除wmsKpiSummary2Mi表
		wmsKpiSummary2MiRouter.PUT("updateWmsKpiSummary2Mi", wmsKpiSummary2MiApi.UpdateWmsKpiSummary2Mi)    // 更新wmsKpiSummary2Mi表
	}
	{
		wmsKpiSummary2MiRouterWithoutRecord.GET("findWmsKpiSummary2Mi", wmsKpiSummary2MiApi.FindWmsKpiSummary2Mi)        // 根据ID获取wmsKpiSummary2Mi表
		wmsKpiSummary2MiRouterWithoutRecord.GET("getWmsKpiSummary2MiList", wmsKpiSummary2MiApi.GetWmsKpiSummary2MiList)  // 获取wmsKpiSummary2Mi表列表
	}
}
