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

type WmsKpi2023SummaryMiApi struct {
}

var wmsKpi2023SummaryMiService = service.ServiceGroupApp.ManualServiceGroup.WmsKpi2023SummaryMiService


// CreateWmsKpi2023SummaryMi 创建wmsKpi2023SummaryMi表
// @Tags WmsKpi2023SummaryMi
// @Summary 创建wmsKpi2023SummaryMi表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body manual.WmsKpi2023SummaryMi true "创建wmsKpi2023SummaryMi表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /wmsKpi2023SummaryMi/createWmsKpi2023SummaryMi [post]
func (wmsKpi2023SummaryMiApi *WmsKpi2023SummaryMiApi) CreateWmsKpi2023SummaryMi(c *gin.Context) {
	var wmsKpi2023SummaryMi manual.WmsKpi2023SummaryMi
	err := c.ShouldBindJSON(&wmsKpi2023SummaryMi)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	// =============== 新增 ==================
	if wmsKpi2023SummaryMi.CreatedName == "" {
		wmsKpi2023SummaryMi.CreatedName = utils.GetUserName(c)
	}
	// =============== 新增 ==================
    wmsKpi2023SummaryMi.CreatedBy = utils.GetUserID(c)
	if err := wmsKpi2023SummaryMiService.CreateWmsKpi2023SummaryMi(&wmsKpi2023SummaryMi); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteWmsKpi2023SummaryMi 删除wmsKpi2023SummaryMi表
// @Tags WmsKpi2023SummaryMi
// @Summary 删除wmsKpi2023SummaryMi表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body manual.WmsKpi2023SummaryMi true "删除wmsKpi2023SummaryMi表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /wmsKpi2023SummaryMi/deleteWmsKpi2023SummaryMi [delete]
func (wmsKpi2023SummaryMiApi *WmsKpi2023SummaryMiApi) DeleteWmsKpi2023SummaryMi(c *gin.Context) {
	var wmsKpi2023SummaryMi manual.WmsKpi2023SummaryMi
	err := c.ShouldBindJSON(&wmsKpi2023SummaryMi)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    wmsKpi2023SummaryMi.DeletedBy = utils.GetUserID(c)
	if err := wmsKpi2023SummaryMiService.DeleteWmsKpi2023SummaryMi(wmsKpi2023SummaryMi); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteWmsKpi2023SummaryMiByIds 批量删除wmsKpi2023SummaryMi表
// @Tags WmsKpi2023SummaryMi
// @Summary 批量删除wmsKpi2023SummaryMi表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除wmsKpi2023SummaryMi表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /wmsKpi2023SummaryMi/deleteWmsKpi2023SummaryMiByIds [delete]
func (wmsKpi2023SummaryMiApi *WmsKpi2023SummaryMiApi) DeleteWmsKpi2023SummaryMiByIds(c *gin.Context) {
	var IDS request.IdsReq
    err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    deletedBy := utils.GetUserID(c)
	if err := wmsKpi2023SummaryMiService.DeleteWmsKpi2023SummaryMiByIds(IDS,deletedBy); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateWmsKpi2023SummaryMi 更新wmsKpi2023SummaryMi表
// @Tags WmsKpi2023SummaryMi
// @Summary 更新wmsKpi2023SummaryMi表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body manual.WmsKpi2023SummaryMi true "更新wmsKpi2023SummaryMi表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /wmsKpi2023SummaryMi/updateWmsKpi2023SummaryMi [put]
func (wmsKpi2023SummaryMiApi *WmsKpi2023SummaryMiApi) UpdateWmsKpi2023SummaryMi(c *gin.Context) {
	var wmsKpi2023SummaryMi manual.WmsKpi2023SummaryMi
	err := c.ShouldBindJSON(&wmsKpi2023SummaryMi)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    wmsKpi2023SummaryMi.UpdatedBy = utils.GetUserID(c)
	if err := wmsKpi2023SummaryMiService.UpdateWmsKpi2023SummaryMi(wmsKpi2023SummaryMi); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindWmsKpi2023SummaryMi 用id查询wmsKpi2023SummaryMi表
// @Tags WmsKpi2023SummaryMi
// @Summary 用id查询wmsKpi2023SummaryMi表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query manual.WmsKpi2023SummaryMi true "用id查询wmsKpi2023SummaryMi表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /wmsKpi2023SummaryMi/findWmsKpi2023SummaryMi [get]
func (wmsKpi2023SummaryMiApi *WmsKpi2023SummaryMiApi) FindWmsKpi2023SummaryMi(c *gin.Context) {
	var wmsKpi2023SummaryMi manual.WmsKpi2023SummaryMi
	err := c.ShouldBindQuery(&wmsKpi2023SummaryMi)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if rewmsKpi2023SummaryMi, err := wmsKpi2023SummaryMiService.GetWmsKpi2023SummaryMi(wmsKpi2023SummaryMi.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rewmsKpi2023SummaryMi": rewmsKpi2023SummaryMi}, c)
	}
}

// GetWmsKpi2023SummaryMiList 分页获取wmsKpi2023SummaryMi表列表
// @Tags WmsKpi2023SummaryMi
// @Summary 分页获取wmsKpi2023SummaryMi表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query manualReq.WmsKpi2023SummaryMiSearch true "分页获取wmsKpi2023SummaryMi表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /wmsKpi2023SummaryMi/getWmsKpi2023SummaryMiList [get]
func (wmsKpi2023SummaryMiApi *WmsKpi2023SummaryMiApi) GetWmsKpi2023SummaryMiList(c *gin.Context) {
	var pageInfo manualReq.WmsKpi2023SummaryMiSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := wmsKpi2023SummaryMiService.GetWmsKpi2023SummaryMiInfoList(pageInfo); err != nil {
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
