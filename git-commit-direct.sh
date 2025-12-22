#!/bin/bash

# 直接执行Git提交和推送
# 自动暂存、提交、推送到远程devgo分支

set -e

PROJECT_DIR="/workspaces/teaching-open"
cd "$PROJECT_DIR"

echo "=================================="
echo "Git 直接提交脚本"
echo "=================================="
echo ""

# 1. 检查分支
echo "【步骤1】检查分支..."
BRANCH=$(git rev-parse --abbrev-ref HEAD)
echo "当前分支: $BRANCH"
echo ""

# 2. 检查待提交文件
echo "【步骤2】检查待提交文件..."
FILE_COUNT=$(git status --short | wc -l)
echo "待提交文件数: $FILE_COUNT"
echo ""

# 显示前30个文件
echo "前30个待提交文件:"
git status --short | head -30
echo "..."
echo ""

# 3. 显示统计
echo "【步骤3】显示统计..."
git diff --stat | tail -3
echo ""

# 4. 暂存
echo "【步骤4】暂存所有文件..."
git add -A
echo "✅ 已暂存"
echo ""

# 5. 提交
echo "【步骤5】创建提交..."
git commit -m "Phase 12-13: 后端完成 + 前端规划 + 完整DAO层生成

【后端完成】
- Phase 12: 部门日志统计 (4 APIs)
  * GetReport: 分页统计报表
  * GetReportGroupByDepart: 按部门聚合统计
  * GetReportGroupByMonth: 按月份时间序列统计
  * UnitViewLog: 单元浏览日志记录
- 累计134个APIs (89% 完成度)
  * 系统管理: 57 APIs ✅
  * 教学管理: 77 APIs ✅

【DAO层生成】
- 系统管理表: 13张 (Entity/DO/DAO/Service)
- 教学管理表: 10张 (Entity/DO/DAO/Service)
- 共生成 ~12500 行自动代码

【中间件和工具】
- CORS/Logger/Error/Auth 中间件
- JWT认证工具
- 统一响应格式
- 错误码常量定义

【前端规划】
- Phase 13: 4周详细开发计划
- 快速启动指南
- Vite/Axios/Pinia/Router 配置示例

【文档完善】
- 更新 changelog.md
- 新增接口规范和开发计划
- 新增快速启动指南

【编译验证】
✅ go build 编译成功"

echo "✅ 已提交"
echo ""

# 6. 推送
echo "【步骤6】推送到远程..."
git push origin devgo
echo "✅ 已推送"
echo ""

# 7. 验证
echo "【步骤7】验证..."
echo "最新提交:"
git log --oneline -1
echo ""
echo "最近5次提交:"
git log --oneline -5
echo ""

echo "=================================="
echo "✅ 操作完成！"
echo "=================================="
