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
// 新增：导入登录组件
import Login from './views/Login.vue';

// 配置路由
const routes = [
  { path: '/', component: Home, meta: { requiresAuth: true, layout: 'default' } },
  { path: '/accounts', component: AccountQuery, meta: { requiresAuth: true, layout: 'default' } },
  { path: '/transfer', component: TransferFunds, meta: { requiresAuth: true, layout: 'default' } },
  { path: '/transactions', component: TransactionHistory, meta: { requiresAuth: true, layout: 'default' } },
  // 登录页面不需要认证，使用empty布局
  { path: '/login', component: Login, meta: { requiresAuth: false, layout: 'empty' } }
];

const router = createRouter({
  history: createWebHistory(),
  routes
});

// 添加路由守卫控制访问权限
router.beforeEach((to, from, next) => {
  // 检查本地存储中是否有token
  const token = localStorage.getItem('token');
  
  // 定义不需要登录的白名单路由
  const whiteList = ['/login'];
  
  // 如果没有token且访问的不是白名单页面，则重定向到登录页
  if (!token && !whiteList.includes(to.path)) {
    next('/login');
  } else {
    next();
  }
});

// 配置Axios
axios.defaults.baseURL = '/api';

// 检查本地存储中的令牌并设置
const token = localStorage.getItem('token');
if (token) {
  axios.defaults.headers.common['Authorization'] = `Bearer ${token}`;
}

// 创建并挂载应用实例
createApp(App)
  .use(router)
  .mount('#app');