package manual

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/manual"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    manualReq "github.com/flipped-aurora/gin-vue-admin/server/model/manual/request"
    "gorm.io/gorm"
)

type WmsLogisticsSalebycountryMiService struct {
}

// CreateWmsLogisticsSalebycountryMi 创建收入by国家记录
// Author [piexlmax](https://github.com/piexlmax)
func (wmsLogisticsSalebycountryMiService *WmsLogisticsSalebycountryMiService) CreateWmsLogisticsSalebycountryMi(wmsLogisticsSalebycountryMi *manual.WmsLogisticsSalebycountryMi) (err error) {
	err = global.MustGetGlobalDBByDBName("ht_smartdata").Create(wmsLogisticsSalebycountryMi).Error
	return err
}

// DeleteWmsLogisticsSalebycountryMi 删除收入by国家记录
// Author [piexlmax](https://github.com/piexlmax)
func (wmsLogisticsSalebycountryMiService *WmsLogisticsSalebycountryMiService)DeleteWmsLogisticsSalebycountryMi(wmsLogisticsSalebycountryMi manual.WmsLogisticsSalebycountryMi) (err error) {
	err = global.MustGetGlobalDBByDBName("ht_smartdata").Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&manual.WmsLogisticsSalebycountryMi{}).Where("id = ?", wmsLogisticsSalebycountryMi.ID).Update("deleted_by", wmsLogisticsSalebycountryMi.DeletedBy).Error; err != nil {
              return err
        }
        if err = tx.Delete(&wmsLogisticsSalebycountryMi).Error; err != nil {
              return err
        }
        return nil
	})
	return err
}

// DeleteWmsLogisticsSalebycountryMiByIds 批量删除收入by国家记录
// Author [piexlmax](https://github.com/piexlmax)
func (wmsLogisticsSalebycountryMiService *WmsLogisticsSalebycountryMiService)DeleteWmsLogisticsSalebycountryMiByIds(ids request.IdsReq,deleted_by uint) (err error) {
	err = global.MustGetGlobalDBByDBName("ht_smartdata").Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&manual.WmsLogisticsSalebycountryMi{}).Where("id in ?", ids.Ids).Update("deleted_by", deleted_by).Error; err != nil {
            return err
        }
        if err := tx.Where("id in ?", ids.Ids).Delete(&manual.WmsLogisticsSalebycountryMi{}).Error; err != nil {
            return err
        }
        return nil
    })
	return err
}

// UpdateWmsLogisticsSalebycountryMi 更新收入by国家记录
// Author [piexlmax](https://github.com/piexlmax)
func (wmsLogisticsSalebycountryMiService *WmsLogisticsSalebycountryMiService)UpdateWmsLogisticsSalebycountryMi(wmsLogisticsSalebycountryMi manual.WmsLogisticsSalebycountryMi) (err error) {
	err = global.MustGetGlobalDBByDBName("ht_smartdata").Save(&wmsLogisticsSalebycountryMi).Error
	return err
}

// GetWmsLogisticsSalebycountryMi 根据id获取收入by国家记录
// Author [piexlmax](https://github.com/piexlmax)
func (wmsLogisticsSalebycountryMiService *WmsLogisticsSalebycountryMiService)GetWmsLogisticsSalebycountryMi(id uint) (wmsLogisticsSalebycountryMi manual.WmsLogisticsSalebycountryMi, err error) {
	err = global.MustGetGlobalDBByDBName("ht_smartdata").Where("id = ?", id).First(&wmsLogisticsSalebycountryMi).Error
	return
}

// GetWmsLogisticsSalebycountryMiInfoList 分页获取收入by国家记录
// Author [piexlmax](https://github.com/piexlmax)
func (wmsLogisticsSalebycountryMiService *WmsLogisticsSalebycountryMiService)GetWmsLogisticsSalebycountryMiInfoList(info manualReq.WmsLogisticsSalebycountryMiSearch) (list []manual.WmsLogisticsSalebycountryMi, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.MustGetGlobalDBByDBName("ht_smartdata").Model(&manual.WmsLogisticsSalebycountryMi{})
    var wmsLogisticsSalebycountryMis []manual.WmsLogisticsSalebycountryMi
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
	
	err = db.Find(&wmsLogisticsSalebycountryMis).Error
	return  wmsLogisticsSalebycountryMis, total, err
}
