# Implementation Plan: Java to Go Migration

## Overview

本实现计划将 Teaching Open 教学管理平台从 Java (Spring Boot + JeecgBoot) 迁移到 Go (GoFrame v2)。采用分阶段迁移策略，优先实现核心功能模块，确保每个阶段都可独立测试和验证。

## Tasks

- [x] 1. 项目基础架构搭建
  - [x] 1.1 初始化GoFrame项目结构
    - 创建标准目录结构 (api, internal, manifest, utility)
    - 配置go.mod依赖
    - 创建main.go入口文件
    - _Requirements: 18.1_

  - [x] 1.2 配置管理模块
    - 创建config.yaml配置文件
    - 实现数据库连接配置
    - 实现Redis连接配置
    - 实现JWT配置
    - 实现文件存储配置
    - _Requirements: 18.1, 19.1_

  - [x] 1.3 统一响应格式实现
    - 创建utility/response/response.go
    - 实现Response结构体
    - 实现PageResult结构体
    - 实现Success/Error辅助函数
    - _Requirements: 15.1, 15.2, 15.3, 15.4_

  - [ ] 1.4 编写响应格式属性测试
    - **Property 12: API Response Format Consistency**
    - **Property 13: Pagination Response Completeness**
    - **Validates: Requirements 15.1, 15.2, 15.3, 15.4**

- [x] 2. 认证授权模块
  - [x] 2.1 JWT工具类实现
    - 创建utility/jwt/jwt.go
    - 实现Token生成 (Sign)
    - 实现Token验证 (Verify)
    - 实现Token解析获取用户名 (GetUsername)
    - _Requirements: 1.1, 1.3_

  - [x] 2.2 密码加密工具实现
    - 创建utility/password/password.go
    - 实现MD5+Salt加密 (Encrypt)
    - 实现Salt生成 (GenerateSalt)
    - _Requirements: 1.8_

  - [ ] 2.3 编写密码加密属性测试
    - **Property 1: Password Encryption Round-Trip**
    - **Validates: Requirements 1.8**

  - [x] 2.4 Redis缓存工具实现
    - 创建utility/redis/redis.go
    - 实现Set/Get/Del操作
    - 实现Expire设置
    - 实现Token缓存操作
    - _Requirements: 19.1, 19.2_

  - [x] 2.5 认证中间件实现
    - 创建api/middleware/auth.go
    - 实现JWT Token验证
    - 实现Token刷新逻辑
    - 实现获取当前用户
    - _Requirements: 1.3, 1.4_

  - [ ] 2.6 编写JWT生命周期属性测试
    - **Property 2: JWT Token Lifecycle**
    - **Property 3: Token Validation Consistency**
    - **Validates: Requirements 1.1, 1.3, 1.5**

- [ ] 3. Checkpoint - 基础架构验证
  - 确保所有测试通过，如有问题请询问用户

- [x] 4. 系统用户模块
  - [x] 4.1 用户实体定义
    - 创建internal/model/entity/sys_user.go
    - 定义SysUser结构体及所有字段
    - 创建internal/model/do/sys_user.go
    - 创建internal/model/vo/sys_user.go
    - _Requirements: 2.1_

  - [x] 4.2 用户DAO层实现
    - 创建internal/dao/internal/sys_user.go
    - 创建internal/dao/sys_user.go
    - 实现基础CRUD操作
    - _Requirements: 2.1_

  - [x] 4.3 用户Service层实现
    - 创建internal/service/sys_user.go
    - 实现ISysUserService接口
    - 实现GetUserByName, GetUserById, GetUserByPhone
    - 实现CreateUser, UpdateUser, DeleteUser
    - 实现GetUserRolesSet, GetUserPermissionsSet
    - _Requirements: 2.1, 2.2, 2.3, 2.4, 2.6, 2.7, 2.8_

  - [x] 4.4 用户Logic层实现
    - 创建internal/logic/system/user.go
    - 实现用户业务逻辑
    - 实现密码加密逻辑
    - 实现用户有效性检查
    - _Requirements: 2.1, 2.2_

  - [ ] 4.5 编写用户CRUD属性测试
    - **Property 4: User CRUD Integrity**
    - **Validates: Requirements 2.1, 2.4**

  - [x] 4.6 用户Controller层实现
    - 创建internal/controller/system/sys_user.go
    - 实现用户列表查询
    - 实现用户增删改查
    - 实现用户批量操作
    - _Requirements: 2.1, 2.3, 2.4, 2.5_

  - [x] 4.7 用户API路由定义
    - 创建api/v1/system/user.go
    - 定义请求/响应结构
    - 注册路由
    - _Requirements: 2.1_

