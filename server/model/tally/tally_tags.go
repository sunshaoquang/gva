// 自动生成模板TallyTags
package tally

// 标签 结构体  TallyTags
type TallyTags struct {
	Id        *int   `json:"id" form:"id" gorm:"primarykey;column:id;comment:标签id;size:11;"`                         //标签id
	TagName   string `json:"tagName" form:"tagName" gorm:"column:tag_name;comment:标签名;size:50;"`                     //标签名
	TagColor  string `json:"tagColor" form:"tagColor" gorm:"default:#1d39c4;column:tag_color;comment:标签颜色;size:15;"` //标签颜色
	UserId    uint   `json:"userId" form:"userId" gorm:"column:user_id;comment:用户的id;size:10;"`
	CreatedBy uint   `gorm:"column:created_by;comment:创建者"`
	UpdatedBy uint   `gorm:"column:updated_by;comment:更新者"`
	DeletedBy uint   `gorm:"column:deleted_by;comment:删除者"`
}

// TableName 标签 TallyTags自定义表名 tally_tags
func (TallyTags) TableName() string {
	return "tally_tags"
}
