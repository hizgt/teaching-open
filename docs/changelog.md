# Teaching Open 项目变更日志

## [3.0.0-dev] - 2025-12-22

### 文档分析与报告生成

#### Added
- 分析了docs目录内容
- 梳理了未开发PRD
- 生成了未完成工作报告到docs/20251222/目录

---

## [3.0.0-dev] - 2025-12-14

### 第十二阶段：部门日志统计 ✅

#### Added
- **部门日志统计相关 DAO 层** ✅
  - 创建 teaching_depart_day_log 表的 Entity、DO、DAO 文件
  - Entity: `internal/model/entity/teaching_depart_day_log.go`
  - DO: `internal/model/do/teaching_depart_day_log.go`
  - Internal DAO: `internal/dao/internal/teaching_depart_day_log.go`
  - DAO: `internal/dao/teaching_depart_day_log.go`

- **部门日志统计接口 (4个)** ✅
  - `GET /api/v1/teaching/teachingDepartDayLog/getReport` - 统计报表（分页、日期筛选、排序）
  - `GET /api/v1/teaching/teachingDepartDayLog/getReportGroupByDepart` - 按部门统计聚合
  - `GET /api/v1/teaching/teachingDepartDayLog/getReportGroupByMonth` - 按月份统计聚合
  - `POST /api/v1/teaching/teachingDepartDayLog/unitViewLog` - 记录单元查看日志

#### 技术实现细节
- 统计维度：班级、日期范围、月份
- 统计指标：开课次数、作业布置/批改/提交次数
- 聚合查询：使用SQL聚合函数进行高效统计
- 分页支持：报表和部门统计均支持分页
- 日志记录：自动记录每天各班级的教学活动数据

### 准备工作

#### Git提交Phase 12代码
```bash
git add -A
git commit -m "Phase 12: 部门日志统计完成"
git push origin devgo
```

#### 第十三阶段:前端Vue3重构 (准备就绪)
- ✅ 初始化脚本: `init-phase13-vue3.sh`
- ✅ 配置示例: Vite、TypeScript、Axios、Pinia、Router
- ✅ 快速启动指南: `PHASE13_QUICK_START.md`
- ⏳ 待执行: `bash init-phase13-vue3.sh`

---

## [3.0.0-dev] - 2025-12-08

### 第十一阶段：Scratch素材库管理 ✅

#### Added
- **Scratch素材相关 DAO 层** ✅
  - 创建 teaching_scratch_assets 表的 Entity、DO、DAO 文件
  - Entity: `internal/model/entity/teaching_scratch_asset.go`
  - DO: `internal/model/do/teaching_scratch_asset.go`
  - Internal DAO: `internal/dao/internal/teaching_scratch_asset.go`
  - DAO: `internal/dao/teaching_scratch_asset.go`

- **Scratch素材管理接口 (7个)** ✅
  - `GET /api/v1/teaching/teachingScratchAssets/list` - 素材列表（分页、筛选）
  - `GET /api/v1/teaching/teachingScratchAssets/getScratchAssets` - 获取素材（Scratch编辑器用）
  - `POST /api/v1/teaching/teachingScratchAssets/add` - 添加素材
  - `PUT /api/v1/teaching/teachingScratchAssets/edit` - 编辑素材
  - `DELETE /api/v1/teaching/teachingScratchAssets/delete` - 删除素材
  - `DELETE /api/v1/teaching/teachingScratchAssets/deleteBatch` - 批量删除素材
  - `GET /api/v1/teaching/teachingScratchAssets/queryById` - 素材详情

#### 技术实现细节
- 素材类型：1背景 2声音 3造型 4角色
- 软删除：del_flag 字段标记删除状态
- JSON数据：asset_data 存储Scratch素材的JSON结构
- MD5标识：md5_ext 用于素材资源定位
- 标签系统：tags 字段支持多标签

### 下一步计划

#### 第十二阶段：部门日志统计
- [ ] 实现部门日志统计接口（teaching_depart_day_log）
- [ ] 支持按日期范围统计部门数据

---

