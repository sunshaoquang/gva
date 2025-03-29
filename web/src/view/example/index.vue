<template>
  <div>
    <router-view v-slot="{ Component }">
      <transition mode="out-in" name="el-fade-in-linear">
        <keep-alive :include="routerStore.keepAliveRouters">
          <component :is="Component" />
        </keep-alive>
      </transition>
    </router-view>
  </div>
</template>

<script setup>
<<<<<<< HEAD
import { useRouterStore } from "@/pinia/modules/router";
const routerStore = useRouterStore();
defineOptions({
  name: "Example",
});
=======
  import { useRouterStore } from '@/pinia/modules/router'
  const routerStore = useRouterStore()
  defineOptions({
    name: 'Example'
  })
>>>>>>> main
</script>
