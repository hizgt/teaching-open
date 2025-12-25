# Design Document: Java to Go Migration

## Overview

本设计文档详细描述了将 Teaching Open 教学管理平台从 Java (Spring Boot 2.1.3 + JeecgBoot 2.8.0) 迁移到 Go (GoFrame v2) 的技术架构和实现方案。

### 技术栈对照

| Java 技术 | Go 替代方案 |
|-----------|-------------|
| Spring Boot 2.1.3 | GoFrame v2 |
| MyBatis-Plus 3.1.2 | GoFrame ORM (gdb) |
| Apache Shiro + JWT | GoFrame JWT + 自定义中间件 |
| Druid 连接池 | GoFrame 内置连接池 |
| Redis (Lettuce) | go-redis/redis |
| Quartz | gocron 或 robfig/cron |
| FastJSON | encoding/json |
| Swagger | swaggo/swag |
| Lombok | Go struct tags |

---

## Architecture

### 项目结构

```
api-go/
├── api/                          # API 定义层
│   ├── middleware/               # 中间件
│   │   ├── auth.go              # JWT认证中间件
│   │   ├── cors.go              # CORS中间件
│   │   └── logger.go            # 日志中间件
│   └── v1/                       # API版本
│       ├── system/              # 系统模块API
│       └── teaching/            # 教学模块API
├── internal/                     # 内部实现
│   ├── cmd/                     # 启动命令
│   ├── consts/                  # 常量定义
│   ├── controller/              # 控制器层
│   │   ├── system/
│   │   └── teaching/
│   ├── dao/                     # 数据访问层
│   │   └── internal/
│   ├── logic/                   # 业务逻辑层
│   │   ├── system/
│   │   └── teaching/
│   ├── model/                   # 数据模型
│   │   ├── do/                  # 数据操作对象
│   │   ├── entity/              # 实体对象
│   │   └── vo/                  # 视图对象
│   └── service/                 # 服务接口
├── manifest/                     # 配置清单
│   └── config/
│       └── config.yaml
├── utility/                      # 工具类
│   ├── jwt/
│   ├── password/
│   ├── response/
│   └── redis/
├── go.mod
├── go.sum
└── main.go
```

### 分层架构

```
┌─────────────────────────────────────────────────────────────┐
│                      API Layer (api/)                        │
│  - 请求/响应结构定义                                          │
│  - 路由定义                                                   │
│  - 中间件                                                     │
├─────────────────────────────────────────────────────────────┤
│                  Controller Layer (controller/)              │
│  - 请求参数验证                                               │
│  - 调用Service层                                              │
│  - 响应格式化                                                 │
├─────────────────────────────────────────────────────────────┤
│                   Logic Layer (logic/)                       │
│  - 业务逻辑实现                                               │
│  - 事务管理                                                   │
│  - 缓存处理                                                   │
├─────────────────────────────────────────────────────────────┤
│                    DAO Layer (dao/)                          │
│  - 数据库操作                                                 │
│  - SQL构建                                                    │
├─────────────────────────────────────────────────────────────┤
│                   Model Layer (model/)                       │
│  - Entity: 数据库表映射                                       │
│  - DO: 数据操作对象                                           │
│  - VO: 视图对象                                               │
└─────────────────────────────────────────────────────────────┘
```

---

## Components and Interfaces

### 1. 认证模块 (Auth Module)

#### JWT Token 结构
```go
type JwtClaims struct {
    Username string `json:"username"`
    UserId   string `json:"userId"`
    jwt.RegisteredClaims
}
```

#### 认证中间件接口
```go
type IAuthMiddleware interface {
    // 验证JWT Token
    Auth(r *ghttp.Request)
    // 验证权限
    CheckPermission(r *ghttp.Request, perms string) bool
    // 获取当前用户
    GetCurrentUser(r *ghttp.Request) *entity.SysUser
}
```

#### 登录服务接口
```go
type ILoginService interface {
    // 用户名密码登录
    Login(ctx context.Context, username, password, captcha, checkKey string) (*vo.LoginResult, error)
    // 手机号登录
    PhoneLogin(ctx context.Context, phone, captcha string) (*vo.LoginResult, error)
    // 退出登录
    Logout(ctx context.Context, token string) error
    // 获取验证码
    GetCaptcha(ctx context.Context, key string) (string, error)
    // 发送短信验证码
    SendSms(ctx context.Context, phone, smsMode string) error
}
```

