package manual

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type WmsLogisticsStoragefeeMiRouter struct {
}

// InitWmsLogisticsStoragefeeMiRouter 初始化 仓储费 路由信息
func (s *WmsLogisticsStoragefeeMiRouter) InitWmsLogisticsStoragefeeMiRouter(Router *gin.RouterGroup) {
	wmsLogisticsStoragefeeMiRouter := Router.Group("wmsLogisticsStoragefeeMi").Use(middleware.OperationRecord())
	wmsLogisticsStoragefeeMiRouterWithoutRecord := Router.Group("wmsLogisticsStoragefeeMi")
	var wmsLogisticsStoragefeeMiApi = v1.ApiGroupApp.ManualApiGroup.WmsLogisticsStoragefeeMiApi
	{
		wmsLogisticsStoragefeeMiRouter.POST("createWmsLogisticsStoragefeeMi", wmsLogisticsStoragefeeMiApi.CreateWmsLogisticsStoragefeeMi)   // 新建仓储费
		wmsLogisticsStoragefeeMiRouter.DELETE("deleteWmsLogisticsStoragefeeMi", wmsLogisticsStoragefeeMiApi.DeleteWmsLogisticsStoragefeeMi) // 删除仓储费
		wmsLogisticsStoragefeeMiRouter.DELETE("deleteWmsLogisticsStoragefeeMiAll", wmsLogisticsStoragefeeMiApi.DeleteWmsLogisticsStoragefeeMiAll) // 全部删除仓储费
		wmsLogisticsStoragefeeMiRouter.DELETE("deleteWmsLogisticsStoragefeeMiByIds", wmsLogisticsStoragefeeMiApi.DeleteWmsLogisticsStoragefeeMiByIds) // 批量删除仓储费
		wmsLogisticsStoragefeeMiRouter.PUT("updateWmsLogisticsStoragefeeMi", wmsLogisticsStoragefeeMiApi.UpdateWmsLogisticsStoragefeeMi)    // 更新仓储费
	}
	{
		wmsLogisticsStoragefeeMiRouterWithoutRecord.GET("findWmsLogisticsStoragefeeMi", wmsLogisticsStoragefeeMiApi.FindWmsLogisticsStoragefeeMi)        // 根据ID获取仓储费
		wmsLogisticsStoragefeeMiRouterWithoutRecord.GET("getWmsLogisticsStoragefeeMiList", wmsLogisticsStoragefeeMiApi.GetWmsLogisticsStoragefeeMiList)  // 获取仓储费列表
	}
}
