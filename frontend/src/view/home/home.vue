<template>
  <div class="flex flex-col">
    <div class="h-8 w-full flex flex-row-reverse z-36" style="-webkit-app-region: drag">
      <WindowControlBar></WindowControlBar>
    </div>
    <div class="font-sans">
      <div class="fixed right-8 bottom-8 z-1000">
        <a-button type="primary" shape="circle" size="large">
          <icon-plus />
        </a-button>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { IconPlus } from '@arco-design/web-vue/es/icon';
import WindowControlBar from "@/components/window_control_bar/window_control_bar.vue";
</script>

<style scoped>
::v-deep .arco-btn-size-large.arco-btn-shape-circle {
  width: 44px;
  height: 44px;
}
</style>
