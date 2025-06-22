import { defineStore } from 'pinia';
import { login as loginApi, getCurrentUser as getCurrentUserApi } from '../api/auth';

export const useAuthStore = defineStore('auth', {
  state: () => ({
    user: null,
    token: localStorage.getItem('token'),
    loading: false,
    error: null
  }),
  
  actions: {
    async login(username, password) {
      this.loading = true;
      this.error = null;
      try {
        const data = await loginApi(username, password);
        this.token = data.token;
        this.user = data.user;
        localStorage.setItem('token', data.token);
        return data;
      } catch (err) {
        this.error = err.response?.data?.error || '登录失败，请检查用户名和密码';
        throw err;
      } finally {
        this.loading = false;
      }
    },
    
    async getCurrentUser() {
      this.loading = true;
      this.error = null;
      try {
        const user = await getCurrentUserApi();
        this.user = user;
        return user;
      } catch (err) {
        this.error = err.response?.data?.error || '获取用户信息失败';
        throw err;
      } finally {
        this.loading = false;
      }
    },
    
    logout() {
      this.user = null;
      this.token = null;
      localStorage.removeItem('token');
    }
  },
  
  getters: {
    isAuthenticated() {
      return !!this.token;
    }
  }
});