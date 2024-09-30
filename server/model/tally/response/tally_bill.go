package response

import "github.com/flipped-aurora/gin-vue-admin/server/model/tally"

type TallyBillDetailResult struct {
	List interface{} `json:"list"`
}

type TallyBillDetailDataList struct {
	tally.TallyBill
	Name string `json:"name" form:"name" gorm:"column:name;comment:分类标题;size:100;"`    //分类标题
	Icon string `json:"icon" form:"icon" gorm:"column:icon;comment:未选中分类图标;size:255;"` //未选中分类图标
}
