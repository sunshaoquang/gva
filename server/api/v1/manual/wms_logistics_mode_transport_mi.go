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

type WmsLogisticsModeTransportMiApi struct {
}

var wmsLogisticsModeTransportMiService = service.ServiceGroupApp.ManualServiceGroup.WmsLogisticsModeTransportMiService


// CreateWmsLogisticsModeTransportMi 创建费用by运输方式
// @Tags WmsLogisticsModeTransportMi
// @Summary 创建费用by运输方式
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body manual.WmsLogisticsModeTransportMi true "创建费用by运输方式"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /wmsLogisticsModeTransportMi/createWmsLogisticsModeTransportMi [post]
func (wmsLogisticsModeTransportMiApi *WmsLogisticsModeTransportMiApi) CreateWmsLogisticsModeTransportMi(c *gin.Context) {
	var wmsLogisticsModeTransportMi manual.WmsLogisticsModeTransportMi
	err := c.ShouldBindJSON(&wmsLogisticsModeTransportMi)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	// =============== 新增 ==================
	if wmsLogisticsModeTransportMi.CreatedName == "" {
		wmsLogisticsModeTransportMi.CreatedName = utils.GetUserName(c)
	}
	// =============== 新增 ==================
    wmsLogisticsModeTransportMi.CreatedBy = utils.GetUserID(c)
	if err := wmsLogisticsModeTransportMiService.CreateWmsLogisticsModeTransportMi(&wmsLogisticsModeTransportMi); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteWmsLogisticsModeTransportMi 删除费用by运输方式
// @Tags WmsLogisticsModeTransportMi
// @Summary 删除费用by运输方式
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body manual.WmsLogisticsModeTransportMi true "删除费用by运输方式"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /wmsLogisticsModeTransportMi/deleteWmsLogisticsModeTransportMi [delete]
func (wmsLogisticsModeTransportMiApi *WmsLogisticsModeTransportMiApi) DeleteWmsLogisticsModeTransportMi(c *gin.Context) {
	var wmsLogisticsModeTransportMi manual.WmsLogisticsModeTransportMi
	err := c.ShouldBindJSON(&wmsLogisticsModeTransportMi)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    wmsLogisticsModeTransportMi.DeletedBy = utils.GetUserID(c)
	if err := wmsLogisticsModeTransportMiService.DeleteWmsLogisticsModeTransportMi(wmsLogisticsModeTransportMi); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteWmsLogisticsModeTransportMiAll 删除所有费用by运输方式
// @Tags WmsLogisticsModeTransportMiAll
// @Summary 删除所有费用by运输方式
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body manual.WmsLogisticsModeTransportMi true "删除所有费用by运输方式"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /wmsLogisticsModeTransportMi/deleteWmsLogisticsModeTransportMiAll [delete]
func (wmsLogisticsModeTransportMiApi *WmsLogisticsModeTransportMiApi) DeleteWmsLogisticsModeTransportMiAll(c *gin.Context) {
	var wmsLogisticsModeTransportMi manual.WmsLogisticsModeTransportMi
	
	if err := global.MustGetGlobalDBByDBName("ht_smartdata").Unscoped().Where("1 = 1").Delete(&wmsLogisticsModeTransportMi).Error; err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteWmsLogisticsModeTransportMiByIds 批量删除费用by运输方式
// @Tags WmsLogisticsModeTransportMi
// @Summary 批量删除费用by运输方式
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除费用by运输方式"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /wmsLogisticsModeTransportMi/deleteWmsLogisticsModeTransportMiByIds [delete]
func (wmsLogisticsModeTransportMiApi *WmsLogisticsModeTransportMiApi) DeleteWmsLogisticsModeTransportMiByIds(c *gin.Context) {
	var IDS request.IdsReq
    err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    deletedBy := utils.GetUserID(c)
	if err := wmsLogisticsModeTransportMiService.DeleteWmsLogisticsModeTransportMiByIds(IDS,deletedBy); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateWmsLogisticsModeTransportMi 更新费用by运输方式
// @Tags WmsLogisticsModeTransportMi
// @Summary 更新费用by运输方式
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body manual.WmsLogisticsModeTransportMi true "更新费用by运输方式"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /wmsLogisticsModeTransportMi/updateWmsLogisticsModeTransportMi [put]
func (wmsLogisticsModeTransportMiApi *WmsLogisticsModeTransportMiApi) UpdateWmsLogisticsModeTransportMi(c *gin.Context) {
	var wmsLogisticsModeTransportMi manual.WmsLogisticsModeTransportMi
	err := c.ShouldBindJSON(&wmsLogisticsModeTransportMi)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    wmsLogisticsModeTransportMi.UpdatedBy = utils.GetUserID(c)
	if err := wmsLogisticsModeTransportMiService.UpdateWmsLogisticsModeTransportMi(wmsLogisticsModeTransportMi); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindWmsLogisticsModeTransportMi 用id查询费用by运输方式
// @Tags WmsLogisticsModeTransportMi
// @Summary 用id查询费用by运输方式
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query manual.WmsLogisticsModeTransportMi true "用id查询费用by运输方式"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /wmsLogisticsModeTransportMi/findWmsLogisticsModeTransportMi [get]
func (wmsLogisticsModeTransportMiApi *WmsLogisticsModeTransportMiApi) FindWmsLogisticsModeTransportMi(c *gin.Context) {
	var wmsLogisticsModeTransportMi manual.WmsLogisticsModeTransportMi
	err := c.ShouldBindQuery(&wmsLogisticsModeTransportMi)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if rewmsLogisticsModeTransportMi, err := wmsLogisticsModeTransportMiService.GetWmsLogisticsModeTransportMi(wmsLogisticsModeTransportMi.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rewmsLogisticsModeTransportMi": rewmsLogisticsModeTransportMi}, c)
	}
}

// GetWmsLogisticsModeTransportMiList 分页获取费用by运输方式列表
// @Tags WmsLogisticsModeTransportMi
// @Summary 分页获取费用by运输方式列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query manualReq.WmsLogisticsModeTransportMiSearch true "分页获取费用by运输方式列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /wmsLogisticsModeTransportMi/getWmsLogisticsModeTransportMiList [get]
func (wmsLogisticsModeTransportMiApi *WmsLogisticsModeTransportMiApi) GetWmsLogisticsModeTransportMiList(c *gin.Context) {
	var pageInfo manualReq.WmsLogisticsModeTransportMiSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := wmsLogisticsModeTransportMiService.GetWmsLogisticsModeTransportMiInfoList(pageInfo); err != nil {
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
