// 自动生成模板WmsLogisticsPcs2023Mi
package manual

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
	
)

// 2023年物流成本明细表 结构体  WmsLogisticsPcs2023Mi
type WmsLogisticsPcs2023Mi struct {
      global.GVA_MODEL
      Country  string `json:"country" form:"country" gorm:"column:country;comment:国家;size:255;"`  //国家 
      ProductAbbreviation  string `json:"productAbbreviation" form:"productAbbreviation" gorm:"column:product_abbreviation;comment:产品简称;size:255;"`  //产品简称 
      Type  string `json:"type" form:"type" gorm:"column:type;comment:类型;size:255;"`  //类型 
      January  string `json:"january" form:"january" gorm:"column:january;comment:1月字段;size:255;"`  //1月 
      February  string `json:"february" form:"february" gorm:"column:february;comment:2月字段;size:255;"`  //2月 
      March  string `json:"march" form:"march" gorm:"column:march;comment:3月字段;size:255;"`  //3月 
      April  string `json:"april" form:"april" gorm:"column:april;comment:4月字段;size:255;"`  //4月 
      May  string `json:"may" form:"may" gorm:"column:may;comment:5月字段;size:255;"`  //5月 
      June  string `json:"june" form:"june" gorm:"column:june;comment:6月字段;size:255;"`  //6月 
      July  string `json:"july" form:"july" gorm:"column:july;comment:7月字段;size:255;"`  //7月 
      August  string `json:"august" form:"august" gorm:"column:august;comment:8月字段;size:255;"`  //8月 
      September  string `json:"september" form:"september" gorm:"column:september;comment:9月字段;size:255;"`  //9月 
      October  string `json:"october" form:"october" gorm:"column:october;comment:10月字段;size:255;"`  //10月 
      November  string `json:"november" form:"november" gorm:"column:november;comment:11月字段;size:255;"`  //11月 
      December  string `json:"december" form:"december" gorm:"column:december;comment:12月字段;size:255;"`  //12月 
      CreatedName  string `json:"createdName" form:"createdName" gorm:"column:created_name;comment:创建人;size:255;"`  //创建人 
      SheetName  string `json:"sheetName" form:"sheetName" gorm:"column:sheet_name;comment:手工表名称;size:255;"`  //手工表名称 
      CreatedBy  uint   `gorm:"column:created_by;comment:创建者"`
      UpdatedBy  uint   `gorm:"column:updated_by;comment:更新者"`
      DeletedBy  uint   `gorm:"column:deleted_by;comment:删除者"`
}


// TableName 2023年物流成本明细表 WmsLogisticsPcs2023Mi自定义表名 wms_logistics_pcs_mi
func (WmsLogisticsPcs2023Mi) TableName() string {
  return "wms_logistics_pcs_mi"
}

