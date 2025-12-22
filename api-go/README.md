# Teaching Open Go 后端

基于 GoFrame V2 的教学管理平台后端服务。

## 功能特性

- ✅ JWT 用户认证与授权
- ✅ Token 自动刷新机制
- ✅ 多设备登录限制
- ✅ 权限验证优化
- ✅ 安全加固功能
- ✅ 统一响应格式
- ✅ 请求限流保护
- ✅ 跨域支持
- ✅ 请求日志记录

## 项目结构

```
api-go/
├── api/                          # API层
│   ├── middleware/               # 中间件
│   │   ├── auth.go              # 认证中间件
│   │   └── ...
│   └── v1/                       # API v1版本
├── internal/                     # 内部代码
│   ├── cmd/                      # 命令行入口
│   ├── controller/               # 控制器层
│   ├── dao/                      # 数据访问层
│   ├── logic/                    # 业务逻辑层
│   ├── model/                    # 数据模型
│   └── service/                  # 服务接口
├── manifest/                     # 配置文件
│   └── config/                   # 配置目录
│       ├── config.yaml          # 主配置文件
│       └── config.example.yaml  # 配置示例
├── utility/                      # 工具库
│   ├── jwt/                     # JWT工具
│   └── response/                # 响应封装
├── go.mod
├── go.sum
└── README.md
```

## 快速开始

### 1. 环境要求

- Go 1.21+
- MySQL 8.0+
- Redis 7.0+

### 2. 安装依赖

```bash
go mod tidy
```

### 3. 配置数据库

复制配置文件并修改数据库连接信息：

```bash
cp manifest/config/config.example.yaml manifest/config/config.yaml
```

编辑 `manifest/config/config.yaml` 中的数据库和Redis配置。

### 4. 运行项目

```bash
go run main.go
```

## 中间件使用指南

### 认证中间件 (Auth)

#### 基本使用

```go
package main

import (
    "github.com/gogf/gf/v2/frame/g"
    "teaching-open/api/middleware"
)

func main() {
    s := g.Server()

    // 全局使用认证中间件
    s.Use(middleware.Auth)

    // 特定路由使用权限验证
    s.Group("/api/v1", func(group *ghttp.RouterGroup) {
        group.Middleware(middleware.Auth)

        // 需要特定权限
        group.GET("/users", middleware.Permission("user:list"), UserList)

        // 需要管理员角色
        group.POST("/users", middleware.Role("admin"), CreateUser)

        // 仅管理员
        group.DELETE("/users/:id", middleware.AdminOnly, DeleteUser)
    })
}
```

#### 配置选项

在 `config.yaml` 中配置认证相关参数：

```yaml
auth:
  maxDevices: 5          # 最大设备数
  enableAutoRefresh: true # 启用自动刷新
  refreshThreshold: "30m" # 自动刷新阈值
  enableRateLimit: true   # 启用限流
  rateLimitDuration: "1m" # 限流时间窗口
  rateLimitMaxRequests: 60 # 最大请求数
```

### JWT Token 使用

#### 生成 Token

```go
jwtUtil := jwt.New()
tokenPair, err := jwtUtil.GenerateTokenPair(
    userId, username, roleIds, deviceId, deviceType,
)
if err != nil {
    // 处理错误
}

// 返回给客户端
response.Success(r, g.Map{
    "token": tokenPair.AccessToken,
    "refreshToken": tokenPair.RefreshToken,
    "expiresIn": tokenPair.ExpiresIn,
})
```

#### 验证 Token

中间件会自动验证请求中的 Token，支持以下方式：

1. `Authorization: Bearer <token>`
2. `X-Access-Token: <token>`
3. 查询参数 `?token=<token>`

#### 自动刷新

当 Token 剩余时间少于配置的阈值时，中间件会自动在响应头中返回新的 Token：

```
X-New-Access-Token: <new_access_token>
X-New-Refresh-Token: <new_refresh_token>
X-Token-Refreshed: true
```

客户端应该检查这些响应头并更新本地存储的 Token。

### 多设备登录限制

默认最多支持 5 个设备同时登录。当超出限制时，新设备登录会失败。

### 权限验证

#### 权限中间件

```go
// 检查特定权限
group.GET("/users", middleware.Permission("user:list"), handler)

// 检查角色
group.POST("/users", middleware.Role("admin", "manager"), handler)

// 仅管理员
group.DELETE("/users/:id", middleware.AdminOnly, handler)
```

### 安全功能

#### 安全头中间件

```go
s.Use(middleware.SecurityHeaders)
```

添加以下安全响应头：
- `X-Content-Type-Options: nosniff`
- `X-Frame-Options: DENY`
- `X-XSS-Protection: 1; mode=block`
- `Strict-Transport-Security: max-age=31536000; includeSubDomains`
- `Content-Security-Policy: default-src 'self'`

#### 请求限流

基于 IP 地址的请求限流，默认每分钟最多 60 次请求。

#### 日志记录

```go
s.Use(middleware.Logging)
```

记录所有 API 请求的详细信息，包括：
- 请求方法和路径
- 用户ID
- IP地址
- User-Agent
- 响应时间
- 状态码

## API 响应格式

所有 API 响应都遵循统一格式：

```json
{
  "code": 200,
  "message": "操作成功",
  "result": {},
  "success": true,
  "timestamp": 1640995200
}
```

### 状态码说明

- `200`: 成功
- `400`: 请求参数错误
- `401`: 未授权（未登录或Token无效）
- `403`: 禁止访问（无权限）
- `404`: 资源不存在
- `429`: 请求过于频繁
- `500`: 服务器内部错误

## 开发规范

### 代码规范

- 遵循 GoFrame 编码规范
- 使用 `gofmt` 格式化代码
- 使用 `go vet` 检查代码
- 使用 `golint` 检查代码风格

### 提交规范

- 使用英文提交信息
- 格式：`type(scope): description`

类型包括：
- `feat`: 新功能
- `fix`: 修复bug
- `docs`: 文档更新
- `style`: 代码格式调整
- `refactor`: 代码重构
- `test`: 测试相关
- `chore`: 构建过程或工具配置更新

## 部署

### Docker 部署

```bash
# 构建镜像
docker build -t teaching-open-api .

# 运行容器
docker run -d \
  --name teaching-open-api \
  -p 8080:8080 \
  -v $(pwd)/manifest:/app/manifest \
  -v $(pwd)/logs:/app/logs \
  teaching-open-api
```

### Kubernetes 部署

参考 `manifest/deploy/k8s/` 目录下的配置文件。

## 监控和日志

### 日志配置

日志文件存储在 `logs/` 目录下：
- `server.log`: 服务器日志
- `access.log`: 访问日志
- `error.log`: 错误日志

### 性能监控

- 使用 `pprof` 进行性能分析
- 集成 Prometheus 监控指标
- 支持链路追踪

## 常见问题

### Q: Token 过期怎么办？

A: 中间件支持自动刷新，当 Token 即将过期时会自动生成新的 Token 返回给客户端。

### Q: 如何修改最大设备数？

A: 在配置文件中修改 `auth.maxDevices` 参数。

### Q: 如何禁用限流？

A: 设置 `auth.enableRateLimit: false`。

### Q: 如何添加新的权限？

A: 在权限验证中间件中添加相应的权限检查逻辑。

## 贡献指南

1. Fork 项目
2. 创建特性分支 (`git checkout -b feature/AmazingFeature`)
3. 提交更改 (`git commit -m 'Add some AmazingFeature'`)
4. 推送到分支 (`git push origin feature/AmazingFeature`)
5. 创建 Pull Request

## 许可证

本项目采用 MIT 许可证。