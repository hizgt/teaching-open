#!/bin/bash

# 启动数据库并使用 GoFrame CLI 生成 DAO 层代码
# 使用方法: chmod +x setup-db-gen-dao.sh && ./setup-db-gen-dao.sh

set -e

echo "================================================"
echo "  Teaching Open - 数据库启动与 DAO 生成"
echo "================================================"
echo ""

cd /root/teaching/api-go

# ==================== 第一步：启动数据库 ====================
echo "[1/5] 检查并启动 MySQL 数据库..."
echo ""

# 检查 Docker 是否可用
if ! command -v docker &> /dev/null; then
    echo "错误: Docker 未安装或不可用"
    exit 1
fi

# 检查 MySQL 是否已运行
if docker ps | grep -q "teaching-open-mysql"; then
    echo "✅ MySQL 容器已运行"
else
    echo "启动 MySQL 容器..."
    cd manifest/docker
    docker-compose up -d mysql
    cd ../..
    
    echo "等待 MySQL 启动..."
    sleep 10
    
    # 等待 MySQL 就绪
    for i in {1..30}; do
        if docker exec teaching-open-mysql mysqladmin ping -uroot -proot &>/dev/null; then
            echo "✅ MySQL 已就绪"
            break
        fi
        echo "  等待中... ($i/30)"
        sleep 2
    done
fi

# ==================== 第二步：检查数据库 ====================
echo ""
echo "[2/5] 检查数据库表..."
echo ""

# 检查表是否存在
TABLES=$(docker exec teaching-open-mysql mysql -uroot -proot -N -e "USE teachingopen; SHOW TABLES LIKE 'sys_dict%';" 2>/dev/null || echo "")
if [ -z "$TABLES" ]; then
    echo "数据库表不存在，需要导入 SQL 文件..."
    echo "导入 teachingopen2.8.sql..."
    docker exec -i teaching-open-mysql mysql -uroot -proot teachingopen < /root/teaching/api/db/teachingopen2.8.sql
    echo "✅ SQL 导入完成"
else
    echo "✅ 数据库表已存在:"
    echo "$TABLES"
fi

# ==================== 第三步：备份现有文件 ====================
echo ""
echo "[3/5] 备份现有 DAO 文件..."
echo ""

BACKUP_DIR="backup_dao_$(date +%Y%m%d_%H%M%S)"
mkdir -p $BACKUP_DIR

for file in internal/dao/sys_dict.go internal/dao/sys_dict_item.go \
            internal/dao/internal/sys_dict.go internal/dao/internal/sys_dict_item.go \
            internal/model/entity/sys_dict.go internal/model/entity/sys_dict_item.go \
            internal/model/do/sys_dict.go internal/model/do/sys_dict_item.go; do
    if [ -f "$file" ]; then
        cp "$file" "$BACKUP_DIR/"
        echo "  备份: $file"
    fi
done

# ==================== 第四步：生成 DAO ====================
echo ""
echo "[4/5] 使用 GoFrame CLI 生成 DAO 层代码..."
echo ""

# 检查 gf 命令
if ! command -v gf &> /dev/null; then
    echo "安装 GoFrame CLI..."
    go install github.com/gogf/gf/cmd/gf/v2@latest
fi

echo "GoFrame CLI 版本: $(gf version | head -1)"
echo ""

# 生成字典相关表的 DAO
echo "生成 sys_dict 和 sys_dict_item 表的 DAO..."
gf gen dao -tables="sys_dict,sys_dict_item"

echo ""
echo "✅ DAO 生成完成"

# ==================== 第五步：验证 ====================
echo ""
echo "[5/5] 验证生成结果..."
echo ""

echo "生成的文件:"
echo "----------------------------------------"
echo "Entity 文件:"
ls -la internal/model/entity/sys_dict*.go 2>/dev/null || echo "  (无)"

echo ""
echo "DO 文件:"
ls -la internal/model/do/sys_dict*.go 2>/dev/null || echo "  (无)"

echo ""
echo "DAO Internal 文件:"
ls -la internal/dao/internal/sys_dict*.go 2>/dev/null || echo "  (无)"

echo ""
echo "DAO 封装文件:"
ls -la internal/dao/sys_dict*.go 2>/dev/null || echo "  (无)"

echo ""
echo "----------------------------------------"
echo "验证编译..."
if go build ./...; then
    echo "✅ 编译通过"
else
    echo "❌ 编译失败，可能需要手动修复"
    echo "备份目录: $BACKUP_DIR"
fi

echo ""
echo "================================================"
echo "  完成!"
echo "================================================"
echo ""
echo "数据库连接信息:"
echo "  Host: 127.0.0.1"
echo "  Port: 3306"
echo "  User: root"
echo "  Pass: root"
echo "  DB:   teachingopen"
echo ""
echo "备份目录: $BACKUP_DIR"
