import service from '@/utils/request'

// @Tags TallyBill
// @Summary 创建记账账单表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TallyBill true "创建记账账单表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /tallyBill/createTallyBill [post]
export const createTallyBill = (data) => {
  return service({
    url: '/tallyBill/createTallyBill',
    method: 'post',
    data
  })
}

// @Tags TallyBill
// @Summary 删除记账账单表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TallyBill true "删除记账账单表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /tallyBill/deleteTallyBill [delete]
export const deleteTallyBill = (params) => {
  return service({
    url: '/tallyBill/deleteTallyBill',
    method: 'delete',
    params
  })
}

// @Tags TallyBill
// @Summary 批量删除记账账单表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除记账账单表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /tallyBill/deleteTallyBill [delete]
export const deleteTallyBillByIds = (params) => {
  return service({
    url: '/tallyBill/deleteTallyBillByIds',
    method: 'delete',
    params
  })
}

// @Tags TallyBill
// @Summary 更新记账账单表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TallyBill true "更新记账账单表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /tallyBill/updateTallyBill [put]
export const updateTallyBill = (data) => {
  return service({
    url: '/tallyBill/updateTallyBill',
    method: 'put',
    data
  })
}

// @Tags TallyBill
// @Summary 用id查询记账账单表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.TallyBill true "用id查询记账账单表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /tallyBill/findTallyBill [get]
export const findTallyBill = (params) => {
  return service({
    url: '/tallyBill/findTallyBill',
    method: 'get',
    params
  })
}

// @Tags TallyBill
// @Summary 分页获取记账账单表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取记账账单表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /tallyBill/getTallyBillList [get]
export const getTallyBillList = (params) => {
  return service({
    url: '/tallyBill/getTallyBillList',
    method: 'get',
    params
  })
}
