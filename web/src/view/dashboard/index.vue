<template>
<<<<<<< HEAD
  <div class="page">
    <div class="gva-card-box">
      <div class="gva-card gva-top-card">
        <div class="gva-top-card-left">
          <div class="gva-top-card-left-title">
            早安，管理员，请开始一天的工作吧
          </div>
          <div class="gva-top-card-left-dot">{{ weatherInfo }}</div>
          <el-row class="my-8 w-[500px]">
            <el-col :span="8" :xs="24" :sm="8">
              <div class="flex items-center">
                <el-icon class="dashboard-icon">
                  <sort />
                </el-icon>
                今日流量 (1231231)
              </div>
            </el-col>
            <el-col :span="8" :xs="24" :sm="8">
              <div class="flex items-center">
                <el-icon class="dashboard-icon">
                  <avatar />
                </el-icon>
                总用户数 (24001)
              </div>
            </el-col>
            <el-col :span="8" :xs="24" :sm="8">
              <div class="flex items-center">
                <el-icon class="dashboard-icon">
                  <comment />
                </el-icon>
                好评率 (99%)
              </div>
            </el-col>
          </el-row>
          <div>
            <div class="gva-top-card-left-item">
              使用教学：
              <a
                style="color: #409eff"
                target="view_window"
                href="https://www.bilibili.com/video/BV1Rg411u7xH/"
                >https://www.bilibili.com/video/BV1Rg411u7xH</a
              >
            </div>
            <div class="gva-top-card-left-item">
              插件仓库：
              <a
                style="color: #409eff"
                target="view_window"
                href="https://plugin.gin-vue-admin.com/#/layout/home"
                >https://plugin.gin-vue-admin.com</a
              >
            </div>
          </div>
        </div>
        <img src="@/assets/dashboard.png" class="gva-top-card-right" alt />
      </div>
    </div>
    <div class="gva-card-box">
      <div class="gva-card quick-entrance">
        <div class="gva-card-title">快捷入口</div>
        <el-row :gutter="20">
          <el-col
            v-for="(card, key) in toolCards"
            :key="key"
            :span="4"
            :xs="8"
            class="quick-entrance-items"
            @click="toTarget(card.name)"
          >
            <div class="quick-entrance-item">
              <div
                class="quick-entrance-item-icon"
                :style="{ backgroundColor: card.bg }"
              >
                <el-icon>
                  <component :is="card.icon" :style="{ color: card.color }" />
                </el-icon>
              </div>
              <p>{{ card.label }}</p>
            </div>
          </el-col>
        </el-row>
      </div>
    </div>
    <div class="gva-card-box">
      <div class="gva-card">
        <div class="gva-card-title">数据统计</div>
        <div class="p-4">
          <el-row :gutter="20">
            <el-col :xs="24" :sm="18">
              <echarts-line />
            </el-col>
            <el-col :xs="24" :sm="6">
              <dashboard-table />
            </el-col>
          </el-row>
        </div>
      </div>
    </div>
=======
  <div
    class="grid grid-cols-1 md:grid-cols-3 lg:grid-cols-7 py-2 gap-4 md:gap-2 gva-container2"
  >
    <gva-card custom-class="col-span-1 lg:col-span-2 ">
      <gva-chart :type="1" title="访问人数" />
    </gva-card>
    <gva-card custom-class="col-span-1 lg:col-span-2  ">
      <gva-chart :type="2" title="新增客户" />
    </gva-card>
    <gva-card custom-class="col-span-1 lg:col-span-2 ">
      <gva-chart :type="3" title="解决数量" />
    </gva-card>
    <gva-card
      title="快捷功能"
      show-action
      custom-class="col-start-1 md:col-start-3 lg:col-start-7 row-span-2 "
    >
      <gva-quick-link />
    </gva-card>
    <gva-card
      title="内容数据"
      custom-class="col-span-1 md:col-span-2 md:row-start-2 lg:col-span-6 col-start-1 row-span-2"
    >
      <gva-chart :type="4" />
    </gva-card>
    <gva-card
      title="文档"
      show-action
      custom-class="md:row-start-8 md:col-start-3 lg:row-start-3 lg:col-start-7"
    >
      <gva-wiki />
    </gva-card>

    <gva-card
      title="最新更新"
      custom-class="col-span-1 md:col-span-3 row-span-2"
    >
      <gva-table />
    </gva-card>
    <gva-card
      title="最新插件"
      custom-class="col-span-1 md:col-span-3 row-span-2"
    >
      <gva-plugin-table />
    </gva-card>

    <gva-card title="公告" show-action custom-class="col-span-1 lg:col-start-7">
      <gva-notice />
    </gva-card>

    <gva-card
      without-padding
      custom-class="overflow-hidden lg:h-40 col-span-1 md:col-start-2 md:col-span-1 lg:col-start-7"
    >
      <gva-banner />
    </gva-card>
>>>>>>> main
  </div>
</template>

<script setup>
<<<<<<< HEAD
import EchartsLine from "@/view/dashboard/dashboardCharts/echartsLine.vue";
import DashboardTable from "@/view/dashboard/dashboardTable/dashboardTable.vue";
import { ref } from "vue";
import { useRouter } from "vue-router";
import { useWeatherInfo } from "@/view/dashboard/weather.js";

defineOptions({
  name: "Dashboard",
});

const weatherInfo = useWeatherInfo();

const toolCards = ref([
  {
    label: "用户管理",
    icon: "monitor",
    name: "user",
    color: "#ff9c6e",
    bg: "rgba(255, 156, 110,.3)",
  },
  {
    label: "角色管理",
    icon: "setting",
    name: "authority",
    color: "#69c0ff",
    bg: "rgba(105, 192, 255,.3)",
  },
  {
    label: "菜单管理",
    icon: "menu",
    name: "menu",
    color: "#b37feb",
    bg: "rgba(179, 127, 235,.3)",
  },
  {
    label: "代码生成器",
    icon: "cpu",
    name: "autoCode",
    color: "#ffd666",
    bg: "rgba(255, 214, 102,.3)",
  },
  {
    label: "表单生成器",
    icon: "document-checked",
    name: "formCreate",
    color: "#ff85c0",
    bg: "rgba(255, 133, 192,.3)",
  },
  {
    label: "关于我们",
    icon: "user",
    name: "about",
    color: "#5cdbd3",
    bg: "rgba(92, 219, 211,.3)",
  },
]);

const router = useRouter();

const toTarget = (name) => {
  router.push({ name });
};
</script>

<style lang="scss" scoped>
.page {
  @apply p-0;
  .gva-card-box {
    @apply p-4;
    & + .gva-card-box {
      @apply pt-0;
    }
  }
  .gva-card {
    @apply box-border bg-white rounded h-auto px-6 py-8 overflow-hidden shadow-sm;
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
=======
  import {
    GvaPluginTable,
    GvaTable,
    GvaChart,
    GvaWiki,
    GvaNotice,
    GvaQuickLink,
    GvaCard,
    GvaBanner
  } from './components'
  defineOptions({
    name: 'Dashboard'
  })
</script>

<style lang="scss" scoped></style>
>>>>>>> main
