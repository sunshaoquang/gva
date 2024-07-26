// 自动生成模板TallyCategory
package tally

import "time"

// 记账类别表 结构体  TallyCategory
type TallyCategory struct {
	Id        *int       `json:"id" form:"id" gorm:"primarykey;column:id;comment:;size:20;"`         //id字段
	CreatedAt *time.Time `json:"createdAt" form:"createdAt" gorm:"column:created_at;comment:;"`      //createdAt字段
	UpdatedAt *time.Time `json:"updatedAt" form:"updatedAt" gorm:"column:updated_at;comment:;"`      //updatedAt字段
	DeletedAt *time.Time `json:"deletedAt" form:"deletedAt" gorm:"column:deleted_at;comment:;"`      //deletedAt字段
	Name      string     `json:"name" form:"name" gorm:"column:name;comment:分类标题;size:100;"`         //分类标题
	Icon      string     `json:"icon" form:"icon" gorm:"column:icon;comment:未选中分类图标;size:255;"`      //未选中分类图标
	ParId     *int       `json:"parId" form:"parId" gorm:"column:par_id;default:0;comment:分类父字段id;"` //分类父字段id
	State     *bool      `json:"state" form:"state" gorm:"column:state;comment:分类字段id;"`             //分类字段id
	Remark    string     `json:"remark" form:"remark" gorm:"column:remark;comment:备注字段;size:255;"`   //备注字段
	CreatedBy uint       `gorm:"column:created_by;comment:创建者"`
	UpdatedBy uint       `gorm:"column:updated_by;comment:更新者"`
	DeletedBy uint       `gorm:"column:deleted_by;comment:删除者"`
}

// TableName 记账类别表 TallyCategory自定义表名 tally_category
func (TallyCategory) TableName() string {
	return "tally_category"
}
