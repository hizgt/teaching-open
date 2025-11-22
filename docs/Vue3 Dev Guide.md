# Vue3 开发指南 - Teaching Open 前端升级方案

## 一、升级概述

### 1.1 升级目标
将现有的 Vue 2.6 + Ant Design Vue 1.x 前端项目升级到 Vue 3.x + Ant Design Vue 4.x，同时优化代码结构，提升性能和开发体验。

### 1.2 技术栈对比

| 组件 | Vue2 原系统 | Vue3 目标系统 |
|------|------------|--------------|
| 核心框架 | Vue 2.6.10 | Vue 3.4+ |
| UI组件库 | Ant Design Vue 1.6.3 | Ant Design Vue 4.x |
| 状态管理 | Vuex 3.1.0 | Pinia 2.x |
| 路由 | Vue Router 3.0.1 | Vue Router 4.x |
| HTTP库 | Axios 0.18.0 | Axios 1.6+ |
| 构建工具 | Vue CLI 3.3.0 / Webpack | Vite 5.x |
| TypeScript | 无 | TypeScript 5.x |
| 代码风格 | Options API | Composition API + Script Setup |
| CSS方案 | Less | Less/Sass + CSS Modules |
| 包管理器 | Yarn/NPM | PNPM (推荐) |

### 1.3 升级收益

- **性能提升**: 更快的渲染速度、更小的打包体积
- **开发体验**: 更好的TypeScript支持、更快的热更新
- **代码质量**: Composition API带来更好的代码组织和复用
- **生态系统**: 更活跃的社区支持、更多的第三方库

## 二、项目结构设计

### 2.1 推荐目录结构

```
teaching-open-web/
├── .vscode/                      # VSCode配置
│   ├── extensions.json
│   └── settings.json
├── public/                       # 静态资源
│   ├── favicon.ico
│   └── index.html
├── src/
│   ├── api/                      # API接口定义
│   │   ├── modules/
│   │   │   ├── system/          # 系统模块API
│   │   │   │   ├── user.ts
│   │   │   │   ├── role.ts
│   │   │   │   └── permission.ts
│   │   │   └── teaching/        # 教学模块API
│   │   │       ├── course.ts
│   │   │       ├── work.ts
│   │   │       └── student.ts
│   │   ├── request.ts           # Axios封装
│   │   └── types.ts             # API类型定义
│   ├── assets/                  # 资源文件
│   │   ├── images/
│   │   ├── styles/
│   │   │   ├── global.less
│   │   │   ├── variables.less
│   │   │   └── mixins.less
│   │   └── fonts/
│   ├── components/              # 通用组件
│   │   ├── common/              # 基础组件
│   │   │   ├── AppHeader.vue
│   │   │   ├── AppFooter.vue
│   │   │   └── AppMenu.vue
│   │   ├── business/            # 业务组件
│   │   │   ├── UserSelect.vue
│   │   │   ├── DeptTree.vue
│   │   │   └── UploadFile.vue
│   │   └── layout/              # 布局组件
│   │       ├── BasicLayout.vue
│   │       ├── UserLayout.vue
│   │       └── HomeLayout.vue
│   ├── composables/             # 组合式函数
│   │   ├── usePermission.ts
│   │   ├── useTable.ts
│   │   ├── useForm.ts
│   │   └── useAuth.ts
│   ├── directives/              # 自定义指令
│   │   ├── permission.ts
│   │   └── loading.ts
│   ├── hooks/                   # 自定义Hooks
│   │   ├── useRequest.ts
│   │   ├── useModal.ts
│   │   └── usePagination.ts
│   ├── router/                  # 路由配置
│   │   ├── index.ts
│   │   ├── routes.ts
│   │   └── guards.ts
│   ├── stores/                  # Pinia状态管理
│   │   ├── modules/
│   │   │   ├── user.ts
│   │   │   ├── permission.ts
│   │   │   └── app.ts
│   │   └── index.ts
│   ├── utils/                   # 工具函数
│   │   ├── auth.ts              # 认证工具
│   │   ├── storage.ts           # 存储工具
│   │   ├── date.ts              # 日期工具
│   │   └── validate.ts          # 验证工具
│   ├── views/                   # 页面组件
│   │   ├── system/              # 系统管理
│   │   │   ├── user/
│   │   │   │   ├── index.vue
│   │   │   │   ├── UserModal.vue
│   │   │   │   └── types.ts
│   │   │   ├── role/
│   │   │   └── permission/
│   │   ├── teaching/            # 教学管理
│   │   │   ├── course/
│   │   │   ├── work/
│   │   │   └── student/
│   │   ├── home/                # 首页
│   │   └── login/               # 登录
│   ├── types/                   # TypeScript类型定义
│   │   ├── global.d.ts
│   │   ├── api.d.ts
│   │   └── components.d.ts
│   ├── App.vue
│   ├── main.ts
│   └── env.d.ts
├── .env                         # 环境变量
├── .env.development             # 开发环境变量
├── .env.production              # 生产环境变量
├── .eslintrc.cjs                # ESLint配置
├── .prettierrc                  # Prettier配置
├── index.html
├── package.json
├── tsconfig.json                # TypeScript配置
├── tsconfig.node.json
├── vite.config.ts               # Vite配置
└── README.md
```

