package manual

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/manual"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    manualReq "github.com/flipped-aurora/gin-vue-admin/server/model/manual/request"
    "gorm.io/gorm"
)

type WmsLogisticsModeTransportMiService struct {
}

// CreateWmsLogisticsModeTransportMi 创建费用by运输方式记录
// Author [piexlmax](https://github.com/piexlmax)
func (wmsLogisticsModeTransportMiService *WmsLogisticsModeTransportMiService) CreateWmsLogisticsModeTransportMi(wmsLogisticsModeTransportMi *manual.WmsLogisticsModeTransportMi) (err error) {
	err = global.MustGetGlobalDBByDBName("ht_smartdata").Create(wmsLogisticsModeTransportMi).Error
	return err
}

// DeleteWmsLogisticsModeTransportMi 删除费用by运输方式记录
// Author [piexlmax](https://github.com/piexlmax)
func (wmsLogisticsModeTransportMiService *WmsLogisticsModeTransportMiService)DeleteWmsLogisticsModeTransportMi(wmsLogisticsModeTransportMi manual.WmsLogisticsModeTransportMi) (err error) {
	err = global.MustGetGlobalDBByDBName("ht_smartdata").Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&manual.WmsLogisticsModeTransportMi{}).Where("id = ?", wmsLogisticsModeTransportMi.ID).Update("deleted_by", wmsLogisticsModeTransportMi.DeletedBy).Error; err != nil {
              return err
        }
        if err = tx.Delete(&wmsLogisticsModeTransportMi).Error; err != nil {
              return err
        }
        return nil
	})
	return err
}

// DeleteWmsLogisticsModeTransportMiByIds 批量删除费用by运输方式记录
// Author [piexlmax](https://github.com/piexlmax)
func (wmsLogisticsModeTransportMiService *WmsLogisticsModeTransportMiService)DeleteWmsLogisticsModeTransportMiByIds(ids request.IdsReq,deleted_by uint) (err error) {
	err = global.MustGetGlobalDBByDBName("ht_smartdata").Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&manual.WmsLogisticsModeTransportMi{}).Where("id in ?", ids.Ids).Update("deleted_by", deleted_by).Error; err != nil {
            return err
        }
        if err := tx.Where("id in ?", ids.Ids).Delete(&manual.WmsLogisticsModeTransportMi{}).Error; err != nil {
            return err
        }
        return nil
    })
	return err
}

// UpdateWmsLogisticsModeTransportMi 更新费用by运输方式记录
// Author [piexlmax](https://github.com/piexlmax)
func (wmsLogisticsModeTransportMiService *WmsLogisticsModeTransportMiService)UpdateWmsLogisticsModeTransportMi(wmsLogisticsModeTransportMi manual.WmsLogisticsModeTransportMi) (err error) {
	err = global.MustGetGlobalDBByDBName("ht_smartdata").Save(&wmsLogisticsModeTransportMi).Error
	return err
}

// GetWmsLogisticsModeTransportMi 根据id获取费用by运输方式记录
// Author [piexlmax](https://github.com/piexlmax)
func (wmsLogisticsModeTransportMiService *WmsLogisticsModeTransportMiService)GetWmsLogisticsModeTransportMi(id uint) (wmsLogisticsModeTransportMi manual.WmsLogisticsModeTransportMi, err error) {
	err = global.MustGetGlobalDBByDBName("ht_smartdata").Where("id = ?", id).First(&wmsLogisticsModeTransportMi).Error
	return
}

// GetWmsLogisticsModeTransportMiInfoList 分页获取费用by运输方式记录
// Author [piexlmax](https://github.com/piexlmax)
func (wmsLogisticsModeTransportMiService *WmsLogisticsModeTransportMiService)GetWmsLogisticsModeTransportMiInfoList(info manualReq.WmsLogisticsModeTransportMiSearch) (list []manual.WmsLogisticsModeTransportMi, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.MustGetGlobalDBByDBName("ht_smartdata").Model(&manual.WmsLogisticsModeTransportMi{})
    var wmsLogisticsModeTransportMis []manual.WmsLogisticsModeTransportMi
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
    if info.Area != "" {
        db = db.Where("area LIKE ?","%"+ info.Area+"%")
    }
    if info.TransMethod != "" {
        db = db.Where("trans_method LIKE ?","%"+ info.TransMethod+"%")
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
	
	err = db.Find(&wmsLogisticsModeTransportMis).Error
	return  wmsLogisticsModeTransportMis, total, err
}
