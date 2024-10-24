<template>
  <div class="flex flex-col">
    <div class="h-8 w-full flex flex-row-reverse z-36" style="-webkit-app-region: drag">
      <WindowControlBar />
    </div>
    <div class="font-sans">
      <div class="fixed right-5 bottom-5 z-1000">
        <a-button type="primary" shape="circle" size="large" @click="showModal = true">
          <icon-plus />
        </a-button>
      </div>
    </div>

    <!-- 弹出的模态框 -->
    <a-modal v-model:visible="showModal" title="设置密钥" @ok="handleOk" @cancel="handleCancel">
      <div class="flex flex-col items-center">

        <a-input v-model="secretKey" placeholder="请输入密钥" class="mt-4 w-full" />
      </div>
    </a-modal>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import { IconPlus } from '@arco-design/web-vue/es/icon';
import WindowControlBar from "@/components/window_control_bar/window_control_bar.vue";

// 控制模态框的显示状态
const showModal = ref(false);

// 密钥的输入值
const secretKey = ref('');

// 处理模态框确定按钮的点击事件
const handleOk = () => {
  console.log('密钥已设置:', secretKey.value);
  showModal.value = false;
};

// 处理模态框取消按钮的点击事件
const handleCancel = () => {
  showModal.value = false;
};
</script>

<style scoped>
/* 添加你需要的样式 */
</style>
