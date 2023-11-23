<template>
  <div class="page">
    <div class="gva-card-box">
      <div class="gva-card">
        <div>
          <h4>一、操作指引说明：</h4>
          <ul class="ml-2">
            <li>Step1:下载最新的模板;</li>
            <li>
              Step2:填充最新的数据（请自行确认数据准确性；<strong
                style="color: red"
                >注：Sheet1名称不要修改</strong
              >）;
            </li>
            <li>Step3:视具体的情况选择【选择删除】&【全部删除】；</li>
            <li>Step4:点击【导入】，选择需要上传的文件；</li>
            <li>Step5:导入成功后,查询确认。</li>
          </ul>
          <h4>二、录入功能联系窗口：</h4>
          <div class="mb-3 ml-2">
            IT信息中心：张坤动；邮箱：zhangkd@hello-tech.com"
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
defineOptions({
  name: "TagInfo",
});
</script>

<style lang="scss" scoped>
.page {
  @apply p-0;
  .gva-card-box {
    @apply pt-4;
    & + .gva-card-box {
      @apply pt-0;
    }
  }
  li {
    color: #999;
  }
  .gva-card {
    @apply box-border bg-white rounded h-auto px-6 py-2 overflow-hidden shadow-sm;
    .gva-card-title {
      @apply pb-5 border-t-0 border-l-0 border-r-0 border-b border-solid border-gray-100;
    }
  }
  .gva-top-card {
    @apply h-72 flex items-center justify-between text-gray-500;
    &-left {
      @apply h-full flex flex-col w-auto;
      &-title {
        @apply text-3xl text-gray-600;
      }
      &-dot {
        @apply mt-4 text-gray-600 text-lg;
      }
      &-item {
        + .gva-top-card-left-item {
          margin-top: 24px;
        }
        margin-top: 14px;
      }
    }
    &-right {
      height: 600px;
      width: 600px;
      margin-top: 28px;
    }
  }
  ::v-deep(.el-card__header) {
    @apply p-0  border-gray-200;
  }
  .card-header {
    @apply pb-5 border-b border-solid border-gray-200 border-t-0 border-l-0 border-r-0;
  }
  .quick-entrance-items {
    @apply flex items-center justify-center text-center text-gray-800;
    .quick-entrance-item {
      @apply px-8 py-6 flex items-center flex-col transition-all duration-100 ease-in-out rounded-lg cursor-pointer;
      &:hover {
        @apply shadow-lg;
      }
      &-icon {
        @apply flex items-center h-16 w-16 rounded-lg justify-center mx-0 my-auto text-2xl;
      }
      p {
        @apply mt-2.5;
      }
    }
  }
}
.dashboard-icon {
  @apply flex items-center text-xl mr-2 text-blue-400;
}
</style>
