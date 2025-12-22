/**
 * 路由守卫
 * 功能: 认证守卫、权限守卫、页面标题设置
 */
import type { Router } from 'vue-router'
import { useUserStore } from '@/stores/user'
import { message } from 'ant-design-vue'

// 白名单 - 不需要登录的页面
const whiteList = ['/login', '/register', '/404']

export function setupRouterGuard(router: Router) {
  // 前置守卫 - 认证和权限检查
  router.beforeEach(async (to, from, next) => {
    // 设置页面标题
    document.title = to.meta.title ? `${to.meta.title} - Teaching Open` : 'Teaching Open'

    const userStore = useUserStore()
    const hasToken = userStore.token

    if (hasToken) {
      // 已登录
      if (to.path === '/login') {
        // 如果已登录,访问登录页则跳转到首页
        next({ path: '/' })
      } else {
        // 检查是否有用户信息
        const hasUserInfo = userStore.userInfo !== null
        if (hasUserInfo) {
          // 有用户信息,检查权限
          const requiresPermission = to.meta.permission as string
          if (requiresPermission) {
            const hasPermission = userStore.hasPermission(requiresPermission)
            if (hasPermission) {
              next()
            } else {
              message.error('没有权限访问该页面')
              next({ path: '/403' })
            }
          } else {
            next()
          }
        } else {
          // 没有用户信息,获取用户信息
          try {
            await userStore.getUserInfo()
            next({ ...to, replace: true })
          } catch (error) {
            // 获取用户信息失败,清除token并跳转登录
            userStore.logout()
            message.error('获取用户信息失败,请重新登录')
            next(`/login?redirect=${to.path}`)
          }
        }
      }
    } else {
      // 未登录
      if (whiteList.includes(to.path)) {
        // 在白名单中,直接放行
        next()
      } else {
        // 不在白名单中,跳转登录页
        next(`/login?redirect=${to.path}`)
      }
    }
  })

  // 后置守卫
  router.afterEach(() => {
    // 可以在这里添加页面访问日志等
  })

  // 错误处理
  router.onError((error) => {
    console.error('路由错误:', error)
    message.error('页面加载失败')
  })
}
