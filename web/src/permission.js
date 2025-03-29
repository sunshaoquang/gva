<<<<<<< HEAD
import { useUserStore } from "@/pinia/modules/user";
import { useRouterStore } from "@/pinia/modules/router";
import getPageTitle from "@/utils/page";
import router from "@/router";
import Nprogress from "nprogress";

const whiteList = ["Login", "Init"];

const getRouter = async (userStore) => {
  const routerStore = useRouterStore();
  await routerStore.SetAsyncRouter();
  await userStore.GetUserInfo();
  const asyncRouters = routerStore.asyncRouters;
  asyncRouters.forEach((asyncRouter) => {
    router.addRoute(asyncRouter);
  });
};

async function handleKeepAlive(to) {
  if (to.matched.some((item) => item.meta.keepAlive)) {
    if (to.matched && to.matched.length > 2) {
      for (let i = 1; i < to.matched.length; i++) {
        const element = to.matched[i - 1];
        if (element.name === "layout") {
          to.matched.splice(i, 1);
          await handleKeepAlive(to);
        }
        // 如果没有按需加载完成则等待加载
        if (typeof element.components.default === "function") {
          await element.components.default();
          await handleKeepAlive(to);
        }
=======
import { useUserStore } from '@/pinia/modules/user'
import { useRouterStore } from '@/pinia/modules/router'
import getPageTitle from '@/utils/page'
import router from '@/router'
import Nprogress from 'nprogress'
import 'nprogress/nprogress.css'

// 配置 NProgress
Nprogress.configure({
  showSpinner: false,
  ease: 'ease',
  speed: 500
})

// 白名单路由
const WHITE_LIST = ['Login', 'Init']

// 处理路由加载
const setupRouter = async (userStore) => {
  try {
    const routerStore = useRouterStore()
    await Promise.all([routerStore.SetAsyncRouter(), userStore.GetUserInfo()])

    routerStore.asyncRouters.forEach((route) => router.addRoute(route))
    return true
  } catch (error) {
    console.error('Setup router failed:', error)
    return false
  }
}

// 移除加载动画
const removeLoading = () => {
  const element = document.getElementById('gva-loading-box')
  element?.remove()
}

// 处理组件缓存
const handleKeepAlive = async (to) => {
  if (!to.matched.some((item) => item.meta.keepAlive)) return

  if (to.matched?.length > 2) {
    for (let i = 1; i < to.matched.length; i++) {
      const element = to.matched[i - 1]

      if (element.name === 'layout') {
        to.matched.splice(i, 1)
        await handleKeepAlive(to)
        continue
      }

      if (typeof element.components.default === 'function') {
        await element.components.default()
        await handleKeepAlive(to)
>>>>>>> main
      }
    }
  }
}

<<<<<<< HEAD
router.beforeEach(async (to, from) => {
  const routerStore = useRouterStore();
  Nprogress.start();
  const userStore = useUserStore();
  to.meta.matched = [...to.matched];
  handleKeepAlive(to);
  const token = userStore.token;
  //TODO 在白名单中的判断情况
  document.title = getPageTitle(to.meta.title, to);
  if (to.meta.client) {
    return true;
=======
// 处理路由重定向
const handleRedirect = (to, userStore) => {
  if (router.hasRoute(userStore.userInfo.authority.defaultRouter)) {
    return { ...to, replace: true }
  }
  return { path: '/layout/404' }
}

// 路由守卫
router.beforeEach(async (to, from) => {
  const userStore = useUserStore()
  const routerStore = useRouterStore()
  const token = userStore.token

  Nprogress.start()

  // 处理元数据和缓存
  to.meta.matched = [...to.matched]
  await handleKeepAlive(to)

  // 设置页面标题
  document.title = getPageTitle(to.meta.title, to)

  if (to.meta.client) {
    return true
>>>>>>> main
  }

  // 白名单路由处理
  if (WHITE_LIST.includes(to.name)) {
    if (token) {
<<<<<<< HEAD
      if (!routerStore.asyncRouterFlag && whiteList.indexOf(from.name) < 0) {
        await getRouter(userStore);
      }
      // token 可以解析但是却是不存在的用户 id 或角色 id 会导致无限调用
      if (userStore.userInfo?.authority?.defaultRouter != null) {
        if (router.hasRoute(userStore.userInfo.authority.defaultRouter)) {
          return { name: userStore.userInfo.authority.defaultRouter };
        } else {
          return { path: "/layout/404" };
        }
      } else {
        // 强制退出账号
        userStore.ClearStorage();
        return {
          name: "Login",
          query: {
            redirect: document.location.hash,
          },
        };
      }
    } else {
      return true;
    }
  } else {
    // 不在白名单中并且已经登录的时候
    if (token) {
      // 添加flag防止多次获取动态路由和栈溢出
      if (!routerStore.asyncRouterFlag && whiteList.indexOf(from.name) < 0) {
        await getRouter(userStore);
        if (userStore.token) {
          if (router.hasRoute(userStore.userInfo.authority.defaultRouter)) {
            return { ...to, replace: true };
          } else {
            return { path: "/layout/404" };
          }
        } else {
          return {
            name: "Login",
            query: { redirect: to.href },
          };
        }
      } else {
        if (to.matched.length) {
          return true;
        } else {
          return { path: "/layout/404" };
        }
=======
      if(!routerStore.asyncRouterFlag){
        await setupRouter(userStore)
      }
      if(userStore.userInfo.authority.defaultRouter){
        return { name: userStore.userInfo.authority.defaultRouter }
>>>>>>> main
      }
    }
    return  true
  }

  // 需要登录的路由处理
  if (token) {
    // 处理需要跳转到首页的情况
    if (sessionStorage.getItem('needToHome') === 'true') {
      sessionStorage.removeItem('needToHome')
      return { path: '/' }
    }

    // 处理异步路由
    if (!routerStore.asyncRouterFlag && !WHITE_LIST.includes(from.name)) {
      const setupSuccess = await setupRouter(userStore)

      if (setupSuccess && userStore.token) {
        return handleRedirect(to, userStore)
      }

      return {
<<<<<<< HEAD
        name: "Login",
        query: {
          redirect: document.location.hash,
        },
      };
=======
        name: 'Login',
        query: { redirect: to.href }
      }
>>>>>>> main
    }

    return to.matched.length ? true : { path: '/layout/404' }
  }

  // 未登录跳转登录页
  return {
    name: 'Login',
    query: {
      redirect: document.location.hash
    }
  }
});

// 路由加载完成
router.afterEach(() => {
<<<<<<< HEAD
  // 路由加载完成后关闭进度条
  document.getElementsByClassName("main-cont main-right")[0]?.scrollTo(0, 0);
  Nprogress.done();
});

router.onError(() => {
  // 路由发生错误后销毁进度条
  Nprogress.remove();
});
=======
  document.querySelector('.main-cont.main-right')?.scrollTo(0, 0)
  Nprogress.done()
})

// 路由错误处理
router.onError((error) => {
  console.error('Router error:', error)
  Nprogress.remove()
})

// 移除初始加载动画
removeLoading()
>>>>>>> main