### 2. 用户模块 (User Module)

#### 用户服务接口
```go
type ISysUserService interface {
    // 根据用户名获取用户
    GetUserByName(ctx context.Context, username string) (*entity.SysUser, error)
    // 根据ID获取用户
    GetUserById(ctx context.Context, id string) (*entity.SysUser, error)
    // 根据手机号获取用户
    GetUserByPhone(ctx context.Context, phone string) (*entity.SysUser, error)
    // 分页查询用户
    GetUserList(ctx context.Context, req *vo.UserListReq) (*vo.PageResult, error)
    // 创建用户
    CreateUser(ctx context.Context, user *entity.SysUser, roles string) error
    // 更新用户
    UpdateUser(ctx context.Context, user *entity.SysUser, roles string) error
    // 删除用户
    DeleteUser(ctx context.Context, userId string) error
    // 批量删除用户
    DeleteBatchUsers(ctx context.Context, userIds string) error
    // 重置密码
    ResetPassword(ctx context.Context, username, oldPwd, newPwd, confirmPwd string) error
    // 获取用户角色集合
    GetUserRolesSet(ctx context.Context, username string) ([]string, error)
    // 获取用户权限集合
    GetUserPermissionsSet(ctx context.Context, username string) ([]string, error)
    // 检查用户有效性
    CheckUserIsEffective(ctx context.Context, user *entity.SysUser) error
    // 更新用户部门
    UpdateUserDepart(ctx context.Context, username, orgCode string) error
}
```

### 3. 角色权限模块 (RBAC Module)

#### 角色服务接口
```go
type ISysRoleService interface {
    GetRoleList(ctx context.Context, req *vo.RoleListReq) (*vo.PageResult, error)
    GetRoleById(ctx context.Context, id string) (*entity.SysRole, error)
    CreateRole(ctx context.Context, role *entity.SysRole) error
    UpdateRole(ctx context.Context, role *entity.SysRole) error
    DeleteRole(ctx context.Context, id string) error
    GetRolePermissions(ctx context.Context, roleId string) ([]string, error)
    SaveRolePermissions(ctx context.Context, roleId string, permissionIds []string) error
}
```

#### 权限服务接口
```go
type ISysPermissionService interface {
    GetPermissionTree(ctx context.Context) ([]*vo.PermissionTreeNode, error)
    GetPermissionById(ctx context.Context, id string) (*entity.SysPermission, error)
    CreatePermission(ctx context.Context, perm *entity.SysPermission) error
    UpdatePermission(ctx context.Context, perm *entity.SysPermission) error
    DeletePermission(ctx context.Context, id string) error
    GetUserMenus(ctx context.Context, userId string) ([]*vo.MenuTreeNode, error)
}
```

### 4. 部门模块 (Department Module)

#### 部门服务接口
```go
type ISysDepartService interface {
    GetDepartTree(ctx context.Context) ([]*vo.DepartTreeNode, error)
    GetDepartById(ctx context.Context, id string) (*entity.SysDepart, error)
    CreateDepart(ctx context.Context, depart *entity.SysDepart) error
    UpdateDepart(ctx context.Context, depart *entity.SysDepart) error
    DeleteDepart(ctx context.Context, id string) error
    QueryUserDeparts(ctx context.Context, userId string) ([]*entity.SysDepart, error)
    GetParentDepartIds(ctx context.Context, departId string) ([]string, error)
    GenerateOrgCode(ctx context.Context, parentId string) (string, error)
}
```

### 5. 课程模块 (Course Module)

#### 课程服务接口
```go
type ITeachingCourseService interface {
    GetCourseList(ctx context.Context, req *vo.CourseListReq) (*vo.PageResult, error)
    GetCourseById(ctx context.Context, id string) (*entity.TeachingCourse, error)
    GetHomeCourses(ctx context.Context, req *vo.HomeCourseReq) (*vo.PageResult, error)
    GetMineCourses(ctx context.Context, userId string) ([]*entity.TeachingCourse, error)
    CreateCourse(ctx context.Context, course *entity.TeachingCourse) error
    UpdateCourse(ctx context.Context, course *entity.TeachingCourse) error
    DeleteCourse(ctx context.Context, id string) error
    DeleteBatchCourses(ctx context.Context, ids string) error
}
```

