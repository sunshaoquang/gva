<template>
  <div class="search-component items-center">
    <div
      class="gvaIcon gvaIcon-refresh"
      :class="[reload ? 'reloading' : '']"
      @click="handleReload"
    />
    <Screenfull class="search-icon" />
    <div class="gvaIcon gvaIcon-customer-service" @click="toService" />
<<<<<<< HEAD
=======
    <el-switch
      v-model="isDark"
      :active-action-icon="Moon"
      :inactive-action-icon="Sunny"
      @change="handleDarkSwitch"
    />
>>>>>>> main
  </div>
</template>

<script setup>
<<<<<<< HEAD
import Screenfull from "@/view/layout/screenfull/index.vue";
import { emitter } from "@/utils/bus.js";
import { ref } from "vue";

defineOptions({
  name: "BtnBox",
});

const reload = ref(false);
const handleReload = () => {
  reload.value = true;
  emitter.emit("reload");
  setTimeout(() => {
    reload.value = false;
  }, 500);
};
const toService = () => {
  window.open("https://support.qq.com/product/371961");
};
</script>
<style scoped lang="scss">
.search-component {
  @apply inline-flex overflow-hidden text-center gap-5 mr-5;
  div {
    @apply cursor-pointer;
  }
  .el-input__inner {
    @apply border-b border-solid border-gray-300;
  }
  .el-dropdown-link {
    @apply cursor-pointer;
  }
}

.reload {
  font-size: 18px;
}

.reloading {
  animation: turn 0.5s linear infinite;
}

@keyframes turn {
  0% {
    transform: rotate(0deg);
=======
  import Screenfull from '@/view/layout/screenfull/index.vue'
  import { emitter } from '@/utils/bus.js'
  import { Sunny, Moon } from '@element-plus/icons-vue'
  import { ref, watchEffect } from 'vue'

  defineOptions({
    name: 'BtnBox'
  })
  const isDark = ref(localStorage.getItem('isDark') === 'true' || true)

  watchEffect(() => {
    if (isDark.value) {
      document.documentElement.classList.add('dark')
      localStorage.setItem('isDark', true)
    } else {
      document.documentElement.classList.remove('dark')
      localStorage.setItem('isDark', false)
    }
  })
  const reload = ref(false)
  const handleReload = () => {
    reload.value = true
    emitter.emit('reload')
    setTimeout(() => {
      reload.value = false
    }, 500)
  }
  const toService = () => {
    window.open('https://support.qq.com/product/371961')
  }

  const handleDarkSwitch = (e) => {
    isDark.value = e
  }
</script>
<style scoped lang="scss">
  .search-component {
    @apply inline-flex overflow-hidden text-center gap-5 mr-5 text-black dark:text-gray-100;
    div {
      @apply cursor-pointer;
    }
    .el-input__inner {
      @apply border-b border-solid border-gray-300;
    }
    .el-dropdown-link {
      @apply cursor-pointer;
    }
>>>>>>> main
  }

  .reload {
    font-size: 18px;
  }

  .reloading {
    animation: turn 0.5s linear infinite;
  }

  @keyframes turn {
    0% {
      transform: rotate(0deg);
    }

<<<<<<< HEAD
  100% {
    transform: rotate(360deg);
  }
}
=======
    25% {
      transform: rotate(90deg);
    }

    50% {
      transform: rotate(180deg);
    }

    75% {
      transform: rotate(270deg);
    }

    100% {
      transform: rotate(360deg);
    }
  }
>>>>>>> main
</style>
