// 自动生成模板TallyBill
package tally

import (
	"time"

	"gorm.io/datatypes"
)

// 记账账单表 结构体  TallyBill
type TallyBill struct {
	Id               *int           `json:"id" form:"id" gorm:"primarykey;column:id;comment:;size:20;"`                              //id字段
	CreatedAt        *time.Time     `json:"createdAt" form:"createdAt" gorm:"column:created_at;comment:;"`                           //createdAt字段
	UpdatedAt        *time.Time     `json:"updatedAt" form:"updatedAt" gorm:"column:updated_at;comment:;"`                           //updatedAt字段
	DeletedAt        *time.Time     `json:"deletedAt" form:"deletedAt" gorm:"column:deleted_at;comment:;"`                           //分类名称
	CategoryId       uint           `json:"categoryId" form:"categoryId" gorm:"column:category_id;comment:选中分类的id;size:20;"`         //选中分类的id
	DeliveryMethodId uint           `json:"deliveryMethodId" form:"deliveryMethodId" gorm:"column:delivery_method_id;comment:交付方式;"` //交付方式的id
	Money            *float64       `json:"money" form:"money" gorm:"column:money;comment:账单金额;"`                                    //账单金额
	Tag              datatypes.JSON `json:"tag" form:"tag" gorm:"column:tag;comment:标签;type:text;"swaggertype:"array,object"`        //标签
	Remark           string         `json:"remark" form:"remark" gorm:"column:remark;comment:备注;size:255;"`                          //备注
	UserId           uint           `json:"userId" form:"userId" gorm:"column:user_id;comment:用户的id;size:20;"`
	CreatedBy        uint           `gorm:"column:created_by;comment:创建者"`
	UpdatedBy        uint           `gorm:"column:updated_by;comment:更新者"`
	DeletedBy        uint           `gorm:"column:deleted_by;comment:删除者"`
}

// TableName 记账账单表 TallyBill自定义表名 tally_bill
func (TallyBill) TableName() string {
	return "tally_bill"
}
