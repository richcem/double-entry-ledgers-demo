import { createPinia } from 'pinia'

const pinia = createPinia()

export default pinia

export * from './auth';
export * from './account';
export * from './transaction';