#### 课程单元服务接口
```go
type ITeachingCourseUnitService interface {
    GetUnitsByCourseId(ctx context.Context, courseId string) ([]*entity.TeachingCourseUnit, error)
    GetUnitById(ctx context.Context, id string) (*entity.TeachingCourseUnit, error)
    CreateUnit(ctx context.Context, unit *entity.TeachingCourseUnit) error
    UpdateUnit(ctx context.Context, unit *entity.TeachingCourseUnit) error
    DeleteUnit(ctx context.Context, id string) error
}
```

### 6. 作业模块 (Work Module)

#### 作业服务接口
```go
type ITeachingWorkService interface {
    GetWorkList(ctx context.Context, req *vo.WorkListReq) (*vo.PageResult, error)
    GetWorkById(ctx context.Context, id string) (*entity.TeachingWork, error)
    GetMyWorks(ctx context.Context, userId string, req *vo.WorkListReq) (*vo.PageResult, error)
    CreateWork(ctx context.Context, work *entity.TeachingWork) error
    UpdateWork(ctx context.Context, work *entity.TeachingWork) error
    DeleteWork(ctx context.Context, id string) error
    CorrectWork(ctx context.Context, correct *entity.TeachingWorkCorrect) error
    CommentWork(ctx context.Context, comment *entity.TeachingWorkComment) error
    IncrementViewCount(ctx context.Context, workId string) error
    IncrementStarCount(ctx context.Context, workId string) error
}
```

### 7. 文件服务模块 (File Module)

#### 文件服务接口
```go
type ISysFileService interface {
    Upload(ctx context.Context, file *ghttp.UploadFile, fileTag string) (*entity.SysFile, error)
    GetFileByKey(ctx context.Context, fileKey string) (*entity.SysFile, error)
    DeleteFile(ctx context.Context, fileKey string) error
    GetFileUrl(ctx context.Context, fileKey string) (string, error)
}

type IStorageProvider interface {
    Upload(ctx context.Context, file *ghttp.UploadFile) (string, error)
    Delete(ctx context.Context, filePath string) error
    GetUrl(ctx context.Context, filePath string) (string, error)
}
```

### 8. 字典服务模块 (Dict Module)

#### 字典服务接口
```go
type ISysDictService interface {
    GetDictList(ctx context.Context, req *vo.DictListReq) (*vo.PageResult, error)
    GetDictById(ctx context.Context, id string) (*entity.SysDict, error)
    CreateDict(ctx context.Context, dict *entity.SysDict) error
    UpdateDict(ctx context.Context, dict *entity.SysDict) error
    DeleteDict(ctx context.Context, id string) error
    GetDictItems(ctx context.Context, dictCode string) ([]*entity.SysDictItem, error)
    QueryAllDictItems(ctx context.Context) (map[string][]*entity.SysDictItem, error)
    QueryTableDictItems(ctx context.Context, table, text, code string) ([]*vo.DictItem, error)
}
```

---

## Data Models

### 系统模块实体

#### SysUser (用户表)
```go
type SysUser struct {
    Id            string     `json:"id" orm:"id,primary"`
    Username      string     `json:"username" orm:"username"`
    Realname      string     `json:"realname" orm:"realname"`
    Password      string     `json:"-" orm:"password"`
    Salt          string     `json:"-" orm:"salt"`
    Avatar        string     `json:"avatar" orm:"avatar"`
    Birthday      *gtime.Time `json:"birthday" orm:"birthday"`
    Sex           int        `json:"sex" orm:"sex"`
    Email         string     `json:"email" orm:"email"`
    Phone         string     `json:"phone" orm:"phone"`
    OrgCode       string     `json:"orgCode" orm:"org_code"`
    Status        int        `json:"status" orm:"status"`
    DelFlag       int        `json:"delFlag" orm:"del_flag"`
    WorkNo        string     `json:"workNo" orm:"work_no"`
    Post          string     `json:"post" orm:"post"`
    School        string     `json:"school" orm:"school"`
    Telephone     string     `json:"telephone" orm:"telephone"`
    DepartIds     string     `json:"departIds" orm:"depart_ids"`
    ThirdId       string     `json:"thirdId" orm:"third_id"`
    ThirdType     string     `json:"thirdType" orm:"third_type"`
    UserIdentity  int        `json:"userIdentity" orm:"user_identity"`
    ActivitiSync  int        `json:"activitiSync" orm:"activiti_sync"`
    CreateBy      string     `json:"createBy" orm:"create_by"`
    CreateTime    *gtime.Time `json:"createTime" orm:"create_time"`
    UpdateBy      string     `json:"updateBy" orm:"update_by"`
    UpdateTime    *gtime.Time `json:"updateTime" orm:"update_time"`
}
```

