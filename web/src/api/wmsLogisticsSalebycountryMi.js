import service from "@/utils/request";

// @Tags WmsLogisticsSalebycountryMi
// @Summary 创建收入by国家
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.WmsLogisticsSalebycountryMi true "创建收入by国家"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /wmsLogisticsSalebycountryMi/createWmsLogisticsSalebycountryMi [post]
export const createWmsLogisticsSalebycountryMi = (data) => {
  return service({
    url: "/wmsLogisticsSalebycountryMi/createWmsLogisticsSalebycountryMi",
    method: "post",
    data,
  });
};

// @Tags WmsLogisticsSalebycountryMi
// @Summary 删除收入by国家
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.WmsLogisticsSalebycountryMi true "删除收入by国家"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /wmsLogisticsSalebycountryMi/deleteWmsLogisticsSalebycountryMi [delete]
export const deleteWmsLogisticsSalebycountryMi = (data) => {
  return service({
    url: "/wmsLogisticsSalebycountryMi/deleteWmsLogisticsSalebycountryMi",
    method: "delete",
    data,
  });
};

// @Tags WmsLogisticsSalebycountryMiAll
// @Summary 删除所有收入by国家
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.WmsLogisticsSalebycountryMiAll true "删除所有收入by国家"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /wmsLogisticsSalebycountryMi/deleteWmsLogisticsSalebycountryMiAll [delete]
export const deleteWmsLogisticsSalebycountryMiAll = () => {
  return service({
    url: "/wmsLogisticsSalebycountryMi/deleteWmsLogisticsSalebycountryMiAll",
    method: "delete",
  });
};

// @Tags WmsLogisticsSalebycountryMi
// @Summary 批量删除收入by国家
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除收入by国家"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /wmsLogisticsSalebycountryMi/deleteWmsLogisticsSalebycountryMi [delete]
export const deleteWmsLogisticsSalebycountryMiByIds = (data) => {
  return service({
    url: "/wmsLogisticsSalebycountryMi/deleteWmsLogisticsSalebycountryMiByIds",
    method: "delete",
    data,
  });
};

// @Tags WmsLogisticsSalebycountryMi
// @Summary 更新收入by国家
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.WmsLogisticsSalebycountryMi true "更新收入by国家"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /wmsLogisticsSalebycountryMi/updateWmsLogisticsSalebycountryMi [put]
export const updateWmsLogisticsSalebycountryMi = (data) => {
  return service({
    url: "/wmsLogisticsSalebycountryMi/updateWmsLogisticsSalebycountryMi",
    method: "put",
    data,
  });
};

// @Tags WmsLogisticsSalebycountryMi
// @Summary 用id查询收入by国家
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.WmsLogisticsSalebycountryMi true "用id查询收入by国家"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /wmsLogisticsSalebycountryMi/findWmsLogisticsSalebycountryMi [get]
export const findWmsLogisticsSalebycountryMi = (params) => {
  return service({
    url: "/wmsLogisticsSalebycountryMi/findWmsLogisticsSalebycountryMi",
    method: "get",
    params,
  });
};

// @Tags WmsLogisticsSalebycountryMi
// @Summary 分页获取收入by国家列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取收入by国家列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /wmsLogisticsSalebycountryMi/getWmsLogisticsSalebycountryMiList [get]
export const getWmsLogisticsSalebycountryMiList = (params) => {
  return service({
    url: "/wmsLogisticsSalebycountryMi/getWmsLogisticsSalebycountryMiList",
    method: "get",
    params,
  });
};
