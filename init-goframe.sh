#!/bin/bash

# Teaching Open GoFrame 项目初始化脚本

set -e

echo "================================"
echo "Teaching Open 项目初始化"
echo "================================"
echo ""

# 1. 检查并安装 GoFrame CLI 工具
echo "步骤 1: 检查 GoFrame CLI 工具..."
if ! command -v gf &> /dev/null; then
    echo "GoFrame CLI 未安装，正在安装..."
    go install github.com/gogf/gf/cmd/gf/v2@latest
    echo "✓ GoFrame CLI 安装完成"
else
    echo "✓ GoFrame CLI 已安装"
    gf version
fi

echo ""

# 2. 创建项目目录
echo "步骤 2: 创建项目目录..."
cd /workspaces/teaching-open
if [ -d "api-go" ]; then
    echo "检测到已存在的 api-go 目录，正在备份..."
    mv api-go api-go.backup.$(date +%Y%m%d%H%M%S)
fi

echo ""

# 3. 使用 gf init 初始化项目
echo "步骤 3: 初始化 GoFrame 项目..."
echo "项目名称: teaching-open"
echo "模式: 单仓模式"
echo ""

# 创建项目目录
mkdir -p api-go
cd api-go

# 初始化 Go 模块
go mod init teaching-open

# 使用 gf init 初始化（如果 gf 命令可用）
if command -v gf &> /dev/null; then
    echo "使用 gf init 初始化项目结构..."
    # gf init 会交互式询问，这里我们使用默认的单仓模式
    # 由于是自动化脚本，我们跳过交互式初始化，手动创建标准结构
    echo "创建标准 GoFrame 项目结构..."
else
    echo "警告: gf 命令不可用，将创建基础项目结构..."
fi

# 4. 创建标准目录结构
echo ""
echo "步骤 4: 创建项目目录结构..."

# API层
mkdir -p api/v1/{system,teaching}
mkdir -p api/middleware

# Internal层
mkdir -p internal/{cmd,controller/{system,teaching},logic/{system,teaching},dao,model/{entity,do,vo},service,consts}

# Manifest层
mkdir -p manifest/{config,docker,deploy/k8s}

# Utility层
mkdir -p utility/{response,jwt}

# Resource层
mkdir -p resource/{public,template,log,upload}

# Hack层（用于开发工具配置）
mkdir -p hack

echo "✓ 目录结构创建完成"

# 5. 添加 GoFrame 依赖
echo ""
echo "步骤 5: 添加项目依赖..."
go get -u github.com/gogf/gf/v2@latest
go get -u github.com/golang-jwt/jwt/v5@latest
go get -u golang.org/x/crypto@latest

echo "✓ 依赖安装完成"

# 6. 整理依赖
echo ""
echo "步骤 6: 整理依赖..."
go mod tidy

echo ""
echo "================================"
echo "✓ 项目初始化完成！"
echo "================================"
echo ""
echo "项目路径: /workspaces/teaching-open/api-go"
echo ""
echo "下一步操作:"
echo "1. cd api-go"
echo "2. 配置 manifest/config/config.yaml"
echo "3. 配置 hack/config.yaml (用于 gf gen dao)"
echo "4. 运行 gf gen dao 生成数据访问层"
echo "5. 运行 go run main.go 启动服务"
echo ""