## 三、核心功能升级方案

### 3.1 Composition API迁移

#### Vue2 Options API (旧)
```javascript
export default {
  name: 'UserList',
  data() {
    return {
      loading: false,
      dataSource: [],
      pagination: {
        current: 1,
        pageSize: 10,
        total: 0
      },
      queryParam: {
        username: ''
      }
    }
  },
  created() {
    this.loadData()
  },
  methods: {
    async loadData() {
      this.loading = true
      try {
        const res = await getUserList({
          page: this.pagination.current,
          pageSize: this.pagination.pageSize,
          ...this.queryParam
        })
        this.dataSource = res.records
        this.pagination.total = res.total
      } finally {
        this.loading = false
      }
    },
    handleSearch() {
      this.pagination.current = 1
      this.loadData()
    }
  }
}
```

#### Vue3 Composition API + Script Setup (新)
```vue
<script setup lang="ts">
import { ref, reactive, onMounted } from 'vue'
import { getUserList } from '@/api/modules/system/user'
import type { User, UserQuery } from '@/types/system/user'

// 响应式数据
const loading = ref(false)
const dataSource = ref<User[]>([])
const pagination = reactive({
  current: 1,
  pageSize: 10,
  total: 0
})
const queryParam = reactive<UserQuery>({
  username: ''
})

// 加载数据
const loadData = async () => {
  loading.value = true
  try {
    const res = await getUserList({
      page: pagination.current,
      pageSize: pagination.pageSize,
      ...queryParam
    })
    dataSource.value = res.records
    pagination.total = res.total
  } finally {
    loading.value = false
  }
}

// 搜索
const handleSearch = () => {
  pagination.current = 1
  loadData()
}

// 生命周期
onMounted(() => {
  loadData()
})
</script>

<template>
  <div class="user-list">
    <a-card>
      <a-form layout="inline">
        <a-form-item label="用户名">
          <a-input v-model:value="queryParam.username" placeholder="请输入用户名" />
        </a-form-item>
        <a-form-item>
          <a-button type="primary" @click="handleSearch">查询</a-button>
        </a-form-item>
      </a-form>

      <a-table
        :loading="loading"
        :dataSource="dataSource"
        :pagination="pagination"
        @change="loadData"
      />
    </a-card>
  </div>
</template>
```

### 3.2 状态管理 (Vuex → Pinia)

#### Vuex (旧)
```javascript
// store/modules/user.js
export default {
  namespaced: true,
  state: {
    token: '',
    userInfo: null
  },
  mutations: {
    SET_TOKEN(state, token) {
      state.token = token
    },
    SET_USER_INFO(state, info) {
      state.userInfo = info
    }
  },
  actions: {
    async login({ commit }, userInfo) {
      const res = await login(userInfo)
      commit('SET_TOKEN', res.token)
      return res
    }
  }
}
```

