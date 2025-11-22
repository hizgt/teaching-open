# Teaching Open 产品需求文档 (PRD)

## 文档信息

| 项目名称 | Teaching Open - 在线教学管理平台 |
|---------|--------------------------------|
| 版本号 | V3.0 (Golang + Vue3重构版) |
| 文档版本 | 1.0 |
| 创建日期 | 2025-11-22 |
| 项目状态 | 技术架构升级规划中 |
| 负责人 | 开发团队 |

## 一、项目概述

### 1.1 项目背景

Teaching Open 是一个面向教育机构的在线教学管理平台，当前基于 Spring Boot + Vue2 技术栈开发。随着业务发展和用户量增长，现有系统面临以下挑战：

- **性能瓶颈**: Java后端在高并发场景下资源占用较高
- **技术债务**: Vue2和相关生态即将停止维护
- **维护成本**: Spring Boot框架较重，部署和运维成本高
- **开发效率**: 代码结构复杂，新功能开发周期长

### 1.2 升级目标

通过技术栈升级实现：

1. **性能提升**: 后端迁移到Golang，提升50%+的并发处理能力，降低70%的资源占用
2. **技术现代化**: 前端升级到Vue3，使用最新的Composition API和TypeScript
3. **降本增效**: 简化部署流程，降低服务器成本，提升开发效率
4. **架构优化**: 采用微服务友好的架构设计，为未来扩展做准备

### 1.3 核心价值主张

- **学生端**: 在线学习编程课程，提交作品，查看评价反馈
- **教师端**: 课程管理，作品批改，学生管理，数据分析
- **管理端**: 系统配置，用户权限，部门管理，数据统计

## 二、用户画像与需求

### 2.1 目标用户

| 用户角色 | 数量占比 | 核心诉求 | 使用频率 |
|---------|---------|---------|---------|
| 学生 | 70% | 在线学习、作品创作、获得反馈 | 每周2-3次 |
| 教师 | 25% | 课程管理、作品批改、教学分析 | 每天使用 |
| 管理员 | 5% | 系统配置、数据统计、权限管理 | 按需使用 |

### 2.2 用户故事

#### 学生用户
- 作为学生，我希望能够浏览课程列表，选择感兴趣的编程课程学习
- 作为学生，我希望能够在线使用Scratch/Python编辑器创作作品
- 作为学生，我希望能够提交作品并查看教师的批改意见
- 作为学生，我希望能够查看我的学习进度和作品历史

#### 教师用户
- 作为教师，我希望能够创建和发布课程，包含课程单元和教学资源
- 作为教师，我希望能够查看学生提交的作品列表，按班级/课程筛选
- 作为教师，我希望能够快速批改作品，给出评分和评语
- 作为教师，我希望能够查看班级的学习统计数据和完成情况

#### 管理员用户
- 作为管理员，我希望能够管理用户账号、角色和权限
- 作为管理员，我希望能够配置系统参数和字典数据
- 作为管理员，我希望能够查看系统日志和操作审计
- 作为管理员，我希望能够导出各类统计报表

## 三、功能模块设计

### 3.1 功能架构图

```
Teaching Open 平台
├── 用户中心
│   ├── 用户注册/登录
│   ├── 个人信息管理
│   ├── 密码修改
│   └── 第三方登录(微信/QQ)
├── 系统管理
│   ├── 用户管理
│   ├── 角色管理
│   ├── 权限管理
│   ├── 部门管理
│   ├── 字典管理
│   ├── 菜单管理
│   ├── 日志管理
│   └── 系统配置
├── 教学管理
│   ├── 课程管理
│   │   ├── 课程创建/编辑
│   │   ├── 课程单元管理
│   │   ├── 课程资源上传
│   │   └── 课程发布/下架
│   ├── 作品管理
│   │   ├── 作品列表
│   │   ├── 作品详情查看
│   │   ├── 作品批改
│   │   └── 作品评论
│   ├── 学生管理
│   │   ├── 学生信息维护
│   │   ├── 班级分配
│   │   └── 学习记录
│   └── 附加作业
│       ├── 作业发布
│       ├── 作业提交
│       └── 作业批改
├── 编程环境
│   ├── Scratch 3.0 编辑器
│   ├── Python 在线编辑器
│   ├── 作品预览运行
│   └── 作品保存/发布
├── 资源管理
│   ├── 文件上传/下载
│   ├── 图片资源管理
│   ├── Scratch素材库
│   └── 素材背包
├── 新闻公告
│   ├── 新闻发布
│   ├── 新闻列表
│   └── 新闻详情
└── 数据统计
    ├── 学习数据统计
    ├── 作品数据统计
    ├── 用户活跃度
    └── 部门日志统计
```

