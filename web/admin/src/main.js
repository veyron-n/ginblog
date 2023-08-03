import { createApp } from 'vue'
import ElementPlus from 'element-plus'
import 'element-plus/dist/index.css'
import App from './App.vue'
import router from './router/index'
import axios from 'axios'
import './assets/css/style.css'
const app = createApp(App)

axios.defaults.baseURL = 'http://localhost:3000/api/v1'

// 使用 Element Plus
app.use(ElementPlus)

// 将 axios 添加到应用实例的原型上，使其在所有组件中可用
app.config.globalProperties.$http = axios

// 关闭生产环境提示
app.config.productionTip = false

// 使用路由
app.use(router)

// 挂载应用实例到 #app
app.mount('#app')