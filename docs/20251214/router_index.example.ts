/**
 * Vue Router 路由配置
 * 功能: 基础路由、动态路由、嵌套路由、路由守卫
 */
import { createRouter, createWebHistory, RouteRecordRaw } from 'vue-router'
import BasicLayout from '@/layouts/BasicLayout.vue'
import UserLayout from '@/layouts/UserLayout.vue'

// 基础路由(不需要权限)
const basicRoutes: RouteRecordRaw[] = [
  {
    path: '/login',
    name: 'Login',
    component: () => import('@/views/login/index.vue'),
    meta: {
      title: '登录',
      requiresAuth: false,
    },
  },
  {
    path: '/',
    component: BasicLayout,
    redirect: '/home',
    children: [
      {
        path: 'home',
        name: 'Home',
        component: () => import('@/views/home/index.vue'),
        meta: {
          title: '首页',
          icon: 'HomeOutlined',
          requiresAuth: true,
        },
      },
    ],
  },
]

// 系统管理路由(需要权限)
const systemRoutes: RouteRecordRaw[] = [
  {
    path: '/system',
    component: BasicLayout,
    redirect: '/system/user',
    meta: {
      title: '系统管理',
      icon: 'SettingOutlined',
      requiresAuth: true,
    },
    children: [
      {
        path: 'user',
        name: 'SystemUser',
        component: () => import('@/views/system/user/index.vue'),
        meta: {
          title: '用户管理',
          icon: 'UserOutlined',
          permission: 'system:user:list',
        },
      },
      {
        path: 'role',
        name: 'SystemRole',
        component: () => import('@/views/system/role/index.vue'),
        meta: {
          title: '角色管理',
          icon: 'TeamOutlined',
          permission: 'system:role:list',
        },
      },
      {
        path: 'permission',
        name: 'SystemPermission',
        component: () => import('@/views/system/permission/index.vue'),
        meta: {
          title: '权限管理',
          icon: 'SafetyOutlined',
          permission: 'system:permission:list',
        },
      },
      {
        path: 'depart',
        name: 'SystemDepart',
        component: () => import('@/views/system/depart/index.vue'),
        meta: {
          title: '部门管理',
          icon: 'ApartmentOutlined',
          permission: 'system:depart:list',
        },
      },
      {
        path: 'dict',
        name: 'SystemDict',
        component: () => import('@/views/system/dict/index.vue'),
        meta: {
          title: '字典管理',
          icon: 'BookOutlined',
          permission: 'system:dict:list',
        },
      },
      {
        path: 'log',
        name: 'SystemLog',
        component: () => import('@/views/system/log/index.vue'),
        meta: {
          title: '日志管理',
          icon: 'FileTextOutlined',
          permission: 'system:log:list',
        },
      },
    ],
  },
]

// 教学管理路由(需要权限)
const teachingRoutes: RouteRecordRaw[] = [
  {
    path: '/teaching',
    component: BasicLayout,
    redirect: '/teaching/course',
    meta: {
      title: '教学管理',
      icon: 'ReadOutlined',
      requiresAuth: true,
    },
    children: [
      {
        path: 'course',
        name: 'TeachingCourse',
        component: () => import('@/views/teaching/course/index.vue'),
        meta: {
          title: '课程管理',
          icon: 'ProjectOutlined',
          permission: 'teaching:course:list',
        },
      },
      {
        path: 'work',
        name: 'TeachingWork',
        component: () => import('@/views/teaching/work/index.vue'),
        meta: {
          title: '作品管理',
          icon: 'CodeOutlined',
          permission: 'teaching:work:list',
        },
      },
      {
        path: 'news',
        name: 'TeachingNews',
        component: () => import('@/views/teaching/news/index.vue'),
        meta: {
          title: '新闻公告',
          icon: 'NotificationOutlined',
          permission: 'teaching:news:list',
        },
      },
      {
        path: 'homework',
        name: 'TeachingHomework',
        component: () => import('@/views/teaching/homework/index.vue'),
        meta: {
          title: '附加作业',
          icon: 'EditOutlined',
          permission: 'teaching:homework:list',
        },
      },
      {
        path: 'material',
        name: 'TeachingMaterial',
        component: () => import('@/views/teaching/material/index.vue'),
        meta: {
          title: 'Scratch素材',
          icon: 'PictureOutlined',
          permission: 'teaching:material:list',
        },
      },
    ],
  },
]

// 404路由
const errorRoutes: RouteRecordRaw[] = [
  {
    path: '/:pathMatch(.*)*',
    name: 'NotFound',
    component: () => import('@/views/error/404.vue'),
    meta: {
      title: '页面不存在',
      requiresAuth: false,
    },
  },
]

// 合并所有路由
const routes: RouteRecordRaw[] = [...basicRoutes, ...systemRoutes, ...teachingRoutes, ...errorRoutes]

// 创建路由实例
const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes,
  scrollBehavior() {
    return { top: 0 }
  },
})

export default router
