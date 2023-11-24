import service from '@/utils/request'

// @Tags OmsCnSalesTargetDaily
// @Summary 创建京东自营销售录入表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.OmsCnSalesTargetDaily true "创建京东自营销售录入表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /omsCnSalesTargetDaily/createOmsCnSalesTargetDaily [post]
export const createOmsCnSalesTargetDaily = (data) => {
  return service({
    url: '/omsCnSalesTargetDaily/createOmsCnSalesTargetDaily',
    method: 'post',
    data
  })
}

// @Tags OmsCnSalesTargetDaily
// @Summary 删除京东自营销售录入表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.OmsCnSalesTargetDaily true "删除京东自营销售录入表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /omsCnSalesTargetDaily/deleteOmsCnSalesTargetDaily [delete]
export const deleteOmsCnSalesTargetDaily = (data) => {
  return service({
    url: '/omsCnSalesTargetDaily/deleteOmsCnSalesTargetDaily',
    method: 'delete',
    data
  })
}

// @Tags OmsCnSalesTargetDailyAll
// @Summary 删除所有京东自营销售录入表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.OmsCnSalesTargetDailyAll true "删除所有京东自营销售录入表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /omsCnSalesTargetDaily/deleteOmsCnSalesTargetDailyAll [delete]
export const deleteOmsCnSalesTargetDailyAll = () => {
  return service({
    url: '/omsCnSalesTargetDaily/deleteOmsCnSalesTargetDailyAll',
    method: 'delete'
  })
}

// @Tags OmsCnSalesTargetDaily
// @Summary 批量删除京东自营销售录入表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除京东自营销售录入表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /omsCnSalesTargetDaily/deleteOmsCnSalesTargetDaily [delete]
export const deleteOmsCnSalesTargetDailyByIds = (data) => {
  return service({
    url: '/omsCnSalesTargetDaily/deleteOmsCnSalesTargetDailyByIds',
    method: 'delete',
    data
  })
}

// @Tags OmsCnSalesTargetDaily
// @Summary 更新京东自营销售录入表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.OmsCnSalesTargetDaily true "更新京东自营销售录入表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /omsCnSalesTargetDaily/updateOmsCnSalesTargetDaily [put]
export const updateOmsCnSalesTargetDaily = (data) => {
  return service({
    url: '/omsCnSalesTargetDaily/updateOmsCnSalesTargetDaily',
    method: 'put',
    data
  })
}

// @Tags OmsCnSalesTargetDaily
// @Summary 用id查询京东自营销售录入表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.OmsCnSalesTargetDaily true "用id查询京东自营销售录入表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /omsCnSalesTargetDaily/findOmsCnSalesTargetDaily [get]
export const findOmsCnSalesTargetDaily = (params) => {
  return service({
    url: '/omsCnSalesTargetDaily/findOmsCnSalesTargetDaily',
    method: 'get',
    params
  })
}

// @Tags OmsCnSalesTargetDaily
// @Summary 分页获取京东自营销售录入表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取京东自营销售录入表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /omsCnSalesTargetDaily/getOmsCnSalesTargetDailyList [get]
export const getOmsCnSalesTargetDailyList = (params) => {
  return service({
    url: '/omsCnSalesTargetDaily/getOmsCnSalesTargetDailyList',
    method: 'get',
    params
  })
}
