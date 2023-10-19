<script setup>
import { useRouter } from 'vue-router'
import { state } from '@/stores/localstorage'

// useRouter 返回当前的router实例对象
const router = useRouter()
console.log(router)

// VueUse基于Vue组合式API的实用工具集 使用localStorage将Login对象变成响应式的数据
console.log(state)

const Home = () => {
  router.push({ name: 'home' })
}
const Login = () => {
  router.push({ name: 'LoginPage' })
}
const Logout = () => {
  state.value.isLogin = false
  router.push({ name: 'home' })
}
const JumpToFrontend = () => {
  router.push({ name: 'FrontendBlogList' })
}
const JumpToBackend = () => {
  router.push({ name: 'BackendBlogList' })
}
</script>

<template>
  <div>
    <!-- 导航菜单 -->
    <div class="menu">
      <a-menu mode="horizontal" :default-selected-keys="['0']">
        <!-- logo显示区 -->
        <a-typography-text class="logo"><icon-bytedance-color /> Vblog</a-typography-text>
        <!-- 竖直分割线 -->
        <a-divider direction="vertical" />
        <!-- 游客访问区 -->
        <a-menu-item key="0" @click="Home">首页</a-menu-item>
        <a-menu-item key="1" @click="JumpToFrontend">博客</a-menu-item>
        <!-- 登陆后可见 -->
        <a-menu-item v-if="state.isLogin" key="2" @click="JumpToBackend">后台管理</a-menu-item>
        <!-- 搜索 -->
        <a-space style="padding-left: 80px">
          <a-input-search placeholder="请输入文章标题或作者名" search-button>
            <template #button-icon>
              <icon-search />
            </template>
            <template #button-default>
              搜索
            </template>
          </a-input-search>
        </a-space>
        <!-- 用户 -->
        <a-space style="float: right">
          <a-button v-if="!state.isLogin" @click="Login" shape="circle">
            <icon-user />
          </a-button>
          <!-- 登陆之后显示 -->
          <a-button v-if="state.isLogin" @click="Logout">退出</a-button>
        </a-space>
      </a-menu>
    </div>

    <!--Vue-Router是Vue.js的官方路由库，Vue-Router的<router-view>组件可以根据页面URL的变化实现组件的动态渲染-->
    <div class="container">
      <router-view></router-view>
    </div>
  </div>
</template>

<style scoped>
.container {
  display: flex;
  justify-content: center;
  align-items: center;
}

.menu {
  box-sizing: border-box;
  width: 100%;
  padding: 10px;
  background-color: var(--color-neutral-2);
}

.logo {
  font-size: 16px;
  color: #555a65;
}
</style>
