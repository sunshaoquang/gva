<template>
  <div>
    <div class="gva-search-box">
      <el-form
        ref="elSearchFormRef"
        :inline="true"
        :model="searchInfo"
        class="demo-form-inline"
        :rules="searchRule"
        @keyup.enter="onSubmit"
      >
        <el-form-item label="创建日期" prop="createdAt">
          <template #label>
            <span>
              创建日期
              <el-tooltip
                content="搜索范围是开始日期（包含）至结束日期（不包含）"
              >
                <el-icon><QuestionFilled /></el-icon>
              </el-tooltip>
            </span>
          </template>
          <el-date-picker
            v-model="searchInfo.startCreatedAt"
            type="datetime"
            placeholder="开始日期"
            :disabled-date="
              (time) =>
                searchInfo.endCreatedAt
                  ? time.getTime() > searchInfo.endCreatedAt.getTime()
                  : false
            "
          ></el-date-picker>
          —
          <el-date-picker
            v-model="searchInfo.endCreatedAt"
            type="datetime"
            placeholder="结束日期"
            :disabled-date="
              (time) =>
                searchInfo.startCreatedAt
                  ? time.getTime() < searchInfo.startCreatedAt.getTime()
                  : false
            "
          ></el-date-picker>
        </el-form-item>
        <el-form-item label="创建人" prop="createdName">
          <el-input v-model="searchInfo.createdName" placeholder="搜索条件" />
        </el-form-item>
        <el-form-item>
          <el-button type="primary" icon="search" @click="onSubmit"
            >查询</el-button
          >
          <el-button icon="refresh" @click="onReset">重置</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
      <div class="gva-btn-list">
        <el-button type="primary" icon="plus" @click="openDialog"
          >新增</el-button
        >
        <el-upload
          class="excel-btn"
          :action="`${path}/excel/importExcel?fileName=${titleName}`"
          :headers="{ 'x-token': userStore.token }"
          :on-success="outputExcel"
          :on-change="
            (e) => {
              sheetName = e.name;
            }
          "
          :show-file-list="false"
        >
          <el-button type="primary" icon="upload">导入</el-button>
        </el-upload>
        <el-button
          class="excel-btn"
          type="success"
          icon="download"
          @click="downloadExcelTemplate()"
          >下载模板</el-button
        >
        <el-popover
          v-model:visible="deleteVisible"
          :disabled="!multipleSelection.length"
          placement="top"
          width="160"
        >
          <p>确定要删除吗？</p>
          <div style="text-align: right; margin-top: 8px">
            <el-button type="primary" link @click="deleteVisible = false"
              >取消</el-button
            >
            <el-button type="primary" @click="onDelete">确定</el-button>
          </div>
          <template #reference>
            <el-button
              icon="delete"
              style="margin-left: 10px"
              :disabled="!multipleSelection.length"
              @click="deleteVisible = true"
              >选择删除</el-button
            >
          </template>
        </el-popover>
        <!--  ================== 增加 ===================== -->
        <el-popover
          v-model:visible="deleteAllVisible"
          placement="top"
          width="160"
        >
          <p>确定要删除吗？</p>
          <div style="text-align: right; margin-top: 8px">
            <el-button type="primary" link @click="deleteAllVisible = false"
              >取消</el-button
            >
            <el-button type="primary" @click="onAllDelete">确定</el-button>
          </div>
          <template #reference>
            <el-button
              icon="delete"
              style="margin-left: 10px"
              @click="deleteAllVisible = true"
              >全部删除</el-button
            >
          </template>
        </el-popover>
        <!--  ================== 增加 ===================== -->
      </div>
      <el-table
        ref="multipleTable"
        style="width: 100%"
        tooltip-effect="dark"
        :data="tableData"
        row-key="ID"
        @selection-change="handleSelectionChange"
      >
        <el-table-column type="selection" width="55" />
        <el-table-column
          align="left"
          show-overflow-tooltip
          label="日期"
          width="180"
        >
          <template #default="scope">{{
            formatDate(scope.row.CreatedAt)
          }}</template>
        </el-table-column>
        <el-table-column
          align="left"
          show-overflow-tooltip
          label="产品简称"
          prop="productAbbreviation"
          width="120"
        />
        <el-table-column
          align="left"
          show-overflow-tooltip
          label="国家"
          prop="country"
          width="120"
        />
        <el-table-column
          align="left"
          show-overflow-tooltip
          label="费用类型CNY"
          prop="feeTypeCNY"
          width="120"
        />
        <el-table-column
          align="left"
          show-overflow-tooltip
          label="降本基数(2022年10月)"
          prop="year2022Month10"
          width="120"
        />
        <el-table-column
          align="left"
          show-overflow-tooltip
          label="2022年11月"
          prop="year2022Month11"
          width="120"
        />
        <el-table-column
          align="left"
          show-overflow-tooltip
          label="2022年12月"
          prop="year2022Month12"
          width="120"
        />
        <el-table-column
          align="left"
          show-overflow-tooltip
          label="2023年1月"
          prop="year2023Month1"
          width="120"
        />
        <el-table-column
          align="left"
          show-overflow-tooltip
          label="2023年2月"
          prop="year2023Month2"
          width="120"
        />
        <el-table-column
          align="left"
          show-overflow-tooltip
          label="2023年3月"
          prop="year2023Month3"
          width="120"
        />
        <el-table-column
          align="left"
          show-overflow-tooltip
          label="2023年4月"
          prop="year2023Month4"
          width="120"
        />
        <el-table-column
          align="left"
          show-overflow-tooltip
          label="2023年5月"
          prop="year2023Month5"
          width="120"
        />
        <el-table-column
          align="left"
          show-overflow-tooltip
          label="2023年6月"
          prop="year2023Month6"
          width="120"
        />
        <el-table-column
          align="left"
          show-overflow-tooltip
          label="2023年7月"
          prop="year2023Month7"
          width="120"
        />
        <el-table-column
          align="left"
          show-overflow-tooltip
          label="2023年8月"
          prop="year2023Month8"
          width="120"
        />
        <el-table-column
          align="left"
          show-overflow-tooltip
          label="2023年9月"
          prop="year2023Month9"
          width="120"
        />
        <el-table-column
          align="left"
          show-overflow-tooltip
          label="2023年10月"
          prop="year2023Month10"
          width="120"
        />
        <el-table-column
          align="left"
          show-overflow-tooltip
          label="2023年11月"
          prop="year2023Month11"
          width="120"
        />
        <el-table-column
          align="left"
          show-overflow-tooltip
          label="2023年12月"
          prop="year2023Month12"
          width="120"
        />
        <el-table-column
          align="left"
          show-overflow-tooltip
          label="创建人"
          prop="createdName"
          width="120"
        />
        <el-table-column
          align="left"
          show-overflow-tooltip
          label="手工表名称"
          prop="sheetName"
          width="120"
        />
        <el-table-column
          fixed="right"
          width="300"
          align="left"
          show-overflow-tooltip
          label="操作"
        >
          <template #default="scope">
            <el-button
              type="primary"
              link
              class="table-button"
              @click="getDetails(scope.row)"
            >
              <el-icon style="margin-right: 5px"><InfoFilled /></el-icon>
              查看详情
            </el-button>
            <el-button
              type="primary"
              link
              icon="edit"
              class="table-button"
              @click="updateWmsLogisticsPcsDetailMiFunc(scope.row)"
              >变更</el-button
            >
            <el-button
              type="primary"
              link
              icon="delete"
              @click="deleteRow(scope.row)"
              >删除</el-button
            >
          </template>
        </el-table-column>
      </el-table>
      <div class="gva-pagination">
        <el-pagination
          layout="total, sizes, prev, pager, next, jumper"
          :current-page="page"
          :page-size="pageSize"
          :page-sizes="[10, 30, 50, 100, 1000, 10000]"
          :total="total"
          @current-change="handleCurrentChange"
          @size-change="handleSizeChange"
        />
      </div>
    </div>
    <el-dialog
      v-model="dialogFormVisible"
      :before-close="closeDialog"
      :title="type === 'create' ? '添加' : '修改'"
      destroy-on-close
    >
      <el-scrollbar height="500px">
        <el-form
          :model="formData"
          label-position="right"
          ref="elFormRef"
          :rules="rule"
          label-width="80px"
        >
          <el-form-item label="产品简称:" prop="productAbbreviation">
            <el-input
              v-model="formData.productAbbreviation"
              :clearable="true"
              placeholder="请输入产品简称"
            />
          </el-form-item>
          <el-form-item label="国家:" prop="country">
            <el-input
              v-model="formData.country"
              :clearable="true"
              placeholder="请输入国家"
            />
          </el-form-item>
          <el-form-item label="费用类型CNY:" prop="feeTypeCNY">
            <el-input
              v-model="formData.feeTypeCNY"
              :clearable="true"
              placeholder="请输入费用类型CNY"
            />
          </el-form-item>
          <el-form-item label="降本基数(2022年10月):" prop="year2022Month10">
            <el-input
              v-model="formData.year2022Month10"
              :clearable="true"
              placeholder="请输入降本基数(2022年10月)"
            />
          </el-form-item>
          <el-form-item label="2022年11月:" prop="year2022Month11">
            <el-input
              v-model="formData.year2022Month11"
              :clearable="true"
              placeholder="请输入2022年11月"
            />
          </el-form-item>
          <el-form-item label="2022年12月:" prop="year2022Month12">
            <el-input
              v-model="formData.year2022Month12"
              :clearable="true"
              placeholder="请输入2022年12月"
            />
          </el-form-item>
          <el-form-item label="2023年1月:" prop="year2023Month1">
            <el-input
              v-model="formData.year2023Month1"
              :clearable="true"
              placeholder="请输入2023年1月"
            />
          </el-form-item>
          <el-form-item label="2023年2月:" prop="year2023Month2">
            <el-input
              v-model="formData.year2023Month2"
              :clearable="true"
              placeholder="请输入2023年2月"
            />
          </el-form-item>
          <el-form-item label="2023年3月:" prop="year2023Month3">
            <el-input
              v-model="formData.year2023Month3"
              :clearable="true"
              placeholder="请输入2023年3月"
            />
          </el-form-item>
          <el-form-item label="2023年4月:" prop="year2023Month4">
            <el-input
              v-model="formData.year2023Month4"
              :clearable="true"
              placeholder="请输入2023年4月"
            />
          </el-form-item>
          <el-form-item label="2023年5月:" prop="year2023Month5">
            <el-input
              v-model="formData.year2023Month5"
              :clearable="true"
              placeholder="请输入2023年5月"
            />
          </el-form-item>
          <el-form-item label="2023年6月:" prop="year2023Month6">
            <el-input
              v-model="formData.year2023Month6"
              :clearable="true"
              placeholder="请输入2023年6月"
            />
          </el-form-item>
          <el-form-item label="2023年7月:" prop="year2023Month7">
            <el-input
              v-model="formData.year2023Month7"
              :clearable="true"
              placeholder="请输入2023年7月"
            />
          </el-form-item>
          <el-form-item label="2023年8月:" prop="year2023Month8">
            <el-input
              v-model="formData.year2023Month8"
              :clearable="true"
              placeholder="请输入2023年8月"
            />
          </el-form-item>
          <el-form-item label="2023年9月:" prop="year2023Month9">
            <el-input
              v-model="formData.year2023Month9"
              :clearable="true"
              placeholder="请输入2023年9月"
            />
          </el-form-item>
          <el-form-item label="2023年10月:" prop="year2023Month10">
            <el-input
              v-model="formData.year2023Month10"
              :clearable="true"
              placeholder="请输入2023年10月"
            />
          </el-form-item>
          <el-form-item label="2023年11月:" prop="year2023Month11">
            <el-input
              v-model="formData.year2023Month11"
              :clearable="true"
              placeholder="请输入2023年11月"
            />
          </el-form-item>
          <el-form-item label="2023年12月:" prop="year2023Month12">
            <el-input
              v-model="formData.year2023Month12"
              :clearable="true"
              placeholder="请输入2023年12月"
            />
          </el-form-item>
          <el-form-item label="创建人:" prop="createdName">
            <el-input
              v-model="formData.createdName"
              :clearable="true"
              placeholder="请输入创建人"
            />
          </el-form-item>
          <el-form-item label="手工表名称:" prop="sheetName">
            <el-input
              v-model="formData.sheetName"
              :clearable="true"
              placeholder="请输入手工表名称"
            />
          </el-form-item>
        </el-form>
      </el-scrollbar>
      <template #footer>
        <div class="dialog-footer">
          <el-button @click="closeDialog">取 消</el-button>
          <el-button type="primary" @click="enterDialog">确 定</el-button>
        </div>
      </template>
    </el-dialog>

    <el-dialog
      v-model="detailShow"
      style="width: 800px"
      lock-scroll
      :before-close="closeDetailShow"
      title="查看详情"
      destroy-on-close
    >
      <el-scrollbar height="550px">
        <el-descriptions column="1" border>
          <el-descriptions-item label="产品简称">
            {{ formData.productAbbreviation }}
          </el-descriptions-item>
          <el-descriptions-item label="国家">
            {{ formData.country }}
          </el-descriptions-item>
          <el-descriptions-item label="费用类型CNY">
            {{ formData.feeTypeCNY }}
          </el-descriptions-item>
          <el-descriptions-item label="降本基数(2022年10月)">
            {{ formData.year2022Month10 }}
          </el-descriptions-item>
          <el-descriptions-item label="2022年11月">
            {{ formData.year2022Month11 }}
          </el-descriptions-item>
          <el-descriptions-item label="2022年12月">
            {{ formData.year2022Month12 }}
          </el-descriptions-item>
          <el-descriptions-item label="2023年1月">
            {{ formData.year2023Month1 }}
          </el-descriptions-item>
          <el-descriptions-item label="2023年2月">
            {{ formData.year2023Month2 }}
          </el-descriptions-item>
          <el-descriptions-item label="2023年3月">
            {{ formData.year2023Month3 }}
          </el-descriptions-item>
          <el-descriptions-item label="2023年4月">
            {{ formData.year2023Month4 }}
          </el-descriptions-item>
          <el-descriptions-item label="2023年5月">
            {{ formData.year2023Month5 }}
          </el-descriptions-item>
          <el-descriptions-item label="2023年6月">
            {{ formData.year2023Month6 }}
          </el-descriptions-item>
          <el-descriptions-item label="2023年7月">
            {{ formData.year2023Month7 }}
          </el-descriptions-item>
          <el-descriptions-item label="2023年8月">
            {{ formData.year2023Month8 }}
          </el-descriptions-item>
          <el-descriptions-item label="2023年9月">
            {{ formData.year2023Month9 }}
          </el-descriptions-item>
          <el-descriptions-item label="2023年10月">
            {{ formData.year2023Month10 }}
          </el-descriptions-item>
          <el-descriptions-item label="2023年11月">
            {{ formData.year2023Month11 }}
          </el-descriptions-item>
          <el-descriptions-item label="2023年12月">
            {{ formData.year2023Month12 }}
          </el-descriptions-item>
          <el-descriptions-item label="创建人">
            {{ formData.createdName }}
          </el-descriptions-item>
          <el-descriptions-item label="手工表名称">
            {{ formData.sheetName }}
          </el-descriptions-item>
        </el-descriptions>
      </el-scrollbar>
    </el-dialog>
  </div>
