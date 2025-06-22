import { defineStore } from 'pinia';
import { 
  createTransaction, 
  getTransactions, 
  createTransfer, 
  createDeposit 
} from '../api/transaction';

export const useTransactionStore = defineStore('transaction', {
  state: () => ({
    transactions: [],
    loading: false,
    error: null
  }),
  
  actions: {
    async fetchTransactions() {
      this.loading = true;
      this.error = null;
      try {
        this.transactions = await getTransactions();
        return this.transactions;
      } catch (err) {
        this.error = err.response?.data?.error || '获取交易记录失败';
        throw err;
      } finally {
        this.loading = false;
      }
    },
    
    async createNewTransaction(transactionData) {
      this.loading = true;
      this.error = null;
      try {
        const newTransaction = await createTransaction(transactionData);
        this.transactions.push(newTransaction);
        return newTransaction;
      } catch (err) {
        this.error = err.response?.data?.error || '创建交易失败';
        throw err;
      } finally {
        this.loading = false;
      }
    },
    
    async performTransfer(transferData) {
      this.loading = true;
      this.error = null;
      try {
        const result = await createTransfer(transferData);
        // 如果需要，可以在这里更新本地交易列表
        return result;
      } catch (err) {
        this.error = err.response?.data?.error || '转账失败';
        throw err;
      } finally {
        this.loading = false;
      }
    },
    
    async performDeposit(depositData) {
      this.loading = true;
      this.error = null;
      try {
        const result = await createDeposit(depositData);
        // 如果需要，可以在这里更新本地交易列表
        return result;
      } catch (err) {
        this.error = err.response?.data?.error || '存款失败';
        throw err;
      } finally {
        this.loading = false;
      }
    }
  }
});
