# GoFrame V2 开发指南 - Teaching Open 后端迁移方案

## 一、项目概述

### 1.1 迁移目标
将现有的 Spring Boot + MyBatis-Plus 架构的 Teaching Open 后端系统迁移到 GoFrame V2 框架，实现:
- 更高的性能和并发能力
- 更低的资源占用
- 更简洁的代码结构
- 保持原有业务功能完整性

### 1.2 技术栈对比

| 组件 | Java 原系统 | Golang 目标系统 |
|------|------------|----------------|
| 基础框架 | Spring Boot 2.1.3 | GoFrame V2.7+ |
| ORM | MyBatis-Plus 3.1.2 | GoFrame Gen + ORM |
| 认证授权 | Shiro + JWT | JWT + GoFrame Middleware |
| 缓存 | Redis + Spring Cache | GoFrame Cache + Redis |
| 数据库连接池 | Druid | GoFrame DB Pool |
| API文档 | Swagger 2.9.2 | OpenAPI 3.0 + Swag |
| 定时任务 | Quartz | GoFrame GCron |
| 日志 | Logback | GoFrame GLog |
| 配置管理 | application.yml | config.yaml |

## 二、项目结构设计

### 2.1 推荐目录结构

```
teaching-open-go/
├── api/                          # API层 - 对外接口定义
│   ├── v1/                       # API v1版本
│   │   ├── system/               # 系统模块API
│   │   │   ├── user.go
│   │   │   ├── role.go
│   │   │   ├── permission.go
│   │   │   └── ...
│   │   └── teaching/             # 教学模块API
│   │       ├── course.go
│   │       ├── work.go
│   │       ├── student.go
│   │       └── ...
│   └── middleware/               # 中间件
│       ├── auth.go               # 认证中间件
│       ├── cors.go               # 跨域中间件
│       ├── logger.go             # 日志中间件
│       └── error.go              # 错误处理中间件
├── internal/
│   ├── cmd/                      # 命令行入口
│   │   └── cmd.go
│   ├── controller/               # 控制器层
│   │   ├── system/
│   │   │   ├── sys_user.go
│   │   │   ├── sys_role.go
│   │   │   ├── sys_permission.go
│   │   │   ├── sys_depart.go
│   │   │   └── ...
│   │   └── teaching/
│   │       ├── teaching_course.go
│   │       ├── teaching_work.go
│   │       ├── teaching_student.go
│   │       └── ...
│   ├── dao/                      # 数据访问层 (自动生成)
│   │   ├── internal/             # 内部DAO实现
│   │   └── system/
│   │       ├── sys_user.go
│   │       └── ...
│   ├── logic/                    # 业务逻辑层
│   │   ├── system/
│   │   │   ├── sys_user.go
│   │   │   ├── sys_role.go
│   │   │   └── ...
│   │   └── teaching/
│   │       ├── teaching_course.go
│   │       ├── teaching_work.go
│   │       └── ...
│   ├── model/                    # 数据模型
│   │   ├── do/                   # Domain Object (自动生成)
│   │   ├── entity/               # 数据实体 (自动生成)
│   │   └── vo/                   # View Object
│   ├── service/                  # 服务接口定义
│   │   ├── system/
│   │   └── teaching/
│   ├── consts/                   # 常量定义
│   │   ├── consts.go
│   │   └── error.go
│   └── packed/                   # 打包资源
├── manifest/                     # 配置文件目录
│   ├── config/
│   │   ├── config.yaml           # 主配置文件
│   │   ├── config.example.yaml   # 配置示例
│   │   └── config.test.yaml      # 测试配置
│   ├── docker/                   # Docker相关
│   │   ├── Dockerfile
│   │   └── docker-compose.yml
│   └── deploy/                   # 部署脚本
│       └── k8s/
├── resource/                     # 资源文件
│   ├── public/                   # 公共资源
│   ├── template/                 # 模板文件
│   └── i18n/                     # 国际化
├── utility/                      # 工具库
│   ├── jwt/                      # JWT工具
│   ├── encrypt/                  # 加密工具
│   ├── upload/                   # 文件上传
│   └── response/                 # 响应封装
├── hack/                         # 工具脚本
│   └── gen/                      # 代码生成脚本
├── go.mod
├── go.sum
├── main.go                       # 程序入口
└── README.md
```

### 2.2 分层架构说明

#### API层 (api/)
- 定义请求和响应结构体
- 参数验证规则
- 接口文档注释 (Swagger)
- 不包含业务逻辑

