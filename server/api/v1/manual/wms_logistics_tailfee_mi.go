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

type WmsLogisticsTailfeeMiApi struct {
}

var wmsLogisticsTailfeeMiService = service.ServiceGroupApp.ManualServiceGroup.WmsLogisticsTailfeeMiService


// CreateWmsLogisticsTailfeeMi 创建尾程费用结构
// @Tags WmsLogisticsTailfeeMi
// @Summary 创建尾程费用结构
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body manual.WmsLogisticsTailfeeMi true "创建尾程费用结构"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /wmsLogisticsTailfeeMi/createWmsLogisticsTailfeeMi [post]
func (wmsLogisticsTailfeeMiApi *WmsLogisticsTailfeeMiApi) CreateWmsLogisticsTailfeeMi(c *gin.Context) {
	var wmsLogisticsTailfeeMi manual.WmsLogisticsTailfeeMi
	err := c.ShouldBindJSON(&wmsLogisticsTailfeeMi)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	// =============== 新增 ==================
	if wmsLogisticsTailfeeMi.CreatedName == "" {
		wmsLogisticsTailfeeMi.CreatedName = utils.GetUserName(c)
	}
	// =============== 新增 ==================
    wmsLogisticsTailfeeMi.CreatedBy = utils.GetUserID(c)
	if err := wmsLogisticsTailfeeMiService.CreateWmsLogisticsTailfeeMi(&wmsLogisticsTailfeeMi); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteWmsLogisticsTailfeeMi 删除尾程费用结构
// @Tags WmsLogisticsTailfeeMi
// @Summary 删除尾程费用结构
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body manual.WmsLogisticsTailfeeMi true "删除尾程费用结构"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /wmsLogisticsTailfeeMi/deleteWmsLogisticsTailfeeMi [delete]
func (wmsLogisticsTailfeeMiApi *WmsLogisticsTailfeeMiApi) DeleteWmsLogisticsTailfeeMi(c *gin.Context) {
	var wmsLogisticsTailfeeMi manual.WmsLogisticsTailfeeMi
	err := c.ShouldBindJSON(&wmsLogisticsTailfeeMi)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    wmsLogisticsTailfeeMi.DeletedBy = utils.GetUserID(c)
	if err := wmsLogisticsTailfeeMiService.DeleteWmsLogisticsTailfeeMi(wmsLogisticsTailfeeMi); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteWmsLogisticsTailfeeMiAll 删除所有尾程费用结构
// @Tags WmsLogisticsTailfeeMiAll
// @Summary 删除所有尾程费用结构
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body manual.WmsLogisticsTailfeeMi true "删除所有尾程费用结构"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /wmsLogisticsTailfeeMi/deleteWmsLogisticsTailfeeMiAll [delete]
func (wmsLogisticsTailfeeMiApi *WmsLogisticsTailfeeMiApi) DeleteWmsLogisticsTailfeeMiAll(c *gin.Context) {
	var wmsLogisticsTailfeeMi manual.WmsLogisticsTailfeeMi
	
	if err := global.MustGetGlobalDBByDBName("ht_smartdata").Unscoped().Where("1 = 1").Delete(&wmsLogisticsTailfeeMi).Error; err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteWmsLogisticsTailfeeMiByIds 批量删除尾程费用结构
// @Tags WmsLogisticsTailfeeMi
// @Summary 批量删除尾程费用结构
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除尾程费用结构"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /wmsLogisticsTailfeeMi/deleteWmsLogisticsTailfeeMiByIds [delete]
func (wmsLogisticsTailfeeMiApi *WmsLogisticsTailfeeMiApi) DeleteWmsLogisticsTailfeeMiByIds(c *gin.Context) {
	var IDS request.IdsReq
    err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    deletedBy := utils.GetUserID(c)
	if err := wmsLogisticsTailfeeMiService.DeleteWmsLogisticsTailfeeMiByIds(IDS,deletedBy); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateWmsLogisticsTailfeeMi 更新尾程费用结构
// @Tags WmsLogisticsTailfeeMi
// @Summary 更新尾程费用结构
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body manual.WmsLogisticsTailfeeMi true "更新尾程费用结构"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /wmsLogisticsTailfeeMi/updateWmsLogisticsTailfeeMi [put]
func (wmsLogisticsTailfeeMiApi *WmsLogisticsTailfeeMiApi) UpdateWmsLogisticsTailfeeMi(c *gin.Context) {
	var wmsLogisticsTailfeeMi manual.WmsLogisticsTailfeeMi
	err := c.ShouldBindJSON(&wmsLogisticsTailfeeMi)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    wmsLogisticsTailfeeMi.UpdatedBy = utils.GetUserID(c)
	if err := wmsLogisticsTailfeeMiService.UpdateWmsLogisticsTailfeeMi(wmsLogisticsTailfeeMi); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindWmsLogisticsTailfeeMi 用id查询尾程费用结构
// @Tags WmsLogisticsTailfeeMi
// @Summary 用id查询尾程费用结构
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query manual.WmsLogisticsTailfeeMi true "用id查询尾程费用结构"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /wmsLogisticsTailfeeMi/findWmsLogisticsTailfeeMi [get]
func (wmsLogisticsTailfeeMiApi *WmsLogisticsTailfeeMiApi) FindWmsLogisticsTailfeeMi(c *gin.Context) {
	var wmsLogisticsTailfeeMi manual.WmsLogisticsTailfeeMi
	err := c.ShouldBindQuery(&wmsLogisticsTailfeeMi)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if rewmsLogisticsTailfeeMi, err := wmsLogisticsTailfeeMiService.GetWmsLogisticsTailfeeMi(wmsLogisticsTailfeeMi.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rewmsLogisticsTailfeeMi": rewmsLogisticsTailfeeMi}, c)
	}
}

// GetWmsLogisticsTailfeeMiList 分页获取尾程费用结构列表
// @Tags WmsLogisticsTailfeeMi
// @Summary 分页获取尾程费用结构列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query manualReq.WmsLogisticsTailfeeMiSearch true "分页获取尾程费用结构列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /wmsLogisticsTailfeeMi/getWmsLogisticsTailfeeMiList [get]
func (wmsLogisticsTailfeeMiApi *WmsLogisticsTailfeeMiApi) GetWmsLogisticsTailfeeMiList(c *gin.Context) {
	var pageInfo manualReq.WmsLogisticsTailfeeMiSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := wmsLogisticsTailfeeMiService.GetWmsLogisticsTailfeeMiInfoList(pageInfo); err != nil {
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
