<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="区域:" prop="area">
          <el-input v-model="formData.area" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="运输方式:" prop="transMethod">
          <el-input v-model="formData.transMethod" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="月份:" prop="mmonth">
          <el-input v-model="formData.mmonth" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="费用:" prop="cost">
          <el-input v-model="formData.cost" :clearable="true" placeholder="请输入" />
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
  createWmsLogisticsModeTransportMi,
  updateWmsLogisticsModeTransportMi,
  findWmsLogisticsModeTransportMi
} from '@/api/wmsLogisticsModeTransportMi'

defineOptions({
    name: 'WmsLogisticsModeTransportMiForm'
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
            area: '',
            transMethod: '',
            mmonth: '',
            cost: '',
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
      const res = await findWmsLogisticsModeTransportMi({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.rewmsLogisticsModeTransportMi
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
               res = await createWmsLogisticsModeTransportMi(formData.value)
               break
             case 'update':
               res = await updateWmsLogisticsModeTransportMi(formData.value)
               break
             default:
               res = await createWmsLogisticsModeTransportMi(formData.value)
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