#### 控制器层 (internal/controller/)
- 处理HTTP请求
- 调用Service层
- 组装响应数据
- 参数转换和校验

#### 业务逻辑层 (internal/logic/)
- 核心业务逻辑实现
- 事务管理
- 数据校验
- 复杂查询组装

#### 数据访问层 (internal/dao/)
- 数据库操作封装
- 由 GoFrame Gen 工具自动生成
- 支持链式操作

#### 服务层 (internal/service/)
- 定义业务接口
- 供Controller调用
- 便于mock和测试

## 三、核心功能迁移方案

### 3.1 用户认证与授权

#### 原Spring Boot + Shiro实现
```java
@RequiresPermissions("user:list")
@GetMapping("/list")
public Result<?> list(@RequestParam Map<String, Object> params) {
    // ...
}
```

#### GoFrame迁移方案

**1. JWT中间件实现**
```go
// utility/jwt/jwt.go
package jwt

import (
    "github.com/gogf/gf/v2/frame/g"
    "github.com/golang-jwt/jwt/v5"
    "time"
)

type Claims struct {
    UserId   string `json:"userId"`
    Username string `json:"username"`
    RoleIds  []string `json:"roleIds"`
    jwt.RegisteredClaims
}

func GenerateToken(userId, username string, roleIds []string) (string, error) {
    claims := Claims{
        UserId:   userId,
        Username: username,
        RoleIds:  roleIds,
        RegisteredClaims: jwt.RegisteredClaims{
            ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
            IssuedAt:  jwt.NewNumericDate(time.Now()),
            Issuer:    "teaching-open",
        },
    }
    
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString([]byte(g.Cfg().MustGet(context.Background(), "jwt.secret").String()))
}

func ParseToken(tokenString string) (*Claims, error) {
    token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
        return []byte(g.Cfg().MustGet(context.Background(), "jwt.secret").String()), nil
    })
    
    if claims, ok := token.Claims.(*Claims); ok && token.Valid {
        return claims, nil
    }
    return nil, err
}
```

**2. 认证中间件**
```go
// api/middleware/auth.go
package middleware

import (
    "github.com/gogf/gf/v2/frame/g"
    "github.com/gogf/gf/v2/net/ghttp"
    "teaching-open-go/utility/jwt"
    "teaching-open-go/utility/response"
)

func Auth(r *ghttp.Request) {
    token := r.Header.Get("X-Access-Token")
    if token == "" {
        response.JsonExit(r, 401, "未登录或token已过期")
        return
    }
    
    claims, err := jwt.ParseToken(token)
    if err != nil {
        response.JsonExit(r, 401, "token验证失败")
        return
    }
    
    // 将用户信息存入上下文
    r.SetCtxVar("userId", claims.UserId)
    r.SetCtxVar("username", claims.Username)
    r.SetCtxVar("roleIds", claims.RoleIds)
    
    r.Middleware.Next()
}

func Permission(permission string) ghttp.HandlerFunc {
    return func(r *ghttp.Request) {
        userId := r.GetCtxVar("userId").String()
        
        // 检查用户权限
        hasPermission, err := service.Permission().CheckUserPermission(r.Context(), userId, permission)
        if err != nil || !hasPermission {
            response.JsonExit(r, 403, "无权限访问")
            return
        }
        
        r.Middleware.Next()
    }
}
```

**3. 使用示例**
```go
// internal/controller/system/sys_user.go
package system

import (
    "github.com/gogf/gf/v2/frame/g"
    "teaching-open-go/api/v1/system"
)

type SysUserController struct{}

func NewSysUser() *SysUserController {
    return &SysUserController{}
}

func (c *SysUserController) List(ctx context.Context, req *system.UserListReq) (res *system.UserListRes, err error) {
    // 业务逻辑
    return service.SysUser().GetList(ctx, req)
}
```

### 3.2 数据库操作

#### 使用GoFrame Gen生成代码

**1. 配置hack/config.yaml**
```yaml
gfcli:
  gen:
    dao:
    - link: "mysql:root:password@tcp(127.0.0.1:3306)/teachingopen"
      tables: "sys_user,sys_role,teaching_course,teaching_work"
      removePrefix: "sys_,teaching_"
      descriptionTag: true
      noModelComment: false
      path: "./internal"
      group: "default"
      jsonCase: "CamelLower"
```

**2. 生成DAO代码**
```bash
gf gen dao
```

