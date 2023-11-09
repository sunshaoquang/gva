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

type PromotionMarketTargetApi struct {
}

var MarketTargetService = service.ServiceGroupApp.ManualServiceGroup.PromotionMarketTargetService


// CreatePromotionMarketTarget 创建大促推广目标表
// @Tags PromotionMarketTarget
// @Summary 创建大促推广目标表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body manual.PromotionMarketTarget true "创建大促推广目标表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /MarketTarget/createPromotionMarketTarget [post]
func (MarketTargetApi *PromotionMarketTargetApi) CreatePromotionMarketTarget(c *gin.Context) {
	var MarketTarget manual.PromotionMarketTarget
	err := c.ShouldBindJSON(&MarketTarget)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	// =============== 新增 ==================
	if MarketTarget.CreatedName == "" {
		MarketTarget.CreatedName = utils.GetUserName(c)
	}
	// =============== 新增 ==================

    MarketTarget.CreatedBy = utils.GetUserID(c)
    verify := utils.Rules{
        "ShopName":{utils.NotEmpty()},
        "Type":{utils.NotEmpty()},
    }
	if err := utils.Verify(MarketTarget, verify); err != nil {
    		response.FailWithMessage(err.Error(), c)
    		return
    	}
	if err := MarketTargetService.CreatePromotionMarketTarget(&MarketTarget); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeletePromotionMarketTarget 删除大促推广目标表
// @Tags PromotionMarketTarget
// @Summary 删除大促推广目标表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body manual.PromotionMarketTarget true "删除大促推广目标表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /MarketTarget/deletePromotionMarketTarget [delete]
func (MarketTargetApi *PromotionMarketTargetApi) DeletePromotionMarketTarget(c *gin.Context) {
	var MarketTarget manual.PromotionMarketTarget
	err := c.ShouldBindJSON(&MarketTarget)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    MarketTarget.DeletedBy = utils.GetUserID(c)
	if err := MarketTargetService.DeletePromotionMarketTarget(MarketTarget); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeletePromotionMarketTargetByIds 批量删除大促推广目标表
// @Tags PromotionMarketTarget
// @Summary 批量删除大促推广目标表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除大促推广目标表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /MarketTarget/deletePromotionMarketTargetByIds [delete]
func (MarketTargetApi *PromotionMarketTargetApi) DeletePromotionMarketTargetByIds(c *gin.Context) {
	var IDS request.IdsReq
    err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    deletedBy := utils.GetUserID(c)
	if err := MarketTargetService.DeletePromotionMarketTargetByIds(IDS,deletedBy); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdatePromotionMarketTarget 更新大促推广目标表
// @Tags PromotionMarketTarget
// @Summary 更新大促推广目标表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body manual.PromotionMarketTarget true "更新大促推广目标表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /MarketTarget/updatePromotionMarketTarget [put]
func (MarketTargetApi *PromotionMarketTargetApi) UpdatePromotionMarketTarget(c *gin.Context) {
	var MarketTarget manual.PromotionMarketTarget
	err := c.ShouldBindJSON(&MarketTarget)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    MarketTarget.UpdatedBy = utils.GetUserID(c)
	
      verify := utils.Rules{
          "ShopName":{utils.NotEmpty()},
          "Type":{utils.NotEmpty()},
      }
    if err := utils.Verify(MarketTarget, verify); err != nil {
      	response.FailWithMessage(err.Error(), c)
      	return
     }
	if err := MarketTargetService.UpdatePromotionMarketTarget(MarketTarget); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindPromotionMarketTarget 用id查询大促推广目标表
// @Tags PromotionMarketTarget
// @Summary 用id查询大促推广目标表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query manual.PromotionMarketTarget true "用id查询大促推广目标表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /MarketTarget/findPromotionMarketTarget [get]
func (MarketTargetApi *PromotionMarketTargetApi) FindPromotionMarketTarget(c *gin.Context) {
	var MarketTarget manual.PromotionMarketTarget
	err := c.ShouldBindQuery(&MarketTarget)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if reMarketTarget, err := MarketTargetService.GetPromotionMarketTarget(MarketTarget.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reMarketTarget": reMarketTarget}, c)
	}
}

// GetPromotionMarketTargetList 分页获取大促推广目标表列表
// @Tags PromotionMarketTarget
// @Summary 分页获取大促推广目标表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query manualReq.PromotionMarketTargetSearch true "分页获取大促推广目标表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /MarketTarget/getPromotionMarketTargetList [get]
func (MarketTargetApi *PromotionMarketTargetApi) GetPromotionMarketTargetList(c *gin.Context) {
	var pageInfo manualReq.PromotionMarketTargetSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := MarketTargetService.GetPromotionMarketTargetInfoList(pageInfo); err != nil {
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
