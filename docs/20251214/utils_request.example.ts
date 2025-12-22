/**
 * Axios HTTP 请求封装
 * 功能: 请求拦截(JWT)、响应拦截(错误处理)、统一配置
 */
import axios, { AxiosInstance, AxiosRequestConfig, AxiosResponse, AxiosError } from 'axios'
import { message } from 'ant-design-vue'
import { useUserStore } from '@/stores/user'

// 创建axios实例
const service: AxiosInstance = axios.create({
  baseURL: import.meta.env.VITE_API_BASE_URL || '/api',
  timeout: 30000,
  headers: {
    'Content-Type': 'application/json;charset=UTF-8',
  },
})

// 请求拦截器 - 添加JWT token
service.interceptors.request.use(
  (config: AxiosRequestConfig) => {
    const userStore = useUserStore()
    const token = userStore.token

    // 如果有token,添加到请求头
    if (token && config.headers) {
      config.headers['Authorization'] = `Bearer ${token}`
    }

    return config
  },
  (error: AxiosError) => {
    console.error('请求错误:', error)
    return Promise.reject(error)
  }
)

// 响应拦截器 - 统一错误处理
service.interceptors.response.use(
  (response: AxiosResponse) => {
    const res = response.data

    // 根据GoFrame后端的响应格式处理
    // 假设后端返回格式: { code: 0, message: 'success', data: {...} }
    if (res.code !== 0) {
      // 业务错误
      message.error(res.message || '请求失败')

      // 401未授权 - 跳转登录
      if (res.code === 401) {
        const userStore = useUserStore()
        userStore.logout()
        window.location.href = '/login'
      }

      return Promise.reject(new Error(res.message || '请求失败'))
    }

    return res.data
  },
  (error: AxiosError) => {
    console.error('响应错误:', error)

    // HTTP错误处理
    if (error.response) {
      const status = error.response.status
      switch (status) {
        case 401:
          message.error('未授权,请登录')
          const userStore = useUserStore()
          userStore.logout()
          window.location.href = '/login'
          break
        case 403:
          message.error('没有权限访问')
          break
        case 404:
          message.error('请求的资源不存在')
          break
        case 500:
          message.error('服务器错误')
          break
        default:
          message.error('网络请求失败')
      }
    } else if (error.request) {
      message.error('网络连接失败,请检查网络')
    } else {
      message.error('请求配置错误')
    }

    return Promise.reject(error)
  }
)

// 导出封装的请求方法
export default {
  get<T = any>(url: string, params?: object): Promise<T> {
    return service.get(url, { params })
  },

  post<T = any>(url: string, data?: object): Promise<T> {
    return service.post(url, data)
  },

  put<T = any>(url: string, data?: object): Promise<T> {
    return service.put(url, data)
  },

  delete<T = any>(url: string, params?: object): Promise<T> {
    return service.delete(url, { params })
  },
}
