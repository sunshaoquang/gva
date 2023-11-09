// 自动生成模板WmsKpiSummary2Mi
package manual

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
	
)

// wmsKpiSummary2Mi表 结构体  WmsKpiSummary2Mi
type WmsKpiSummary2Mi struct {
      global.GVA_MODEL
      Project  string `json:"project" form:"project" gorm:"column:project;comment:;size:255;"`  //项目字段 
      Year2021  string `json:"year2021" form:"year2021" gorm:"column:year2021;comment:;size:255;"`  //2021年字段 
      CreatedName  string `json:"createdName" form:"createdName" gorm:"column:created_name;comment:创建人;size:255;"`  //创建人 
      Year2022  string `json:"year2022" form:"year2022" gorm:"column:year2022;comment:;size:255;"`  //2022年字段 
      SheetName  string `json:"sheetName" form:"sheetName" gorm:"column:sheet_name;comment:手工表名称;size:255;"`  //手工表名称 
      Year2023  string `json:"year2023" form:"year2023" gorm:"column:year2023;comment:;size:255;"`  //2023年字段 
      CreatedBy  uint   `gorm:"column:created_by;comment:创建者"`
      UpdatedBy  uint   `gorm:"column:updated_by;comment:更新者"`
      DeletedBy  uint   `gorm:"column:deleted_by;comment:删除者"`
}


// TableName wmsKpiSummary2Mi表 WmsKpiSummary2Mi自定义表名 wms_kpi_summary2_mi
func (WmsKpiSummary2Mi) TableName() string {
  return "wms_kpi_summary2_mi"
}

