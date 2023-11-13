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
        <el-form-item label="手工表名称" prop="sheetName">
          <el-input v-model="searchInfo.sheetName" placeholder="搜索条件" />
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
        @sort-change="sortChange"
      >
        <el-table-column type="selection" width="55" />
        <el-table-column align="left" label="日期" width="180">
          <template #default="scope">{{
            formatDate(scope.row.CreatedAt)
          }}</template>
        </el-table-column>
        <el-table-column
          show-overflow-tooltip
          sortable
          align="left"
          label="指标分类"
          prop="indicatorClassification"
          width="120"
        />
        <el-table-column
          show-overflow-tooltip
          sortable
          align="left"
          label="指标名称"
          prop="indicatorName"
          width="120"
        />
        <el-table-column
          show-overflow-tooltip
          align="left"
          label="目标值"
          prop="target"
          width="120"
        />
        <el-table-column align="left" label="1月" prop="january" width="120" />
        <el-table-column align="left" label="2月" prop="february" width="120" />
        <el-table-column align="left" label="3月" prop="march" width="120" />
        <el-table-column align="left" label="4月" prop="april" width="120" />
        <el-table-column align="left" label="5月" prop="may" width="120" />
        <el-table-column align="left" label="6月" prop="june" width="120" />
        <el-table-column align="left" label="7月" prop="july" width="120" />
        <el-table-column align="left" label="8月" prop="august" width="120" />
        <el-table-column
          align="left"
          label="9月"
          prop="september"
          width="120"
        />
        <el-table-column align="left" label="10月" prop="october" width="120" />
        <el-table-column
          align="left"
          label="11月"
          prop="november"
          width="120"
        />
        <el-table-column
          align="left"
          label="12月"
          prop="december"
          width="120"
        />
        <el-table-column
          show-overflow-tooltip
          sortable
          align="left"
          label="创建人"
          prop="createdName"
          width="120"
        />
        <el-table-column
          show-overflow-tooltip
          sortable
          align="left"
          label="手工表名称"
          prop="sheetName"
          width="120"
        />
        <el-table-column fixed="right" width="300" align="left" label="操作">
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
              @click="updateWmsKpi2023SummaryMiFunc(scope.row)"
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
          <el-form-item label="指标分类:" prop="indicatorClassification">
            <el-input
              v-model="formData.indicatorClassification"
              :clearable="true"
              placeholder="请输入指标分类"
            />
          </el-form-item>
          <el-form-item label="指标名称:" prop="indicatorName">
            <el-input
              v-model="formData.indicatorName"
              :clearable="true"
              placeholder="请输入指标名称"
            />
          </el-form-item>
          <el-form-item label="目标值:" prop="target">
            <el-input
              v-model="formData.target"
              style="width: 100%"
              :clearable="true"
            />
          </el-form-item>
          <el-form-item label="1月:" prop="january">
            <el-input
              v-model="formData.january"
              style="width: 100%"
              :clearable="true"
            />
          </el-form-item>
          <el-form-item label="2月:" prop="february">
            <el-input
              v-model="formData.february"
              style="width: 100%"
              :clearable="true"
            />
          </el-form-item>
          <el-form-item label="3月:" prop="march">
            <el-input
              v-model="formData.march"
              style="width: 100%"
              :clearable="true"
            />
          </el-form-item>
          <el-form-item label="4月:" prop="april">
            <el-input
              v-model="formData.april"
              style="width: 100%"
              :clearable="true"
            />
          </el-form-item>
          <el-form-item label="5月:" prop="may">
            <el-input
              v-model="formData.may"
              style="width: 100%"
              :clearable="true"
            />
          </el-form-item>
          <el-form-item label="6月:" prop="june">
            <el-input
              v-model="formData.june"
              style="width: 100%"
              :clearable="true"
            />
          </el-form-item>
          <el-form-item label="7月:" prop="july">
            <el-input
              v-model="formData.july"
              style="width: 100%"
              :clearable="true"
            />
          </el-form-item>
          <el-form-item label="8月:" prop="august">
            <el-input
              v-model="formData.august"
              style="width: 100%"
              :clearable="true"
            />
          </el-form-item>
          <el-form-item label="9月:" prop="september">
            <el-input
              v-model="formData.september"
              style="width: 100%"
              :clearable="true"
            />
          </el-form-item>
          <el-form-item label="10月:" prop="october">
            <el-input
              v-model="formData.october"
              style="width: 100%"
              :clearable="true"
            />
          </el-form-item>
          <el-form-item label="11月:" prop="november">
            <el-input
              v-model="formData.november"
              style="width: 100%"
              :clearable="true"
            />
          </el-form-item>
          <el-form-item label="12月:" prop="december">
            <el-input
              v-model="formData.december"
              style="width: 100%"
              :clearable="true"
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
          <el-descriptions-item label="指标分类">
            {{ formData.indicatorClassification }}
          </el-descriptions-item>
          <el-descriptions-item label="指标名称">
            {{ formData.indicatorName }}
          </el-descriptions-item>
          <el-descriptions-item label="目标值">
            {{ formData.target }}
          </el-descriptions-item>
          <el-descriptions-item label="1月">
            {{ formData.january }}
          </el-descriptions-item>
          <el-descriptions-item label="2月">
            {{ formData.february }}
          </el-descriptions-item>
          <el-descriptions-item label="3月">
            {{ formData.march }}
          </el-descriptions-item>
          <el-descriptions-item label="4月">
            {{ formData.april }}
          </el-descriptions-item>
          <el-descriptions-item label="5月">
            {{ formData.may }}
          </el-descriptions-item>
          <el-descriptions-item label="6月">
            {{ formData.june }}
          </el-descriptions-item>
          <el-descriptions-item label="7月">
            {{ formData.july }}
          </el-descriptions-item>
          <el-descriptions-item label="8月">
            {{ formData.august }}
          </el-descriptions-item>
          <el-descriptions-item label="9月">
            {{ formData.september }}
          </el-descriptions-item>
          <el-descriptions-item label="10月">
            {{ formData.october }}
          </el-descriptions-item>
          <el-descriptions-item label="11月">
            {{ formData.november }}
          </el-descriptions-item>
          <el-descriptions-item label="12月">
            {{ formData.december }}
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
  createWmsKpi2023SummaryMi,
  deleteWmsKpi2023SummaryMi,
  deleteWmsKpi2023SummaryMiAll,
  deleteWmsKpi2023SummaryMiByIds,
  findWmsKpi2023SummaryMi,
  getWmsKpi2023SummaryMiList,
} from "@/api/wmsKpi2023SummaryMi";

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
  name: "WmsKpi2023SummaryMi",
});

