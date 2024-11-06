<template>
  <div class="flex flex-wrap gap-x-4">
    <a-card
      v-for="item in keysList"
      hoverable
      :style="{
        width: 'calc(33.33% - 1rem)', // 确保一行显示 3 个卡片，并留出 gap
        height: '70px',
        marginBottom: '20px',
        borderRadius: '12px',
      }"
    >
      <div class="flex items-center justify-between">
        <div :style="{ display: 'flex', alignItems: 'center', color: '#1D2129' }">
          <a-avatar
            :style="{ marginRight: '8px', backgroundColor: '#165DFF' }"
            :size="34"
          >
            <p>{{ item.name.slice(0, 1) }}</p>
          </a-avatar>
          <div class="flex flex-col">
            <div class="text-4.2 absolute top-3">{{ item.name }}</div>
            <div class="color-[#165dff] absolute top-9 font-extralight">
              {{ item.key }}
            </div>
          </div>
        </div>
        <a-progress size="mini" status="warning" :percent="item.validity / 30" />
      </div>
    </a-card>
  </div>
</template>

<script lang="ts" setup>
import { useGlobalStore } from "@/store/global";
import { processKeys } from "@/utils/totp";
import { ref, toRaw } from "vue";

const globalStore = useGlobalStore();
const keysList = ref(<any>[]);

setInterval(async () => {
  keysList.value = await processKeys(toRaw(globalStore.keysList));
  console.log(keysList.value);
}, 100);
</script>
