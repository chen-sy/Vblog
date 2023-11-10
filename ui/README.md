# 前端框架搭建

### 项目初始化 [Vue](https://cn.vuejs.org/) 
```sh
npm create vue@latest
```  

### UI 组件 [Arco.Design](https://arco.design/) 

##### 安装
```sh
npm install --save-dev @arco-design/web-vue
```

##### 引入
```js
// 引入UI组件
import ArcoVue from '@arco-design/web-vue'
// 引入样式
import '@arco-design/web-vue/dist/arco.css'
// 引入图标
import ArcoVueIcon from '@arco-design/web-vue/es/icon'
// 使用UI组件库
app.use(ArcoVue)
app.use(ArcoVueIcon)
```

### HTTP 客户端 [Axios](https://www.axios-http.cn/) 

##### 安装
```sh
npm install axios
```

### Vue 实用工具集合 [VueUse](https://vueuse.org/)

##### 安装

```sh
npm install @vueuse/core
```

##### 引入
```js
// 使用useStorage创建一个响应式的存储值
import { useStorage } from '@vueuse/core'
```

### 日期库 [Day.js](https://dayjs.fenxianglu.cn/)

##### 安装

```sh
npm install dayjs
```

##### 引入
```js
import dayjs from 'dayjs'
```

###  Markdown 编辑器 [md-editor-v3]

##### 安装

```sh
npm install md-editor-v3
```

##### 引入
```js
import MdEditor from 'md-editor-v3'
import 'md-editor-v3/lib/style.css';
app.use(MdEditor)
```


#### 常用命令

```sh
# 安装
npm install
# 格式化代码
npm run format
# 运行
npm run dev
```  
