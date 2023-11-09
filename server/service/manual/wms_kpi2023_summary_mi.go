package manual

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/manual"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    manualReq "github.com/flipped-aurora/gin-vue-admin/server/model/manual/request"
    "gorm.io/gorm"
)

type WmsKpi2023SummaryMiService struct {
}

// CreateWmsKpi2023SummaryMi 创建wmsKpi2023SummaryMi表记录
// Author [piexlmax](https://github.com/piexlmax)
func (wmsKpi2023SummaryMiService *WmsKpi2023SummaryMiService) CreateWmsKpi2023SummaryMi(wmsKpi2023SummaryMi *manual.WmsKpi2023SummaryMi) (err error) {
	err = global.MustGetGlobalDBByDBName("ht_smartdata").Create(wmsKpi2023SummaryMi).Error
	return err
}

// DeleteWmsKpi2023SummaryMi 删除wmsKpi2023SummaryMi表记录
// Author [piexlmax](https://github.com/piexlmax)
func (wmsKpi2023SummaryMiService *WmsKpi2023SummaryMiService)DeleteWmsKpi2023SummaryMi(wmsKpi2023SummaryMi manual.WmsKpi2023SummaryMi) (err error) {
	err = global.MustGetGlobalDBByDBName("ht_smartdata").Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&manual.WmsKpi2023SummaryMi{}).Where("id = ?", wmsKpi2023SummaryMi.ID).Update("deleted_by", wmsKpi2023SummaryMi.DeletedBy).Error; err != nil {
              return err
        }
        if err = tx.Delete(&wmsKpi2023SummaryMi).Error; err != nil {
              return err
        }
        return nil
	})
	return err
}

// DeleteWmsKpi2023SummaryMiByIds 批量删除wmsKpi2023SummaryMi表记录
// Author [piexlmax](https://github.com/piexlmax)
func (wmsKpi2023SummaryMiService *WmsKpi2023SummaryMiService)DeleteWmsKpi2023SummaryMiByIds(ids request.IdsReq,deleted_by uint) (err error) {
	err = global.MustGetGlobalDBByDBName("ht_smartdata").Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&manual.WmsKpi2023SummaryMi{}).Where("id in ?", ids.Ids).Update("deleted_by", deleted_by).Error; err != nil {
            return err
        }
        if err := tx.Where("id in ?", ids.Ids).Delete(&manual.WmsKpi2023SummaryMi{}).Error; err != nil {
            return err
        }
        return nil
    })
	return err
}

// UpdateWmsKpi2023SummaryMi 更新wmsKpi2023SummaryMi表记录
// Author [piexlmax](https://github.com/piexlmax)
func (wmsKpi2023SummaryMiService *WmsKpi2023SummaryMiService)UpdateWmsKpi2023SummaryMi(wmsKpi2023SummaryMi manual.WmsKpi2023SummaryMi) (err error) {
	err = global.MustGetGlobalDBByDBName("ht_smartdata").Save(&wmsKpi2023SummaryMi).Error
	return err
}

// GetWmsKpi2023SummaryMi 根据id获取wmsKpi2023SummaryMi表记录
// Author [piexlmax](https://github.com/piexlmax)
func (wmsKpi2023SummaryMiService *WmsKpi2023SummaryMiService)GetWmsKpi2023SummaryMi(id uint) (wmsKpi2023SummaryMi manual.WmsKpi2023SummaryMi, err error) {
	err = global.MustGetGlobalDBByDBName("ht_smartdata").Where("id = ?", id).First(&wmsKpi2023SummaryMi).Error
	return
}

// GetWmsKpi2023SummaryMiInfoList 分页获取wmsKpi2023SummaryMi表记录
// Author [piexlmax](https://github.com/piexlmax)
func (wmsKpi2023SummaryMiService *WmsKpi2023SummaryMiService)GetWmsKpi2023SummaryMiInfoList(info manualReq.WmsKpi2023SummaryMiSearch) (list []manual.WmsKpi2023SummaryMi, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.MustGetGlobalDBByDBName("ht_smartdata").Model(&manual.WmsKpi2023SummaryMi{})
    var wmsKpi2023SummaryMis []manual.WmsKpi2023SummaryMi
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
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
         	orderMap["indicator_classification"] = true
         	orderMap["indicator_name"] = true
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
	
	err = db.Find(&wmsKpi2023SummaryMis).Error
	return  wmsKpi2023SummaryMis, total, err
}
