# Requirements Document

## Introduction

本文档定义了将 Teaching Open 教学管理平台从 Java (Spring Boot + JeecgBoot) 迁移到 Go (GoFrame) 的完整需求规范。该系统是一个教育管理平台，包含用户管理、课程管理、作业管理、权限控制等核心功能模块。

## Glossary

- **Teaching_Open_System**: 教学开放平台系统，提供课程管理、作业管理、用户管理等功能
- **GoFrame**: Go语言的Web框架，用于替代Spring Boot
- **JWT_Service**: JSON Web Token服务，用于用户认证和授权
- **RBAC_System**: 基于角色的访问控制系统
- **Course_Service**: 课程管理服务
- **Work_Service**: 作业管理服务
- **User_Service**: 用户管理服务
- **Depart_Service**: 部门/班级管理服务
- **Dict_Service**: 数据字典服务
- **File_Service**: 文件存储服务
- **Redis_Cache**: Redis缓存服务

---

## Requirements

### Requirement 1: 用户认证与授权系统

**User Story:** As a 系统用户, I want to 通过用户名密码或手机验证码登录系统, so that I can 安全地访问系统功能。

#### Acceptance Criteria

1. WHEN a user submits valid username, password, and captcha, THE JWT_Service SHALL generate a JWT token and store it in Redis with configurable expiration time
2. WHEN a user submits invalid credentials, THE JWT_Service SHALL return an error message without revealing which field is incorrect
3. WHEN a user submits an expired or invalid token, THE JWT_Service SHALL return 401 Unauthorized status
4. WHEN a user's token is about to expire but still valid in Redis cache, THE JWT_Service SHALL refresh the token automatically
5. WHEN a user logs out, THE JWT_Service SHALL invalidate the token in Redis cache
6. WHEN a user attempts phone login with valid SMS code, THE JWT_Service SHALL authenticate and generate token
7. IF the system is configured for single-device login, THEN THE JWT_Service SHALL invalidate previous tokens when a new login occurs
8. THE JWT_Service SHALL encrypt passwords using MD5 with salt before storage and comparison

---

### Requirement 2: 用户管理模块

**User Story:** As a 系统管理员, I want to 管理系统用户信息, so that I can 控制用户访问权限和维护用户数据。

#### Acceptance Criteria

1. THE User_Service SHALL support CRUD operations for user entities with fields: id, username, realname, password, salt, avatar, birthday, sex, email, phone, orgCode, status, delFlag, workNo, post, school, telephone, departIds, thirdId, thirdType
2. WHEN creating a user, THE User_Service SHALL generate a unique salt and encrypt the password
3. WHEN querying users, THE User_Service SHALL support pagination, filtering by department, role, and status
4. WHEN deleting a user, THE User_Service SHALL perform logical deletion by setting delFlag=1
5. THE User_Service SHALL support batch import/export of users via Excel
6. WHEN a user is assigned to departments, THE User_Service SHALL maintain the user-department relationship in sys_user_depart table
7. WHEN a user is assigned roles, THE User_Service SHALL maintain the user-role relationship in sys_user_role table
8. THE User_Service SHALL support querying users by phone, email, or third-party ID (WeChat, etc.)

---

### Requirement 3: 角色与权限管理

**User Story:** As a 系统管理员, I want to 配置角色和权限, so that I can 实现细粒度的访问控制。

#### Acceptance Criteria

1. THE RBAC_System SHALL support role entities with fields: id, roleName, roleCode, roleLevel, description
2. THE RBAC_System SHALL support permission entities with fields: id, parentId, name, perms, icon, component, url, menuType, sortNo, hidden, status
3. WHEN assigning permissions to a role, THE RBAC_System SHALL maintain the role-permission relationship in sys_role_permission table
4. THE RBAC_System SHALL support three permission types: menu (menuType=0), submenu (menuType=1), and button (menuType=2)
5. WHEN a user requests a protected resource, THE RBAC_System SHALL verify the user has the required permission
6. THE RBAC_System SHALL support data-level permission rules via sys_permission_data_rule table
7. THE RBAC_System SHALL cache user permissions in Redis for performance optimization
8. WHEN permissions are modified, THE RBAC_System SHALL invalidate the affected user's permission cache

---

### Requirement 4: 部门/班级管理

**User Story:** As a 管理员, I want to 管理组织架构和班级信息, so that I can 组织用户和课程资源。

#### Acceptance Criteria

1. THE Depart_Service SHALL support department entities with fields: id, parentId, departName, departNameEn, departOrder, description, orgCategory, orgType, orgCode, mobile, fax, address, memo, status, delFlag
2. THE Depart_Service SHALL support tree structure for department hierarchy
3. WHEN querying departments, THE Depart_Service SHALL return tree-structured data with parent-child relationships
4. THE Depart_Service SHALL generate unique orgCode for each department following hierarchical pattern
5. WHEN a department is deleted, THE Depart_Service SHALL check for child departments and associated users
6. THE Depart_Service SHALL support department-level roles via sys_depart_role table
7. THE Depart_Service SHALL support querying all parent department IDs for permission inheritance

