package App

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/App"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    AppReq "github.com/flipped-aurora/gin-vue-admin/server/model/App/request"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/service"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

type AppBillApi struct {
}

var appBillService = service.ServiceGroupApp.AppServiceGroup.AppBillService


// CreateAppBill 创建AppBill
// @Tags AppBill
// @Summary 创建AppBill
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body App.AppBill true "创建AppBill"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /appBill/createAppBill [post]
func (appBillApi *AppBillApi) CreateAppBill(c *gin.Context) {
	var appBill App.AppBill
	err := c.ShouldBindJSON(&appBill)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := appBillService.CreateAppBill(appBill); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteAppBill 删除AppBill
// @Tags AppBill
// @Summary 删除AppBill
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body App.AppBill true "删除AppBill"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /appBill/deleteAppBill [delete]
func (appBillApi *AppBillApi) DeleteAppBill(c *gin.Context) {
	var appBill App.AppBill
	err := c.ShouldBindJSON(&appBill)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := appBillService.DeleteAppBill(appBill); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteAppBillByIds 批量删除AppBill
// @Tags AppBill
// @Summary 批量删除AppBill
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除AppBill"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /appBill/deleteAppBillByIds [delete]
func (appBillApi *AppBillApi) DeleteAppBillByIds(c *gin.Context) {
	var IDS request.IdsReq
    err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := appBillService.DeleteAppBillByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateAppBill 更新AppBill
// @Tags AppBill
// @Summary 更新AppBill
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body App.AppBill true "更新AppBill"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /appBill/updateAppBill [put]
func (appBillApi *AppBillApi) UpdateAppBill(c *gin.Context) {
	var appBill App.AppBill
	err := c.ShouldBindJSON(&appBill)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := appBillService.UpdateAppBill(appBill); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindAppBill 用id查询AppBill
// @Tags AppBill
// @Summary 用id查询AppBill
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query App.AppBill true "用id查询AppBill"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /appBill/findAppBill [get]
func (appBillApi *AppBillApi) FindAppBill(c *gin.Context) {
	var appBill App.AppBill
	err := c.ShouldBindQuery(&appBill)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if reappBill, err := appBillService.GetAppBill(appBill.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reappBill": reappBill}, c)
	}
}

// GetAppBillList 分页获取AppBill列表
// @Tags AppBill
// @Summary 分页获取AppBill列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query AppReq.AppBillSearch true "分页获取AppBill列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /appBill/getAppBillList [get]
func (appBillApi *AppBillApi) GetAppBillList(c *gin.Context) {
	var pageInfo AppReq.AppBillSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := appBillService.GetAppBillInfoList(pageInfo); err != nil {
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
