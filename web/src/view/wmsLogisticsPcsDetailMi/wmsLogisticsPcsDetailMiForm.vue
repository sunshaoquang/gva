<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="产品简称:" prop="productAbbreviation">
          <el-input v-model="formData.productAbbreviation" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="国家:" prop="country">
          <el-input v-model="formData.country" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="费用类型CNY:" prop="feeTypeCNY">
          <el-input v-model="formData.feeTypeCNY" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="降本基数(2022年10月):" prop="year2022Month10">
          <el-input v-model="formData.year2022Month10" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="2022年11月:" prop="year2022Month11">
          <el-input v-model="formData.year2022Month11" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="2022年12月:" prop="year2022Month12">
          <el-input v-model="formData.year2022Month12" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="2023年1月:" prop="year2023Month1">
          <el-input v-model="formData.year2023Month1" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="2023年2月:" prop="year2023Month2">
          <el-input v-model="formData.year2023Month2" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="2023年3月:" prop="year2023Month3">
          <el-input v-model="formData.year2023Month3" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="2023年4月:" prop="year2023Month4">
          <el-input v-model="formData.year2023Month4" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="2023年5月:" prop="year2023Month5">
          <el-input v-model="formData.year2023Month5" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="2023年6月:" prop="year2023Month6">
          <el-input v-model="formData.year2023Month6" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="2023年7月:" prop="year2023Month7">
          <el-input v-model="formData.year2023Month7" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="2023年8月:" prop="year2023Month8">
          <el-input v-model="formData.year2023Month8" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="2023年9月:" prop="year2023Month9">
          <el-input v-model="formData.year2023Month9" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="2023年10月:" prop="year2023Month10">
          <el-input v-model="formData.year2023Month10" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="2023年11月:" prop="year2023Month11">
          <el-input v-model="formData.year2023Month11" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="2023年12月:" prop="year2023Month12">
          <el-input v-model="formData.year2023Month12" :clearable="true" placeholder="请输入" />
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
  createWmsLogisticsPcsDetailMi,
  updateWmsLogisticsPcsDetailMi,
  findWmsLogisticsPcsDetailMi
} from '@/api/wmsLogisticsPcsDetailMi'

defineOptions({
    name: 'WmsLogisticsPcsDetailMiForm'
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
            productAbbreviation: '',
            country: '',
            feeTypeCNY: '',
            year2022Month10: '',
            year2022Month11: '',
            year2022Month12: '',
            year2023Month1: '',
            year2023Month2: '',
            year2023Month3: '',
            year2023Month4: '',
            year2023Month5: '',
            year2023Month6: '',
            year2023Month7: '',
            year2023Month8: '',
            year2023Month9: '',
            year2023Month10: '',
            year2023Month11: '',
            year2023Month12: '',
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
      const res = await findWmsLogisticsPcsDetailMi({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.rewmsLogisticsPcsDetailMi
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
               res = await createWmsLogisticsPcsDetailMi(formData.value)
               break
             case 'update':
               res = await updateWmsLogisticsPcsDetailMi(formData.value)
               break
             default:
               res = await createWmsLogisticsPcsDetailMi(formData.value)
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
