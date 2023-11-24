// 自动生成模板OmsCnChannelSales
package manual

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
	
)

// 按天销售目标录入表 结构体  OmsCnChannelSales
type OmsCnChannelSales struct {
      global.GVA_MODEL
      Date  string `json:"date" form:"date" gorm:"column:date;comment:月份;size:255;"`  //月份 
      Channel  string `json:"channel" form:"channel" gorm:"column:channel;comment:渠道;size:255;"`  //渠道 
      Department  string `json:"department" form:"department" gorm:"column:department;comment:类型;size:255;"`  //类型 
      Principal  string `json:"principal" form:"principal" gorm:"column:principal;comment:负责人;size:255;"`  //负责人 
      Amount  string `json:"amount" form:"amount" gorm:"column:amount;comment:回款收入(未税);size:255;"`  //回款收入(未税) 
      CreatedName  string `json:"createdName" form:"createdName" gorm:"column:created_name;comment:创建人;size:32;"`  //创建人 
      SheetName  string `json:"sheetName" form:"sheetName" gorm:"column:sheet_name;comment:手工表名称;size:64;"`  //手工表名称 
      CreatedBy  uint   `gorm:"column:created_by;comment:创建者"`
      UpdatedBy  uint   `gorm:"column:updated_by;comment:更新者"`
      DeletedBy  uint   `gorm:"column:deleted_by;comment:删除者"`
}


// TableName 按天销售目标录入表 OmsCnChannelSales自定义表名 oms_cn_channel_sales
func (OmsCnChannelSales) TableName() string {
  return "oms_cn_channel_sales"
}

