import service from '@/utils/request'

// @Tags WmsLogisticsModeTransportMi
// @Summary 创建费用by运输方式
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.WmsLogisticsModeTransportMi true "创建费用by运输方式"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /wmsLogisticsModeTransportMi/createWmsLogisticsModeTransportMi [post]
export const createWmsLogisticsModeTransportMi = (data) => {
  return service({
    url: '/wmsLogisticsModeTransportMi/createWmsLogisticsModeTransportMi',
    method: 'post',
    data
  })
}

// @Tags WmsLogisticsModeTransportMi
// @Summary 删除费用by运输方式
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.WmsLogisticsModeTransportMi true "删除费用by运输方式"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /wmsLogisticsModeTransportMi/deleteWmsLogisticsModeTransportMi [delete]
export const deleteWmsLogisticsModeTransportMi = (data) => {
  return service({
    url: '/wmsLogisticsModeTransportMi/deleteWmsLogisticsModeTransportMi',
    method: 'delete',
    data
  })
}

// @Tags WmsLogisticsModeTransportMiAll
// @Summary 删除所有费用by运输方式
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.WmsLogisticsModeTransportMiAll true "删除所有费用by运输方式"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /wmsLogisticsModeTransportMi/deleteWmsLogisticsModeTransportMiAll [delete]
export const deleteWmsLogisticsModeTransportMiAll = () => {
  return service({
    url: '/wmsLogisticsModeTransportMi/deleteWmsLogisticsModeTransportMiAll',
    method: 'delete'
  })
}

// @Tags WmsLogisticsModeTransportMi
// @Summary 批量删除费用by运输方式
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除费用by运输方式"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /wmsLogisticsModeTransportMi/deleteWmsLogisticsModeTransportMi [delete]
export const deleteWmsLogisticsModeTransportMiByIds = (data) => {
  return service({
    url: '/wmsLogisticsModeTransportMi/deleteWmsLogisticsModeTransportMiByIds',
    method: 'delete',
    data
  })
}

// @Tags WmsLogisticsModeTransportMi
// @Summary 更新费用by运输方式
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.WmsLogisticsModeTransportMi true "更新费用by运输方式"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /wmsLogisticsModeTransportMi/updateWmsLogisticsModeTransportMi [put]
export const updateWmsLogisticsModeTransportMi = (data) => {
  return service({
    url: '/wmsLogisticsModeTransportMi/updateWmsLogisticsModeTransportMi',
    method: 'put',
    data
  })
}

// @Tags WmsLogisticsModeTransportMi
// @Summary 用id查询费用by运输方式
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.WmsLogisticsModeTransportMi true "用id查询费用by运输方式"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /wmsLogisticsModeTransportMi/findWmsLogisticsModeTransportMi [get]
export const findWmsLogisticsModeTransportMi = (params) => {
  return service({
    url: '/wmsLogisticsModeTransportMi/findWmsLogisticsModeTransportMi',
    method: 'get',
    params
  })
}

// @Tags WmsLogisticsModeTransportMi
// @Summary 分页获取费用by运输方式列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取费用by运输方式列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /wmsLogisticsModeTransportMi/getWmsLogisticsModeTransportMiList [get]
export const getWmsLogisticsModeTransportMiList = (params) => {
  return service({
    url: '/wmsLogisticsModeTransportMi/getWmsLogisticsModeTransportMiList',
    method: 'get',
    params
  })
}
