<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="项目字段:" prop="project">
          <el-input v-model="formData.project" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="2021年字段:" prop="year2021">
          <el-input-number v-model="formData.year2021" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="2022年字段:" prop="year2022">
          <el-input-number v-model="formData.year2022" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="2023年字段:" prop="year2023">
          <el-input-number v-model="formData.year2023" :precision="2" :clearable="true"></el-input-number>
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
  createWmsKpiSummary2Mi,
  updateWmsKpiSummary2Mi,
  findWmsKpiSummary2Mi
} from '@/api/wmsKpiSummary2Mi'

defineOptions({
    name: 'WmsKpiSummary2MiForm'
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
            project: '',
            year2021: 0,
            year2022: 0,
            year2023: 0,
            createdName: '',
            sheetName: '',
        })
// 验证规则
const rule = reactive({
               project : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findWmsKpiSummary2Mi({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.rewmsKpiSummary2Mi
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
               res = await createWmsKpiSummary2Mi(formData.value)
               break
             case 'update':
               res = await updateWmsKpiSummary2Mi(formData.value)
               break
             default:
               res = await createWmsKpiSummary2Mi(formData.value)
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
