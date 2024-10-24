import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import UnoCSS from 'unocss/vite'  
import { resolve} from "path";

// https://vitejs.dev/config/
export default defineConfig({
  resolve: {
    alias: {
      "@": resolve(__dirname, "./src"),
      "@wails": resolve(__dirname, "./wailsjs")
    },
  },
  plugins: [vue(),UnoCSS()],
})