---

### Requirement 5: 数据字典管理

**User Story:** As a 开发者, I want to 使用数据字典管理系统枚举值, so that I can 统一管理下拉选项和状态码。

#### Acceptance Criteria

1. THE Dict_Service SHALL support dictionary entities with fields: id, dictName, dictCode, description, type, delFlag
2. THE Dict_Service SHALL support dictionary item entities with fields: id, dictId, itemText, itemValue, description, sortOrder, status
3. WHEN querying dictionary items, THE Dict_Service SHALL return items sorted by sortOrder
4. THE Dict_Service SHALL support querying all dictionary items for frontend initialization
5. THE Dict_Service SHALL support table-based dictionary lookup (dictTable, dicText, dicCode)
6. THE Dict_Service SHALL cache dictionary data in Redis for performance

---

### Requirement 6: 课程管理模块

**User Story:** As a 教师, I want to 创建和管理课程内容, so that I can 为学生提供学习资源。

#### Acceptance Criteria

1. THE Course_Service SHALL support course entities with fields: id, courseName, courseDesc, courseIcon, courseCover, courseMap, courseType, courseCategory, showHome, isShared, departIds, orderNum, showType, delFlag
2. THE Course_Service SHALL support course unit entities with fields: id, courseId, unitName, unitDesc, courseVideo, courseCase, coursePpt, courseWork, courseWorkAnswer, coursePlan, mapX, mapY, mediaContent, orderNum
3. WHEN querying courses, THE Course_Service SHALL support filtering by department authorization (departIds)
4. WHEN a course is marked as showHome=1, THE Course_Service SHALL include it in homepage course list
5. THE Course_Service SHALL support course-department authorization via teaching_course_dept table
6. WHEN deleting a course, THE Course_Service SHALL also delete associated files (icon, cover, map)
7. THE Course_Service SHALL support querying user's enrolled courses based on department membership

---

### Requirement 7: 作业管理模块

**User Story:** As a 学生, I want to 提交和管理我的作业, so that I can 完成课程学习任务。

#### Acceptance Criteria

1. THE Work_Service SHALL support work entities with fields: id, userId, departId, courseId, additionalId, workName, workType, workStatus, workFile, workCover, viewNum, starNum, collectNum, delFlag, workScene, hasCloudData
2. WHEN a student submits work, THE Work_Service SHALL associate it with the student's current department
3. THE Work_Service SHALL support work types: scratch, python, blockly, etc. (via work_type dictionary)
4. THE Work_Service SHALL support work status tracking: draft, submitted, corrected (via work_status dictionary)
5. THE Work_Service SHALL support work comments via teaching_work_comment table
6. THE Work_Service SHALL support work correction/grading via teaching_work_correct table with score and comment
7. WHEN querying works, THE Work_Service SHALL support filtering by user, department, course, and status
8. THE Work_Service SHALL track view count, star count, and collect count for each work

---

### Requirement 8: 附加作业管理

**User Story:** As a 教师, I want to 布置额外的作业任务, so that I can 补充课程内容。

#### Acceptance Criteria

1. THE Work_Service SHALL support additional work entities with fields: id, workName, workDesc, workFile, departId
2. WHEN assigning additional work, THE Work_Service SHALL associate it with specific departments
3. THE Work_Service SHALL support linking student submissions to additional work via additionalId field

---

### Requirement 9: 文件存储服务

**User Story:** As a 用户, I want to 上传和管理文件, so that I can 存储课程资源和作业文件。

#### Acceptance Criteria

1. THE File_Service SHALL support file entities with fields: id, fileName, filePath, fileType, fileLocation, fileTag, delFlag
2. THE File_Service SHALL support multiple storage backends: local disk, Qiniu Cloud, Aliyun OSS, MinIO
3. WHEN uploading a file, THE File_Service SHALL generate a unique file key and store metadata in database
4. THE File_Service SHALL support file type detection and validation
5. WHEN deleting a file record, THE File_Service SHALL also delete the physical file from storage
6. THE File_Service SHALL support generating signed URLs for private file access

---

### Requirement 10: 系统日志与审计

**User Story:** As a 管理员, I want to 查看系统操作日志, so that I can 审计用户行为和排查问题。

#### Acceptance Criteria

1. THE Teaching_Open_System SHALL log all user operations with fields: id, logType, logContent, operateType, userid, username, ip, method, requestUrl, requestParam, requestType, costTime
2. THE Teaching_Open_System SHALL support two log types: login log (logType=1) and operation log (logType=2)
3. WHEN a user performs CRUD operations, THE Teaching_Open_System SHALL automatically log the operation via AOP
4. THE Teaching_Open_System SHALL support querying logs with pagination and filtering by user, type, and date range
5. THE Teaching_Open_System SHALL log request execution time for performance monitoring

---

### Requirement 11: 消息与通知系统

**User Story:** As a 用户, I want to 接收系统通知和消息, so that I can 及时了解重要信息。

#### Acceptance Criteria