</template>

<script setup>
import {
  createWmsLogisticsPcsDetailMi,
  deleteWmsLogisticsPcsDetailMi,
  deleteWmsLogisticsPcsDetailMiAll,
  deleteWmsLogisticsPcsDetailMiByIds,
  updateWmsLogisticsPcsDetailMi,
  findWmsLogisticsPcsDetailMi,
  getWmsLogisticsPcsDetailMiList,
} from "@/api/wmsLogisticsPcsDetailMi";

// 全量引入格式化工具 请按需保留
import {
  getDictFunc,
  formatDate,
  formatBoolean,
  filterDict,
  ReturnArrImg,
  onDownloadFile,
} from "@/utils/format";
import { ElMessage, ElMessageBox, ElUpload } from "element-plus"; // 更新 新增 ElUpload
import { ref, reactive, computed } from "vue"; // 更新 新增 computed

// ================== 增加 =====================
import { outputExcelData, downloadTemplate } from "@/api/customer";
import { useUserStore } from "@/pinia/modules/user";
import config from "@/core/config";
const sheetName = ref("");
const path = ref(import.meta.env.VITE_BASE_API);
// ================== 增加 =====================

defineOptions({
  name: "WmsLogisticsPcsDetailMi",
});

// 自动化生成的字典（可能为空）以及字段
const formData = ref({
  productAbbreviation: "",
  country: "",
  feeTypeCNY: "",
  year2022Month10: "",
  year2022Month11: "",
  year2022Month12: "",
  year2023Month1: "",
  year2023Month2: "",
  year2023Month3: "",
  year2023Month4: "",
  year2023Month5: "",
  year2023Month6: "",
  year2023Month7: "",
  year2023Month8: "",
  year2023Month9: "",
  year2023Month10: "",
  year2023Month11: "",
  year2023Month12: "",
  createdName: "",
  sheetName: "",
});

