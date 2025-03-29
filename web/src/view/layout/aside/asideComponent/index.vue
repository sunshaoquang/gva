<template>
  <component
    :is="menuComponent"
    v-if="!routerInfo.hidden"
    :router-info="routerInfo"
  >
    <template v-if="routerInfo.children && routerInfo.children.length">
      <AsideComponent
        v-for="item in routerInfo.children"
        :key="item.name"
        :router-info="item"
      />
    </template>
  </component>
</template>

<script setup>
<<<<<<< HEAD
import MenuItem from "./menuItem.vue";
import AsyncSubmenu from "./asyncSubmenu.vue";
import { computed } from "vue";

defineOptions({
  name: "AsideComponent",
});

const props = defineProps({
  routerInfo: {
    type: Object,
    default: () => null,
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

const menuComponent = computed(() => {
  if (
    props.routerInfo.children &&
    props.routerInfo.children.filter((item) => !item.hidden).length
  ) {
    return AsyncSubmenu;
  } else {
    return MenuItem;
  }
});
=======
  import MenuItem from './menuItem.vue'
  import AsyncSubmenu from './asyncSubmenu.vue'
  import { computed } from 'vue'

  defineOptions({
    name: 'AsideComponent'
  })

  const props = defineProps({
    routerInfo: {
      type: Object,
      default: () => null
    },
    mode: {
      type: String,
      default: 'vertical'
    }
  })

  const menuComponent = computed(() => {
    if (
      props.routerInfo.children &&
      props.routerInfo.children.filter((item) => !item.hidden).length
    ) {
      return AsyncSubmenu
    } else {
      return MenuItem
    }
  })
>>>>>>> main
</script>
