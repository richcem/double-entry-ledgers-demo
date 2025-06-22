import apiClient from './client';

export const getAccounts = async () => {
  const response = await apiClient.get('/accounts');
  return response.data;
};

export const createAccount = async (accountData) => {
  const response = await apiClient.post('/accounts', accountData);
  return response.data;
};

export const updateAccount = async (id, accountData) => {
  const response = await apiClient.put(`/accounts/${id}`, accountData);
  return response.data;
};

export const deleteAccount = async (id) => {
  const response = await apiClient.delete(`/accounts/${id}`);
  return response.data;
};