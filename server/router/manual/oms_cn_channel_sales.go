package manual

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type OmsCnChannelSalesRouter struct {
}

// InitOmsCnChannelSalesRouter 初始化 按天销售目标录入表 路由信息
func (s *OmsCnChannelSalesRouter) InitOmsCnChannelSalesRouter(Router *gin.RouterGroup) {
	omsCnChannelSalesRouter := Router.Group("omsCnChannelSales").Use(middleware.OperationRecord())
	omsCnChannelSalesRouterWithoutRecord := Router.Group("omsCnChannelSales")
	var omsCnChannelSalesApi = v1.ApiGroupApp.ManualApiGroup.OmsCnChannelSalesApi
	{
		omsCnChannelSalesRouter.POST("createOmsCnChannelSales", omsCnChannelSalesApi.CreateOmsCnChannelSales)   // 新建按天销售目标录入表
		omsCnChannelSalesRouter.DELETE("deleteOmsCnChannelSales", omsCnChannelSalesApi.DeleteOmsCnChannelSales) // 删除按天销售目标录入表
		omsCnChannelSalesRouter.DELETE("deleteOmsCnChannelSalesAll", omsCnChannelSalesApi.DeleteOmsCnChannelSalesAll) // 全部删除按天销售目标录入表
		omsCnChannelSalesRouter.DELETE("deleteOmsCnChannelSalesByIds", omsCnChannelSalesApi.DeleteOmsCnChannelSalesByIds) // 批量删除按天销售目标录入表
		omsCnChannelSalesRouter.PUT("updateOmsCnChannelSales", omsCnChannelSalesApi.UpdateOmsCnChannelSales)    // 更新按天销售目标录入表
	}
	{
		omsCnChannelSalesRouterWithoutRecord.GET("findOmsCnChannelSales", omsCnChannelSalesApi.FindOmsCnChannelSales)        // 根据ID获取按天销售目标录入表
		omsCnChannelSalesRouterWithoutRecord.GET("getOmsCnChannelSalesList", omsCnChannelSalesApi.GetOmsCnChannelSalesList)  // 获取按天销售目标录入表列表
	}
}
