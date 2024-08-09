package tally

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/tally"
	tallyReq "github.com/flipped-aurora/gin-vue-admin/server/model/tally/request"
	"gorm.io/gorm"
)

type TallyTagsService struct{}

// CreateTallyTags 创建标签记录
// Author [piexlmax](https://github.com/piexlmax)
func (tallyTagsService *TallyTagsService) CreateTallyTags(tallyTags *tally.TallyTags) (err error) {
	err = global.GVA_DB.Create(tallyTags).Error
	return err
}

// DeleteTallyTags 删除标签记录
// Author [piexlmax](https://github.com/piexlmax)
func (tallyTagsService *TallyTagsService) DeleteTallyTags(id string, userID uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&tally.TallyTags{}).Where("id = ?", id).Update("deleted_by", userID).Error; err != nil {
			return err
		}
		if err = tx.Delete(&tally.TallyTags{}, "id = ?", id).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// DeleteTallyTagsByIds 批量删除标签记录
// Author [piexlmax](https://github.com/piexlmax)
func (tallyTagsService *TallyTagsService) DeleteTallyTagsByIds(ids []string, deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&tally.TallyTags{}).Where("id in ?", ids).Update("deleted_by", deleted_by).Error; err != nil {
			return err
		}
		if err := tx.Where("id in ?", ids).Delete(&tally.TallyTags{}).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// UpdateTallyTags 更新标签记录
// Author [piexlmax](https://github.com/piexlmax)
func (tallyTagsService *TallyTagsService) UpdateTallyTags(tallyTags tally.TallyTags) (err error) {
	err = global.GVA_DB.Model(&tally.TallyTags{}).Where("id = ?", tallyTags.Id).Updates(&tallyTags).Error
	return err
}

// GetTallyTags 根据id获取标签记录
// Author [piexlmax](https://github.com/piexlmax)
func (tallyTagsService *TallyTagsService) GetTallyTags(id string) (tallyTags tally.TallyTags, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&tallyTags).Error
	return
}

// GetTallyTags 根据用户id获取标签列表
// Author [piexlmax](https://github.com/piexlmax)
func (tallyTagsService *TallyTagsService) GetTallyTagsUserInfoList(userId uint) (list []tally.TallyTags, total int64, err error) {
	// 创建db
	db := global.GVA_DB.Model(&tally.TallyTags{})
	var tallyTags []tally.TallyTags
	// 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = global.GVA_DB.Where("user_id = ?", userId).Find(&tallyTags).Error
	return tallyTags, total, err
}

// GetTallyTagsInfoList 分页获取标签记录
// Author [piexlmax](https://github.com/piexlmax)
func (tallyTagsService *TallyTagsService) GetTallyTagsInfoList(info tallyReq.TallyTagsSearch) (list []tally.TallyTags, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&tally.TallyTags{})
	var tallyTagss []tally.TallyTags
	// 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&tallyTagss).Error
	return tallyTagss, total, err
}
