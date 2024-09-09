<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="ID:" prop="id">
          <el-input v-model.number="formData.id" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="分类图标:" prop="icon">
          <el-input v-model="formData.icon" :clearable="true"  placeholder="请输入分类图标" />
       </el-form-item>
        <el-form-item label="分类名称:" prop="name">
          <el-input v-model="formData.name" :clearable="true"  placeholder="请输入分类名称" />
       </el-form-item>
        <el-form-item label="交付方式:" prop="deliveryMethod">
        <el-select v-model="formData.deliveryMethod" placeholder="请选择" style="width:100%" :clearable="true">
          <el-option v-for="item in [20]" :key="item" :label="item" :value="item" />
        </el-select>
       </el-form-item>
        <el-form-item label="账单金额:" prop="money">
          <el-input-number v-model="formData.money" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="标签:" prop="tag">
       </el-form-item>
        <el-form-item label="备注:" prop="tag">
          <el-input v-model="formData.tag" :clearable="true"  placeholder="请输入备注" />
       </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="save">保存</el-button>
          <el-button type="primary" @click="back">返回</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script setup>
import {
  createTallyBill,
  updateTallyBill,
  findTallyBill
} from '@/api/tally/tallyBill'

defineOptions({
    name: 'TallyBillForm'
})

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'

const route = useRoute()
const router = useRouter()

const type = ref('')
const formData = ref({
            id: undefined,
            icon: '',
            name: '',
            money: 0,
            tag: [],
            tag: '',
        })
// 验证规则
const rule = reactive({
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findTallyBill({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
}

init()
// 保存按钮
const save = async() => {
      elFormRef.value?.validate( async (valid) => {
         if (!valid) return
            let res
           switch (type.value) {
             case 'create':
               res = await createTallyBill(formData.value)
               break
             case 'update':
               res = await updateTallyBill(formData.value)
               break
             default:
               res = await createTallyBill(formData.value)
               break
           }
           if (res.code === 0) {
             ElMessage({
               type: 'success',
               message: '创建/更改成功'
             })
           }
       })
}

// 返回按钮
const back = () => {
    router.go(-1)
}

</script>

<style>
</style>
