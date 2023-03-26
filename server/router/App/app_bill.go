package App

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type AppBillRouter struct {
}

// InitAppBillRouter 初始化 AppBill 路由信息
func (s *AppBillRouter) InitAppBillRouter(Router *gin.RouterGroup) {
	appBillRouter := Router.Group("appBill").Use(middleware.OperationRecord())
	appBillRouterWithoutRecord := Router.Group("appBill")
	var appBillApi = v1.ApiGroupApp.AppApiGroup.AppBillApi
	{
		appBillRouter.POST("createAppBill", appBillApi.CreateAppBill)   // 新建AppBill
		appBillRouter.DELETE("deleteAppBill", appBillApi.DeleteAppBill) // 删除AppBill
		appBillRouter.DELETE("deleteAppBillByIds", appBillApi.DeleteAppBillByIds) // 批量删除AppBill
		appBillRouter.PUT("updateAppBill", appBillApi.UpdateAppBill)    // 更新AppBill
	}
	{
		appBillRouterWithoutRecord.GET("findAppBill", appBillApi.FindAppBill)        // 根据ID获取AppBill
		appBillRouterWithoutRecord.GET("getAppBillList", appBillApi.GetAppBillList)  // 获取AppBill列表
	}
}
