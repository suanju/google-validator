import { defineStore } from "pinia";
import { ref } from "vue";

export const useglobalStore = defineStore(
  "global",
  () => {
    const selectedCatalog = ref("");
    return {
        selectedCatalog,
    };
  },
  {
    persist: {
      storage: localStorage,
      paths: ["selectedCatalog"],
    },
  }
);
