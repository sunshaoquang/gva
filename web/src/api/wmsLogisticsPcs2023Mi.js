import service from '@/utils/request'

// @Tags WmsLogisticsPcs2023Mi
// @Summary 创建2023年物流成本明细表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.WmsLogisticsPcs2023Mi true "创建2023年物流成本明细表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /wmsLogisticsPcs2023Mi/createWmsLogisticsPcs2023Mi [post]
export const createWmsLogisticsPcs2023Mi = (data) => {
  return service({
    url: '/wmsLogisticsPcs2023Mi/createWmsLogisticsPcs2023Mi',
    method: 'post',
    data
  })
}

// @Tags WmsLogisticsPcs2023Mi
// @Summary 删除2023年物流成本明细表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.WmsLogisticsPcs2023Mi true "删除2023年物流成本明细表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /wmsLogisticsPcs2023Mi/deleteWmsLogisticsPcs2023Mi [delete]
export const deleteWmsLogisticsPcs2023Mi = (data) => {
  return service({
    url: '/wmsLogisticsPcs2023Mi/deleteWmsLogisticsPcs2023Mi',
    method: 'delete',
    data
  })
}

// @Tags WmsLogisticsPcs2023MiAll
// @Summary 删除2023年物流成本明细表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.WmsLogisticsPcs2023MiAll true "删除2023年物流成本明细表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /wmsLogisticsPcs2023Mi/deleteWmsLogisticsPcs2023MiAll [delete]
export const deleteWmsLogisticsPcs2023MiAll = () => {
  return service({
    url: '/wmsLogisticsPcs2023Mi/deleteWmsLogisticsPcs2023MiAll',
    method: 'delete'
  })
}

// @Tags WmsLogisticsPcs2023Mi
// @Summary 批量删除2023年物流成本明细表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除2023年物流成本明细表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /wmsLogisticsPcs2023Mi/deleteWmsLogisticsPcs2023Mi [delete]
export const deleteWmsLogisticsPcs2023MiByIds = (data) => {
  return service({
    url: '/wmsLogisticsPcs2023Mi/deleteWmsLogisticsPcs2023MiByIds',
    method: 'delete',
    data
  })
}

// @Tags WmsLogisticsPcs2023Mi
// @Summary 更新2023年物流成本明细表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.WmsLogisticsPcs2023Mi true "更新2023年物流成本明细表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /wmsLogisticsPcs2023Mi/updateWmsLogisticsPcs2023Mi [put]
export const updateWmsLogisticsPcs2023Mi = (data) => {
  return service({
    url: '/wmsLogisticsPcs2023Mi/updateWmsLogisticsPcs2023Mi',
    method: 'put',
    data
  })
}

// @Tags WmsLogisticsPcs2023Mi
// @Summary 用id查询2023年物流成本明细表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.WmsLogisticsPcs2023Mi true "用id查询2023年物流成本明细表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /wmsLogisticsPcs2023Mi/findWmsLogisticsPcs2023Mi [get]
export const findWmsLogisticsPcs2023Mi = (params) => {
  return service({
    url: '/wmsLogisticsPcs2023Mi/findWmsLogisticsPcs2023Mi',
    method: 'get',
    params
  })
}

// @Tags WmsLogisticsPcs2023Mi
// @Summary 分页获取2023年物流成本明细表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取2023年物流成本明细表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /wmsLogisticsPcs2023Mi/getWmsLogisticsPcs2023MiList [get]
export const getWmsLogisticsPcs2023MiList = (params) => {
  return service({
    url: '/wmsLogisticsPcs2023Mi/getWmsLogisticsPcs2023MiList',
    method: 'get',
    params
  })
}
