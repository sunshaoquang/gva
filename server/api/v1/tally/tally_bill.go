package tally

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/tally"
	tallyReq "github.com/flipped-aurora/gin-vue-admin/server/model/tally/request"
	tallyRes "github.com/flipped-aurora/gin-vue-admin/server/model/tally/response"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type TallyBillApi struct{}

var tallyBillService = service.ServiceGroupApp.TallyServiceGroup.TallyBillService

// CreateTallyBill 创建记账账单表
// @Tags TallyBill
// @Summary 创建记账账单表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body tally.TallyBill true "创建记账账单表"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /tallyBill/createTallyBill [post]
func (tallyBillApi *TallyBillApi) CreateTallyBill(c *gin.Context) {
	var tallyBill tally.TallyBill
	err := c.ShouldBindJSON(&tallyBill)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	tallyBill.CreatedBy = utils.GetUserID(c)
	tallyBill.UserId = utils.GetUserID(c)

	if err := tallyBillService.CreateTallyBill(&tallyBill); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteTallyBill 删除记账账单表
// @Tags TallyBill
// @Summary 删除记账账单表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body tally.TallyBill true "删除记账账单表"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /tallyBill/deleteTallyBill [delete]
func (tallyBillApi *TallyBillApi) DeleteTallyBill(c *gin.Context) {
	Id := c.Query("Id")
	userID := utils.GetUserID(c)
	if err := tallyBillService.DeleteTallyBill(Id, userID); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteTallyBillByIds 批量删除记账账单表
// @Tags TallyBill
// @Summary 批量删除记账账单表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /tallyBill/deleteTallyBillByIds [delete]
func (tallyBillApi *TallyBillApi) DeleteTallyBillByIds(c *gin.Context) {
	ids := c.QueryArray("ids[]")
	userID := utils.GetUserID(c)
	if err := tallyBillService.DeleteTallyBillByIds(ids, userID); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateTallyBill 更新记账账单表
// @Tags TallyBill
// @Summary 更新记账账单表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body tally.TallyBill true "更新记账账单表"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /tallyBill/updateTallyBill [put]
func (tallyBillApi *TallyBillApi) UpdateTallyBill(c *gin.Context) {
	var tallyBill tally.TallyBill
	err := c.ShouldBindJSON(&tallyBill)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	tallyBill.UpdatedBy = utils.GetUserID(c)

	if err := tallyBillService.UpdateTallyBill(tallyBill); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindTallyBill 用id查询记账账单表
// @Tags TallyBill
// @Summary 用id查询记账账单表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query tally.TallyBill true "用id查询记账账单表"
// @Success 200 {object} response.Response{data=object{retallyBill=tally.TallyBill},msg=string} "查询成功"
// @Router /tallyBill/findTallyBill [get]
func (tallyBillApi *TallyBillApi) FindTallyBill(c *gin.Context) {
	Id := c.Query("id")
	if retallyBill, err := tallyBillService.GetTallyBill(Id); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(retallyBill, c)
	}
}

// GetTallyBillList 分页获取记账账单表列表
// @Tags TallyBill
// @Summary 分页获取记账账单表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query tallyReq.TallyBillSearch true "分页获取记账账单表列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /tallyBill/getTallyBillList [get]
func (tallyBillApi *TallyBillApi) GetTallyBillList(c *gin.Context) {
	var pageInfo tallyReq.TallyBillSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := tallyBillService.GetTallyBillInfoList(pageInfo); err != nil {
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

// GetTallyBillDetailDataList 获取记账账单表详细列表接口
// @Tags TallyBill
// @Summary 获取记账账单表详细列表接口
// @accept application/json
// @Produce application/json
// @Param data query tallyReq.TallyBillDetailSearch true "获取记账账单表详细列表接口"
// @Success 200 {object} tallyRes.TallyBillDetailResult{data=object,msg=string} "获取成功"
// @Router /tallyBill/getTallyBillDetailDataList [get]
func (tallyBillApi *TallyBillApi) GetTallyBillDetailDataList(c *gin.Context) {
	var pageInfo tallyReq.TallyBillDetailSearch
	UserId := utils.GetUserID(c)
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, err := tallyBillService.GetTallyBillDetailDataInfoList(UserId, pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(tallyRes.TallyBillDetailResult{
			List: list,
		}, "获取成功", c)
	}
}

// GetTallyBillPublic 不需要鉴权的记账账单表接口
// @Tags TallyBill
// @Summary 不需要鉴权的记账账单表接口
// @accept application/json
// @Produce application/json
// @Param data query tallyReq.TallyBillSearch true "分页获取记账账单表列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /tallyBill/getTallyBillPublic [get]
func (tallyBillApi *TallyBillApi) GetTallyBillPublic(c *gin.Context) {
	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的记账账单表接口信息",
	}, "获取成功", c)
}
