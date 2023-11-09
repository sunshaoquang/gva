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

type WmsKpiSummary2MiApi struct {
}

var wmsKpiSummary2MiService = service.ServiceGroupApp.ManualServiceGroup.WmsKpiSummary2MiService


// CreateWmsKpiSummary2Mi 创建wmsKpiSummary2Mi表
// @Tags WmsKpiSummary2Mi
// @Summary 创建wmsKpiSummary2Mi表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body manual.WmsKpiSummary2Mi true "创建wmsKpiSummary2Mi表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /wmsKpiSummary2Mi/createWmsKpiSummary2Mi [post]
func (wmsKpiSummary2MiApi *WmsKpiSummary2MiApi) CreateWmsKpiSummary2Mi(c *gin.Context) {
	var wmsKpiSummary2Mi manual.WmsKpiSummary2Mi
	err := c.ShouldBindJSON(&wmsKpiSummary2Mi)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	// =============== 新增 ==================
	if wmsKpiSummary2Mi.CreatedName == "" {
		wmsKpiSummary2Mi.CreatedName = utils.GetUserName(c)
	}
	// =============== 新增 ==================
    wmsKpiSummary2Mi.CreatedBy = utils.GetUserID(c)
    verify := utils.Rules{
        "project":{utils.NotEmpty()},
    }
	if err := utils.Verify(wmsKpiSummary2Mi, verify); err != nil {
    		response.FailWithMessage(err.Error(), c)
    		return
    	}
	if err := wmsKpiSummary2MiService.CreateWmsKpiSummary2Mi(&wmsKpiSummary2Mi); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteWmsKpiSummary2Mi 删除wmsKpiSummary2Mi表
// @Tags WmsKpiSummary2Mi
// @Summary 删除wmsKpiSummary2Mi表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body manual.WmsKpiSummary2Mi true "删除wmsKpiSummary2Mi表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /wmsKpiSummary2Mi/deleteWmsKpiSummary2Mi [delete]
func (wmsKpiSummary2MiApi *WmsKpiSummary2MiApi) DeleteWmsKpiSummary2Mi(c *gin.Context) {
	var wmsKpiSummary2Mi manual.WmsKpiSummary2Mi
	err := c.ShouldBindJSON(&wmsKpiSummary2Mi)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    wmsKpiSummary2Mi.DeletedBy = utils.GetUserID(c)
	if err := wmsKpiSummary2MiService.DeleteWmsKpiSummary2Mi(wmsKpiSummary2Mi); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteWmsKpiSummary2MiByIds 批量删除wmsKpiSummary2Mi表
// @Tags WmsKpiSummary2Mi
// @Summary 批量删除wmsKpiSummary2Mi表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除wmsKpiSummary2Mi表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /wmsKpiSummary2Mi/deleteWmsKpiSummary2MiByIds [delete]
func (wmsKpiSummary2MiApi *WmsKpiSummary2MiApi) DeleteWmsKpiSummary2MiByIds(c *gin.Context) {
	var IDS request.IdsReq
    err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    deletedBy := utils.GetUserID(c)
	if err := wmsKpiSummary2MiService.DeleteWmsKpiSummary2MiByIds(IDS,deletedBy); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateWmsKpiSummary2Mi 更新wmsKpiSummary2Mi表
// @Tags WmsKpiSummary2Mi
// @Summary 更新wmsKpiSummary2Mi表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body manual.WmsKpiSummary2Mi true "更新wmsKpiSummary2Mi表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /wmsKpiSummary2Mi/updateWmsKpiSummary2Mi [put]
func (wmsKpiSummary2MiApi *WmsKpiSummary2MiApi) UpdateWmsKpiSummary2Mi(c *gin.Context) {
	var wmsKpiSummary2Mi manual.WmsKpiSummary2Mi
	err := c.ShouldBindJSON(&wmsKpiSummary2Mi)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

    wmsKpiSummary2Mi.UpdatedBy = utils.GetUserID(c)
      verify := utils.Rules{
          "project":{utils.NotEmpty()},
      }
    if err := utils.Verify(wmsKpiSummary2Mi, verify); err != nil {
      	response.FailWithMessage(err.Error(), c)
      	return
     }
	if err := wmsKpiSummary2MiService.UpdateWmsKpiSummary2Mi(wmsKpiSummary2Mi); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindWmsKpiSummary2Mi 用id查询wmsKpiSummary2Mi表
// @Tags WmsKpiSummary2Mi
// @Summary 用id查询wmsKpiSummary2Mi表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query manual.WmsKpiSummary2Mi true "用id查询wmsKpiSummary2Mi表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /wmsKpiSummary2Mi/findWmsKpiSummary2Mi [get]
func (wmsKpiSummary2MiApi *WmsKpiSummary2MiApi) FindWmsKpiSummary2Mi(c *gin.Context) {
	var wmsKpiSummary2Mi manual.WmsKpiSummary2Mi
	err := c.ShouldBindQuery(&wmsKpiSummary2Mi)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if rewmsKpiSummary2Mi, err := wmsKpiSummary2MiService.GetWmsKpiSummary2Mi(wmsKpiSummary2Mi.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rewmsKpiSummary2Mi": rewmsKpiSummary2Mi}, c)
	}
}

// GetWmsKpiSummary2MiList 分页获取wmsKpiSummary2Mi表列表
// @Tags WmsKpiSummary2Mi
// @Summary 分页获取wmsKpiSummary2Mi表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query manualReq.WmsKpiSummary2MiSearch true "分页获取wmsKpiSummary2Mi表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /wmsKpiSummary2Mi/getWmsKpiSummary2MiList [get]
func (wmsKpiSummary2MiApi *WmsKpiSummary2MiApi) GetWmsKpiSummary2MiList(c *gin.Context) {
	var pageInfo manualReq.WmsKpiSummary2MiSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := wmsKpiSummary2MiService.GetWmsKpiSummary2MiInfoList(pageInfo); err != nil {
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
