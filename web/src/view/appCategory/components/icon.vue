<template>
  <div>
    <el-select
      v-model="metaData.name"
      clearable
      filterable
      placeholder="请选择"
    >
      <template #prefix>
        <div :class="`ssqIcon ssqIcon-${metaData.icon}`" />
      </template>
      <el-option
        v-for="item in options"
        :key="item.label"
        class="select__option_item"
        :label="item.label"
        :value="item.label"
      >
        <span class="gva-icon" style="padding: 3px 0 0" :class="item.extend">
          <div :class="`ssqIcon ssqIcon-${item.extend}`" />
        </span>
        <span style="text-align: left">{{ item.label }}</span>
      </el-option>
    </el-select>
  </div>
</template>

<script setup>
import { ref, reactive, nextTick, watch } from 'vue'
import { getDictFunc, filterDict } from '@/utils/format'
import { showDictLabel } from '@/utils/dictionary'

defineOptions({
  name: 'AppCategoryIcon'
})

const props = defineProps({
  meta: {
    default: function () {
      return {}
    },
    type: Object
  }
})

// 获取字典
const options = ref([])
const metaData = ref(props.meta)

watch(
  () => metaData.value.name,
  (newVal) => {
    metaData.value.icon = showDictLabel(
      options.value,
      newVal,
      'label',
      'extend'
    )
  }
)
nextTick(async () => {
  options.value = await getDictFunc('io_icon_dict')
  if (!metaData.value.name) {
    metaData.value.name = options.value[0].label
  }
})
</script>

<style lang="scss">
.gva-icon {
  color: rgb(132, 146, 166);
  font-size: 14px;
  margin-right: 10px;
}

.select__option_item {
  display: flex;
  align-items: center;
  justify-content: flex-start;
}
</style>