- [x] 5. 登录模块
  - [x] 5.1 登录Service实现
    - 创建internal/service/login.go
    - 实现ILoginService接口
    - 实现Login方法
    - 实现PhoneLogin方法
    - 实现Logout方法
    - 实现GetCaptcha方法
    - _Requirements: 1.1, 1.5, 1.6_

  - [x] 5.2 登录Controller实现
    - 创建internal/controller/system/login.go
    - 实现登录接口
    - 实现退出接口
    - 实现验证码接口
    - 实现短信发送接口
    - _Requirements: 1.1, 1.2, 1.5, 1.6_

  - [ ] 5.3 编写单设备登录属性测试
    - **Property 15: Single-Device Login Enforcement**
    - **Validates: Requirements 1.7**

- [ ] 6. Checkpoint - 用户认证模块验证
  - 确保所有测试通过，如有问题请询问用户

- [x] 7. 角色权限模块
  - [x] 7.1 角色实体定义
    - 创建internal/model/entity/sys_role.go
    - 创建internal/model/entity/sys_permission.go
    - 创建internal/model/entity/sys_role_permission.go
    - 创建internal/model/entity/sys_user_role.go
    - _Requirements: 3.1, 3.2_

  - [x] 7.2 角色DAO层实现
    - 创建internal/dao/sys_role.go
    - 创建internal/dao/sys_permission.go
    - 创建internal/dao/sys_role_permission.go
    - 创建internal/dao/sys_user_role.go
    - _Requirements: 3.1, 3.2, 3.3_

  - [x] 7.3 角色Service层实现
    - 创建internal/service/sys_role.go
    - 实现ISysRoleService接口
    - 实现角色CRUD
    - 实现角色权限分配
    - _Requirements: 3.1, 3.3_

  - [x] 7.4 权限Service层实现
    - 创建internal/service/sys_permission.go
    - 实现ISysPermissionService接口
    - 实现权限树查询
    - 实现用户菜单查询
    - _Requirements: 3.2, 3.4, 3.5_

  - [ ] 7.5 编写角色权限关系属性测试
    - **Property 5: User-Role-Permission Relationship Integrity**
    - **Validates: Requirements 2.7, 3.3, 3.5**

  - [ ] 7.6 权限缓存实现
    - 实现用户权限Redis缓存
    - 实现缓存失效逻辑
    - _Requirements: 3.7, 3.8_

  - [ ] 7.7 编写缓存失效属性测试
    - **Property 14: Cache Invalidation on Permission Change**
    - **Validates: Requirements 3.8**

  - [x] 7.8 角色权限Controller实现
    - 创建internal/controller/system/sys_role.go
    - 创建internal/controller/system/sys_permission.go
    - 实现角色管理接口
    - 实现权限管理接口
    - _Requirements: 3.1, 3.2, 3.3, 3.4_

- [x] 8. 部门模块
  - [x] 8.1 部门实体定义
    - 创建internal/model/entity/sys_depart.go
    - 创建internal/model/entity/sys_user_depart.go
    - 创建internal/model/vo/depart_tree.go
    - _Requirements: 4.1_

  - [x] 8.2 部门DAO层实现
    - 创建internal/dao/sys_depart.go
    - 创建internal/dao/sys_user_depart.go
    - _Requirements: 4.1_

  - [x] 8.3 部门Service层实现
    - 创建internal/service/sys_depart.go
    - 实现ISysDepartService接口
    - 实现部门树查询
    - 实现OrgCode生成
    - 实现父部门ID查询
    - _Requirements: 4.1, 4.2, 4.3, 4.4, 4.5, 4.7_

  - [ ] 8.4 编写部门树结构属性测试
    - **Property 6: Department Tree Structure Integrity**
    - **Property 7: OrgCode Uniqueness**
    - **Validates: Requirements 4.2, 4.4, 4.7**

  - [x] 8.5 部门Controller实现
    - 创建internal/controller/system/sys_depart.go
    - 实现部门树查询接口
    - 实现部门CRUD接口
    - _Requirements: 4.1, 4.2, 4.3_

