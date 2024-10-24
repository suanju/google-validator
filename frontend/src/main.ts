import { router } from "@/router/router";
import { createApp } from "vue";

import "@mdi/font/css/materialdesignicons.css";
import "@/style.css";
import 'virtual:uno.css'
import App from "./App.vue";
import { createPinia } from "pinia";
import piniaPluginPersistedstate from 'pinia-plugin-persistedstate'

//arco
import ArcoVue from '@arco-design/web-vue';
import '@arco-design/web-vue/dist/arco.css';

const pinia = createPinia();
pinia.use(piniaPluginPersistedstate)
createApp(App).use(router).use(pinia).use(ArcoVue).mount("#app");