## [3.0.0-dev] - 2025-12-07

### 第十阶段：新闻公告与附加作业 ✅

#### Added
- **新闻公告相关 DAO 层** ✅
  - 创建 teaching_news 表的 Entity、DO、DAO 文件

- **附加作业相关 DAO 层** ✅
  - 创建 teaching_additional_work 表的 Entity、DO、DAO 文件

- **新闻公告管理接口 (9个)** ✅
  - `GET /api/v1/teaching/teachingNews/list` - 新闻列表（分页、筛选）
  - `POST /api/v1/teaching/teachingNews/add` - 添加新闻
  - `PUT /api/v1/teaching/teachingNews/edit` - 编辑新闻
  - `DELETE /api/v1/teaching/teachingNews/delete` - 删除新闻
  - `DELETE /api/v1/teaching/teachingNews/deleteBatch` - 批量删除新闻
  - `GET /api/v1/teaching/teachingNews/queryById` - 新闻详情
  - `PUT /api/v1/teaching/teachingNews/publish` - 发布新闻
  - `PUT /api/v1/teaching/teachingNews/offline` - 下架新闻
  - `GET /api/v1/teaching/teachingNews/publicList` - 公开新闻列表（无需登录）

- **附加作业管理接口 (9个)** ✅
  - `GET /api/v1/teaching/teachingAdditionalWork/list` - 附加作业列表（分页、筛选）
  - `POST /api/v1/teaching/teachingAdditionalWork/add` - 添加附加作业
  - `PUT /api/v1/teaching/teachingAdditionalWork/edit` - 编辑附加作业
  - `DELETE /api/v1/teaching/teachingAdditionalWork/delete` - 删除附加作业
  - `DELETE /api/v1/teaching/teachingAdditionalWork/deleteBatch` - 批量删除附加作业
  - `GET /api/v1/teaching/teachingAdditionalWork/queryById` - 附加作业详情
  - `PUT /api/v1/teaching/teachingAdditionalWork/publish` - 发布附加作业
  - `PUT /api/v1/teaching/teachingAdditionalWork/offline` - 下架附加作业
  - `GET /api/v1/teaching/teachingAdditionalWork/listByDept` - 按班级获取附加作业

#### 技术实现细节
- 新闻状态：草稿(0)/已发布(1)
- 附加作业状态：未发布(0)/已发布(1)
- 代码类型支持：Scratch/Python/Blockly等
- 班级分配：work_dept字段支持多班级逗号分隔

#### 第十一阶段：Scratch素材库管理
- [ ] 实现Scratch素材库管理接口（teaching_scratch_asset）
- [ ] 支持自定义角色、背景、音效素材

---

## [3.0.0-dev] - 2025-12-06

### 第九阶段：教学管理-作品模块 ✅

#### Added
- **作品相关 DAO 层** ✅
  - 创建 teaching_work 表的 Entity、DO、DAO 文件
  - 创建 teaching_work_comment 表的 Entity、DO、DAO 文件
  - 创建 teaching_work_correct 表的 Entity、DO、DAO 文件

- **作品管理接口 (15个)** ✅
  - `GET /api/v1/teaching/teachingWork/list` - 作品列表（分页、筛选）
  - `GET /api/v1/teaching/teachingWork/mine` - 我的作品
  - `GET /api/v1/teaching/teachingWork/greatWork` - 优秀作品
  - `GET /api/v1/teaching/teachingWork/starWork` - 收藏作品
  - `GET /api/v1/teaching/teachingWork/leaderboard` - 作品排行榜
  - `POST /api/v1/teaching/teachingWork/add` - 添加作品
  - `PUT /api/v1/teaching/teachingWork/edit` - 编辑作品
  - `DELETE /api/v1/teaching/teachingWork/delete` - 删除作品
  - `DELETE /api/v1/teaching/teachingWork/deleteBatch` - 批量删除作品
  - `GET /api/v1/teaching/teachingWork/queryById` - 作品详情
  - `POST /api/v1/teaching/teachingWork/submit` - 提交作品
  - `GET /api/v1/teaching/teachingWork/studentWorkInfo` - 学生作品信息
  - `POST /api/v1/teaching/teachingWork/sendWork` - 发送作品给其他用户
  - `GET /api/v1/teaching/teachingWork/mineAdditionalWork` - 我的附加作业作品
  - `POST /api/v1/teaching/teachingWork/star` - 点赞/取消点赞
  - `POST /api/v1/teaching/teachingWork/collect` - 收藏/取消收藏

