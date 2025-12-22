/**
 * 用户状态管理 (Pinia Store)
 * 功能: 用户登录、登出、用户信息管理、token管理
 */
import { defineStore } from 'pinia'
import { ref } from 'vue'
import { loginApi, getUserInfoApi } from '@/api/user'
import { message } from 'ant-design-vue'

export interface UserInfo {
  id: string
  username: string
  realname: string
  avatar: string
  phone: string
  email: string
  status: number
  school: string
}

export const useUserStore = defineStore(
  'user',
  () => {
    // 状态
    const token = ref<string>('')
    const userInfo = ref<UserInfo | null>(null)
    const roles = ref<string[]>([])
    const permissions = ref<string[]>([])

    // 登录
    const login = async (username: string, password: string) => {
      try {
        const res = await loginApi({ username, password })
        token.value = res.token
        userInfo.value = res.userInfo
        message.success('登录成功')
        return res
      } catch (error) {
        console.error('登录失败:', error)
        throw error
      }
    }

    // 获取用户信息
    const getUserInfo = async () => {
      try {
        const res = await getUserInfoApi()
        userInfo.value = res.userInfo
        roles.value = res.roles || []
        permissions.value = res.permissions || []
        return res
      } catch (error) {
        console.error('获取用户信息失败:', error)
        throw error
      }
    }

    // 登出
    const logout = () => {
      token.value = ''
      userInfo.value = null
      roles.value = []
      permissions.value = []
      message.info('已退出登录')
    }

    // 更新用户信息
    const updateUserInfo = (info: Partial<UserInfo>) => {
      if (userInfo.value) {
        userInfo.value = { ...userInfo.value, ...info }
      }
    }

    // 判断是否有权限
    const hasPermission = (permission: string) => {
      return permissions.value.includes(permission)
    }

    // 判断是否有角色
    const hasRole = (role: string) => {
      return roles.value.includes(role)
    }

    return {
      token,
      userInfo,
      roles,
      permissions,
      login,
      logout,
      getUserInfo,
      updateUserInfo,
      hasPermission,
      hasRole,
    }
  },
  {
    // 持久化配置
    persist: {
      enabled: true,
      strategies: [
        {
          key: 'user-store',
          storage: localStorage,
          paths: ['token', 'userInfo', 'roles', 'permissions'],
        },
      ],
    },
  }
)
