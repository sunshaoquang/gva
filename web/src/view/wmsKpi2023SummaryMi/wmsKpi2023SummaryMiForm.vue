<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="指标分类字段:" prop="indicatorClassification">
          <el-input v-model="formData.indicatorClassification" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="指标名称字段:" prop="indicatorName">
          <el-input v-model="formData.indicatorName" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="目标值字段:" prop="target">
          <el-input-number v-model="formData.target" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="1月字段:" prop="january">
          <el-input-number v-model="formData.january" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="2月字段:" prop="february">
          <el-input-number v-model="formData.february" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="3月字段:" prop="march">
          <el-input-number v-model="formData.march" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="4月字段:" prop="april">
          <el-input-number v-model="formData.april" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="5月字段:" prop="may">
          <el-input-number v-model="formData.may" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="6月字段:" prop="june">
          <el-input-number v-model="formData.june" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="7月字段:" prop="july">
          <el-input-number v-model="formData.july" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="8月字段:" prop="august">
          <el-input-number v-model="formData.august" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="9月字段:" prop="september">
          <el-input-number v-model="formData.september" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="10月字段:" prop="october">
          <el-input-number v-model="formData.october" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="11月字段:" prop="november">
          <el-input-number v-model="formData.november" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="12月字段:" prop="december">
          <el-input-number v-model="formData.december" :precision="2" :clearable="true"></el-input-number>
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
  createWmsKpi2023SummaryMi,
  updateWmsKpi2023SummaryMi,
  findWmsKpi2023SummaryMi
} from '@/api/wmsKpi2023SummaryMi'

defineOptions({
    name: 'WmsKpi2023SummaryMiForm'
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
            indicatorClassification: '',
            indicatorName: '',
            target: 0,
            january: 0,
            february: 0,
            march: 0,
            april: 0,
            may: 0,
            june: 0,
            july: 0,
            august: 0,
            september: 0,
            october: 0,
            november: 0,
            december: 0,
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
      const res = await findWmsKpi2023SummaryMi({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.rewmsKpi2023SummaryMi
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
               res = await createWmsKpi2023SummaryMi(formData.value)
               break
             case 'update':
               res = await updateWmsKpi2023SummaryMi(formData.value)
               break
             default:
               res = await createWmsKpi2023SummaryMi(formData.value)
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
