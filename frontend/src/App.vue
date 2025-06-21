<template>
  <div class="min-h-screen flex flex-col bg-gray-50">
    <!-- 顶部导航栏 -->
    <header class="bg-blue-600 text-white shadow-md">
      <div class="container mx-auto px-4 py-3">
        <h1 class="text-xl font-bold">双-entry记账系统</h1>
        <nav class="mt-2">
          <ul class="flex space-x-4">
            <li><router-link to="/" class="hover:text-blue-200 transition-colors">首页</router-link></li>
            <li><router-link to="/accounts" class="hover:text-blue-200 transition-colors">账户查询</router-link></li>
            <li><router-link to="/transfer" class="hover:text-blue-200 transition-colors">转账</router-link></li>
            <li><router-link to="/transactions" class="hover:text-blue-200 transition-colors">交易记录</router-link></li>
          </ul>
        </nav>
      </div>
    </header>

    <!-- 主内容区域 -->
    <main class="flex-grow container mx-auto px-4 py-6">
      <router-view />
    </main>

    <!-- 页脚 -->
    <footer class="bg-gray-800 text-white py-4">
      <div class="container mx-auto px-4 text-center">
        <p>双-entry记账系统 &copy; { { new Date().getFullYear() } }</p>
      </div>
    </footer>
  </div>
</template>

<style scoped>
/* 基础样式可以在这里添加 */
</style>