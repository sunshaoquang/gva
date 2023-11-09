package manual

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/manual"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    manualReq "github.com/flipped-aurora/gin-vue-admin/server/model/manual/request"
    "gorm.io/gorm"
)

type WmsKpiSummary2MiService struct {
}

// CreateWmsKpiSummary2Mi 创建wmsKpiSummary2Mi表记录
// Author [piexlmax](https://github.com/piexlmax)
func (wmsKpiSummary2MiService *WmsKpiSummary2MiService) CreateWmsKpiSummary2Mi(wmsKpiSummary2Mi *manual.WmsKpiSummary2Mi) (err error) {
	err = global.MustGetGlobalDBByDBName("ht_smartdata").Create(wmsKpiSummary2Mi).Error
	return err
}

// DeleteWmsKpiSummary2Mi 删除wmsKpiSummary2Mi表记录
// Author [piexlmax](https://github.com/piexlmax)
func (wmsKpiSummary2MiService *WmsKpiSummary2MiService)DeleteWmsKpiSummary2Mi(wmsKpiSummary2Mi manual.WmsKpiSummary2Mi) (err error) {
	err = global.MustGetGlobalDBByDBName("ht_smartdata").Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&manual.WmsKpiSummary2Mi{}).Where("id = ?", wmsKpiSummary2Mi.ID).Update("deleted_by", wmsKpiSummary2Mi.DeletedBy).Error; err != nil {
              return err
        }
        if err = tx.Delete(&wmsKpiSummary2Mi).Error; err != nil {
              return err
        }
        return nil
	})
	return err
}

// DeleteWmsKpiSummary2MiByIds 批量删除wmsKpiSummary2Mi表记录
// Author [piexlmax](https://github.com/piexlmax)
func (wmsKpiSummary2MiService *WmsKpiSummary2MiService)DeleteWmsKpiSummary2MiByIds(ids request.IdsReq,deleted_by uint) (err error) {
	err = global.MustGetGlobalDBByDBName("ht_smartdata").Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&manual.WmsKpiSummary2Mi{}).Where("id in ?", ids.Ids).Update("deleted_by", deleted_by).Error; err != nil {
            return err
        }
        if err := tx.Where("id in ?", ids.Ids).Delete(&manual.WmsKpiSummary2Mi{}).Error; err != nil {
            return err
        }
        return nil
    })
	return err
}

// UpdateWmsKpiSummary2Mi 更新wmsKpiSummary2Mi表记录
// Author [piexlmax](https://github.com/piexlmax)
func (wmsKpiSummary2MiService *WmsKpiSummary2MiService)UpdateWmsKpiSummary2Mi(wmsKpiSummary2Mi manual.WmsKpiSummary2Mi) (err error) {
	err = global.MustGetGlobalDBByDBName("ht_smartdata").Save(&wmsKpiSummary2Mi).Error
	return err
}

// GetWmsKpiSummary2Mi 根据id获取wmsKpiSummary2Mi表记录
// Author [piexlmax](https://github.com/piexlmax)
func (wmsKpiSummary2MiService *WmsKpiSummary2MiService)GetWmsKpiSummary2Mi(id uint) (wmsKpiSummary2Mi manual.WmsKpiSummary2Mi, err error) {
	err = global.MustGetGlobalDBByDBName("ht_smartdata").Where("id = ?", id).First(&wmsKpiSummary2Mi).Error
	return
}

// GetWmsKpiSummary2MiInfoList 分页获取wmsKpiSummary2Mi表记录
// Author [piexlmax](https://github.com/piexlmax)
func (wmsKpiSummary2MiService *WmsKpiSummary2MiService)GetWmsKpiSummary2MiInfoList(info manualReq.WmsKpiSummary2MiSearch) (list []manual.WmsKpiSummary2Mi, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.MustGetGlobalDBByDBName("ht_smartdata").Model(&manual.WmsKpiSummary2Mi{})
    var wmsKpiSummary2Mis []manual.WmsKpiSummary2Mi
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
    if info.Project != "" {
        db = db.Where("project LIKE ?","%"+ info.Project+"%")
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
         	orderMap["project"] = true
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
	
	err = db.Find(&wmsKpiSummary2Mis).Error
	return  wmsKpiSummary2Mis, total, err
}
