package manual

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/manual"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    manualReq "github.com/flipped-aurora/gin-vue-admin/server/model/manual/request"
    "gorm.io/gorm"
)

type WmsLogisticsRateMonthsService struct {
}

// CreateWmsLogisticsRateMonths 创建wmsLogisticsRateMonths表记录
// Author [piexlmax](https://github.com/piexlmax)
func (wmsLogisticsRateMonthsService *WmsLogisticsRateMonthsService) CreateWmsLogisticsRateMonths(wmsLogisticsRateMonths *manual.WmsLogisticsRateMonths) (err error) {
	err = global.MustGetGlobalDBByDBName("ht_smartdata").Create(wmsLogisticsRateMonths).Error
	return err
}

// DeleteWmsLogisticsRateMonths 删除wmsLogisticsRateMonths表记录
// Author [piexlmax](https://github.com/piexlmax)
func (wmsLogisticsRateMonthsService *WmsLogisticsRateMonthsService)DeleteWmsLogisticsRateMonths(wmsLogisticsRateMonths manual.WmsLogisticsRateMonths) (err error) {
	err = global.MustGetGlobalDBByDBName("ht_smartdata").Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&manual.WmsLogisticsRateMonths{}).Where("id = ?", wmsLogisticsRateMonths.ID).Update("deleted_by", wmsLogisticsRateMonths.DeletedBy).Error; err != nil {
              return err
        }
        if err = tx.Delete(&wmsLogisticsRateMonths).Error; err != nil {
              return err
        }
        return nil
	})
	return err
}

// DeleteWmsLogisticsRateMonthsByIds 批量删除wmsLogisticsRateMonths表记录
// Author [piexlmax](https://github.com/piexlmax)
func (wmsLogisticsRateMonthsService *WmsLogisticsRateMonthsService)DeleteWmsLogisticsRateMonthsByIds(ids request.IdsReq,deleted_by uint) (err error) {
	err = global.MustGetGlobalDBByDBName("ht_smartdata").Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&manual.WmsLogisticsRateMonths{}).Where("id in ?", ids.Ids).Update("deleted_by", deleted_by).Error; err != nil {
            return err
        }
        if err := tx.Where("id in ?", ids.Ids).Delete(&manual.WmsLogisticsRateMonths{}).Error; err != nil {
            return err
        }
        return nil
    })
	return err
}

// UpdateWmsLogisticsRateMonths 更新wmsLogisticsRateMonths表记录
// Author [piexlmax](https://github.com/piexlmax)
func (wmsLogisticsRateMonthsService *WmsLogisticsRateMonthsService)UpdateWmsLogisticsRateMonths(wmsLogisticsRateMonths manual.WmsLogisticsRateMonths) (err error) {
	err = global.MustGetGlobalDBByDBName("ht_smartdata").Save(&wmsLogisticsRateMonths).Error
	return err
}

// GetWmsLogisticsRateMonths 根据id获取wmsLogisticsRateMonths表记录
// Author [piexlmax](https://github.com/piexlmax)
func (wmsLogisticsRateMonthsService *WmsLogisticsRateMonthsService)GetWmsLogisticsRateMonths(id uint) (wmsLogisticsRateMonths manual.WmsLogisticsRateMonths, err error) {
	err = global.MustGetGlobalDBByDBName("ht_smartdata").Where("id = ?", id).First(&wmsLogisticsRateMonths).Error
	return
}

// GetWmsLogisticsRateMonthsInfoList 分页获取wmsLogisticsRateMonths表记录
// Author [piexlmax](https://github.com/piexlmax)
func (wmsLogisticsRateMonthsService *WmsLogisticsRateMonthsService)GetWmsLogisticsRateMonthsInfoList(info manualReq.WmsLogisticsRateMonthsSearch) (list []manual.WmsLogisticsRateMonths, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.MustGetGlobalDBByDBName("ht_smartdata").Model(&manual.WmsLogisticsRateMonths{})
    var wmsLogisticsRateMonthss []manual.WmsLogisticsRateMonths
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
    if info.Area != "" {
        db = db.Where("area LIKE ?","%"+ info.Area+"%")
    }
    if info.Mmonth != "" {
        db = db.Where("mmonth LIKE ?","%"+ info.Mmonth+"%")
    }
    if info.Name != "" {
        db = db.Where("name LIKE ?","%"+ info.Name+"%")
    }
    if info.CreatedName != "" {
        db = db.Where("created_name LIKE ?","%"+ info.CreatedName+"%")
    }
    if info.SheetName != "" {
        db = db.Where("sheet_name LIKE ?","%"+ info.SheetName+"%")
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }
        var OrderStr string
        orderMap := make(map[string]bool)
         	orderMap["created_name"] = true
         	orderMap["sheet_name"] = true
       if orderMap[info.Sort] {
          OrderStr = info.Sort
          if info.Order == "descending" {
             OrderStr = OrderStr + " desc"
          }
          db = db.Order(OrderStr)
       }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }
	
	err = db.Find(&wmsLogisticsRateMonthss).Error
	return  wmsLogisticsRateMonthss, total, err
}
