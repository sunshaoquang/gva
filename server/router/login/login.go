package login

import (
	"github.com/gin-gonic/gin"
)

type loginRouter struct{}

func (e *loginRouter) InitloginRouter(Router *gin.RouterGroup) {
	// loginRouter := Router.Group("wx").Use(middleware.OperationRecord())
	loginRouterWithoutRecord := Router.Group("wx")

	{
		loginRouterWithoutRecord.POST("login", loginApi.WxLogin)                     // 操作微信用户登录
		loginRouterWithoutRecord.GET("checkUserOpenById", baseApi.CheckUserOpenById) // 检查用户OpenId状态信息

	}
}
