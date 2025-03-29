<template>
<<<<<<< HEAD
  <div style="height: 80vh">
    <iframe
      width="100%"
      height="100%"
      :src="`${basePath}:${basePort}/form-generator/#/`"
      frameborder="0"
    />
  </div>
</template>

<script setup>
import { ref } from "vue";
const basePath = ref(import.meta.env.VITE_BASE_PATH);
const basePort = ref(import.meta.env.VITE_SERVER_PORT);
defineOptions({
  name: "FormGenerator",
});
=======
  <fc-designer ref="designer" :config="config" height="calc(100vh - 160px)" />
</template>

<script setup>
  import { ref } from 'vue'
  import FcDesigner from '@form-create/designer'

  defineOptions({
    name: 'FormGenerator'
  })

  const designer = ref(null)

  const config = {
    fieldReadonly: false
  }
>>>>>>> main
</script>
