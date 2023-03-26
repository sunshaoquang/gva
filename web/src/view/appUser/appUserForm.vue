<template>
  <div>
    <div class="gva-form-box">
      <el-form
        ref="elFormRef"
        :model="formData"
        label-position="right"
        :rules="rule"
        label-width="80px"
      >
        <el-form-item label="用户名:" prop="username">
          <el-input
            v-model="formData.username"
            :clearable="true"
            placeholder="请输入"
          />
        </el-form-item>
        <el-form-item label="微信唯一openid:" prop="openid">
          <el-input
            v-model="formData.openid"
            :clearable="true"
            placeholder="请输入"
          />
        </el-form-item>
        <el-form-item label="unionid:" prop="unionid">
          <el-input
            v-model="formData.unionid"
            :clearable="true"
            placeholder="请输入"
          />
        </el-form-item>
        <el-form-item label="用户登陆的次数:" prop="loginCount">
          <el-input
            v-model.number="formData.loginCount"
            :clearable="true"
            placeholder="请输入"
          />
        </el-form-item>
        <el-form-item label="性别:" prop="sex">
          <el-input
            v-model.number="formData.sex"
            :clearable="true"
            placeholder="请输入"
          />
        </el-form-item>
        <el-form-item label="微信头像:" prop="avatar">
          <el-input
            v-model="formData.avatar"
            :clearable="true"
            placeholder="请输入"
          />
        </el-form-item>
        <el-form-item>
          <el-button size="small" type="primary" @click="save">保存</el-button>
          <el-button size="small" type="primary" @click="back">返回</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script>
export default {
  name: "AppUser",
};
</script>

<script setup>
import { createAppUser, updateAppUser, findAppUser } from "@/api/appUser";

// 自动获取字典
import { getDictFunc } from "@/utils/format";
import { useRoute, useRouter } from "vue-router";
import { ElMessage } from "element-plus";
import { ref, reactive } from "vue";
const route = useRoute();
const router = useRouter();

const type = ref("");
const formData = ref({
  username: "",
  openid: "",
  unionid: "",
  loginCount: 0,
  sex: 0,
  avatar: "",
});
// 验证规则
const rule = reactive({
  openid: [
    {
      required: true,
      message: "",
      trigger: ["input", "blur"],
    },
  ],
});

const elFormRef = ref();

// 初始化方法
const init = async () => {
  // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
  if (route.query.id) {
    const res = await findAppUser({ ID: route.query.id });
    if (res.code === 0) {
      formData.value = res.data.reappUser;
      type.value = "update";
    }
  } else {
    type.value = "create";
  }
};

init();
// 保存按钮
const save = async () => {
  elFormRef.value?.validate(async (valid) => {
    if (!valid) return;
    let res;
    switch (type.value) {
      case "create":
        res = await createAppUser(formData.value);
        break;
      case "update":
        res = await updateAppUser(formData.value);
        break;
      default:
        res = await createAppUser(formData.value);
        break;
    }
    if (res.code === 0) {
      ElMessage({
        type: "success",
        message: "创建/更改成功",
      });
    }
  });
};

// 返回按钮
const back = () => {
  router.go(-1);
};
</script>

<style></style>
