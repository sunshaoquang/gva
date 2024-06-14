import service from '@/utils/request'

// @Tags AppCategory
// @Summary 创建AppCategory
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AppCategory true "创建AppCategory"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /appCategory/createAppCategory [post]
export const createAppCategory = (data) => {
  return service({
    url: '/appCategory/createAppCategory',
    method: 'post',
    data
  })
}

// @Tags AppCategory
// @Summary 删除AppCategory
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AppCategory true "删除AppCategory"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /appCategory/deleteAppCategory [delete]
export const deleteAppCategory = (data) => {
  return service({
    url: '/appCategory/deleteAppCategory',
    method: 'delete',
    data
  })
}

// @Tags AppCategory
// @Summary 删除AppCategory
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除AppCategory"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /appCategory/deleteAppCategory [delete]
export const deleteAppCategoryByIds = (data) => {
  return service({
    url: '/appCategory/deleteAppCategoryByIds',
    method: 'delete',
    data
  })
}

// @Tags AppCategory
// @Summary 更新AppCategory
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AppCategory true "更新AppCategory"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /appCategory/updateAppCategory [put]
export const updateAppCategory = (data) => {
  return service({
    url: '/appCategory/updateAppCategory',
    method: 'put',
    data
  })
}

// @Tags AppCategory
// @Summary 用id查询AppCategory
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.AppCategory true "用id查询AppCategory"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /appCategory/findAppCategory [get]
export const findAppCategory = (params) => {
  return service({
    url: '/appCategory/findAppCategory',
    method: 'get',
    params
  })
}

// @Tags AppCategory
// @Summary 分页获取AppCategory列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取AppCategory列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /appCategory/getAppCategoryList [get]
export const getAppCategoryList = (params) => {
  return service({
    url: '/appCategory/getAppCategoryList',
    method: 'get',
    params
  })
}
