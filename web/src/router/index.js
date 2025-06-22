import { createRouter, createWebHistory } from 'vue-router';
import Home from '../views/Home.vue';
import AccountQuery from '../views/AccountQuery.vue';
import TransferFunds from '../views/TransferFunds.vue';
import TransactionHistory from '../views/TransactionHistory.vue';
import Login from '../views/Login.vue';
import Deposit from '../views/Deposit.vue';

// 合并所有路由配置
const routes = [
  { 
    path: '/', 
    component: Home, 
    meta: { requiresAuth: true, layout: 'default' } 
  },
  { 
    path: '/accounts', 
    component: AccountQuery, 
    meta: { requiresAuth: true, layout: 'default' } 
  },
  { 
    path: '/transfer', 
    component: TransferFunds, 
    meta: { requiresAuth: true, layout: 'default' } 
  },
  { 
    path: '/transactions', 
    component: TransactionHistory, 
    meta: { requiresAuth: true, layout: 'default' } 
  },
  { 
    path: '/login', 
    component: Login, 
    meta: { requiresAuth: false, layout: 'empty' } 
  },
  {
    path: '/deposit',
    name: 'Deposit',
    component: Deposit,
    meta: { requiresAuth: true }
  }
];

// 创建路由实例
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

export default router;