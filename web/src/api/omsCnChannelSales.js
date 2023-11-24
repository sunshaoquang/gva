import service from '@/utils/request'

// @Tags OmsCnChannelSales
// @Summary 创建按天销售目标录入表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.OmsCnChannelSales true "创建按天销售目标录入表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /omsCnChannelSales/createOmsCnChannelSales [post]
export const createOmsCnChannelSales = (data) => {
  return service({
    url: '/omsCnChannelSales/createOmsCnChannelSales',
    method: 'post',
    data
  })
}

// @Tags OmsCnChannelSales
// @Summary 删除按天销售目标录入表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.OmsCnChannelSales true "删除按天销售目标录入表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /omsCnChannelSales/deleteOmsCnChannelSales [delete]
export const deleteOmsCnChannelSales = (data) => {
  return service({
    url: '/omsCnChannelSales/deleteOmsCnChannelSales',
    method: 'delete',
    data
  })
}

// @Tags OmsCnChannelSalesAll
// @Summary 删除所有按天销售目标录入表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.OmsCnChannelSalesAll true "删除所有按天销售目标录入表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /omsCnChannelSales/deleteOmsCnChannelSalesAll [delete]
export const deleteOmsCnChannelSalesAll = () => {
  return service({
    url: '/omsCnChannelSales/deleteOmsCnChannelSalesAll',
    method: 'delete'
  })
}

// @Tags OmsCnChannelSales
// @Summary 批量删除按天销售目标录入表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除按天销售目标录入表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /omsCnChannelSales/deleteOmsCnChannelSales [delete]
export const deleteOmsCnChannelSalesByIds = (data) => {
  return service({
    url: '/omsCnChannelSales/deleteOmsCnChannelSalesByIds',
    method: 'delete',
    data
  })
}

// @Tags OmsCnChannelSales
// @Summary 更新按天销售目标录入表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.OmsCnChannelSales true "更新按天销售目标录入表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /omsCnChannelSales/updateOmsCnChannelSales [put]
export const updateOmsCnChannelSales = (data) => {
  return service({
    url: '/omsCnChannelSales/updateOmsCnChannelSales',
    method: 'put',
    data
  })
}

// @Tags OmsCnChannelSales
// @Summary 用id查询按天销售目标录入表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.OmsCnChannelSales true "用id查询按天销售目标录入表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /omsCnChannelSales/findOmsCnChannelSales [get]
export const findOmsCnChannelSales = (params) => {
  return service({
    url: '/omsCnChannelSales/findOmsCnChannelSales',
    method: 'get',
    params
  })
}

// @Tags OmsCnChannelSales
// @Summary 分页获取按天销售目标录入表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取按天销售目标录入表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /omsCnChannelSales/getOmsCnChannelSalesList [get]
export const getOmsCnChannelSalesList = (params) => {
  return service({
    url: '/omsCnChannelSales/getOmsCnChannelSalesList',
    method: 'get',
    params
  })
}
