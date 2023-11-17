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

type WmsLogisticsStoragefeeMiApi struct {
}

var wmsLogisticsStoragefeeMiService = service.ServiceGroupApp.ManualServiceGroup.WmsLogisticsStoragefeeMiService


// CreateWmsLogisticsStoragefeeMi 创建仓储费
// @Tags WmsLogisticsStoragefeeMi
// @Summary 创建仓储费
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body manual.WmsLogisticsStoragefeeMi true "创建仓储费"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /wmsLogisticsStoragefeeMi/createWmsLogisticsStoragefeeMi [post]
func (wmsLogisticsStoragefeeMiApi *WmsLogisticsStoragefeeMiApi) CreateWmsLogisticsStoragefeeMi(c *gin.Context) {
	var wmsLogisticsStoragefeeMi manual.WmsLogisticsStoragefeeMi
	err := c.ShouldBindJSON(&wmsLogisticsStoragefeeMi)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	// =============== 新增 ==================
	if wmsLogisticsStoragefeeMi.CreatedName == "" {
		wmsLogisticsStoragefeeMi.CreatedName = utils.GetUserName(c)
	}
	// =============== 新增 ==================
    wmsLogisticsStoragefeeMi.CreatedBy = utils.GetUserID(c)
	if err := wmsLogisticsStoragefeeMiService.CreateWmsLogisticsStoragefeeMi(&wmsLogisticsStoragefeeMi); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteWmsLogisticsStoragefeeMi 删除仓储费
// @Tags WmsLogisticsStoragefeeMi
// @Summary 删除仓储费
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body manual.WmsLogisticsStoragefeeMi true "删除仓储费"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /wmsLogisticsStoragefeeMi/deleteWmsLogisticsStoragefeeMi [delete]
func (wmsLogisticsStoragefeeMiApi *WmsLogisticsStoragefeeMiApi) DeleteWmsLogisticsStoragefeeMi(c *gin.Context) {
	var wmsLogisticsStoragefeeMi manual.WmsLogisticsStoragefeeMi
	err := c.ShouldBindJSON(&wmsLogisticsStoragefeeMi)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    wmsLogisticsStoragefeeMi.DeletedBy = utils.GetUserID(c)
	if err := wmsLogisticsStoragefeeMiService.DeleteWmsLogisticsStoragefeeMi(wmsLogisticsStoragefeeMi); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteWmsLogisticsStoragefeeMiAll 删除所有仓储费
// @Tags WmsLogisticsStoragefeeMiAll
// @Summary 删除所有仓储费
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body manual.WmsLogisticsStoragefeeMi true "删除所有仓储费"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /wmsLogisticsStoragefeeMi/deleteWmsLogisticsStoragefeeMiAll [delete]
func (wmsLogisticsStoragefeeMiApi *WmsLogisticsStoragefeeMiApi) DeleteWmsLogisticsStoragefeeMiAll(c *gin.Context) {
	var wmsLogisticsStoragefeeMi manual.WmsLogisticsStoragefeeMi
	
	if err := global.MustGetGlobalDBByDBName("ht_smartdata").Unscoped().Where("1 = 1").Delete(&wmsLogisticsStoragefeeMi).Error; err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteWmsLogisticsStoragefeeMiByIds 批量删除仓储费
// @Tags WmsLogisticsStoragefeeMi
// @Summary 批量删除仓储费
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除仓储费"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /wmsLogisticsStoragefeeMi/deleteWmsLogisticsStoragefeeMiByIds [delete]
func (wmsLogisticsStoragefeeMiApi *WmsLogisticsStoragefeeMiApi) DeleteWmsLogisticsStoragefeeMiByIds(c *gin.Context) {
	var IDS request.IdsReq
    err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    deletedBy := utils.GetUserID(c)
	if err := wmsLogisticsStoragefeeMiService.DeleteWmsLogisticsStoragefeeMiByIds(IDS,deletedBy); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateWmsLogisticsStoragefeeMi 更新仓储费
// @Tags WmsLogisticsStoragefeeMi
// @Summary 更新仓储费
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body manual.WmsLogisticsStoragefeeMi true "更新仓储费"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /wmsLogisticsStoragefeeMi/updateWmsLogisticsStoragefeeMi [put]
func (wmsLogisticsStoragefeeMiApi *WmsLogisticsStoragefeeMiApi) UpdateWmsLogisticsStoragefeeMi(c *gin.Context) {
	var wmsLogisticsStoragefeeMi manual.WmsLogisticsStoragefeeMi
	err := c.ShouldBindJSON(&wmsLogisticsStoragefeeMi)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    wmsLogisticsStoragefeeMi.UpdatedBy = utils.GetUserID(c)
	if err := wmsLogisticsStoragefeeMiService.UpdateWmsLogisticsStoragefeeMi(wmsLogisticsStoragefeeMi); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindWmsLogisticsStoragefeeMi 用id查询仓储费
// @Tags WmsLogisticsStoragefeeMi
// @Summary 用id查询仓储费
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query manual.WmsLogisticsStoragefeeMi true "用id查询仓储费"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /wmsLogisticsStoragefeeMi/findWmsLogisticsStoragefeeMi [get]
func (wmsLogisticsStoragefeeMiApi *WmsLogisticsStoragefeeMiApi) FindWmsLogisticsStoragefeeMi(c *gin.Context) {
	var wmsLogisticsStoragefeeMi manual.WmsLogisticsStoragefeeMi
	err := c.ShouldBindQuery(&wmsLogisticsStoragefeeMi)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if rewmsLogisticsStoragefeeMi, err := wmsLogisticsStoragefeeMiService.GetWmsLogisticsStoragefeeMi(wmsLogisticsStoragefeeMi.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rewmsLogisticsStoragefeeMi": rewmsLogisticsStoragefeeMi}, c)
	}
}

// GetWmsLogisticsStoragefeeMiList 分页获取仓储费列表
// @Tags WmsLogisticsStoragefeeMi
// @Summary 分页获取仓储费列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query manualReq.WmsLogisticsStoragefeeMiSearch true "分页获取仓储费列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /wmsLogisticsStoragefeeMi/getWmsLogisticsStoragefeeMiList [get]
func (wmsLogisticsStoragefeeMiApi *WmsLogisticsStoragefeeMiApi) GetWmsLogisticsStoragefeeMiList(c *gin.Context) {
	var pageInfo manualReq.WmsLogisticsStoragefeeMiSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := wmsLogisticsStoragefeeMiService.GetWmsLogisticsStoragefeeMiInfoList(pageInfo); err != nil {
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