#### Pinia (新)
```typescript
// stores/modules/user.ts
import { defineStore } from 'pinia'
import { ref } from 'vue'
import { login, getUserInfo } from '@/api/modules/system/user'
import type { LoginParams, UserInfo } from '@/types/system/user'
import { setToken, getToken, removeToken } from '@/utils/auth'

export const useUserStore = defineStore('user', () => {
  // State
  const token = ref(getToken() || '')
  const userInfo = ref<UserInfo | null>(null)
  const permissions = ref<string[]>([])
  
  // Actions
  const setUserToken = (newToken: string) => {
    token.value = newToken
    setToken(newToken)
  }
  
  const setUserInfo = (info: UserInfo) => {
    userInfo.value = info
  }
  
  const userLogin = async (params: LoginParams) => {
    const res = await login(params)
    setUserToken(res.token)
    return res
  }
  
  const fetchUserInfo = async () => {
    const res = await getUserInfo()
    setUserInfo(res.userInfo)
    permissions.value = res.permissions
    return res
  }
  
  const logout = () => {
    token.value = ''
    userInfo.value = null
    permissions.value = []
    removeToken()
  }
  
  // Getters
  const isLogin = computed(() => !!token.value)
  const username = computed(() => userInfo.value?.username || '')
  
  return {
    // State
    token,
    userInfo,
    permissions,
    // Actions
    setUserToken,
    setUserInfo,
    userLogin,
    fetchUserInfo,
    logout,
    // Getters
    isLogin,
    username
  }
})
```

使用示例:
```vue
<script setup lang="ts">
import { useUserStore } from '@/stores/modules/user'

const userStore = useUserStore()

// 访问状态
console.log(userStore.username)

// 调用方法
const handleLogin = async () => {
  await userStore.userLogin({ username: 'admin', password: '123456' })
}
</script>
```

### 3.3 路由配置升级

#### Vue Router 4配置
```typescript
// router/index.ts
import { createRouter, createWebHistory } from 'vue-router'
import type { RouteRecordRaw } from 'vue-router'
import { setupRouterGuard } from './guards'

const routes: RouteRecordRaw[] = [
  {
    path: '/login',
    name: 'Login',
    component: () => import('@/views/login/index.vue'),
    meta: {
      title: '登录',
      requiresAuth: false
    }
  },
  {
    path: '/',
    name: 'Layout',
    component: () => import('@/components/layout/BasicLayout.vue'),
    redirect: '/dashboard',
    meta: {
      requiresAuth: true
    },
    children: [
      {
        path: 'dashboard',
        name: 'Dashboard',
        component: () => import('@/views/dashboard/index.vue'),
        meta: {
          title: '工作台',
          icon: 'DashboardOutlined'
        }
      },
      {
        path: 'system',
        name: 'System',
        meta: {
          title: '系统管理',
          icon: 'SettingOutlined'
        },
        children: [
          {
            path: 'user',
            name: 'SystemUser',
            component: () => import('@/views/system/user/index.vue'),
            meta: {
              title: '用户管理',
              permission: 'system:user:list'
            }
          },
          {
            path: 'role',
            name: 'SystemRole',
            component: () => import('@/views/system/role/index.vue'),
            meta: {
              title: '角色管理',
              permission: 'system:role:list'
            }
          }
        ]
      }
    ]
  },
  {
    path: '/:pathMatch(.*)*',
    name: 'NotFound',
    component: () => import('@/views/error/404.vue')
  }
]

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes
})

// 设置路由守卫
setupRouterGuard(router)

export default router
```

#### 路由守卫
```typescript
// router/guards.ts
import type { Router } from 'vue-router'
import { useUserStore } from '@/stores/modules/user'
import { message } from 'ant-design-vue'

export function setupRouterGuard(router: Router) {
  // 前置守卫
  router.beforeEach(async (to, from, next) => {
    // 设置页面标题
    document.title = to.meta.title ? `${to.meta.title} - Teaching Open` : 'Teaching Open'
    
    const userStore = useUserStore()
    
    // 白名单路由
    const whiteList = ['/login', '/register']
    
    if (userStore.token) {
      if (to.path === '/login') {
        next({ path: '/' })
      } else {
        // 获取用户信息
        if (!userStore.userInfo) {
          try {
            await userStore.fetchUserInfo()
            next({ ...to, replace: true })
          } catch (error) {
            userStore.logout()
            message.error('获取用户信息失败，请重新登录')
            next({ path: '/login', query: { redirect: to.fullPath } })
          }
        } else {
          // 检查权限
          if (to.meta.permission && !hasPermission(to.meta.permission as string)) {
            message.error('没有权限访问该页面')
            next({ path: '/403' })
          } else {
            next()
          }
        }
      }
    } else {
      if (whiteList.includes(to.path)) {
        next()
      } else {
        next({ path: '/login', query: { redirect: to.fullPath } })
      }
    }
  })
  
  // 后置守卫
  router.afterEach(() => {
    // 进度条结束等
  })
}

function hasPermission(permission: string): boolean {
  const userStore = useUserStore()
  return userStore.permissions.includes(permission) || userStore.permissions.includes('*')
}
```