1. THE Teaching_Open_System SHALL support announcement entities with fields: id, title, msgContent, startTime, endTime, sender, priority, msgCategory, msgType, sendStatus
2. THE Teaching_Open_System SHALL support sending announcements to all users or specific users
3. THE Teaching_Open_System SHALL track announcement read status per user via sys_announcement_send table
4. THE Teaching_Open_System SHALL support SMS sending via Aliyun SMS service
5. THE Teaching_Open_System SHALL support email sending via SMTP

---

### Requirement 12: 前端菜单管理

**User Story:** As a 前端用户, I want to 看到个性化的导航菜单, so that I can 快速访问相关功能。

#### Acceptance Criteria

1. THE Teaching_Open_System SHALL support teaching menu entities with fields: id, parentId, name, icon, url, menuType, sortNo, isLeaf, isRoute, hidden, needLogin, internalOrExternal
2. THE Teaching_Open_System SHALL return menu tree structure for frontend rendering
3. WHEN needLogin=1, THE Teaching_Open_System SHALL require authentication to access the menu item
4. THE Teaching_Open_System SHALL support external link menus via internalOrExternal field

---

### Requirement 13: 班级日志统计

**User Story:** As a 管理员, I want to 查看班级教学活动统计, so that I can 监控教学进度。

#### Acceptance Criteria

1. THE Teaching_Open_System SHALL support department day log entities with fields: id, departId, departName, unitOpenCount, courseWorkAssignCount, additionalWorkAssignCount, courseWorkCorrectCount, additionalWorkCorrectCount, courseWorkSubmitCount, additionalWorkSubmitCount, createTime
2. THE Teaching_Open_System SHALL aggregate daily statistics for each department
3. WHEN querying statistics, THE Teaching_Open_System SHALL support date range filtering

---

### Requirement 14: Scratch素材管理

**User Story:** As a Scratch用户, I want to 管理编程素材库, so that I can 在创作中使用丰富的素材。

#### Acceptance Criteria

1. THE Teaching_Open_System SHALL support scratch asset entities with fields: id, assetType, assetName, assetData, md5Ext, tags, delFlag
2. THE Teaching_Open_System SHALL support asset types: background (1), sound (2), costume (3), sprite (4)
3. WHEN querying assets, THE Teaching_Open_System SHALL support filtering by type and tags
4. THE Teaching_Open_System SHALL store asset JSON data for Scratch editor integration

---

### Requirement 15: API响应格式标准化

**User Story:** As a 前端开发者, I want to 接收统一格式的API响应, so that I can 一致地处理后端返回数据。

#### Acceptance Criteria

1. THE Teaching_Open_System SHALL return all API responses in format: {success: boolean, message: string, code: number, result: any, timestamp: number}
2. WHEN an operation succeeds, THE Teaching_Open_System SHALL return success=true with code=200
3. WHEN an operation fails, THE Teaching_Open_System SHALL return success=false with appropriate error code and message
4. THE Teaching_Open_System SHALL support pagination response format: {records: [], total: number, size: number, current: number, pages: number}

---

### Requirement 16: 定时任务管理

**User Story:** As a 管理员, I want to 配置和管理定时任务, so that I can 自动化执行周期性操作。

#### Acceptance Criteria

1. THE Teaching_Open_System SHALL support quartz job entities with fields: id, jobClassName, cronExpression, parameter, description, status
2. THE Teaching_Open_System SHALL support starting, stopping, and executing jobs immediately
3. THE Teaching_Open_System SHALL persist job state across system restarts

---

### Requirement 17: 微信集成

**User Story:** As a 用户, I want to 使用微信账号登录系统, so that I can 便捷地访问平台。

#### Acceptance Criteria

1. THE Teaching_Open_System SHALL support wechat user entities with fields: id, userId, openId, unionId, nickname, sex, city, province, headimgurl, appId
2. WHEN a user logs in via WeChat, THE Teaching_Open_System SHALL create or link the wechat user record
3. THE Teaching_Open_System SHALL support binding WeChat account to existing system user

---

### Requirement 18: 数据库兼容性

**User Story:** As a 运维人员, I want to 系统支持MySQL数据库, so that I can 使用现有的数据库基础设施。

#### Acceptance Criteria

1. THE Teaching_Open_System SHALL use MySQL as the primary database
2. THE Teaching_Open_System SHALL use UUID (ID_WORKER_STR) as primary key generation strategy
3. THE Teaching_Open_System SHALL support soft delete via delFlag field
4. THE Teaching_Open_System SHALL maintain audit fields: createBy, createTime, updateBy, updateTime, sysOrgCode

---

### Requirement 19: Redis缓存集成

**User Story:** As a 系统, I want to 使用Redis缓存热点数据, so that I can 提高系统响应速度。

#### Acceptance Criteria

1. THE Teaching_Open_System SHALL cache JWT tokens in Redis with configurable TTL
2. THE Teaching_Open_System SHALL cache user permissions in Redis
3. THE Teaching_Open_System SHALL cache dictionary data in Redis
4. THE Teaching_Open_System SHALL support Redis connection pool configuration
5. WHEN cached data is modified, THE Teaching_Open_System SHALL invalidate the corresponding cache entries
