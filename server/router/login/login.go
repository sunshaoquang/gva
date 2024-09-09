package login

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/gin-gonic/gin"
)

type loginRouter struct{}

func (e *loginRouter) InitloginRouter(Router *gin.RouterGroup) {
	// loginRouter := Router.Group("wx").Use(middleware.OperationRecord())
	loginRouterWithoutRecord := Router.Group("wx")
	loginApi := v1.ApiGroupApp.LoginApiGroup.LoginApi

	{
		loginRouterWithoutRecord.POST("login", loginApi.WxLogin) // 操作微信用户登录
	}
}