- **作品批改接口 (2个)** ✅
  - `GET /api/v1/teaching/teachingWork/queryTeachingWorkCorrectByMainId` - 批改记录列表
  - `POST /api/v1/teaching/teachingWork/correct` - 批改作品

- **作品评论接口 (3个)** ✅
  - `GET /api/v1/teaching/teachingWork/getWorkComments` - 获取评论列表
  - `POST /api/v1/teaching/teachingWork/saveComment` - 添加评论
  - `DELETE /api/v1/teaching/teachingWork/deleteComment` - 删除评论

- **作品标签接口 (3个)** ✅
  - `GET /api/v1/teaching/teachingWork/getWorkTags` - 获取作品标签
  - `POST /api/v1/teaching/teachingWork/setWorkTag` - 设置作品标签
  - `DELETE /api/v1/teaching/teachingWork/delWorkTag` - 删除作品标签

#### 技术实现细节
- 作品状态：草稿(0)/已提交(1)/已批改(2)
- 软删除：del_flag 标记删除
- 点赞/收藏：star_num/collect_num 计数
- 查看次数：view_num 自动增加
- 来源场景：course(课程作业)/additional(附加作业)/create(自由创作)
- 云变量支持：has_cloud_data 标识

### 下一步计划

#### 第十阶段：新闻公告与附加作业
- [ ] 实现新闻公告管理接口（teaching_news）
- [ ] 实现附加作业管理接口（teaching_additional_work）

---

## [3.0.0-dev] - 2025-12-06

### 第八阶段：教学管理-课程模块 ✅

#### Added
- **课程相关 DAO 层** ✅
  - 创建 teaching_course 表的 Entity、DO、DAO 文件
  - 创建 teaching_course_unit 表的 Entity、DO、DAO 文件
  - 创建 teaching_course_dept 表的 Entity、DO、DAO 文件

- **课程管理接口 (10个)** ✅
  - `GET /api/v1/teaching/teachingCourse/list` - 课程列表（分页、筛选）
  - `GET /api/v1/teaching/teachingCourse/getHomeCourse` - 首页课程列表
  - `POST /api/v1/teaching/teachingCourse/add` - 添加课程
  - `PUT /api/v1/teaching/teachingCourse/edit` - 编辑课程
  - `DELETE /api/v1/teaching/teachingCourse/delete` - 删除课程
  - `DELETE /api/v1/teaching/teachingCourse/deleteBatch` - 批量删除课程
  - `GET /api/v1/teaching/teachingCourse/queryById` - 课程详情
  - `PUT /api/v1/teaching/teachingCourse/publish` - 发布/下架课程
  - `PUT /api/v1/teaching/teachingCourse/setShared` - 设置共享状态
  - `POST /api/v1/teaching/teachingCourse/authorizeDept` - 授权课程给部门

- **课程单元管理接口 (8个)** ✅
  - `GET /api/v1/teaching/teachingCourseUnit/list` - 单元列表（分页）
  - `GET /api/v1/teaching/teachingCourseUnit/queryByCourseId` - 课程所有单元
  - `POST /api/v1/teaching/teachingCourseUnit/add` - 添加单元
  - `PUT /api/v1/teaching/teachingCourseUnit/edit` - 编辑单元
  - `DELETE /api/v1/teaching/teachingCourseUnit/delete` - 删除单元
  - `DELETE /api/v1/teaching/teachingCourseUnit/deleteBatch` - 批量删除单元
  - `GET /api/v1/teaching/teachingCourseUnit/queryById` - 单元详情
  - `PUT /api/v1/teaching/teachingCourseUnit/sort` - 单元排序

