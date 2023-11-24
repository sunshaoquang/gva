import { createRouter, createWebHashHistory } from "vue-router";

const promotionChildren = [
  {
    path: "/template/wmsKpiSummary2Mi",
    name: "WmsKpiSummary2Mi",
    component: () => import("@/view/wmsKpiSummary2Mi/wmsKpiSummary2Mi.vue"),
  },
  {
    path: "/template/promotionMarketTarget",
    name: "PromotionMarketTarget",
    component: () =>
      import("@/view/promotionMarketTarget/promotionMarketTarget.vue"),
  },
  {
    path: "/template/wmsKpiSummary2Mi",
    name: "WmsKpiSummary2Mi",
    component: () => import("@/view/wmsKpiSummary2Mi/wmsKpiSummary2Mi.vue"),
  },
  {
    path: "/template/wmsKpi2023SummaryMi",
    name: "WmsKpi2023SummaryMi",
    component: () =>
      import("@/view/wmsKpi2023SummaryMi/wmsKpi2023SummaryMi.vue"),
  },
  {
    path: "/template/wmsLogisticsPcsDetailMi",
    name: "WmsLogisticsPcsDetailMi",
    component: () =>
      import("@/view/wmsLogisticsPcsDetailMi/wmsLogisticsPcsDetailMi.vue"),
  },
  {
    path: "/template/wmsLogisticsPcsSummaryMi",
    name: "WmsLogisticsPcsSummaryMi",
    component: () =>
      import("@/view/wmsLogisticsPcsSummaryMi/wmsLogisticsPcsSummaryMi.vue"),
  },
  {
    path: "/template/wmsLogisticsPcs2023Mi",
    name: "WmsLogisticsPcs2023Mi",
    component: () =>
      import("@/view/wmsLogisticsPcs2023Mi/wmsLogisticsPcs2023Mi.vue"),
  },
  {
    path: "/template/wmsLogisticsModeTransportMi", //费用by运输方式
    name: "WmsLogisticsModeTransportMi",
    component: () =>
      import(
        "@/view/wmsLogisticsModeTransportMi/wmsLogisticsModeTransportMi.vue"
      ),
  },
  {
    path: "/template/wmsLogisticsProviderMi", //费用by供应商
    name: "WmsLogisticsProviderMi",
    component: () =>
      import("@/view/wmsLogisticsProviderMi/wmsLogisticsProviderMi.vue"),
  },
  {
    path: "/template/wmsLogisticsSalebycountryMi", //收入by国家
    name: "WmsLogisticsSalebycountryMi",
    component: () =>
      import(
        "@/view/wmsLogisticsSalebycountryMi/wmsLogisticsSalebycountryMi.vue"
      ),
  },
  {
    path: "/template/wmsLogisticsStoragefeeMi", //仓储费
    name: "WmsLogisticsStoragefeeMi",
    component: () =>
      import("@/view/wmsLogisticsStoragefeeMi/wmsLogisticsStoragefeeMi.vue"),
  },
  {
    path: "/template/wmsLogisticsRateMonths", //物流占销比
    name: "WmsLogisticsRateMonths",
    component: () =>
      import("@/view/wmsLogisticsRateMonths/wmsLogisticsRateMonths.vue"),
  },
  {
    path: "/template/omsCnSalesTargetDaily", //按天销售目标录入表
    name: "OmsCnSalesTargetDaily",
    component: () =>
      import("@/view/omsCnSalesTargetDaily/omsCnSalesTargetDaily.vue"),
  },
  {
    path: "/template/omsCnChannelSales", //京东自营销售录入表
    name: "OmsCnChannelSales",
    component: () => import("@/view/omsCnChannelSales/omsCnChannelSales.vue"),
  },
];
const routes = [
  {
    path: "/",
    redirect: "/login",
  },
  {
    path: "/init",
    name: "Init",
    component: () => import("@/view/init/index.vue"),
  },
  {
    path: "/login",
    name: "Login",
    component: () => import("@/view/login/index.vue"),
  },
  {
    path: "/template",
    name: "template",
    component: () => import("@/view/routerHolder.vue"),
    children: promotionChildren,
  },
  {
    path: "/:catchAll(.*)",
    meta: {
      closeTab: true,
    },
    component: () => import("@/view/error/index.vue"),
  },
];

const router = createRouter({
  history: createWebHashHistory(),
  routes,
});

export default router;