### 3.2 核心功能详细设计

#### 3.2.1 用户认证与授权

**功能描述**: 基于JWT的用户认证，支持RBAC权限模型

**功能点**:
- 用户名密码登录
- 手机号验证码登录
- 第三方登录(微信、QQ)
- JWT Token刷新机制
- 基于角色的访问控制(RBAC)
- 数据权限控制(部门数据隔离)
- 登录日志记录

**接口设计**:
```
POST /sys/login              # 用户登录
POST /sys/logout             # 用户登出
POST /sys/phoneLogin         # 手机号登录
GET  /sys/user/info          # 获取当前用户信息
POST /sys/user/changePassword # 修改密码
GET  /sys/permission/getUserPermissionByToken # 获取用户权限
```

#### 3.2.2 课程管理

**功能描述**: 教师创建课程，管理课程单元和教学内容

**功能点**:
- 课程CRUD操作
- 课程分类管理(Scratch/Python/其他)
- 课程单元管理(章节结构)
- 课程资源上传(视频、文档、素材)
- 课程发布/下架
- 课程关联部门
- 课程封面图设置

**数据模型**:
```sql
teaching_course (课程表)
- id: 课程ID
- name: 课程名称
- type: 课程类型(scratch/python)
- description: 课程描述
- cover_image: 封面图
- status: 状态(draft/published)
- create_by: 创建人
- create_time: 创建时间

teaching_course_unit (课程单元表)
- id: 单元ID
- course_id: 课程ID
- name: 单元名称
- content: 单元内容
- sort_order: 排序
- resource_url: 资源链接

teaching_course_dept (课程部门关联表)
- course_id: 课程ID
- dept_id: 部门ID
```

**接口设计**:
```
GET    /teaching/course/list          # 课程列表
POST   /teaching/course/add           # 创建课程
PUT    /teaching/course/edit          # 编辑课程
DELETE /teaching/course/delete        # 删除课程
GET    /teaching/course/detail/:id    # 课程详情
GET    /teaching/courseUnit/list      # 课程单元列表
POST   /teaching/courseUnit/add       # 添加单元
PUT    /teaching/courseUnit/edit      # 编辑单元
DELETE /teaching/courseUnit/delete    # 删除单元
```

#### 3.2.3 作品管理

**功能描述**: 学生提交作品，教师批改评价

**功能点**:
- 作品在线创作(Scratch/Python)
- 作品提交
- 作品列表查看(支持按课程、学生、状态筛选)
- 作品详情查看
- 作品批改(评分、评语)
- 作品评论
- 优秀作品展示
- 作品分享

**数据模型**:
```sql
teaching_work (作品表)
- id: 作品ID
- title: 作品标题
- type: 作品类型(scratch/python)
- content: 作品内容(JSON/代码)
- student_id: 学生ID
- course_id: 课程ID
- unit_id: 单元ID
- status: 状态(draft/submitted/corrected)
- score: 分数
- create_time: 创建时间
- submit_time: 提交时间

teaching_work_correct (批改记录表)
- id: 批改ID
- work_id: 作品ID
- teacher_id: 教师ID
- score: 分数
- comment: 评语
- correct_time: 批改时间

teaching_work_comment (评论表)
- id: 评论ID
- work_id: 作品ID
- user_id: 评论用户ID
- content: 评论内容
- create_time: 评论时间
```