#### SysRole (角色表)
```go
type SysRole struct {
    Id          string     `json:"id" orm:"id,primary"`
    RoleName    string     `json:"roleName" orm:"role_name"`
    RoleCode    string     `json:"roleCode" orm:"role_code"`
    RoleLevel   int        `json:"roleLevel" orm:"role_level"`
    Description string     `json:"description" orm:"description"`
    CreateBy    string     `json:"createBy" orm:"create_by"`
    CreateTime  *gtime.Time `json:"createTime" orm:"create_time"`
    UpdateBy    string     `json:"updateBy" orm:"update_by"`
    UpdateTime  *gtime.Time `json:"updateTime" orm:"update_time"`
}
```

#### SysPermission (权限表)
```go
type SysPermission struct {
    Id                 string     `json:"id" orm:"id,primary"`
    ParentId           string     `json:"parentId" orm:"parent_id"`
    Name               string     `json:"name" orm:"name"`
    Url                string     `json:"url" orm:"url"`
    Component          string     `json:"component" orm:"component"`
    ComponentName      string     `json:"componentName" orm:"component_name"`
    Redirect           string     `json:"redirect" orm:"redirect"`
    MenuType           int        `json:"menuType" orm:"menu_type"`
    Perms              string     `json:"perms" orm:"perms"`
    PermsType          string     `json:"permsType" orm:"perms_type"`
    SortNo             float64    `json:"sortNo" orm:"sort_no"`
    AlwaysShow         bool       `json:"alwaysShow" orm:"always_show"`
    Icon               string     `json:"icon" orm:"icon"`
    IsRoute            bool       `json:"isRoute" orm:"is_route"`
    IsLeaf             bool       `json:"isLeaf" orm:"is_leaf"`
    KeepAlive          bool       `json:"keepAlive" orm:"keep_alive"`
    Hidden             int        `json:"hidden" orm:"hidden"`
    Description        string     `json:"description" orm:"description"`
    DelFlag            int        `json:"delFlag" orm:"del_flag"`
    RuleFlag           int        `json:"ruleFlag" orm:"rule_flag"`
    Status             string     `json:"status" orm:"status"`
    InternalOrExternal bool       `json:"internalOrExternal" orm:"internal_or_external"`
    CreateBy           string     `json:"createBy" orm:"create_by"`
    CreateTime         *gtime.Time `json:"createTime" orm:"create_time"`
    UpdateBy           string     `json:"updateBy" orm:"update_by"`
    UpdateTime         *gtime.Time `json:"updateTime" orm:"update_time"`
}
```

#### SysDepart (部门表)
```go
type SysDepart struct {
    Id            string     `json:"id" orm:"id,primary"`
    ParentId      string     `json:"parentId" orm:"parent_id"`
    DepartName    string     `json:"departName" orm:"depart_name"`
    DepartNameEn  string     `json:"departNameEn" orm:"depart_name_en"`
    DepartNameAbbr string    `json:"departNameAbbr" orm:"depart_name_abbr"`
    DepartOrder   int        `json:"departOrder" orm:"depart_order"`
    Description   string     `json:"description" orm:"description"`
    OrgCategory   string     `json:"orgCategory" orm:"org_category"`
    OrgType       string     `json:"orgType" orm:"org_type"`
    OrgCode       string     `json:"orgCode" orm:"org_code"`
    Mobile        string     `json:"mobile" orm:"mobile"`
    Fax           string     `json:"fax" orm:"fax"`
    Address       string     `json:"address" orm:"address"`
    Memo          string     `json:"memo" orm:"memo"`
    Status        string     `json:"status" orm:"status"`
    DelFlag       string     `json:"delFlag" orm:"del_flag"`
    CreateBy      string     `json:"createBy" orm:"create_by"`
    CreateTime    *gtime.Time `json:"createTime" orm:"create_time"`
    UpdateBy      string     `json:"updateBy" orm:"update_by"`
    UpdateTime    *gtime.Time `json:"updateTime" orm:"update_time"`
}
```

