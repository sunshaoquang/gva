// 自动生成模板WmsLogisticsProviderMi
package manual

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
	
)

// 费用by供应商 结构体  WmsLogisticsProviderMi
type WmsLogisticsProviderMi struct {
      global.GVA_MODEL
      Type  string `json:"type" form:"type" gorm:"column:type;comment:类型;size:255;"`  //类型 
      Area  string `json:"area" form:"area" gorm:"column:area;comment:区域;size:255;"`  //区域 
      Country  string `json:"country" form:"country" gorm:"column:country;comment:目的国;size:255;"`  //目的国 
      Supplier  string `json:"supplier" form:"supplier" gorm:"column:supplier;comment:供应商;size:255;"`  //供应商 
      Jan  string `json:"jan" form:"jan" gorm:"column:jan;comment:2023年1月;size:255;"`  //2023年1月 
      Feb  string `json:"feb" form:"feb" gorm:"column:feb;comment:2023年2月;size:255;"`  //2023年2月 
      Mar  string `json:"mar" form:"mar" gorm:"column:mar;comment:2023年3月;size:255;"`  //2023年3月 
      Apr  string `json:"apr" form:"apr" gorm:"column:apr;comment:2023年4月;size:255;"`  //2023年4月 
      May  string `json:"may" form:"may" gorm:"column:may;comment:2023年5月;size:255;"`  //2023年5月 
      Jun  string `json:"jun" form:"jun" gorm:"column:jun;comment:2023年6月;size:255;"`  //2023年6月 
      Jul  string `json:"jul" form:"jul" gorm:"column:jul;comment:2023年7月;size:255;"`  //2023年7月 
      Aug  string `json:"aug" form:"aug" gorm:"column:aug;comment:2023年8月;size:255;"`  //2023年8月 
      Sept  string `json:"sept" form:"sept" gorm:"column:sept;comment:2023年9月;size:255;"`  //2023年9月 
      Oct  string `json:"oct" form:"oct" gorm:"column:oct;comment:2023年10月;size:255;"`  //2023年10月 
      Nov  string `json:"nov" form:"nov" gorm:"column:nov;comment:2023年11月;size:255;"`  //2023年11月 
      Dec  string `json:"dec" form:"dec" gorm:"column:dec;comment:2023年12月;size:255;"`  //2023年12月 
      CreatedName  string `json:"createdName" form:"createdName" gorm:"column:created_name;comment:创建人;size:255;"`  //创建人 
      SheetName  string `json:"sheetName" form:"sheetName" gorm:"column:sheet_name;comment:手工表名称;size:255;"`  //手工表名称 
      CreatedBy  uint   `gorm:"column:created_by;comment:创建者"`
      UpdatedBy  uint   `gorm:"column:updated_by;comment:更新者"`
      DeletedBy  uint   `gorm:"column:deleted_by;comment:删除者"`
}


// TableName 费用by供应商 WmsLogisticsProviderMi自定义表名 wms_logistics_provider_mi
func (WmsLogisticsProviderMi) TableName() string {
  return "wms_logistics_provider_mi"
}

