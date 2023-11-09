package manual

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/manual"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    manualReq "github.com/flipped-aurora/gin-vue-admin/server/model/manual/request"
    "gorm.io/gorm"
)

type PromotionMarketTargetService struct {
}

// CreatePromotionMarketTarget 创建大促推广目标表记录
// Author [piexlmax](https://github.com/piexlmax)
func (MarketTargetService *PromotionMarketTargetService) CreatePromotionMarketTarget(MarketTarget *manual.PromotionMarketTarget) (err error) {
	err = global.GVA_DB.Create(MarketTarget).Error
	return err
}

// DeletePromotionMarketTarget 删除大促推广目标表记录
// Author [piexlmax](https://github.com/piexlmax)
func (MarketTargetService *PromotionMarketTargetService)DeletePromotionMarketTarget(MarketTarget manual.PromotionMarketTarget) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&manual.PromotionMarketTarget{}).Where("id = ?", MarketTarget.ID).Update("deleted_by", MarketTarget.DeletedBy).Error; err != nil {
              return err
        }
        if err = tx.Delete(&MarketTarget).Error; err != nil {
              return err
        }
        return nil
	})
	return err
}

// DeletePromotionMarketTargetByIds 批量删除大促推广目标表记录
// Author [piexlmax](https://github.com/piexlmax)
func (MarketTargetService *PromotionMarketTargetService)DeletePromotionMarketTargetByIds(ids request.IdsReq,deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&manual.PromotionMarketTarget{}).Where("id in ?", ids.Ids).Update("deleted_by", deleted_by).Error; err != nil {
            return err
        }
        if err := tx.Where("id in ?", ids.Ids).Delete(&manual.PromotionMarketTarget{}).Error; err != nil {
            return err
        }
        return nil
    })
	return err
}

// UpdatePromotionMarketTarget 更新大促推广目标表记录
// Author [piexlmax](https://github.com/piexlmax)
func (MarketTargetService *PromotionMarketTargetService)UpdatePromotionMarketTarget(MarketTarget manual.PromotionMarketTarget) (err error) {
	err = global.GVA_DB.Save(&MarketTarget).Error
	return err
}

// GetPromotionMarketTarget 根据id获取大促推广目标表记录
// Author [piexlmax](https://github.com/piexlmax)
func (MarketTargetService *PromotionMarketTargetService)GetPromotionMarketTarget(id uint) (MarketTarget manual.PromotionMarketTarget, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&MarketTarget).Error
	return
}

// GetPromotionMarketTargetInfoList 分页获取大促推广目标表记录
// Author [piexlmax](https://github.com/piexlmax)
func (MarketTargetService *PromotionMarketTargetService)GetPromotionMarketTargetInfoList(info manualReq.PromotionMarketTargetSearch) (list []manual.PromotionMarketTarget, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&manual.PromotionMarketTarget{})
    var MarketTargets []manual.PromotionMarketTarget
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
    if info.ShopName != "" {
        db = db.Where("shop_name LIKE ?","%"+ info.ShopName+"%")
    }
    if info.ShopCode != "" {
        db = db.Where("shop_code LIKE ?","%"+ info.ShopCode+"%")
    }
        if info.StartTarget != nil && info.EndTarget != nil {
            db = db.Where("target BETWEEN ? AND ? ",info.StartTarget,info.EndTarget)
        }
    if info.Type != "" {
        db = db.Where("type LIKE ?","%"+ info.Type+"%")
    }
    if info.CreatedName != "" {
        db = db.Where("created_name = ?",info.CreatedName)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }
        var OrderStr string
        orderMap := make(map[string]bool)
         	orderMap["shop_name"] = true
         	orderMap["shop_code"] = true
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
	
	err = db.Find(&MarketTargets).Error
	return  MarketTargets, total, err
}
