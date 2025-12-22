# Phase 13开发计划: 前端Vue3重构

**创建日期**: 2025-12-14  
**预计工期**: 3-4周  
**当前分支**: devgo  
**依赖**: Phase 1-12 后端API已完成

---

## 一、项目初始化 (Week 1, Day 1-2)

### 1.1 创建Vue3项目

```bash
cd /root/teaching
npm create vite@latest web-vue3 -- --template vue-ts
cd web-vue3
npm install
```

### 1.2 安装核心依赖

```bash
# 路由和状态管理
npm install vue-router@4 pinia

# UI框架 - Ant Design Vue 4.x
npm install ant-design-vue

# HTTP请求
npm install axios

# 工具库
npm install dayjs
npm install lodash-es
npm install @vueuse/core

# 开发依赖
npm install -D @types/lodash-es
npm install -D @types/node
npm install -D unplugin-vue-components
npm install -D unplugin-auto-import
```

### 1.3 目录结构创建

```
web-vue3/
├── src/
│   ├── api/                 # API接口定义
│   │   ├── sys/            # 系统管理API
│   │   │   ├── user.ts
│   │   │   ├── role.ts
│   │   │   ├── permission.ts
│   │   │   ├── depart.ts
│   │   │   ├── dict.ts
│   │   │   ├── log.ts
│   │   │   └── file.ts
│   │   └── teaching/       # 教学管理API
│   │       ├── course.ts
│   │       ├── courseUnit.ts
│   │       ├── work.ts
│   │       ├── news.ts
│   │       ├── additionalWork.ts
│   │       ├── scratchAsset.ts
│   │       └── departDayLog.ts
│   ├── assets/             # 静态资源
│   ├── components/         # 通用组件
│   │   ├── BasicLayout/
│   │   ├── UserLayout/
│   │   ├── UploadFile/
│   │   ├── ImagePreview/
│   │   ├── DeptTree/
│   │   └── DictSelect/
│   ├── composables/        # 组合式函数
│   │   ├── useRequest.ts
│   │   ├── useTable.ts
│   │   ├── useForm.ts
│   │   └── useDict.ts
│   ├── router/             # 路由配置
│   │   ├── index.ts
│   │   ├── routes.ts
│   │   └── guards.ts
│   ├── stores/             # Pinia状态管理
│   │   ├── user.ts
│   │   ├── permission.ts
│   │   ├── app.ts
│   │   └── dict.ts
│   ├── styles/             # 全局样式
│   │   ├── index.less
│   │   ├── variables.less
│   │   └── reset.less
│   ├── types/              # TypeScript类型
│   │   ├── api.d.ts
│   │   ├── user.d.ts
│   │   └── common.d.ts
│   ├── utils/              # 工具函数
│   │   ├── request.ts      # Axios封装
│   │   ├── auth.ts         # 认证工具
│   │   ├── storage.ts      # 本地存储
│   │   └── helpers.ts
│   ├── views/              # 页面组件
│   │   ├── login/
│   │   ├── sys/            # 系统管理页面
│   │   │   ├── user/
│   │   │   ├── role/
│   │   │   ├── permission/
│   │   │   ├── depart/
│   │   │   ├── dict/
│   │   │   └── log/
│   │   └── teaching/       # 教学管理页面
│   │       ├── course/
│   │       ├── work/
│   │       ├── news/
│   │       ├── additionalWork/
│   │       ├── scratchAsset/
│   │       └── stat/
│   ├── App.vue
│   └── main.ts
├── public/
├── index.html
├── vite.config.ts
├── tsconfig.json
├── .eslintrc.js
└── package.json
```

---

## 二、基础设施搭建 (Week 1, Day 3-5)

### 2.1 Vite配置 (vite.config.ts)

```typescript
import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import { resolve } from 'path'
import Components from 'unplugin-vue-components/vite'
import { AntDesignVueResolver } from 'unplugin-vue-components/resolvers'
import AutoImport from 'unplugin-auto-import/vite'

export default defineConfig({
  plugins: [
    vue(),
    Components({
      resolvers: [AntDesignVueResolver({ importStyle: false })],
    }),
    AutoImport({
      imports: ['vue', 'vue-router', 'pinia'],
      dts: 'src/types/auto-imports.d.ts',
    }),
  ],
  resolve: {
    alias: {
      '@': resolve(__dirname, 'src'),
      '~': resolve(__dirname, 'src'),
    },
  },
  server: {
    port: 3000,
    proxy: {
      '/api': {
        target: 'http://localhost:8199',
        changeOrigin: true,
      },
    },
  },
})
```

### 2.2 Axios封装 (utils/request.ts)

