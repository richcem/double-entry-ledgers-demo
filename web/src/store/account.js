import { defineStore } from 'pinia';
// 导入api/account.js中的所有接口函数
import { getAccounts, createAccount, updateAccount, deleteAccount } from '@/api/account';

export const useAccountStore = defineStore('account', {
  state: () => ({
    accounts: [],
    loading: false,
    error: null
  }),
  
  actions: {
    async fetchAccounts() {
      this.loading = true;
      this.error = null;
      try {
        // 使用api/account.js中的getAccounts函数
        this.accounts = await getAccounts();
      } catch (err) {
        this.error = '获取账户数据失败: ' + (err.response?.data?.error || err.message);
        console.error('Error fetching accounts:', err);
        throw err;
      } finally {
        this.loading = false;
      }
    },
    
    async createAccount(accountData) {
      this.loading = true;
      this.error = null;
      try {
        // 使用api/account.js中的createAccount函数
        const newAccount = await createAccount(accountData);
        this.accounts.push(newAccount);
        return newAccount;
      } catch (err) {
        this.error = '创建账户失败: ' + (err.response?.data?.error || err.message);
        console.error('Error creating account:', err);
        throw err;
      } finally {
        this.loading = false;
      }
    },
    
    async updateAccount(id, accountData) {
      this.loading = true;
      this.error = null;
      try {
        // 使用api/account.js中的updateAccount函数
        const updatedAccount = await updateAccount(id, accountData);
        const index = this.accounts.findIndex(account => account.id === id);
        if (index !== -1) {
          this.accounts[index] = updatedAccount;
        }
        return updatedAccount;
      } catch (err) {
        this.error = '更新账户失败: ' + (err.response?.data?.error || err.message);
        console.error('Error updating account:', err);
        throw err;
      } finally {
        this.loading = false;
      }
    },
    
    async deleteAccount(id) {
      this.loading = true;
      this.error = null;
      try {
        // 使用api/account.js中的deleteAccount函数
        await deleteAccount(id);
        this.accounts = this.accounts.filter(account => account.id !== id);
      } catch (err) {
        this.error = '删除账户失败: ' + (err.response?.data?.error || err.message);
        console.error('Error deleting account:', err);
        throw err;
      } finally {
        this.loading = false;
      }
    }
  }
});