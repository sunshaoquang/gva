import service from '@/utils/request'

// @Tags WmsLogisticsStoragefeeMi
// @Summary 创建仓储费
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.WmsLogisticsStoragefeeMi true "创建仓储费"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /wmsLogisticsStoragefeeMi/createWmsLogisticsStoragefeeMi [post]
export const createWmsLogisticsStoragefeeMi = (data) => {
  return service({
    url: '/wmsLogisticsStoragefeeMi/createWmsLogisticsStoragefeeMi',
    method: 'post',
    data
  })
}

// @Tags WmsLogisticsStoragefeeMi
// @Summary 删除仓储费
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.WmsLogisticsStoragefeeMi true "删除仓储费"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /wmsLogisticsStoragefeeMi/deleteWmsLogisticsStoragefeeMi [delete]
export const deleteWmsLogisticsStoragefeeMi = (data) => {
  return service({
    url: '/wmsLogisticsStoragefeeMi/deleteWmsLogisticsStoragefeeMi',
    method: 'delete',
    data
  })
}

// @Tags WmsLogisticsStoragefeeMiAll
// @Summary 删除所有仓储费
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.WmsLogisticsStoragefeeMiAll true "删除所有仓储费"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /wmsLogisticsStoragefeeMi/deleteWmsLogisticsStoragefeeMiAll [delete]
export const deleteWmsLogisticsStoragefeeMiAll = () => {
  return service({
    url: '/wmsLogisticsStoragefeeMi/deleteWmsLogisticsStoragefeeMiAll',
    method: 'delete'
  })
}

// @Tags WmsLogisticsStoragefeeMi
// @Summary 批量删除仓储费
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除仓储费"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /wmsLogisticsStoragefeeMi/deleteWmsLogisticsStoragefeeMi [delete]
export const deleteWmsLogisticsStoragefeeMiByIds = (data) => {
  return service({
    url: '/wmsLogisticsStoragefeeMi/deleteWmsLogisticsStoragefeeMiByIds',
    method: 'delete',
    data
  })
}

// @Tags WmsLogisticsStoragefeeMi
// @Summary 更新仓储费
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.WmsLogisticsStoragefeeMi true "更新仓储费"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /wmsLogisticsStoragefeeMi/updateWmsLogisticsStoragefeeMi [put]
export const updateWmsLogisticsStoragefeeMi = (data) => {
  return service({
    url: '/wmsLogisticsStoragefeeMi/updateWmsLogisticsStoragefeeMi',
    method: 'put',
    data
  })
}

// @Tags WmsLogisticsStoragefeeMi
// @Summary 用id查询仓储费
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.WmsLogisticsStoragefeeMi true "用id查询仓储费"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /wmsLogisticsStoragefeeMi/findWmsLogisticsStoragefeeMi [get]
export const findWmsLogisticsStoragefeeMi = (params) => {
  return service({
    url: '/wmsLogisticsStoragefeeMi/findWmsLogisticsStoragefeeMi',
    method: 'get',
    params
  })
}

// @Tags WmsLogisticsStoragefeeMi
// @Summary 分页获取仓储费列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取仓储费列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /wmsLogisticsStoragefeeMi/getWmsLogisticsStoragefeeMiList [get]
export const getWmsLogisticsStoragefeeMiList = (params) => {
  return service({
    url: '/wmsLogisticsStoragefeeMi/getWmsLogisticsStoragefeeMiList',
    method: 'get',
    params
  })
}