### 教学模块实体

#### TeachingCourse (课程表)
```go
type TeachingCourse struct {
    Id             string     `json:"id" orm:"id,primary"`
    CourseName     string     `json:"courseName" orm:"course_name"`
    CourseDesc     string     `json:"courseDesc" orm:"course_desc"`
    CourseIcon     string     `json:"courseIcon" orm:"course_icon"`
    CourseCover    string     `json:"courseCover" orm:"course_cover"`
    CourseMap      string     `json:"courseMap" orm:"course_map"`
    CourseType     string     `json:"courseType" orm:"course_type"`
    CourseCategory string     `json:"courseCategory" orm:"course_category"`
    ShowHome       bool       `json:"showHome" orm:"show_home"`
    IsShared       bool       `json:"isShared" orm:"is_shared"`
    ShowType       int        `json:"showType" orm:"show_type"`
    DepartIds      string     `json:"departIds" orm:"depart_ids"`
    OrderNum       int        `json:"orderNum" orm:"order_num"`
    DelFlag        int        `json:"delFlag" orm:"del_flag"`
    SysOrgCode     string     `json:"sysOrgCode" orm:"sys_org_code"`
    CreateBy       string     `json:"createBy" orm:"create_by"`
    CreateTime     *gtime.Time `json:"createTime" orm:"create_time"`
    UpdateBy       string     `json:"updateBy" orm:"update_by"`
    UpdateTime     *gtime.Time `json:"updateTime" orm:"update_time"`
}
```

#### TeachingCourseUnit (课程单元表)
```go
type TeachingCourseUnit struct {
    Id               string     `json:"id" orm:"id,primary"`
    CourseId         string     `json:"courseId" orm:"course_id"`
    UnitName         string     `json:"unitName" orm:"unit_name"`
    UnitDesc         string     `json:"unitDesc" orm:"unit_desc"`
    CourseVideo      string     `json:"courseVideo" orm:"course_video"`
    CourseVideoSource int       `json:"courseVideoSource" orm:"course_video_source"`
    ShowCourseVideo  int        `json:"showCourseVideo" orm:"show_course_video"`
    CourseCase       string     `json:"courseCase" orm:"course_case"`
    ShowCourseCase   int        `json:"showCourseCase" orm:"show_course_case"`
    CoursePpt        string     `json:"coursePpt" orm:"course_ppt"`
    ShowCoursePpt    int        `json:"showCoursePpt" orm:"show_course_ppt"`
    CourseWorkType   int        `json:"courseWorkType" orm:"course_work_type"`
    CourseWork       string     `json:"courseWork" orm:"course_work"`
    CourseWorkAnswer string     `json:"courseWorkAnswer" orm:"course_work_answer"`
    CoursePlan       string     `json:"coursePlan" orm:"course_plan"`
    ShowCoursePlan   int        `json:"showCoursePlan" orm:"show_course_plan"`
    MapX             int        `json:"mapX" orm:"map_x"`
    MapY             int        `json:"mapY" orm:"map_y"`
    MediaContent     string     `json:"mediaContent" orm:"media_content"`
    OrderNum         int        `json:"orderNum" orm:"order_num"`
    CreateBy         string     `json:"createBy" orm:"create_by"`
    CreateTime       *gtime.Time `json:"createTime" orm:"create_time"`
    UpdateBy         string     `json:"updateBy" orm:"update_by"`
    UpdateTime       *gtime.Time `json:"updateTime" orm:"update_time"`
}
```