// 自动化生成的字典（可能为空）以及
const formData = ref({
  indicatorClassification: "",
  indicatorName: "",
  target: "",
  january: "",
  february: "",
  march: "",
  april: "",
  may: "",
  june: "",
  july: "",
  august: "",
  september: "",
  october: "",
  november: "",
  december: "",
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
// 排序
const sortChange = ({ prop, order }) => {
  searchInfo.value.sort = prop;
  searchInfo.value.order = order;
  getTableData();
};

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
  const table = await getWmsKpi2023SummaryMiList({
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
    deleteWmsKpi2023SummaryMiFunc(row);
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
  const res = await deleteWmsKpi2023SummaryMiByIds({ ids });
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
  const res = await deleteWmsKpi2023SummaryMiAll();
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
const updateWmsKpi2023SummaryMiFunc = async (row) => {
  const res = await findWmsKpi2023SummaryMi({ ID: row.ID });
  type.value = "update";
  if (res.code === 0) {
    formData.value = res.data.rewmsKpi2023SummaryMi;
    dialogFormVisible.value = true;
  }
};

// 删除行
const deleteWmsKpi2023SummaryMiFunc = async (row) => {
  const res = await deleteWmsKpi2023SummaryMi({ ID: row.ID });
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
  const res = await findWmsKpi2023SummaryMi({ ID: row.ID });
  if (res.code === 0) {
    formData.value = res.data.rewmsKpi2023SummaryMi;
    openDetailShow();
  }
};

// 关闭详情弹窗
const closeDetailShow = () => {
  detailShow.value = false;
  formData.value = {
    indicatorClassification: "",
    indicatorName: "",
    target: "",
    january: "",
    february: "",
    march: "",
    april: "",
    may: "",
    june: "",
    july: "",
    august: "",
    september: "",
    october: "",
    november: "",
    december: "",
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
    indicatorClassification: "",
    indicatorName: "",
    target: "",
    january: "",
    february: "",
    march: "",
    april: "",
    may: "",
    june: "",
    july: "",
    august: "",
    september: "",
    october: "",
    november: "",
    december: "",
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
        res = await createWmsKpi2023SummaryMi(formData.value);
        break;
      case "update":
        res = await updateWmsKpi2023SummaryMi(formData.value);
        break;
      default:
        res = await createWmsKpi2023SummaryMi(formData.value);
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
