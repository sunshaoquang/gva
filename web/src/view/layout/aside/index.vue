<template>
  <div>
    <normal-mode
      v-if="
        config.side_mode === 'normal' ||
        (device === 'mobile' && config.side_mode == 'head') ||
        (device === 'mobile' && config.side_mode == 'combination')
      "
    />
    <head-mode v-if="config.side_mode === 'head' && device !== 'mobile'" />
    <combination-mode
      v-if="config.side_mode === 'combination' && device !== 'mobile'"
      :mode="mode"
    />
  </div>
</template>

<script setup>
  import NormalMode from './normalMode.vue'
  import HeadMode from './headMode.vue'
  import CombinationMode from './combinationMode.vue'

<<<<<<< HEAD
defineOptions({
  name: 'Aside'
})

const route = useRoute()
const router = useRouter()

const userStore = useUserStore()
const routerStore = useRouterStore()

const theme = ref({})

const getTheme = () => {
  switch (userStore.sideMode) {
    case '#fff':
      theme.value = {
        background: '#fff',
        activeBackground: 'var(--el-color-primary)',
        activeText: '#fff',
        normalText: '#333',
        hoverBackground: 'rgba(64, 158, 255, 0.08)',
        hoverText: '#333'
      }
      break
    case '#191a23':
      theme.value = {
        background: '#191a23',
        activeBackground: 'var(--el-color-primary)',
        activeText: '#fff',
        normalText: '#fff',
        hoverBackground: 'rgba(64, 158, 255, 0.08)',
        hoverText: '#fff'
      }
      break
  }
}

getTheme()

const active = ref('')
watch(
  () => route,
  () => {
    active.value = route.meta.activeName || route.name
  },
  { deep: true }
)

watch(
  () => userStore.sideMode,
  () => {
    getTheme()
  }
)

const isCollapse = ref(false)
const initPage = () => {
  active.value = route.meta.activeName || route.name
  const screenWidth = document.body.clientWidth
  if (screenWidth < 1000) {
    isCollapse.value = !isCollapse.value
  }

  emitter.on('collapse', (item) => {
    isCollapse.value = item
=======
  defineProps({
    mode: {
      type: String,
      default: 'normal'
    }
>>>>>>> main
  })

  import { storeToRefs } from 'pinia'
  import { useAppStore } from '@/pinia'
  const appStore = useAppStore()
  const { config, device } = storeToRefs(appStore)
</script>
<<<<<<< HEAD

<style lang="scss">
.el-sub-menu__title:hover,
.el-menu-item:hover {
  background: transparent;
}

.el-scrollbar {
  .el-scrollbar__view {
    height: 100%;
  }
}
.menu-info {
  .menu-contorl {
    line-height: 52px;
    font-size: 20px;
    display: table-cell;
    vertical-align: middle;
  }
}
</style>
=======
>>>>>>> main
