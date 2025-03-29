<template>
<<<<<<< HEAD
  <div>
    <el-button
      type="primary"
      class="drawer-container"
      icon="setting"
      @click="showSettingDrawer"
    />
    <el-drawer
      v-model="drawer"
      title="系统配置"
      :direction="direction"
      :before-close="handleClose"
    >
      <div class="setting_body">
        <div class="setting_card">
          <div class="setting_content">
            <div class="theme-box">
              <div class="item" @click="changeMode('light')">
                <div class="item-top">
                  <el-icon v-if="userStore.mode === 'light'" class="check">
                    <check />
                  </el-icon>
                  <img
                    src="https://gw.alipayobjects.com/zos/antfincdn/NQ%24zoisaD2/jpRkZQMyYRryryPNtyIC.svg"
                  />
                </div>
                <p>简约白</p>
              </div>
              <div class="item" @click="changeMode('dark')">
                <div class="item-top">
                  <el-icon v-if="userStore.mode === 'dark'" class="check">
                    <check />
                  </el-icon>
                  <img
                    src="https://gw.alipayobjects.com/zos/antfincdn/XwFOFbLkSM/LCkqqYNmvBEbokSDscrm.svg"
                  />
                </div>
                <p>商务黑</p>
              </div>
            </div>
          </div>
        </div>
      </div>
    </el-drawer>
  </div>
</template>

<script setup>
import { ref } from "vue";
import { useUserStore } from "@/pinia/modules/user";

defineOptions({
  name: "Setting",
});

const drawer = ref(false);
const direction = ref("rtl");

const userStore = useUserStore();

const handleClose = () => {
  drawer.value = false;
};
const showSettingDrawer = () => {
  drawer.value = true;
};
const changeMode = (e) => {
  if (e === null) {
    userStore.changeSideMode("dark");
    return;
  }
  userStore.changeSideMode(e);
};
</script>

<style lang="scss" scoped>
.drawer-container {
  transition: all 0.2s;
  &:hover {
    right: 0;
  }
  position: fixed;
  right: -20px;
  bottom: 15%;
  height: 40px;
  width: 40px;
  display: flex;
  align-items: center;
  justify-content: center;
  z-index: 999;
  color: #fff;
  border-radius: 4px 0 0 4px;
  cursor: pointer;
  -webkit-box-shadow: inset 0 0 6px rgba(0, 0, 0, 10%);
}
.setting_body {
  padding: 20px;
  .setting_card {
    margin-bottom: 20px;
  }
  .setting_content {
    margin-top: 20px;
    display: flex;
    flex-direction: column;
    > .theme-box {
      display: flex;
    }
    > .color-box {
      div {
        display: flex;
        flex-direction: column;
      }
    }
    .item {
      display: flex;
      align-items: center;
      justify-content: center;
      flex-direction: column;
      margin-right: 20px;
      .item-top {
        position: relative;
      }
      .check {
        position: absolute;
        font-size: 20px;
        color: #00afff;
        right: 10px;
        bottom: 10px;
      }
      p {
        text-align: center;
        font-size: 12px;
      }
    }
  }
}
=======
  <el-drawer
    v-model="drawer"
    title="系统配置"
    direction="rtl"
    :size="width"
    :show-close="false"
  >
    <template #header>
      <div class="flex justify-between items-center">
        <span class="text-lg">系统配置</span>
        <el-button type="primary" @click="saveConfig">保存配置</el-button>
      </div>
    </template>
    <div class="flex flex-col">
      <div class="mb-8">
        <Title title="默认主题"></Title>
        <div class="mt-2 text-sm p-2 flex items-center justify-center gap-2">
          <el-segmented
            v-model="config.darkMode"
            :options="options"
            size="default"
            @change="appStore.toggleDarkMode"
          />
        </div>
      </div>
      <div class="mb-8">
        <Title title="主题色"></Title>
        <div class="mt-2 text-sm p-2 flex items-center gap-2 justify-center">
          <div
            v-for="item in colors"
            :key="item"
            class="w-5 h-5 rounded cursor-pointer flex items-center justify-center"
            :style="`background:${item}`"
            @click="appStore.togglePrimaryColor(item)"
          >
            <el-icon v-if="config.primaryColor === item">
              <Select />
            </el-icon>
          </div>
          <el-color-picker
            v-model="customColor"
            @change="appStore.togglePrimaryColor"
          />
        </div>
      </div>
      <div class="mb-8">
        <Title title="主题配置"></Title>
        <div class="mt-2 text-md p-2 flex flex-col gap-2">
          <div class="flex items-center justify-between">
            <div>展示水印</div>
            <el-switch
              v-model="config.show_watermark"
              @change="appStore.toggleConfigWatermark"
            />
          </div>
          <div class="flex items-center justify-between">
            <div>灰色模式</div>
            <el-switch v-model="config.grey" @change="appStore.toggleGrey" />
          </div>
          <div class="flex items-center justify-between">
            <div>色弱模式</div>
            <el-switch
              v-model="config.weakness"
              @change="appStore.toggleWeakness"
            />
          </div>
          <div class="flex items-center justify-between">
            <div>菜单模式</div>
            <el-segmented
              v-model="config.side_mode"
              :options="sideModes"
              size="default"
              @change="appStore.toggleSideMode"
            />
          </div>

          <div class="flex items-center justify-between">
            <div>显示标签页</div>
            <el-switch
              v-model="config.showTabs"
              @change="appStore.toggleTabs"
            />
          </div>
          <div class="flex items-center justify-between gap-2">
            <div class="flex-shrink-0">页面切换动画</div>
            <el-select
              v-model="config.transition_type"
              @change="appStore.toggleTransition"
              class="w-40"
            >
              <el-option value="fade" label="淡入淡出" />
              <el-option value="slide" label="滑动" />
              <el-option value="zoom" label="缩放" />
              <el-option value="none" label="无动画" />
            </el-select>
          </div>
        </div>
      </div>

      <div class="mb-8">
        <Title title="layout 大小配置"></Title>
        <div class="mt-2 text-md p-2 flex flex-col gap-2">
          <div class="flex items-center justify-between mb-2">
            <div>侧边栏展开宽度</div>
            <el-input-number
              v-model="config.layout_side_width"
              :min="150"
              :max="400"
              :step="10"
            ></el-input-number>
          </div>
          <div class="flex items-center justify-between mb-2">
            <div>侧边栏收缩宽度</div>
            <el-input-number
              v-model="config.layout_side_collapsed_width"
              :min="60"
              :max="100"
            ></el-input-number>
          </div>
          <div class="flex items-center justify-between mb-2">
            <div>侧边栏子项高度</div>
            <el-input-number
              v-model="config.layout_side_item_height"
              :min="30"
              :max="50"
            ></el-input-number>
          </div>
        </div>
      </div>

      <!--      <el-alert type="warning" :closable="false">
        请注意，所有配置请保存到本地文件的
        <el-tag>config.json</el-tag> 文件中，否则刷新页面后会丢失配置
      </el-alert>-->
    </div>
  </el-drawer>
