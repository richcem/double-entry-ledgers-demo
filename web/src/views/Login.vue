<template>
  <div class="min-h-screen flex items-center justify-center bg-gray-50">
    <div class="bg-white p-8 rounded-lg shadow-md w-full max-w-md">
      <h2 class="text-2xl font-bold mb-6 text-center">用户登录</h2>
      
      <!-- 错误提示 -->
      <div v-if="error" class="bg-red-100 border border-red-400 text-red-700 px-4 py-3 rounded mb-4">
        {{ error }}
      </div>
      
      <!-- 登录表单 -->
      <form @submit.prevent="handleLogin">
        <div class="mb-4">
          <label class="block text-gray-700 text-sm font-bold mb-2" for="username">
            用户名
          </label>
          <input
            v-model="username"
            class="form-input w-full"
            id="username"
            type="text"
            required
            placeholder="请输入用户名"
          />
        </div>
        <div class="mb-6">
          <label class="block text-gray-700 text-sm font-bold mb-2" for="password">
            密码
          </label>
          <input
            v-model="password"
            class="form-input w-full"
            id="password"
            type="password"
            required
            placeholder="请输入密码"
          />
        </div>
        <div class="flex items-center justify-between">
          <button
            class="btn-primary w-full"
            type="submit"
            :disabled="loading"
          >
            <span v-if="loading" class="mr-2 animate-spin rounded-full h-4 w-4 border-b-2 border-white"></span>
            登录
          </button>
        </div>
      </form>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue';
import { useRouter } from 'vue-router';
import axios from 'axios';

const username = ref('');
const password = ref('');
const error = ref('');
const loading = ref(false);
const router = useRouter();

const handleLogin = async () => {
  error.value = '';
  loading.value = true;
  
  try {
    const response = await axios.post('/auth/login', {
      username: username.value,
      password: password.value
    });
    
    // 存储JWT令牌
    localStorage.setItem('token', response.data.token);
    // 登录成功后
    localStorage.setItem('username', response.data.user.username); 
    
    // 设置Axios默认Authorization头
    axios.defaults.headers.common['Authorization'] = `Bearer ${response.data.token}`;
    
    // 跳转到首页
    router.push('/');
  } catch (err) {
    error.value = err.response?.data?.error || '登录失败，请检查用户名和密码';
    console.error('Login error:', err);
  } finally {
    loading.value = false;
  }
};
</script>