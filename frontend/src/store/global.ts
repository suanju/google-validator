import { defineStore } from "pinia";
import { ref } from "vue";

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
      keysList.value.push({...key});
    };
    return {
      keysList,
      addKeys
    };
  },
  {
    persist: {
      storage: localStorage,
      paths: ["keysList"],
    },
  }
);
