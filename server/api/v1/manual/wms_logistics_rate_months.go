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

type WmsLogisticsRateMonthsApi struct {
}

var wmsLogisticsRateMonthsService = service.ServiceGroupApp.ManualServiceGroup.WmsLogisticsRateMonthsService


// CreateWmsLogisticsRateMonths 创建wmsLogisticsRateMonths表
// @Tags WmsLogisticsRateMonths
// @Summary 创建wmsLogisticsRateMonths表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body manual.WmsLogisticsRateMonths true "创建wmsLogisticsRateMonths表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /wmsLogisticsRateMonths/createWmsLogisticsRateMonths [post]
func (wmsLogisticsRateMonthsApi *WmsLogisticsRateMonthsApi) CreateWmsLogisticsRateMonths(c *gin.Context) {
	var wmsLogisticsRateMonths manual.WmsLogisticsRateMonths
	err := c.ShouldBindJSON(&wmsLogisticsRateMonths)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	// =============== 新增 ==================
	if wmsLogisticsRateMonths.CreatedName == "" {
		wmsLogisticsRateMonths.CreatedName = utils.GetUserName(c)
	}
	// =============== 新增 ==================
    wmsLogisticsRateMonths.CreatedBy = utils.GetUserID(c)
	if err := wmsLogisticsRateMonthsService.CreateWmsLogisticsRateMonths(&wmsLogisticsRateMonths); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteWmsLogisticsRateMonths 删除wmsLogisticsRateMonths表
// @Tags WmsLogisticsRateMonths
// @Summary 删除wmsLogisticsRateMonths表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body manual.WmsLogisticsRateMonths true "删除wmsLogisticsRateMonths表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /wmsLogisticsRateMonths/deleteWmsLogisticsRateMonths [delete]
func (wmsLogisticsRateMonthsApi *WmsLogisticsRateMonthsApi) DeleteWmsLogisticsRateMonths(c *gin.Context) {
	var wmsLogisticsRateMonths manual.WmsLogisticsRateMonths
	err := c.ShouldBindJSON(&wmsLogisticsRateMonths)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    wmsLogisticsRateMonths.DeletedBy = utils.GetUserID(c)
	if err := wmsLogisticsRateMonthsService.DeleteWmsLogisticsRateMonths(wmsLogisticsRateMonths); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteWmsLogisticsRateMonthsAll 删除所有wmsLogisticsRateMonths表
// @Tags WmsLogisticsRateMonthsAll
// @Summary 删除所有wmsLogisticsRateMonths表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body manual.WmsLogisticsRateMonths true "删除所有wmsLogisticsRateMonths表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /wmsLogisticsRateMonths/deleteWmsLogisticsRateMonthsAll [delete]
func (wmsLogisticsRateMonthsApi *WmsLogisticsRateMonthsApi) DeleteWmsLogisticsRateMonthsAll(c *gin.Context) {
	var wmsLogisticsRateMonths manual.WmsLogisticsRateMonths
	
	if err := global.MustGetGlobalDBByDBName("ht_smartdata").Unscoped().Where("1 = 1").Delete(&wmsLogisticsRateMonths).Error; err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteWmsLogisticsRateMonthsByIds 批量删除wmsLogisticsRateMonths表
// @Tags WmsLogisticsRateMonths
// @Summary 批量删除wmsLogisticsRateMonths表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除wmsLogisticsRateMonths表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /wmsLogisticsRateMonths/deleteWmsLogisticsRateMonthsByIds [delete]
func (wmsLogisticsRateMonthsApi *WmsLogisticsRateMonthsApi) DeleteWmsLogisticsRateMonthsByIds(c *gin.Context) {
	var IDS request.IdsReq
    err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    deletedBy := utils.GetUserID(c)
	if err := wmsLogisticsRateMonthsService.DeleteWmsLogisticsRateMonthsByIds(IDS,deletedBy); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateWmsLogisticsRateMonths 更新wmsLogisticsRateMonths表
// @Tags WmsLogisticsRateMonths
// @Summary 更新wmsLogisticsRateMonths表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body manual.WmsLogisticsRateMonths true "更新wmsLogisticsRateMonths表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /wmsLogisticsRateMonths/updateWmsLogisticsRateMonths [put]
func (wmsLogisticsRateMonthsApi *WmsLogisticsRateMonthsApi) UpdateWmsLogisticsRateMonths(c *gin.Context) {
	var wmsLogisticsRateMonths manual.WmsLogisticsRateMonths
	err := c.ShouldBindJSON(&wmsLogisticsRateMonths)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    wmsLogisticsRateMonths.UpdatedBy = utils.GetUserID(c)
	if err := wmsLogisticsRateMonthsService.UpdateWmsLogisticsRateMonths(wmsLogisticsRateMonths); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindWmsLogisticsRateMonths 用id查询wmsLogisticsRateMonths表
// @Tags WmsLogisticsRateMonths
// @Summary 用id查询wmsLogisticsRateMonths表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query manual.WmsLogisticsRateMonths true "用id查询wmsLogisticsRateMonths表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /wmsLogisticsRateMonths/findWmsLogisticsRateMonths [get]
func (wmsLogisticsRateMonthsApi *WmsLogisticsRateMonthsApi) FindWmsLogisticsRateMonths(c *gin.Context) {
	var wmsLogisticsRateMonths manual.WmsLogisticsRateMonths
	err := c.ShouldBindQuery(&wmsLogisticsRateMonths)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if rewmsLogisticsRateMonths, err := wmsLogisticsRateMonthsService.GetWmsLogisticsRateMonths(wmsLogisticsRateMonths.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rewmsLogisticsRateMonths": rewmsLogisticsRateMonths}, c)
	}
}

// GetWmsLogisticsRateMonthsList 分页获取wmsLogisticsRateMonths表列表
// @Tags WmsLogisticsRateMonths
// @Summary 分页获取wmsLogisticsRateMonths表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query manualReq.WmsLogisticsRateMonthsSearch true "分页获取wmsLogisticsRateMonths表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /wmsLogisticsRateMonths/getWmsLogisticsRateMonthsList [get]
func (wmsLogisticsRateMonthsApi *WmsLogisticsRateMonthsApi) GetWmsLogisticsRateMonthsList(c *gin.Context) {
	var pageInfo manualReq.WmsLogisticsRateMonthsSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := wmsLogisticsRateMonthsService.GetWmsLogisticsRateMonthsInfoList(pageInfo); err != nil {
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