// 验证规则
const rule = reactive({});

const searchRule = reactive({
  createdAt: [
    {
      validator: (rule, value, callback) => {
        if (searchInfo.value.startCreatedAt && !searchInfo.value.endCreatedAt) {
          callback(new Error("请填写结束日期"));
        } else if (
          !searchInfo.value.startCreatedAt &&
          searchInfo.value.endCreatedAt
        ) {
          callback(new Error("请填写开始日期"));
        } else if (
          searchInfo.value.startCreatedAt &&
          searchInfo.value.endCreatedAt &&
          (searchInfo.value.startCreatedAt.getTime() ===
            searchInfo.value.endCreatedAt.getTime() ||
            searchInfo.value.startCreatedAt.getTime() >
              searchInfo.value.endCreatedAt.getTime())
        ) {
          callback(new Error("开始日期应当早于结束日期"));
        } else {
          callback();
        }
      },
      trigger: "change",
    },
  ],
});

const elFormRef = ref();
const elSearchFormRef = ref();

// =========== 表格控制部分 ===========
const page = ref(1);
const total = ref(0);
const pageSize = ref(10);
const tableData = ref([]);
const searchInfo = ref({});

// 重置
const onReset = () => {
  searchInfo.value = {};
  getTableData();
};

// 搜索
const onSubmit = () => {
  elSearchFormRef.value?.validate(async (valid) => {
    if (!valid) return;
    page.value = 1;
    pageSize.value = 10;
    getTableData();
  });
};

