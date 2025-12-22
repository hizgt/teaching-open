# Phase 13 Vue3 前端开发快速启动指南

## 📋 当前状态

✅ **Phase 12 已完成** (134 APIs, 89%)
- 后端GoFrame部分全部开发完成
- 部门日志统计模块已实现并修复
- 编译成功,等待Git提交

⏳ **Phase 13 待启动** (Vue3前端重构)
- 基于Vite 5.x + Vue 3.4+ + TypeScript
- UI框架: Ant Design Vue 4.x
- 状态管理: Pinia 2.x
- 路由: Vue Router 4.x

---

## 🚀 快速启动步骤

### 1. Git提交Phase 12代码

```bash
cd /root/teaching
git add -A
git commit -m "Phase 12: 部门日志统计完成 - 4 APIs (teaching_depart_day_log)

- DAO层: teaching_depart_day_log CRUD
- Entity/DO: 11字段数据库映射
- API层: 4个统计接口(按日期/部门/月份统计、日志记录)
- Controller: 4个HTTP处理器
- Logic: 4个业务方法
- 累计完成134 APIs (89%)
- 文档更新: changelog.md, PHASE13_PLAN.md"

git push origin devgo
```

### 2. 执行Phase 13初始化脚本

```bash
# 方式1: 自动化脚本
cd /root/teaching
bash init-phase13-vue3.sh

# 方式2: 手动执行(如果脚本失败)
npm create vite@latest web-vue3 -- --template vue-ts
cd web-vue3
npm install
npm install vue-router@4 pinia@2 ant-design-vue@4 axios@1.6 dayjs @vueuse/core
npm install -D @types/node unplugin-vue-components unplugin-auto-import
```

### 3. 配置项目文件

#### 3.1 Vite配置 (`vite.config.ts`)

复制 `docs/20251214/vite.config.example.ts` 的内容到 `web-vue3/vite.config.ts`

**关键配置**:
- ✅ 路径别名 `@` 指向 `src`
- ✅ 代理配置 `/api` -> `http://localhost:8000`
- ✅ 自动导入组件(Ant Design Vue)
- ✅ 自动导入API(vue, vue-router, pinia)

#### 3.2 TypeScript配置 (`tsconfig.json`)

```json
{
  "compilerOptions": {
    "target": "ES2020",
    "useDefineForClassFields": true,
    "module": "ESNext",
    "lib": ["ES2020", "DOM", "DOM.Iterable"],
    "skipLibCheck": true,

    "moduleResolution": "bundler",
    "allowImportingTsExtensions": true,
    "resolveJsonModule": true,
    "isolatedModules": true,
    "noEmit": true,
    "jsx": "preserve",

    "strict": true,
    "noUnusedLocals": true,
    "noUnusedParameters": true,
    "noFallthroughCasesInSwitch": true,

    "baseUrl": ".",
    "paths": {
      "@/*": ["src/*"],
      "~/*": ["src/*"]
    }
  },
  "include": ["src/**/*.ts", "src/**/*.d.ts", "src/**/*.tsx", "src/**/*.vue"],
  "references": [{ "path": "./tsconfig.node.json" }]
}
```

#### 3.3 环境变量 (`.env.development`)

```env
# 开发环境
VITE_APP_TITLE=Teaching Open
VITE_API_BASE_URL=http://localhost:8000
VITE_PORT=3000
```

`.env.production`:

```env
# 生产环境
VITE_APP_TITLE=Teaching Open
VITE_API_BASE_URL=https://api.teachingopen.com
```

### 4. 创建目录结构

```bash
cd web-vue3
mkdir -p src/{api,assets/{images,styles},components,composables,hooks,layouts,router,stores,types,utils,views/{system,teaching,home,login,error}}
```

### 5. 创建核心文件

#### 5.1 Axios封装 (`src/utils/request.ts`)

复制 `docs/20251214/utils_request.example.ts` 内容

**功能**:
- ✅ JWT token自动添加
- ✅ 统一错误处理
- ✅ 401自动跳转登录
- ✅ 响应数据解包

#### 5.2 用户状态管理 (`src/stores/user.ts`)

复制 `docs/20251214/stores_user.example.ts` 内容

