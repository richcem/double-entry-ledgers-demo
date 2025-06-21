<template>
  <div class="grid grid-cols-1 md:grid-cols-3 gap-6">
    <!-- 欢迎信息 -->
    <div class="md:col-span-3 card">
      <h2 class="text-2xl font-bold mb-4">欢迎使用双-entry记账系统</h2>
      <p class="text-gray-600 mb-6">
        这是一个基于复式记账法的财务管理系统，帮助您记录和管理账户交易。
        通过本系统，您可以查询账户余额、进行账户间转账以及查看交易历史记录。
      </p>
      <div class="flex flex-wrap gap-4">
        <router-link to="/accounts" class="btn-primary">
          <i class="mr-2"></i>查看账户
        </router-link>
        <router-link to="/transfer" class="btn-primary">
          <i class="mr-2"></i>进行转账
        </router-link>
        <router-link to="/transactions" class="btn-primary">
          <i class="mr-2"></i>交易记录
        </router-link>
      </div>
    </div>

    <!-- 功能卡片 -->
    <div class="card hover:shadow-lg transition-shadow">
      <div class="text-blue-600 text-4xl mb-4">
        <i class=""></i>
      </div>
      <h3 class="text-xl font-semibold mb-2">账户查询</h3>
      <p class="text-gray-600 mb-4">查看所有账户余额和详细信息，掌握您的财务状况。</p>
      <router-link to="/accounts" class="text-blue-600 hover:underline">
        进入查询 <i class="ml-1"></i>
      </router-link>
    </div>

    <div class="card hover:shadow-lg transition-shadow">
      <div class="text-green-600 text-4xl mb-4">
        <i class=""></i>
      </div>
      <h3 class="text-xl font-semibold mb-2">账户转账</h3>
      <p class="text-gray-600 mb-4">在不同账户间进行资金转移，记录交易详情。</p>
      <router-link to="/transfer" class="text-green-600 hover:underline">
        进行转账 <i class="ml-1"></i>
      </router-link>
    </div>

    <div class="card hover:shadow-lg transition-shadow">
      <div class="text-purple-600 text-4xl mb-4">
        <i class=""></i>
      </div>
      <h3 class="text-xl font-semibold mb-2">交易记录</h3>
      <p class="text-gray-600 mb-4">查看所有历史交易记录，支持按日期和账户筛选。</p>
      <router-link to="/transactions" class="text-purple-600 hover:underline">
        查看记录 <i class="ml-1"></i>
      </router-link>
    </div>
  </div>
</template>

<style scoped>
/* 可以添加组件特定样式 */
</style>