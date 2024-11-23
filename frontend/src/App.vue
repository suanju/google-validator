<template>
  <div class="w-screen h-screen flex flex-col">
    <div class="h-full">
      <a-spin class="h-full" :loading="loading.show" :tip="loading.text">
        <RouterView />
      </a-spin>
    </div>
  </div>
</template>

<script setup lang="ts">
import { useGlobalStore } from "@/store/global";
const loading = useGlobalStore().loading;
</script>
