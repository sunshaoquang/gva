package App

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/App"
	AppReq "github.com/flipped-aurora/gin-vue-admin/server/model/App/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type AppCategoryService struct {
}

// CreateAppCategory 创建AppCategory记录
// Author [piexlmax](https://github.com/piexlmax)
func (appCategoryService *AppCategoryService) CreateAppCategory(appCategory App.AppCategory, id int) (err error) {
	appUserDb := global.GVA_DB.Model(&App.AppUser{})
	var appUser []App.AppUser
	appUserDb.Where("sys_user_id = ?", id).Find(&appUser)
	appCategory.User = appUser
	err = global.GVA_DB.Create(&appCategory).Error
	return err
}

// DeleteAppCategory 删除AppCategory记录
// Author [piexlmax](https://github.com/piexlmax)
func (appCategoryService *AppCategoryService) DeleteAppCategory(appCategory App.AppCategory, id int) (err error) {
	appUserDb := global.GVA_DB.Model(&App.AppUser{})
	var appUser []App.AppUser
	appUserDb.Where("sys_user_id = ?", id).Find(&appUser)
	appCategory.User = appUser
	err = global.GVA_DB.Delete(&appCategory).Error
	return err
}

// DeleteAppCategoryByIds 批量删除AppCategory记录
// Author [piexlmax](https://github.com/piexlmax)
func (appCategoryService *AppCategoryService) DeleteAppCategoryByIds(ids request.IdsReq, id int) (err error) {
	appUserDb := global.GVA_DB.Model(&App.AppUser{})
	var appUser []App.AppUser
	appUserDb.Where("sys_user_id = ?", id).Find(&appUser)
	var appCategory []App.AppCategory
	for i, _ := range ids.Ids {
		appCategory[i].User = appUser
	}
	err = global.GVA_DB.Delete(&appCategory, "id in ?", ids.Ids).Error
	return err
}

// UpdateAppCategory 更新AppCategory记录
// Author [piexlmax](https://github.com/piexlmax)
func (appCategoryService *AppCategoryService) UpdateAppCategory(appCategory App.AppCategory, id int) (err error) {
	appUserDb := global.GVA_DB.Model(&App.AppUser{})
	var appUser []App.AppUser
	appUserDb.Where("sys_user_id = ?", id).Find(&appUser)
	appCategory.User = appUser
	err = global.GVA_DB.Save(&appCategory).Error
	return err
}

// GetAppCategory 根据id获取AppCategory记录
// Author [piexlmax](https://github.com/piexlmax)
func (appCategoryService *AppCategoryService) GetAppCategory(id uint) (appCategory App.AppCategory, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&appCategory).Error
	return
}

// GetAppCategoryInfoList 分页获取AppCategory记录
// Author [piexlmax](https://github.com/piexlmax)
func (appCategoryService *AppCategoryService) GetAppCategoryInfoList(info AppReq.AppCategorySearch) (list []App.AppCategory, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&App.AppCategory{})
	var appCategorys []App.AppCategory
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.Name != "" {
		db = db.Where("name LIKE ?", "%"+info.Name+"%")
	}
	if info.Icon != "" {
		db = db.Where("icon LIKE ?", "%"+info.Icon+"%")
	}
	if info.State != nil {
		db = db.Where("state = ?", info.State)
	}
	if info.ID != 0 {
		db = db.Where("id = ?", info.ID)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Preload("User").Find(&appCategorys).Error
	return appCategorys, total, err
}