```typescript
import axios, { AxiosInstance, AxiosRequestConfig, AxiosResponse } from 'axios'
import { message } from 'ant-design-vue'
import { useUserStore } from '@/stores/user'

const service: AxiosInstance = axios.create({
  baseURL: '/api/v1',
  timeout: 15000,
})

// 请求拦截器
service.interceptors.request.use(
  (config) => {
    const userStore = useUserStore()
    if (userStore.token) {
      config.headers['X-Token'] = `Bearer ${userStore.token}`
    }
    return config
  },
  (error) => {
    return Promise.reject(error)
  }
)

// 响应拦截器
service.interceptors.response.use(
  (response: AxiosResponse) => {
    const res = response.data
    
    if (res.code !== 0) {
      message.error(res.message || '请求失败')
      
      // Token过期处理
      if (res.code === 2002 || res.code === 2003) {
        const userStore = useUserStore()
        userStore.logout()
        window.location.href = '/login'
      }
      
      return Promise.reject(new Error(res.message || 'Error'))
    }
    
    return res.result
  },
  (error) => {
    message.error(error.message || '网络错误')
    return Promise.reject(error)
  }
)

export default service
```

### 2.3 Pinia状态管理

#### 用户Store (stores/user.ts)

```typescript
import { defineStore } from 'pinia'
import { login, getUserInfo } from '@/api/sys/user'
import type { UserInfo, LoginParams } from '@/types/user'

export const useUserStore = defineStore('user', {
  state: () => ({
    token: localStorage.getItem('token') || '',
    userInfo: null as UserInfo | null,
  }),
  
  getters: {
    isLoggedIn: (state) => !!state.token,
  },
  
  actions: {
    async login(params: LoginParams) {
      const res = await login(params)
      this.token = res.token
      this.userInfo = res.userInfo
      localStorage.setItem('token', res.token)
    },
    
    async getUserInfo() {
      const userInfo = await getUserInfo()
      this.userInfo = userInfo
    },
    
    logout() {
      this.token = ''
      this.userInfo = null
      localStorage.removeItem('token')
    },
  },
})
```

### 2.4 路由配置 (router/index.ts)

```typescript
import { createRouter, createWebHistory } from 'vue-router'
import type { RouteRecordRaw } from 'vue-router'
import { useUserStore } from '@/stores/user'

const routes: RouteRecordRaw[] = [
  {
    path: '/login',
    name: 'Login',
    component: () => import('@/views/login/index.vue'),
    meta: { requiresAuth: false },
  },
  {
    path: '/',
    component: () => import('@/components/BasicLayout/index.vue'),
    redirect: '/dashboard',
    children: [
      {
        path: 'dashboard',
        name: 'Dashboard',
        component: () => import('@/views/dashboard/index.vue'),
        meta: { title: '首页' },
      },
      // 系统管理路由
      {
        path: 'sys',
        children: [
          {
            path: 'user',
            name: 'SysUser',
            component: () => import('@/views/sys/user/index.vue'),
            meta: { title: '用户管理' },
          },
          // ... 其他系统管理路由
        ],
      },
      // 教学管理路由
      {
        path: 'teaching',
        children: [
          {
            path: 'course',
            name: 'TeachingCourse',
            component: () => import('@/views/teaching/course/index.vue'),
            meta: { title: '课程管理' },
          },
          // ... 其他教学管理路由
        ],
      },
    ],
  },
]

const router = createRouter({
  history: createWebHistory(),
  routes,
})

// 路由守卫
router.beforeEach((to, from, next) => {
  const userStore = useUserStore()
  
  if (to.meta.requiresAuth !== false && !userStore.isLoggedIn) {
    next('/login')
  } else {
    next()
  }
})

export default router
```

---

## 三、系统管理页面开发 (Week 2)

### 3.1 登录页面 (Day 1)

**功能清单**:
- [ ] 用户名密码登录
- [ ] 表单验证
- [ ] 记住密码
- [ ] 登录状态管理

**文件**: `src/views/login/index.vue`

### 3.2 用户管理 (Day 2-3)

**功能清单**:
- [ ] 用户列表展示（表格、分页）
- [ ] 新增用户弹窗
- [ ] 编辑用户
- [ ] 删除用户
- [ ] 搜索筛选（用户名、真实姓名、状态）
- [ ] 密码修改

**文件**:
- `src/views/sys/user/index.vue` - 主页面
- `src/views/sys/user/UserModal.vue` - 用户弹窗
- `src/api/sys/user.ts` - API接口

### 3.3 角色管理 (Day 4)

**功能清单**:
- [ ] 角色列表
- [ ] 新增/编辑角色
- [ ] 删除角色
- [ ] 权限分配弹窗（树形选择）

**文件**:
- `src/views/sys/role/index.vue`
- `src/views/sys/role/RoleModal.vue`
- `src/views/sys/role/PermissionModal.vue`
- `src/api/sys/role.ts`

### 3.4 权限管理 (Day 5)

**功能清单**:
- [ ] 权限列表（树形展示）
- [ ] 新增/编辑权限
- [ ] 删除权限

**文件**:
- `src/views/sys/permission/index.vue`
- `src/api/sys/permission.ts`

---

## 四、教学管理页面开发 (Week 3)

### 4.1 课程管理 (Day 1-2)

