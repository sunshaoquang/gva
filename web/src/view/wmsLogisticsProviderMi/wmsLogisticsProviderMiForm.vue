<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="类型:" prop="type">
          <el-input v-model="formData.type" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="区域:" prop="area">
          <el-input v-model="formData.area" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="目的国:" prop="country">
          <el-input v-model="formData.country" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="供应商:" prop="supplier">
          <el-input v-model="formData.supplier" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="1月:" prop="jan">
          <el-input v-model="formData.jan" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="2月:" prop="feb">
          <el-input v-model="formData.feb" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="3月:" prop="mar">
          <el-input v-model="formData.mar" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="4月:" prop="apr">
          <el-input v-model="formData.apr" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="5月:" prop="may">
          <el-input v-model="formData.may" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="6月:" prop="jun">
          <el-input v-model="formData.jun" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="7月:" prop="jul">
          <el-input v-model="formData.jul" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="8月:" prop="aug">
          <el-input v-model="formData.aug" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="9月:" prop="sept">
          <el-input v-model="formData.sept" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="10月:" prop="oct">
          <el-input v-model="formData.oct" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="11月:" prop="nov">
          <el-input v-model="formData.nov" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="12月:" prop="dec">
          <el-input v-model="formData.dec" :clearable="true" placeholder="请输入" />
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
  createWmsLogisticsProviderMi,
  updateWmsLogisticsProviderMi,
  findWmsLogisticsProviderMi
} from '@/api/wmsLogisticsProviderMi'

defineOptions({
    name: 'WmsLogisticsProviderMiForm'
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
            type: '',
            area: '',
            country: '',
            supplier: '',
            jan: '',
            feb: '',
            mar: '',
            apr: '',
            may: '',
            jun: '',
            jul: '',
            aug: '',
            sept: '',
            oct: '',
            nov: '',
            dec: '',
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
      const res = await findWmsLogisticsProviderMi({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.rewmsLogisticsProviderMi
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
               res = await createWmsLogisticsProviderMi(formData.value)
               break
             case 'update':
               res = await updateWmsLogisticsProviderMi(formData.value)
               break
             default:
               res = await createWmsLogisticsProviderMi(formData.value)
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