</template>

<script setup>
  import { useAppStore } from '@/pinia'
  import { storeToRefs } from 'pinia'
  import { ref, computed } from 'vue'
  import { ElMessage } from 'element-plus'
  import { setSelfSetting } from '@/api/user'
  import Title from './title.vue'
  const appStore = useAppStore()
  const { config, device } = storeToRefs(appStore)
  defineOptions({
    name: 'GvaSetting'
  })

  const width = computed(() => {
    return device.value === 'mobile' ? '100%' : '500px'
  })

  const colors = [
    '#EB2F96',
    '#3b82f6',
    '#2FEB54',
    '#EBEB2F',
    '#EB2F2F',
    '#2FEBEB'
  ]

  const drawer = defineModel('drawer', {
    default: true,
    type: Boolean
  })

  const options = ['dark', 'light', 'auto']
  const sideModes = [
    {
      label: '正常模式',
      value: 'normal'
    },
    {
      label: '顶部菜单栏模式',
      value: 'head'
    },
    {
      label: '组合模式',
      value: 'combination'
    }
  ]

  const saveConfig = async () => {
    /*const input = document.createElement("textarea");
  input.value = JSON.stringify(config.value);
  // 添加回车
  input.value = input.value.replace(/,/g, ",\n");
  document.body.appendChild(input);
  input.select();
  document.execCommand("copy");
  document.body.removeChild(input);
  ElMessage.success("复制成功, 请自行保存到本地文件中");*/
    const res = await setSelfSetting(config.value)
    if (res.code === 0) {
      localStorage.setItem('originSetting', JSON.stringify(config.value))
      ElMessage.success('保存成功')
      drawer.value = false
    }
  }

  const customColor = ref('')
</script>

<style lang="scss" scoped>
  ::v-deep(.el-drawer__header) {
    @apply border-gray-400 dark:border-gray-600;
  }
>>>>>>> main
</style>
