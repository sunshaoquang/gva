package manual

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type PromotionMarketTargetRouter struct {
}

// InitPromotionMarketTargetRouter 初始化 大促推广目标表 路由信息
func (s *PromotionMarketTargetRouter) InitPromotionMarketTargetRouter(Router *gin.RouterGroup) {
	MarketTargetRouter := Router.Group("MarketTarget").Use(middleware.OperationRecord())
	MarketTargetRouterWithoutRecord := Router.Group("MarketTarget")
	var MarketTargetApi = v1.ApiGroupApp.ManualApiGroup.PromotionMarketTargetApi
	{
		MarketTargetRouter.POST("createPromotionMarketTarget", MarketTargetApi.CreatePromotionMarketTarget)   // 新建大促推广目标表
		MarketTargetRouter.DELETE("deletePromotionMarketTarget", MarketTargetApi.DeletePromotionMarketTarget) // 删除大促推广目标表
		MarketTargetRouter.DELETE("deletePromotionMarketTargetByIds", MarketTargetApi.DeletePromotionMarketTargetByIds) // 批量删除大促推广目标表
		MarketTargetRouter.PUT("updatePromotionMarketTarget", MarketTargetApi.UpdatePromotionMarketTarget)    // 更新大促推广目标表
	}
	{
		MarketTargetRouterWithoutRecord.GET("findPromotionMarketTarget", MarketTargetApi.FindPromotionMarketTarget)        // 根据ID获取大促推广目标表
		MarketTargetRouterWithoutRecord.GET("getPromotionMarketTargetList", MarketTargetApi.GetPromotionMarketTargetList)  // 获取大促推广目标表列表
	}
}
