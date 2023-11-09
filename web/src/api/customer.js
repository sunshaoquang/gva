import service from "@/utils/request";
import { ElMessage } from "element-plus";

// @Tags SysApi
// @Summary 删除客户
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body dbModel.ExaCustomer true "删除客户"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /customer/customer [post]
export const createExaCustomer = (data) => {
  return service({
    url: "/customer/customer",
    method: "post",
    data,
  });
};

// @Tags SysApi
// @Summary 更新客户信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body dbModel.ExaCustomer true "更新客户信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /customer/customer [put]
export const updateExaCustomer = (data) => {
  return service({
    url: "/customer/customer",
    method: "put",
    data,
  });
};

// @Tags SysApi
// @Summary 创建客户
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body dbModel.ExaCustomer true "创建客户"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /customer/customer [delete]
export const deleteExaCustomer = (data) => {
  return service({
    url: "/customer/customer",
    method: "delete",
    data,
  });
};

// @Tags SysApi
// @Summary 获取单一客户信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body dbModel.ExaCustomer true "获取单一客户信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /customer/customer [get]
export const getExaCustomer = (params) => {
  return service({
    url: "/customer/customer",
    method: "get",
    params,
  });
};

// @Tags SysApi
// @Summary 获取权限客户列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body modelInterface.PageInfo true "获取权限客户列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /customer/customerList [get]
export const getExaCustomerList = (params) => {
  return service({
    url: "/customer/customerList",
    method: "get",
    params,
  });
};

const handleFileError = (res, fileName) => {
  if (typeof res.data !== "undefined") {
    if (res.data.type === "application/json") {
      const reader = new FileReader();
      reader.onload = function () {
        const message = JSON.parse(reader.result).msg;
        ElMessage({
          showClose: true,
          message: message,
          type: "error",
        });
      };
      reader.readAsText(new Blob([res.data]));
    }
  } else {
    var downloadUrl = window.URL.createObjectURL(new Blob([res]));
    var a = document.createElement("a");
    a.style.display = "none";
    a.href = downloadUrl;
    a.download = fileName;
    var event = new MouseEvent("click");
    a.dispatchEvent(event);
  }
};

// @Tags excel
// @Summary 导出Excel
// @Security ApiKeyAuth
// @accept application/json
// @Produce  application/octet-stream
// @Param data body model.ExcelInfo true "导出Excel文件信息"
// @Success 200
// @Router /excel/exportExcel [post]
export const exportExcel = (tableData, fileName) => {
  service({
    url: "/excel/exportExcel",
    method: "post",
    data: {
      fileName: fileName,
      infoList: tableData,
    },
    responseType: "blob",
  }).then((res) => {
    handleFileError(res, fileName);
  });
};

// @Tags excel
// @Summary 导入Excel文件
// @Security ApiKeyAuth
// @accept multipart/form-data
// @Produce  application/json
// @Param file formData file true "导入Excel文件"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"导入成功"}"
// @Router /excel/importExcel [post]
export const loadExcelData = () => {
  return service({
    url: "/excel/loadExcel",
    method: "get",
  });
};

// @Tags excel
// @Summary 下载模板
// @Security ApiKeyAuth
// @accept multipart/form-data
// @Produce  application/json
// @Param fileName query fileName true "模板名称"
// @Success 200
// @Router /excel/downloadTemplate [get]
export const downloadTemplate = (fileName) => {
  return service({
    url: "/excel/downloadTemplate",
    method: "get",
    params: {
      fileName: fileName,
    },
    responseType: "blob",
  }).then((res) => {
    handleFileError(res, fileName);
  });
};

// @Tags excel
// @Summary 上传Excel数据
// @Security ApiKeyAuth
// @accept multipart/form-data
// @Produce  application/json
// @Param data body example.ExcelInfo true "上传Excel数据"
// @Success 200 {object} response.Response{msg=string} "上传Excel数据"
// @Router /excel/outputExcel [post]
export const outputExcelData = (fileName, sheetName, createdName) => {
  return service({
    url: "/excel/outputExcel",
    method: "post",
    params: {
      fileName: fileName,
      sheetName: sheetName,
      createdName: createdName,
    },
  }).then((res) => {
    handleFileError(res, fileName);
  });
};