- **班级课程管理接口 (6个)** ✅
  - `GET /api/v1/teaching/teachingCourseDept/list` - 班级课程列表
  - `GET /api/v1/teaching/teachingCourseDept/queryByDeptId` - 班级的课程列表
  - `GET /api/v1/teaching/teachingCourseDept/queryByCourseId` - 课程授权班级列表
  - `POST /api/v1/teaching/teachingCourseDept/addOrUpdate` - 添加/更新班级课程
  - `DELETE /api/v1/teaching/teachingCourseDept/delete` - 删除班级课程
  - `POST /api/v1/teaching/teachingCourseDept/batchAdd` - 批量添加班级课程

#### 技术实现细节
- 课程软删除：del_flag 标记删除
- 首页展示：show_home 字段控制首页展示
- 共享课程：is_shared 字段控制课程共享
- 部门授权：depart_ids 字段存储授权部门
- 课程单元：支持视频、PPT、作业、教案等多种资源
- 单元排序：支持拖拽排序

### 下一步计划

#### 第九阶段：教学管理-作品模块
- [ ] 实现作品管理接口（teaching_work）
- [ ] 实现作品批改接口（teaching_work_correct）
- [ ] 实现作品评论接口（teaching_work_comment）

---

## [3.0.0-dev] - 2025-12-06

### 第七阶段：文件上传管理 ✅

#### Added
- **文件相关 DAO 层** ✅
  - 手动创建 sys_file 表的 Entity、DO、DAO 文件

- **文件上传管理接口** ✅
  - `POST /api/v1/sys/file/upload` - 单文件上传
  - `POST /api/v1/sys/file/uploadBatch` - 批量文件上传
  - `GET /api/v1/sys/file/list` - 文件列表（分页、类型筛选、标签筛选）
  - `GET /api/v1/sys/file/queryById` - 文件详情
  - `DELETE /api/v1/sys/file/delete` - 删除文件（逻辑删除+物理文件删除）
  - `DELETE /api/v1/sys/file/deleteBatch` - 批量删除文件
  - `GET /api/v1/sys/file/view/:id` - 文件预览
  - `GET /api/v1/sys/file/download/:id` - 文件下载

- **Service层架构扩展**
  - 创建 `internal/service/sys_file.go` 文件服务接口
  - 创建 `internal/logic/sys/sys_file.go` 文件业务逻辑
  - 创建 `internal/controller/sys/sys_file.go` 文件控制器
  - 创建 `api/v1/sys/file.go` 文件API定义

#### 技术实现细节
- 文件类型识别：自动识别图片(1)/文档(2)/视频(3)/音频(4)/压缩包(5)/其他(0)
- 文件大小限制：默认50MB，可配置
- 存储位置：本地存储(1)，支持扩展云存储
- 文件命名：使用 UUID + 原始扩展名，避免重名
- 目录结构：按日期分目录存储（YYYY/MM/DD）
- 逻辑删除：文件记录逻辑删除，同时删除物理文件

### 下一步计划

#### 第八阶段：系统配置与公告管理
- [ ] 实现系统配置管理接口
- [ ] 实现公告管理接口

---

## [3.0.0-dev] - 2025-12-06

### 第六阶段：日志管理 ✅

#### Added
- **日志相关 DAO 层** ✅
  - 手动创建 sys_log 表的 Entity、DO、DAO 文件
  - 手动创建 sys_data_log 表的 Entity、DO、DAO 文件

- **系统日志管理接口** ✅
  - `GET /api/v1/sys/log/list` - 系统日志列表（分页、多条件筛选）
  - `DELETE /api/v1/sys/log/delete` - 删除系统日志
  - `DELETE /api/v1/sys/log/deleteBatch` - 批量删除系统日志
  - `DELETE /api/v1/sys/log/clear` - 清空系统日志（可按类型清空）

- **数据日志管理接口** ✅
  - `GET /api/v1/sys/dataLog/list` - 数据日志列表（分页、筛选）
  - `GET /api/v1/sys/dataLog/queryById` - 数据日志详情
  - `GET /api/v1/sys/dataLog/history` - 数据变更历史