**接口设计**:
```
GET    /teaching/work/list            # 作品列表
POST   /teaching/work/add             # 创建作品
PUT    /teaching/work/edit            # 编辑作品
DELETE /teaching/work/delete          # 删除作品
GET    /teaching/work/detail/:id      # 作品详情
POST   /teaching/work/submit          # 提交作品
POST   /teaching/work/correct         # 批改作品
GET    /teaching/work/comment/list    # 评论列表
POST   /teaching/work/comment/add     # 添加评论
```

#### 3.2.4 Scratch编辑器集成

**功能描述**: 集成Scratch 3.0在线编辑器

**功能点**:
- Scratch项目创建/编辑
- 作品运行预览
- 作品保存(自动保存)
- 素材管理(角色、背景、音效)
- 素材背包(个人素材收藏)
- 作品导出/导入(.sb3格式)

**技术方案**:
- 使用官方Scratch 3.0 GUI
- WebSocket实时保存
- IndexedDB本地缓存
- 素材CDN加速

#### 3.2.5 文件管理

**功能描述**: 统一的文件上传、存储和访问

**功能点**:
- 文件上传(支持图片、视频、文档)
- 文件大小限制
- 文件类型校验
- 文件访问权限控制
- 支持多种存储方式(本地/阿里云OSS/七牛云)
- 图片压缩和缩略图

**接口设计**:
```
POST /sys/upload/file              # 上传文件
GET  /sys/upload/view/:id          # 查看文件
DELETE /sys/upload/delete/:id      # 删除文件
POST /sys/upload/batch             # 批量上传
```

### 3.3 非功能需求

#### 3.3.1 性能指标

| 指标 | 目标值 | 说明 |
|-----|-------|------|
| 接口响应时间 | <200ms | 90%接口在200ms内响应 |
| 并发用户数 | 5000+ | 支持5000+在线用户 |
| 数据库查询 | <100ms | 单表查询100ms内完成 |
| 页面加载时间 | <2s | 首屏加载2秒内完成 |
| API可用性 | 99.9% | 年度可用性99.9% |

#### 3.3.2 安全要求

- **认证安全**: JWT Token + RefreshToken机制
- **数据加密**: 密码BCrypt加密，敏感数据AES加密
- **SQL注入防护**: 使用ORM参数化查询
- **XSS防护**: 前端输入过滤，后端输出转义
- **CSRF防护**: Token验证
- **接口限流**: 防止暴力攻击
- **操作审计**: 关键操作日志记录

#### 3.3.3 兼容性

**浏览器支持**:
- Chrome 90+
- Firefox 88+
- Safari 14+
- Edge 90+

**移动端适配**:
- 响应式设计，支持平板和手机访问
- 关键功能移动端优化

## 四、技术架构升级方案

### 4.1 技术选型对比

#### 后端技术栈

| 对比项 | Spring Boot (现状) | GoFrame V2 (目标) | 优势 |
|-------|-------------------|------------------|------|
| 语言 | Java 8 | Golang 1.21+ | 性能提升3-5倍 |
| 框架 | Spring Boot 2.1.3 | GoFrame V2.7+ | 更轻量，启动快 |
| ORM | MyBatis-Plus | GoFrame ORM | 代码自动生成 |
| 内存占用 | 500MB+ | 50MB+ | 降低90% |
| 编译产物 | 80MB+ JAR | 10MB+ 二进制 | 部署简单 |
| 并发能力 | 1000 QPS | 5000+ QPS | 大幅提升 |

#### 前端技术栈

| 对比项 | Vue2 (现状) | Vue3 (目标) | 优势 |
|-------|-----------|-----------|------|
| 核心版本 | Vue 2.6 | Vue 3.4 | 性能提升 |
| API风格 | Options API | Composition API | 更好的复用性 |
| TypeScript | 无 | 完整支持 | 类型安全 |
| 状态管理 | Vuex 3 | Pinia 2 | 更简洁的API |
| 构建工具 | Vue CLI + Webpack | Vite 5 | 快10倍以上 |
| 打包体积 | 800KB+ | 400KB+ | 减少50% |

### 4.2 系统架构图

