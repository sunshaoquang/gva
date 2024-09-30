package login

import api "github.com/flipped-aurora/gin-vue-admin/server/api/v1"

type RouterGroup struct {
	loginRouter
}

var (
	loginApi = api.ApiGroupApp.LoginApiGroup.LoginApi
	baseApi  = api.ApiGroupApp.SystemApiGroup.BaseApi
)
