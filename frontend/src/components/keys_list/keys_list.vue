<template>
  <div ref="cardList" class="flex flex-wrap gap-x-4">
    <a-card
      v-for="item in keysList"
      hoverable
      :style="{
        width: 'calc(33.33% - 1rem)',
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
            <div
              class="color-[#165dff] absolute top-9 font-extralight"
              @click="copy(item.declassify)"
            >
              {{ item.declassify }}
            </div>
          </div>
        </div>
        <a-progress
          size="mini"
          :color="{
            '0%': 'rgb(var(--primary-6))',
            '100%': 'rgb(var(--success-6))',
          }"
          :percent="item.validity / 30"
        />
      </div>
    </a-card>
  </div>
</template>

<script lang="ts" setup>
import { useGlobalStore } from "@/store/global";
import message from "@arco-design/web-vue/es/message";
import { useDraggable } from "vue-draggable-plus";
import { processKeys } from "@/utils/totp";
import { ref, toRaw } from "vue";

const cardList = ref();
const globalStore = useGlobalStore();
const keysList = ref(<any>[]);

useDraggable(cardList, keysList, {
  animation: 150,
  onStart() {
    console.log("start");
  },
  onUpdate() {
    globalStore.keysList = keysList.value;
    console.log("update");
  },
});
const copy = async (text: string): Promise<void> => {
  try {
    await navigator.clipboard.writeText(text);
    message.success("复制成功");
  } catch (err) {
    message.success("复制失败");
    console.error("Failed to copy text: ", err);
  }
};

setInterval(async () => {
  keysList.value = await processKeys(toRaw(globalStore.keysList));
}, 100);
</script>
