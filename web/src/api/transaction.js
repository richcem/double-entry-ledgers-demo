import apiClient from './client';

export const createTransaction = async (transactionData) => {
  const response = await apiClient.post('/transactions', transactionData);
  return response.data;
};

export const getTransactions = async () => {
  const response = await apiClient.get('/transactions');
  return response.data;
};

export const createTransfer = async (transferData) => {
  const response = await apiClient.post('/transactions/transfer', transferData);
  return response.data;
};

export const createDeposit = async (depositData) => {
  const response = await apiClient.post('/transactions/deposit', depositData);
  return response.data;
};