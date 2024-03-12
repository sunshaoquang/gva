package manual

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type WmsLogisticsTailfeeMiRouter struct {
}

// InitWmsLogisticsTailfeeMiRouter 初始化 尾程费用结构 路由信息
func (s *WmsLogisticsTailfeeMiRouter) InitWmsLogisticsTailfeeMiRouter(Router *gin.RouterGroup) {
	wmsLogisticsTailfeeMiRouter := Router.Group("wmsLogisticsTailfeeMi").Use(middleware.OperationRecord())
	wmsLogisticsTailfeeMiRouterWithoutRecord := Router.Group("wmsLogisticsTailfeeMi")
	var wmsLogisticsTailfeeMiApi = v1.ApiGroupApp.ManualApiGroup.WmsLogisticsTailfeeMiApi
	{
		wmsLogisticsTailfeeMiRouter.POST("createWmsLogisticsTailfeeMi", wmsLogisticsTailfeeMiApi.CreateWmsLogisticsTailfeeMi)   // 新建尾程费用结构
		wmsLogisticsTailfeeMiRouter.DELETE("deleteWmsLogisticsTailfeeMi", wmsLogisticsTailfeeMiApi.DeleteWmsLogisticsTailfeeMi) // 删除尾程费用结构
		wmsLogisticsTailfeeMiRouter.DELETE("deleteWmsLogisticsTailfeeMiAll", wmsLogisticsTailfeeMiApi.DeleteWmsLogisticsTailfeeMiAll) // 全部删除尾程费用结构
		wmsLogisticsTailfeeMiRouter.DELETE("deleteWmsLogisticsTailfeeMiByIds", wmsLogisticsTailfeeMiApi.DeleteWmsLogisticsTailfeeMiByIds) // 批量删除尾程费用结构
		wmsLogisticsTailfeeMiRouter.PUT("updateWmsLogisticsTailfeeMi", wmsLogisticsTailfeeMiApi.UpdateWmsLogisticsTailfeeMi)    // 更新尾程费用结构
	}
	{
		wmsLogisticsTailfeeMiRouterWithoutRecord.GET("findWmsLogisticsTailfeeMi", wmsLogisticsTailfeeMiApi.FindWmsLogisticsTailfeeMi)        // 根据ID获取尾程费用结构
		wmsLogisticsTailfeeMiRouterWithoutRecord.GET("getWmsLogisticsTailfeeMiList", wmsLogisticsTailfeeMiApi.GetWmsLogisticsTailfeeMiList)  // 获取尾程费用结构列表
	}
}
