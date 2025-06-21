<template>
  <component :is="currentLayout" />
</template>

<script setup>
import { computed, watch } from 'vue';
import { useRoute, useRouter } from 'vue-router';
import DefaultLayout from './layouts/MainLayout.vue';
import EmptyLayout from './layouts/EmptyLayout.vue';

const route = useRoute();
const router = useRouter();

// 根据路由meta信息动态选择布局
const currentLayout = computed(() => {
  const token = localStorage.getItem('token');
  const requiresAuth = route.meta.requiresAuth !== false;
  
  // 如果需要认证但没有token，重定向到登录页
  if (requiresAuth && !token) {
    router.push('/login');
    return EmptyLayout;
  }
  
  return route.meta.layout === 'empty' ? EmptyLayout : DefaultLayout;
});

// 监听路由变化，确保布局正确切换
watch(route, () => {
  currentLayout.value;
});

// 登出功能实现
const logout = () => {
  // 清除本地存储的token
  localStorage.removeItem('token');
  // 重定向到登录页
  router.push('/login');
};
</script>

<style scoped>
/* 基础样式可以在这里添加 */
</style>