import service from '@/utils/request'

// @Tags WmsLogisticsPcsDetailMi
// @Summary 创建主要产品成本明细表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.WmsLogisticsPcsDetailMi true "创建主要产品成本明细表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /wmsLogisticsPcsDetailMi/createWmsLogisticsPcsDetailMi [post]
export const createWmsLogisticsPcsDetailMi = (data) => {
  return service({
    url: '/wmsLogisticsPcsDetailMi/createWmsLogisticsPcsDetailMi',
    method: 'post',
    data
  })
}

// @Tags WmsLogisticsPcsDetailMi
// @Summary 删除主要产品成本明细表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.WmsLogisticsPcsDetailMi true "删除主要产品成本明细表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /wmsLogisticsPcsDetailMi/deleteWmsLogisticsPcsDetailMi [delete]
export const deleteWmsLogisticsPcsDetailMi = (data) => {
  return service({
    url: '/wmsLogisticsPcsDetailMi/deleteWmsLogisticsPcsDetailMi',
    method: 'delete',
    data
  })
}

// @Tags WmsLogisticsPcsDetailMiAll
// @Summary 删除主要产品成本明细表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.WmsLogisticsPcsDetailMiAll true "删除主要产品成本明细表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /wmsLogisticsPcsDetailMi/deleteWmsLogisticsPcsDetailMiAll [delete]
export const deleteWmsLogisticsPcsDetailMiAll = () => {
  return service({
    url: '/wmsLogisticsPcsDetailMi/deleteWmsLogisticsPcsDetailMiAll',
    method: 'delete'
  })
}


// @Tags WmsLogisticsPcsDetailMi
// @Summary 批量删除主要产品成本明细表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除主要产品成本明细表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /wmsLogisticsPcsDetailMi/deleteWmsLogisticsPcsDetailMi [delete]
export const deleteWmsLogisticsPcsDetailMiByIds = (data) => {
  return service({
    url: '/wmsLogisticsPcsDetailMi/deleteWmsLogisticsPcsDetailMiByIds',
    method: 'delete',
    data
  })
}

// @Tags WmsLogisticsPcsDetailMi
// @Summary 更新主要产品成本明细表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.WmsLogisticsPcsDetailMi true "更新主要产品成本明细表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /wmsLogisticsPcsDetailMi/updateWmsLogisticsPcsDetailMi [put]
export const updateWmsLogisticsPcsDetailMi = (data) => {
  return service({
    url: '/wmsLogisticsPcsDetailMi/updateWmsLogisticsPcsDetailMi',
    method: 'put',
    data
  })
}

// @Tags WmsLogisticsPcsDetailMi
// @Summary 用id查询主要产品成本明细表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.WmsLogisticsPcsDetailMi true "用id查询主要产品成本明细表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /wmsLogisticsPcsDetailMi/findWmsLogisticsPcsDetailMi [get]
export const findWmsLogisticsPcsDetailMi = (params) => {
  return service({
    url: '/wmsLogisticsPcsDetailMi/findWmsLogisticsPcsDetailMi',
    method: 'get',
    params
  })
}

// @Tags WmsLogisticsPcsDetailMi
// @Summary 分页获取主要产品成本明细表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取主要产品成本明细表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /wmsLogisticsPcsDetailMi/getWmsLogisticsPcsDetailMiList [get]
export const getWmsLogisticsPcsDetailMiList = (params) => {
  return service({
    url: '/wmsLogisticsPcsDetailMi/getWmsLogisticsPcsDetailMiList',
    method: 'get',
    params
  })
}
