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
        <el-form-item label="降本基数10月成本:" prop="reduceBaseCostInOctober">
          <el-input v-model="formData.reduceBaseCostInOctober" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="承诺达成降本金额:" prop="commitmentToAchieveCostReductionAmount">
          <el-input v-model="formData.commitmentToAchieveCostReductionAmount" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="1月VS10月降本:" prop="costReductionInMonth1VsMonth10">
          <el-input v-model="formData.costReductionInMonth1VsMonth10" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="2月VS10月降本:" prop="costReductionInMonth2VsMonth10">
          <el-input v-model="formData.costReductionInMonth2VsMonth10" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="3月VS10月降本:" prop="costReductionInMonth3VsMonth10">
          <el-input v-model="formData.costReductionInMonth3VsMonth10" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="4月VS10月降本:" prop="costReductionInMonth4VsMonth10">
          <el-input v-model="formData.costReductionInMonth4VsMonth10" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="5月VS10月降本:" prop="costReductionInMonth5VsMonth10">
          <el-input v-model="formData.costReductionInMonth5VsMonth10" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="6月VS10月降本:" prop="costReductionInMonth6VsMonth10">
          <el-input v-model="formData.costReductionInMonth6VsMonth10" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="7月VS10月降本:" prop="costReductionInMonth7VsMonth10">
          <el-input v-model="formData.costReductionInMonth7VsMonth10" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="8月VS10月降本:" prop="costReductionInMonth8VsMonth10">
          <el-input v-model="formData.costReductionInMonth8VsMonth10" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="9月VS10月降本:" prop="costReductionInMonth9VsMonth10">
          <el-input v-model="formData.costReductionInMonth9VsMonth10" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="10月VS10月降本:" prop="costReductionInMonth10VsMonth10">
          <el-input v-model="formData.costReductionInMonth10VsMonth10" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="11月VS10月降本:" prop="costReductionInMonth11VsMonth10">
          <el-input v-model="formData.costReductionInMonth11VsMonth10" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="12月VS10月降本:" prop="costReductionInMonth12VsMonth10">
          <el-input v-model="formData.costReductionInMonth12VsMonth10" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="1-当前月累计加权降本金额:" prop="monthlyCumulativeWeightedPrincipalReductionAmount">
          <el-input v-model="formData.monthlyCumulativeWeightedPrincipalReductionAmount" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="1-当前月累计加权降本完成度:" prop="monthlyCumulativeWeightedCostReductionCompletionDegree">
          <el-input v-model="formData.monthlyCumulativeWeightedCostReductionCompletionDegree" :clearable="true" placeholder="请输入" />
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
  createWmsLogisticsPcsSummaryMi,
  updateWmsLogisticsPcsSummaryMi,
  findWmsLogisticsPcsSummaryMi
} from '@/api/wmsLogisticsPcsSummaryMi'

defineOptions({
    name: 'WmsLogisticsPcsSummaryMiForm'
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
            reduceBaseCostInOctober: '',
            commitmentToAchieveCostReductionAmount: '',
            costReductionInMonth1VsMonth10: '',
            costReductionInMonth2VsMonth10: '',
            costReductionInMonth3VsMonth10: '',
            costReductionInMonth4VsMonth10: '',
            costReductionInMonth5VsMonth10: '',
            costReductionInMonth6VsMonth10: '',
            costReductionInMonth7VsMonth10: '',
            costReductionInMonth8VsMonth10: '',
            costReductionInMonth9VsMonth10: '',
            costReductionInMonth10VsMonth10: '',
            costReductionInMonth11VsMonth10: '',
            costReductionInMonth12VsMonth10: '',
            monthlyCumulativeWeightedPrincipalReductionAmount: '',
            monthlyCumulativeWeightedCostReductionCompletionDegree: '',
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
      const res = await findWmsLogisticsPcsSummaryMi({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.rewmsLogisticsPcsSummaryMi
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
               res = await createWmsLogisticsPcsSummaryMi(formData.value)
               break
             case 'update':
               res = await updateWmsLogisticsPcsSummaryMi(formData.value)
               break
             default:
               res = await createWmsLogisticsPcsSummaryMi(formData.value)
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
