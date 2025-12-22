#!/bin/bash

# 生成DAO层代码的脚本
cd /workspaces/teaching-open/api-go

echo "开始生成DAO层代码..."
echo "工作目录: $(pwd)"
echo ""

# 检查gf命令
if ! command -v gf &> /dev/null; then
    echo "错误: gf命令未找到"
    echo "请先安装GoFrame CLI: go install github.com/gogf/gf/cmd/gf/v2@latest"
    exit 1
fi

echo "GoFrame CLI版本:"
gf version
echo ""

echo "配置文件检查: hack/config.yaml"
if [ -f "hack/config.yaml" ]; then
    echo "✓ 配置文件存在"
    echo "数据库连接: mysql:root:root@tcp(127.0.0.1:3306)/teachingopen"
else
    echo "✗ 配置文件不存在"
    exit 1
fi
echo ""

echo "执行: gf gen dao"
gf gen dao

echo ""
echo "检查生成的文件..."
echo ""

echo "DAO文件:"
ls -lh internal/dao/*.go 2>/dev/null || echo "无DAO文件生成"
echo ""

echo "DO文件:"
ls -lh internal/model/do/*.go 2>/dev/null || echo "无DO文件生成"
echo ""

echo "Entity文件:"
ls -lh internal/model/entity/*.go 2>/dev/null || echo "无Entity文件生成"
echo ""

echo "完成!"
