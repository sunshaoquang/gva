import service from '@/utils/request'

// @Tags WmsKpiSummary2Mi
// @Summary 创建wmsKpiSummary2Mi表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.WmsKpiSummary2Mi true "创建wmsKpiSummary2Mi表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /wmsKpiSummary2Mi/createWmsKpiSummary2Mi [post]
export const createWmsKpiSummary2Mi = (data) => {
  return service({
    url: '/wmsKpiSummary2Mi/createWmsKpiSummary2Mi',
    method: 'post',
    data
  })
}

// @Tags WmsKpiSummary2Mi
// @Summary 删除wmsKpiSummary2Mi表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.WmsKpiSummary2Mi true "删除wmsKpiSummary2Mi表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /wmsKpiSummary2Mi/deleteWmsKpiSummary2Mi [delete]
export const deleteWmsKpiSummary2Mi = (data) => {
  return service({
    url: '/wmsKpiSummary2Mi/deleteWmsKpiSummary2Mi',
    method: 'delete',
    data
  })
}

// @Tags WmsKpiSummary2Mi
// @Summary 批量删除wmsKpiSummary2Mi表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除wmsKpiSummary2Mi表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /wmsKpiSummary2Mi/deleteWmsKpiSummary2Mi [delete]
export const deleteWmsKpiSummary2MiByIds = (data) => {
  return service({
    url: '/wmsKpiSummary2Mi/deleteWmsKpiSummary2MiByIds',
    method: 'delete',
    data
  })
}

// @Tags WmsKpiSummary2Mi
// @Summary 更新wmsKpiSummary2Mi表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.WmsKpiSummary2Mi true "更新wmsKpiSummary2Mi表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /wmsKpiSummary2Mi/updateWmsKpiSummary2Mi [put]
export const updateWmsKpiSummary2Mi = (data) => {
  return service({
    url: '/wmsKpiSummary2Mi/updateWmsKpiSummary2Mi',
    method: 'put',
    data
  })
}

// @Tags WmsKpiSummary2Mi
// @Summary 用id查询wmsKpiSummary2Mi表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.WmsKpiSummary2Mi true "用id查询wmsKpiSummary2Mi表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /wmsKpiSummary2Mi/findWmsKpiSummary2Mi [get]
export const findWmsKpiSummary2Mi = (params) => {
  return service({
    url: '/wmsKpiSummary2Mi/findWmsKpiSummary2Mi',
    method: 'get',
    params
  })
}

// @Tags WmsKpiSummary2Mi
// @Summary 分页获取wmsKpiSummary2Mi表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取wmsKpiSummary2Mi表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /wmsKpiSummary2Mi/getWmsKpiSummary2MiList [get]
export const getWmsKpiSummary2MiList = (params) => {
  return service({
    url: '/wmsKpiSummary2Mi/getWmsKpiSummary2MiList',
    method: 'get',
    params
  })
}
