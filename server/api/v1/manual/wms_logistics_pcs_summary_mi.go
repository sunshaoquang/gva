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

type WmsLogisticsPcsSummaryMiApi struct {
}

var wmsLogisticsPcsSummaryMiService = service.ServiceGroupApp.ManualServiceGroup.WmsLogisticsPcsSummaryMiService


// CreateWmsLogisticsPcsSummaryMi 创建主要产品成本降本汇总表
// @Tags WmsLogisticsPcsSummaryMi
// @Summary 创建主要产品成本降本汇总表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body manual.WmsLogisticsPcsSummaryMi true "创建主要产品成本降本汇总表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /wmsLogisticsPcsSummaryMi/createWmsLogisticsPcsSummaryMi [post]
func (wmsLogisticsPcsSummaryMiApi *WmsLogisticsPcsSummaryMiApi) CreateWmsLogisticsPcsSummaryMi(c *gin.Context) {
	var wmsLogisticsPcsSummaryMi manual.WmsLogisticsPcsSummaryMi
	err := c.ShouldBindJSON(&wmsLogisticsPcsSummaryMi)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	// =============== 新增 ==================
	if wmsLogisticsPcsSummaryMi.CreatedName == "" {
		wmsLogisticsPcsSummaryMi.CreatedName = utils.GetUserName(c)
	}
	// =============== 新增 ==================
    wmsLogisticsPcsSummaryMi.CreatedBy = utils.GetUserID(c)
	if err := wmsLogisticsPcsSummaryMiService.CreateWmsLogisticsPcsSummaryMi(&wmsLogisticsPcsSummaryMi); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteWmsLogisticsPcsSummaryMi 删除主要产品成本降本汇总表
// @Tags WmsLogisticsPcsSummaryMi
// @Summary 删除主要产品成本降本汇总表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body manual.WmsLogisticsPcsSummaryMi true "删除主要产品成本降本汇总表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /wmsLogisticsPcsSummaryMi/deleteWmsLogisticsPcsSummaryMi [delete]
func (wmsLogisticsPcsSummaryMiApi *WmsLogisticsPcsSummaryMiApi) DeleteWmsLogisticsPcsSummaryMi(c *gin.Context) {
	var wmsLogisticsPcsSummaryMi manual.WmsLogisticsPcsSummaryMi
	err := c.ShouldBindJSON(&wmsLogisticsPcsSummaryMi)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    wmsLogisticsPcsSummaryMi.DeletedBy = utils.GetUserID(c)
	if err := wmsLogisticsPcsSummaryMiService.DeleteWmsLogisticsPcsSummaryMi(wmsLogisticsPcsSummaryMi); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteWmsLogisticsPcsSummaryMiAll 删除主要产品成本降本汇总表
// @Tags WmsLogisticsPcsSummaryMiAll
// @Summary 删除主要产品成本降本汇总表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body manual.WmsLogisticsPcsSummaryMiAll true "删除主要产品成本降本汇总表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /wmsLogisticsPcsSummaryMi/deleteWmsLogisticsPcsSummaryMiAll [delete]
func (wmsLogisticsPcsSummaryMiApi *WmsLogisticsPcsSummaryMiApi) DeleteWmsLogisticsPcsSummaryMiAll(c *gin.Context) {
	var wmsLogisticsPcsSummaryMi manual.WmsLogisticsPcsSummaryMi
	if err := global.MustGetGlobalDBByDBName("ht_smartdata").Unscoped().Where("1 = 1").Delete(&wmsLogisticsPcsSummaryMi).Error; err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteWmsLogisticsPcsSummaryMiByIds 批量删除主要产品成本降本汇总表
// @Tags WmsLogisticsPcsSummaryMi
// @Summary 批量删除主要产品成本降本汇总表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除主要产品成本降本汇总表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /wmsLogisticsPcsSummaryMi/deleteWmsLogisticsPcsSummaryMiByIds [delete]
func (wmsLogisticsPcsSummaryMiApi *WmsLogisticsPcsSummaryMiApi) DeleteWmsLogisticsPcsSummaryMiByIds(c *gin.Context) {
	var IDS request.IdsReq
    err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    deletedBy := utils.GetUserID(c)
	if err := wmsLogisticsPcsSummaryMiService.DeleteWmsLogisticsPcsSummaryMiByIds(IDS,deletedBy); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateWmsLogisticsPcsSummaryMi 更新主要产品成本降本汇总表
// @Tags WmsLogisticsPcsSummaryMi
// @Summary 更新主要产品成本降本汇总表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body manual.WmsLogisticsPcsSummaryMi true "更新主要产品成本降本汇总表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /wmsLogisticsPcsSummaryMi/updateWmsLogisticsPcsSummaryMi [put]
func (wmsLogisticsPcsSummaryMiApi *WmsLogisticsPcsSummaryMiApi) UpdateWmsLogisticsPcsSummaryMi(c *gin.Context) {
	var wmsLogisticsPcsSummaryMi manual.WmsLogisticsPcsSummaryMi
	err := c.ShouldBindJSON(&wmsLogisticsPcsSummaryMi)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    wmsLogisticsPcsSummaryMi.UpdatedBy = utils.GetUserID(c)
	if err := wmsLogisticsPcsSummaryMiService.UpdateWmsLogisticsPcsSummaryMi(wmsLogisticsPcsSummaryMi); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindWmsLogisticsPcsSummaryMi 用id查询主要产品成本降本汇总表
// @Tags WmsLogisticsPcsSummaryMi
// @Summary 用id查询主要产品成本降本汇总表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query manual.WmsLogisticsPcsSummaryMi true "用id查询主要产品成本降本汇总表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /wmsLogisticsPcsSummaryMi/findWmsLogisticsPcsSummaryMi [get]
func (wmsLogisticsPcsSummaryMiApi *WmsLogisticsPcsSummaryMiApi) FindWmsLogisticsPcsSummaryMi(c *gin.Context) {
	var wmsLogisticsPcsSummaryMi manual.WmsLogisticsPcsSummaryMi
	err := c.ShouldBindQuery(&wmsLogisticsPcsSummaryMi)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if rewmsLogisticsPcsSummaryMi, err := wmsLogisticsPcsSummaryMiService.GetWmsLogisticsPcsSummaryMi(wmsLogisticsPcsSummaryMi.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rewmsLogisticsPcsSummaryMi": rewmsLogisticsPcsSummaryMi}, c)
	}
}

// GetWmsLogisticsPcsSummaryMiList 分页获取主要产品成本降本汇总表列表
// @Tags WmsLogisticsPcsSummaryMi
// @Summary 分页获取主要产品成本降本汇总表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query manualReq.WmsLogisticsPcsSummaryMiSearch true "分页获取主要产品成本降本汇总表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /wmsLogisticsPcsSummaryMi/getWmsLogisticsPcsSummaryMiList [get]
func (wmsLogisticsPcsSummaryMiApi *WmsLogisticsPcsSummaryMiApi) GetWmsLogisticsPcsSummaryMiList(c *gin.Context) {
	var pageInfo manualReq.WmsLogisticsPcsSummaryMiSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := wmsLogisticsPcsSummaryMiService.GetWmsLogisticsPcsSummaryMiInfoList(pageInfo); err != nil {
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