#### TeachingWork (作业表)
```go
type TeachingWork struct {
    Id           string     `json:"id" orm:"id,primary"`
    UserId       string     `json:"userId" orm:"user_id"`
    DepartId     string     `json:"departId" orm:"depart_id"`
    CourseId     string     `json:"courseId" orm:"course_id"`
    AdditionalId string     `json:"additionalId" orm:"additional_id"`
    WorkName     string     `json:"workName" orm:"work_name"`
    WorkType     string     `json:"workType" orm:"work_type"`
    WorkStatus   int        `json:"workStatus" orm:"work_status"`
    WorkFile     string     `json:"workFile" orm:"work_file"`
    WorkCover    string     `json:"workCover" orm:"work_cover"`
    ViewNum      int        `json:"viewNum" orm:"view_num"`
    StarNum      int        `json:"starNum" orm:"star_num"`
    CollectNum   int        `json:"collectNum" orm:"collect_num"`
    DelFlag      int        `json:"delFlag" orm:"del_flag"`
    WorkScene    string     `json:"workScene" orm:"work_scene"`
    HasCloudData bool       `json:"hasCloudData" orm:"has_cloud_data"`
    SysOrgCode   string     `json:"sysOrgCode" orm:"sys_org_code"`
    CreateBy     string     `json:"createBy" orm:"create_by"`
    CreateTime   *gtime.Time `json:"createTime" orm:"create_time"`
    UpdateBy     string     `json:"updateBy" orm:"update_by"`
    UpdateTime   *gtime.Time `json:"updateTime" orm:"update_time"`
}
```

---

## Correctness Properties

*A property is a characteristic or behavior that should hold true across all valid executions of a system-essentially, a formal statement about what the system should do. Properties serve as the bridge between human-readable specifications and machine-verifiable correctness guarantees.*



### Property 1: Password Encryption Round-Trip
*For any* valid password string and generated salt, encrypting the password with the salt and then comparing with the stored encrypted password SHALL produce a match.
**Validates: Requirements 1.8**

### Property 2: JWT Token Lifecycle
*For any* valid user credentials, logging in SHALL generate a token stored in Redis, and logging out SHALL remove that token from Redis.
**Validates: Requirements 1.1, 1.5**

### Property 3: Token Validation Consistency
*For any* expired or invalid JWT token, the authentication middleware SHALL return 401 Unauthorized status.
**Validates: Requirements 1.3**

### Property 4: User CRUD Integrity
*For any* user entity, creating then reading SHALL return the same data, and deleting SHALL set delFlag=1 without physical deletion.
**Validates: Requirements 2.1, 2.4**

### Property 5: User-Role-Permission Relationship Integrity
*For any* user assigned to roles, querying user permissions SHALL return the union of all permissions from assigned roles.
**Validates: Requirements 2.7, 3.3, 3.5**

### Property 6: Department Tree Structure Integrity
*For any* department with parentId, the parent department SHALL exist, and querying parent IDs SHALL return a valid chain to root.
**Validates: Requirements 4.2, 4.7**

### Property 7: OrgCode Uniqueness
*For any* two departments, their orgCode values SHALL be unique.
**Validates: Requirements 4.4**

### Property 8: Dictionary Items Ordering
*For any* dictionary code, querying items SHALL return results sorted by sortOrder in ascending order.
**Validates: Requirements 5.3**

### Property 9: Course Department Authorization
*For any* course with departIds, querying courses for a user SHALL only return courses where user's department is in departIds or departIds is empty.
**Validates: Requirements 6.3**

### Property 10: Work-Department Association
*For any* work submission, the work's departId SHALL match the submitting user's current department.
**Validates: Requirements 7.2**

### Property 11: Work Counter Monotonicity
*For any* work, incrementing view/star/collect count SHALL increase the respective counter by exactly 1.
**Validates: Requirements 7.8**

### Property 12: API Response Format Consistency
*For any* API endpoint, the response SHALL contain success, message, code, and result fields, with success=true and code=200 for successful operations.
**Validates: Requirements 15.1, 15.2, 15.3**

### Property 13: Pagination Response Completeness
*For any* paginated query, the response SHALL contain records, total, size, current, and pages fields with mathematically consistent values (pages = ceil(total/size)).
**Validates: Requirements 15.4**

