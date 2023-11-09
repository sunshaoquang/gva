// 自动生成模板PromotionMarketTarget
package manual

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
	
)

// 大促推广目标表 结构体  PromotionMarketTarget
type PromotionMarketTarget struct {
      global.GVA_MODEL
      ShopName  string `json:"shopName" form:"shopName" gorm:"column:shop_name;comment:店铺名称;size:64;"`  //店铺名称 
      ShopCode  string `json:"shopCode" form:"shopCode" gorm:"column:shop_code;comment:店铺代码;size:64;"`  //店铺代码 
      Target  uint `json:"target" form:"target" gorm:"column:target;comment:任务目标;"`  //目标 
      Type  string `json:"type" form:"type" gorm:"column:type;comment:场景类型;"`  //场景类型 
      CreatedName  string `json:"createdName" form:"createdName" gorm:"column:created_name;comment:创建人;size:32;"`  //创建人 
      SheetName  string `json:"sheetName" form:"sheetName" gorm:"column:sheet_name;comment:手工表名称;size:64;"`  //手工表名称 
      CreatedBy  uint   `gorm:"column:created_by;comment:创建者"`
      UpdatedBy  uint   `gorm:"column:updated_by;comment:更新者"`
      DeletedBy  uint   `gorm:"column:deleted_by;comment:删除者"`
}


// TableName 大促推广目标表 PromotionMarketTarget自定义表名 dim_promotion_market_target
func (PromotionMarketTarget) TableName() string {
  return "dim_promotion_market_target"
}

