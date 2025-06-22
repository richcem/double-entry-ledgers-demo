<template>
  <div class="max-w-2xl mx-auto p-6 bg-white rounded-lg shadow-md">
    <h1 class="text-2xl font-bold mb-6 text-gray-800">账户入账</h1>
    
    <div v-if="transactionStore.error" class="mb-4 p-3 bg-red-100 text-red-700 rounded">{{ transactionStore.error }}</div>
    <div v-if="transactionStore.successMessage" class="mb-4 p-3 bg-green-100 text-green-700 rounded">{{ transactionStore.successMessage }}</div>
    
    <form @submit.prevent="handleSubmit" class="space-y-4">
      <!-- 账户选择 -->
      <div>
        <label for="accountId" class="block text-sm font-medium text-gray-700 mb-1">
          选择账户 <span class="text-red-500">*</span>
        </label>
        <select
          id="accountId"
          v-model="deposit.account_id"
          class="w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500"
          required
          :disabled="accountStore.loading || transactionStore.loading"
        >
          <option value="">请选择账户</option>
          <option v-for="account in accountStore.accounts" :key="account.id" :value="account.id">
            {{ account.name }} (¥{{ account.balance.toFixed(2) }})
          </option>
        </select>
      </div>

      <!-- 入账金额 -->
      <div>
        <label for="amount" class="block text-sm font-medium text-gray-700 mb-1">
          入账金额 (元) <span class="text-red-500">*</span>
        </label>
        <input
          id="amount"
          type="number"
          v-model.number="deposit.amount"
          class="w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500"
          min="0.01"
          step="0.01"
          required
          :disabled="loading"
        >
        <p v-if="amountError" class="mt-1 text-sm text-red-600">{{ amountError }}</p>
      </div>

      <!-- 入账描述 -->
      <div>
        <label for="description" class="block text-sm font-medium text-gray-700 mb-1">
          入账描述
        </label>
        <textarea
          id="description"
          v-model="deposit.description"
          class="w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500"
          rows="3"
          :disabled="loading"
        ></textarea>
      </div>

      <!-- 提交按钮 -->
      <div class="pt-2">
        <button
          type="submit"
          :disabled="loading"
          class="w-full flex justify-center py-2 px-4 border border-transparent rounded-md shadow-sm text-sm font-medium text-white bg-indigo-600 hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500 disabled:opacity-50"
        >
          <span v-if="loading" class="mr-2 animate-spin h-4 w-4 border-2 border-white border-opacity-50 border-t-white rounded-full"></span>
          确认入账
        </button>
      </div>
    </form>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue';
import { useRouter } from 'vue-router';
import { useAccountStore } from '@/store/account';
import { useTransactionStore } from '@/store/transaction';

const amountError = ref('');
const loading = ref(false);

const router = useRouter();
const accountStore = useAccountStore();
const transactionStore = useTransactionStore();

// 入账表单数据
const deposit = ref({
  account_id: '',
  amount: null,
  description: ''
});

// 验证表单数据
const validateForm = () => {
  transactionStore.error = '';
  let amountError = '';

  if (!deposit.value.account_id) {
    transactionStore.error = '请选择入账账户';
    return false;
  }

  if (!deposit.value.amount || deposit.value.amount <= 0) {
    amountError = '请输入有效的入账金额';
    return false;
  }

  return true;
};

// 处理入账提交
const handleSubmit = async () => {
  if (!validateForm()) return;

  const result = await transactionStore.performDeposit(deposit.value);
  if (result) {
    // 重置表单
    deposit.value = {
      account_id: '',
      amount: null,
      description: ''
    };

    // 2秒后可以选择跳转到交易记录页面
    setTimeout(() => {
      if (confirm('入账成功，是否前往交易记录页面查看？')) {
        router.push('/transactions');
      }
    }, 2000);
  }
};

// 组件挂载时获取账户列表
onMounted(() => {
  accountStore.fetchAccounts();
});
</script>