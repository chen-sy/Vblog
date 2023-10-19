import { state } from '@/stores/localstorage'

// 业务守卫的业务逻辑
// 只需要对去往backend后台的请求进行鉴权: /backend/blog
// next： 理解为router.push 函数
export var beforeEachHanler = function (to, from, next) {
  // 使用indexOf来判断当前url 是否已 /backend开头
  if (to.path.indexOf('/backend') === 0) {
    // 需要判断当前用户是否已经登录
    if (!state.value.isLogin) {
      // 如果没有登录，需要重定向到登录页面去
      // 需要获取router对象? 这么不能不用useRoute
      // 直接跳转到LoginPage去登录
      // 记录下用户需要 去往的目标页面
      // /login?to=TagList
      next({ name: "LoginPage", query: { to: to.name } })
      return
    }
  }

  // 直接继续后面的路由处理
  next()
}