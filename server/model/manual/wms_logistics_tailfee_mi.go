// 自动生成模板WmsLogisticsTailfeeMi
package manual

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
	
)

// 尾程费用结构 结构体  WmsLogisticsTailfeeMi
type WmsLogisticsTailfeeMi struct {
      global.GVA_MODEL
      Area  string `json:"area" form:"area" gorm:"column:area;comment:区域;size:255;"`  //区域 
      CostType  string `json:"costType" form:"costType" gorm:"column:cost_type;comment:费用类型;size:255;"`  //费用类型 
      Mmonth  string `json:"mmonth" form:"mmonth" gorm:"column:mmonth;comment:月份;size:255;"`  //月份 
      ExcludingTaxCost  string `json:"excludingTaxCost" form:"excludingTaxCost" gorm:"column:excluding_tax_cost;comment:不含税费用;size:255;"`  //不含税费用 
      CreatedName  string `json:"createdName" form:"createdName" gorm:"column:created_name;comment:创建人;size:64;"`  //创建人 
      SheetName  string `json:"sheetName" form:"sheetName" gorm:"column:sheet_name;comment:手工表名称;size:64;"`  //手工表名称 
      CreatedBy  uint   `gorm:"column:created_by;comment:创建者"`
      UpdatedBy  uint   `gorm:"column:updated_by;comment:更新者"`
      DeletedBy  uint   `gorm:"column:deleted_by;comment:删除者"`
}


// TableName 尾程费用结构 WmsLogisticsTailfeeMi自定义表名 wms_logistics_tailfee_mi
func (WmsLogisticsTailfeeMi) TableName() string {
  return "wms_logistics_tailfee_mi"
}

