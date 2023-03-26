package App

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/App"
	AppReq "github.com/flipped-aurora/gin-vue-admin/server/model/App/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
)

type AppCategoryApi struct {
}

var appCategoryService = service.ServiceGroupApp.AppServiceGroup.AppCategoryService

// CreateAppCategory 创建AppCategory
// @Tags AppCategory
// @Summary 创建AppCategory
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body App.AppCategory true "创建AppCategory"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /appCategory/createAppCategory [post]
func (appCategoryApi *AppCategoryApi) CreateAppCategory(c *gin.Context) {
	var appCategory App.AppCategory
	id, err := strconv.Atoi(c.Request.Header.Get("x-user-id"))
	if err != nil {
		response.FailWithMessage("创建失败", c)
	}

	err = c.ShouldBindJSON(&appCategory)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := appCategoryService.CreateAppCategory(appCategory, id); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteAppCategory 删除AppCategory
// @Tags AppCategory
// @Summary 删除AppCategory
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body App.AppCategory true "删除AppCategory"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /appCategory/deleteAppCategory [delete]
func (appCategoryApi *AppCategoryApi) DeleteAppCategory(c *gin.Context) {
	var appCategory App.AppCategory
	id, err := strconv.Atoi(c.Request.Header.Get("x-user-id"))
	if err != nil {
		response.FailWithMessage("创建失败", c)
	}
	err = c.ShouldBindJSON(&appCategory)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := appCategoryService.DeleteAppCategory(appCategory, id); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteAppCategoryByIds 批量删除AppCategory
// @Tags AppCategory
// @Summary 批量删除AppCategory
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除AppCategory"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /appCategory/deleteAppCategoryByIds [delete]
func (appCategoryApi *AppCategoryApi) DeleteAppCategoryByIds(c *gin.Context) {
	var IDS request.IdsReq
	id, err := strconv.Atoi(c.Request.Header.Get("x-user-id"))
	if err != nil {
		response.FailWithMessage("创建失败", c)
	}
	err = c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := appCategoryService.DeleteAppCategoryByIds(IDS, id); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateAppCategory 更新AppCategory
// @Tags AppCategory
// @Summary 更新AppCategory
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body App.AppCategory true "更新AppCategory"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /appCategory/updateAppCategory [put]
func (appCategoryApi *AppCategoryApi) UpdateAppCategory(c *gin.Context) {
	var appCategory App.AppCategory
	id, err := strconv.Atoi(c.Request.Header.Get("x-user-id"))
	if err != nil {
		response.FailWithMessage("创建失败", c)
	}
	err = c.ShouldBindJSON(&appCategory)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := appCategoryService.UpdateAppCategory(appCategory, id); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindAppCategory 用id查询AppCategory
// @Tags AppCategory
// @Summary 用id查询AppCategory
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query App.AppCategory true "用id查询AppCategory"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /appCategory/findAppCategory [get]
func (appCategoryApi *AppCategoryApi) FindAppCategory(c *gin.Context) {
	var appCategory App.AppCategory
	err := c.ShouldBindQuery(&appCategory)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if reappCategory, err := appCategoryService.GetAppCategory(appCategory.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reappCategory": reappCategory}, c)
	}
}

// GetAppCategoryList 分页获取AppCategory列表
// @Tags AppCategory
// @Summary 分页获取AppCategory列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query AppReq.AppCategorySearch true "分页获取AppCategory列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /appCategory/getAppCategoryList [get]
func (appCategoryApi *AppCategoryApi) GetAppCategoryList(c *gin.Context) {
	var pageInfo AppReq.AppCategorySearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := appCategoryService.GetAppCategoryInfoList(pageInfo); err != nil {
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
