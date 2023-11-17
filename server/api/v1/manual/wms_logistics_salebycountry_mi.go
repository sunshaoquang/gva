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

type WmsLogisticsSalebycountryMiApi struct {
}

var wmsLogisticsSalebycountryMiService = service.ServiceGroupApp.ManualServiceGroup.WmsLogisticsSalebycountryMiService


// CreateWmsLogisticsSalebycountryMi 创建收入by国家
// @Tags WmsLogisticsSalebycountryMi
// @Summary 创建收入by国家
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body manual.WmsLogisticsSalebycountryMi true "创建收入by国家"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /wmsLogisticsSalebycountryMi/createWmsLogisticsSalebycountryMi [post]
func (wmsLogisticsSalebycountryMiApi *WmsLogisticsSalebycountryMiApi) CreateWmsLogisticsSalebycountryMi(c *gin.Context) {
	var wmsLogisticsSalebycountryMi manual.WmsLogisticsSalebycountryMi
	err := c.ShouldBindJSON(&wmsLogisticsSalebycountryMi)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	// =============== 新增 ==================
	if wmsLogisticsSalebycountryMi.CreatedName == "" {
		wmsLogisticsSalebycountryMi.CreatedName = utils.GetUserName(c)
	}
	// =============== 新增 ==================
    wmsLogisticsSalebycountryMi.CreatedBy = utils.GetUserID(c)
	if err := wmsLogisticsSalebycountryMiService.CreateWmsLogisticsSalebycountryMi(&wmsLogisticsSalebycountryMi); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteWmsLogisticsSalebycountryMi 删除收入by国家
// @Tags WmsLogisticsSalebycountryMi
// @Summary 删除收入by国家
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body manual.WmsLogisticsSalebycountryMi true "删除收入by国家"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /wmsLogisticsSalebycountryMi/deleteWmsLogisticsSalebycountryMi [delete]
func (wmsLogisticsSalebycountryMiApi *WmsLogisticsSalebycountryMiApi) DeleteWmsLogisticsSalebycountryMi(c *gin.Context) {
	var wmsLogisticsSalebycountryMi manual.WmsLogisticsSalebycountryMi
	err := c.ShouldBindJSON(&wmsLogisticsSalebycountryMi)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    wmsLogisticsSalebycountryMi.DeletedBy = utils.GetUserID(c)
	if err := wmsLogisticsSalebycountryMiService.DeleteWmsLogisticsSalebycountryMi(wmsLogisticsSalebycountryMi); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteWmsLogisticsSalebycountryMiAll 删除所有收入by国家
// @Tags WmsLogisticsSalebycountryMiAll
// @Summary 删除所有收入by国家
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body manual.WmsLogisticsSalebycountryMi true "删除所有收入by国家"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /wmsLogisticsSalebycountryMi/deleteWmsLogisticsSalebycountryMiAll [delete]
func (wmsLogisticsSalebycountryMiApi *WmsLogisticsSalebycountryMiApi) DeleteWmsLogisticsSalebycountryMiAll(c *gin.Context) {
	var wmsLogisticsSalebycountryMi manual.WmsLogisticsSalebycountryMi
	
	if err := global.MustGetGlobalDBByDBName("ht_smartdata").Unscoped().Where("1 = 1").Delete(&wmsLogisticsSalebycountryMi).Error; err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteWmsLogisticsSalebycountryMiByIds 批量删除收入by国家
// @Tags WmsLogisticsSalebycountryMi
// @Summary 批量删除收入by国家
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除收入by国家"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /wmsLogisticsSalebycountryMi/deleteWmsLogisticsSalebycountryMiByIds [delete]
func (wmsLogisticsSalebycountryMiApi *WmsLogisticsSalebycountryMiApi) DeleteWmsLogisticsSalebycountryMiByIds(c *gin.Context) {
	var IDS request.IdsReq
    err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    deletedBy := utils.GetUserID(c)
	if err := wmsLogisticsSalebycountryMiService.DeleteWmsLogisticsSalebycountryMiByIds(IDS,deletedBy); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateWmsLogisticsSalebycountryMi 更新收入by国家
// @Tags WmsLogisticsSalebycountryMi
// @Summary 更新收入by国家
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body manual.WmsLogisticsSalebycountryMi true "更新收入by国家"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /wmsLogisticsSalebycountryMi/updateWmsLogisticsSalebycountryMi [put]
func (wmsLogisticsSalebycountryMiApi *WmsLogisticsSalebycountryMiApi) UpdateWmsLogisticsSalebycountryMi(c *gin.Context) {
	var wmsLogisticsSalebycountryMi manual.WmsLogisticsSalebycountryMi
	err := c.ShouldBindJSON(&wmsLogisticsSalebycountryMi)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    wmsLogisticsSalebycountryMi.UpdatedBy = utils.GetUserID(c)
	if err := wmsLogisticsSalebycountryMiService.UpdateWmsLogisticsSalebycountryMi(wmsLogisticsSalebycountryMi); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindWmsLogisticsSalebycountryMi 用id查询收入by国家
// @Tags WmsLogisticsSalebycountryMi
// @Summary 用id查询收入by国家
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query manual.WmsLogisticsSalebycountryMi true "用id查询收入by国家"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /wmsLogisticsSalebycountryMi/findWmsLogisticsSalebycountryMi [get]
func (wmsLogisticsSalebycountryMiApi *WmsLogisticsSalebycountryMiApi) FindWmsLogisticsSalebycountryMi(c *gin.Context) {
	var wmsLogisticsSalebycountryMi manual.WmsLogisticsSalebycountryMi
	err := c.ShouldBindQuery(&wmsLogisticsSalebycountryMi)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if rewmsLogisticsSalebycountryMi, err := wmsLogisticsSalebycountryMiService.GetWmsLogisticsSalebycountryMi(wmsLogisticsSalebycountryMi.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rewmsLogisticsSalebycountryMi": rewmsLogisticsSalebycountryMi}, c)
	}
}

// GetWmsLogisticsSalebycountryMiList 分页获取收入by国家列表
// @Tags WmsLogisticsSalebycountryMi
// @Summary 分页获取收入by国家列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query manualReq.WmsLogisticsSalebycountryMiSearch true "分页获取收入by国家列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /wmsLogisticsSalebycountryMi/getWmsLogisticsSalebycountryMiList [get]
func (wmsLogisticsSalebycountryMiApi *WmsLogisticsSalebycountryMiApi) GetWmsLogisticsSalebycountryMiList(c *gin.Context) {
	var pageInfo manualReq.WmsLogisticsSalebycountryMiSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := wmsLogisticsSalebycountryMiService.GetWmsLogisticsSalebycountryMiInfoList(pageInfo); err != nil {
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
