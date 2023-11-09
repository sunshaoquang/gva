package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/manual"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type PromotionMarketTargetSearch struct{
    manual.PromotionMarketTarget
    StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
    EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
    StartTarget  *int  `json:"startTarget" form:"startTarget"`
    EndTarget  *int  `json:"endTarget" form:"endTarget"`
    request.PageInfo
    Sort  string `json:"sort" form:"sort"`
    Order string `json:"order" form:"order"`
}
