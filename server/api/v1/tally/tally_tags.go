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

type TallyTagsApi struct{}

var tallyTagsService = service.ServiceGroupApp.TallyServiceGroup.TallyTagsService

// CreateTallyTags 创建标签
// @Tags TallyTags
// @Summary 创建标签
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body tally.TallyTags true "创建标签"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /tallyTags/createTallyTags [post]
func (tallyTagsApi *TallyTagsApi) CreateTallyTags(c *gin.Context) {
	var tallyTags tally.TallyTags
	err := c.ShouldBindJSON(&tallyTags)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	tallyTags.CreatedBy = utils.GetUserID(c)
	tallyTags.UserId = utils.GetUserID(c)

	if err := tallyTagsService.CreateTallyTags(&tallyTags); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteTallyTags 删除标签
// @Tags TallyTags
// @Summary 删除标签
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body tally.TallyTags true "删除标签"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /tallyTags/deleteTallyTags [delete]
func (tallyTagsApi *TallyTagsApi) DeleteTallyTags(c *gin.Context) {
	id := c.Query("id")
	userID := utils.GetUserID(c)
	if err := tallyTagsService.DeleteTallyTags(id, userID); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteTallyTagsByIds 批量删除标签
// @Tags TallyTags
// @Summary 批量删除标签
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /tallyTags/deleteTallyTagsByIds [delete]
func (tallyTagsApi *TallyTagsApi) DeleteTallyTagsByIds(c *gin.Context) {
	ids := c.QueryArray("ids[]")
	userID := utils.GetUserID(c)
	if err := tallyTagsService.DeleteTallyTagsByIds(ids, userID); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateTallyTags 更新标签
// @Tags TallyTags
// @Summary 更新标签
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body tally.TallyTags true "更新标签"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /tallyTags/updateTallyTags [put]
func (tallyTagsApi *TallyTagsApi) UpdateTallyTags(c *gin.Context) {
	var tallyTags tally.TallyTags
	err := c.ShouldBindJSON(&tallyTags)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	tallyTags.UpdatedBy = utils.GetUserID(c)
	tallyTags.UserId = utils.GetUserID(c)

	if err := tallyTagsService.UpdateTallyTags(tallyTags); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindTallyTags 用id查询标签
// @Tags TallyTags
// @Summary 用id查询标签
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query tally.TallyTags true "用id查询标签"
// @Success 200 {object} response.Response{data=object{retallyTags=tally.TallyTags},msg=string} "查询成功"
// @Router /tallyTags/findTallyTags [get]
func (tallyTagsApi *TallyTagsApi) FindTallyTags(c *gin.Context) {
	id := c.Query("id")
	if retallyTags, err := tallyTagsService.GetTallyTags(id); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(retallyTags, c)
	}
}

// GetTallyTagsList 分页获取标签列表
// @Tags TallyTags
// @Summary 分页获取标签列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query tallyReq.TallyTagsSearch true "分页获取标签列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /tallyTags/getTallyTagsList [get]
func (tallyTagsApi *TallyTagsApi) GetTallyTagsList(c *gin.Context) {
	var pageInfo tallyReq.TallyTagsSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := tallyTagsService.GetTallyTagsInfoList(pageInfo); err != nil {
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

// GetTallyTagsListByUser 获取用户的标签列表
// @Tags TallyTags
// @Summary 获取用户的标签列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /tallyTags/getTallyTagsListByUser [get]
func (tallyTagsApi *TallyTagsApi) GetTallyTagsListByUser(c *gin.Context) {

	UserId := utils.GetUserID(c)
	if list, total, err := tallyTagsService.GetTallyTagsUserInfoList(UserId); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:  list,
			Total: total,
		}, "获取成功", c)
	}
}

// GetTallyTagsPublic 不需要鉴权的标签接口
// @Tags TallyTags
// @Summary 不需要鉴权的标签接口
// @accept application/json
// @Produce application/json
// @Param data query tallyReq.TallyTagsSearch true "分页获取标签列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /tallyTags/getTallyTagsPublic [get]
func (tallyTagsApi *TallyTagsApi) GetTallyTagsPublic(c *gin.Context) {
	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的标签接口信息",
	}, "获取成功", c)
}
