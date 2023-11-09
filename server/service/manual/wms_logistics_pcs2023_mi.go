package manual

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/manual"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    manualReq "github.com/flipped-aurora/gin-vue-admin/server/model/manual/request"
    "gorm.io/gorm"
)

type WmsLogisticsPcs2023MiService struct {
}

// CreateWmsLogisticsPcs2023Mi 创建2023年物流成本明细表记录
// Author [piexlmax](https://github.com/piexlmax)
func (wmsLogisticsPcs2023MiService *WmsLogisticsPcs2023MiService) CreateWmsLogisticsPcs2023Mi(wmsLogisticsPcs2023Mi *manual.WmsLogisticsPcs2023Mi) (err error) {
	err = global.MustGetGlobalDBByDBName("ht_smartdata").Create(wmsLogisticsPcs2023Mi).Error
	return err
}

// DeleteWmsLogisticsPcs2023Mi 删除2023年物流成本明细表记录
// Author [piexlmax](https://github.com/piexlmax)
func (wmsLogisticsPcs2023MiService *WmsLogisticsPcs2023MiService)DeleteWmsLogisticsPcs2023Mi(wmsLogisticsPcs2023Mi manual.WmsLogisticsPcs2023Mi) (err error) {
	err = global.MustGetGlobalDBByDBName("ht_smartdata").Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&manual.WmsLogisticsPcs2023Mi{}).Where("id = ?", wmsLogisticsPcs2023Mi.ID).Update("deleted_by", wmsLogisticsPcs2023Mi.DeletedBy).Error; err != nil {
              return err
        }
        if err = tx.Delete(&wmsLogisticsPcs2023Mi).Error; err != nil {
              return err
        }
        return nil
	})
	return err
}

// DeleteWmsLogisticsPcs2023MiByIds 批量删除2023年物流成本明细表记录
// Author [piexlmax](https://github.com/piexlmax)
func (wmsLogisticsPcs2023MiService *WmsLogisticsPcs2023MiService)DeleteWmsLogisticsPcs2023MiByIds(ids request.IdsReq,deleted_by uint) (err error) {
	err = global.MustGetGlobalDBByDBName("ht_smartdata").Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&manual.WmsLogisticsPcs2023Mi{}).Where("id in ?", ids.Ids).Update("deleted_by", deleted_by).Error; err != nil {
            return err
        }
        if err := tx.Where("id in ?", ids.Ids).Delete(&manual.WmsLogisticsPcs2023Mi{}).Error; err != nil {
            return err
        }
        return nil
    })
	return err
}

// UpdateWmsLogisticsPcs2023Mi 更新2023年物流成本明细表记录
// Author [piexlmax](https://github.com/piexlmax)
func (wmsLogisticsPcs2023MiService *WmsLogisticsPcs2023MiService)UpdateWmsLogisticsPcs2023Mi(wmsLogisticsPcs2023Mi manual.WmsLogisticsPcs2023Mi) (err error) {
	err = global.MustGetGlobalDBByDBName("ht_smartdata").Save(&wmsLogisticsPcs2023Mi).Error
	return err
}

// GetWmsLogisticsPcs2023Mi 根据id获取2023年物流成本明细表记录
// Author [piexlmax](https://github.com/piexlmax)
func (wmsLogisticsPcs2023MiService *WmsLogisticsPcs2023MiService)GetWmsLogisticsPcs2023Mi(id uint) (wmsLogisticsPcs2023Mi manual.WmsLogisticsPcs2023Mi, err error) {
	err = global.MustGetGlobalDBByDBName("ht_smartdata").Where("id = ?", id).First(&wmsLogisticsPcs2023Mi).Error
	return
}

// GetWmsLogisticsPcs2023MiInfoList 分页获取2023年物流成本明细表记录
// Author [piexlmax](https://github.com/piexlmax)
func (wmsLogisticsPcs2023MiService *WmsLogisticsPcs2023MiService)GetWmsLogisticsPcs2023MiInfoList(info manualReq.WmsLogisticsPcs2023MiSearch) (list []manual.WmsLogisticsPcs2023Mi, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.MustGetGlobalDBByDBName("ht_smartdata").Model(&manual.WmsLogisticsPcs2023Mi{})
    var wmsLogisticsPcs2023Mis []manual.WmsLogisticsPcs2023Mi
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
    if info.Country != "" {
        db = db.Where("country LIKE ?","%"+ info.Country+"%")
    }
    if info.ProductAbbreviation != "" {
        db = db.Where("product_abbreviation LIKE ?","%"+ info.ProductAbbreviation+"%")
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
         	orderMap["country"] = true
         	orderMap["product_abbreviation"] = true
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
	
	err = db.Find(&wmsLogisticsPcs2023Mis).Error
	return  wmsLogisticsPcs2023Mis, total, err
}
