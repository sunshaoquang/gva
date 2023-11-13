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

type WmsLogisticsPcs2023MiApi struct {
}

var wmsLogisticsPcs2023MiService = service.ServiceGroupApp.ManualServiceGroup.WmsLogisticsPcs2023MiService


// CreateWmsLogisticsPcs2023Mi 创建2023年物流成本明细表
// @Tags WmsLogisticsPcs2023Mi
// @Summary 创建2023年物流成本明细表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body manual.WmsLogisticsPcs2023Mi true "创建2023年物流成本明细表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /wmsLogisticsPcs2023Mi/createWmsLogisticsPcs2023Mi [post]
func (wmsLogisticsPcs2023MiApi *WmsLogisticsPcs2023MiApi) CreateWmsLogisticsPcs2023Mi(c *gin.Context) {
	var wmsLogisticsPcs2023Mi manual.WmsLogisticsPcs2023Mi
	err := c.ShouldBindJSON(&wmsLogisticsPcs2023Mi)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	// =============== 新增 ==================
	if wmsLogisticsPcs2023Mi.CreatedName == "" {
		wmsLogisticsPcs2023Mi.CreatedName = utils.GetUserName(c)
	}
	// =============== 新增 ==================
    wmsLogisticsPcs2023Mi.CreatedBy = utils.GetUserID(c)
	if err := wmsLogisticsPcs2023MiService.CreateWmsLogisticsPcs2023Mi(&wmsLogisticsPcs2023Mi); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteWmsLogisticsPcs2023Mi 删除2023年物流成本明细表
// @Tags WmsLogisticsPcs2023Mi
// @Summary 删除2023年物流成本明细表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body manual.WmsLogisticsPcs2023Mi true "删除2023年物流成本明细表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /wmsLogisticsPcs2023Mi/deleteWmsLogisticsPcs2023Mi [delete]
func (wmsLogisticsPcs2023MiApi *WmsLogisticsPcs2023MiApi) DeleteWmsLogisticsPcs2023Mi(c *gin.Context) {
	var wmsLogisticsPcs2023Mi manual.WmsLogisticsPcs2023Mi
	err := c.ShouldBindJSON(&wmsLogisticsPcs2023Mi)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    wmsLogisticsPcs2023Mi.DeletedBy = utils.GetUserID(c)
	if err := wmsLogisticsPcs2023MiService.DeleteWmsLogisticsPcs2023Mi(wmsLogisticsPcs2023Mi); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteWmsLogisticsPcs2023MiAll 删除所有2023年物流成本明细表
// @Tags WmsLogisticsPcs2023MiAll
// @Summary 删除所有2023年物流成本明细表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body manual.WmsLogisticsPcs2023MiAll true "删除所有2023年物流成本明细表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /wmsLogisticsPcs2023Mi/deleteWmsLogisticsPcs2023MiAll [delete]
func (wmsLogisticsPcs2023MiApi *WmsLogisticsPcs2023MiApi) DeleteWmsLogisticsPcs2023MiAll(c *gin.Context) {
	var wmsLogisticsPcs2023Mi manual.WmsLogisticsPcs2023Mi
	// 用gorm硬删除所有表数据
	if err := global.MustGetGlobalDBByDBName("ht_smartdata").Unscoped().Where("1 = 1").Delete(&wmsLogisticsPcs2023Mi).Error; err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteWmsLogisticsPcs2023MiByIds 批量删除2023年物流成本明细表
// @Tags WmsLogisticsPcs2023Mi
// @Summary 批量删除2023年物流成本明细表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除2023年物流成本明细表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /wmsLogisticsPcs2023Mi/deleteWmsLogisticsPcs2023MiByIds [delete]
func (wmsLogisticsPcs2023MiApi *WmsLogisticsPcs2023MiApi) DeleteWmsLogisticsPcs2023MiByIds(c *gin.Context) {
	var IDS request.IdsReq
    err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    deletedBy := utils.GetUserID(c)
	if err := wmsLogisticsPcs2023MiService.DeleteWmsLogisticsPcs2023MiByIds(IDS,deletedBy); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateWmsLogisticsPcs2023Mi 更新2023年物流成本明细表
// @Tags WmsLogisticsPcs2023Mi
// @Summary 更新2023年物流成本明细表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body manual.WmsLogisticsPcs2023Mi true "更新2023年物流成本明细表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /wmsLogisticsPcs2023Mi/updateWmsLogisticsPcs2023Mi [put]
func (wmsLogisticsPcs2023MiApi *WmsLogisticsPcs2023MiApi) UpdateWmsLogisticsPcs2023Mi(c *gin.Context) {
	var wmsLogisticsPcs2023Mi manual.WmsLogisticsPcs2023Mi
	err := c.ShouldBindJSON(&wmsLogisticsPcs2023Mi)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    wmsLogisticsPcs2023Mi.UpdatedBy = utils.GetUserID(c)
	if err := wmsLogisticsPcs2023MiService.UpdateWmsLogisticsPcs2023Mi(wmsLogisticsPcs2023Mi); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindWmsLogisticsPcs2023Mi 用id查询2023年物流成本明细表
// @Tags WmsLogisticsPcs2023Mi
// @Summary 用id查询2023年物流成本明细表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query manual.WmsLogisticsPcs2023Mi true "用id查询2023年物流成本明细表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /wmsLogisticsPcs2023Mi/findWmsLogisticsPcs2023Mi [get]
func (wmsLogisticsPcs2023MiApi *WmsLogisticsPcs2023MiApi) FindWmsLogisticsPcs2023Mi(c *gin.Context) {
	var wmsLogisticsPcs2023Mi manual.WmsLogisticsPcs2023Mi
	err := c.ShouldBindQuery(&wmsLogisticsPcs2023Mi)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if rewmsLogisticsPcs2023Mi, err := wmsLogisticsPcs2023MiService.GetWmsLogisticsPcs2023Mi(wmsLogisticsPcs2023Mi.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rewmsLogisticsPcs2023Mi": rewmsLogisticsPcs2023Mi}, c)
	}
}

// GetWmsLogisticsPcs2023MiList 分页获取2023年物流成本明细表列表
// @Tags WmsLogisticsPcs2023Mi
// @Summary 分页获取2023年物流成本明细表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query manualReq.WmsLogisticsPcs2023MiSearch true "分页获取2023年物流成本明细表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /wmsLogisticsPcs2023Mi/getWmsLogisticsPcs2023MiList [get]
func (wmsLogisticsPcs2023MiApi *WmsLogisticsPcs2023MiApi) GetWmsLogisticsPcs2023MiList(c *gin.Context) {
	var pageInfo manualReq.WmsLogisticsPcs2023MiSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := wmsLogisticsPcs2023MiService.GetWmsLogisticsPcs2023MiInfoList(pageInfo); err != nil {
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
