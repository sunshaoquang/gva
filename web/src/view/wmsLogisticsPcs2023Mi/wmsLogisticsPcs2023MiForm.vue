<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="国家:" prop="country">
          <el-input v-model="formData.country" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="产品简称:" prop="productAbbreviation">
          <el-input v-model="formData.productAbbreviation" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="类型:" prop="type">
          <el-input v-model="formData.type" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="1月:" prop="january">
          <el-input v-model="formData.january" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="2月:" prop="february">
          <el-input v-model="formData.february" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="3月:" prop="march">
          <el-input v-model="formData.march" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="4月:" prop="april">
          <el-input v-model="formData.april" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="5月:" prop="may">
          <el-input v-model="formData.may" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="6月:" prop="june">
          <el-input v-model="formData.june" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="7月:" prop="july">
          <el-input v-model="formData.july" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="8月:" prop="august">
          <el-input v-model="formData.august" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="9月:" prop="september">
          <el-input v-model="formData.september" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="10月:" prop="october">
          <el-input v-model="formData.october" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="11月:" prop="november">
          <el-input v-model="formData.november" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="12月:" prop="december">
          <el-input v-model="formData.december" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="创建人:" prop="createdName">
          <el-input v-model="formData.createdName" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="手工表名称:" prop="sheetName">
          <el-input v-model="formData.sheetName" :clearable="true" placeholder="请输入" />
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
  createWmsLogisticsPcs2023Mi,
  updateWmsLogisticsPcs2023Mi,
  findWmsLogisticsPcs2023Mi
} from '@/api/wmsLogisticsPcs2023Mi'

defineOptions({
    name: 'WmsLogisticsPcs2023MiForm'
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
            country: '',
            productAbbreviation: '',
            type: '',
            january: '',
            february: '',
            march: '',
            april: '',
            may: '',
            june: '',
            july: '',
            august: '',
            september: '',
            october: '',
            november: '',
            december: '',
            createdName: '',
            sheetName: '',
        })
// 验证规则
const rule = reactive({
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findWmsLogisticsPcs2023Mi({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.rewmsLogisticsPcs2023Mi
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
               res = await createWmsLogisticsPcs2023Mi(formData.value)
               break
             case 'update':
               res = await updateWmsLogisticsPcs2023Mi(formData.value)
               break
             default:
               res = await createWmsLogisticsPcs2023Mi(formData.value)
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
