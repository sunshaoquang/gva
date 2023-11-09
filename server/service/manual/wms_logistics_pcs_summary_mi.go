package manual

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/manual"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    manualReq "github.com/flipped-aurora/gin-vue-admin/server/model/manual/request"
    "gorm.io/gorm"
)

type WmsLogisticsPcsSummaryMiService struct {
}

// CreateWmsLogisticsPcsSummaryMi 创建主要产品成本降本汇总表记录
// Author [piexlmax](https://github.com/piexlmax)
func (wmsLogisticsPcsSummaryMiService *WmsLogisticsPcsSummaryMiService) CreateWmsLogisticsPcsSummaryMi(wmsLogisticsPcsSummaryMi *manual.WmsLogisticsPcsSummaryMi) (err error) {
	err = global.MustGetGlobalDBByDBName("ht_smartdata").Create(wmsLogisticsPcsSummaryMi).Error
	return err
}

// DeleteWmsLogisticsPcsSummaryMi 删除主要产品成本降本汇总表记录
// Author [piexlmax](https://github.com/piexlmax)
func (wmsLogisticsPcsSummaryMiService *WmsLogisticsPcsSummaryMiService)DeleteWmsLogisticsPcsSummaryMi(wmsLogisticsPcsSummaryMi manual.WmsLogisticsPcsSummaryMi) (err error) {
	err = global.MustGetGlobalDBByDBName("ht_smartdata").Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&manual.WmsLogisticsPcsSummaryMi{}).Where("id = ?", wmsLogisticsPcsSummaryMi.ID).Update("deleted_by", wmsLogisticsPcsSummaryMi.DeletedBy).Error; err != nil {
              return err
        }
        if err = tx.Delete(&wmsLogisticsPcsSummaryMi).Error; err != nil {
              return err
        }
        return nil
	})
	return err
}

// DeleteWmsLogisticsPcsSummaryMiByIds 批量删除主要产品成本降本汇总表记录
// Author [piexlmax](https://github.com/piexlmax)
func (wmsLogisticsPcsSummaryMiService *WmsLogisticsPcsSummaryMiService)DeleteWmsLogisticsPcsSummaryMiByIds(ids request.IdsReq,deleted_by uint) (err error) {
	err = global.MustGetGlobalDBByDBName("ht_smartdata").Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&manual.WmsLogisticsPcsSummaryMi{}).Where("id in ?", ids.Ids).Update("deleted_by", deleted_by).Error; err != nil {
            return err
        }
        if err := tx.Where("id in ?", ids.Ids).Delete(&manual.WmsLogisticsPcsSummaryMi{}).Error; err != nil {
            return err
        }
        return nil
    })
	return err
}

// UpdateWmsLogisticsPcsSummaryMi 更新主要产品成本降本汇总表记录
// Author [piexlmax](https://github.com/piexlmax)
func (wmsLogisticsPcsSummaryMiService *WmsLogisticsPcsSummaryMiService)UpdateWmsLogisticsPcsSummaryMi(wmsLogisticsPcsSummaryMi manual.WmsLogisticsPcsSummaryMi) (err error) {
	err = global.MustGetGlobalDBByDBName("ht_smartdata").Save(&wmsLogisticsPcsSummaryMi).Error
	return err
}

// GetWmsLogisticsPcsSummaryMi 根据id获取主要产品成本降本汇总表记录
// Author [piexlmax](https://github.com/piexlmax)
func (wmsLogisticsPcsSummaryMiService *WmsLogisticsPcsSummaryMiService)GetWmsLogisticsPcsSummaryMi(id uint) (wmsLogisticsPcsSummaryMi manual.WmsLogisticsPcsSummaryMi, err error) {
	err = global.MustGetGlobalDBByDBName("ht_smartdata").Where("id = ?", id).First(&wmsLogisticsPcsSummaryMi).Error
	return
}

// GetWmsLogisticsPcsSummaryMiInfoList 分页获取主要产品成本降本汇总表记录
// Author [piexlmax](https://github.com/piexlmax)
func (wmsLogisticsPcsSummaryMiService *WmsLogisticsPcsSummaryMiService)GetWmsLogisticsPcsSummaryMiInfoList(info manualReq.WmsLogisticsPcsSummaryMiSearch) (list []manual.WmsLogisticsPcsSummaryMi, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.MustGetGlobalDBByDBName("ht_smartdata").Model(&manual.WmsLogisticsPcsSummaryMi{})
    var wmsLogisticsPcsSummaryMis []manual.WmsLogisticsPcsSummaryMi
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
    if info.ProductAbbreviation != "" {
        db = db.Where("product_abbreviation LIKE ?","%"+ info.ProductAbbreviation+"%")
    }
    if info.Country != "" {
        db = db.Where("country LIKE ?","%"+ info.Country+"%")
    }
    if info.CreatedName != "" {
        db = db.Where("created_name LIKE ?","%"+ info.CreatedName+"%")
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }
	
	err = db.Find(&wmsLogisticsPcsSummaryMis).Error
	return  wmsLogisticsPcsSummaryMis, total, err
}
