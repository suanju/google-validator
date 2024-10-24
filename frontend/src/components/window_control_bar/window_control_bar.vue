<template>
  <div>
  
    <div class="h-8 w-100vw flex flex-row-reverse z-36" style="-webkit-app-region: drag">
      <div class="i-system-uicons-cross h-6 w-6 bg-gray-500 m-2" @click="cross"></div>
      <div class="i-system-uicons-minus h-6 w-6 bg-gray-500 m-2" @click="minus"></div>

     <div class="flex items-center justify-start ml-4 float-left flex-1">
      <a-image width="50" :src="Logo"/>
     </div>
    </div>
  </div>
</template>

<script lang="ts" setup>
import Logo from "@/assets/google.png";
import { Quit, WindowMinimise } from "@wails/runtime";

const cross = () => {
  Quit();
};
const minus = () => {
  WindowMinimise();
};
</script>
