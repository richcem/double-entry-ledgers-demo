<template>
  <div class="card">
    <h2 class="text-2xl font-bold mb-6">账户查询</h2>

    <!-- 搜索和筛选 -->
    <div class="mb-6 flex flex-col md:flex-row gap-4">
      <div class="flex-grow">
        <input
          v-model="searchQuery"
          type="text"
          placeholder="搜索账户名称..."
          class="form-input"
        />
      </div>
      <div class="flex gap-2">
        <button @click="fetchAccounts()" class="btn-primary whitespace-nowrap">
          刷新账户
        </button>
        <button @click="exportAccounts()" class="btn-secondary whitespace-nowrap">
          导出数据
        </button>
      </div>
    </div>

    <!-- 加载状态 -->
    <div v-if="loading" class="text-center py-10">
      <div class="inline-block animate-spin rounded-full h-8 w-8 border-b-2 border-blue-600 mb-2"></div>
      <p>正在加载账户数据...</p>
    </div>

    <!-- 错误提示 -->
    <div v-else-if="error" class="bg-red-100 border border-red-400 text-red-700 px-4 py-3 rounded mb-6">
      <p>{{ error }}</p>
    </div>

    <!-- 账户列表 -->
    <div v-else-if="accounts.length > 0" class="overflow-x-auto">
      <table class="min-w-full bg-white">
        <thead>
          <tr class="bg-gray-100 text-left">
            <th class="py-3 px-4 border-b">账户ID</th>
            <th class="py-3 px-4 border-b">账户名称</th>
            <th class="py-3 px-4 border-b">账户类型</th>
            <th class="py-3 px-4 border-b">余额</th>
            <th class="py-3 px-4 border-b">创建日期</th>
            <th class="py-3 px-4 border-b">操作</th>
          </tr>
        </thead>
        <tbody>
          <tr
            v-for="account in filteredAccounts"
            :key="account.id"
            class="hover:bg-gray-50 transition-colors"
          >
            <td class="py-3 px-4 border-b">{{ account.id }}</td>
            <td class="py-3 px-4 border-b">{{ account.name }}</td>
            <td class="py-3 px-4 border-b">{{ account.type }}</td>
            <td class="py-3 px-4 border-b font-medium"
              :class="account.balance >= 0 ? 'text-green-600' : 'text-red-600'">
              ¥{{ account.balance.toFixed(2) }}
            </td>
            <td class="py-3 px-4 border-b">{{ formatDate(account.created_at) }}</td>
            <td class="py-3 px-4 border-b">
              <button
                @click="viewAccountDetails(account.id)"
                class="text-blue-600 hover:underline"
              >
                详情
              </button>
            </td>
          </tr>
        </tbody>
      </table>
    </div>

    <!-- 无数据状态 -->
    <div v-else class="text-center py-10">
      <p class="text-gray-500">暂无账户数据，请先创建账户</p>
      <button class="mt-4 btn-primary" @click="goToCreateAccount()">
        创建账户
      </button>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted, computed } from 'vue';
import axios from 'axios';
import { useRouter } from 'vue-router';

const router = useRouter();
const accounts = ref([]);
const loading = ref(false);
const error = ref('');
const searchQuery = ref('');

// 获取账户列表
const fetchAccounts = async () => {
  try {
    loading.value = true;
    error.value = '';
    const response = await axios.get('/accounts');
    accounts.value = response.data;
  } catch (err) {
    error.value = '获取账户数据失败: ' + (err.response?.data?.error || err.message);
    console.error('Error fetching accounts:', err);
  } finally {
    loading.value = false;
  }
};

// 过滤账户
const filteredAccounts = computed(() => {
  if (!searchQuery.value) return accounts.value;
  const query = searchQuery.value.toLowerCase();
  return accounts.value.filter(account =>
    account.name.toLowerCase().includes(query) ||
    account.type.toLowerCase().includes(query)
  );
});

// 格式化日期
const formatDate = (dateString) => {
  if (!dateString) return '';
  const date = new Date(dateString);
  return date.toLocaleString('zh-CN', {
    year: 'numeric',
    month: '2-digit',
    day: '2-digit',
    hour: '2-digit',
    minute: '2-digit'
  });
};

// 查看账户详情
const viewAccountDetails = (accountId) => {
  // 在实际应用中，这里可以导航到账户详情页
  // router.push(`/accounts/${accountId}`);
  alert(`查看账户 ${accountId} 的详情`);
};

// 导出账户数据
const exportAccounts = () => {
  // 简单实现CSV导出
  const headers = ['账户ID,账户名称,账户类型,余额,创建日期\n'];
  const rows = accounts.value.map(account =>
    `${account.id},${account.name},${account.type},${account.balance},${formatDate(account.created_at)}`
  );
  const csvContent = headers.concat(rows).join('\n');
  const blob = new Blob([csvContent], { type: 'text/csv;charset=utf-8;' });
  const url = URL.createObjectURL(blob);
  const link = document.createElement('a');
  link.setAttribute('href', url);
  link.setAttribute('download', `accounts_${new Date().toISOString().slice(0,10)}.csv`);
  link.style.visibility = 'hidden';
  document.body.appendChild(link);
  link.click();
  document.body.removeChild(link);
};

// 前往创建账户页面
const goToCreateAccount = () => {
  // 在实际应用中，这里可以导航到创建账户页
  // router.push('/accounts/create');
  alert('跳转到创建账户页面');
};

// 组件挂载时获取账户数据
onMounted(() => {
  fetchAccounts();
});
</script>