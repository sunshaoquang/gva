<template>
<<<<<<< HEAD
  <el-sub-menu ref="subMenu" :index="routerInfo.name">
    <template #title>
      <div v-if="!isCollapse" class="gva-subMenu">
=======
  <el-sub-menu
    ref="subMenu"
    :index="routerInfo.name"
    class="gva-sub-menu dark:text-slate-300 relative"
  >
    <template #title>
      <div
        v-if="!isCollapse"
        class="flex items-center"
        :style="{
          height: sideHeight
        }"
      >
>>>>>>> main
        <el-icon v-if="routerInfo.meta.icon">
          <component :is="routerInfo.meta.icon" />
        </el-icon>
        <span>{{ routerInfo.meta.title }}</span>
      </div>
      <template v-else>
        <el-icon v-if="routerInfo.meta.icon">
          <component :is="routerInfo.meta.icon" />
        </el-icon>
        <span>{{ routerInfo.meta.title }}</span>
      </template>
    </template>
    <slot />
  </el-sub-menu>
</template>

<script setup>
<<<<<<< HEAD
import { ref, watch } from "vue";

defineOptions({
  name: "AsyncSubmenu",
});

const props = defineProps({
  routerInfo: {
    default: function () {
      return null;
    },
    type: Object,
  },
  isCollapse: {
    default: function () {
      return false;
    },
    type: Boolean,
  },
  theme: {
    default: function () {
      return {};
    },
    type: Object,
  },
});

const activeBackground = ref(props.theme.activeBackground);
const activeText = ref(props.theme.activeText);
const normalText = ref(props.theme.normalText);
// const hoverBackground = ref(props.theme.hoverBackground)
// const hoverText = ref(props.theme.hoverText)

watch(
  () => props.theme,
  () => {
    activeBackground.value = props.theme.activeBackground;
    activeText.value = props.theme.activeText;
    normalText.value = props.theme.normalText;
    // hoverBackground.value = props.theme.hoverBackground
    // hoverText.value = props.theme.hoverText
  }
);
</script>

<style lang="scss" scoped>
.el-sub-menu {
  ::v-deep(.el-sub-menu__title) {
    padding: 6px;
    color: v-bind(normalText);
  }
}

.is-active:not(.is-opened) {
  ::v-deep(.el-sub-menu__title) .gva-subMenu {
    flex: 1;
    height: 100%;
    line-height: 44px;
    background: v-bind(activeBackground) !important;
    border-radius: 4px;
    box-shadow: 0 0 2px 1px v-bind(activeBackground) !important;
    i {
      color: v-bind(activeText);
    }
    span {
      color: v-bind(activeText);
    }
  }
}

.gva-subMenu {
  padding-left: 4px;
}
=======
  import { inject, computed } from 'vue'
  import { useAppStore } from '@/pinia'
  import { storeToRefs } from 'pinia'
  const appStore = useAppStore()
  const { config } = storeToRefs(appStore)

  defineOptions({
    name: 'AsyncSubmenu'
  })

  defineProps({
    routerInfo: {
      default: function () {
        return null
      },
      type: Object
    }
  })

  const isCollapse = inject('isCollapse', {
    default: false
  })

  const sideHeight = computed(() => {
    return config.value.layout_side_item_height + 'px'
  })
</script>

<style lang="scss">
  .gva-sub-menu {
    .el-sub-menu__title {
      height: v-bind('sideHeight') !important;
    }
  }
>>>>>>> main
</style>
