import { createRouter, createWebHistory } from 'vue-router'
import HomeView from '../views/HomeView.vue'
import { beforeEach } from "./permission";

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'home',
      component: HomeView
    },
    {
      // 登录页面
      path: "/login",
      name: "Login",
      component: () => import("@/views/login/LoginView.vue")
    },
    {
      // 前台页面
      path: "/frontend",
      name: "FrontendLayout",
      component: () => import("@/views/frontend/LayoutView.vue"),
      children: [
        {
          path: "blog/list",
          name: "FrontendBlogList",
          component: () => import("@/views/frontend/blog/ListView.vue"),
        }
      ]
    },
    {
      // 后台页面
      path: "/backend",
      name: "BackendLayout",
      component: () => import("@/views/backend/LayoutView.vue"),
      // 当访问/backend路径时，应该重定向到BackendBlogList页面。
      redirect: { name: 'BackendBlogList' },
      children: [
        {
          path: "blog/list",
          name: "BackendBlogList",
          component: () => import("@/views/backend/blog/ListView.vue"),
        },
        {
          path: 'blog/edit',
          name: 'BlogEdit',
          //..能直接跳转到页面
          component: () => import('../views/backend/blog/EditView.vue')
        },
      ]
    },
    {
      // 无权限页面
      path: "/errors/403",
      name: "PermissionDeny",
      component: () => import("@/views/errors/PermissionDeny.vue")
    },
    {
      // 前面所有路由都无法被匹配时, 指向404页面
      path: "/:pathMatch(.*)*",
      name: "notFound",
      component: () => import("@/views/errors/NotFound.vue")
    }
  ]
})


// 补充导航守卫
router.beforeEach(beforeEach)

export default router
