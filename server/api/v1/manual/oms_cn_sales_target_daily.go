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

type OmsCnSalesTargetDailyApi struct {
}

var omsCnSalesTargetDailyService = service.ServiceGroupApp.ManualServiceGroup.OmsCnSalesTargetDailyService


// CreateOmsCnSalesTargetDaily 创建京东自营销售录入表
// @Tags OmsCnSalesTargetDaily
// @Summary 创建京东自营销售录入表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body manual.OmsCnSalesTargetDaily true "创建京东自营销售录入表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /omsCnSalesTargetDaily/createOmsCnSalesTargetDaily [post]
func (omsCnSalesTargetDailyApi *OmsCnSalesTargetDailyApi) CreateOmsCnSalesTargetDaily(c *gin.Context) {
	var omsCnSalesTargetDaily manual.OmsCnSalesTargetDaily
	err := c.ShouldBindJSON(&omsCnSalesTargetDaily)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	// =============== 新增 ==================
	if omsCnSalesTargetDaily.CreatedName == "" {
		omsCnSalesTargetDaily.CreatedName = utils.GetUserName(c)
	}
	// =============== 新增 ==================
    omsCnSalesTargetDaily.CreatedBy = utils.GetUserID(c)
	if err := omsCnSalesTargetDailyService.CreateOmsCnSalesTargetDaily(&omsCnSalesTargetDaily); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteOmsCnSalesTargetDaily 删除京东自营销售录入表
// @Tags OmsCnSalesTargetDaily
// @Summary 删除京东自营销售录入表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body manual.OmsCnSalesTargetDaily true "删除京东自营销售录入表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /omsCnSalesTargetDaily/deleteOmsCnSalesTargetDaily [delete]
func (omsCnSalesTargetDailyApi *OmsCnSalesTargetDailyApi) DeleteOmsCnSalesTargetDaily(c *gin.Context) {
	var omsCnSalesTargetDaily manual.OmsCnSalesTargetDaily
	err := c.ShouldBindJSON(&omsCnSalesTargetDaily)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    omsCnSalesTargetDaily.DeletedBy = utils.GetUserID(c)
	if err := omsCnSalesTargetDailyService.DeleteOmsCnSalesTargetDaily(omsCnSalesTargetDaily); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteOmsCnSalesTargetDailyAll 删除所有京东自营销售录入表
// @Tags OmsCnSalesTargetDailyAll
// @Summary 删除所有京东自营销售录入表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body manual.OmsCnSalesTargetDaily true "删除所有京东自营销售录入表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /omsCnSalesTargetDaily/deleteOmsCnSalesTargetDailyAll [delete]
func (omsCnSalesTargetDailyApi *OmsCnSalesTargetDailyApi) DeleteOmsCnSalesTargetDailyAll(c *gin.Context) {
	var omsCnSalesTargetDaily manual.OmsCnSalesTargetDaily
	
	if err := global.MustGetGlobalDBByDBName("ht_smartdata").Unscoped().Where("1 = 1").Delete(&omsCnSalesTargetDaily).Error; err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteOmsCnSalesTargetDailyByIds 批量删除京东自营销售录入表
// @Tags OmsCnSalesTargetDaily
// @Summary 批量删除京东自营销售录入表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除京东自营销售录入表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /omsCnSalesTargetDaily/deleteOmsCnSalesTargetDailyByIds [delete]
func (omsCnSalesTargetDailyApi *OmsCnSalesTargetDailyApi) DeleteOmsCnSalesTargetDailyByIds(c *gin.Context) {
	var IDS request.IdsReq
    err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    deletedBy := utils.GetUserID(c)
	if err := omsCnSalesTargetDailyService.DeleteOmsCnSalesTargetDailyByIds(IDS,deletedBy); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateOmsCnSalesTargetDaily 更新京东自营销售录入表
// @Tags OmsCnSalesTargetDaily
// @Summary 更新京东自营销售录入表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body manual.OmsCnSalesTargetDaily true "更新京东自营销售录入表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /omsCnSalesTargetDaily/updateOmsCnSalesTargetDaily [put]
func (omsCnSalesTargetDailyApi *OmsCnSalesTargetDailyApi) UpdateOmsCnSalesTargetDaily(c *gin.Context) {
	var omsCnSalesTargetDaily manual.OmsCnSalesTargetDaily
	err := c.ShouldBindJSON(&omsCnSalesTargetDaily)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    omsCnSalesTargetDaily.UpdatedBy = utils.GetUserID(c)
	if err := omsCnSalesTargetDailyService.UpdateOmsCnSalesTargetDaily(omsCnSalesTargetDaily); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindOmsCnSalesTargetDaily 用id查询京东自营销售录入表
// @Tags OmsCnSalesTargetDaily
// @Summary 用id查询京东自营销售录入表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query manual.OmsCnSalesTargetDaily true "用id查询京东自营销售录入表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /omsCnSalesTargetDaily/findOmsCnSalesTargetDaily [get]
func (omsCnSalesTargetDailyApi *OmsCnSalesTargetDailyApi) FindOmsCnSalesTargetDaily(c *gin.Context) {
	var omsCnSalesTargetDaily manual.OmsCnSalesTargetDaily
	err := c.ShouldBindQuery(&omsCnSalesTargetDaily)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if reomsCnSalesTargetDaily, err := omsCnSalesTargetDailyService.GetOmsCnSalesTargetDaily(omsCnSalesTargetDaily.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reomsCnSalesTargetDaily": reomsCnSalesTargetDaily}, c)
	}
}

// GetOmsCnSalesTargetDailyList 分页获取京东自营销售录入表列表
// @Tags OmsCnSalesTargetDaily
// @Summary 分页获取京东自营销售录入表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query manualReq.OmsCnSalesTargetDailySearch true "分页获取京东自营销售录入表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /omsCnSalesTargetDaily/getOmsCnSalesTargetDailyList [get]
func (omsCnSalesTargetDailyApi *OmsCnSalesTargetDailyApi) GetOmsCnSalesTargetDailyList(c *gin.Context) {
	var pageInfo manualReq.OmsCnSalesTargetDailySearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := omsCnSalesTargetDailyService.GetOmsCnSalesTargetDailyInfoList(pageInfo); err != nil {
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
