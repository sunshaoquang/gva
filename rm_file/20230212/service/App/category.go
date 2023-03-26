package App

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/App"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    AppReq "github.com/flipped-aurora/gin-vue-admin/server/model/App/request"
)

type CategoryService struct {
}

// CreateCategory 创建Category记录
// Author [piexlmax](https://github.com/piexlmax)
func (categoryService *CategoryService) CreateCategory(category App.Category) (err error) {
	err = global.GVA_DB.Create(&category).Error
	return err
}

// DeleteCategory 删除Category记录
// Author [piexlmax](https://github.com/piexlmax)
func (categoryService *CategoryService)DeleteCategory(category App.Category) (err error) {
	err = global.GVA_DB.Delete(&category).Error
	return err
}

// DeleteCategoryByIds 批量删除Category记录
// Author [piexlmax](https://github.com/piexlmax)
func (categoryService *CategoryService)DeleteCategoryByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]App.Category{},"id in ?",ids.Ids).Error
	return err
}

// UpdateCategory 更新Category记录
// Author [piexlmax](https://github.com/piexlmax)
func (categoryService *CategoryService)UpdateCategory(category App.Category) (err error) {
	err = global.GVA_DB.Save(&category).Error
	return err
}

// GetCategory 根据id获取Category记录
// Author [piexlmax](https://github.com/piexlmax)
func (categoryService *CategoryService)GetCategory(id uint) (category App.Category, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&category).Error
	return
}

// GetCategoryInfoList 分页获取Category记录
// Author [piexlmax](https://github.com/piexlmax)
func (categoryService *CategoryService)GetCategoryInfoList(info AppReq.CategorySearch) (list []App.Category, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&App.Category{})
    var categorys []App.Category
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
    if info.Name != "" {
        db = db.Where("name LIKE ?","%"+ info.Name+"%")
    }
    if info.Icon != "" {
        db = db.Where("icon LIKE ?","%"+ info.Icon+"%")
    }
    if info.state != nil {
        db = db.Where("state = ?",info.state)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	err = db.Limit(limit).Offset(offset).Find(&categorys).Error
	return  categorys, total, err
}
