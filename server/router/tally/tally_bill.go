package tally

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type TallyBillRouter struct{}

// InitTallyBillRouter 初始化 记账账单表 路由信息
func (s *TallyBillRouter) InitTallyBillRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	tallyBillRouter := Router.Group("tallyBill").Use(middleware.OperationRecord())
	tallyBillRouterWithoutRecord := Router.Group("tallyBill")
	tallyBillRouterWithoutAuth := PublicRouter.Group("tallyBill")

	var tallyBillApi = v1.ApiGroupApp.TallyApiGroup.TallyBillApi
	{
		tallyBillRouter.POST("createTallyBill", tallyBillApi.CreateTallyBill)             // 新建记账账单表
		tallyBillRouter.DELETE("deleteTallyBill", tallyBillApi.DeleteTallyBill)           // 删除记账账单表
		tallyBillRouter.DELETE("deleteTallyBillByIds", tallyBillApi.DeleteTallyBillByIds) // 批量删除记账账单表
		tallyBillRouter.PUT("updateTallyBill", tallyBillApi.UpdateTallyBill)              // 更新记账账单表
	}
	{
		tallyBillRouterWithoutRecord.GET("findTallyBill", tallyBillApi.FindTallyBill)       // 根据ID获取记账账单表
		tallyBillRouterWithoutRecord.GET("getTallyBillList", tallyBillApi.GetTallyBillList) // 获取记账账单表列表
	}
	{
		tallyBillRouterWithoutAuth.GET("getTallyBillDetailDataList", tallyBillApi.GetTallyBillDetailDataList) // 获取记账账单表详细列表接口
		tallyBillRouterWithoutAuth.GET("getTallyBillPublic", tallyBillApi.GetTallyBillPublic)                 // 获取记账账单表列表
	}
}