- **Service层架构扩展**
  - 创建 `internal/service/sys_log.go` 日志服务接口
  - 创建 `internal/logic/sys/sys_log.go` 日志业务逻辑
  - 创建 `internal/controller/sys/sys_log.go` 日志控制器
  - 创建 `api/v1/sys/log.go` 日志API定义

#### 技术实现细节
- 日志类型：支持登录日志(1)和操作日志(2)
- 多条件筛选：用户名、IP、时间范围、日志内容
- 数据版本追踪：数据日志按版本号追踪变更
- 批量操作：支持批量删除和清空

### 下一步计划

#### 第七阶段：文件上传管理
- [ ] 实现文件上传接口（sys_file）
- [ ] 支持本地存储和云存储
- [ ] 实现文件预览和下载

---

## [3.0.0-dev] - 2025-12-05

### 第五阶段：字典管理 ✅

#### Added
- **字典相关 DAO 层** ✅
  - 手动创建 sys_dict 表的 Entity、DO、DAO 文件
  - 手动创建 sys_dict_item 表的 Entity、DO、DAO 文件

- **字典管理接口** ✅
  - `GET /api/v1/sys/dict/list` - 字典列表（分页、模糊搜索）
  - `POST /api/v1/sys/dict/add` - 添加字典（编码/名称唯一性校验）
  - `PUT /api/v1/sys/dict/edit` - 编辑字典
  - `DELETE /api/v1/sys/dict/delete` - 删除字典（逻辑删除，级联删除字典项）
  - `GET /api/v1/sys/dict/queryById` - 字典详情

- **字典项管理接口** ✅
  - `GET /api/v1/sys/dictItem/list` - 字典项列表（按字典ID查询）
  - `GET /api/v1/sys/dict/getDictItems/:dictCode` - 根据字典编码获取字典项
  - `POST /api/v1/sys/dictItem/add` - 添加字典项
  - `PUT /api/v1/sys/dictItem/edit` - 编辑字典项
  - `DELETE /api/v1/sys/dictItem/delete` - 删除字典项

- **Service层架构扩展**
  - 创建 `internal/service/sys_dict.go` 字典服务接口
  - 创建 `internal/logic/sys/sys_dict.go` 字典业务逻辑
  - 创建 `internal/controller/sys/sys_dict.go` 字典控制器
  - 创建 `api/v1/sys/dict.go` 字典API定义

#### 技术实现细节
- 逻辑删除：字典使用 del_flag=1 标记删除
- 级联删除：删除字典时同时删除所有字典项
- 排序支持：字典项按 sort_order 字段排序
- 状态控制：字典项支持启用/禁用状态

### 下一步计划

#### 第六阶段：日志与其他管理
- [ ] 实现操作日志管理接口（sys_log）
- [ ] 实现登录日志管理接口
- [ ] 完善系统配置管理

---

## [3.0.0-dev] - 2025-12-04

### 第四阶段：部门管理 ✅

#### Added
- **部门管理相关 DAO 层** ✅
  - 手动创建 sys_depart 表的 Entity、DO、DAO 文件
  - 手动创建 sys_user_depart 表的 Entity、DO、DAO 文件

- **部门管理接口** ✅
  - `GET /api/v1/sys/sysDepart/queryTreeList` - 部门树列表
  - `POST /api/v1/sys/sysDepart/add` - 添加部门（自动生成orgCode）
  - `PUT /api/v1/sys/sysDepart/edit` - 编辑部门
  - `DELETE /api/v1/sys/sysDepart/delete` - 删除部门（递归删除子部门）
  - `GET /api/v1/sys/sysDepart/queryIdTree` - 部门ID树结构
  - `GET /api/v1/sys/sysDepart/searchBy` - 搜索部门
  - `GET /api/v1/sys/sysDepart/queryDepartTreeSync` - 部门用户树

