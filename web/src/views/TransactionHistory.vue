<template>
  <div class="card">
    <h2 class="text-2xl font-bold mb-6">交易记录</h2>

    <!-- 筛选和搜索 -->
    <div class="mb-6 grid grid-cols-1 md:grid-cols-3 gap-4">
      <div>
        <label class="block text-sm font-medium text-gray-700 mb-1">开始日期</label>
        <input
          type="date"
          v-model="dateRange.start"
          class="form-input"
          :disabled="loading"
        >
      </div>
      <div>
        <label class="block text-sm font-medium text-gray-700 mb-1">结束日期</label>
        <input
          type="date"
          v-model="dateRange.end"
          class="form-input"
          :disabled="loading"
        >
      </div>
      <div>
        <label class="block text-sm font-medium text-gray-700 mb-1">账户筛选</label>
        <select
          v-model="filterAccountId"
          class="form-input"
          :disabled="loading"
        >
          <option value="">所有账户</option>
          <option v-for="account in accounts" :key="account.id" :value="account.id">
            {{ account.name }}
          </option>
        </select>
      </div>
    </div>

    <div class="mb-6 flex flex-col md:flex-row gap-4">
      <div class="flex-grow">
        <input
          v-model="searchQuery"
          type="text"
          placeholder="搜索描述或交易ID..."
          class="form-input"
          :disabled="loading"
        >
      </div>
      <div class="flex gap-2">
        <button @click="fetchTransactions()" class="btn-primary whitespace-nowrap">
          <span v-if="loading" class="inline-block animate-spin rounded-full h-4 w-4 border-b-2 border-white mr-2"></span>
          刷新记录
        </button>
        <button @click="exportTransactions()" class="btn-secondary whitespace-nowrap">
          导出记录
        </button>
      </div>
    </div>

    <!-- 加载状态 -->
    <div v-if="loading" class="text-center py-10">
      <div class="inline-block animate-spin rounded-full h-8 w-8 border-b-2 border-blue-600 mb-2"></div>
      <p>正在加载交易记录...</p>
    </div>

    <!-- 错误提示 -->
    <div v-else-if="error" class="bg-red-100 border border-red-400 text-red-700 px-4 py-3 rounded mb-6">
      <p>{{ error }}</p>
    </div>

    <!-- 交易记录列表 -->
    <div v-else-if="transactions.length > 0" class="overflow-x-auto">
      <table class="min-w-full bg-white">
        <thead>
          <tr class="bg-gray-100 text-left">
            <th class="py-3 px-4 border-b">交易ID</th>
            <th class="py-3 px-4 border-b">日期时间</th>
            <th class="py-3 px-4 border-b">描述</th>
            <th class="py-3 px-4 border-b">涉及账户</th>
            <th class="py-3 px-4 border-b">金额</th>
            <th class="py-3 px-4 border-b">状态</th>
          </tr>
        </thead>
        <tbody>
          <tr
            v-for="transaction in filteredTransactions"
            :key="transaction.id"
            class="hover:bg-gray-50 transition-colors"
          >
            <td class="py-3 px-4 border-b">{{ transaction.id }}</td>
            <td class="py-3 px-4 border-b">{{ formatDate(transaction.date) }}</td>
            <td class="py-3 px-4 border-b">{{ transaction.description }}</td>
            <td class="py-3 px-4 border-b">
              <div v-for="entry in transaction.entries" :key="entry.id">
                <span :class="entry.type === 'debit' ? 'text-green-600' : 'text-red-600'">{{ entry.type === 'debit' ? '收入: ' : '支出: ' }}</span>
                {{ entry.account.name }} (ID: {{ entry.account_id }})
              </div>
            </td>
            <td class="py-3 px-4 border-b font-medium">
              ¥{{ transaction.amount.toFixed(2) }}
            </td>
            <td class="py-3 px-4 border-b">
              <span :class="transaction.status === 'completed' ? 'bg-green-100 text-green-800' : 'bg-yellow-100 text-yellow-800'" class="px-2 py-1 rounded-full text-xs">
                {{ transaction.status === 'completed' ? '已完成' : transaction.status }}
              </span>
            </td>
          </tr>
        </tbody>
      </table>
    </div>

    <!-- 无数据状态 -->
    <div v-else class="text-center py-10">
      <p class="text-gray-500">暂无交易记录</p>
      <button class="mt-4 btn-primary" @click="goToTransfer()">
        进行转账
      </button>
    </div>

    <!-- 分页控件 -->
    <div v-if="transactions.length > 0 && !loading" class="flex justify-between items-center mt-6">
      <p class="text-sm text-gray-600">显示 {{ filteredTransactions.length }} 条，共 {{ transactions.length }} 条记录</p>
      <div class="flex gap-2">
        <button
          class="btn-secondary"
          @click="prevPage"
          :disabled="currentPage <= 1"
        >
          上一页
        </button>
        <span class="px-4 py-2">第 {{ currentPage }} 页</span>
        <button
          class="btn-secondary"
          @click="nextPage"
          :disabled="currentPage >= totalPages"
        >
          下一页
        </button>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted, computed, watch } from 'vue';
