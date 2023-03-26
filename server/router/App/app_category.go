package App

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type AppCategoryRouter struct {
}

// InitAppCategoryRouter 初始化 AppCategory 路由信息
func (s *AppCategoryRouter) InitAppCategoryRouter(Router *gin.RouterGroup) {
	appCategoryRouter := Router.Group("appCategory").Use(middleware.OperationRecord())
	appCategoryRouterWithoutRecord := Router.Group("appCategory")
	var appCategoryApi = v1.ApiGroupApp.AppApiGroup.AppCategoryApi
	{
		appCategoryRouter.POST("createAppCategory", appCategoryApi.CreateAppCategory)             // 新建AppCategory
		appCategoryRouter.DELETE("deleteAppCategory", appCategoryApi.DeleteAppCategory)           // 删除AppCategory
		appCategoryRouter.DELETE("deleteAppCategoryByIds", appCategoryApi.DeleteAppCategoryByIds) // 批量删除AppCategory
		appCategoryRouter.PUT("updateAppCategory", appCategoryApi.UpdateAppCategory)              // 更新AppCategory
	}
	{
		appCategoryRouterWithoutRecord.GET("findAppCategory", appCategoryApi.FindAppCategory)       // 根据ID获取AppCategory
		appCategoryRouterWithoutRecord.GET("getAppCategoryList", appCategoryApi.GetAppCategoryList) // 获取AppCategory列表
	}
}
