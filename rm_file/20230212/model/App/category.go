// 自动生成模板Category
package App

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
)

// Category 结构体
type Category struct {
      global.GVA_MODEL
      Name  string `json:"name" form:"name" gorm:"column:name;comment:分类标题;size:100;"`
      Icon  string `json:"icon" form:"icon" gorm:"column:icon;comment:未选中分类图标;size:255;"`
      ParId  *int `json:"categoryPar" form:"categoryPar" gorm:"column:par_id;comment:分类父字段id;size:10;"`
      state  *bool `json:"state" form:"state" gorm:"column:state;comment:分类字段id;"`
      Remark  string `json:"remark" form:"remark" gorm:"column:remark;comment:备注字段;size:255;"`
}


// TableName Category 表名
func (Category) TableName() string {
  return "category"
}