**3. 使用生成的DAO**
```go
// internal/logic/system/sys_user.go
package system

import (
    "context"
    "teaching-open-go/internal/dao"
    "teaching-open-go/internal/model/entity"
    "github.com/gogf/gf/v2/database/gdb"
    "github.com/gogf/gf/v2/frame/g"
)

type sSysUser struct{}

func NewSysUser() *sSysUser {
    return &sSysUser{}
}

func (s *sSysUser) GetList(ctx context.Context, page, pageSize int, username string) (list []*entity.SysUser, total int, err error) {
    model := dao.SysUser.Ctx(ctx)
    
    // 条件查询
    if username != "" {
        model = model.WhereLike("username", "%"+username+"%")
    }
    
    // 分页查询
    total, err = model.Count()
    if err != nil {
        return
    }
    
    err = model.Page(page, pageSize).OrderDesc("create_time").Scan(&list)
    return
}

func (s *sSysUser) Create(ctx context.Context, user *entity.SysUser) error {
    _, err := dao.SysUser.Ctx(ctx).Data(user).Insert()
    return err
}

func (s *sSysUser) Update(ctx context.Context, user *entity.SysUser) error {
    _, err := dao.SysUser.Ctx(ctx).
        Data(user).
        Where("id", user.Id).
        Update()
    return err
}

func (s *sSysUser) Delete(ctx context.Context, id string) error {
    _, err := dao.SysUser.Ctx(ctx).Where("id", id).Delete()
    return err
}

// 事务示例
func (s *sSysUser) CreateWithRole(ctx context.Context, user *entity.SysUser, roleIds []string) error {
    return dao.SysUser.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
        // 插入用户
        result, err := dao.SysUser.Ctx(ctx).Data(user).Insert()
        if err != nil {
            return err
        }
        
        userId, _ := result.LastInsertId()
        
        // 插入用户角色关系
        for _, roleId := range roleIds {
            _, err = dao.SysUserRole.Ctx(ctx).Data(g.Map{
                "user_id": userId,
                "role_id": roleId,
            }).Insert()
            if err != nil {
                return err
            }
        }
        
        return nil
    })
}
```

### 3.3 Redis缓存

```go
// internal/logic/system/sys_dict.go
package system

import (
    "context"
    "encoding/json"
    "github.com/gogf/gf/v2/frame/g"
    "teaching-open-go/internal/dao"
    "time"
)

func (s *sSysDict) GetDictItems(ctx context.Context, dictCode string) (items []entity.SysDictItem, err error) {
    cacheKey := "dict:" + dictCode
    
    // 先从缓存获取
    cacheData, err := g.Redis().Get(ctx, cacheKey)
    if err == nil && !cacheData.IsEmpty() {
        err = json.Unmarshal(cacheData.Bytes(), &items)
        if err == nil {
            return items, nil
        }
    }
    
    // 缓存未命中，查询数据库
    err = dao.SysDictItem.Ctx(ctx).
        Where("dict_id", dictCode).
        OrderAsc("sort_order").
        Scan(&items)
    
    if err != nil {
        return
    }
    
    // 写入缓存
    data, _ := json.Marshal(items)
    g.Redis().Set(ctx, cacheKey, data)
    g.Redis().Expire(ctx, cacheKey, time.Hour*24)
    
    return
}
```

### 3.4 文件上传

```go
// utility/upload/upload.go
package upload

import (
    "context"
    "fmt"
    "github.com/gogf/gf/v2/frame/g"
    "github.com/gogf/gf/v2/net/ghttp"
    "github.com/gogf/gf/v2/os/gfile"
    "github.com/gogf/gf/v2/os/gtime"
    "path"
)

type UploadService struct{}

func New() *UploadService {
    return &UploadService{}
}

func (s *UploadService) UploadFile(ctx context.Context, r *ghttp.Request) (url string, err error) {
    file := r.GetUploadFile("file")
    if file == nil {
        return "", fmt.Errorf("请选择文件")
    }
    
    // 生成文件名
    uploadPath := g.Cfg().MustGet(ctx, "upload.path").String()
    datePath := gtime.Now().Format("Y/m/d")
    savePath := path.Join(uploadPath, datePath)
    
    // 创建目录
    if !gfile.Exists(savePath) {
        gfile.Mkdir(savePath)
    }
    
    // 保存文件
    fileName := gtime.TimestampMicroStr() + path.Ext(file.Filename)
    filePath := path.Join(savePath, fileName)
    
    _, err = file.Save(filePath)
    if err != nil {
        return "", err
    }
    
    // 返回访问URL
    url = fmt.Sprintf("/upload/%s/%s", datePath, fileName)
    return url, nil
}

// 支持OSS上传
func (s *UploadService) UploadToOSS(ctx context.Context, r *ghttp.Request) (url string, err error) {
    // 对接阿里云OSS、七牛云等
    // ...
    return
}
```

