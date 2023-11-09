import { createRouter, createWebHistory } from 'vue-router'
import HomeView from '../views/HomeView.vue'
//import { beforeEachHandler } from "./permission";

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
      name: "LoginPage",
      component: () => import("@/views/login/LoginPage.vue")
    },
    {
      // 前台页面
      path: "/frontend",
      name: "FrontendLayout",
      component: () => import("@/views/frontend/FrontendLayout.vue"),
      children: [
        {
          path: "blog/list",
          name: "FrontendBlogList",
          component: () => import("@/views/frontend/BlogList.vue"),
        }
      ]
    },
    {
      // 后台页面
      path: "/backend",
      name: "BackendLayout",
      component: () => import("@/views/backend/BackendLayout.vue"),
      children: [
        {
          path: "blog/list",
          name: "BackendBlogList",
          component: () => import("@/views/backend/blog/BlogList.vue"),
        }
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
//router.beforeEach(beforeEachHandler)

export default router