**功能**:
- ✅ 登录/登出
- ✅ 用户信息管理
- ✅ 权限判断
- ✅ 本地持久化

#### 5.3 路由配置 (`src/router/index.ts` + `guards.ts`)

复制示例文件内容:
- `docs/20251214/router_index.example.ts` -> `src/router/index.ts`
- `docs/20251214/router_guards.example.ts` -> `src/router/guards.ts`

**路由结构**:
```
/login              登录页
/home               首页
/system             系统管理
  /user             用户管理
  /role             角色管理
  /permission       权限管理
  /depart           部门管理
  /dict             字典管理
  /log              日志管理
/teaching           教学管理
  /course           课程管理
  /work             作品管理
  /news             新闻公告
  /homework         附加作业
  /material         Scratch素材
```

### 6. 创建API接口文件

#### 6.1 用户API (`src/api/user.ts`)

```typescript
import request from '@/utils/request'

// 登录
export function loginApi(data: { username: string; password: string }) {
  return request.post('/api/login', data)
}

// 获取用户信息
export function getUserInfoApi() {
  return request.get('/api/user/info')
}

// 获取用户列表
export function getUserListApi(params: any) {
  return request.get('/api/user/list', params)
}

// 新增用户
export function addUserApi(data: any) {
  return request.post('/api/user/add', data)
}

// 编辑用户
export function editUserApi(data: any) {
  return request.put('/api/user/edit', data)
}

// 删除用户
export function deleteUserApi(id: string) {
  return request.delete(`/api/user/delete?id=${id}`)
}
```

### 7. 创建基础布局组件

#### 7.1 主布局 (`src/layouts/BasicLayout.vue`)

```vue
<template>
  <a-layout class="basic-layout">
    <a-layout-sider v-model:collapsed="collapsed" collapsible>
      <div class="logo">Teaching Open</div>
      <a-menu
        theme="dark"
        mode="inline"
        v-model:selectedKeys="selectedKeys"
        :items="menuItems"
        @click="handleMenuClick"
      />
    </a-layout-sider>

    <a-layout>
      <a-layout-header class="header">
        <div class="header-left">
          <menu-unfold-outlined
            v-if="collapsed"
            class="trigger"
            @click="() => (collapsed = !collapsed)"
          />
          <menu-fold-outlined v-else class="trigger" @click="() => (collapsed = !collapsed)" />
        </div>
        <div class="header-right">
          <a-dropdown>
            <a-avatar :src="userInfo?.avatar || '/avatar-default.png'" />
            <template #overlay>
              <a-menu>
                <a-menu-item key="profile">个人中心</a-menu-item>
                <a-menu-item key="settings">设置</a-menu-item>
                <a-menu-divider />
                <a-menu-item key="logout" @click="handleLogout">退出登录</a-menu-item>
              </a-menu>
            </template>
          </a-dropdown>
        </div>
      </a-layout-header>

      <a-layout-content class="content">
        <router-view />
      </a-layout-content>

      <a-layout-footer class="footer">
        Teaching Open ©2025 Created by Teaching Team
      </a-layout-footer>
    </a-layout>
  </a-layout>
</template>

<script setup lang="ts">
import { ref, computed } from 'vue'
import { useRouter } from 'vue-router'
import { useUserStore } from '@/stores/user'
import { MenuUnfoldOutlined, MenuFoldOutlined } from '@ant-design/icons-vue'

const router = useRouter()
const userStore = useUserStore()

const collapsed = ref(false)
const selectedKeys = ref<string[]>(['home'])

const userInfo = computed(() => userStore.userInfo)

const menuItems = [
  { key: 'home', icon: 'HomeOutlined', label: '首页', path: '/home' },
  {
    key: 'system',
    icon: 'SettingOutlined',
    label: '系统管理',
    children: [
      { key: 'system-user', label: '用户管理', path: '/system/user' },
      { key: 'system-role', label: '角色管理', path: '/system/role' },
      { key: 'system-permission', label: '权限管理', path: '/system/permission' },
      { key: 'system-depart', label: '部门管理', path: '/system/depart' },
      { key: 'system-dict', label: '字典管理', path: '/system/dict' },
      { key: 'system-log', label: '日志管理', path: '/system/log' },
    ],
  },
  {
    key: 'teaching',
    icon: 'ReadOutlined',
    label: '教学管理',
    children: [
      { key: 'teaching-course', label: '课程管理', path: '/teaching/course' },
      { key: 'teaching-work', label: '作品管理', path: '/teaching/work' },
      { key: 'teaching-news', label: '新闻公告', path: '/teaching/news' },
      { key: 'teaching-homework', label: '附加作业', path: '/teaching/homework' },
      { key: 'teaching-material', label: 'Scratch素材', path: '/teaching/material' },
    ],
  },
]

const handleMenuClick = ({ key, keyPath }: any) => {
  const findPath = (items: any[], targetKey: string): string | null => {
    for (const item of items) {
      if (item.key === targetKey && item.path) return item.path
      if (item.children) {
        const childPath = findPath(item.children, targetKey)
        if (childPath) return childPath
      }
    }
    return null
  }

  const path = findPath(menuItems, key)
  if (path) router.push(path)
}

const handleLogout = () => {
  userStore.logout()
  router.push('/login')
}
</script>

<style scoped lang="less">
.basic-layout {
  min-height: 100vh;

  .logo {
    height: 64px;
    display: flex;
    align-items: center;
    justify-content: center;
    color: #fff;
    font-size: 18px;
    font-weight: bold;
  }

  .header {
    background: #fff;
    padding: 0 24px;
    display: flex;
    justify-content: space-between;
    align-items: center;
    box-shadow: 0 1px 4px rgba(0, 21, 41, 0.08);

    .trigger {
      font-size: 18px;
      cursor: pointer;
      transition: color 0.3s;

      &:hover {
        color: #1890ff;
      }
    }
  }

  .content {
    margin: 16px;
    padding: 24px;
    background: #fff;
    min-height: 280px;
  }

  .footer {
    text-align: center;
  }
}
</style>
```

