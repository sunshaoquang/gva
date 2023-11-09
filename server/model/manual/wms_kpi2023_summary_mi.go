// 自动生成模板WmsKpi2023SummaryMi
package manual

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
	
)

// wmsKpi2023SummaryMi表 结构体  WmsKpi2023SummaryMi
type WmsKpi2023SummaryMi struct {
      global.GVA_MODEL
      IndicatorClassification  string `json:"indicatorClassification" form:"indicatorClassification" gorm:"column:indicator_classification;comment:指标分类字段;size:255;"`  //指标分类字段 
      IndicatorName  string `json:"indicatorName" form:"indicatorName" gorm:"column:indicator_name;comment:指标名称字段;size:255;"`  //指标名称字段 
      Target  string `json:"target" form:"target" gorm:"column:target;comment:目标值字段;size:255;"`  //目标值字段 
      January  string `json:"january" form:"january" gorm:"column:january;comment:1月字段;size:255;"`  //1月字段 
      February  string `json:"february" form:"february" gorm:"column:february;comment:2月字段;size:255;"`  //2月字段 
      March  string `json:"march" form:"march" gorm:"column:march;comment:3月字段;size:255;"`  //3月字段 
      April  string `json:"april" form:"april" gorm:"column:april;comment:4月字段;size:255;"`  //4月字段 
      May  string `json:"may" form:"may" gorm:"column:may;comment:5月字段;size:255;"`  //5月字段 
      June  string `json:"june" form:"june" gorm:"column:june;comment:6月字段;size:255;"`  //6月字段 
      July  string `json:"july" form:"july" gorm:"column:july;comment:7月字段;size:255;"`  //7月字段 
      August  string `json:"august" form:"august" gorm:"column:august;comment:8月字段;size:255;"`  //8月字段 
      September  string `json:"september" form:"september" gorm:"column:september;comment:9月字段;size:255;"`  //9月字段 
      October  string `json:"october" form:"october" gorm:"column:october;comment:10月字段;size:255;"`  //10月字段 
      November  string `json:"november" form:"november" gorm:"column:november;comment:11月字段;size:255;"`  //11月字段 
      December  string `json:"december" form:"december" gorm:"column:december;comment:12月字段;size:255;"`  //12月字段 
      CreatedName  string `json:"createdName" form:"createdName" gorm:"column:created_name;comment:创建人;size:255;"`  //创建人 
      SheetName  string `json:"sheetName" form:"sheetName" gorm:"column:sheet_name;comment:手工表名称;size:64;"`  //手工表名称 
      CreatedBy  uint   `gorm:"column:created_by;comment:创建者"`
      UpdatedBy  uint   `gorm:"column:updated_by;comment:更新者"`
      DeletedBy  uint   `gorm:"column:deleted_by;comment:删除者"`
}


// TableName wmsKpi2023SummaryMi表 WmsKpi2023SummaryMi自定义表名 wms_kpi_2023_summary_mi
func (WmsKpi2023SummaryMi) TableName() string {
  return "wms_kpi_2023_summary_mi"
}