- [x] 9. 数据字典模块
  - [x] 9.1 字典实体定义
    - 创建internal/model/entity/sys_dict.go
    - 创建internal/model/entity/sys_dict_item.go
    - _Requirements: 5.1, 5.2_

  - [x] 9.2 字典DAO层实现
    - 创建internal/dao/sys_dict.go
    - 创建internal/dao/sys_dict_item.go
    - _Requirements: 5.1, 5.2_

  - [x] 9.3 字典Service层实现
    - 创建internal/service/sys_dict.go
    - 实现ISysDictService接口
    - 实现字典项查询(按sortOrder排序)
    - 实现表字典查询
    - 实现字典缓存
    - _Requirements: 5.1, 5.2, 5.3, 5.4, 5.5, 5.6_

  - [ ] 9.4 编写字典排序属性测试
    - **Property 8: Dictionary Items Ordering**
    - **Validates: Requirements 5.3**

  - [ ] 9.5 字典Controller实现
    - 创建internal/controller/system/sys_dict.go
    - 实现字典管理接口
    - 实现字典项查询接口
    - _Requirements: 5.1, 5.2, 5.3, 5.4_

- [ ] 10. Checkpoint - 系统模块验证
  - 确保所有测试通过，如有问题请询问用户

- [ ] 11. 课程模块
  - [ ] 11.1 课程实体定义
    - 创建internal/model/entity/teaching_course.go
    - 创建internal/model/entity/teaching_course_unit.go
    - 创建internal/model/entity/teaching_course_dept.go
    - _Requirements: 6.1, 6.2_

  - [ ] 11.2 课程DAO层实现
    - 创建internal/dao/teaching_course.go
    - 创建internal/dao/teaching_course_unit.go
    - 创建internal/dao/teaching_course_dept.go
    - _Requirements: 6.1, 6.2, 6.5_

  - [ ] 11.3 课程Service层实现
    - 创建internal/service/teaching_course.go
    - 实现ITeachingCourseService接口
    - 实现课程列表查询(部门授权过滤)
    - 实现首页课程查询
    - 实现我的课程查询
    - 实现课程CRUD
    - _Requirements: 6.1, 6.3, 6.4, 6.5, 6.6, 6.7_

  - [ ] 11.4 编写课程授权属性测试
    - **Property 9: Course Department Authorization**
    - **Validates: Requirements 6.3**

  - [ ] 11.5 课程单元Service实现
    - 创建internal/service/teaching_course_unit.go
    - 实现ITeachingCourseUnitService接口
    - _Requirements: 6.2_

  - [ ] 11.6 课程Controller实现
    - 创建internal/controller/teaching/teaching_course.go
    - 创建internal/controller/teaching/teaching_course_unit.go
    - 实现课程管理接口
    - 实现课程单元管理接口
    - _Requirements: 6.1, 6.2_

- [ ] 12. 作业模块
  - [ ] 12.1 作业实体定义
    - 创建internal/model/entity/teaching_work.go
    - 创建internal/model/entity/teaching_work_comment.go
    - 创建internal/model/entity/teaching_work_correct.go
    - 创建internal/model/entity/teaching_additional_work.go
    - _Requirements: 7.1, 8.1_

  - [ ] 12.2 作业DAO层实现
    - 创建internal/dao/teaching_work.go
    - 创建internal/dao/teaching_work_comment.go
    - 创建internal/dao/teaching_work_correct.go
    - 创建internal/dao/teaching_additional_work.go
    - _Requirements: 7.1, 7.5, 7.6, 8.1_

  - [ ] 12.3 作业Service层实现
    - 创建internal/service/teaching_work.go
    - 实现ITeachingWorkService接口
    - 实现作业提交(关联部门)
    - 实现作业查询(多条件过滤)
    - 实现作业批改
    - 实现作业评论
    - 实现计数器增加
    - _Requirements: 7.1, 7.2, 7.3, 7.4, 7.5, 7.6, 7.7, 7.8_

  - [ ] 12.4 编写作业部门关联属性测试
    - **Property 10: Work-Department Association**
    - **Validates: Requirements 7.2**

  - [ ] 12.5 编写作业计数器属性测试
    - **Property 11: Work Counter Monotonicity**
    - **Validates: Requirements 7.8**

  - [ ] 12.6 作业Controller实现
    - 创建internal/controller/teaching/teaching_work.go
    - 实现作业管理接口
    - 实现作业批改接口
    - 实现作业评论接口
    - _Requirements: 7.1, 7.5, 7.6, 7.7_

