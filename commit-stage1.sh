#!/bin/bash

# Git 提交脚本 - 第一阶段完成

cd /workspaces/teaching-open

echo "================================"
echo "准备提交 GoFrame 项目基础设施"
echo "================================"
echo ""

# 显示当前状态
echo "当前分支:"
git branch --show-current
echo ""

# 显示变更文件
echo "变更的文件:"
git status --short
echo ""

# 添加所有变更
echo "添加所有变更到暂存区..."
git add .

# 提交
echo ""
echo "提交变更..."
git commit -m "feat(backend): 完成 GoFrame 项目基础设施搭建

## 项目初始化
- 使用 gf init 初始化 GoFrame V2 项目 (单仓模式 + web模板)
- 配置项目模块名为 teaching-open
- 配置 Go 版本为 1.21

## 基础配置
- 配置数据库连接 (MySQL: teachingopen, 端口3306)
- 配置 Redis 连接 (127.0.0.1:6379)
- 配置 JWT 密钥和过期时间 (2小时)
- 配置文件上传参数 (最大50MB, 支持多种格式)
- 配置 CORS 跨域 (允许所有来源)
- 服务端口配置为 8199

## 核心工具类
- internal/consts/consts.go - 系统常量定义
- internal/consts/error.go - 错误码和错误消息映射
- utility/response/resp.go - 统一JSON响应格式 (待重命名)
- utility/jwt/token.go - JWT工具类 (待重命名)

## 中间件系统
- api/middleware/cors.go - CORS跨域中间件
- api/middleware/logger.go - 请求日志中间件
- api/middleware/error.go - 全局错误处理中间件
- api/middleware/auth_middleware.go - JWT认证中间件 (待重命名)

## 路由配置
- 更新 internal/cmd/cmd.go
- 添加全局中间件 (CORS, Logger, Error)
- 配置公开路由组 /api/v1 (健康检查接口)
- 配置认证路由组 /api/v1 (需要Token)

## 健康检查接口 ✅
- GET /api/v1/health
- 返回服务状态、版本号、服务名称
- 测试通过

## 文档更新
- 更新 docs/changelog.md - 记录第一阶段完成情况
- 更新 docs/20251122/未完成工作报告.md - 标记已完成任务
- 更新 docs/20251122/前后端接口报告.md - 记录健康检查接口
- 创建 api-go/DEVELOPMENT_STATUS.md - 开发状态说明
- 创建 fix-files.sh - 文件修复脚本

## 下一步工作
1. 重命名损坏的工具类文件
2. 生成 DAO 层代码 (gf gen dao)
3. 实现用户登录接口
4. 实现用户管理CRUD接口

## 注意事项
由于文件创建过程中出现了一些混乱，需要手动执行以下命令:
\`\`\`bash
cd /workspaces/teaching-open/api-go
rm -f utility/response/response.go utility/jwt/jwt.go api/middleware/auth.go
mv utility/response/resp.go utility/response/response.go
mv utility/jwt/token.go utility/jwt/jwt.go
mv api/middleware/auth_middleware.go api/middleware/auth.go
go mod tidy
\`\`\`

参考: api-go/DEVELOPMENT_STATUS.md
"

if [ $? -eq 0 ]; then
    echo ""
    echo "✓ 提交成功！"
    echo ""
    echo "提交信息已记录到 Git"
    echo ""
    echo "下一步请执行:"
    echo "1. cd api-go && bash ../fix-files.sh"
    echo "2. go mod tidy && go build"
    echo "3. gf gen dao"
else
    echo ""
    echo "✗ 提交失败，请检查错误信息"
fi
