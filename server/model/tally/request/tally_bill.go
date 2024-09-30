package request

import (
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type TallyBillSearch struct {
	request.PageInfo
}

type TallyBillDetailSearch struct {
	QueryTime time.Time `json:"queryTime" form:"queryTime"`
	TimeType  string    `json:"timeType" form:"timeType"`
}
