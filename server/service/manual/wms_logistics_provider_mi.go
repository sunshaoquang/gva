package manual

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/manual"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    manualReq "github.com/flipped-aurora/gin-vue-admin/server/model/manual/request"
    "gorm.io/gorm"
)

type WmsLogisticsProviderMiService struct {
}

// CreateWmsLogisticsProviderMi 创建费用by供应商记录
// Author [piexlmax](https://github.com/piexlmax)
func (wmsLogisticsProviderMiService *WmsLogisticsProviderMiService) CreateWmsLogisticsProviderMi(wmsLogisticsProviderMi *manual.WmsLogisticsProviderMi) (err error) {
	err = global.MustGetGlobalDBByDBName("ht_smartdata").Create(wmsLogisticsProviderMi).Error
	return err
}

// DeleteWmsLogisticsProviderMi 删除费用by供应商记录
// Author [piexlmax](https://github.com/piexlmax)
func (wmsLogisticsProviderMiService *WmsLogisticsProviderMiService)DeleteWmsLogisticsProviderMi(wmsLogisticsProviderMi manual.WmsLogisticsProviderMi) (err error) {
	err = global.MustGetGlobalDBByDBName("ht_smartdata").Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&manual.WmsLogisticsProviderMi{}).Where("id = ?", wmsLogisticsProviderMi.ID).Update("deleted_by", wmsLogisticsProviderMi.DeletedBy).Error; err != nil {
              return err
        }
        if err = tx.Delete(&wmsLogisticsProviderMi).Error; err != nil {
              return err
        }
        return nil
	})
	return err
}

// DeleteWmsLogisticsProviderMiByIds 批量删除费用by供应商记录
// Author [piexlmax](https://github.com/piexlmax)
func (wmsLogisticsProviderMiService *WmsLogisticsProviderMiService)DeleteWmsLogisticsProviderMiByIds(ids request.IdsReq,deleted_by uint) (err error) {
	err = global.MustGetGlobalDBByDBName("ht_smartdata").Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&manual.WmsLogisticsProviderMi{}).Where("id in ?", ids.Ids).Update("deleted_by", deleted_by).Error; err != nil {
            return err
        }
        if err := tx.Where("id in ?", ids.Ids).Delete(&manual.WmsLogisticsProviderMi{}).Error; err != nil {
            return err
        }
        return nil
    })
	return err
}

// UpdateWmsLogisticsProviderMi 更新费用by供应商记录
// Author [piexlmax](https://github.com/piexlmax)
func (wmsLogisticsProviderMiService *WmsLogisticsProviderMiService)UpdateWmsLogisticsProviderMi(wmsLogisticsProviderMi manual.WmsLogisticsProviderMi) (err error) {
	err = global.MustGetGlobalDBByDBName("ht_smartdata").Save(&wmsLogisticsProviderMi).Error
	return err
}

// GetWmsLogisticsProviderMi 根据id获取费用by供应商记录
// Author [piexlmax](https://github.com/piexlmax)
func (wmsLogisticsProviderMiService *WmsLogisticsProviderMiService)GetWmsLogisticsProviderMi(id uint) (wmsLogisticsProviderMi manual.WmsLogisticsProviderMi, err error) {
	err = global.MustGetGlobalDBByDBName("ht_smartdata").Where("id = ?", id).First(&wmsLogisticsProviderMi).Error
	return
}

// GetWmsLogisticsProviderMiInfoList 分页获取费用by供应商记录
// Author [piexlmax](https://github.com/piexlmax)
func (wmsLogisticsProviderMiService *WmsLogisticsProviderMiService)GetWmsLogisticsProviderMiInfoList(info manualReq.WmsLogisticsProviderMiSearch) (list []manual.WmsLogisticsProviderMi, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.MustGetGlobalDBByDBName("ht_smartdata").Model(&manual.WmsLogisticsProviderMi{})
    var wmsLogisticsProviderMis []manual.WmsLogisticsProviderMi
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
    if info.Type != "" {
        db = db.Where("type LIKE ?","%"+ info.Type+"%")
    }
    if info.Area != "" {
        db = db.Where("area LIKE ?","%"+ info.Area+"%")
    }
    if info.Country != "" {
        db = db.Where("country LIKE ?","%"+ info.Country+"%")
    }
    if info.Supplier != "" {
        db = db.Where("supplier LIKE ?","%"+ info.Supplier+"%")
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

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }
	
	err = db.Find(&wmsLogisticsProviderMis).Error
	return  wmsLogisticsProviderMis, total, err
}