- **Service层架构扩展**
  - 创建 `internal/service/sys_depart.go` 部门服务接口
  - 创建 `internal/logic/sys/sys_depart.go` 部门业务逻辑
  - 创建 `internal/controller/sys/sys_depart.go` 部门控制器
  - 创建 `api/v1/sys/depart.go` 部门API定义

#### 技术实现细节
- 树形结构构建：递归构建部门树
- orgCode自动生成：根据父部门编码自动生成子部门编码（格式：A01, A01A01）
- 递归删除：删除部门时自动删除所有子部门
- 用户关联：支持部门-用户关联查询

### 下一步计划

#### 第五阶段：字典与日志管理
- [ ] 实现字典管理接口（sys_dict, sys_dict_item）
- [ ] 实现日志管理接口

---

## [3.0.0-dev] - 2025-12-03

### 第三阶段：角色权限管理 ✅

#### Added
- **角色相关表的 DAO 层** ✅
  - 生成 sys_role 表的 Entity、DO、DAO 文件
  - 生成 sys_user_role 表的 Entity、DO、DAO 文件
  - 生成 sys_permission 表的 Entity、DO、DAO 文件
  - 生成 sys_role_permission 表的 Entity、DO、DAO 文件

- **角色管理接口** ✅
  - `GET /api/v1/sys/role/list` - 角色列表（分页、模糊搜索）
  - `POST /api/v1/sys/role/add` - 添加角色（编码/名称唯一性校验）
  - `PUT /api/v1/sys/role/edit` - 编辑角色
  - `DELETE /api/v1/sys/role/delete` - 删除角色（批量，检查使用情况）
  - `GET /api/v1/sys/role/getById` - 角色详情
  - `GET /api/v1/sys/role/all` - 获取所有角色（下拉选择用）

- **用户角色关联接口** ✅
  - `GET /api/v1/sys/role/getUserRoles` - 获取用户角色
  - `POST /api/v1/sys/role/saveUserRoles` - 保存用户角色

- **权限管理接口** ✅
  - `GET /api/v1/sys/permission/list` - 权限列表（平铺）
  - `GET /api/v1/sys/permission/queryTreeList` - 权限树
  - `POST /api/v1/sys/permission/add` - 添加权限
  - `PUT /api/v1/sys/permission/edit` - 编辑权限
  - `DELETE /api/v1/sys/permission/delete` - 删除权限（批量，检查子菜单）
  - `GET /api/v1/sys/permission/getById` - 权限详情
  - `GET /api/v1/sys/permission/queryRolePermission` - 查询角色权限
  - `POST /api/v1/sys/permission/saveRolePermission` - 保存角色权限
  - `GET /api/v1/sys/permission/getUserPermissionByToken` - 获取用户权限

- **Service层架构扩展**
  - 创建 `internal/service/sys_role.go` 角色服务接口
  - 创建 `internal/logic/sys/sys_role.go` 角色业务逻辑
  - 创建 `internal/controller/sys/sys_role.go` 角色控制器
  - 创建 `api/v1/sys/role.go` 角色API定义
  - 创建 `internal/service/sys_permission.go` 权限服务接口
  - 创建 `internal/logic/sys/sys_permission.go` 权限业务逻辑
  - 创建 `internal/controller/sys/sys_permission.go` 权限控制器
  - 创建 `api/v1/sys/permission.go` 权限API定义

### 下一步计划

#### 第四阶段：部门与字典管理
- [ ] 实现部门管理接口
- [ ] 实现字典管理接口

---

### 第二阶段：数据访问层与认证系统 ✅

#### Added
- **DAO层代码生成**
  - 生成 sys_user 表的 Entity、DO、DAO 文件
  - 配置 hack/config.yaml 用于 DAO 生成

- **用户登录接口** ✅
  - 实现 `POST /api/v1/sys/login` 登录接口
  - MD5+salt 密码验证
  - JWT Token 生成（2小时有效期）
  - 自动过滤已删除用户
  - 用户状态检查（冻结/正常）