#### 7.2 登录布局 (`src/layouts/UserLayout.vue`)

```vue
<template>
  <div class="user-layout">
    <div class="user-layout-content">
      <router-view />
    </div>
    <div class="user-layout-footer">
      Teaching Open ©2025 Created by Teaching Team
    </div>
  </div>
</template>

<style scoped lang="less">
.user-layout {
  min-height: 100vh;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;

  &-content {
    width: 400px;
    padding: 40px;
    background: #fff;
    border-radius: 8px;
    box-shadow: 0 4px 12px rgba(0, 0, 0, 0.15);
  }

  &-footer {
    margin-top: 24px;
    color: #fff;
    font-size: 14px;
  }
}
</style>
```

### 8. 创建登录页面

#### 8.1 登录页 (`src/views/login/index.vue`)

```vue
<template>
  <div class="login-container">
    <div class="login-title">Teaching Open</div>
    <a-form
      :model="formState"
      :rules="rules"
      @finish="handleSubmit"
      layout="vertical"
    >
      <a-form-item label="用户名" name="username">
        <a-input
          v-model:value="formState.username"
          placeholder="请输入用户名"
          size="large"
        >
          <template #prefix>
            <UserOutlined />
          </template>
        </a-input>
      </a-form-item>

      <a-form-item label="密码" name="password">
        <a-input-password
          v-model:value="formState.password"
          placeholder="请输入密码"
          size="large"
        >
          <template #prefix>
            <LockOutlined />
          </template>
        </a-input-password>
      </a-form-item>

      <a-form-item>
        <a-button type="primary" html-type="submit" size="large" block :loading="loading">
          登录
        </a-button>
      </a-form-item>
    </a-form>
  </div>
</template>

<script setup lang="ts">
import { reactive, ref } from 'vue'
import { useRouter } from 'vue-router'
import { useUserStore } from '@/stores/user'
import { UserOutlined, LockOutlined } from '@ant-design/icons-vue'

const router = useRouter()
const userStore = useUserStore()

const loading = ref(false)
const formState = reactive({
  username: '',
  password: '',
})

const rules = {
  username: [{ required: true, message: '请输入用户名', trigger: 'blur' }],
  password: [{ required: true, message: '请输入密码', trigger: 'blur' }],
}

const handleSubmit = async () => {
  loading.value = true
  try {
    await userStore.login(formState.username, formState.password)
    router.push('/')
  } catch (error) {
    console.error('登录失败:', error)
  } finally {
    loading.value = false
  }
}
</script>

<style scoped lang="less">
.login-container {
  .login-title {
    text-align: center;
    font-size: 28px;
    font-weight: bold;
    margin-bottom: 32px;
    color: #333;
  }
}
</style>
```

