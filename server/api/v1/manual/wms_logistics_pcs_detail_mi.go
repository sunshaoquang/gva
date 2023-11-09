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

type WmsLogisticsPcsDetailMiApi struct {
}

var wmsLogisticsPcsDetailMiService = service.ServiceGroupApp.ManualServiceGroup.WmsLogisticsPcsDetailMiService


// CreateWmsLogisticsPcsDetailMi 创建主要产品成本明细表
// @Tags WmsLogisticsPcsDetailMi
// @Summary 创建主要产品成本明细表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body manual.WmsLogisticsPcsDetailMi true "创建主要产品成本明细表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /wmsLogisticsPcsDetailMi/createWmsLogisticsPcsDetailMi [post]
func (wmsLogisticsPcsDetailMiApi *WmsLogisticsPcsDetailMiApi) CreateWmsLogisticsPcsDetailMi(c *gin.Context) {
	var wmsLogisticsPcsDetailMi manual.WmsLogisticsPcsDetailMi
	err := c.ShouldBindJSON(&wmsLogisticsPcsDetailMi)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	// =============== 新增 ==================
	if wmsLogisticsPcsDetailMi.CreatedName == "" {
		wmsLogisticsPcsDetailMi.CreatedName = utils.GetUserName(c)
	}
	// =============== 新增 ==================
    wmsLogisticsPcsDetailMi.CreatedBy = utils.GetUserID(c)
	if err := wmsLogisticsPcsDetailMiService.CreateWmsLogisticsPcsDetailMi(&wmsLogisticsPcsDetailMi); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteWmsLogisticsPcsDetailMi 删除主要产品成本明细表
// @Tags WmsLogisticsPcsDetailMi
// @Summary 删除主要产品成本明细表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body manual.WmsLogisticsPcsDetailMi true "删除主要产品成本明细表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /wmsLogisticsPcsDetailMi/deleteWmsLogisticsPcsDetailMi [delete]
func (wmsLogisticsPcsDetailMiApi *WmsLogisticsPcsDetailMiApi) DeleteWmsLogisticsPcsDetailMi(c *gin.Context) {
	var wmsLogisticsPcsDetailMi manual.WmsLogisticsPcsDetailMi
	err := c.ShouldBindJSON(&wmsLogisticsPcsDetailMi)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    wmsLogisticsPcsDetailMi.DeletedBy = utils.GetUserID(c)
	if err := wmsLogisticsPcsDetailMiService.DeleteWmsLogisticsPcsDetailMi(wmsLogisticsPcsDetailMi); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteWmsLogisticsPcsDetailMiByIds 批量删除主要产品成本明细表
// @Tags WmsLogisticsPcsDetailMi
// @Summary 批量删除主要产品成本明细表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除主要产品成本明细表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /wmsLogisticsPcsDetailMi/deleteWmsLogisticsPcsDetailMiByIds [delete]
func (wmsLogisticsPcsDetailMiApi *WmsLogisticsPcsDetailMiApi) DeleteWmsLogisticsPcsDetailMiByIds(c *gin.Context) {
	var IDS request.IdsReq
    err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    deletedBy := utils.GetUserID(c)
	if err := wmsLogisticsPcsDetailMiService.DeleteWmsLogisticsPcsDetailMiByIds(IDS,deletedBy); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateWmsLogisticsPcsDetailMi 更新主要产品成本明细表
// @Tags WmsLogisticsPcsDetailMi
// @Summary 更新主要产品成本明细表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body manual.WmsLogisticsPcsDetailMi true "更新主要产品成本明细表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /wmsLogisticsPcsDetailMi/updateWmsLogisticsPcsDetailMi [put]
func (wmsLogisticsPcsDetailMiApi *WmsLogisticsPcsDetailMiApi) UpdateWmsLogisticsPcsDetailMi(c *gin.Context) {
	var wmsLogisticsPcsDetailMi manual.WmsLogisticsPcsDetailMi
	err := c.ShouldBindJSON(&wmsLogisticsPcsDetailMi)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    wmsLogisticsPcsDetailMi.UpdatedBy = utils.GetUserID(c)
	if err := wmsLogisticsPcsDetailMiService.UpdateWmsLogisticsPcsDetailMi(wmsLogisticsPcsDetailMi); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindWmsLogisticsPcsDetailMi 用id查询主要产品成本明细表
// @Tags WmsLogisticsPcsDetailMi
// @Summary 用id查询主要产品成本明细表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query manual.WmsLogisticsPcsDetailMi true "用id查询主要产品成本明细表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /wmsLogisticsPcsDetailMi/findWmsLogisticsPcsDetailMi [get]
func (wmsLogisticsPcsDetailMiApi *WmsLogisticsPcsDetailMiApi) FindWmsLogisticsPcsDetailMi(c *gin.Context) {
	var wmsLogisticsPcsDetailMi manual.WmsLogisticsPcsDetailMi
	err := c.ShouldBindQuery(&wmsLogisticsPcsDetailMi)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if rewmsLogisticsPcsDetailMi, err := wmsLogisticsPcsDetailMiService.GetWmsLogisticsPcsDetailMi(wmsLogisticsPcsDetailMi.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rewmsLogisticsPcsDetailMi": rewmsLogisticsPcsDetailMi}, c)
	}
}

// GetWmsLogisticsPcsDetailMiList 分页获取主要产品成本明细表列表
// @Tags WmsLogisticsPcsDetailMi
// @Summary 分页获取主要产品成本明细表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query manualReq.WmsLogisticsPcsDetailMiSearch true "分页获取主要产品成本明细表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /wmsLogisticsPcsDetailMi/getWmsLogisticsPcsDetailMiList [get]
func (wmsLogisticsPcsDetailMiApi *WmsLogisticsPcsDetailMiApi) GetWmsLogisticsPcsDetailMiList(c *gin.Context) {
	var pageInfo manualReq.WmsLogisticsPcsDetailMiSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := wmsLogisticsPcsDetailMiService.GetWmsLogisticsPcsDetailMiInfoList(pageInfo); err != nil {
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
