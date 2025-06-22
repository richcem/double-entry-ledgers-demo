<template>
  <div class="min-h-screen flex flex-col bg-gray-50">
    <!-- 顶部导航栏 -->
    <header class="bg-blue-600 text-white shadow-md">
      <div class="container mx-auto px-4 py-3">
        <div class="flex justify-between items-center">
          <h1 class="text-xl font-bold">复式记账系统</h1>
          <div class="flex items-center space-x-8">
            <!-- 功能导航 -->
            <nav>
              <ul class="flex space-x-6 items-center">
                <li><router-link to="/" class="hover:text-blue-200 transition-colors">首页</router-link></li>
                <li><router-link to="/accounts" class="hover:text-blue-200 transition-colors">账户查询</router-link></li>
                <li><router-link to="/deposit" class="hover:text-blue-200 transition-colors">账户入账</router-link></li>
                <li><router-link to="/transfer" class="hover:text-blue-200 transition-colors">转账</router-link></li>
                <li><router-link to="/transactions" class="hover:text-blue-200 transition-colors">交易记录</router-link></li>
              </ul>
            </nav>
            
            <!-- 用户信息和登出按钮 -->
            <div class="flex items-center space-x-4">
              <span class="whitespace-nowrap">欢迎，{{ username }}</span>
              <button 
                @click="logout" 
                class="hover:text-blue-200 transition-colors px-3 py-1 rounded hover:bg-blue-700"
              >
                登出
              </button>
            </div>
          </div>
        </div>
      </div>
    </header>

    <!-- 主内容区域 - 使用router-view渲染页面内容 -->
    <main class="flex-grow container mx-auto px-4 py-6">
      <router-view />
    </main>

    <!-- 页脚 -->
    <footer class="bg-gray-800 text-white py-4">
      <div class="container mx-auto px-4 text-center">
        <p>双-entry记账系统 &copy; {{ new Date().getFullYear() }}</p>
      </div>
    </footer>
  </div>
</template>

<script setup>
import { ref } from 'vue';
import { useRouter } from 'vue-router';

const router = useRouter();
// 获取当前登录用户名（假设登录时已存储）
const username = ref(localStorage.getItem('username') || '用户');

// 登出功能实现
const logout = () => {
  // 清除本地存储的认证信息
  localStorage.removeItem('token');
  localStorage.removeItem('username');
  // 重定向到登录页
  router.push('/login');
};
</script>

<style scoped>
/* 布局相关样式可以在这里添加 */
</style>
