package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/App"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type AppBillSearch struct{
    App.AppBill
    StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
    EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
    request.PageInfo
}
