package manual

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/manual"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    manualReq "github.com/flipped-aurora/gin-vue-admin/server/model/manual/request"
    "gorm.io/gorm"
)

type WmsLogisticsPcsDetailMiService struct {
}

// CreateWmsLogisticsPcsDetailMi 创建主要产品成本明细表记录
// Author [piexlmax](https://github.com/piexlmax)
func (wmsLogisticsPcsDetailMiService *WmsLogisticsPcsDetailMiService) CreateWmsLogisticsPcsDetailMi(wmsLogisticsPcsDetailMi *manual.WmsLogisticsPcsDetailMi) (err error) {
	err = global.MustGetGlobalDBByDBName("ht_smartdata").Create(wmsLogisticsPcsDetailMi).Error
	return err
}

// DeleteWmsLogisticsPcsDetailMi 删除主要产品成本明细表记录
// Author [piexlmax](https://github.com/piexlmax)
func (wmsLogisticsPcsDetailMiService *WmsLogisticsPcsDetailMiService)DeleteWmsLogisticsPcsDetailMi(wmsLogisticsPcsDetailMi manual.WmsLogisticsPcsDetailMi) (err error) {
	err = global.MustGetGlobalDBByDBName("ht_smartdata").Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&manual.WmsLogisticsPcsDetailMi{}).Where("id = ?", wmsLogisticsPcsDetailMi.ID).Update("deleted_by", wmsLogisticsPcsDetailMi.DeletedBy).Error; err != nil {
              return err
        }
        if err = tx.Delete(&wmsLogisticsPcsDetailMi).Error; err != nil {
              return err
        }
        return nil
	})
	return err
}

// DeleteWmsLogisticsPcsDetailMiByIds 批量删除主要产品成本明细表记录
// Author [piexlmax](https://github.com/piexlmax)
func (wmsLogisticsPcsDetailMiService *WmsLogisticsPcsDetailMiService)DeleteWmsLogisticsPcsDetailMiByIds(ids request.IdsReq,deleted_by uint) (err error) {
	err = global.MustGetGlobalDBByDBName("ht_smartdata").Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&manual.WmsLogisticsPcsDetailMi{}).Where("id in ?", ids.Ids).Update("deleted_by", deleted_by).Error; err != nil {
            return err
        }
        if err := tx.Where("id in ?", ids.Ids).Delete(&manual.WmsLogisticsPcsDetailMi{}).Error; err != nil {
            return err
        }
        return nil
    })
	return err
}

// UpdateWmsLogisticsPcsDetailMi 更新主要产品成本明细表记录
// Author [piexlmax](https://github.com/piexlmax)
func (wmsLogisticsPcsDetailMiService *WmsLogisticsPcsDetailMiService)UpdateWmsLogisticsPcsDetailMi(wmsLogisticsPcsDetailMi manual.WmsLogisticsPcsDetailMi) (err error) {
	err = global.MustGetGlobalDBByDBName("ht_smartdata").Save(&wmsLogisticsPcsDetailMi).Error
	return err
}

// GetWmsLogisticsPcsDetailMi 根据id获取主要产品成本明细表记录
// Author [piexlmax](https://github.com/piexlmax)
func (wmsLogisticsPcsDetailMiService *WmsLogisticsPcsDetailMiService)GetWmsLogisticsPcsDetailMi(id uint) (wmsLogisticsPcsDetailMi manual.WmsLogisticsPcsDetailMi, err error) {
	err = global.MustGetGlobalDBByDBName("ht_smartdata").Where("id = ?", id).First(&wmsLogisticsPcsDetailMi).Error
	return
}

// GetWmsLogisticsPcsDetailMiInfoList 分页获取主要产品成本明细表记录
// Author [piexlmax](https://github.com/piexlmax)
func (wmsLogisticsPcsDetailMiService *WmsLogisticsPcsDetailMiService)GetWmsLogisticsPcsDetailMiInfoList(info manualReq.WmsLogisticsPcsDetailMiSearch) (list []manual.WmsLogisticsPcsDetailMi, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.MustGetGlobalDBByDBName("ht_smartdata").Model(&manual.WmsLogisticsPcsDetailMi{})
    var wmsLogisticsPcsDetailMis []manual.WmsLogisticsPcsDetailMi
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
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
	
	err = db.Find(&wmsLogisticsPcsDetailMis).Error
	return  wmsLogisticsPcsDetailMis, total, err
}
