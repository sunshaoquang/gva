package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/router/example"
	"github.com/flipped-aurora/gin-vue-admin/server/router/login"
	"github.com/flipped-aurora/gin-vue-admin/server/router/system"
	"github.com/flipped-aurora/gin-vue-admin/server/router/tally"
)

type RouterGroup struct {
	System  system.RouterGroup
	Example example.RouterGroup
	Login   login.RouterGroup
	Tally   tally.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
