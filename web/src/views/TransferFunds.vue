<template>
  <div class="card">
    <h2 class="text-2xl font-bold mb-6">账户转账</h2>

    <!-- 转账表单 -->
    <form @submit.prevent="handleSubmit" class="space-y-6">
      <!-- 错误提示 -->
      <div v-if="error" class="bg-red-100 border border-red-400 text-red-700 px-4 py-3 rounded">
        <p>{{ error }}</p>
      </div>

      <!-- 成功提示 -->
      <div v-if="successMessage" class="bg-green-100 border border-green-400 text-green-700 px-4 py-3 rounded">
        <p>{{ successMessage }}</p>
      </div>

      <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
        <!-- 转出账户 -->
        <div>
          <label for="fromAccount" class="block text-sm font-medium text-gray-700 mb-1">
            转出账户 <span class="text-red-500">*</span>
          </label>
          <select
            id="fromAccount"
            v-model="transfer.fromAccountId"
            class="form-input"
            required
            :disabled="loading"
          >
            <option value="">请选择转出账户</option>
            <option v-for="account in accounts" :key="account.id" :value="account.id">
              {{ account.name }} (¥{{ account.balance.toFixed(2) }})
            </option>
          </select>
        </div>

        <!-- 转入账户 -->
        <div>
          <label for="toAccount" class="block text-sm font-medium text-gray-700 mb-1">
            转入账户 <span class="text-red-500">*</span>
          </label>
          <select
            id="toAccount"
            v-model="transfer.toAccountId"
            class="form-input"
            required
            :disabled="loading"
          >
            <option value="">请选择转入账户</option>
            <option v-for="account in accounts" :key="account.id" :value="account.id">
              {{ account.name }} (¥{{ account.balance.toFixed(2) }})
            </option>
          </select>
        </div>

        <!-- 转账金额 -->
        <div>
          <label for="amount" class="block text-sm font-medium text-gray-700 mb-1">
            转账金额 (元) <span class="text-red-500">*</span>
          </label>
          <input
            id="amount"
            type="number"
            v-model.number="transfer.amount"
            class="form-input"
            min="0.01"
            step="0.01"
            required
            :disabled="loading"
          >
          <p v-if="amountError" class="mt-1 text-sm text-red-600">{{ amountError }}</p>
        </div>

        <!-- 转账描述 -->
        <div class="md:col-span-2">
          <label for="description" class="block text-sm font-medium text-gray-700 mb-1">
            转账描述
          </label>
          <textarea
            id="description"
            v-model="transfer.description"
            class="form-input"
            rows="3"
            :disabled="loading"
          ></textarea>
        </div>
      </div>

      <!-- 操作按钮 -->
      <div class="flex justify-end gap-4 pt-4 border-t">
        <button
          type="button"
          @click="resetForm()"
          class="btn-secondary"
          :disabled="loading"
        >
          重置
        </button>
        <button
          type="submit"
          class="btn-primary"
          :disabled="loading"
        >
          <span v-if="loading" class="inline-block animate-spin rounded-full h-4 w-4 border-b-2 border-white mr-2"></span>
          确认转账
        </button>
      </div>
    </form>
  </div>
</template>

<script setup>
import { ref, onMounted, watch } from 'vue';
import axios from 'axios';
import { useRouter } from 'vue-router';

const router = useRouter();
const accounts = ref([]);
const loading = ref(false);
const error = ref('');
const successMessage = ref('');
const amountError = ref('');

// 转账表单数据
const transfer = ref({
  fromAccountId: '',
  toAccountId: '',
  amount: null,
  description: ''
});

// 获取账户列表
const fetchAccounts = async () => {
  try {
    loading.value = true;
    const response = await axios.get('/accounts');
    accounts.value = response.data;
  } catch (err) {
    error.value = '获取账户列表失败: ' + (err.response?.data?.error || err.message);
    console.error('Error fetching accounts:', err);
  } finally {
    loading.value = false;
  }
};

// 验证表单数据
const validateForm = () => {
  error.value = '';
  amountError.value = '';

  if (!transfer.value.fromAccountId || !transfer.value.toAccountId) {
    error.value = '请选择转出和转入账户';
    return false;
  }

  if (transfer.value.fromAccountId === transfer.value.toAccountId) {
    error.value = '转出账户和转入账户不能相同';
    return false;
  }

  if (!transfer.value.amount || transfer.value.amount <= 0) {
    amountError.value = '请输入有效的转账金额';
    return false;
  }

  return true;
};

// 处理转账提交
const handleSubmit = async () => {
  if (!validateForm()) return;
  
  try {
    loading.value = true;
    const response = await transactionStore.createTransaction(transfer.value);
    
    // 在组件中处理路由跳转
    successMessage.value = '转账成功！交易ID: ' + response.id;
    setTimeout(() => {
      router.push('/transactions');
    }, 2000);
  } catch (err) {
    error.value = '转账失败: ' + (err.response?.data?.error || err.message);
    console.error('Transfer error:', err);
  } finally {
    loading.value = false;
  }
};

// 重置表单
const resetForm = () => {
  transfer.value = {
    fromAccountId: '',
    toAccountId: '',
    amount: null,
    description: ''
  };
  error.value = '';
  successMessage.value = '';
  amountError.value = '';
};

// 监听金额变化，清除金额错误提示
watch(
  () => transfer.value.amount,
  () => {
    if (amountError.value) {
      amountError.value = '';
    }
  }
);

// 组件挂载时获取账户数据
onMounted(() => {
  fetchAccounts();
});
</script>