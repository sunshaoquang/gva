import service from '@/utils/request'

// @Tags WmsLogisticsRateMonths
// @Summary 创建wmsLogisticsRateMonths表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.WmsLogisticsRateMonths true "创建wmsLogisticsRateMonths表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /wmsLogisticsRateMonths/createWmsLogisticsRateMonths [post]
export const createWmsLogisticsRateMonths = (data) => {
  return service({
    url: '/wmsLogisticsRateMonths/createWmsLogisticsRateMonths',
    method: 'post',
    data
  })
}

// @Tags WmsLogisticsRateMonths
// @Summary 删除wmsLogisticsRateMonths表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.WmsLogisticsRateMonths true "删除wmsLogisticsRateMonths表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /wmsLogisticsRateMonths/deleteWmsLogisticsRateMonths [delete]
export const deleteWmsLogisticsRateMonths = (data) => {
  return service({
    url: '/wmsLogisticsRateMonths/deleteWmsLogisticsRateMonths',
    method: 'delete',
    data
  })
}

// @Tags WmsLogisticsRateMonthsAll
// @Summary 删除所有wmsLogisticsRateMonths表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.WmsLogisticsRateMonthsAll true "删除所有wmsLogisticsRateMonths表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /wmsLogisticsRateMonths/deleteWmsLogisticsRateMonthsAll [delete]
export const deleteWmsLogisticsRateMonthsAll = () => {
  return service({
    url: '/wmsLogisticsRateMonths/deleteWmsLogisticsRateMonthsAll',
    method: 'delete'
  })
}

// @Tags WmsLogisticsRateMonths
// @Summary 批量删除wmsLogisticsRateMonths表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除wmsLogisticsRateMonths表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /wmsLogisticsRateMonths/deleteWmsLogisticsRateMonths [delete]
export const deleteWmsLogisticsRateMonthsByIds = (data) => {
  return service({
    url: '/wmsLogisticsRateMonths/deleteWmsLogisticsRateMonthsByIds',
    method: 'delete',
    data
  })
}

// @Tags WmsLogisticsRateMonths
// @Summary 更新wmsLogisticsRateMonths表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.WmsLogisticsRateMonths true "更新wmsLogisticsRateMonths表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /wmsLogisticsRateMonths/updateWmsLogisticsRateMonths [put]
export const updateWmsLogisticsRateMonths = (data) => {
  return service({
    url: '/wmsLogisticsRateMonths/updateWmsLogisticsRateMonths',
    method: 'put',
    data
  })
}

// @Tags WmsLogisticsRateMonths
// @Summary 用id查询wmsLogisticsRateMonths表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.WmsLogisticsRateMonths true "用id查询wmsLogisticsRateMonths表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /wmsLogisticsRateMonths/findWmsLogisticsRateMonths [get]
export const findWmsLogisticsRateMonths = (params) => {
  return service({
    url: '/wmsLogisticsRateMonths/findWmsLogisticsRateMonths',
    method: 'get',
    params
  })
}

// @Tags WmsLogisticsRateMonths
// @Summary 分页获取wmsLogisticsRateMonths表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取wmsLogisticsRateMonths表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /wmsLogisticsRateMonths/getWmsLogisticsRateMonthsList [get]
export const getWmsLogisticsRateMonthsList = (params) => {
  return service({
    url: '/wmsLogisticsRateMonths/getWmsLogisticsRateMonthsList',
    method: 'get',
    params
  })
}