// 分页
const handleSizeChange = (val) => {
  pageSize.value = val;
  getTableData();
};

// 修改页面容量
const handleCurrentChange = (val) => {
  page.value = val;
  getTableData();
};

// 查询
const getTableData = async () => {
  const table = await getWmsLogisticsPcsDetailMiList({
    page: page.value,
    pageSize: pageSize.value,
    ...searchInfo.value,
  });
  if (table.code === 0) {
    tableData.value = table.data.list;
    total.value = table.data.total;
    page.value = table.data.page;
    pageSize.value = table.data.pageSize;
  }
};

getTableData();

// ============== 表格控制部分结束 ===============

// 获取需要的字典 可能为空 按需保留
const setOptions = async () => {};

// 获取需要的字典 可能为空 按需保留
setOptions();

// 多选数据
const multipleSelection = ref([]);
// 多选
const handleSelectionChange = (val) => {
  multipleSelection.value = val;
};

// 删除行
const deleteRow = (row) => {
  ElMessageBox.confirm("确定要删除吗?", "提示", {
    confirmButtonText: "确定",
    cancelButtonText: "取消",
    type: "warning",
  }).then(() => {
    deleteWmsLogisticsPcsDetailMiFunc(row);
  });
};

// 批量删除控制标记
const deleteVisible = ref(false);

