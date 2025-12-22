#!/bin/bash

# 使用 GoFrame CLI 生成字典相关表的 DAO 层代码
# 使用方法: chmod +x gen-dict-dao.sh && ./gen-dict-dao.sh

echo "================================================"
echo "  GoFrame CLI - 生成字典管理 DAO 层代码"
echo "================================================"
echo ""

cd /root/teaching/api-go

# 检查 gf 命令
if ! command -v gf &> /dev/null; then
    echo "错误: gf 命令未找到，请先安装 GoFrame CLI"
    echo "安装命令: go install github.com/gogf/gf/cmd/gf/v2@latest"
    exit 1
fi

echo "GoFrame CLI 版本:"
gf version
echo ""

# 检查数据库连接配置
echo "数据库配置 (hack/config.yaml):"
cat hack/config.yaml | grep -A 10 "dao:"
echo ""

# 先备份现有文件（如果存在）
echo "备份现有 DAO 文件..."
BACKUP_DIR="backup_dao_$(date +%Y%m%d_%H%M%S)"
mkdir -p $BACKUP_DIR

# 备份字典相关文件
for file in internal/dao/sys_dict.go internal/dao/sys_dict_item.go \
            internal/dao/internal/sys_dict.go internal/dao/internal/sys_dict_item.go \
            internal/model/entity/sys_dict.go internal/model/entity/sys_dict_item.go \
            internal/model/do/sys_dict.go internal/model/do/sys_dict_item.go; do
    if [ -f "$file" ]; then
        cp "$file" "$BACKUP_DIR/"
        echo "  已备份: $file"
    fi
done
echo ""

# 生成 sys_dict 表的 DAO
echo "生成 sys_dict 表 DAO..."
gf gen dao -tables="sys_dict" -removePrefix="" -jsonCase="CamelLower"

# 生成 sys_dict_item 表的 DAO
echo "生成 sys_dict_item 表 DAO..."
gf gen dao -tables="sys_dict_item" -removePrefix="" -jsonCase="CamelLower"

echo ""
echo "================================================"
echo "  DAO 生成完成!"
echo "================================================"
echo ""

# 显示生成的文件
echo "生成的文件列表:"
echo "Entity 文件:"
ls -la internal/model/entity/sys_dict*.go 2>/dev/null || echo "  未找到"

echo ""
echo "DO 文件:"
ls -la internal/model/do/sys_dict*.go 2>/dev/null || echo "  未找到"

echo ""
echo "DAO Internal 文件:"
ls -la internal/dao/internal/sys_dict*.go 2>/dev/null || echo "  未找到"

echo ""
echo "DAO 文件:"
ls -la internal/dao/sys_dict*.go 2>/dev/null || echo "  未找到"

echo ""
echo "验证编译..."
go build ./... && echo "✅ 编译通过" || echo "❌ 编译失败"

echo ""
echo "备份目录: $BACKUP_DIR"
echo "如需回滚，可从备份目录恢复文件"
