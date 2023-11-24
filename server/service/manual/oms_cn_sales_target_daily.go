package manual

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/manual"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    manualReq "github.com/flipped-aurora/gin-vue-admin/server/model/manual/request"
    "gorm.io/gorm"
)

type OmsCnSalesTargetDailyService struct {
}

// CreateOmsCnSalesTargetDaily 创建京东自营销售录入表记录
// Author [piexlmax](https://github.com/piexlmax)
func (omsCnSalesTargetDailyService *OmsCnSalesTargetDailyService) CreateOmsCnSalesTargetDaily(omsCnSalesTargetDaily *manual.OmsCnSalesTargetDaily) (err error) {
	err = global.MustGetGlobalDBByDBName("ht_smartdata").Create(omsCnSalesTargetDaily).Error
	return err
}

// DeleteOmsCnSalesTargetDaily 删除京东自营销售录入表记录
// Author [piexlmax](https://github.com/piexlmax)
func (omsCnSalesTargetDailyService *OmsCnSalesTargetDailyService)DeleteOmsCnSalesTargetDaily(omsCnSalesTargetDaily manual.OmsCnSalesTargetDaily) (err error) {
	err = global.MustGetGlobalDBByDBName("ht_smartdata").Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&manual.OmsCnSalesTargetDaily{}).Where("id = ?", omsCnSalesTargetDaily.ID).Update("deleted_by", omsCnSalesTargetDaily.DeletedBy).Error; err != nil {
              return err
        }
        if err = tx.Delete(&omsCnSalesTargetDaily).Error; err != nil {
              return err
        }
        return nil
	})
	return err
}

// DeleteOmsCnSalesTargetDailyByIds 批量删除京东自营销售录入表记录
// Author [piexlmax](https://github.com/piexlmax)
func (omsCnSalesTargetDailyService *OmsCnSalesTargetDailyService)DeleteOmsCnSalesTargetDailyByIds(ids request.IdsReq,deleted_by uint) (err error) {
	err = global.MustGetGlobalDBByDBName("ht_smartdata").Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&manual.OmsCnSalesTargetDaily{}).Where("id in ?", ids.Ids).Update("deleted_by", deleted_by).Error; err != nil {
            return err
        }
        if err := tx.Where("id in ?", ids.Ids).Delete(&manual.OmsCnSalesTargetDaily{}).Error; err != nil {
            return err
        }
        return nil
    })
	return err
}

// UpdateOmsCnSalesTargetDaily 更新京东自营销售录入表记录
// Author [piexlmax](https://github.com/piexlmax)
func (omsCnSalesTargetDailyService *OmsCnSalesTargetDailyService)UpdateOmsCnSalesTargetDaily(omsCnSalesTargetDaily manual.OmsCnSalesTargetDaily) (err error) {
	err = global.MustGetGlobalDBByDBName("ht_smartdata").Save(&omsCnSalesTargetDaily).Error
	return err
}

// GetOmsCnSalesTargetDaily 根据id获取京东自营销售录入表记录
// Author [piexlmax](https://github.com/piexlmax)
func (omsCnSalesTargetDailyService *OmsCnSalesTargetDailyService)GetOmsCnSalesTargetDaily(id uint) (omsCnSalesTargetDaily manual.OmsCnSalesTargetDaily, err error) {
	err = global.MustGetGlobalDBByDBName("ht_smartdata").Where("id = ?", id).First(&omsCnSalesTargetDaily).Error
	return
}

// GetOmsCnSalesTargetDailyInfoList 分页获取京东自营销售录入表记录
// Author [piexlmax](https://github.com/piexlmax)
func (omsCnSalesTargetDailyService *OmsCnSalesTargetDailyService)GetOmsCnSalesTargetDailyInfoList(info manualReq.OmsCnSalesTargetDailySearch) (list []manual.OmsCnSalesTargetDaily, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.MustGetGlobalDBByDBName("ht_smartdata").Model(&manual.OmsCnSalesTargetDaily{})
    var omsCnSalesTargetDailys []manual.OmsCnSalesTargetDaily
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
    if info.Month != "" {
        db = db.Where("month LIKE ?","%"+ info.Month+"%")
    }
    if info.Channel != "" {
        db = db.Where("channel LIKE ?","%"+ info.Channel+"%")
    }
    if info.SheetName != "" {
        db = db.Where("sheet_name LIKE ?","%"+ info.SheetName+"%")
    }
    if info.CreatedName != "" {
        db = db.Where("created_name LIKE ?","%"+ info.CreatedName+"%")
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }
        var OrderStr string
        orderMap := make(map[string]bool)
         	orderMap["department"] = true
         	orderMap["month"] = true
         	orderMap["sheet_name"] = true
         	orderMap["created_name"] = true
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
	
	err = db.Find(&omsCnSalesTargetDailys).Error
	return  omsCnSalesTargetDailys, total, err
}
