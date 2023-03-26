// 自动生成模板AppUser
package App

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// AppUser 结构体
type AppUser struct {
	global.GVA_MODEL
	Username   string        `json:"username" form:"username" gorm:"column:username;comment:用户名;size:60;"`
	Openid     string        `json:"openid" form:"openid" gorm:"uniqueIndex;column:openid;comment:微信唯一openid;size:150;"`
	Unionid    string        `json:"unionid" form:"unionid" gorm:"column:unionid;comment:unionid;size:100;"`
	LoginCount *int          `json:"loginCount" form:"loginCount" gorm:"column:login_count;comment:用户登陆的次数;size:10;"`
	Sex        *int          `json:"sex" form:"sex" gorm:"column:sex;comment:性别;size:10;"`
	Avatar     string        `json:"avatar" form:"avatar" gorm:"column:avatar;comment:微信头像;size:150;"`
	Category   []AppCategory `gorm:"many2many:user_category;"`
	Bill       []AppBill     `gorm:"foreignKey:UserId"`
	SysUserID  *int          `json:"sysUserId" form:"sysUserId"`
}

// TableName AppUser 表名
func (AppUser) TableName() string {
	return "app_user"
}
