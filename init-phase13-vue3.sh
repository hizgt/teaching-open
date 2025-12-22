#!/bin/bash
# Phase 13 Vue3项目初始化启动脚本

set -e

echo "========================================="
echo "Phase 13: Vue3 Frontend项目初始化"
echo "========================================="

# 1. 创建Vue3项目
echo ""
echo "步骤1: 创建Vue3+Vite项目..."
cd /workspaces/teaching-open
npm create vite@latest web-vue3 -- --template vue-ts

# 2. 进入项目目录
cd web-vue3

# 3. 安装核心依赖
echo ""
echo "步骤2: 安装核心依赖..."
npm install

echo ""
echo "步骤3: 安装额外依赖..."
npm install vue-router@4 pinia@2 ant-design-vue@4 axios@1.6 dayjs @vueuse/core
npm install -D @types/node unplugin-vue-components unplugin-auto-import

# 4. 创建目录结构
echo ""
echo "步骤4: 创建目录结构..."
mkdir -p src/api
mkdir -p src/assets/{images,styles}
mkdir -p src/components
mkdir -p src/composables
mkdir -p src/hooks
mkdir -p src/layouts
mkdir -p src/router
mkdir -p src/stores
mkdir -p src/types
mkdir -p src/utils
mkdir -p src/views/{system,teaching,home,login}

echo ""
echo "✅ Phase 13项目初始化完成!"
echo ""
echo "下一步:"
echo "1. cd web-vue3"
echo "2. 复制配置文件: vite.config.ts, tsconfig.json"
echo "3. 创建基础代码: utils/request.ts, stores/user.ts, router/index.ts"
echo "4. npm run dev 启动开发服务器"
echo ""