### 3.4 Axios封装与TypeScript类型

```typescript
// api/request.ts
import axios, { AxiosError, AxiosRequestConfig, AxiosResponse } from 'axios'
import { message } from 'ant-design-vue'
import { useUserStore } from '@/stores/modules/user'

// 响应数据格式
export interface ResponseData<T = any> {
  code: number
  message: string
  result: T
  success: boolean
  timestamp: number
}

// 创建axios实例
const service = axios.create({
  baseURL: import.meta.env.VITE_API_BASE_URL,
  timeout: 30000,
  headers: {
    'Content-Type': 'application/json;charset=UTF-8'
  }
})

// 请求拦截器
service.interceptors.request.use(
  (config) => {
    const userStore = useUserStore()
    if (userStore.token) {
      config.headers['X-Access-Token'] = userStore.token
    }
    return config
  },
  (error: AxiosError) => {
    console.error('Request error:', error)
    return Promise.reject(error)
  }
)

// 响应拦截器
service.interceptors.response.use(
  (response: AxiosResponse<ResponseData>) => {
    const res = response.data
    
    if (res.code !== 200) {
      message.error(res.message || '请求失败')
      
      // token过期
      if (res.code === 401) {
        const userStore = useUserStore()
        userStore.logout()
        window.location.href = '/login'
      }
      
      return Promise.reject(new Error(res.message || '请求失败'))
    }
    
    return res.result
  },
  (error: AxiosError) => {
    console.error('Response error:', error)
    
    if (error.response) {
      const status = error.response.status
      const messages: Record<number, string> = {
        400: '请求参数错误',
        401: '未授权，请重新登录',
        403: '拒绝访问',
        404: '请求资源不存在',
        500: '服务器错误',
        502: '网关错误',
        503: '服务不可用',
        504: '网关超时'
      }
      message.error(messages[status] || `请求失败(${status})`)
    } else {
      message.error('网络连接异常，请检查网络设置')
    }
    
    return Promise.reject(error)
  }
)

// 封装请求方法
export const request = {
  get<T = any>(url: string, params?: any, config?: AxiosRequestConfig): Promise<T> {
    return service.get(url, { params, ...config })
  },
  
  post<T = any>(url: string, data?: any, config?: AxiosRequestConfig): Promise<T> {
    return service.post(url, data, config)
  },
  
  put<T = any>(url: string, data?: any, config?: AxiosRequestConfig): Promise<T> {
    return service.put(url, data, config)
  },
  
  delete<T = any>(url: string, params?: any, config?: AxiosRequestConfig): Promise<T> {
    return service.delete(url, { params, ...config })
  }
}

export default service
```

#### API定义示例
```typescript
// api/modules/system/user.ts
import { request } from '@/api/request'

export interface LoginParams {
  username: string
  password: string
}

export interface LoginResult {
  token: string
  userInfo: UserInfo
}

export interface UserInfo {
  id: string
  username: string
  realname: string
  avatar: string
  email: string
}

export interface UserQuery {
  username?: string
  realname?: string
  page: number
  pageSize: number
}

export interface UserListResult {
  records: UserInfo[]
  total: number
  pageNo: number
  pageSize: number
}

// 登录
export function login(data: LoginParams) {
  return request.post<LoginResult>('/sys/login', data)
}

// 获取用户信息
export function getUserInfo() {
  return request.get<LoginResult>('/sys/user/info')
}

// 获取用户列表
export function getUserList(params: UserQuery) {
  return request.get<UserListResult>('/sys/user/list', params)
}

// 添加用户
export function addUser(data: Partial<UserInfo>) {
  return request.post('/sys/user/add', data)
}

// 编辑用户
export function editUser(data: Partial<UserInfo>) {
  return request.put('/sys/user/edit', data)
}

// 删除用户
export function deleteUser(id: string) {
  return request.delete(`/sys/user/delete`, { id })
}
```

