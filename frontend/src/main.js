import { createApp } from 'vue';
import { createRouter, createWebHistory } from 'vue-router';
import App from './App.vue';
import axios from 'axios';

// 导入CSS样式
import './assets/style.css';

// 导入页面组件
import AccountQuery from './views/AccountQuery.vue';
import TransferFunds from './views/TransferFunds.vue';
import TransactionHistory from './views/TransactionHistory.vue';
import Home from './views/Home.vue';

// 配置路由
const routes = [
  { path: '/', component: Home },
  { path: '/accounts', component: AccountQuery },
  { path: '/transfer', component: TransferFunds },
  { path: '/transactions', component: TransactionHistory }
];

const router = createRouter({
  history: createWebHistory(),
  routes
});

// 配置Axios
axios.defaults.baseURL = '/api';

// 创建并挂载应用实例
createApp(App)
  .use(router)
  .mount('#app');