```
┌─────────────────────────────────────────────────────────┐
│                      客户端层                             │
├─────────────────┬───────────────────┬───────────────────┤
│   Web浏览器      │    移动H5         │    小程序(未来)    │
│  (Vue3 + TS)    │   (响应式)         │                  │
└────────┬────────┴──────────┬────────┴──────────┬────────┘
         │                   │                   │
         └───────────────────┴───────────────────┘
                             │
                    ┌────────▼────────┐
                    │   Nginx/CDN     │
                    │   (负载均衡)      │
                    └────────┬────────┘
                             │
         ┌───────────────────┴───────────────────┐
         │                                       │
┌────────▼────────┐                   ┌─────────▼────────┐
│  API Gateway    │                   │   Static Files   │
│  (可选，未来扩展)  │                   │   (静态资源)      │
└────────┬────────┘                   └──────────────────┘
         │
┌────────▼────────────────────────────────────────────────┐
│               GoFrame应用层 (API Server)                 │
├────────┬────────┬────────┬────────┬────────┬───────────┤
│ 用户中心 │ 系统管理 │ 教学管理 │ 文件管理 │ 定时任务 │ WebSocket │
└────────┴────────┴────────┴────────┴────────┴───────────┘
         │                   │                   │
┌────────▼────────┐  ┌──────▼──────┐   ┌───────▼──────┐
│   MySQL 8.0     │  │   Redis 7   │   │  OSS/本地存储 │
│   (主数据库)     │  │   (缓存)     │   │  (文件存储)   │
└─────────────────┘  └─────────────┘   └──────────────┘
```

### 4.3 数据库设计

#### 核心表结构

**系统管理模块**:
- sys_user: 用户表
- sys_role: 角色表
- sys_permission: 权限表
- sys_user_role: 用户角色关联表
- sys_role_permission: 角色权限关联表
- sys_depart: 部门表
- sys_user_depart: 用户部门关联表
- sys_dict: 字典表
- sys_dict_item: 字典项表
- sys_log: 操作日志表
- sys_data_log: 数据日志表

**教学管理模块**:
- teaching_course: 课程表
- teaching_course_unit: 课程单元表
- teaching_course_dept: 课程部门关联表
- teaching_work: 作品表
- teaching_work_correct: 作品批改表
- teaching_work_comment: 作品评论表
- teaching_student: 学生信息表
- teaching_additional_work: 附加作业表
- teaching_scratch_assets: Scratch素材表
- teaching_news: 新闻公告表
- teaching_menu: 菜单表
- teaching_depart_day_log: 部门日志表

### 4.4 接口设计规范

#### RESTful API规范

```
# 资源命名
GET    /api/v1/users          # 获取用户列表
GET    /api/v1/users/:id      # 获取单个用户
POST   /api/v1/users          # 创建用户
PUT    /api/v1/users/:id      # 更新用户
DELETE /api/v1/users/:id      # 删除用户

# 响应格式
{
  "code": 200,              # 状态码
  "message": "操作成功",     # 提示信息
  "result": {},             # 数据
  "success": true,          # 是否成功
  "timestamp": 1234567890   # 时间戳
}

# 分页格式
{
  "code": 200,
  "result": {
    "records": [],          # 数据列表
    "total": 100,          # 总数
    "pageNo": 1,           # 当前页
    "pageSize": 10         # 每页条数
  }
}
```

## 五、开发计划

### 5.1 项目里程碑

| 阶段 | 时间周期 | 主要产出 | 负责人 |
|-----|---------|---------|-------|
| 需求分析 | 1周 | PRD文档、技术方案 | 产品+技术 |
| 环境搭建 | 1周 | 项目脚手架、CI/CD | 后端+前端 |
| 基础设施 | 2周 | 认证、权限、数据库 | 后端 |
| 系统模块 | 3周 | 用户、角色、权限等 | 全栈 |
| 教学模块 | 4周 | 课程、作品、学生等 | 全栈 |
| 集成测试 | 2周 | 功能测试、性能测试 | 测试 |
| 上线部署 | 1周 | 灰度发布、监控 | 运维 |
| **总计** | **14周** | 完整系统 | 全团队 |

### 5.2 详细开发计划

