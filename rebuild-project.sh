#!/bin/bash

# 重建 GoFrame 项目

cd /workspaces/teaching-open

# 删除旧目录
rm -rf api-go

# 使用 gf init 初始化项目
echo "请手动执行以下命令来初始化 GoFrame 项目："
echo ""
echo "cd /workspaces/teaching-open"
echo "gf init api-go"
echo ""
echo "然后按照提示选择:"
echo "- 模式: 单仓(mono-repo)"
echo "- 模板: web"
echo ""
echo "完成后项目将创建在 /workspaces/teaching-open/api-go 目录下"