### 3.5 自定义Hooks封装

#### 表格Hook
```typescript
// hooks/useTable.ts
import { ref, reactive } from 'vue'
import type { TableProps } from 'ant-design-vue'

interface UseTableOptions<T = any> {
  api: (params: any) => Promise<{
    records: T[]
    total: number
  }>
  immediate?: boolean
  queryParams?: Record<string, any>
}

export function useTable<T = any>(options: UseTableOptions<T>) {
  const { api, immediate = true, queryParams = {} } = options
  
  const loading = ref(false)
  const dataSource = ref<T[]>([])
  const pagination = reactive({
    current: 1,
    pageSize: 10,
    total: 0,
    showSizeChanger: true,
    showQuickJumper: true,
    showTotal: (total: number) => `共 ${total} 条`
  })
  
  const query = reactive(queryParams)
  
  const loadData = async (resetPage = false) => {
    if (resetPage) {
      pagination.current = 1
    }
    
    loading.value = true
    try {
      const res = await api({
        page: pagination.current,
        pageSize: pagination.pageSize,
        ...query
      })
      dataSource.value = res.records
      pagination.total = res.total
    } finally {
      loading.value = false
    }
  }
  
  const handleTableChange: TableProps['onChange'] = (pag) => {
    pagination.current = pag.current!
    pagination.pageSize = pag.pageSize!
    loadData()
  }
  
  const refresh = () => loadData()
  const search = () => loadData(true)
  const reset = () => {
    Object.keys(query).forEach(key => {
      query[key] = undefined
    })
    loadData(true)
  }
  
  if (immediate) {
    loadData()
  }
  
  return {
    loading,
    dataSource,
    pagination,
    query,
    loadData,
    handleTableChange,
    refresh,
    search,
    reset
  }
}
```

使用示例:
```vue
<script setup lang="ts">
import { useTable } from '@/hooks/useTable'
import { getUserList } from '@/api/modules/system/user'
import type { UserInfo } from '@/types/system/user'

const {
  loading,
  dataSource,
  pagination,
  query,
  handleTableChange,
  search,
  reset
} = useTable<UserInfo>({
  api: getUserList,
  queryParams: {
    username: '',
    realname: ''
  }
})

const columns = [
  { title: '用户名', dataIndex: 'username' },
  { title: '姓名', dataIndex: 'realname' },
  { title: '邮箱', dataIndex: 'email' }
]
</script>

<template>
  <div>
    <a-form layout="inline">
      <a-form-item label="用户名">
        <a-input v-model:value="query.username" />
      </a-form-item>
      <a-form-item>
        <a-button type="primary" @click="search">查询</a-button>
        <a-button @click="reset">重置</a-button>
      </a-form-item>
    </a-form>
    
    <a-table
      :loading="loading"
      :dataSource="dataSource"
      :columns="columns"
      :pagination="pagination"
      @change="handleTableChange"
    />
  </div>
</template>
```

### 3.6 权限指令

```typescript
// directives/permission.ts
import type { Directive } from 'vue'
import { useUserStore } from '@/stores/modules/user'

export const permission: Directive = {
  mounted(el, binding) {
    const { value } = binding
    const userStore = useUserStore()
    
    if (value) {
      const permissions = userStore.permissions
      const hasPermission = permissions.includes(value) || permissions.includes('*')
      
      if (!hasPermission) {
        el.style.display = 'none'
        // 或者直接移除元素
        // el.parentNode?.removeChild(el)
      }
    }
  }
}

// 注册全局指令
// main.ts
import { permission } from '@/directives/permission'
app.directive('permission', permission)
```

使用:
```vue
<template>
  <a-button v-permission="'user:add'" type="primary">添加用户</a-button>
  <a-button v-permission="'user:delete'" danger>删除</a-button>
</template>
```

## 四、Vite配置

