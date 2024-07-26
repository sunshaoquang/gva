package tally

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/tally"
	tallyReq "github.com/flipped-aurora/gin-vue-admin/server/model/tally/request"
	"gorm.io/gorm"
)

type TallyCategoryService struct{}

// CreateTallyCategory 创建记账类别表记录
// Author [piexlmax](https://github.com/piexlmax)
func (tallyCategoryService *TallyCategoryService) CreateTallyCategory(tallyCategory *tally.TallyCategory) (err error) {
	err = global.GVA_DB.Create(tallyCategory).Error
	return err
}

// DeleteTallyCategory 删除记账类别表记录
// Author [piexlmax](https://github.com/piexlmax)
func (tallyCategoryService *TallyCategoryService) DeleteTallyCategory(id string, userID uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&tally.TallyCategory{}).Where("id = ?", id).Update("deleted_by", userID).Error; err != nil {
			return err
		}
		if err = tx.Delete(&tally.TallyCategory{}, "id = ?", id).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// DeleteTallyCategoryByIds 批量删除记账类别表记录
// Author [piexlmax](https://github.com/piexlmax)
func (tallyCategoryService *TallyCategoryService) DeleteTallyCategoryByIds(ids []string, deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&tally.TallyCategory{}).Where("id in ?", ids).Update("deleted_by", deleted_by).Error; err != nil {
			return err
		}
		if err := tx.Where("id in ?", ids).Delete(&tally.TallyCategory{}).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// UpdateTallyCategory 更新记账类别表记录
// Author [piexlmax](https://github.com/piexlmax)
func (tallyCategoryService *TallyCategoryService) UpdateTallyCategory(tallyCategory tally.TallyCategory) (err error) {
	err = global.GVA_DB.Model(&tally.TallyCategory{}).Where("id = ?", tallyCategory.Id).Updates(&tallyCategory).Error
	return err
}

// GetTallyCategory 根据id获取记账类别表记录
// Author [piexlmax](https://github.com/piexlmax)
func (tallyCategoryService *TallyCategoryService) GetTallyCategory(id string) (tallyCategory tally.TallyCategory, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&tallyCategory).Error
	return
}

// GetTallyCategoryInfoList 分页获取记账类别表记录
// Author [piexlmax](https://github.com/piexlmax)
func (tallyCategoryService *TallyCategoryService) GetTallyCategoryInfoList(info tallyReq.TallyCategorySearch) (list []tally.TallyCategory, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&tally.TallyCategory{})
	var tallyCategorys []tally.TallyCategory
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
	if info.Id != nil {
		db = db.Where("id = ?", info.Id)
	}
	if info.ParId != nil {
		db = db.Where("par_id = ?", info.ParId)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&tallyCategorys).Error
	return tallyCategorys, total, err
}
