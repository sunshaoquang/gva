import service from '@/utils/request'
// @Tags wxpay
// @Summary 预下单
// @Security ApiKeyAuth
// @Produce  application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"返回成功"}"
// @Router /wxpay/getPayCode [post]   需要根据业务设计传入商品信息
export const getPayCode = (data) => {
  return service({
    url: '/wxpay/getPayCode',
    method: 'post',
    data
  })
}

// @Tags wxpay
// @Summary 获取订单状态
// @Security ApiKeyAuth
// @Produce  application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"返回成功"}"
// @Router /wxpay/getOrderById [get] {orderID:0}
export const getOrderById = (data) => {
  return service({
    url: '/wxpay/getOrderById',
    method: 'get',
    data
  })
}

// @Tags wxpay
// @Summary 支付回调
// @Security ApiKeyAuth
// @Produce  application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"返回成功"}"
// @Router /wxpay/payAction [post] 供微信使用
export const payAction = (data) => {
  return service({
    url: '/wxpay/payAction',
    method: 'post',
    data
  })
}
