package tally

import (
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/tally"
	tallyReq "github.com/flipped-aurora/gin-vue-admin/server/model/tally/request"
	tallyRes "github.com/flipped-aurora/gin-vue-admin/server/model/tally/response"
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

// GetTallyBillInfoList 根据用户Id获取用户记账账单表详细列表接口
// Author [piexlmax](https://github.com/piexlmax)
func (tallyBillService *TallyBillService) GetTallyBillDetailDataInfoList(id uint, info tallyReq.TallyBillDetailSearch) (list []tallyRes.TallyBillDetailDataList, err error) {
	// 创建db
	db := global.GVA_DB.Model(&tally.TallyBill{})
	var tallyBills []tallyRes.TallyBillDetailDataList
	if id != 0 {
		db = db.Where("user_id = ? AND tally_bill.deleted_at is null", id)
	}
	// 根据 info.TimeType 来判断当前需要过滤的info.queryTime 时间是开始年 周 月 日 到结束 时间
	switch info.TimeType {
	case "year":
		db = db.Where("tally_bill.created_at BETWEEN ? AND ?", time.Date(info.QueryTime.Year(), 1, 1, 0, 0, 0, 0, time.Local), time.Date(info.QueryTime.Year()+1, 1, 1, 0, 0, 0, 0, time.Local))
	case "month":
		db = db.Where("tally_bill.created_at BETWEEN ? AND ?", time.Date(info.QueryTime.Year(), info.QueryTime.Month(), 1, 0, 0, 0, 0, time.Local), time.Date(info.QueryTime.Year(), info.QueryTime.Month()+1, 1, 0, 0, 0, 0, time.Local))
	case "day":
		db = db.Where("tally_bill.created_at BETWEEN ? AND ?", time.Date(info.QueryTime.Year(), info.QueryTime.Month(), info.QueryTime.Day(), 0, 0, 0, 0, time.Local), time.Date(info.QueryTime.Year(), info.QueryTime.Month(), info.QueryTime.Day()+1, 0, 0, 0, 0, time.Local))
	}
	err = db.Debug().Joins("JOIN tally_category as c ON tally_bill.delivery_method_id = c.id").Select("tally_bill.*", "c.name", "c.icon").Order("tally_bill.created_at desc").Find(&tallyBills).Error
	if err != nil {
		return
	}
	return tallyBills, err
}
