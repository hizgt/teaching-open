#!/bin/bash

# Git 提交脚本

cd /workspaces/teaching-open

echo "正在提交 GoFrame 项目初始化相关文件..."

# 添加所有变更
git add .

# 提交
git commit -m "feat: GoFrame 项目初始化准备

- 删除手动创建的项目结构
- 创建 GoFrame 项目初始化脚本 (init-goframe.sh, rebuild-project.sh)
- 创建项目初始化指南 (api-go/INIT_GUIDE.md)
- 更新 changelog.md，记录 v3.0.0-dev 版本信息
- 更新未完成工作报告，标记项目初始化进度
- 创建前后端接口报告模板 (docs/20251122/前后端接口报告.md)

下一步需要手动执行:
1. go install github.com/gogf/gf/cmd/gf/v2@latest
2. cd /workspaces/teaching-open && gf init api-go
3. 选择单仓模式和 web 模板
"

echo "✓ 提交完成"
