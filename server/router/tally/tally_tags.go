package tally

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type TallyTagsRouter struct{}

// InitTallyTagsRouter 初始化 标签 路由信息
func (s *TallyTagsRouter) InitTallyTagsRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	tallyTagsRouter := Router.Group("tallyTags").Use(middleware.OperationRecord())
	tallyTagsRouterWithoutRecord := Router.Group("tallyTags")
	tallyTagsRouterWithoutAuth := PublicRouter.Group("tallyTags")

	var tallyTagsApi = v1.ApiGroupApp.TallyApiGroup.TallyTagsApi
	{
		tallyTagsRouter.POST("createTallyTags", tallyTagsApi.CreateTallyTags)             // 新建标签
		tallyTagsRouter.DELETE("deleteTallyTags", tallyTagsApi.DeleteTallyTags)           // 删除标签
		tallyTagsRouter.DELETE("deleteTallyTagsByIds", tallyTagsApi.DeleteTallyTagsByIds) // 批量删除标签
		tallyTagsRouter.PUT("updateTallyTags", tallyTagsApi.UpdateTallyTags)              // 更新标签
	}
	{
		tallyTagsRouterWithoutRecord.GET("findTallyTags", tallyTagsApi.FindTallyTags)                   // 根据ID获取标签
		tallyTagsRouterWithoutRecord.GET("getTallyTagsList", tallyTagsApi.GetTallyTagsList)             // 获取标签列表
		tallyTagsRouterWithoutRecord.GET("getTallyTagsListByUser", tallyTagsApi.GetTallyTagsListByUser) // 根据用户ID获取标签列表

	}
	{
		tallyTagsRouterWithoutAuth.GET("getTallyTagsPublic", tallyTagsApi.GetTallyTagsPublic) // 获取标签列表
	}
}
