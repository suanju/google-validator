import { defineStore } from "pinia";
import { reactive, ref } from "vue";

export type KeysListItem = {
  name: string;
  key: string;
  keyType: string;
}

export type KeysList = Array<KeysListItem>

export const useGlobalStore = defineStore(
  "global",
  () => {
    const keysList = ref(<KeysList>[]);
    const addKeys = (key: KeysListItem) => {
      keysList.value.push({ ...key });
    };

    const loading = reactive({
      text: "截图中-请稍后 ...",
      show: false
    })
    return {
      keysList,
      addKeys,
      loading
    };
  },
  {
    persist: {
      storage: localStorage,
      paths: ["keysList"],
    },
  }
);