### 3.5 定时任务

```go
// internal/logic/system/sys_quartz.go
package system

import (
    "context"
    "github.com/gogf/gf/v2/frame/g"
    "github.com/gogf/gf/v2/os/gcron"
)

type sQuartz struct{}

func NewQuartz() *sQuartz {
    return &sQuartz{}
}

func (s *sQuartz) Init(ctx context.Context) error {
    // 数据备份任务 - 每天凌晨2点
    _, err := gcron.Add(ctx, "0 2 * * *", func(ctx context.Context) {
        g.Log().Info(ctx, "执行数据备份任务")
        // 备份逻辑
    }, "DataBackup")
    
    // 清理临时文件 - 每小时
    _, err = gcron.Add(ctx, "0 * * * *", func(ctx context.Context) {
        g.Log().Info(ctx, "清理临时文件")
        // 清理逻辑
    }, "CleanTemp")
    
    return err
}

func (s *sQuartz) AddJob(ctx context.Context, name, cron string, job func(ctx context.Context)) error {
    _, err := gcron.Add(ctx, cron, job, name)
    return err
}

func (s *sQuartz) RemoveJob(ctx context.Context, name string) {
    gcron.Remove(name)
}
```

## 四、API接口设计规范

### 4.1 统一响应格式

```go
// utility/response/response.go
package response

import (
    "github.com/gogf/gf/v2/net/ghttp"
)

type Response struct {
    Code    int         `json:"code"`
    Message string      `json:"message"`
    Result  interface{} `json:"result,omitempty"`
    Success bool        `json:"success"`
    Timestamp int64     `json:"timestamp"`
}

func Json(r *ghttp.Request, code int, message string, data ...interface{}) {
    responseData := interface{}(nil)
    if len(data) > 0 {
        responseData = data[0]
    }
    
    r.Response.WriteJson(Response{
        Code:      code,
        Message:   message,
        Result:    responseData,
        Success:   code == 200,
        Timestamp: gtime.Now().Timestamp(),
    })
}

func JsonExit(r *ghttp.Request, code int, message string) {
    Json(r, code, message)
    r.Exit()
}

func Success(r *ghttp.Request, data ...interface{}) {
    Json(r, 200, "操作成功", data...)
}

func Error(r *ghttp.Request, message string) {
    Json(r, 500, message)
}
```

### 4.2 API定义示例

```go
// api/v1/teaching/course.go
package teaching

import (
    "github.com/gogf/gf/v2/frame/g"
)

// 课程列表请求
type CourseListReq struct {
    g.Meta   `path:"/course/list" method:"get" tags:"教学管理" summary:"课程列表"`
    Page     int    `json:"page" v:"required|min:1" dc:"页码"`
    PageSize int    `json:"pageSize" v:"required|min:1|max:100" dc:"每页数量"`
    Name     string `json:"name" dc:"课程名称"`
    Type     string `json:"type" dc:"课程类型"`
}

type CourseListRes struct {
    List     []CourseItem `json:"records"`
    Total    int          `json:"total"`
    Page     int          `json:"pageNo"`
    PageSize int          `json:"pageSize"`
}

type CourseItem struct {
    Id          string `json:"id"`
    Name        string `json:"name"`
    Type        string `json:"type"`
    Description string `json:"description"`
    CreateTime  string `json:"createTime"`
}

// 创建课程
type CourseCreateReq struct {
    g.Meta      `path:"/course/add" method:"post" tags:"教学管理" summary:"创建课程"`
    Name        string   `json:"name" v:"required" dc:"课程名称"`
    Type        string   `json:"type" v:"required" dc:"课程类型"`
    Description string   `json:"description" dc:"课程描述"`
    DeptIds     []string `json:"deptIds" dc:"关联部门"`
}

type CourseCreateRes struct {
    Id string `json:"id"`
}
```

## 五、配置管理

### 5.1 主配置文件

