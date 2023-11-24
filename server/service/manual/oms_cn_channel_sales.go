package manual

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/manual"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    manualReq "github.com/flipped-aurora/gin-vue-admin/server/model/manual/request"
    "gorm.io/gorm"
)

type OmsCnChannelSalesService struct {
}

// CreateOmsCnChannelSales 创建按天销售目标录入表记录
// Author [piexlmax](https://github.com/piexlmax)
func (omsCnChannelSalesService *OmsCnChannelSalesService) CreateOmsCnChannelSales(omsCnChannelSales *manual.OmsCnChannelSales) (err error) {
	err = global.MustGetGlobalDBByDBName("ht_smartdata").Create(omsCnChannelSales).Error
	return err
}

// DeleteOmsCnChannelSales 删除按天销售目标录入表记录
// Author [piexlmax](https://github.com/piexlmax)
func (omsCnChannelSalesService *OmsCnChannelSalesService)DeleteOmsCnChannelSales(omsCnChannelSales manual.OmsCnChannelSales) (err error) {
	err = global.MustGetGlobalDBByDBName("ht_smartdata").Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&manual.OmsCnChannelSales{}).Where("id = ?", omsCnChannelSales.ID).Update("deleted_by", omsCnChannelSales.DeletedBy).Error; err != nil {
              return err
        }
        if err = tx.Delete(&omsCnChannelSales).Error; err != nil {
              return err
        }
        return nil
	})
	return err
}

// DeleteOmsCnChannelSalesByIds 批量删除按天销售目标录入表记录
// Author [piexlmax](https://github.com/piexlmax)
func (omsCnChannelSalesService *OmsCnChannelSalesService)DeleteOmsCnChannelSalesByIds(ids request.IdsReq,deleted_by uint) (err error) {
	err = global.MustGetGlobalDBByDBName("ht_smartdata").Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&manual.OmsCnChannelSales{}).Where("id in ?", ids.Ids).Update("deleted_by", deleted_by).Error; err != nil {
            return err
        }
        if err := tx.Where("id in ?", ids.Ids).Delete(&manual.OmsCnChannelSales{}).Error; err != nil {
            return err
        }
        return nil
    })
	return err
}

// UpdateOmsCnChannelSales 更新按天销售目标录入表记录
// Author [piexlmax](https://github.com/piexlmax)
func (omsCnChannelSalesService *OmsCnChannelSalesService)UpdateOmsCnChannelSales(omsCnChannelSales manual.OmsCnChannelSales) (err error) {
	err = global.MustGetGlobalDBByDBName("ht_smartdata").Save(&omsCnChannelSales).Error
	return err
}

// GetOmsCnChannelSales 根据id获取按天销售目标录入表记录
// Author [piexlmax](https://github.com/piexlmax)
func (omsCnChannelSalesService *OmsCnChannelSalesService)GetOmsCnChannelSales(id uint) (omsCnChannelSales manual.OmsCnChannelSales, err error) {
	err = global.MustGetGlobalDBByDBName("ht_smartdata").Where("id = ?", id).First(&omsCnChannelSales).Error
	return
}

// GetOmsCnChannelSalesInfoList 分页获取按天销售目标录入表记录
// Author [piexlmax](https://github.com/piexlmax)
func (omsCnChannelSalesService *OmsCnChannelSalesService)GetOmsCnChannelSalesInfoList(info manualReq.OmsCnChannelSalesSearch) (list []manual.OmsCnChannelSales, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.MustGetGlobalDBByDBName("ht_smartdata").Model(&manual.OmsCnChannelSales{})
    var omsCnChannelSaless []manual.OmsCnChannelSales
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
    if info.Date != "" {
        db = db.Where("date LIKE ?","%"+ info.Date+"%")
    }
    if info.Channel != "" {
        db = db.Where("channel LIKE ?","%"+ info.Channel+"%")
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
         	orderMap["date"] = true
         	orderMap["department"] = true
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
	
	err = db.Find(&omsCnChannelSaless).Error
	return  omsCnChannelSaless, total, err
}