#### 第1-2周：项目准备与基础搭建
**后端任务**:
- [x] 创建GoFrame项目结构
- [x] 配置数据库连接和Redis
- [x] 使用gf gen dao生成DAO层
- [x] 实现JWT认证中间件
- [x] 实现统一响应格式
- [x] 配置日志系统
- [x] 编写开发文档

**前端任务**:
- [x] 创建Vue3+Vite项目
- [x] 配置TypeScript和ESLint
- [x] 封装Axios和API
- [x] 配置Pinia状态管理
- [x] 配置路由和路由守卫
- [x] 搭建基础布局组件
- [x] 编写开发规范

#### 第3-5周：系统管理模块
**功能列表**:
- [ ] 用户管理(CRUD、冻结、密码重置)
- [ ] 角色管理(CRUD、权限分配)
- [ ] 权限管理(菜单树、数据权限)
- [ ] 部门管理(树形结构、CRUD)
- [ ] 字典管理(CRUD、缓存)
- [ ] 菜单管理(动态菜单)
- [ ] 日志管理(操作日志、数据日志)
- [ ] 文件上传(本地/OSS)

#### 第6-9周：教学管理模块
**功能列表**:
- [ ] 课程管理(CRUD、发布、下架)
- [ ] 课程单元管理
- [ ] 作品管理(列表、详情、批改)
- [ ] 学生管理(信息维护、班级)
- [ ] Scratch编辑器集成
- [ ] Python编辑器集成
- [ ] 素材管理
- [ ] 评论系统
- [ ] 新闻公告

#### 第10-11周：集成测试与优化
- [ ] 单元测试覆盖
- [ ] 接口测试
- [ ] 前端E2E测试
- [ ] 性能测试与优化
- [ ] 压力测试
- [ ] 安全测试
- [ ] 浏览器兼容性测试

#### 第12周：部署上线
- [ ] Docker镜像构建
- [ ] Kubernetes部署配置
- [ ] 灰度发布方案
- [ ] 监控告警配置
- [ ] 数据迁移方案
- [ ] 回滚方案

### 5.3 风险与应对

| 风险项 | 影响 | 概率 | 应对措施 |
|-------|-----|------|---------|
| 技术栈不熟悉 | 高 | 中 | 提前学习、技术分享会 |
| 数据迁移失败 | 高 | 低 | 充分测试、准备回滚方案 |
| 性能不达标 | 中 | 低 | 压力测试、性能优化 |
| 进度延期 | 中 | 中 | 敏捷开发、及时调整 |
| API不兼容 | 中 | 低 | 版本管理、灰度发布 |

## 六、测试方案

### 6.1 测试策略

**测试金字塔**:
```
        ┌─────────┐
        │  E2E    │  (10%)
        ├─────────┤
        │  集成    │  (20%)
        ├─────────┤
        │  单元    │  (70%)
        └─────────┘
```

### 6.2 测试用例

#### 用户登录测试
| 用例ID | 测试场景 | 预期结果 |
|-------|---------|---------|
| TC001 | 正确的用户名密码 | 登录成功，跳转首页 |
| TC002 | 错误的密码 | 提示密码错误 |
| TC003 | 不存在的用户 | 提示用户不存在 |
| TC004 | 账号被冻结 | 提示账号已被冻结 |
| TC005 | Token过期 | 自动跳转登录页 |

#### 作品提交测试
| 用例ID | 测试场景 | 预期结果 |
|-------|---------|---------|
| TC101 | 正常提交作品 | 提交成功，状态变更 |
| TC102 | 重复提交 | 更新作品内容 |
| TC103 | 空内容提交 | 提示内容不能为空 |
| TC104 | 超大文件 | 提示文件过大 |

### 6.3 性能测试

**测试场景**:
- 1000并发用户登录
- 5000并发查询课程列表
- 100并发提交作品
- 1000并发查看作品详情

**性能指标**:
- TPS > 1000
- 平均响应时间 < 200ms
- 错误率 < 0.1%
- CPU使用率 < 70%
- 内存使用 < 2GB

## 七、部署方案

