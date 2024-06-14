package App

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/App"
	AppReq "github.com/flipped-aurora/gin-vue-admin/server/model/App/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type AppUserService struct {
}

// CreateAppUser 创建AppUser记录
// Author [piexlmax](https://github.com/piexlmax)
func (appUserService *AppUserService) CreateAppUser(appUser App.AppUser) (err error) {
	err = global.GVA_DB.Create(&appUser).Error
	return err
}

// DeleteAppUser 删除AppUser记录
// Author [piexlmax](https://github.com/piexlmax)
func (appUserService *AppUserService) DeleteAppUser(appUser App.AppUser) (err error) {
	err = global.GVA_DB.Delete(&appUser).Error
	return err
}

// DeleteAppUserByIds 批量删除AppUser记录
// Author [piexlmax](https://github.com/piexlmax)
func (appUserService *AppUserService) DeleteAppUserByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]App.AppUser{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateAppUser 更新AppUser记录
// Author [piexlmax](https://github.com/piexlmax)
func (appUserService *AppUserService) UpdateAppUser(appUser App.AppUser) (err error) {
	err = global.GVA_DB.Save(&appUser).Error
	return err
}

// GetAppUser 根据id获取AppUser记录
// Author [piexlmax](https://github.com/piexlmax)
func (appUserService *AppUserService) GetAppUser(id uint) (appUser App.AppUser, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&appUser).Error
	return
}

// GetAppUserInfoList 分页获取AppUser记录
// Author [piexlmax](https://github.com/piexlmax)
func (appUserService *AppUserService) GetAppUserInfoList(info AppReq.AppUserSearch) (list []App.AppUser, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&App.AppUser{})
	var appUsers []App.AppUser
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.ID != 0 {
		db = db.Where("id = ?", info.ID)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&appUsers).Error
	return appUsers, total, err
}
