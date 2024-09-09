import service from "@/utils/request";

// @Tags TallyCategory
// @Summary 创建记账类别表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TallyCategory true "创建记账类别表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /tallyCategory/createTallyCategory [post]
export const createTallyCategory = (data) => {
  return service({
    url: "/tallyCategory/createTallyCategory",
    method: "post",
    data,
  });
};

// @Tags TallyCategory
// @Summary 删除记账类别表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TallyCategory true "删除记账类别表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /tallyCategory/deleteTallyCategory [delete]
export const deleteTallyCategory = (params) => {
  return service({
    url: "/tallyCategory/deleteTallyCategory",
    method: "delete",
    params,
  });
};

// @Tags TallyCategory
// @Summary 批量删除记账类别表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除记账类别表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /tallyCategory/deleteTallyCategory [delete]
export const deleteTallyCategoryByIds = (params) => {
  return service({
    url: "/tallyCategory/deleteTallyCategoryByIds",
    method: "delete",
    params,
  });
};

// @Tags TallyCategory
// @Summary 更新记账类别表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TallyCategory true "更新记账类别表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /tallyCategory/updateTallyCategory [put]
export const updateTallyCategory = (data) => {
  return service({
    url: "/tallyCategory/updateTallyCategory",
    method: "put",
    data,
  });
};

// @Tags TallyCategory
// @Summary 用id查询记账类别表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.TallyCategory true "用id查询记账类别表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /tallyCategory/findTallyCategory [get]
export const findTallyCategory = (params) => {
  return service({
    url: "/tallyCategory/findTallyCategory",
    method: "get",
    params,
  });
};

// @Tags TallyCategory
// @Summary 分页获取记账类别表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取记账类别表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /tallyCategory/getTallyCategoryList [get]
export const getTallyCategoryList = (params) => {
  return service({
    url: "/tallyCategory/getTallyCategoryList",
    method: "get",
    params,
  });
};

// @Tags TallyCategory
// @Summary 分页获取记账类别表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取记账类别表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /tallyCategory/getTallyCategoryPublic [get]
export const getTallyCategoryPublic = (params) => {
  return service({
    url: "/tallyCategory/getTallyCategoryPublic",
    method: "get",
    params,
  });
};
