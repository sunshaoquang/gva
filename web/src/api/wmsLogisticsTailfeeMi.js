import service from '@/utils/request'

// @Tags WmsLogisticsTailfeeMi
// @Summary 创建尾程费用结构
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.WmsLogisticsTailfeeMi true "创建尾程费用结构"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /wmsLogisticsTailfeeMi/createWmsLogisticsTailfeeMi [post]
export const createWmsLogisticsTailfeeMi = (data) => {
  return service({
    url: '/wmsLogisticsTailfeeMi/createWmsLogisticsTailfeeMi',
    method: 'post',
    data
  })
}

// @Tags WmsLogisticsTailfeeMi
// @Summary 删除尾程费用结构
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.WmsLogisticsTailfeeMi true "删除尾程费用结构"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /wmsLogisticsTailfeeMi/deleteWmsLogisticsTailfeeMi [delete]
export const deleteWmsLogisticsTailfeeMi = (data) => {
  return service({
    url: '/wmsLogisticsTailfeeMi/deleteWmsLogisticsTailfeeMi',
    method: 'delete',
    data
  })
}

// @Tags WmsLogisticsTailfeeMiAll
// @Summary 删除所有尾程费用结构
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.WmsLogisticsTailfeeMiAll true "删除所有尾程费用结构"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /wmsLogisticsTailfeeMi/deleteWmsLogisticsTailfeeMiAll [delete]
export const deleteWmsLogisticsTailfeeMiAll = () => {
  return service({
    url: '/wmsLogisticsTailfeeMi/deleteWmsLogisticsTailfeeMiAll',
    method: 'delete'
  })
}

// @Tags WmsLogisticsTailfeeMi
// @Summary 批量删除尾程费用结构
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除尾程费用结构"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /wmsLogisticsTailfeeMi/deleteWmsLogisticsTailfeeMi [delete]
export const deleteWmsLogisticsTailfeeMiByIds = (data) => {
  return service({
    url: '/wmsLogisticsTailfeeMi/deleteWmsLogisticsTailfeeMiByIds',
    method: 'delete',
    data
  })
}

// @Tags WmsLogisticsTailfeeMi
// @Summary 更新尾程费用结构
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.WmsLogisticsTailfeeMi true "更新尾程费用结构"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /wmsLogisticsTailfeeMi/updateWmsLogisticsTailfeeMi [put]
export const updateWmsLogisticsTailfeeMi = (data) => {
  return service({
    url: '/wmsLogisticsTailfeeMi/updateWmsLogisticsTailfeeMi',
    method: 'put',
    data
  })
}

// @Tags WmsLogisticsTailfeeMi
// @Summary 用id查询尾程费用结构
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.WmsLogisticsTailfeeMi true "用id查询尾程费用结构"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /wmsLogisticsTailfeeMi/findWmsLogisticsTailfeeMi [get]
export const findWmsLogisticsTailfeeMi = (params) => {
  return service({
    url: '/wmsLogisticsTailfeeMi/findWmsLogisticsTailfeeMi',
    method: 'get',
    params
  })
}

// @Tags WmsLogisticsTailfeeMi
// @Summary 分页获取尾程费用结构列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取尾程费用结构列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /wmsLogisticsTailfeeMi/getWmsLogisticsTailfeeMiList [get]
export const getWmsLogisticsTailfeeMiList = (params) => {
  return service({
    url: '/wmsLogisticsTailfeeMi/getWmsLogisticsTailfeeMiList',
    method: 'get',
    params
  })
}
