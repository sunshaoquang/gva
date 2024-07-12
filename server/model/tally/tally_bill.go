// 自动生成模板TallyBill
package tally

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// 记账账单表 结构体  TallyBill
type TallyBill struct {
    global.GVA_MODEL
    Icon  string `json:"icon" form:"icon" gorm:"column:icon;comment:选中分类图标;size:100;"`  //分类图标 
    Name  string `json:"name" form:"name" gorm:"column:name;comment:选中分类名称;size:100;"`  //分类名称 
    DeliveryMethod  string `json:"deliveryMethod" form:"deliveryMethod" gorm:"column:delivery_method;type:enum(20);comment:交付方式;"`  //交付方式 
    Money  *float64 `json:"money" form:"money" gorm:"column:money;comment:账单金额;"`  //账单金额 
    Tag  string `json:"tag" form:"tag" gorm:"column:tag;comment:备注;size:255;"`  //备注 
    CreatedBy  uint   `gorm:"column:created_by;comment:创建者"`
    UpdatedBy  uint   `gorm:"column:updated_by;comment:更新者"`
    DeletedBy  uint   `gorm:"column:deleted_by;comment:删除者"`
}


// TableName 记账账单表 TallyBill自定义表名 tally_bill
func (TallyBill) TableName() string {
    return "tally_bill"
}