- **用户管理CRUD接口** ✅
  - `GET /api/v1/sys/user/list` - 用户列表（分页、模糊搜索、状态筛选）
  - `POST /api/v1/sys/user` - 添加用户（UUID生成、密码加密、唯一性校验）
  - `PUT /api/v1/sys/user` - 编辑用户（部分字段更新、唯一性校验）
  - `DELETE /api/v1/sys/user/:id` - 删除用户（逻辑删除）
  - `GET /api/v1/sys/user/:id` - 用户详情

- **Service层架构**
  - 创建 `internal/service/sys_user.go` 服务接口
  - 创建 `internal/logic/sys/sys_user.go` 业务逻辑实现
  - 创建 `internal/controller/sys/sys_user.go` 控制器
  - 创建 `api/v1/sys/user.go` API定义

#### Documentation
- 生成工作总结文档 `api-go/WORK_SUMMARY.md`
- 更新接口报告 `docs/20251122/前后端接口报告.md`
- 创建 Git 提交脚本 `git-commit-user-module.sh`
- 创建编译测试脚本 `compile-test.sh`

#### 技术实现细节
- 密码加密：MD5(password + 8位随机salt)
- ID生成：使用 guid.S() 生成 UUID
- 分页查询：支持 page/pageSize 参数，最大100条/页
- 唯一性校验：username/phone/email 自动检查重复
- 逻辑删除：设置 DelFlag=1

### 下一步计划

#### 第三阶段：角色权限管理
- [ ] 生成角色相关表的 DAO 层
- [ ] 实现角色管理 CRUD 接口
- [ ] 实现权限管理接口
- [ ] 实现部门管理接口
- [ ] 实现用户角色关联

---

## [3.0.0-dev] - 2025-11-22

### 第一阶段：项目初始化与基础设施 ✅

#### Added
- 使用 `gf init` 初始化 GoFrame V2 后端项目
- 创建标准项目目录结构（api/, internal/, manifest/, utility/, resource/）
- 配置项目基础信息（go.mod, config.yaml）
- 实现核心常量定义（consts.go, error.go）
- 实现统一响应格式工具类（response包）
- 实现JWT认证工具类（jwt包）
- 实现中间件系统：
  - CORS 跨域中间件
  - Logger 日志中间件
  - Error 错误处理中间件
  - Auth 认证中间件
- 配置路由系统和健康检查接口

#### Configuration
- 数据库配置：MySQL (teachingopen数据库)
- Redis配置：127.0.0.1:6379
- JWT配置：2小时过期时间
- 服务端口：8199
- 文件上传配置：最大50MB，支持多种文件类型
- CORS跨域配置：允许所有来源

#### Documentation
- 添加项目初始化脚本 `init-goframe.sh`
- 添加项目初始化指南 `api-go/INIT_GUIDE.md`
- 添加开发状态文档 `api-go/DEVELOPMENT_STATUS.md`
- 添加文件修复脚本 `fix-files.sh`
- 生成未完成工作报告（docs/20251122/未完成工作报告.md）
- 创建前后端接口规范文档（docs/20251122/前后端接口报告.md）

#### Infrastructure
- 设置项目分支策略（master -> devgo）
- 配置 Git 仓库结构
- 配置 GoFrame CLI 工具
- 配置 DAO 生成参数（hack/config.yaml）

### 初始化阶段

#### Added
- 创建 GoFrame V2 后端项目基础结构
- 创建标准项目目录结构（api/, internal/, manifest/, utility/, resource/）
- 添加项目初始化脚本 `init-goframe.sh`
- 添加项目初始化指南 `api-go/INIT_GUIDE.md`
- 创建重建项目脚本 `rebuild-project.sh`

#### Documentation
- 完成 PRD 产品需求文档
- 完成 GoFrame V2 开发指南
- 完成 Vue3 开发指南
- 生成未完成工作报告（docs/20251122/未完成工作报告.md）

#### Infrastructure
- 设置项目分支策略（master -> devgo）
- 配置 Git 仓库结构
- 准备开发环境

### 计划中的功能

