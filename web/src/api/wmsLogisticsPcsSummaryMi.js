import service from '@/utils/request'

// @Tags WmsLogisticsPcsSummaryMi
// @Summary 创建主要产品成本降本汇总表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.WmsLogisticsPcsSummaryMi true "创建主要产品成本降本汇总表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /wmsLogisticsPcsSummaryMi/createWmsLogisticsPcsSummaryMi [post]
export const createWmsLogisticsPcsSummaryMi = (data) => {
  return service({
    url: '/wmsLogisticsPcsSummaryMi/createWmsLogisticsPcsSummaryMi',
    method: 'post',
    data
  })
}

// @Tags WmsLogisticsPcsSummaryMi
// @Summary 删除主要产品成本降本汇总表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.WmsLogisticsPcsSummaryMi true "删除主要产品成本降本汇总表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /wmsLogisticsPcsSummaryMi/deleteWmsLogisticsPcsSummaryMi [delete]
export const deleteWmsLogisticsPcsSummaryMi = (data) => {
  return service({
    url: '/wmsLogisticsPcsSummaryMi/deleteWmsLogisticsPcsSummaryMi',
    method: 'delete',
    data
  })
}

// @Tags WmsLogisticsPcsSummaryMi
// @Summary 批量删除主要产品成本降本汇总表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除主要产品成本降本汇总表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /wmsLogisticsPcsSummaryMi/deleteWmsLogisticsPcsSummaryMi [delete]
export const deleteWmsLogisticsPcsSummaryMiByIds = (data) => {
  return service({
    url: '/wmsLogisticsPcsSummaryMi/deleteWmsLogisticsPcsSummaryMiByIds',
    method: 'delete',
    data
  })
}

// @Tags WmsLogisticsPcsSummaryMi
// @Summary 更新主要产品成本降本汇总表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.WmsLogisticsPcsSummaryMi true "更新主要产品成本降本汇总表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /wmsLogisticsPcsSummaryMi/updateWmsLogisticsPcsSummaryMi [put]
export const updateWmsLogisticsPcsSummaryMi = (data) => {
  return service({
    url: '/wmsLogisticsPcsSummaryMi/updateWmsLogisticsPcsSummaryMi',
    method: 'put',
    data
  })
}

// @Tags WmsLogisticsPcsSummaryMi
// @Summary 用id查询主要产品成本降本汇总表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.WmsLogisticsPcsSummaryMi true "用id查询主要产品成本降本汇总表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /wmsLogisticsPcsSummaryMi/findWmsLogisticsPcsSummaryMi [get]
export const findWmsLogisticsPcsSummaryMi = (params) => {
  return service({
    url: '/wmsLogisticsPcsSummaryMi/findWmsLogisticsPcsSummaryMi',
    method: 'get',
    params
  })
}

// @Tags WmsLogisticsPcsSummaryMi
// @Summary 分页获取主要产品成本降本汇总表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取主要产品成本降本汇总表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /wmsLogisticsPcsSummaryMi/getWmsLogisticsPcsSummaryMiList [get]
export const getWmsLogisticsPcsSummaryMiList = (params) => {
  return service({
    url: '/wmsLogisticsPcsSummaryMi/getWmsLogisticsPcsSummaryMiList',
    method: 'get',
    params
  })
}