```yaml
# manifest/config/config.yaml
server:
  address: ":8080"
  serverRoot: "resource/public"
  logPath: "logs/server"
  logStdout: true
  errorLogEnabled: true
  accessLogEnabled: true
  graceful: true

database:
  default:
    link: "mysql:root:password@tcp(127.0.0.1:3306)/teachingopen"
    debug: true
    charset: "utf8mb4"
    maxIdle: 10
    maxOpen: 100
    maxLifetime: 30

redis:
  default:
    address: "127.0.0.1:6379"
    db: 0
    pass: ""
    maxActive: 100

logger:
  path: "logs"
  level: "all"
  stdout: true

jwt:
  secret: "teaching-open-secret-key"
  expire: 86400

upload:
  path: "resource/upload"
  maxSize: 10485760  # 10MB

# 跨域配置
cors:
  allowOrigin: "*"
  allowMethods: "GET,POST,PUT,DELETE,OPTIONS"
  allowHeaders: "Origin,Content-Type,Accept,X-Access-Token"
```

## 六、部署方案

### 6.1 Docker部署

```dockerfile
# manifest/docker/Dockerfile
FROM golang:1.21-alpine AS builder

WORKDIR /app
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o teaching-open-go main.go

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /app
COPY --from=builder /app/teaching-open-go .
COPY --from=builder /app/manifest ./manifest
COPY --from=builder /app/resource ./resource

EXPOSE 8080
CMD ["./teaching-open-go"]
```

```yaml
# manifest/docker/docker-compose.yml
version: '3.8'

services:
  app:
    build: .
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
      - ./resource/upload:/app/resource/upload

  mysql:
    image: mysql:8.0
    environment:
      MYSQL_ROOT_PASSWORD: password
      MYSQL_DATABASE: teachingopen
    ports:
      - "3306:3306"
    volumes:
      - mysql_data:/var/lib/mysql

  redis:
    image: redis:7-alpine
    ports:
      - "6379:6379"

volumes:
  mysql_data:
```

### 6.2 启动命令

```bash
# 开发环境
go run main.go

# 编译
go build -o teaching-open-go main.go

# 运行
./teaching-open-go

# Docker部署
docker-compose up -d
```

## 七、迁移步骤

### 7.1 阶段一：基础框架搭建 (1-2周)
1. 创建项目结构
2. 配置数据库连接
3. 使用gf gen dao生成DAO层代码
4. 实现JWT认证中间件
5. 实现统一响应格式

### 7.2 阶段二：系统模块迁移 (2-3周)
1. 用户管理模块
2. 角色权限模块
3. 部门管理模块
4. 菜单管理模块
5. 字典管理模块
6. 文件上传模块

### 7.3 阶段三：教学模块迁移 (3-4周)
1. 课程管理
2. 作品管理
3. 学生管理
4. 评论系统
5. Scratch资源管理
6. 新闻公告

### 7.4 阶段四：测试与优化 (1-2周)
1. 单元测试
2. 接口测试
3. 性能测试
4. 压力测试
5. 安全测试

## 八、性能优化建议

### 8.1 数据库优化
- 使用索引优化查询
- 合理使用事务
- 避免N+1查询问题
- 使用缓存减少数据库压力

### 8.2 缓存策略
- 热点数据缓存(字典、菜单等)
- 用户信息缓存
- 查询结果缓存
- 设置合理的过期时间

### 8.3 并发处理
- 使用goroutine处理耗时任务
- 使用channel进行协程通信
- 合理设置连接池大小

## 九、监控与日志

### 9.1 日志规范
```go
// 使用GoFrame内置日志
g.Log().Info(ctx, "用户登录", g.Map{"username": username})
g.Log().Error(ctx, "数据库错误", err)
g.Log().Debug(ctx, "调试信息", data)
```

### 9.2 性能监控
- 使用pprof进行性能分析
- 接入Prometheus监控
- 接入链路追踪(Jaeger/Zipkin)

## 十、开发工具推荐

1. **IDE**: GoLand / VSCode with Go extension
2. **API测试**: Postman / Apifox
3. **数据库工具**: Navicat / DBeaver
4. **性能分析**: pprof / Grafana
5. **代码生成**: GoFrame CLI (`gf`)

## 十一、参考资料

- [GoFrame官方文档](https://goframe.org)
- [GoFrame V2教程](https://goframe.org/pages/viewpage.action?pageId=1114119)
- [Go语言标准库](https://pkg.go.dev/std)
- [项目示例代码](https://github.com/gogf/gf-demos)