### 7.1 环境规划

| 环境 | 用途 | 配置 | 域名 |
|-----|-----|-----|------|
| 开发环境 | 开发调试 | 2C4G | dev.example.com |
| 测试环境 | 功能测试 | 4C8G | test.example.com |
| 预发布环境 | 上线验证 | 4C8G | pre.example.com |
| 生产环境 | 正式服务 | 8C16G*3 | www.example.com |

### 7.2 Docker部署

```yaml
# docker-compose.yml
version: '3.8'

services:
  # 后端服务
  api:
    image: teaching-open-api:latest
    ports:
      - "8080:8080"
    environment:
      - GF_DATABASE_DEFAULT_LINK=mysql:root:password@tcp(mysql:3306)/teachingopen
      - GF_REDIS_DEFAULT_ADDRESS=redis:6379
    depends_on:
      - mysql
      - redis
    volumes:
      - ./logs:/app/logs
      - ./upload:/app/resource/upload
    restart: always

  # 前端服务
  web:
    image: teaching-open-web:latest
    ports:
      - "80:80"
    depends_on:
      - api
    restart: always

  # MySQL
  mysql:
    image: mysql:8.0
    environment:
      MYSQL_ROOT_PASSWORD: password
      MYSQL_DATABASE: teachingopen
    ports:
      - "3306:3306"
    volumes:
      - mysql_data:/var/lib/mysql
    restart: always

  # Redis
  redis:
    image: redis:7-alpine
    ports:
      - "6379:6379"
    volumes:
      - redis_data:/data
    restart: always

volumes:
  mysql_data:
  redis_data:
```

### 7.3 Kubernetes部署

```yaml
# k8s-deployment.yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: teaching-open-api
spec:
  replicas: 3
  selector:
    matchLabels:
      app: teaching-open-api
  template:
    metadata:
      labels:
        app: teaching-open-api
    spec:
      containers:
      - name: api
        image: teaching-open-api:latest
        ports:
        - containerPort: 8080
        env:
        - name: GF_DATABASE_DEFAULT_LINK
          valueFrom:
            secretKeyRef:
              name: db-secret
              key: connection-string
        resources:
          requests:
            memory: "512Mi"
            cpu: "500m"
          limits:
            memory: "1Gi"
            cpu: "1000m"
        livenessProbe:
          httpGet:
            path: /health
            port: 8080
          initialDelaySeconds: 30
          periodSeconds: 10
        readinessProbe:
          httpGet:
            path: /health
            port: 8080
          initialDelaySeconds: 5
          periodSeconds: 5
---
apiVersion: v1
kind: Service
metadata:
  name: teaching-open-api-service
spec:
  selector:
    app: teaching-open-api
  ports:
  - protocol: TCP
    port: 80
    targetPort: 8080
  type: LoadBalancer
```

### 7.4 CI/CD流程

```yaml
# .github/workflows/deploy.yml
name: Deploy

on:
  push:
    branches: [ master ]

jobs:
  build-and-deploy:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v2
      
      - name: Build Backend
        run: |
          cd api-go
          go build -o teaching-open-go main.go
          
      - name: Build Frontend
        run: |
          cd web-vue3
          npm install
          npm run build
          
      - name: Build Docker Images
        run: |
          docker build -t teaching-open-api:latest ./api-go
          docker build -t teaching-open-web:latest ./web-vue3
          
      - name: Push to Registry
        run: |
          docker push teaching-open-api:latest
          docker push teaching-open-web:latest
          
      - name: Deploy to Kubernetes
        run: |
          kubectl apply -f k8s/
```

## 八、监控与运维

### 8.1 监控指标

**系统指标**:
- CPU使用率
- 内存使用率
- 磁盘使用率
- 网络流量

**应用指标**:
- API响应时间
- API错误率
- 并发连接数
- 数据库连接池

**业务指标**:
- 活跃用户数
- 作品提交量
- 课程访问量
- 系统在线人数

### 8.2 日志管理

**日志分类**:
- 访问日志(access.log)
- 错误日志(error.log)
- 业务日志(business.log)
- SQL日志(sql.log)