- [ ] 13. Checkpoint - 教学模块验证
  - 确保所有测试通过，如有问题请询问用户

- [ ] 14. 文件存储模块
  - [ ] 14.1 文件实体定义
    - 创建internal/model/entity/sys_file.go
    - _Requirements: 9.1_

  - [ ] 14.2 存储提供者实现
    - 创建utility/storage/storage.go (接口定义)
    - 创建utility/storage/local.go (本地存储)
    - 创建utility/storage/qiniu.go (七牛云)
    - 创建utility/storage/aliyun.go (阿里云OSS)
    - _Requirements: 9.2_

  - [ ] 14.3 文件Service实现
    - 创建internal/service/sys_file.go
    - 实现ISysFileService接口
    - 实现文件上传
    - 实现文件删除
    - 实现文件URL获取
    - _Requirements: 9.1, 9.2, 9.3, 9.4, 9.5, 9.6_

  - [ ] 14.4 文件Controller实现
    - 创建internal/controller/system/sys_file.go
    - 实现文件上传接口
    - 实现文件访问接口
    - _Requirements: 9.1, 9.3_

- [ ] 15. 系统日志模块
  - [ ] 15.1 日志实体定义
    - 创建internal/model/entity/sys_log.go
    - _Requirements: 10.1_

  - [ ] 15.2 日志DAO层实现
    - 创建internal/dao/sys_log.go
    - _Requirements: 10.1_

  - [ ] 15.3 日志Service实现
    - 创建internal/service/sys_log.go
    - 实现日志记录
    - 实现日志查询
    - _Requirements: 10.1, 10.2, 10.4, 10.5_

  - [ ] 15.4 日志中间件实现
    - 创建api/middleware/logger.go
    - 实现自动操作日志记录
    - 实现请求耗时记录
    - _Requirements: 10.3, 10.5_

  - [ ] 15.5 日志Controller实现
    - 创建internal/controller/system/sys_log.go
    - 实现日志查询接口
    - _Requirements: 10.4_

- [ ] 16. 前端菜单模块
  - [ ] 16.1 菜单实体定义
    - 创建internal/model/entity/teaching_menu.go
    - _Requirements: 12.1_

  - [ ] 16.2 菜单Service实现
    - 创建internal/service/teaching_menu.go
    - 实现菜单树查询
    - _Requirements: 12.1, 12.2, 12.3_

  - [ ] 16.3 菜单Controller实现
    - 创建internal/controller/teaching/teaching_menu.go
    - 实现菜单查询接口
    - _Requirements: 12.1, 12.2_

- [ ] 17. 其他辅助模块
  - [ ] 17.1 班级日志模块
    - 创建internal/model/entity/teaching_depart_day_log.go
    - 创建internal/service/teaching_depart_day_log.go
    - 创建internal/controller/teaching/teaching_depart_day_log.go
    - _Requirements: 13.1, 13.2, 13.3_

  - [ ] 17.2 Scratch素材模块
    - 创建internal/model/entity/teaching_scratch_assets.go
    - 创建internal/service/teaching_scratch_assets.go
    - 创建internal/controller/teaching/teaching_scratch_assets.go
    - _Requirements: 14.1, 14.2, 14.3, 14.4_

  - [ ] 17.3 新闻公告模块
    - 创建internal/model/entity/teaching_news.go
    - 创建internal/service/teaching_news.go
    - 创建internal/controller/teaching/teaching_news.go
    - _Requirements: 11.1, 11.2_

- [ ] 18. 路由整合与API文档
  - [ ] 18.1 路由注册整合
    - 创建internal/cmd/cmd.go
    - 注册所有Controller路由
    - 配置中间件链
    - _Requirements: 15.1_

  - [ ] 18.2 Swagger文档生成
    - 添加swaggo注解
    - 生成API文档
    - _Requirements: 15.1_

- [ ] 19. Final Checkpoint - 完整系统验证
  - 确保所有测试通过
  - 验证所有API端点可用
  - 验证与前端的兼容性
  - 如有问题请询问用户

## Notes

- 所有任务均为必需，包括属性测试任务
- 每个Checkpoint确保当前阶段功能完整可用
- 属性测试使用gopter库，每个测试至少100次迭代
- 所有API保持与Java版本兼容，确保前端无需修改
- 数据库表结构完全兼容，无需数据迁移
