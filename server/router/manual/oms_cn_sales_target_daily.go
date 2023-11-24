package manual

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type OmsCnSalesTargetDailyRouter struct {
}

// InitOmsCnSalesTargetDailyRouter 初始化 京东自营销售录入表 路由信息
func (s *OmsCnSalesTargetDailyRouter) InitOmsCnSalesTargetDailyRouter(Router *gin.RouterGroup) {
	omsCnSalesTargetDailyRouter := Router.Group("omsCnSalesTargetDaily").Use(middleware.OperationRecord())
	omsCnSalesTargetDailyRouterWithoutRecord := Router.Group("omsCnSalesTargetDaily")
	var omsCnSalesTargetDailyApi = v1.ApiGroupApp.ManualApiGroup.OmsCnSalesTargetDailyApi
	{
		omsCnSalesTargetDailyRouter.POST("createOmsCnSalesTargetDaily", omsCnSalesTargetDailyApi.CreateOmsCnSalesTargetDaily)   // 新建京东自营销售录入表
		omsCnSalesTargetDailyRouter.DELETE("deleteOmsCnSalesTargetDaily", omsCnSalesTargetDailyApi.DeleteOmsCnSalesTargetDaily) // 删除京东自营销售录入表
		omsCnSalesTargetDailyRouter.DELETE("deleteOmsCnSalesTargetDailyAll", omsCnSalesTargetDailyApi.DeleteOmsCnSalesTargetDailyAll) // 全部删除京东自营销售录入表
		omsCnSalesTargetDailyRouter.DELETE("deleteOmsCnSalesTargetDailyByIds", omsCnSalesTargetDailyApi.DeleteOmsCnSalesTargetDailyByIds) // 批量删除京东自营销售录入表
		omsCnSalesTargetDailyRouter.PUT("updateOmsCnSalesTargetDaily", omsCnSalesTargetDailyApi.UpdateOmsCnSalesTargetDaily)    // 更新京东自营销售录入表
	}
	{
		omsCnSalesTargetDailyRouterWithoutRecord.GET("findOmsCnSalesTargetDaily", omsCnSalesTargetDailyApi.FindOmsCnSalesTargetDaily)        // 根据ID获取京东自营销售录入表
		omsCnSalesTargetDailyRouterWithoutRecord.GET("getOmsCnSalesTargetDailyList", omsCnSalesTargetDailyApi.GetOmsCnSalesTargetDailyList)  // 获取京东自营销售录入表列表
	}
}
