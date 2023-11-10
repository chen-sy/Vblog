// 引入样式
import '@arco-design/web-vue/dist/arco.css'
import './assets/main.css'

import { createApp } from 'vue'
import { createPinia } from 'pinia'

import App from './App.vue'
import router from './router'

const app = createApp(App)

app.use(createPinia())
app.use(router)

// 引入UI组件
import ArcoVue from '@arco-design/web-vue'
// 引入图标
import ArcoVueIcon from '@arco-design/web-vue/es/icon'
// 使用UI组件库
app.use(ArcoVue)
app.use(ArcoVueIcon)

app.mount('#app')