### Property 14: Cache Invalidation on Permission Change
*For any* role permission modification, the affected users' permission cache in Redis SHALL be invalidated.
**Validates: Requirements 3.8**

### Property 15: Single-Device Login Enforcement
*For any* user with single-device login enabled, a new login SHALL invalidate all previous tokens for that user.
**Validates: Requirements 1.7**

---

## Error Handling

### 错误码定义

| 错误码 | 描述 | HTTP状态码 |
|--------|------|------------|
| 200 | 成功 | 200 |
| 400 | 请求参数错误 | 400 |
| 401 | 未授权/Token无效 | 401 |
| 403 | 权限不足 | 403 |
| 404 | 资源不存在 | 404 |
| 500 | 服务器内部错误 | 500 |
| 10001 | 用户名或密码错误 | 200 |
| 10002 | 验证码错误 | 200 |
| 10003 | 账号已被锁定 | 200 |
| 10004 | Token已过期 | 200 |
| 10005 | 用户不存在 | 200 |
| 10006 | 手机号已注册 | 200 |
| 10007 | 短信发送失败 | 200 |

### 统一响应结构

```go
type Response struct {
    Success   bool        `json:"success"`
    Message   string      `json:"message"`
    Code      int         `json:"code"`
    Result    interface{} `json:"result"`
    Timestamp int64       `json:"timestamp"`
}

type PageResult struct {
    Records []interface{} `json:"records"`
    Total   int64         `json:"total"`
    Size    int           `json:"size"`
    Current int           `json:"current"`
    Pages   int           `json:"pages"`
}
```

### 错误处理中间件

```go
func ErrorHandler(r *ghttp.Request) {
    r.Middleware.Next()
    
    if err := r.GetError(); err != nil {
        r.Response.ClearBuffer()
        
        var code int
        var message string
        
        switch e := err.(type) {
        case *gerror.Error:
            code = e.Code().Code()
            message = e.Message()
        default:
            code = 500
            message = "服务器内部错误"
        }
        
        r.Response.WriteJson(Response{
            Success:   false,
            Message:   message,
            Code:      code,
            Timestamp: gtime.Now().Unix(),
        })
    }
}
```

---

## Testing Strategy

### 测试框架选择

- **单元测试**: Go标准库 `testing` + `testify/assert`
- **属性测试**: `gopter` (Go Property Testing)
- **集成测试**: GoFrame 内置测试支持
- **API测试**: `httptest` + GoFrame

### 测试配置

```go
// 属性测试配置
const (
    PropertyTestIterations = 100  // 每个属性测试最少100次迭代
)
```

### 测试目录结构

```
api-go/
├── internal/
│   ├── logic/
│   │   ├── system/
│   │   │   ├── user_test.go
│   │   │   ├── role_test.go
│   │   │   └── permission_test.go
│   │   └── teaching/
│   │       ├── course_test.go
│   │       └── work_test.go
│   └── service/
│       └── *_test.go
├── utility/
│   ├── jwt/
│   │   └── jwt_test.go
│   └── password/
│       └── password_test.go
└── test/
    ├── integration/
    └── property/
```

### 单元测试示例

```go
func TestPasswordEncrypt(t *testing.T) {
    username := "testuser"
    password := "testpass123"
    salt := GenerateSalt()
    
    encrypted := Encrypt(username, password, salt)
    
    assert.NotEmpty(t, encrypted)
    assert.NotEqual(t, password, encrypted)
    
    // 验证相同输入产生相同输出
    encrypted2 := Encrypt(username, password, salt)
    assert.Equal(t, encrypted, encrypted2)
}
```

### 属性测试示例

```go
// Feature: java-to-go-migration, Property 1: Password Encryption Round-Trip
func TestPasswordEncryptionRoundTrip(t *testing.T) {
    properties := gopter.NewProperties(gopter.DefaultTestParameters())
    properties.Property("password encryption is deterministic", prop.ForAll(
        func(password string) bool {
            if len(password) == 0 {
                return true // skip empty passwords
            }
            salt := GenerateSalt()
            encrypted1 := Encrypt("user", password, salt)
            encrypted2 := Encrypt("user", password, salt)
            return encrypted1 == encrypted2
        },
        gen.AnyString(),
    ))
    properties.TestingRun(t)
}
```

---

