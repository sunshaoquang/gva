import service from '@/utils/request'

// @Tags WmsLogisticsProviderMi
// @Summary 创建费用by供应商
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.WmsLogisticsProviderMi true "创建费用by供应商"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /wmsLogisticsProviderMi/createWmsLogisticsProviderMi [post]
export const createWmsLogisticsProviderMi = (data) => {
  return service({
    url: '/wmsLogisticsProviderMi/createWmsLogisticsProviderMi',
    method: 'post',
    data
  })
}

// @Tags WmsLogisticsProviderMi
// @Summary 删除费用by供应商
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.WmsLogisticsProviderMi true "删除费用by供应商"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /wmsLogisticsProviderMi/deleteWmsLogisticsProviderMi [delete]
export const deleteWmsLogisticsProviderMi = (data) => {
  return service({
    url: '/wmsLogisticsProviderMi/deleteWmsLogisticsProviderMi',
    method: 'delete',
    data
  })
}

// @Tags WmsLogisticsProviderMiAll
// @Summary 删除所有费用by供应商
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.WmsLogisticsProviderMiAll true "删除所有费用by供应商"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /wmsLogisticsProviderMi/deleteWmsLogisticsProviderMiAll [delete]
export const deleteWmsLogisticsProviderMiAll = () => {
  return service({
    url: '/wmsLogisticsProviderMi/deleteWmsLogisticsProviderMiAll',
    method: 'delete'
  })
}

// @Tags WmsLogisticsProviderMi
// @Summary 批量删除费用by供应商
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除费用by供应商"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /wmsLogisticsProviderMi/deleteWmsLogisticsProviderMi [delete]
export const deleteWmsLogisticsProviderMiByIds = (data) => {
  return service({
    url: '/wmsLogisticsProviderMi/deleteWmsLogisticsProviderMiByIds',
    method: 'delete',
    data
  })
}

// @Tags WmsLogisticsProviderMi
// @Summary 更新费用by供应商
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.WmsLogisticsProviderMi true "更新费用by供应商"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /wmsLogisticsProviderMi/updateWmsLogisticsProviderMi [put]
export const updateWmsLogisticsProviderMi = (data) => {
  return service({
    url: '/wmsLogisticsProviderMi/updateWmsLogisticsProviderMi',
    method: 'put',
    data
  })
}

// @Tags WmsLogisticsProviderMi
// @Summary 用id查询费用by供应商
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.WmsLogisticsProviderMi true "用id查询费用by供应商"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /wmsLogisticsProviderMi/findWmsLogisticsProviderMi [get]
export const findWmsLogisticsProviderMi = (params) => {
  return service({
    url: '/wmsLogisticsProviderMi/findWmsLogisticsProviderMi',
    method: 'get',
    params
  })
}

// @Tags WmsLogisticsProviderMi
// @Summary 分页获取费用by供应商列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取费用by供应商列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /wmsLogisticsProviderMi/getWmsLogisticsProviderMiList [get]
export const getWmsLogisticsProviderMiList = (params) => {
  return service({
    url: '/wmsLogisticsProviderMi/getWmsLogisticsProviderMiList',
    method: 'get',
    params
  })
}
