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
        <button @click="showCreateModal = true" class="btn-primary whitespace-nowrap">
          创建账户
        </button>
      </div>
    </div>

    <!-- 账户列表 -->
    <div v-if="accounts.length > 0" class="overflow-x-auto">
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
                class="text-blue-600 hover:underline mr-2"
              >
                详情
              </button>
              <button
                @click="editAccount(account)"
                class="text-green-600 hover:underline mr-2"
              >
                编辑
              </button>
              <button
                @click="confirmDeleteAccount(account.id)"
                class="text-red-600 hover:underline"
              >
                删除
              </button>
            </td>
          </tr>
        </tbody>
      </table>
    </div>

    <!-- 创建账户模态框 -->
    <div v-if="showCreateModal" class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center p-4">
      <div class="bg-white rounded-lg w-full max-w-md">
        <div class="p-6">
          <h3 class="text-xl font-bold mb-4">创建新账户</h3>
          <div v-if="error" class="bg-red-100 text-red-700 p-3 rounded mb-4">{{ error }}</div>
          <div class="space-y-4">
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-1">账户名称 *</label>
              <input v-model="accountForm.name" type="text" class="form-input" required>
            </div>
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-1">账户类型 *</label>
              <select v-model="accountForm.type" class="form-input" required>
                <option value="">请选择账户类型</option>
                <option v-for="type in accountTypes" :key="type.value" :value="type.value">{{ type.label }}</option>
              </select>
            </div>
          </div>
          <div class="mt-6 flex justify-end gap-3">
            <button @click="showCreateModal = false" class="btn-secondary">取消</button>
            <button @click="submitCreateForm" class="btn-primary">创建</button>
          </div>
        </div>
      </div>
    </div>

    <!-- 编辑账户模态框 -->
    <div v-if="showEditModal" class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center p-4">
      <div class="bg-white rounded-lg w-full max-w-md">
        <div class="p-6">
          <h3 class="text-xl font-bold mb-4">编辑账户</h3>
          <div v-if="error" class="bg-red-100 text-red-700 p-3 rounded mb-4">{{ error }}</div>
          <div class="space-y-4">
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-1">账户名称 *</label>
              <input v-model="accountForm.name" type="text" class="form-input" required>
            </div>
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-1">账户类型 *</label>
              <select v-model="accountForm.type" class="form-input" required>
                <option value="">请选择账户类型</option>
                <option v-for="type in accountTypes" :key="type.value" :value="type.value">{{ type.label }}</option>
              </select>
            </div>
          </div>
          <div class="mt-6 flex justify-end gap-3">
            <button @click="showEditModal = false" class="btn-secondary">取消</button>
            <button @click="submitEditForm" class="btn-primary">保存</button>
          </div>
        </div>
      </div>
    </div>

    <!-- 删除确认对话框 -->
    <div v-if="deleteDialogVisible" class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center p-4">
      <div class="bg-white rounded-lg w-full max-w-md p-6">
        <h3 class="text-xl font-bold mb-4">确认删除</h3>
        <p class="mb-6">确定要删除此账户吗？此操作不可撤销。</p>
        <div class="flex justify-end gap-3">
          <button @click="deleteDialogVisible = false" class="btn-secondary">取消</button>
          <button @click="deleteAccount()" class="btn-danger">删除</button>
        </div>
      </div>
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
const successMessage = ref('');

// 模态框状态
const showCreateModal = ref(false);
const showEditModal = ref(false);
const deleteDialogVisible = ref(false);
const currentAccount = ref(null);
const accountToDelete = ref(null);

// 表单数据
const accountForm = ref({
  name: '',
  type: ''
});

// 账户类型选项
const accountTypes = [
  { value: 'asset', label: '资产' },
  { value: 'liability', label: '负债' },
  { value: 'equity', label: '权益' },
  { value: 'income', label: '收入' },
  { value: 'expense', label: '支出' }
];

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

// 编辑账户
const editAccount = (account) => {
  currentAccount.value = { ...account };
  accountForm.value.name = account.name;
  accountForm.value.type = account.type;
  showEditModal.value = true;
};

// 确认删除账户
const confirmDeleteAccount = (id) => {
  accountToDelete.value = id;
  deleteDialogVisible.value = true;
};

// 删除账户
const deleteAccount = async () => {
  try {
    loading.value = true;
    await axios.delete(`/accounts/${accountToDelete.value}`);
    fetchAccounts(); // 重新获取账户列表
    deleteDialogVisible.value = false;
    successMessage.value = '账户删除成功';
    setTimeout(() => successMessage.value = '', 3000);
  } catch (err) {
    error.value = '删除账户失败: ' + (err.response?.data?.error || err.message);
  } finally {
    loading.value = false;
  }
};

// 提交创建账户表单
const submitCreateForm = async () => {
  try {
    loading.value = true;
    await axios.post('/accounts', accountForm.value);
    showCreateModal.value = false;
    fetchAccounts();
    accountForm.value = { name: '', type: '' };
    successMessage.value = '账户创建成功';
    setTimeout(() => successMessage.value = '', 3000);
  } catch (err) {
    error.value = '创建账户失败: ' + (err.response?.data?.error || err.message);
  } finally {
    loading.value = false;
  }
};

// 提交编辑账户表单
const submitEditForm = async () => {
  try {
    loading.value = true;
    await axios.put(`/accounts/${currentAccount.value.id}`, accountForm.value);
    showEditModal.value = false;
    fetchAccounts();
    successMessage.value = '账户更新成功';
    setTimeout(() => successMessage.value = '', 3000);
  } catch (err) {
    error.value = '更新账户失败: ' + (err.response?.data?.error || err.message);
  } finally {
    loading.value = false;
  }
};
</script>