### 9. 启动开发服务器

```bash
cd web-vue3
npm run dev
```

访问: http://localhost:3000

---

## 📦 项目结构

```
web-vue3/
├── public/                 # 静态资源
├── src/
│   ├── api/               # API接口
│   │   ├── user.ts
│   │   ├── role.ts
│   │   └── ...
│   ├── assets/            # 资源文件
│   │   ├── images/
│   │   └── styles/
│   ├── components/        # 通用组件
│   ├── composables/       # 组合式函数
│   ├── hooks/             # 自定义Hooks
│   ├── layouts/           # 布局组件
│   │   ├── BasicLayout.vue
│   │   └── UserLayout.vue
│   ├── router/            # 路由
│   │   ├── index.ts
│   │   └── guards.ts
│   ├── stores/            # Pinia状态管理
│   │   ├── user.ts
│   │   ├── app.ts
│   │   └── permission.ts
│   ├── types/             # TypeScript类型
│   ├── utils/             # 工具函数
│   │   └── request.ts     # Axios封装
│   ├── views/             # 页面组件
│   │   ├── login/
│   │   ├── home/
│   │   ├── system/
│   │   │   ├── user/
│   │   │   ├── role/
│   │   │   └── ...
│   │   └── teaching/
│   │       ├── course/
│   │       ├── work/
│   │       └── ...
│   ├── App.vue
│   └── main.ts
├── .env.development       # 开发环境变量
├── .env.production        # 生产环境变量
├── vite.config.ts         # Vite配置
├── tsconfig.json          # TypeScript配置
└── package.json
```

---

## 🔧 后续开发任务

### Week 1: 系统管理页面 (优先级P0)

1. ✅ 登录页面
2. ⏳ 用户管理页面
3. ⏳ 角色管理页面
4. ⏳ 权限管理页面
5. ⏳ 部门管理页面
6. ⏳ 字典管理页面
7. ⏳ 日志管理页面

### Week 2-3: 教学管理页面 (优先级P1)

1. ⏳ 课程管理页面
2. ⏳ 作品管理页面
3. ⏳ 新闻公告页面
4. ⏳ 附加作业页面
5. ⏳ Scratch素材页面

### Week 4: 编辑器集成 + 优化

1. ⏳ Scratch 3.0编辑器集成
2. ⏳ Python编辑器集成
3. ⏳ 性能优化(懒加载、虚拟滚动)
4. ⏳ 用户体验优化(骨架屏、Loading)

---

## 📝 注意事项

1. **API接口对接**: 确保后端GoFrame服务已启动 (默认端口8000)
2. **JWT认证**: 登录成功后token自动存储在localStorage
3. **权限控制**: 路由守卫会检查用户权限,没有权限会跳转403
4. **错误处理**: 所有API请求失败会自动弹出错误提示
5. **代码规范**: 使用ESLint + Prettier保持代码风格一致

---

## 🎯 开发规范

1. **组件命名**: PascalCase (如 `UserList.vue`)
2. **文件命名**: kebab-case (如 `user-list.ts`)
3. **变量命名**: camelCase (如 `userName`)
4. **常量命名**: UPPER_SNAKE_CASE (如 `API_BASE_URL`)
5. **注释规范**: 
   - 组件顶部写功能说明
   - 复杂逻辑写清楚注释
   - API接口写清楚参数和返回值

---

## 📚 参考文档

- Vue 3: https://cn.vuejs.org/
- Vite: https://cn.vitejs.dev/
- Vue Router: https://router.vuejs.org/zh/
- Pinia: https://pinia.vuejs.org/zh/
- Ant Design Vue: https://antdv.com/components/overview-cn
- GoFrame V2: https://goframe.org/

---

**生成日期**: 2025-12-14  
**文档版本**: 1.0  
**维护人**: AI Assistant