## API Endpoints Migration Reference

### 系统模块 API

| Java Endpoint | Go Endpoint | 方法 | 描述 |
|---------------|-------------|------|------|
| /sys/login | /api/sys/login | POST | 用户登录 |
| /sys/logout | /api/sys/logout | POST | 用户退出 |
| /sys/randomImage/{key} | /api/sys/randomImage/{key} | GET | 获取验证码 |
| /sys/sms | /api/sys/sms | POST | 发送短信 |
| /sys/phoneLogin | /api/sys/phoneLogin | POST | 手机号登录 |
| /sys/user/list | /api/sys/user/list | GET | 用户列表 |
| /sys/user/add | /api/sys/user/add | POST | 添加用户 |
| /sys/user/edit | /api/sys/user/edit | PUT | 编辑用户 |
| /sys/user/delete | /api/sys/user/delete | DELETE | 删除用户 |
| /sys/role/list | /api/sys/role/list | GET | 角色列表 |
| /sys/permission/list | /api/sys/permission/list | GET | 权限列表 |
| /sys/depart/queryTreeList | /api/sys/depart/queryTreeList | GET | 部门树 |
| /sys/dict/list | /api/sys/dict/list | GET | 字典列表 |
| /sys/dict/getDictItems/{dictCode} | /api/sys/dict/getDictItems/{dictCode} | GET | 字典项 |

### 教学模块 API

| Java Endpoint | Go Endpoint | 方法 | 描述 |
|---------------|-------------|------|------|
| /teaching/teachingCourse/list | /api/teaching/course/list | GET | 课程列表 |
| /teaching/teachingCourse/add | /api/teaching/course/add | POST | 添加课程 |
| /teaching/teachingCourse/edit | /api/teaching/course/edit | PUT | 编辑课程 |
| /teaching/teachingCourse/delete | /api/teaching/course/delete | DELETE | 删除课程 |
| /teaching/teachingCourse/mineCourse | /api/teaching/course/mine | GET | 我的课程 |
| /teaching/teachingCourse/getHomeCourse | /api/teaching/course/home | GET | 首页课程 |
| /teaching/teachingCourseUnit/list | /api/teaching/unit/list | GET | 课程单元列表 |
| /teaching/teachingWork/list | /api/teaching/work/list | GET | 作业列表 |
| /teaching/teachingWork/add | /api/teaching/work/add | POST | 提交作业 |
| /teaching/teachingWork/correct | /api/teaching/work/correct | POST | 批改作业 |

---

## Database Migration Notes

### 表结构兼容性

Go版本将完全兼容现有MySQL数据库表结构，无需进行数据库迁移。

### 主键生成策略

使用GoFrame的UUID生成器替代MyBatis-Plus的ID_WORKER_STR：

```go
import "github.com/gogf/gf/v2/util/guid"

func GenerateId() string {
    return guid.S()
}
```

### 时间字段处理

```go
import "github.com/gogf/gf/v2/os/gtime"

// 使用 *gtime.Time 替代 java.util.Date
type Entity struct {
    CreateTime *gtime.Time `json:"createTime" orm:"create_time"`
    UpdateTime *gtime.Time `json:"updateTime" orm:"update_time"`
}
```

---

## Configuration Migration

### Java配置 (application-dev.yml) → Go配置 (config.yaml)

```yaml
server:
  address: ":8081"
  serverRoot: "/api"

database:
  default:
    type: "mysql"
    host: "127.0.0.1"
    port: "3306"
    user: "teachingopen"
    pass: "teachingopen"
    name: "teachingopen"
    charset: "utf8mb4"
    maxIdle: 10
    maxOpen: 100
    maxLifetime: "30s"

redis:
  default:
    address: "127.0.0.1:6379"
    db: 1
    pass: ""

jwt:
  secret: "your-jwt-secret-key"
  expireTime: 7200  # 2 hours in seconds

upload:
  type: "qiniu"  # local, qiniu, aliyun, minio
  path: "/opt/upFiles"

qiniu:
  accessKey: "your-access-key"
  secretKey: "your-secret-key"
  bucket: "teaching-open"
  domain: "http://open.qn.teaching.vip"
  zone: "z0"

logger:
  level: "all"
  stdout: true
```