**日志格式**:
```json
{
  "time": "2025-11-22 10:00:00",
  "level": "INFO",
  "module": "user",
  "message": "用户登录",
  "userId": "123456",
  "ip": "192.168.1.1",
  "traceId": "abc123"
}
```

### 8.3 告警策略

| 告警项 | 阈值 | 级别 | 通知方式 |
|-------|-----|------|---------|
| CPU使用率 | >80% | P1 | 电话+短信 |
| 内存使用率 | >85% | P1 | 电话+短信 |
| API错误率 | >5% | P2 | 邮件+企业微信 |
| 响应时间 | >1s | P2 | 邮件+企业微信 |
| 服务宕机 | - | P0 | 电话+短信+企业微信 |

## 九、成本分析

### 9.1 人力成本

| 角色 | 人数 | 周期 | 人天 | 单价 | 小计 |
|-----|-----|-----|-----|-----|-----|
| 后端工程师 | 2 | 14周 | 140 | 1000元/天 | 14万 |
| 前端工程师 | 2 | 14周 | 140 | 1000元/天 | 14万 |
| 测试工程师 | 1 | 4周 | 20 | 800元/天 | 1.6万 |
| 产品经理 | 1 | 2周 | 10 | 1000元/天 | 1万 |
| **合计** | 6 | 14周 | 310 | - | **30.6万** |

### 9.2 服务器成本(年)

| 项目 | 配置 | 数量 | 单价(月) | 年费用 |
|-----|-----|-----|---------|--------|
| 应用服务器 | 8C16G | 3 | 500元 | 1.8万 |
| 数据库服务器 | 8C32G | 1 | 800元 | 0.96万 |
| Redis服务器 | 4C8G | 1 | 300元 | 0.36万 |
| OSS存储 | 1TB | 1 | 150元 | 0.18万 |
| CDN流量 | 1TB/月 | - | 200元 | 0.24万 |
| **合计** | - | - | - | **3.54万** |

### 9.3 ROI分析

**成本节约**:
- 服务器成本降低: 30% (从Java到Go，资源占用降低)
- 运维成本降低: 40% (部署更简单，故障更少)
- 开发效率提升: 20% (更现代的技术栈)

**收益提升**:
- 性能提升带来更好的用户体验
- 支持更多并发用户，业务增长空间更大
- 技术债务清理，维护成本降低

## 十、后续规划

### 10.1 短期目标 (3个月内)
- [ ] 完成核心模块迁移
- [ ] 灰度发布50%流量
- [ ] 全量上线新系统
- [ ] 监控系统完善

### 10.2 中期目标 (6-12个月)
- [ ] 移动端App开发
- [ ] 小程序开发
- [ ] AI智能批改
- [ ] 数据分析平台
- [ ] 微服务拆分

### 10.3 长期目标 (1年+)
- [ ] 多租户SaaS化
- [ ] 国际化多语言
- [ ] 在线协作功能
- [ ] 直播教学功能
- [ ] 开放API平台

## 十一、附录

### 11.1 术语表

- **RBAC**: Role-Based Access Control，基于角色的访问控制
- **JWT**: JSON Web Token，一种开放标准的身份验证方案
- **ORM**: Object-Relational Mapping，对象关系映射
- **OSS**: Object Storage Service，对象存储服务
- **CDN**: Content Delivery Network，内容分发网络
- **CI/CD**: Continuous Integration/Continuous Deployment，持续集成/持续部署

### 11.2 参考文档

- [GoFrame官方文档](https://goframe.org)
- [Vue3官方文档](https://cn.vuejs.org/)
- [Ant Design Vue文档](https://antdv.com/)
- [MySQL 8.0文档](https://dev.mysql.com/doc/)
- [Redis文档](https://redis.io/documentation)

### 11.3 变更记录

| 版本 | 日期 | 修改内容 | 修改人 |
|-----|-----|---------|-------|
| 1.0 | 2025-11-22 | 初始版本创建 | 开发团队 |

---

**文档状态**: ✅ 已完成  
**审核状态**: 待审核  
**下次更新**: 根据开发进展持续更新
