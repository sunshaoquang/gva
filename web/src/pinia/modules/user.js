import { login, getUserInfo, setSelfInfo } from "@/api/user";
import { jsonInBlacklist } from "@/api/jwt";
import router from "@/router/index";
import { ElLoading, ElMessage } from "element-plus";
import { defineStore } from "pinia";
import { ref, computed, watch } from "vue";
import { useRouterStore } from "./router";

export const useUserStore = defineStore("user", () => {
  const loadingInstance = ref(null);

  const userInfo = ref({
    uuid: "",
    nickName: "",
    headerImg: "",
    authority: {},
<<<<<<< HEAD
    sideMode: "dark",
    activeColor: "#4D70FF",
    baseColor: "#fff",
    appUser: {},
  });
  const token = ref(window.localStorage.getItem("token") || "");
=======
    sideMode: 'dark',
    activeColor: 'var(--el-color-primary)',
    baseColor: '#fff'
  })
  const token = ref(window.localStorage.getItem('token') || '')
>>>>>>> origin
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

  const ResetUserInfo = (value = {}) => {
    userInfo.value = {
      ...userInfo.value,
      ...value,
    };
  };
  /* 获取用户信息*/
  const GetUserInfo = async () => {
    const res = await getUserInfo();
    if (res.code === 0) {
      setUserInfo(res.data.userInfo);
    }
    return res;
  };
  /* 登录*/
  const LoginIn = async (loginInfo) => {
    loadingInstance.value = ElLoading.service({
      fullscreen: true,
      text: "登录中，请稍候...",
    });
    try {
      const res = await login(loginInfo);
      if (res.code === 0) {
<<<<<<< HEAD
        setUserInfo(res.data.user);
        setToken(res.data.token);
        const routerStore = useRouterStore();
        await routerStore.SetAsyncRouter();
        const asyncRouters = routerStore.asyncRouters;
        asyncRouters.forEach((asyncRouter) => {
          router.addRoute(asyncRouter);
        });
        await router.push({ name: userInfo.value.authority.defaultRouter });
        loadingInstance.value.close();
        return true;
=======
        setUserInfo(res.data.user)
        setToken(res.data.token)
        const routerStore = useRouterStore()
        await routerStore.SetAsyncRouter()
        const asyncRouters = routerStore.asyncRouters
        asyncRouters.forEach(asyncRouter => {
          router.addRoute(asyncRouter)
        })
        await router.replace({ name: userInfo.value.authority.defaultRouter })
        loadingInstance.value.close()
        return true
>>>>>>> origin
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
      token.value = "";
      sessionStorage.clear();
      localStorage.clear();
      router.push({ name: "Login", replace: true });
      window.location.reload();
    }
  };
  /* 清理数据 */
  const ClearStorage = async () => {
    token.value = "";
    sessionStorage.clear();
    localStorage.clear();
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
<<<<<<< HEAD
    if (
      userInfo.value.sideMode === "dark" ||
      userInfo.value.sideMode === "light"
    ) {
      return "#4D70FF";
    }
    return userInfo.activeColor;
  });
=======
    return 'var(--el-color-primary)'
  })
>>>>>>> origin

  watch(
    () => token.value,
    () => {
      window.localStorage.setItem("token", token.value);
    }
  );

  const userId = computed(() => userInfo.value.appUser?.ID ?? -1);
  const userIsAdmin = computed(() => {
    const authorityId = userInfo.value?.authorityId;
    const userName = userInfo.value?.userName;
    return authorityId === 888 && userName === "admin";
  });

  return {
    userInfo,
    token,
    NeedInit,
    ResetUserInfo,
    GetUserInfo,
    LoginIn,
    LoginOut,
    changeSideMode,
    mode,
    sideMode,
    setToken,
    baseColor,
    activeColor,
    loadingInstance,
    ClearStorage,
    userIsAdmin,
    userId,
  };
});
