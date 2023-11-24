// 自动生成模板OmsCnSalesTargetDaily
package manual

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
	
)

// 京东自营销售录入表 结构体  OmsCnSalesTargetDaily
type OmsCnSalesTargetDaily struct {
      global.GVA_MODEL
      Department  string `json:"department" form:"department" gorm:"column:department;comment:类型;size:255;"`  //类型 
      Month  string `json:"month" form:"month" gorm:"column:month;comment:月份;size:255;"`  //月份 
      Channel  string `json:"channel" form:"channel" gorm:"column:channel;comment:渠道;size:255;"`  //渠道 
      Version  string `json:"version" form:"version" gorm:"column:version;comment:版本;size:255;"`  //版本 
      Date  string `json:"date" form:"date" gorm:"column:date;comment:日期;size:255;"`  //日期 
      AmountTarget  string `json:"amountTarget" form:"amountTarget" gorm:"column:amount_target;comment:销售目标;size:255;"`  //销售目标 
      SheetName  string `json:"sheetName" form:"sheetName" gorm:"column:sheet_name;comment:手工表名称;size:64;"`  //手工表名称 
      CreatedName  string `json:"createdName" form:"createdName" gorm:"column:created_name;comment:创建人;size:32;"`  //创建人 
      CreatedBy  uint   `gorm:"column:created_by;comment:创建者"`
      UpdatedBy  uint   `gorm:"column:updated_by;comment:更新者"`
      DeletedBy  uint   `gorm:"column:deleted_by;comment:删除者"`
}


// TableName 京东自营销售录入表 OmsCnSalesTargetDaily自定义表名 oms_cn_sales_target_daily
func (OmsCnSalesTargetDaily) TableName() string {
  return "oms_cn_sales_target_daily"
}