// 多选删除
const onDelete = async () => {
  const ids = [];
  if (multipleSelection.value.length === 0) {
    ElMessage({
      type: "warning",
      message: "请选择要删除的数据",
    });
    return;
  }
  multipleSelection.value &&
    multipleSelection.value.map((item) => {
      ids.push(item.ID);
    });
  const res = await deleteWmsLogisticsPcsDetailMiByIds({ ids });
  if (res.code === 0) {
    ElMessage({
      type: "success",
      message: "删除成功",
    });
    if (tableData.value.length === ids.length && page.value > 1) {
      page.value--;
    }
    deleteVisible.value = false;
    getTableData();
  }
};

// ================== 增加 =====================
// 全部删除控制标记
const deleteAllVisible = ref(false);

const onAllDelete = async () => {
  const res = await deleteWmsLogisticsPcsDetailMiAll();
  if (res.code === 0) {
    ElMessage({
      type: "success",
      message: "删除成功",
    });
    deleteAllVisible.value = false;
    getTableData();
  }
};
// ================== 增加 =====================

// 行为控制标记（弹窗内部需要增还是改）
const type = ref("");

// 更新行
const updateWmsLogisticsPcsDetailMiFunc = async (row) => {
  const res = await findWmsLogisticsPcsDetailMi({ ID: row.ID });
  type.value = "update";
  if (res.code === 0) {
    formData.value = res.data.rewmsLogisticsPcsDetailMi;
    dialogFormVisible.value = true;
  }
};

