package manual

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/manual"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    manualReq "github.com/flipped-aurora/gin-vue-admin/server/model/manual/request"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/service"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
    "github.com/flipped-aurora/gin-vue-admin/server/utils"
)

type OmsCnChannelSalesApi struct {
}

var omsCnChannelSalesService = service.ServiceGroupApp.ManualServiceGroup.OmsCnChannelSalesService


// CreateOmsCnChannelSales 创建按天销售目标录入表
// @Tags OmsCnChannelSales
// @Summary 创建按天销售目标录入表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body manual.OmsCnChannelSales true "创建按天销售目标录入表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /omsCnChannelSales/createOmsCnChannelSales [post]
func (omsCnChannelSalesApi *OmsCnChannelSalesApi) CreateOmsCnChannelSales(c *gin.Context) {
	var omsCnChannelSales manual.OmsCnChannelSales
	err := c.ShouldBindJSON(&omsCnChannelSales)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	// =============== 新增 ==================
	if omsCnChannelSales.CreatedName == "" {
		omsCnChannelSales.CreatedName = utils.GetUserName(c)
	}
	// =============== 新增 ==================
    omsCnChannelSales.CreatedBy = utils.GetUserID(c)
	if err := omsCnChannelSalesService.CreateOmsCnChannelSales(&omsCnChannelSales); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteOmsCnChannelSales 删除按天销售目标录入表
// @Tags OmsCnChannelSales
// @Summary 删除按天销售目标录入表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body manual.OmsCnChannelSales true "删除按天销售目标录入表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /omsCnChannelSales/deleteOmsCnChannelSales [delete]
func (omsCnChannelSalesApi *OmsCnChannelSalesApi) DeleteOmsCnChannelSales(c *gin.Context) {
	var omsCnChannelSales manual.OmsCnChannelSales
	err := c.ShouldBindJSON(&omsCnChannelSales)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    omsCnChannelSales.DeletedBy = utils.GetUserID(c)
	if err := omsCnChannelSalesService.DeleteOmsCnChannelSales(omsCnChannelSales); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteOmsCnChannelSalesAll 删除所有按天销售目标录入表
// @Tags OmsCnChannelSalesAll
// @Summary 删除所有按天销售目标录入表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body manual.OmsCnChannelSales true "删除所有按天销售目标录入表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /omsCnChannelSales/deleteOmsCnChannelSalesAll [delete]
func (omsCnChannelSalesApi *OmsCnChannelSalesApi) DeleteOmsCnChannelSalesAll(c *gin.Context) {
	var omsCnChannelSales manual.OmsCnChannelSales
	
	if err := global.MustGetGlobalDBByDBName("ht_smartdata").Unscoped().Where("1 = 1").Delete(&omsCnChannelSales).Error; err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteOmsCnChannelSalesByIds 批量删除按天销售目标录入表
// @Tags OmsCnChannelSales
// @Summary 批量删除按天销售目标录入表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除按天销售目标录入表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /omsCnChannelSales/deleteOmsCnChannelSalesByIds [delete]
func (omsCnChannelSalesApi *OmsCnChannelSalesApi) DeleteOmsCnChannelSalesByIds(c *gin.Context) {
	var IDS request.IdsReq
    err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    deletedBy := utils.GetUserID(c)
	if err := omsCnChannelSalesService.DeleteOmsCnChannelSalesByIds(IDS,deletedBy); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateOmsCnChannelSales 更新按天销售目标录入表
// @Tags OmsCnChannelSales
// @Summary 更新按天销售目标录入表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body manual.OmsCnChannelSales true "更新按天销售目标录入表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /omsCnChannelSales/updateOmsCnChannelSales [put]
func (omsCnChannelSalesApi *OmsCnChannelSalesApi) UpdateOmsCnChannelSales(c *gin.Context) {
	var omsCnChannelSales manual.OmsCnChannelSales
	err := c.ShouldBindJSON(&omsCnChannelSales)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    omsCnChannelSales.UpdatedBy = utils.GetUserID(c)
	if err := omsCnChannelSalesService.UpdateOmsCnChannelSales(omsCnChannelSales); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindOmsCnChannelSales 用id查询按天销售目标录入表
// @Tags OmsCnChannelSales
// @Summary 用id查询按天销售目标录入表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query manual.OmsCnChannelSales true "用id查询按天销售目标录入表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /omsCnChannelSales/findOmsCnChannelSales [get]
func (omsCnChannelSalesApi *OmsCnChannelSalesApi) FindOmsCnChannelSales(c *gin.Context) {
	var omsCnChannelSales manual.OmsCnChannelSales
	err := c.ShouldBindQuery(&omsCnChannelSales)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if reomsCnChannelSales, err := omsCnChannelSalesService.GetOmsCnChannelSales(omsCnChannelSales.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reomsCnChannelSales": reomsCnChannelSales}, c)
	}
}

// GetOmsCnChannelSalesList 分页获取按天销售目标录入表列表
// @Tags OmsCnChannelSales
// @Summary 分页获取按天销售目标录入表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query manualReq.OmsCnChannelSalesSearch true "分页获取按天销售目标录入表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /omsCnChannelSales/getOmsCnChannelSalesList [get]
func (omsCnChannelSalesApi *OmsCnChannelSalesApi) GetOmsCnChannelSalesList(c *gin.Context) {
	var pageInfo manualReq.OmsCnChannelSalesSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := omsCnChannelSalesService.GetOmsCnChannelSalesInfoList(pageInfo); err != nil {
	    global.GVA_LOG.Error("获取失败!", zap.Error(err))
        response.FailWithMessage("获取失败", c)
    } else {
        response.OkWithDetailed(response.PageResult{
            List:     list,
            Total:    total,
            Page:     pageInfo.Page,
            PageSize: pageInfo.PageSize,
        }, "获取成功", c)
    }
}
