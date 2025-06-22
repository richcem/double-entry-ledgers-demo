import { createApp } from 'vue'
import App from './App.vue'
import pinia from './store'  // 新增
import router from './router'  // 导入路由实例

import './assets/style.css'

createApp(App)
  .use(pinia)  // 新增
  .use(router)  // 使用导入的路由实例
  .mount('#app')