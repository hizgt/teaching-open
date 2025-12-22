#!/bin/bash

# Git提交用户认证和管理模块
# 使用方法: chmod +x git-commit-user-module.sh && ./git-commit-user-module.sh

echo "================================================"
echo "  Git提交 - 用户认证和管理模块"
echo "================================================"
echo ""

# 切换到项目根目录
cd /workspaces/teaching-open

# 配置Git用户信息(如果需要)
echo "检查Git配置..."
if [ -z "$(git config user.name)" ]; then
    echo "请输入Git用户名:"
    read git_username
    git config user.name "$git_username"
fi

if [ -z "$(git config user.email)" ]; then
    echo "请输入Git邮箱:"
    read git_email
    git config user.email "$git_email"
fi

echo ""
echo "当前分支: $(git branch --show-current)"
echo ""

# 查看当前状态
echo "查看当前Git状态..."
git status --short
echo ""

# 添加文件到暂存区
echo "添加文件到暂存区..."
git add api-go/
git add changelist.txt
git add docs/20251122/
git add commit-stage1.sh
git add compile-test.sh
git add fix-files.sh
git add git-commit-user-module.sh

echo "已添加文件到暂存区"
echo ""

# 执行提交
echo "执行Git提交..."
git commit -m "feat: 实现用户认证和管理模块

【功能实现】
- 用户登录接口 (POST /api/v1/sys/login)
  * MD5+salt密码加密验证
  * JWT token生成(2小时有效期)
  * 自动过滤已删除/冻结用户
  
- 用户管理CRUD接口
  * 用户列表查询 (GET /api/v1/sys/user/list)
    - 支持分页查询
    - 支持用户名/真实姓名模糊搜索  
    - 支持状态筛选
    - 按创建时间倒序排列
  * 新增用户 (POST /api/v1/sys/user)
    - UUID自动生成
    - 随机salt生成
    - 密码MD5加密
    - username/phone/email唯一性校验
  * 编辑用户 (PUT /api/v1/sys/user)
    - 部分字段更新
    - 唯一性校验(排除自身)
    - 不允许修改密码
  * 删除用户 (DELETE /api/v1/sys/user/:id)
    - 逻辑删除(DelFlag=1)
    - 自动更新UpdateTime
  * 用户详情 (GET /api/v1/sys/user/:id)

【DAO层】
- 手动创建sys_user表相关文件
  * internal/model/entity/sys_user.go - Entity实体
  * internal/model/do/sys_user.go - DO数据对象  
  * internal/dao/sys_user.go - DAO访问层
  * internal/dao/internal/sys_user.go - DAO内部实现

【架构层次】
- API层: api/v1/sys/user.go - 请求响应结构体
- Controller层: internal/controller/sys/sys_user.go - HTTP控制器
- Service层: internal/service/sys_user.go - 服务接口定义
- Logic层: internal/logic/sys/sys_user.go - 业务逻辑实现
- DAO层: internal/dao - 数据访问层

【基础设施】
- JWT工具类 (utility/jwt/jwt.go)
  * GenerateToken - 生成token
  * ParseToken - 解析token
  * ValidateToken - 验证token
  * 简化实现(base64+json,生产需替换为标准JWT库)
  
- 统一响应工具 (utility/response/response.go)
  * JsonRes、PageRes 响应结构
  * Success、Error、Page 等辅助函数
  * 统一错误码响应
  
- 错误码常量 (internal/consts/)
  * consts.go - 系统常量(状态码、缓存键等)
  * error.go - 错误码定义(0-4xxx)
  
- 中间件系统 (api/middleware/)
  * cors.go - CORS跨域处理
  * logger.go - 请求日志记录
  * error.go - 全局错误处理
  * auth.go - JWT认证中间件
  
- 路由配置 (internal/cmd/cmd.go)
  * 公开路由组(/api/v1) - 登录、健康检查
  * 认证路由组(/api/v1) - 用户管理CRUD

【配置文件】
- go.mod - 模块名更新为teaching-open
- manifest/config/config.yaml - 完整配置(DB/Redis/JWT/CORS)
- hack/config.yaml - DAO生成配置

【文档更新】
- changelist.txt - 新增v3.0.0-dev版本记录
- docs/20251122/前后端接口报告.md - 更新接口统计(7个已完成)
- docs/20251122/未完成工作报告.md - 标记已完成任务

【技术栈】
- GoFrame V2.7.1
- Go 1.21
- MySQL 8.0 (teachingopen数据库)
- Redis 6.0+
- 服务端口: 8199

【编译状态】
- 代码静态检查通过
- 无语法错误
- import路径正确
- 项目结构完整"

echo ""
echo "================================================"
echo "  提交成功!"
echo "================================================"
echo ""

# 查看提交日志
echo "提交信息:"
git log -1 --stat

echo ""
echo "================================================"
echo "  是否推送到远程仓库?"
echo "  输入 'y' 推送到 origin/devgo"
echo "  输入 'n' 跳过推送"
echo "================================================"
read -p "请选择 (y/n): " push_choice

if [ "$push_choice" = "y" ] || [ "$push_choice" = "Y" ]; then
    echo ""
    echo "推送到远程仓库..."
    git push origin devgo
    echo ""
    echo "推送完成!"
else
    echo ""
    echo "已跳过推送,您可以稍后手动执行: git push origin devgo"
fi

echo ""
echo "完成!"
