import service from '@/utils/request'

// @Tags AppBill
// @Summary 创建AppBill
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AppBill true "创建AppBill"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /appBill/createAppBill [post]
export const createAppBill = (data) => {
  return service({
    url: '/appBill/createAppBill',
    method: 'post',
    data
  })
}

// @Tags AppBill
// @Summary 删除AppBill
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AppBill true "删除AppBill"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /appBill/deleteAppBill [delete]
export const deleteAppBill = (data) => {
  return service({
    url: '/appBill/deleteAppBill',
    method: 'delete',
    data
  })
}

// @Tags AppBill
// @Summary 删除AppBill
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除AppBill"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /appBill/deleteAppBill [delete]
export const deleteAppBillByIds = (data) => {
  return service({
    url: '/appBill/deleteAppBillByIds',
    method: 'delete',
    data
  })
}

// @Tags AppBill
// @Summary 更新AppBill
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AppBill true "更新AppBill"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /appBill/updateAppBill [put]
export const updateAppBill = (data) => {
  return service({
    url: '/appBill/updateAppBill',
    method: 'put',
    data
  })
}

// @Tags AppBill
// @Summary 用id查询AppBill
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.AppBill true "用id查询AppBill"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /appBill/findAppBill [get]
export const findAppBill = (params) => {
  return service({
    url: '/appBill/findAppBill',
    method: 'get',
    params
  })
}

// @Tags AppBill
// @Summary 分页获取AppBill列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取AppBill列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /appBill/getAppBillList [get]
export const getAppBillList = (params) => {
  return service({
    url: '/appBill/getAppBillList',
    method: 'get',
    params
  })
}
