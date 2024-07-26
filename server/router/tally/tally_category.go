package tally

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type TallyCategoryRouter struct{}

// InitTallyCategoryRouter 初始化 记账类别表 路由信息
func (s *TallyCategoryRouter) InitTallyCategoryRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	tallyCategoryRouter := Router.Group("tallyCategory").Use(middleware.OperationRecord())
	tallyCategoryRouterWithoutRecord := Router.Group("tallyCategory")
	tallyCategoryRouterWithoutAuth := PublicRouter.Group("tallyCategory")

	var tallyCategoryApi = v1.ApiGroupApp.TallyApiGroup.TallyCategoryApi
	{
		tallyCategoryRouter.POST("createTallyCategory", tallyCategoryApi.CreateTallyCategory)             // 新建记账类别表
		tallyCategoryRouter.DELETE("deleteTallyCategory", tallyCategoryApi.DeleteTallyCategory)           // 删除记账类别表
		tallyCategoryRouter.DELETE("deleteTallyCategoryByIds", tallyCategoryApi.DeleteTallyCategoryByIds) // 批量删除记账类别表
		tallyCategoryRouter.PUT("updateTallyCategory", tallyCategoryApi.UpdateTallyCategory)              // 更新记账类别表
	}
	{
		tallyCategoryRouterWithoutRecord.GET("findTallyCategory", tallyCategoryApi.FindTallyCategory)       // 根据ID获取记账类别表
		tallyCategoryRouterWithoutRecord.GET("getTallyCategoryList", tallyCategoryApi.GetTallyCategoryList) // 获取记账类别表列表
	}
	{
		tallyCategoryRouterWithoutAuth.GET("getTallyCategoryPublic", tallyCategoryApi.GetTallyCategoryPublic) // 获取记账类别表列表
	}
}
