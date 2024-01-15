// 自动生成模板AppBill
package App

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
)

// AppBill 结构体
type AppBill struct {
      global.GVA_MODEL
      Icon  string `json:"icon" form:"icon" gorm:"column:icon;comment:选中分类图标;size:100;"`
      Name  string `json:"name" form:"name" gorm:"column:name;comment:选中分类名称;size:100;"`
      DeliveryMethod  string `json:"deliveryMethod" form:"deliveryMethod" gorm:"column:delivery_method;comment:交付方式;size:20;"`
      Money  *float64 `json:"money" form:"money" gorm:"column:money;comment:账单金额;size:10;"`
      Tag  string `json:"tag" form:"tag" gorm:"column:tag;comment:备注;size:255;"`
      UserId  *int `json:"userId" form:"userId" gorm:"column:user_id;comment:某个用户创建的id;size:10;"`
}


// TableName AppBill 表名
func (AppBill) TableName() string {
  return "app_bill"
}

