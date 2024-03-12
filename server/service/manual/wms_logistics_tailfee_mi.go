package manual

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/manual"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    manualReq "github.com/flipped-aurora/gin-vue-admin/server/model/manual/request"
    "gorm.io/gorm"
)

type WmsLogisticsTailfeeMiService struct {
}

// CreateWmsLogisticsTailfeeMi 创建尾程费用结构记录
// Author [piexlmax](https://github.com/piexlmax)
func (wmsLogisticsTailfeeMiService *WmsLogisticsTailfeeMiService) CreateWmsLogisticsTailfeeMi(wmsLogisticsTailfeeMi *manual.WmsLogisticsTailfeeMi) (err error) {
	err = global.MustGetGlobalDBByDBName("ht_smartdata").Create(wmsLogisticsTailfeeMi).Error
	return err
}

// DeleteWmsLogisticsTailfeeMi 删除尾程费用结构记录
// Author [piexlmax](https://github.com/piexlmax)
func (wmsLogisticsTailfeeMiService *WmsLogisticsTailfeeMiService)DeleteWmsLogisticsTailfeeMi(wmsLogisticsTailfeeMi manual.WmsLogisticsTailfeeMi) (err error) {
	err = global.MustGetGlobalDBByDBName("ht_smartdata").Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&manual.WmsLogisticsTailfeeMi{}).Where("id = ?", wmsLogisticsTailfeeMi.ID).Update("deleted_by", wmsLogisticsTailfeeMi.DeletedBy).Error; err != nil {
              return err
        }
        if err = tx.Delete(&wmsLogisticsTailfeeMi).Error; err != nil {
              return err
        }
        return nil
	})
	return err
}

// DeleteWmsLogisticsTailfeeMiByIds 批量删除尾程费用结构记录
// Author [piexlmax](https://github.com/piexlmax)
func (wmsLogisticsTailfeeMiService *WmsLogisticsTailfeeMiService)DeleteWmsLogisticsTailfeeMiByIds(ids request.IdsReq,deleted_by uint) (err error) {
	err = global.MustGetGlobalDBByDBName("ht_smartdata").Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&manual.WmsLogisticsTailfeeMi{}).Where("id in ?", ids.Ids).Update("deleted_by", deleted_by).Error; err != nil {
            return err
        }
        if err := tx.Where("id in ?", ids.Ids).Delete(&manual.WmsLogisticsTailfeeMi{}).Error; err != nil {
            return err
        }
        return nil
    })
	return err
}

// UpdateWmsLogisticsTailfeeMi 更新尾程费用结构记录
// Author [piexlmax](https://github.com/piexlmax)
func (wmsLogisticsTailfeeMiService *WmsLogisticsTailfeeMiService)UpdateWmsLogisticsTailfeeMi(wmsLogisticsTailfeeMi manual.WmsLogisticsTailfeeMi) (err error) {
	err = global.MustGetGlobalDBByDBName("ht_smartdata").Save(&wmsLogisticsTailfeeMi).Error
	return err
}

// GetWmsLogisticsTailfeeMi 根据id获取尾程费用结构记录
// Author [piexlmax](https://github.com/piexlmax)
func (wmsLogisticsTailfeeMiService *WmsLogisticsTailfeeMiService)GetWmsLogisticsTailfeeMi(id uint) (wmsLogisticsTailfeeMi manual.WmsLogisticsTailfeeMi, err error) {
	err = global.MustGetGlobalDBByDBName("ht_smartdata").Where("id = ?", id).First(&wmsLogisticsTailfeeMi).Error
	return
}

// GetWmsLogisticsTailfeeMiInfoList 分页获取尾程费用结构记录
// Author [piexlmax](https://github.com/piexlmax)
func (wmsLogisticsTailfeeMiService *WmsLogisticsTailfeeMiService)GetWmsLogisticsTailfeeMiInfoList(info manualReq.WmsLogisticsTailfeeMiSearch) (list []manual.WmsLogisticsTailfeeMi, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.MustGetGlobalDBByDBName("ht_smartdata").Model(&manual.WmsLogisticsTailfeeMi{})
    var wmsLogisticsTailfeeMis []manual.WmsLogisticsTailfeeMi
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
	
	err = db.Find(&wmsLogisticsTailfeeMis).Error
	return  wmsLogisticsTailfeeMis, total, err
}
