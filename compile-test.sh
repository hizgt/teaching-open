#!/bin/bash

echo "=== GoFrame项目编译测试 ==="
echo ""

# 步骤1: 清理临时文件
echo "步骤1: 清理临时文件..."
cd /root/teaching
rm -f api-go/internal/service/sys_user_new.go
rm -f api-go/internal/service/sys_user_service.go
rm -f api-go/utility/response/res.go
echo "✓ 临时文件清理完成"
echo ""

# 步骤2: 整理依赖
echo "步骤2: 整理依赖 (go mod tidy)..."
cd /root/teaching
cd api-go
go mod tidy
if [ $? -eq 0 ]; then
    echo "✓ go mod tidy 成功"
else
    echo "✗ go mod tidy 失败"
    exit 1
fi
echo ""

# 步骤3: 编译检查
echo "步骤3: 编译检查 (go build -v)..."
go build -v
BUILD_STATUS=$?
echo ""

if [ $BUILD_STATUS -eq 0 ]; then
    echo "=== 编译成功 ==="
    echo ""
    echo "生成的可执行文件信息:"
    ls -lh teaching-open 2>/dev/null || ls -lh api-go 2>/dev/null || echo "未找到可执行文件"
    echo ""
    echo "文件大小和路径:"
    if [ -f "teaching-open" ]; then
        FILE_SIZE=$(stat -c%s "teaching-open" 2>/dev/null || stat -f%z "teaching-open" 2>/dev/null)
        echo "- 文件: $PWD/teaching-open"
        echo "- 大小: $FILE_SIZE 字节 ($(numfmt --to=iec-i --suffix=B $FILE_SIZE 2>/dev/null || echo "$FILE_SIZE bytes"))"
    fi
else
    echo "=== 编译失败 ==="
    echo ""
    echo "请查看上面的错误信息"
    exit 1
fi
