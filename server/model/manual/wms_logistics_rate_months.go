// 自动生成模板WmsLogisticsRateMonths
package manual

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
	
)

// wmsLogisticsRateMonths表 结构体  WmsLogisticsRateMonths
type WmsLogisticsRateMonths struct {
      global.GVA_MODEL
      Area  string `json:"area" form:"area" gorm:"column:area;comment:区域;size:255;"`  //区域 
      Mmonth  string `json:"mmonth" form:"mmonth" gorm:"column:mmonth;comment:月份;size:255;"`  //月份 
      Name  string `json:"name" form:"name" gorm:"column:name;comment:指标;size:255;"`  //指标 
      Value  string `json:"value" form:"value" gorm:"column:value;comment:指标值;size:255;"`  //指标值 
      CreatedName  string `json:"createdName" form:"createdName" gorm:"column:created_name;comment:创建人;size:64;"`  //创建人 
      SheetName  string `json:"sheetName" form:"sheetName" gorm:"column:sheet_name;comment:手工表名称;size:64;"`  //手工表名称 
      CreatedBy  uint   `gorm:"column:created_by;comment:创建者"`
      UpdatedBy  uint   `gorm:"column:updated_by;comment:更新者"`
      DeletedBy  uint   `gorm:"column:deleted_by;comment:删除者"`
}


// TableName wmsLogisticsRateMonths表 WmsLogisticsRateMonths自定义表名 wms_logistics_rate_months
func (WmsLogisticsRateMonths) TableName() string {
  return "wms_logistics_rate_months"
}

