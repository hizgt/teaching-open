#!/bin/bash

# Git提交脚本 - Phase 12完成 + 前端规划 + DAO层生成

set -e

echo "================================================"
echo "Git 提交脚本"
echo "================================================"
echo ""

# 获取当前分支
CURRENT_BRANCH=$(git rev-parse --abbrev-ref HEAD)
echo "📍 当前分支: $CURRENT_BRANCH"
echo ""

# 显示变更统计
echo "📊 变更统计:"
git diff --stat HEAD
echo ""

# 查看待提交文件
echo "📝 待提交的文件:"
git status --short | head -20
echo "... (共 $(git status --short | wc -l) 个文件)"
echo ""

# 确认提交
read -p "确认提交? (y/n) " -n 1 -r
echo
if [[ ! $REPLY =~ ^[Yy]$ ]]; then
    echo "❌ 取消提交"
    exit 1
fi

# 暂存所有文件
echo ""
echo "⏳ 正在暂存所有文件..."
git add -A
echo "✅ 文件已暂存"

# 提交
echo ""
echo "⏳ 正在提交..."
git commit -m "Phase 12-13: 后端完成 + 前端规划 + 完整DAO层生成

【后端完成】
- Phase 12: 部门日志统计 (4 APIs)
  * GetReport: 分页统计报表
  * GetReportGroupByDepart: 按部门聚合统计
  * GetReportGroupByMonth: 按月份时间序列统计
  * UnitViewLog: 单元浏览日志记录
  
- 累计完成134个APIs (89%)
  * 系统管理: 57 APIs ✅
  * 教学管理: 77 APIs ✅

【DAO层生成】
- 生成所有系统管理表的Entity、DO、DAO文件:
  * 系统管理: sys_user, sys_role, sys_permission, sys_depart, 
             sys_dict, sys_log, sys_file, sys_data_log, 
             sys_user_role, sys_user_depart, sys_role_permission
  * 教学管理: teaching_course, teaching_course_unit, teaching_work,
             teaching_additional_work, teaching_news, teaching_scratch_asset,
             teaching_depart_day_log, teaching_work_correct, teaching_work_comment,
             teaching_course_dept

【中间件系统】
- CORS跨域中间件
- Logger日志中间件
- Error错误处理中间件
- Auth认证中间件

【工具类】
- JWT认证工具
- 统一响应格式
- 错误码常量定义

【文档更新】
- docs/changelog.md: 更新Phase 12-13进度
- docs/20251203/前后端接口报告.md: 接口统计和规范
- docs/20251203/未完成工作报告.md: 详细待办事项
- PHASE12_FIX_GUIDE.md: Phase 12修复指南
- api-go/WORK_SUMMARY.md: 工作总结
- api-go/DEVELOPMENT_STATUS.md: 开发状态

【前端规划】
- 创建Phase 13 Vue3开发计划
- 创建PHASE13_QUICK_START.md快速启动指南
- 创建init-phase13-vue3.sh初始化脚本
- 提供Vite/Axios/Pinia/Router配置示例

【编译验证】
- go build 编译成功 ✅

【Git分支】
- 当前分支: devgo
- 更新记录已同步

变更统计:
- 修改文件: ~10 个
- 新增文件: ~130 个
- 总代码行数: ~15000+ 行"

echo ""
echo "✅ 提交成功！"
echo ""

# 显示提交信息
echo "📝 最新提交:"
git log --oneline -1
echo ""

# 显示待push的提交
echo "📤 待推送的提交:"
git log --oneline origin/$CURRENT_BRANCH..$CURRENT_BRANCH 2>/dev/null || echo "本地分支与远程分支相同,无待推送提交"
echo ""

# 提示推送
echo "💡 提示: 使用 'git push origin $CURRENT_BRANCH' 推送到远程仓库"
echo ""
