#!/bin/bash

# Git 提交脚本 - 文件上传管理模块
# 使用方法: chmod +x git-commit.sh && ./git-commit.sh

echo "================================================"
echo "  Git提交 - 文件上传管理模块"
echo "================================================"
echo ""

cd /root/teaching

echo "当前分支: $(git branch --show-current)"
echo ""
echo "查看Git状态..."
git status --short
echo ""

echo "添加文件到暂存区..."
git add api-go/
git add docs/
git add *.sh

echo "执行Git提交..."
git commit -m "feat: 实现文件上传管理模块 (8个API接口)

【新增功能】
文件上传管理:
- 单文件上传 (POST /sys/file/upload)
- 批量上传 (POST /sys/file/uploadBatch)
- 文件列表 (GET /sys/file/list) - 分页、类型筛选、标签筛选
- 文件详情 (GET /sys/file/queryById)
- 删除文件 (DELETE /sys/file/delete) - 逻辑删除+物理删除
- 批量删除 (DELETE /sys/file/deleteBatch)
- 文件预览 (GET /sys/file/view/:id)
- 文件下载 (GET /sys/file/download/:id)

【DAO层】
- sys_file表: Entity/DO/DAO (12字段)

【技术实现】
- 文件类型自动识别: 图片(1)/文档(2)/视频(3)/音频(4)/压缩包(5)/其他(0)
- 文件大小限制: 默认50MB，可配置
- 存储位置: 本地存储(1)，支持扩展云存储
- 文件命名: UUID + 原始扩展名
- 目录结构: 按日期分目录存储（YYYY/MM/DD）
- 逻辑删除: 文件记录逻辑删除，同时删除物理文件

【文档更新】
- 前后端接口报告: 57个已完成接口
- changelog: 文件上传模块记录 (第七阶段)
- 未完成工作报告: 进度更新
"

echo ""
echo "提交完成!"
git log -1 --oneline
echo ""
echo "================================================"
