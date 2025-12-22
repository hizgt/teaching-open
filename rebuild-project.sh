#!/bin/bash

# 重建 GoFrame 项目

cd /root/teaching

# 删除旧目录
rm -rf api-go

# 使用 gf init 初始化项目
echo "请手动执行以下命令来初始化 GoFrame 项目："
echo ""
echo "cd /root/teaching"
echo "gf init api-go"
echo ""
echo "然后按照提示选择:"
echo "- 模式: 单仓(mono-repo)"
echo "- 模板: web"
echo ""
echo "完成后项目将创建在 /root/teaching/api-go 目录下"
