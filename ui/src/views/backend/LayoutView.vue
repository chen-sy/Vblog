<template>
  <div>
    <!-- 顶部导航组件 -->
    <TopBar></TopBar>
    <!-- 需要显示子路由的页面, 必须要添加RouteView -->
    <div>博客后台布局页</div>
    <router-view></router-view>
  </div>
</template>

<script setup>
import TopBar from '../../components/TopBar.vue'
</script>

<style lang="css" scoped></style>