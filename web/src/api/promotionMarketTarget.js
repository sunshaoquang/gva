import service from '@/utils/request'

// @Tags PromotionMarketTarget
// @Summary 创建大促推广目标表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PromotionMarketTarget true "创建大促推广目标表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /MarketTarget/createPromotionMarketTarget [post]
export const createPromotionMarketTarget = (data) => {
  return service({
    url: '/MarketTarget/createPromotionMarketTarget',
    method: 'post',
    data
  })
}

// @Tags PromotionMarketTarget
// @Summary 删除大促推广目标表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PromotionMarketTarget true "删除大促推广目标表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /MarketTarget/deletePromotionMarketTarget [delete]
export const deletePromotionMarketTarget = (data) => {
  return service({
    url: '/MarketTarget/deletePromotionMarketTarget',
    method: 'delete',
    data
  })
}

// @Tags PromotionMarketTarget
// @Summary 批量删除大促推广目标表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除大促推广目标表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /MarketTarget/deletePromotionMarketTarget [delete]
export const deletePromotionMarketTargetByIds = (data) => {
  return service({
    url: '/MarketTarget/deletePromotionMarketTargetByIds',
    method: 'delete',
    data
  })
}

// @Tags PromotionMarketTarget
// @Summary 更新大促推广目标表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PromotionMarketTarget true "更新大促推广目标表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /MarketTarget/updatePromotionMarketTarget [put]
export const updatePromotionMarketTarget = (data) => {
  return service({
    url: '/MarketTarget/updatePromotionMarketTarget',
    method: 'put',
    data
  })
}

// @Tags PromotionMarketTarget
// @Summary 用id查询大促推广目标表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.PromotionMarketTarget true "用id查询大促推广目标表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /MarketTarget/findPromotionMarketTarget [get]
export const findPromotionMarketTarget = (params) => {
  return service({
    url: '/MarketTarget/findPromotionMarketTarget',
    method: 'get',
    params
  })
}

// @Tags PromotionMarketTarget
// @Summary 分页获取大促推广目标表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取大促推广目标表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /MarketTarget/getPromotionMarketTargetList [get]
export const getPromotionMarketTargetList = (params) => {
  return service({
    url: '/MarketTarget/getPromotionMarketTargetList',
    method: 'get',
    params
  })
}
