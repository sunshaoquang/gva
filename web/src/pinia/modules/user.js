<<<<<<< HEAD
import { login, getUserInfo, setSelfInfo } from "@/api/user";
import { jsonInBlacklist } from "@/api/jwt";
import router from "@/router/index";
import { ElLoading, ElMessage } from "element-plus";
import { defineStore } from "pinia";
import { ref, computed, watch } from "vue";
import { useRouterStore } from "./router";
import cookie from "js-cookie";

export const useUserStore = defineStore("user", () => {
  const loadingInstance = ref(null);

  const userInfo = ref({
    uuid: "",
    nickName: "",
    headerImg: "",
    authority: {},
    sideMode: "dark",
    activeColor: "var(--el-color-primary)",
    baseColor: "#fff",
  });
  const token = ref(
    window.localStorage.getItem("token") || cookie.get("x-token") || ""
  );
  const setUserInfo = (val) => {
    userInfo.value = val;
  };

  const setToken = (val) => {
    token.value = val;
  };

  const NeedInit = () => {
    token.value = "";
    window.localStorage.removeItem("token");
    localStorage.clear();
    router.push({ name: "Init", replace: true });
  };
=======
import { login, getUserInfo } from '@/api/user'
import { jsonInBlacklist } from '@/api/jwt'
import router from '@/router/index'
import { ElLoading, ElMessage } from 'element-plus'
import { defineStore } from 'pinia'
import { ref, computed } from 'vue'
import { useRouterStore } from './router'
import { useCookies } from '@vueuse/integrations/useCookies'
import { useStorage } from '@vueuse/core'

import { useAppStore } from '@/pinia'

export const useUserStore = defineStore('user', () => {
  const appStore = useAppStore()
  const loadingInstance = ref(null)

  const userInfo = ref({
    uuid: '',
    nickName: '',
    headerImg: '',
    authority: {}
  })
  const token = useStorage('token', '')
  const xToken = useCookies('x-token')
  const currentToken = computed(() => token.value || xToken.value || '')

  const setUserInfo = (val) => {
    userInfo.value = val
    if (val.originSetting) {
      Object.keys(appStore.config).forEach((key) => {
        if (val.originSetting[key] !== undefined) {
          appStore.config[key] = val.originSetting[key]
        }
      })
    }
    console.log(appStore.config)
  }

  const setToken = (val) => {
    token.value = val
    xToken.value = val
  }

  const NeedInit = async () => {
    await ClearStorage()
    await router.push({ name: 'Init', replace: true })
  }
>>>>>>> main

  const ResetUserInfo = (value = {}) => {
    userInfo.value = {
      ...userInfo.value,
      ...value,
    };
  };
  /* 获取用户信息*/
  const GetUserInfo = async () => {
<<<<<<< HEAD
    const res = await getUserInfo();
=======
    const res = await getUserInfo()
>>>>>>> main
    if (res.code === 0) {
      setUserInfo(res.data.userInfo);
    }
    return res;
  };
  /* 登录*/
  const LoginIn = async (loginInfo) => {
<<<<<<< HEAD
    loadingInstance.value = ElLoading.service({
      fullscreen: true,
      text: "登录中，请稍候...",
    });
    try {
      const res = await login(loginInfo);
      if (res.code === 0) {
        setUserInfo(res.data.user);
        setToken(res.data.token);
        const routerStore = useRouterStore();
        await routerStore.SetAsyncRouter();
        const asyncRouters = routerStore.asyncRouters;
        asyncRouters.forEach((asyncRouter) => {
          router.addRoute(asyncRouter);
        });

        if (!router.hasRoute(userInfo.value.authority.defaultRouter)) {
          ElMessage.error("请联系管理员进行授权");
        } else {
          await router.replace({
            name: userInfo.value.authority.defaultRouter,
          });
        }

        loadingInstance.value.close();

        const isWin = ref(/windows/i.test(navigator.userAgent));
        if (isWin.value) {
          window.localStorage.setItem("osType", "WIN");
        } else {
          window.localStorage.setItem("osType", "MAC");
        }
        return true;
      }
    } catch (e) {
      loadingInstance.value.close();
    }
    loadingInstance.value.close();
  };
  /* 登出*/
  const LoginOut = async () => {
    const res = await jsonInBlacklist();
    if (res.code === 0) {
      await ClearStorage();
      router.push({ name: "Login", replace: true });
      window.location.reload();
    }
  };
  /* 清理数据 */
  const ClearStorage = async () => {
    token.value = "";
    sessionStorage.clear();
    localStorage.clear();
    cookie.remove("x-token");
  };
  /* 设置侧边栏模式*/
  const changeSideMode = async (data) => {
    const res = await setSelfInfo({ sideMode: data });
    if (res.code === 0) {
      userInfo.value.sideMode = data;
      ElMessage({
        type: "success",
        message: "设置成功",
      });
    }
  };

  const mode = computed(() => userInfo.value.sideMode);
  const sideMode = computed(() => {
    if (userInfo.value.sideMode === "dark") {
      return "#191a23";
    } else if (userInfo.value.sideMode === "light") {
      return "#fff";
    } else {
      return userInfo.value.sideMode;
    }
  });
  const baseColor = computed(() => {
    if (userInfo.value.sideMode === "dark") {
      return "#fff";
    } else if (userInfo.value.sideMode === "light") {
      return "#191a23";
    } else {
      return userInfo.value.baseColor;
    }
  });
  const activeColor = computed(() => {
    return "var(--el-color-primary)";
  });

  watch(
    () => token.value,
    () => {
      window.localStorage.setItem("token", token.value);
    }
  );
=======
    try {
      loadingInstance.value = ElLoading.service({
        fullscreen: true,
        text: '登录中，请稍候...'
      })

      const res = await login(loginInfo)

      if (res.code !== 0) {
        ElMessage.error(res.message || '登录失败')
        return false
      }
      // 登陆成功，设置用户信息和权限相关信息
      setUserInfo(res.data.user)
      setToken(res.data.token)

      // 初始化路由信息
      const routerStore = useRouterStore()
      await routerStore.SetAsyncRouter()
      const asyncRouters = routerStore.asyncRouters

      // 注册到路由表里
      asyncRouters.forEach((asyncRouter) => {
        router.addRoute(asyncRouter)
      })

      if (!router.hasRoute(userInfo.value.authority.defaultRouter)) {
        ElMessage.error('请联系管理员进行授权')
      } else {
        await router.replace({ name: userInfo.value.authority.defaultRouter })
      }

      const isWindows = /windows/i.test(navigator.userAgent)
      window.localStorage.setItem('osType', isWindows ? 'WIN' : 'MAC')

      // 全部操作均结束，关闭loading并返回
      return true
    } catch (error) {
      console.error('LoginIn error:', error)
      return false
    } finally {
      loadingInstance.value?.close()
    }
  }
  /* 登出*/
  const LoginOut = async () => {
    const res = await jsonInBlacklist()

    // 登出失败
    if (res.code !== 0) {
      return
    }

    await ClearStorage()

    // 把路由定向到登录页，无需等待直接reload
    router.push({ name: 'Login', replace: true })
    window.location.reload()
  }
  /* 清理数据 */
  const ClearStorage = async () => {
    token.value = ''
    xToken.value = ''
    sessionStorage.clear()
    localStorage.removeItem('originSetting')
  }
>>>>>>> main

  return {
    userInfo,
    token: currentToken,
    NeedInit,
    ResetUserInfo,
    GetUserInfo,
    LoginIn,
    LoginOut,
    setToken,
    loadingInstance,
    ClearStorage,
  };
});
