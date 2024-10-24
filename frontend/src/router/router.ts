import { createRouter, createWebHistory } from "vue-router";

export const router = createRouter({
  history: createWebHistory(),
  routes: [
    { path: "", name: "Login", component: () => import("@/view/home/home.vue") },
  ],
});