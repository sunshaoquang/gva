<template>
  <button @click="getPayCodeAction">发起预下单</button>
  <!-- 渲染二维码 -->
  <div>
    <qrcode-vue :value="qrcode.value" :size="qrcode.size" />
  </div>
</template>

<script setup>
import { reactive } from 'vue'
import QrcodeVue from 'qrcode.vue'
import { getPayCode } from '../api/wxpay'

defineOptions({
  name: 'Wxpay'
})

// 一、发起api请求，携带商品ID（千万别带金额等，不要把金钱等敏感信息控制在前端）给后端，调用后端/wxpay/getPayCode接口，创建预下单 获取订单号和二维码

// 二、将二维码渲染到下面提供的两个二维码工具 生成二维码的同时，创建轮询 用第一步获得的ID持续调用/wxpay/getOrderById[get] 参数：{orderID:第一步给的ID}
// （前端可以用setInterval开启定时器
//     const flag(定时器唯一标识) = setInterval(()=>{要做的事},毫秒)
// ）

// 三、当第二步获取订单信息中 status状态成为 SUCCESS时，发放商品等。

const qrcode = reactive({
  value: 'www.baidu.com',
  size: 150
})

const getPayCodeAction = async () => {
  const res = await getPayCode({
    orderID: '123456789'
  })
  console.log(res, '......')
}
</script>
