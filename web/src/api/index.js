// 创建axios实例
const api = axios.create({
  baseURL: import.meta.env.VITE_API_BASE_URL,
  timeout: 5000
});

// 请求拦截器：转换请求参数为snake_case
api.interceptors.request.use(config => {
  if (config.data && typeof config.data === 'object') {
    config.data = camelToSnakeCase(config.data);
  }
  if (config.params && typeof config.params === 'object') {
    config.params = camelToSnakeCase(config.params);
  }
  return config;
});

// 工具函数：camelCase转snake_case
function camelToSnakeCase(obj) {
  if (typeof obj !== 'object' || obj === null) return obj;
  
  if (Array.isArray(obj)) {
    return obj.map(item => camelToSnakeCase(item));
  }
  
  const result = {};
  for (const key in obj) {
    if (obj.hasOwnProperty(key)) {
      const snakeKey = key.replace(/[A-Z]/g, letter => `_${letter.toLowerCase()}`);
      result[snakeKey] = camelToSnakeCase(obj[key]);
    }
  }
  return result;
}