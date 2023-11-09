<template>
  <div class="login-form">
    <a-form :model="form" @submit="handleSubmit">
      <a-form-item>
        <div class="login-title">系统登录</div>
      </a-form-item>
      <a-form-item field="username" :rules="[{ required: true, message: '请输入用户名' }]"
        :validate-trigger="['change', 'input']" tooltip="请输入用户名" label="用户名">
        <a-input v-model="form.username" placeholder="请输入用户名" />
      </a-form-item>
      <a-form-item field="password" :rules="[
        { required: true, message: '请输入密码' },
        { minLength: 6, message: '密码长度不能低于6位数' }
      ]" :validate-trigger="['change', 'input']" tooltip="请输入密码" label="密码">
        <a-input-password v-model="form.password" placeholder="请输入密码" allow-clear />
      </a-form-item>
      <a-form-item>
        <a-button style="width: 100%" html-type="submit">登录</a-button>
      </a-form-item>
    </a-form>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { Message } from '@arco-design/web-vue'
import { useRouter } from 'vue-router'
import { state } from '@/stores/app'
import { LOGIN } from '@/api/token'

// 路由对象
const router = useRouter()

// 用户登录数据
const form = ref({
  username: '',
  password: ''
})


// 提交处理
const handleSubmit = async (data) => {
  console.log(data)
  if (!data.errors) {
    // 需要对接API
    const form = data.values
    try {
      let resp = await LOGIN(form)
      debugger
      if (resp) {
        Message.success(`登录成功`)
        let to = 'BackendBlogList'
        router.push({ name: to })
      }
      state.value.isLogin = true
      state.value.username = resp.username
    } catch (error) {
      Message.error(`登录失败: ${error}`)
    }
  }
}
</script>

<style lang="css" scoped>
.login-form {
  margin-top: 200px;
  width: 600px;
}

.login-title {
  display: flex;
  justify-content: center;
  width: 100%;
  font-size: 24px;
  color: #555a65;
}
</style>
