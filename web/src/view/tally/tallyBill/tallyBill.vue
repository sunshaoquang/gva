<template>
  <div>
    <div class="gva-search-box">
      <el-form
        :inline="true"
        :model="searchInfo"
        class="demo-form-inline"
        @keyup.enter="onSubmit"
      >
        <el-form-item label="创建时间">
          <el-date-picker
            v-model="searchInfo.startCreatedAt"
            type="datetime"
            placeholder="开始时间"
          ></el-date-picker>
          —
          <el-date-picker
            v-model="searchInfo.endCreatedAt"
            type="datetime"
            placeholder="结束时间"
          ></el-date-picker>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" icon="search" @click="onSubmit">
            查询</el-button
          >
          <el-button icon="refresh" @click="onReset">重置 </el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
      <div class="gva-btn-list">
        <el-button type="primary" icon="plus" @click="openDialog"
          >新增</el-button
        >
        <el-popover v-model:visible="deleteVisible" placement="top" width="160">
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
              >删除</el-button
            >
          </template>
        </el-popover>
      </div>
      <el-table
        ref="multipleTable"
        style="width: 100%"
        tooltip-effect="dark"
        :data="tableData"
        row-key="id"
        @selection-change="handleSelectionChange"
      >
        <el-table-column type="selection" width="55" />
        <el-table-column align="left" label="日期" width="180">
          <template #default="scope">{{
            formatDate(scope.row.createdAt)
          }}</template>
        </el-table-column>
        <el-table-column align="left" label="交付方式" width="120">
          <template #default="scope">
            <el-tag
              :type="
                ['primary', 'danger', 'success'][scope.row.categoryId]
              "
            >
              {{
                showDictLabel(
                  categoryListData,
                  scope.row.categoryId,
                  "id",
                  "name"
                )
              }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column align="left" label="分类类别" width="120">
          <template #default="scope">
            {{
              showDictLabel(
                categoryListData,
                scope.row.deliveryMethodId,
                "id",
                "name"
              )
            }}
          </template>
        </el-table-column>
        <el-table-column
          align="left"
          label="账单金额"
          prop="money"
          width="120"
        />
        <el-table-column align="left" label="标签" prop="tag" width="120">
          <template #default="scope">
            <div class="flex gap-2">
              <div v-for="item in scope.row.tag" :key="item">
                <el-tag
                  :key="item"
                  class="text-gray-100"
                  :color="showDictLabel(tagList, item, 'id', 'tagColor')"
                >
                  {{ showDictLabel(tagList, item, "id", "tagName") }}
                </el-tag>
              </div>
            </div>
          </template>
        </el-table-column>
        <el-table-column align="left" label="备注" prop="remark" width="120" />
        <el-table-column align="left" label="按钮组">
          <template #default="scope">
            <el-button
              type="primary"
              link
              icon="edit"
              class="table-button"
              @click="updateTallyBillFunc(scope.row)"
            >
              变更
            </el-button>
            <el-button
              type="primary"
              link
              icon="delete"
              @click="deleteRow(scope.row)"
            >
              删除
            </el-button>
          </template>
        </el-table-column>
      </el-table>
      <div class="gva-pagination">
        <el-pagination
          layout="total, sizes, prev, pager, next, jumper"
          :current-page="page"
          :page-size="pageSize"
          :page-sizes="[10, 30, 50, 100]"
          :total="total"
          @current-change="handleCurrentChange"
          @size-change="handleSizeChange"
        />
      </div>
    </div>
    <el-drawer
      destroy-on-close
      size="800"
      v-model="dialogFormVisible"
      :show-close="false"
      :before-close="closeDialog"
    >
      <template #header>
        <div class="flex items-center justify-between">
          <span class="text-lg">{{ type === "create" ? "添加" : "修改" }}</span>
          <div>
            <el-button type="primary" @click="enterDialog">确 定</el-button>
            <el-button @click="closeDialog">取 消</el-button>
          </div>
        </div>
      </template>

      <el-form
        :model="formData"
        label-position="right"
        ref="elFormRef"
        :rules="rule"
        label-width="80px"
      >
        <el-form-item label="账单时间">
          <el-date-picker
            v-model="formData.createdAt"
            :clearable="true"
            type="datetime"
            placeholder="请选择账单时间"
          ></el-date-picker>
        </el-form-item>
        <el-form-item label="交付方式:" prop="categoryId">
          <el-select
            v-model="formData.categoryId"
            :clearable="true"
            placeholder="请选择"
            @change="onDeliveryMethodChange"
          >
            <div v-for="item in categoryListData" :key="item?.id">
              <el-option
                v-if="item?.state && item?.parId === 0"
                :label="item?.name"
                :value="item?.id"
              >
              </el-option>
            </div>
          </el-select>
        </el-form-item>
        <el-form-item label="分类类别:" prop="deliveryMethodId">
          <el-select
            v-model="formData.deliveryMethodId"
            :clearable="true"
            placeholder="请选择"
          >
            <div v-for="item in categoryListData1" :key="item?.id">
              <el-option
                v-if="item?.state"
                :label="item?.name"
                :value="item?.id"
              >
              </el-option>
            </div>
          </el-select>
        </el-form-item>
        <!-- <el-form-item label="分类类别:" prop="categoryId">
          <template v-for="item in categoryListData">
            <el-option
              v-if="
                item?.state &&
                item?.parId !== 0 &&
                item?.id === formData.deliveryMethodId
              "
              :key="item?.id"
              :label="item?.name"
              :value="item?.id"
            >
            </el-option>
          </template>
        </el-form-item> -->

        <el-form-item label="账单金额:" prop="money">
          <el-input-number
            v-model="formData.money"
            style="width: 100%"
            :precision="2"
            :clearable="true"
          />
        </el-form-item>
        <el-form-item label="标签:" prop="tag">
          <el-select
            v-model="formData.tag"
            :clearable="true"
            :multiple="true"
            value-key="id"
            placeholder="请选择"
          >
            <div v-for="item in tagList" :key="item.id">
              <el-option :label="item.tagName" :value="item.id"> </el-option>
            </div>
          </el-select>
        </el-form-item>
        <el-form-item label="备注:" prop="remark">
          <el-input
            v-model="formData.remark"
            :clearable="true"
            placeholder="请输入"
          />
        </el-form-item>
      </el-form>
    </el-drawer>
  </div>
</template>

<script>
export default {
  name: "TallyBill",
};
</script>

<script setup>
import {
  createTallyBill,
  deleteTallyBill,
  deleteTallyBillByIds,
  updateTallyBill,
  findTallyBill,
  getTallyBillList,
} from "@/api/tally/tallyBill";
import { getTagsByUser } from "@/api/tally/tallyTags";

import { getTallyCategoryList } from "@/api/tally/tallyCategory";

// 全量引入格式化工具 请按需保留
import {
  getDictFunc,
  formatDate,
  formatBoolean,
  filterDict,
} from "@/utils/format";
import { showDictLabel } from "@/utils/dictionary";
import { ElMessage, ElMessageBox } from "element-plus";
import { ref, reactive, onMounted } from "vue";
// 自动化生成的字典（可能为空）以及字段
const formData = ref({
  categoryId: void 0,
  deliveryMethodId: void 0,
  money: 0,
  tag: [],
  remark: "",
});

// 用户标签列表
const tagList = ref([]);
// 验证规则
const rule = reactive({});

const elFormRef = ref();
onMounted(async () => {
  const res = await getTagsByUser();
  if (res.code === 0) {
    tagList.value = res.data.list;
  }
});

// =========== 表格控制部分 ===========
const page = ref(1);
const total = ref(0);
const pageSize = ref(10);
const tableData = ref([]);
const categoryListData = ref([]);
const categoryListData1 = ref([]);
const searchInfo = ref({});

// 重置
const onReset = () => {
  searchInfo.value = {};
  getTableData();
};

// 搜索
const onSubmit = () => {
  page.value = 1;
  pageSize.value = 10;
  getTableData();
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
  let options = { page: page.value, pageSize: pageSize.value };
  options = { ...options, ...searchInfo.value };
  const table = await getTallyBillList(options);
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
const setOptions = async () => {
  const res = await getTallyCategoryList();
  if (res.code === 0) {
    categoryListData.value = res.data.list;
  }
};

// 获取需要的字典 可能为空 按需保留
setOptions();

// 多选数据
const multipleSelection = ref([]);
// 多选
const handleSelectionChange = (val) => {
  multipleSelection.value = val;
};

// onDeliveryMethodChange
const onDeliveryMethodChange = async (val) => {
  console.log(formData.value,val);
  
  if(formData.value.categoryId !== val){
    formData.value.deliveryMethodId = void 0;
  }
  if (val === void 0) {
    categoryListData1.value = [];
    return;
  }
  const res = await getTallyCategoryList({ parId: val });
  if (res.code === 0) {
    categoryListData1.value = res.data.list;
  }
};

// 删除行
const deleteRow = (row) => {
  ElMessageBox.confirm("确定要删除吗?", "提示", {
    confirmButtonText: "确定",
    cancelButtonText: "取消",
    type: "warning",
  }).then(() => {
    deleteTallyBillFunc(row);
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
      ids.push(item.id);
    });
  const res = await deleteTallyBillByIds({ ids });
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

// 行为控制标记（弹窗内部需要增还是改）
const type = ref("");

// 更新行
const updateTallyBillFunc = async (row) => {
  const res = await findTallyBill({ id: row.id });
  type.value = "update";
  if (res.code === 0) {
    formData.value = res.data;
    onDeliveryMethodChange(res.data.categoryId);
    dialogFormVisible.value = true;
  }
};

// 删除行
const deleteTallyBillFunc = async (row) => {
  const res = await deleteTallyBill({ id: row.id });
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

// 打开弹窗
const openDialog = () => {
  type.value = "create";
  dialogFormVisible.value = true;
};

// 关闭弹窗
const closeDialog = () => {
  dialogFormVisible.value = false;
  formData.value = {
    categoryId: void 0,
    deliveryMethodId: void 0,
    money: 0,
    tag: [],
    remark: "",
  };
};
// 弹窗确定
const enterDialog = async () => {
  elFormRef.value?.validate(async (valid) => {
    if (!valid) return;
    let res;
    switch (type.value) {
      case "create":
        res = await createTallyBill(formData.value);
        break;
      case "update":
        res = await updateTallyBill(formData.value);
        break;
      default:
        res = await createTallyBill(formData.value);
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
</script>

<style></style>
