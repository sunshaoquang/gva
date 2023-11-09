import service from '@/utils/request'

// @Tags WmsKpi2023SummaryMi
// @Summary 创建wmsKpi2023SummaryMi表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.WmsKpi2023SummaryMi true "创建wmsKpi2023SummaryMi表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /wmsKpi2023SummaryMi/createWmsKpi2023SummaryMi [post]
export const createWmsKpi2023SummaryMi = (data) => {
  return service({
    url: '/wmsKpi2023SummaryMi/createWmsKpi2023SummaryMi',
    method: 'post',
    data
  })
}

// @Tags WmsKpi2023SummaryMi
// @Summary 删除wmsKpi2023SummaryMi表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.WmsKpi2023SummaryMi true "删除wmsKpi2023SummaryMi表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /wmsKpi2023SummaryMi/deleteWmsKpi2023SummaryMi [delete]
export const deleteWmsKpi2023SummaryMi = (data) => {
  return service({
    url: '/wmsKpi2023SummaryMi/deleteWmsKpi2023SummaryMi',
    method: 'delete',
    data
  })
}

// @Tags WmsKpi2023SummaryMi
// @Summary 批量删除wmsKpi2023SummaryMi表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除wmsKpi2023SummaryMi表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /wmsKpi2023SummaryMi/deleteWmsKpi2023SummaryMi [delete]
export const deleteWmsKpi2023SummaryMiByIds = (data) => {
  return service({
    url: '/wmsKpi2023SummaryMi/deleteWmsKpi2023SummaryMiByIds',
    method: 'delete',
    data
  })
}

// @Tags WmsKpi2023SummaryMi
// @Summary 更新wmsKpi2023SummaryMi表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.WmsKpi2023SummaryMi true "更新wmsKpi2023SummaryMi表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /wmsKpi2023SummaryMi/updateWmsKpi2023SummaryMi [put]
export const updateWmsKpi2023SummaryMi = (data) => {
  return service({
    url: '/wmsKpi2023SummaryMi/updateWmsKpi2023SummaryMi',
    method: 'put',
    data
  })
}

// @Tags WmsKpi2023SummaryMi
// @Summary 用id查询wmsKpi2023SummaryMi表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.WmsKpi2023SummaryMi true "用id查询wmsKpi2023SummaryMi表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /wmsKpi2023SummaryMi/findWmsKpi2023SummaryMi [get]
export const findWmsKpi2023SummaryMi = (params) => {
  return service({
    url: '/wmsKpi2023SummaryMi/findWmsKpi2023SummaryMi',
    method: 'get',
    params
  })
}

// @Tags WmsKpi2023SummaryMi
// @Summary 分页获取wmsKpi2023SummaryMi表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取wmsKpi2023SummaryMi表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /wmsKpi2023SummaryMi/getWmsKpi2023SummaryMiList [get]
export const getWmsKpi2023SummaryMiList = (params) => {
  return service({
    url: '/wmsKpi2023SummaryMi/getWmsKpi2023SummaryMiList',
    method: 'get',
    params
  })
}
