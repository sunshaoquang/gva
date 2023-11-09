// 自动生成模板WmsLogisticsPcsDetailMi
package manual

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
	
)

// 主要产品成本明细表 结构体  WmsLogisticsPcsDetailMi
type WmsLogisticsPcsDetailMi struct {
      global.GVA_MODEL
      ProductAbbreviation  string `json:"productAbbreviation" form:"productAbbreviation" gorm:"column:product_abbreviation;comment:产品简称;size:255;"`  //产品简称 
      Country  string `json:"country" form:"country" gorm:"column:country;comment:国家;size:255;"`  //国家 
      FeeTypeCNY  string `json:"feeTypeCNY" form:"feeTypeCNY" gorm:"column:fee_type_cny;comment:费用类型CNY;size:255;"`  //费用类型CNY 
      Year2022Month10  string `json:"year2022Month10" form:"year2022Month10" gorm:"column:year2022_month10;comment:降本基数(2022年10月);size:255;"`  //降本基数(2022年10月) 
      Year2022Month11  string `json:"year2022Month11" form:"year2022Month11" gorm:"column:year2022_month11;comment:2022年11月;size:255;"`  //2022年11月 
      Year2022Month12  string `json:"year2022Month12" form:"year2022Month12" gorm:"column:year2022_month12;comment:2022年12月;size:255;"`  //2022年12月 
      Year2023Month1  string `json:"year2023Month1" form:"year2023Month1" gorm:"column:year2023_month1;comment:2023年1月;size:255;"`  //2023年1月 
      Year2023Month2  string `json:"year2023Month2" form:"year2023Month2" gorm:"column:year2023_month2;comment:2023年2月;size:255;"`  //2023年2月 
      Year2023Month3  string `json:"year2023Month3" form:"year2023Month3" gorm:"column:year2023_month3;comment:2023年3月;size:255;"`  //2023年3月 
      Year2023Month4  string `json:"year2023Month4" form:"year2023Month4" gorm:"column:year2023_month4;comment:2023年4月;size:255;"`  //2023年4月 
      Year2023Month5  string `json:"year2023Month5" form:"year2023Month5" gorm:"column:year2023_month5;comment:2023年5月;size:255;"`  //2023年5月 
      Year2023Month6  string `json:"year2023Month6" form:"year2023Month6" gorm:"column:year2023_month6;comment:2023年6月;size:255;"`  //2023年6月 
      Year2023Month7  string `json:"year2023Month7" form:"year2023Month7" gorm:"column:year2023_month7;comment:2023年7月;size:255;"`  //2023年7月 
      Year2023Month8  string `json:"year2023Month8" form:"year2023Month8" gorm:"column:year2023_month8;comment:2023年8月;size:255;"`  //2023年8月 
      Year2023Month9  string `json:"year2023Month9" form:"year2023Month9" gorm:"column:year2023_month9;comment:2023年9月;size:255;"`  //2023年9月 
      Year2023Month10  string `json:"year2023Month10" form:"year2023Month10" gorm:"column:year2023_month10;comment:2023年10月;size:255;"`  //2023年10月 
      Year2023Month11  string `json:"year2023Month11" form:"year2023Month11" gorm:"column:year2023_month11;comment:2023年11月;size:255;"`  //2023年11月 
      Year2023Month12  string `json:"year2023Month12" form:"year2023Month12" gorm:"column:year2023_month12;comment:2023年12月;size:255;"`  //2023年12月 
      CreatedName  string `json:"createdName" form:"createdName" gorm:"column:created_name;comment:创建人;size:255;"`  //创建人 
      SheetName  string `json:"sheetName" form:"sheetName" gorm:"column:sheet_name;comment:手工表名称;size:255;"`  //手工表名称 
      CreatedBy  uint   `gorm:"column:created_by;comment:创建者"`
      UpdatedBy  uint   `gorm:"column:updated_by;comment:更新者"`
      DeletedBy  uint   `gorm:"column:deleted_by;comment:删除者"`
}


// TableName 主要产品成本明细表 WmsLogisticsPcsDetailMi自定义表名 wms_logistics_pcs_detail_mi
func (WmsLogisticsPcsDetailMi) TableName() string {
  return "wms_logistics_pcs_detail_mi"
}

