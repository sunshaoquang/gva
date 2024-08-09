import service from "@/utils/request";

// @Tags TallyTags
// @Summary 创建标签
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TallyTags true "创建标签"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /tallyTags/createTallyTags [post]
export const createTallyTags = (data) => {
  return service({
    url: "/tallyTags/createTallyTags",
    method: "post",
    data,
  });
};

// @Tags TallyTags
// @Summary 删除标签
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TallyTags true "删除标签"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /tallyTags/deleteTallyTags [delete]
export const deleteTallyTags = (params) => {
  return service({
    url: "/tallyTags/deleteTallyTags",
    method: "delete",
    params,
  });
};

// @Tags TallyTags
// @Summary 批量删除标签
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除标签"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /tallyTags/deleteTallyTags [delete]
export const deleteTallyTagsByIds = (params) => {
  return service({
    url: "/tallyTags/deleteTallyTagsByIds",
    method: "delete",
    params,
  });
};

// @Tags TallyTags
// @Summary 更新标签
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TallyTags true "更新标签"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /tallyTags/updateTallyTags [put]
export const updateTallyTags = (data) => {
  return service({
    url: "/tallyTags/updateTallyTags",
    method: "put",
    data,
  });
};

// @Tags TallyTags
// @Summary 用id查询标签
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.TallyTags true "用id查询标签"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /tallyTags/findTallyTags [get]
export const findTallyTags = (params) => {
  return service({
    url: "/tallyTags/findTallyTags",
    method: "get",
    params,
  });
};

// @Tags TallyTags
// @Summary 用用户id查询标签列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.TallyTags true "用户id查询标签列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /tallyTags/getTallyTagsListByUser [get]
export const getTagsByUser = () => {
  return service({
    url: "/tallyTags/getTallyTagsListByUser",
    method: "get",
  });
};

// @Tags TallyTags
// @Summary 分页获取标签列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取标签列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /tallyTags/getTallyTagsList [get]
export const getTallyTagsList = (params) => {
  return service({
    url: "/tallyTags/getTallyTagsList",
    method: "get",
    params,
  });
};
