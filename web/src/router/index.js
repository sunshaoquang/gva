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