// 删除行
const deleteWmsLogisticsPcsDetailMiFunc = async (row) => {
  const res = await deleteWmsLogisticsPcsDetailMi({ ID: row.ID });
  if (res.code === 0) {
    ElMessage({
      type: "success",
      message: "删除成功",
    });
    if (tableData.value.length === 1 && page.value > 1) {
      page.value--;
    }
    getTableData();
  }
};

// 弹窗控制标记
const dialogFormVisible = ref(false);

// 查看详情控制标记
const detailShow = ref(false);

// 打开详情弹窗
const openDetailShow = () => {
  detailShow.value = true;
};

// 打开详情
const getDetails = async (row) => {
  // 打开弹窗
  const res = await findWmsLogisticsPcsDetailMi({ ID: row.ID });
  if (res.code === 0) {
    formData.value = res.data.rewmsLogisticsPcsDetailMi;
    openDetailShow();
  }
};

// 关闭详情弹窗
const closeDetailShow = () => {
  detailShow.value = false;
  formData.value = {
    productAbbreviation: "",
    country: "",
    feeTypeCNY: "",
    year2022Month10: "",
    year2022Month11: "",
    year2022Month12: "",
    year2023Month1: "",
    year2023Month2: "",
    year2023Month3: "",
    year2023Month4: "",
    year2023Month5: "",
    year2023Month6: "",
    year2023Month7: "",
    year2023Month8: "",
    year2023Month9: "",
    year2023Month10: "",
    year2023Month11: "",
    year2023Month12: "",
    createdName: "",
    sheetName: "",
  };
};

// 打开弹窗
const openDialog = () => {
  type.value = "create";
  dialogFormVisible.value = true;
};

// 关闭弹窗
const closeDialog = () => {
  dialogFormVisible.value = false;
  formData.value = {
    productAbbreviation: "",
    country: "",
    feeTypeCNY: "",
    year2022Month10: "",
    year2022Month11: "",
    year2022Month12: "",
    year2023Month1: "",
    year2023Month2: "",
    year2023Month3: "",
    year2023Month4: "",
    year2023Month5: "",
    year2023Month6: "",
    year2023Month7: "",
    year2023Month8: "",
    year2023Month9: "",
    year2023Month10: "",
    year2023Month11: "",
    year2023Month12: "",
    createdName: "",
    sheetName: "",
  };
};
// 弹窗确定
const enterDialog = async () => {
  elFormRef.value?.validate(async (valid) => {
    if (!valid) return;
    let res;
    switch (type.value) {
      case "create":
        res = await createWmsLogisticsPcsDetailMi(formData.value);
        break;
      case "update":
        res = await updateWmsLogisticsPcsDetailMi(formData.value);
        break;
      default:
        res = await createWmsLogisticsPcsDetailMi(formData.value);
        break;
    }
    if (res.code === 0) {
      ElMessage({
        type: "success",
        message: "创建/更改成功",
      });
      closeDialog();
      getTableData();
    }
  });
};

// ================== 增加 =====================
const userStore = useUserStore();
const titleName = computed(() =>
  document?.title?.replace(` - ${config.appName}`, "")
);
const outputExcel = async () => {
  await outputExcelData(titleName.value, sheetName.value, window.name);
  getTableData();
};
const downloadExcelTemplate = () => {
  downloadTemplate(`${titleName.value}__ExcelTemplate.xlsx`);
};
// ================== 增加 =====================
</script>

<style scoped>
::v-deep(.el-upload__input) {
  display: none;
}
</style>
