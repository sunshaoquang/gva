package tally

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/tally"
	tallyReq "github.com/flipped-aurora/gin-vue-admin/server/model/tally/request"
	"gorm.io/gorm"
)

type TallyBillService struct{}

// CreateTallyBill 创建记账账单表记录
// Author [piexlmax](https://github.com/piexlmax)
func (tallyBillService *TallyBillService) CreateTallyBill(tallyBill *tally.TallyBill) (err error) {
	err = global.GVA_DB.Create(tallyBill).Error
	return err
}

// DeleteTallyBill 删除记账账单表记录
// Author [piexlmax](https://github.com/piexlmax)
func (tallyBillService *TallyBillService) DeleteTallyBill(Id string, userID uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&tally.TallyBill{}).Where("id = ?", Id).Update("deleted_by", userID).Error; err != nil {
			return err
		}
		if err = tx.Delete(&tally.TallyBill{}, "id = ?", Id).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// DeleteTallyBillByIds 批量删除记账账单表记录
// Author [piexlmax](https://github.com/piexlmax)
func (tallyBillService *TallyBillService) DeleteTallyBillByIds(IDs []string, deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&tally.TallyBill{}).Where("id in ?", IDs).Update("deleted_by", deleted_by).Error; err != nil {
			return err
		}
		if err := tx.Where("id in ?", IDs).Delete(&tally.TallyBill{}).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// UpdateTallyBill 更新记账账单表记录
// Author [piexlmax](https://github.com/piexlmax)
func (tallyBillService *TallyBillService) UpdateTallyBill(tallyBill tally.TallyBill) (err error) {
	err = global.GVA_DB.Model(&tally.TallyBill{}).Where("id = ?", tallyBill.Id).Updates(&tallyBill).Error
	return err
}

// GetTallyBill 根据ID获取记账账单表记录
// Author [piexlmax](https://github.com/piexlmax)
func (tallyBillService *TallyBillService) GetTallyBill(Id string) (tallyBill tally.TallyBill, err error) {
	err = global.GVA_DB.Where("id = ?", Id).First(&tallyBill).Error
	return
}

// GetTallyBillInfoList 分页获取记账账单表记录
// Author [piexlmax](https://github.com/piexlmax)
func (tallyBillService *TallyBillService) GetTallyBillInfoList(info tallyReq.TallyBillSearch) (list []tally.TallyBill, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&tally.TallyBill{})
	var tallyBills []tally.TallyBill
	// 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}
	err = db.Find(&tallyBills).Error
	return tallyBills, total, err
}
