package tally

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/tally"
	tallyReq "github.com/flipped-aurora/gin-vue-admin/server/model/tally/request"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type TallyCategoryApi struct{}

var tallyCategoryService = service.ServiceGroupApp.TallyServiceGroup.TallyCategoryService

// CreateTallyCategory 创建记账类别表
// @Tags TallyCategory
// @Summary 创建记账类别表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body tally.TallyCategory true "创建记账类别表"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /tallyCategory/createTallyCategory [post]
func (tallyCategoryApi *TallyCategoryApi) CreateTallyCategory(c *gin.Context) {
	var tallyCategory tally.TallyCategory
	err := c.ShouldBindJSON(&tallyCategory)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	tallyCategory.CreatedBy = utils.GetUserID(c)

	if err := tallyCategoryService.CreateTallyCategory(&tallyCategory); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteTallyCategory 删除记账类别表
// @Tags TallyCategory
// @Summary 删除记账类别表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body tally.TallyCategory true "删除记账类别表"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /tallyCategory/deleteTallyCategory [delete]
func (tallyCategoryApi *TallyCategoryApi) DeleteTallyCategory(c *gin.Context) {
	id := c.Query("id")
	userID := utils.GetUserID(c)
	if err := tallyCategoryService.DeleteTallyCategory(id, userID); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteTallyCategoryByIds 批量删除记账类别表
// @Tags TallyCategory
// @Summary 批量删除记账类别表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /tallyCategory/deleteTallyCategoryByIds [delete]
func (tallyCategoryApi *TallyCategoryApi) DeleteTallyCategoryByIds(c *gin.Context) {
	ids := c.QueryArray("ids[]")
	userID := utils.GetUserID(c)
	if err := tallyCategoryService.DeleteTallyCategoryByIds(ids, userID); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateTallyCategory 更新记账类别表
// @Tags TallyCategory
// @Summary 更新记账类别表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body tally.TallyCategory true "更新记账类别表"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /tallyCategory/updateTallyCategory [put]
func (tallyCategoryApi *TallyCategoryApi) UpdateTallyCategory(c *gin.Context) {
	var tallyCategory tally.TallyCategory
	err := c.ShouldBindJSON(&tallyCategory)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	tallyCategory.UpdatedBy = utils.GetUserID(c)

	if err := tallyCategoryService.UpdateTallyCategory(tallyCategory); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindTallyCategory 用id查询记账类别表
// @Tags TallyCategory
// @Summary 用id查询记账类别表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query tally.TallyCategory true "用id查询记账类别表"
// @Success 200 {object} response.Response{data=object{retallyCategory=tally.TallyCategory},msg=string} "查询成功"
// @Router /tallyCategory/findTallyCategory [get]
func (tallyCategoryApi *TallyCategoryApi) FindTallyCategory(c *gin.Context) {
	id := c.Query("id")
	if retallyCategory, err := tallyCategoryService.GetTallyCategory(id); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(retallyCategory, c)
	}
}

// GetTallyCategoryList 分页获取记账类别表列表
// @Tags TallyCategory
// @Summary 分页获取记账类别表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query tallyReq.TallyCategorySearch true "分页获取记账类别表列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /tallyCategory/getTallyCategoryList [get]
func (tallyCategoryApi *TallyCategoryApi) GetTallyCategoryList(c *gin.Context) {
	var pageInfo tallyReq.TallyCategorySearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := tallyCategoryService.GetTallyCategoryInfoList(pageInfo); err != nil {
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

// GetTallyCategoryPublic 不需要鉴权的记账类别表接口
// @Tags TallyCategory
// @Summary 不需要鉴权的记账类别表接口
// @accept application/json
// @Produce application/json
// @Param data query tallyReq.TallyCategorySearch true "分页获取记账类别表列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /tallyCategory/getTallyCategoryPublic [get]
func (tallyCategoryApi *TallyCategoryApi) GetTallyCategoryPublic(c *gin.Context) {
	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的记账类别表接口信息",
	}, "获取成功", c)
}
