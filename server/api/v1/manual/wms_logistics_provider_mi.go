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

type WmsLogisticsProviderMiApi struct {
}

var wmsLogisticsProviderMiService = service.ServiceGroupApp.ManualServiceGroup.WmsLogisticsProviderMiService


// CreateWmsLogisticsProviderMi 创建费用by供应商
// @Tags WmsLogisticsProviderMi
// @Summary 创建费用by供应商
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body manual.WmsLogisticsProviderMi true "创建费用by供应商"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /wmsLogisticsProviderMi/createWmsLogisticsProviderMi [post]
func (wmsLogisticsProviderMiApi *WmsLogisticsProviderMiApi) CreateWmsLogisticsProviderMi(c *gin.Context) {
	var wmsLogisticsProviderMi manual.WmsLogisticsProviderMi
	err := c.ShouldBindJSON(&wmsLogisticsProviderMi)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	// =============== 新增 ==================
	if wmsLogisticsProviderMi.CreatedName == "" {
		wmsLogisticsProviderMi.CreatedName = utils.GetUserName(c)
	}
	// =============== 新增 ==================
    wmsLogisticsProviderMi.CreatedBy = utils.GetUserID(c)
	if err := wmsLogisticsProviderMiService.CreateWmsLogisticsProviderMi(&wmsLogisticsProviderMi); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteWmsLogisticsProviderMi 删除费用by供应商
// @Tags WmsLogisticsProviderMi
// @Summary 删除费用by供应商
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body manual.WmsLogisticsProviderMi true "删除费用by供应商"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /wmsLogisticsProviderMi/deleteWmsLogisticsProviderMi [delete]
func (wmsLogisticsProviderMiApi *WmsLogisticsProviderMiApi) DeleteWmsLogisticsProviderMi(c *gin.Context) {
	var wmsLogisticsProviderMi manual.WmsLogisticsProviderMi
	err := c.ShouldBindJSON(&wmsLogisticsProviderMi)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    wmsLogisticsProviderMi.DeletedBy = utils.GetUserID(c)
	if err := wmsLogisticsProviderMiService.DeleteWmsLogisticsProviderMi(wmsLogisticsProviderMi); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteWmsLogisticsProviderMiAll 删除所有费用by供应商
// @Tags WmsLogisticsProviderMiAll
// @Summary 删除所有费用by供应商
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body manual.WmsLogisticsProviderMi true "删除所有费用by供应商"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /wmsLogisticsProviderMi/deleteWmsLogisticsProviderMiAll [delete]
func (wmsLogisticsProviderMiApi *WmsLogisticsProviderMiApi) DeleteWmsLogisticsProviderMiAll(c *gin.Context) {
	var wmsLogisticsProviderMi manual.WmsLogisticsProviderMi
	
	if err := global.MustGetGlobalDBByDBName("ht_smartdata").Unscoped().Where("1 = 1").Delete(&wmsLogisticsProviderMi).Error; err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteWmsLogisticsProviderMiByIds 批量删除费用by供应商
// @Tags WmsLogisticsProviderMi
// @Summary 批量删除费用by供应商
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除费用by供应商"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /wmsLogisticsProviderMi/deleteWmsLogisticsProviderMiByIds [delete]
func (wmsLogisticsProviderMiApi *WmsLogisticsProviderMiApi) DeleteWmsLogisticsProviderMiByIds(c *gin.Context) {
	var IDS request.IdsReq
    err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    deletedBy := utils.GetUserID(c)
	if err := wmsLogisticsProviderMiService.DeleteWmsLogisticsProviderMiByIds(IDS,deletedBy); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateWmsLogisticsProviderMi 更新费用by供应商
// @Tags WmsLogisticsProviderMi
// @Summary 更新费用by供应商
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body manual.WmsLogisticsProviderMi true "更新费用by供应商"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /wmsLogisticsProviderMi/updateWmsLogisticsProviderMi [put]
func (wmsLogisticsProviderMiApi *WmsLogisticsProviderMiApi) UpdateWmsLogisticsProviderMi(c *gin.Context) {
	var wmsLogisticsProviderMi manual.WmsLogisticsProviderMi
	err := c.ShouldBindJSON(&wmsLogisticsProviderMi)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    wmsLogisticsProviderMi.UpdatedBy = utils.GetUserID(c)
	if err := wmsLogisticsProviderMiService.UpdateWmsLogisticsProviderMi(wmsLogisticsProviderMi); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindWmsLogisticsProviderMi 用id查询费用by供应商
// @Tags WmsLogisticsProviderMi
// @Summary 用id查询费用by供应商
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query manual.WmsLogisticsProviderMi true "用id查询费用by供应商"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /wmsLogisticsProviderMi/findWmsLogisticsProviderMi [get]
func (wmsLogisticsProviderMiApi *WmsLogisticsProviderMiApi) FindWmsLogisticsProviderMi(c *gin.Context) {
	var wmsLogisticsProviderMi manual.WmsLogisticsProviderMi
	err := c.ShouldBindQuery(&wmsLogisticsProviderMi)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if rewmsLogisticsProviderMi, err := wmsLogisticsProviderMiService.GetWmsLogisticsProviderMi(wmsLogisticsProviderMi.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rewmsLogisticsProviderMi": rewmsLogisticsProviderMi}, c)
	}
}

// GetWmsLogisticsProviderMiList 分页获取费用by供应商列表
// @Tags WmsLogisticsProviderMi
// @Summary 分页获取费用by供应商列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query manualReq.WmsLogisticsProviderMiSearch true "分页获取费用by供应商列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /wmsLogisticsProviderMi/getWmsLogisticsProviderMiList [get]
func (wmsLogisticsProviderMiApi *WmsLogisticsProviderMiApi) GetWmsLogisticsProviderMiList(c *gin.Context) {
	var pageInfo manualReq.WmsLogisticsProviderMiSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := wmsLogisticsProviderMiService.GetWmsLogisticsProviderMiInfoList(pageInfo); err != nil {
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
