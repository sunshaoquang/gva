package manual

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/manual"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    manualReq "github.com/flipped-aurora/gin-vue-admin/server/model/manual/request"
    "gorm.io/gorm"
)

type WmsLogisticsStoragefeeMiService struct {
}

// CreateWmsLogisticsStoragefeeMi 创建仓储费记录
// Author [piexlmax](https://github.com/piexlmax)
func (wmsLogisticsStoragefeeMiService *WmsLogisticsStoragefeeMiService) CreateWmsLogisticsStoragefeeMi(wmsLogisticsStoragefeeMi *manual.WmsLogisticsStoragefeeMi) (err error) {
	err = global.MustGetGlobalDBByDBName("ht_smartdata").Create(wmsLogisticsStoragefeeMi).Error
	return err
}

// DeleteWmsLogisticsStoragefeeMi 删除仓储费记录
// Author [piexlmax](https://github.com/piexlmax)
func (wmsLogisticsStoragefeeMiService *WmsLogisticsStoragefeeMiService)DeleteWmsLogisticsStoragefeeMi(wmsLogisticsStoragefeeMi manual.WmsLogisticsStoragefeeMi) (err error) {
	err = global.MustGetGlobalDBByDBName("ht_smartdata").Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&manual.WmsLogisticsStoragefeeMi{}).Where("id = ?", wmsLogisticsStoragefeeMi.ID).Update("deleted_by", wmsLogisticsStoragefeeMi.DeletedBy).Error; err != nil {
              return err
        }
        if err = tx.Delete(&wmsLogisticsStoragefeeMi).Error; err != nil {
              return err
        }
        return nil
	})
	return err
}

// DeleteWmsLogisticsStoragefeeMiByIds 批量删除仓储费记录
// Author [piexlmax](https://github.com/piexlmax)
func (wmsLogisticsStoragefeeMiService *WmsLogisticsStoragefeeMiService)DeleteWmsLogisticsStoragefeeMiByIds(ids request.IdsReq,deleted_by uint) (err error) {
	err = global.MustGetGlobalDBByDBName("ht_smartdata").Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&manual.WmsLogisticsStoragefeeMi{}).Where("id in ?", ids.Ids).Update("deleted_by", deleted_by).Error; err != nil {
            return err
        }
        if err := tx.Where("id in ?", ids.Ids).Delete(&manual.WmsLogisticsStoragefeeMi{}).Error; err != nil {
            return err
        }
        return nil
    })
	return err
}

// UpdateWmsLogisticsStoragefeeMi 更新仓储费记录
// Author [piexlmax](https://github.com/piexlmax)
func (wmsLogisticsStoragefeeMiService *WmsLogisticsStoragefeeMiService)UpdateWmsLogisticsStoragefeeMi(wmsLogisticsStoragefeeMi manual.WmsLogisticsStoragefeeMi) (err error) {
	err = global.MustGetGlobalDBByDBName("ht_smartdata").Save(&wmsLogisticsStoragefeeMi).Error
	return err
}

// GetWmsLogisticsStoragefeeMi 根据id获取仓储费记录
// Author [piexlmax](https://github.com/piexlmax)
func (wmsLogisticsStoragefeeMiService *WmsLogisticsStoragefeeMiService)GetWmsLogisticsStoragefeeMi(id uint) (wmsLogisticsStoragefeeMi manual.WmsLogisticsStoragefeeMi, err error) {
	err = global.MustGetGlobalDBByDBName("ht_smartdata").Where("id = ?", id).First(&wmsLogisticsStoragefeeMi).Error
	return
}

// GetWmsLogisticsStoragefeeMiInfoList 分页获取仓储费记录
// Author [piexlmax](https://github.com/piexlmax)
func (wmsLogisticsStoragefeeMiService *WmsLogisticsStoragefeeMiService)GetWmsLogisticsStoragefeeMiInfoList(info manualReq.WmsLogisticsStoragefeeMiSearch) (list []manual.WmsLogisticsStoragefeeMi, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.MustGetGlobalDBByDBName("ht_smartdata").Model(&manual.WmsLogisticsStoragefeeMi{})
    var wmsLogisticsStoragefeeMis []manual.WmsLogisticsStoragefeeMi
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
    if info.Area != "" {
        db = db.Where("area LIKE ?","%"+ info.Area+"%")
    }
    if info.CostType != "" {
        db = db.Where("cost_type LIKE ?","%"+ info.CostType+"%")
    }
    if info.Mmonth != "" {
        db = db.Where("mmonth LIKE ?","%"+ info.Mmonth+"%")
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
	
	err = db.Find(&wmsLogisticsStoragefeeMis).Error
	return  wmsLogisticsStoragefeeMis, total, err
}
