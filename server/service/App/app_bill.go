package App

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/App"
	AppReq "github.com/flipped-aurora/gin-vue-admin/server/model/App/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type AppBillService struct {
}

// CreateAppBill 创建AppBill记录
// Author [piexlmax](https://github.com/piexlmax)
func (appBillService *AppBillService) CreateAppBill(appBill App.AppBill) (err error) {
	err = global.GVA_DB.Create(&appBill).Error
	return err
}

// DeleteAppBill 删除AppBill记录
// Author [piexlmax](https://github.com/piexlmax)
func (appBillService *AppBillService) DeleteAppBill(appBill App.AppBill) (err error) {
	err = global.GVA_DB.Delete(&appBill).Error
	return err
}

// DeleteAppBillByIds 批量删除AppBill记录
// Author [piexlmax](https://github.com/piexlmax)
func (appBillService *AppBillService) DeleteAppBillByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]App.AppBill{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateAppBill 更新AppBill记录
// Author [piexlmax](https://github.com/piexlmax)
func (appBillService *AppBillService) UpdateAppBill(appBill App.AppBill) (err error) {
	err = global.GVA_DB.Save(&appBill).Error
	return err
}

// GetAppBill 根据id获取AppBill记录
// Author [piexlmax](https://github.com/piexlmax)
func (appBillService *AppBillService) GetAppBill(id uint) (appBill App.AppBill, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&appBill).Error
	return
}

// GetAppBillInfoList 分页获取AppBill记录
// Author [piexlmax](https://github.com/piexlmax)
func (appBillService *AppBillService) GetAppBillInfoList(info AppReq.AppBillSearch) (list []App.AppBill, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&App.AppBill{})
	var appBills []App.AppBill
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.ID != 0 {
		db = db.Where("user_id = ?", info.ID)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Order("created_at desc").Limit(limit).Offset(offset).Find(&appBills).Error
	return appBills, total, err
}
