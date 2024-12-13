<template>
  <div class="flex flex-col home-element" ref="homeRef">
    <div class="h-8 w-full flex flex-row-reverse z-36" style="-webkit-app-region: drag">
      <WindowControlBar></WindowControlBar>
      <AddKeys :visible="addKeysVisible" @cloos="addKeysVisible = false"></AddKeys>
    </div>
    <!-- 添加按钮 -->
    <div class="font-sans">
      <div class="fixed right-8 bottom-8">
        <a-dropdown
          position="top"
          :popup-container="homeRef"
          @popup-visible-change="handlePopupVisibleChange"
        >
          <a-button type="primary" shape="circle" size="large">
            <icon-plus
              class="transition-transform duration-300 ease"
              :style="{ transform: pushButtonRotation }"
            />
          </a-button>
          <template #content>
            <a-doption @click="takeScreenshot">
              <div class="flex items-center justify-end">
                <div class="text-xs mr-2.4 bg-gray-100 px-2 py-1.2 rounded-xl">
                  扫描二维码
                </div>
                <icon-camera :size="20" />
              </div>
            </a-doption>
            <a-doption @click="addKeysVisible = !addKeysVisible">
              <div class="flex items-center justify-end">
                <div class="text-xs mr-2.4 bg-gray-100 px-2 py-1.2 rounded-xl">
                  输入设置密钥
                </div>
                <icon-edit :size="20" />
              </div>
            </a-doption>
          </template>
        </a-dropdown>
      </div>
    </div>
    <!-- 主体显示 -->
    <div class="mt-2 mx-2">
      <KeysList></KeysList>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from "vue";
import { IconPlus } from "@arco-design/web-vue/es/icon";
import { ScreenshotRecognize } from "@wails/go/backend/App";
import { useGlobalStore, type KeysListItem } from "@/store/global";
import { WindowHide, WindowShow } from "@wails/runtime/runtime";
import { parseOTPAuth } from "@/utils/totp";
import WindowControlBar from "@/components/window_control_bar/window_control_bar.vue";
import message from "@arco-design/web-vue/es/message";
import AddKeys from "@/components/add_keys/add_keys.vue";
import KeysList from "@/components/keys_list/keys_list.vue";

const homeRef = ref();
const globalStore = useGlobalStore()
const addKeysVisible = ref(false);
const pushButtonRotation = ref("rotate(0deg)");
const loading = useGlobalStore().loading;

const takeScreenshot = async () => {
  try {
    loading.show = true;
    WindowHide();
    //阻塞一会防止还没收起窗口
    await new Promise((resolve) => setTimeout(resolve, 300));
    const resp = await ScreenshotRecognize();
    const { otpArr, hasMatch } = parseOTPAuth(resp);
    if (!hasMatch) {
      message.error("未检出到验证器相关二维码~");
      return;
    }
    otpArr.forEach(item => {
      globalStore.addKeys(<KeysListItem>{
        name: item.name,
        key: item.key,
        keyType: "time",
      });
    });
    console.log(resp, otpArr);
    message.success("识别成功");
  } catch (error) {
    message.error("识别失败,可能当前屏幕中不包含二维码~");
    console.error("Error taking screenshot:", error);
  } finally {
    loading.show = false;
    WindowShow();
  }
};

const handlePopupVisibleChange = (visible: boolean) => {
  if (visible) {
    pushButtonRotation.value = "rotate(45deg)";
  } else {
    pushButtonRotation.value = "rotate(0deg)";
  }
};
</script>

<style scoped>
:deep(.arco-btn-size-large.arco-btn-shape-circle) {
  width: 44px;
  height: 44px;
}

:deep(.arco-dropdown) {
  background-color: transparent;
  box-shadow: none;
  border: none;
}

:deep(.arco-dropdown-option) {
  height: 40px;
  width: 100%;
}

:deep(.arco-dropdown-option:hover) {
  background-color: transparent;
}

:deep(.arco-dropdown-option-content) {
  width: 100%;
}

:deep(.arco-dropdown-list) {
  margin-bottom: 16px;
  padding-right: 16px;
}
</style>