import axios from 'axios';
import { useRouter } from 'vue-router';

const router = useRouter();
const transactions = ref([]);
const accounts = ref([]);
const loading = ref(false);
const error = ref('');
const searchQuery = ref('');
const filterAccountId = ref('');
const dateRange = ref({
  start: '',
  end: ''
});

// 分页相关
const currentPage = ref(1);
const pageSize = ref(10);

// 获取交易记录
const fetchTransactions = async () => {
  try {
    loading.value = true;
    error.value = '';
    // 将查询参数改为snake_case格式
    const params = {
      filter_account_id: filterAccountId.value,
      date_start: dateRange.start,
      date_end: dateRange.end,
      search_query: searchQuery.value,
      page: currentPage.value,
      page_size: pageSize.value
    };
    const response = await axios.get('/transactions', { params });
    transactions.value = response.data;
    currentPage.value = 1; // 重置到第一页
  } catch (err) {
    error.value = '获取交易记录失败: ' + (err.response?.data?.error || err.message);
    console.error('Error fetching transactions:', err);
  } finally {
    loading.value = false;
  }
};

// 获取账户列表（用于筛选）
const fetchAccounts = async () => {
  try {
    const response = await axios.get('/accounts');
    accounts.value = response.data;
  } catch (err) {
    console.error('Error fetching accounts for filter:', err);
  }
};

// 格式化日期
const formatDate = (dateString) => {
  if (!dateString) return '';
  const date = new Date(dateString);
  return date.toLocaleString('zh-CN', {
    year: 'numeric',
    month: '2-digit',
    day: '2-digit',
    hour: '2-digit',
    minute: '2-digit',
    second: '2-digit'
  });
};

// 筛选交易记录
const filteredTransactions = computed(() => {
  let result = [...transactions.value];

  // 按账户筛选
  if (filterAccountId.value) {
    result = result.filter(transaction =>
      transaction.entries.some(entry => entry.account_id === Number(filterAccountId.value))
    );
  }

  // 按日期范围筛选
  if (dateRange.start) {
    const startDate = new Date(dateRange.start);
    result = result.filter(transaction => new Date(transaction.date) >= startDate);
  }
  if (dateRange.end) {
    const endDate = new Date(dateRange.end);
    endDate.setHours(23, 59, 59, 999);
    result = result.filter(transaction => new Date(transaction.date) <= endDate);
  }

  // 按搜索关键词筛选
  if (searchQuery.value) {
    const query = searchQuery.value.toLowerCase();
    result = result.filter(transaction =>
      transaction.description.toLowerCase().includes(query) ||
      transaction.id.toString().includes(query)
    );
  }

  // 按日期排序（最新的在前）
  result.sort((a, b) => new Date(b.date) - new Date(a.date));

  // 分页处理
  const startIndex = (currentPage.value - 1) * pageSize.value;
  return result.slice(startIndex, startIndex + pageSize.value);
});

// 总页数
const totalPages = computed(() => {
  return Math.ceil(filteredTransactions.value.length / pageSize.value);
});

// 分页控制
const prevPage = () => {
  if (currentPage.value > 1) {
    currentPage.value--;
  }
};

const nextPage = () => {
  if (currentPage.value < totalPages.value) {
    currentPage.value++;
  }
};

// 导出交易记录
const exportTransactions = () => {
  // 简单实现CSV导出
  const headers = ['交易ID,日期时间,描述,金额,状态,涉及账户\n'];
  const rows = filteredTransactions.value.map(transaction => {
    const accountsInfo = transaction.entries.map(entry =>
      `${entry.type === 'debit' ? '收入' : '支出'}:${entry.account.name}`
    ).join(';');
    return `${transaction.id},${formatDate(transaction.date)},${transaction.description},${transaction.amount},${transaction.status === 'completed' ? '已完成' : transaction.status},${accountsInfo}`;
  });
  const csvContent = headers.concat(rows).join('\n');
  const blob = new Blob([csvContent], { type: 'text/csv;charset=utf-8;' });
  const url = URL.createObjectURL(blob);
  const link = document.createElement('a');
  link.setAttribute('href', url);
  link.setAttribute('download', `transactions_${new Date().toISOString().slice(0,10)}.csv`);
  link.style.visibility = 'hidden';
  document.body.appendChild(link);
  link.click();
  document.body.removeChild(link);
};

// 前往转账页面
const goToTransfer = () => {
  router.push('/transfer');
};

// 监听筛选条件变化，重置到第一页
watch([filterAccountId, dateRange, searchQuery], () => {
  currentPage.value = 1;
});

// 组件挂载时获取数据
onMounted(() => {
  fetchTransactions();
  fetchAccounts();
});
</script>