**功能清单**:
- [ ] 课程列表（卡片或列表视图）
- [ ] 新增/编辑课程
- [ ] 删除课程
- [ ] 课程详情页
- [ ] 课程单元管理

**文件**:
- `src/views/teaching/course/index.vue`
- `src/views/teaching/course/CourseModal.vue`
- `src/views/teaching/course/CourseDetail.vue`
- `src/views/teaching/course/UnitManager.vue`
- `src/api/teaching/course.ts`

### 4.2 作品管理 (Day 3-4)

**功能清单**:
- [ ] 作品列表
- [ ] 作品详情
- [ ] 作品批改
- [ ] 作品评论
- [ ] 作品标签

**文件**:
- `src/views/teaching/work/index.vue`
- `src/views/teaching/work/WorkDetail.vue`
- `src/views/teaching/work/WorkCorrect.vue`
- `src/api/teaching/work.ts`

### 4.3 新闻公告 (Day 5)

**功能清单**:
- [ ] 新闻列表
- [ ] 新增/编辑新闻
- [ ] 新闻详情
- [ ] 富文本编辑器集成

**文件**:
- `src/views/teaching/news/index.vue`
- `src/views/teaching/news/NewsModal.vue`
- `src/views/teaching/news/NewsDetail.vue`
- `src/api/teaching/news.ts`

---

## 五、统计报表页面 (Week 4, Day 1-2)

### 5.1 部门日志统计

**功能清单**:
- [ ] 统计报表（表格展示）
- [ ] 按部门统计图表（柱状图）
- [ ] 按月份统计图表（折线图）
- [ ] 日期范围筛选
- [ ] ECharts图表集成

**文件**:
- `src/views/teaching/stat/DepartDayLog.vue`
- `src/components/Charts/BarChart.vue`
- `src/components/Charts/LineChart.vue`
- `src/api/teaching/departDayLog.ts`

---

## 六、通用组件开发 (Week 4, Day 3-5)

### 6.1 文件上传组件

```vue
<template>
  <a-upload
    v-model:file-list="fileList"
    :action="uploadUrl"
    :headers="headers"
    @change="handleChange"
  >
    <a-button>
      <upload-outlined />
      上传文件
    </a-button>
  </a-upload>
</template>
```

### 6.2 部门树选择器

```vue
<template>
  <a-tree-select
    v-model:value="selectedValue"
    :tree-data="treeData"
    :load-data="loadData"
    placeholder="请选择部门"
  />
</template>
```

### 6.3 字典选择器

```vue
<template>
  <a-select
    v-model:value="selectedValue"
    :options="dictOptions"
    placeholder="请选择"
  />
</template>
```

---

## 七、测试与优化 (Week 4, Day 5+)

### 7.1 功能测试清单

- [ ] 登录/登出功能
- [ ] 路由导航
- [ ] 权限控制
- [ ] API请求/响应
- [ ] 表单验证
- [ ] 文件上传
- [ ] 分页功能
- [ ] 搜索筛选

### 7.2 性能优化

- [ ] 路由懒加载
- [ ] 组件懒加载
- [ ] 图片懒加载
- [ ] 防抖节流
- [ ] Vite构建优化

### 7.3 浏览器兼容性测试

- [ ] Chrome 90+
- [ ] Firefox 88+
- [ ] Safari 14+
- [ ] Edge 90+

---

## 八、部署配置

### 8.1 生产构建

```bash
npm run build
```

### 8.2 Nginx配置示例

```nginx
server {
    listen 80;
    server_name teaching-open.com;
    
    root /var/www/teaching-open/web-vue3/dist;
    index index.html;
    
    location / {
        try_files $uri $uri/ /index.html;
    }
    
    location /api {
        proxy_pass http://localhost:8199;
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
    }
}
```

---

## 九、Git提交规范

每完成一个模块后提交:

```bash
git add .
git commit -m "feat(web): [模块名] 实现[功能描述]"
git push origin devgo
```

示例:
```bash
git commit -m "feat(web): [login] 实现登录页面和认证逻辑"
git commit -m "feat(web): [user] 实现用户管理CRUD功能"
git commit -m "feat(web): [course] 实现课程管理和单元管理"
```

---

## 十、注意事项

1. **代码规范**: 遵循 `docs/Vue3 Dev Guide.md` 中的代码规范
2. **类型定义**: 充分利用TypeScript类型系统
3. **组件复用**: 提取通用组件,避免代码重复
4. **性能优化**: 注意虚拟滚动、分页加载
5. **错误处理**: 统一错误提示和边界情况处理
6. **响应式设计**: 考虑移动端适配

---

## 十一、后续扩展 (可选)

- [ ] 单元测试 (Vitest)
- [ ] E2E测试 (Playwright)
- [ ] 国际化 (i18n)
- [ ] 主题切换
- [ ] PWA支持
- [ ] 性能监控

---

**计划制定人**: AI Assistant  
**制定日期**: 2025-12-14  
**审核状态**: 待审核  
**开始时间**: Phase 12完成后