```typescript
// vite.config.ts
import { defineConfig, loadEnv } from 'vite'
import vue from '@vitejs/plugin-vue'
import { resolve } from 'path'
import Components from 'unplugin-vue-components/vite'
import { AntDesignVueResolver } from 'unplugin-vue-components/resolvers'

export default defineConfig(({ mode }) => {
  const env = loadEnv(mode, process.cwd())
  
  return {
    plugins: [
      vue(),
      // 组件自动导入
      Components({
        resolvers: [
          AntDesignVueResolver({
            importStyle: false
          })
        ]
      })
    ],
    resolve: {
      alias: {
        '@': resolve(__dirname, 'src'),
        '~': resolve(__dirname, 'src')
      }
    },
    css: {
      preprocessorOptions: {
        less: {
          javascriptEnabled: true,
          modifyVars: {
            'primary-color': '#1890ff',
            'link-color': '#1890ff',
            'border-radius-base': '4px'
          }
        }
      }
    },
    server: {
      host: '0.0.0.0',
      port: 3000,
      open: true,
      proxy: {
        '/api': {
          target: env.VITE_API_BASE_URL,
          changeOrigin: true,
          rewrite: (path) => path.replace(/^\/api/, '')
        }
      }
    },
    build: {
      outDir: 'dist',
      sourcemap: false,
      chunkSizeWarningLimit: 1500,
      rollupOptions: {
        output: {
          manualChunks: {
            'vue-vendor': ['vue', 'vue-router', 'pinia'],
            'antdv-vendor': ['ant-design-vue', '@ant-design/icons-vue']
          }
        }
      }
    }
  }
})
```

## 五、环境配置

```bash
# .env.development
VITE_APP_TITLE=Teaching Open
VITE_API_BASE_URL=http://localhost:8080
VITE_UPLOAD_URL=http://localhost:8080/upload

# .env.production
VITE_APP_TITLE=Teaching Open
VITE_API_BASE_URL=https://api.example.com
VITE_UPLOAD_URL=https://api.example.com/upload
```

## 六、升级步骤

### 6.1 阶段一：环境准备 (1周)
1. 创建新的Vue3项目
2. 配置TypeScript、ESLint、Prettier
3. 安装依赖包
4. 配置Vite构建工具
5. 配置路径别名和环境变量

### 6.2 阶段二：基础设施迁移 (2周)
1. API封装和类型定义
2. Pinia状态管理
3. 路由配置
4. 权限控制
5. 通用组件封装
6. 工具函数迁移

### 6.3 阶段三：页面组件迁移 (4-5周)
1. 登录页面
2. 系统管理模块 (用户、角色、权限、部门)
3. 教学管理模块 (课程、作品、学生)
4. 首页和工作台
5. 其他业务页面

### 6.4 阶段四：测试优化 (1-2周)
1. 功能测试
2. 兼容性测试
3. 性能优化
4. 打包优化

## 七、最佳实践

### 7.1 组件设计原则
- 单一职责：每个组件只负责一个功能
- 可复用性：提取通用逻辑到Hooks
- Props验证：使用TypeScript进行类型检查
- 事件命名：使用kebab-case命名自定义事件

### 7.2 代码规范
- 使用ESLint + Prettier统一代码风格
- 组件命名使用PascalCase
- 文件命名使用kebab-case
- 常量使用UPPER_CASE

### 7.3 性能优化
- 使用`v-memo`优化列表渲染
- 路由懒加载
- 组件异步加载
- 图片懒加载
- 防抖节流

## 八、部署

```bash
# 安装依赖
pnpm install

# 开发
pnpm dev

# 构建
pnpm build

# 预览
pnpm preview
```

### Docker部署
```dockerfile
# Dockerfile
FROM node:18-alpine as build
WORKDIR /app
COPY package.json pnpm-lock.yaml ./
RUN npm install -g pnpm && pnpm install
COPY . .
RUN pnpm build

FROM nginx:alpine
COPY --from=build /app/dist /usr/share/nginx/html
COPY nginx.conf /etc/nginx/conf.d/default.conf
EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]
```

## 九、参考资源

- [Vue 3 官方文档](https://cn.vuejs.org/)
- [Ant Design Vue 4.x](https://antdv.com/)
- [Pinia 文档](https://pinia.vuejs.org/zh/)
- [Vite 文档](https://cn.vitejs.dev/)
- [TypeScript 文档](https://www.typescriptlang.org/)
