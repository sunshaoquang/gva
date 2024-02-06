package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/wxpay/api"
	"github.com/gin-gonic/gin"
)

type WxpayRouter struct {
}

func (s *WxpayRouter) InitWxpayRouter(Router *gin.RouterGroup) {
	plugRouter := Router
	plugApi := api.ApiGroupApp.WxpayApi
	{
		plugRouter.POST("getPayCode", plugApi.GetPayCode)
		plugRouter.GET("getOrderById", plugApi.GetOrderById)
		plugRouter.POST("payAction", plugApi.PayAction)
	}
}
