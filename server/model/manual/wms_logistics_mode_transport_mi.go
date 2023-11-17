// 自动生成模板WmsLogisticsModeTransportMi
package manual

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
	
)

// 费用by运输方式 结构体  WmsLogisticsModeTransportMi
type WmsLogisticsModeTransportMi struct {
      global.GVA_MODEL
      Area  string `json:"area" form:"area" gorm:"column:area;comment:区域;size:255;"`  //区域 
      TransMethod  string `json:"transMethod" form:"transMethod" gorm:"column:trans_method;comment:运输方式;size:255;"`  //运输方式 
      Mmonth  string `json:"mmonth" form:"mmonth" gorm:"column:mmonth;comment:月份;size:255;"`  //月份 
      Cost  string `json:"cost" form:"cost" gorm:"column:cost;comment:费用;size:255;"`  //费用 
      CreatedName  string `json:"createdName" form:"createdName" gorm:"column:created_name;comment:创建人;size:64;"`  //创建人 
      SheetName  string `json:"sheetName" form:"sheetName" gorm:"column:sheet_name;comment:手工表名称;size:64;"`  //手工表名称 
      CreatedBy  uint   `gorm:"column:created_by;comment:创建者"`
      UpdatedBy  uint   `gorm:"column:updated_by;comment:更新者"`
      DeletedBy  uint   `gorm:"column:deleted_by;comment:删除者"`
}


// TableName 费用by运输方式 WmsLogisticsModeTransportMi自定义表名 wms_logistics_mode_transport_mi
func (WmsLogisticsModeTransportMi) TableName() string {
  return "wms_logistics_mode_transport_mi"
}