#### 后端 (GoFrame V2)
- [x] JWT 认证系统 ✅ 2025-12-03 完成
- [x] 用户管理模块 ✅ 2025-12-03 完成
- [ ] 角色权限管理
- [ ] 部门管理
- [ ] 课程管理
- [ ] 作品管理
- [ ] 学生管理
- [ ] 文件上传
- [ ] 数据字典
- [ ] 日志管理

#### 前端 (Vue3)
- [ ] 项目初始化
- [ ] 登录页面
- [ ] 用户管理页面
- [ ] 角色管理页面
- [ ] 课程管理页面
- [ ] 作品管理页面
- [ ] Scratch 编辑器集成
- [ ] Python 编辑器集成

---

## [2.8] - 历史版本

### Added
- 文章资讯功能
- 作品打标签功能
- 自定义首页编辑功能
- ScratchJr升级
- Scratch升级
- 本地文件模式支持视频进度拖动
- 文件预览优化
- 用户信息增加学校字段
- 可设置是否开放作品评论
- 增加学生作品页面

### Fixed
- SQL注入漏洞修复
- swagger漏洞修复
- 其他bug修复

---

## [2.7]

### Added
- 首页移动端适配和UI优化
- 作品分享页面优化
- 增加Google blockly编辑器
- 增加角色等级配置：低等级权限的角色无法操作高能级权限的角色
- 权限系统完善：非管理员只能看到自己负责的部门、学生、作业
- 课程权限：课程授权给指定部门，可能在班级中分配该课程
- 增加地图编辑器
- 可编辑后台首页内容
- 可编辑作品分享页文案
- 可添加自定义JS和CSS
- 部门管理优化
- 查询条件优化
- 简化前端配置
- 支持富文本课程内容
- 课程支持首页展示
- PPT资料在线预览

### Fixed
- 修复课程开始时间无效bug
- 修复七牛云删除文件bug
- 修复JUpload组件bug
- 修复首页作品排序bug
- 修复作品预览bug
- 修复本地存储bug

---

## [2.6]

### Added
- 支持布置自定义作业
- 增加在线网站配置
- 增加前端自定义菜单配置
- 课程视频支持外链和外部播放器
- 课程视频屏蔽右键和下载按钮
- 富文本编辑器支持源码编辑、插入代码、预览
- 课程增加排序功能
- 自定义素材库优先显示
- 显示老师评分和评语

### Fixed
- 修复Scratch提交作业bug
- 修复上传组件文件数量限制bug
- 自定义素材库管理bug
- 注册用户名字段bug
- 修复其他bug

---

## [2.5]

### Changed
- 首页改为社区

### Fixed
- 修复其他bug

---

## [2.4]

### Added
- Scratch素材库管理，支持上传自定义素材库
- 可设置本地存储，脱离云存储
- 可设置共享课程，所有学生都可学习
- 删除记录的同时物理删除文件，释放硬盘和云存储空间
- Scratch同步至官方最新
- 样式优化

### Fixed
- 修复其他bug

---

## [2.3]

### Added
- 延长用户登录失效时间（5小时）
- 提交Scratch可选是否分享，提供获取作品排行接口
- 增加用户头像上传入口
- 课程视频屏蔽右键下载
- 手机端加载项目时出现小猫
- 开放注册（可在Login.vue中关闭）

### Fixed
- 修复ScratchJr移动端无法使用

---

## [2.2]

### Added
- 接入Python turtle编辑器
- 课件支持Scratch案例

### Fixed
- 修复若干BUG

---

## [2.1]

### Added
- 集成ScratchJr
- 简化配置

### Fixed
- 修复若干BUG

---

## 版本说明

### 版本号规则
- 主版本号：重大架构变更（如 3.0.0 从 Java 迁移到 Golang）
- 次版本号：新功能添加
- 修订号：Bug 修复和小改进

### 当前开发状态
- 版本：3.0.0-dev
- 阶段：技术架构迁移
- 计划：14周开发周期

### 相关文档
- PRD：docs/PRD.md
- 开发指南：docs/goFrameV2 dev guide.md, docs/Vue3 Dev Guide.md
- 未完成工作：docs/20251122/未完成工作报告.md
