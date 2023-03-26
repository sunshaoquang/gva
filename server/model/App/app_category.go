// 自动生成模板AppCategory
package App

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// AppCategory 结构体
type AppCategory struct {
	global.GVA_MODEL
	Name   string    `json:"name" form:"name" gorm:"column:name;comment:分类标题;size:100;"`
	Icon   string    `json:"icon" form:"icon" gorm:"column:icon;comment:未选中分类图标;size:255;"`
	ParId  *int      `json:"categoryPar" form:"categoryPar" gorm:"column:par_id;default:0;comment:分类父字段id;size:10;"`
	State  *bool     `json:"state" form:"state" gorm:"column:state;comment:分类字段id;"`
	Remark string    `json:"remark" form:"remark" gorm:"column:remark;comment:备注字段;size:255;"`
	User   []AppUser `gorm:"many2many:user_category;"`
}

// TableName AppCategory 表名
func (AppCategory) TableName() string {
	return "app_category"
}
