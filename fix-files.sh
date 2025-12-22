#!/bin/bash

# 清理损坏的文件并测试编译

cd /workspaces/teaching-open/api-go

echo "=== 清理损坏的文件 ==="

# 删除损坏的utility文件
rm -f utility/response/response.go
rm -f utility/jwt/jwt.go
rm -f api/middleware/auth.go

# 重命名正确的文件
if [ -f "utility/response/resp.go" ]; then
    mv utility/response/resp.go utility/response/response.go
    echo "✓ 恢复 response.go"
fi

if [ -f "utility/jwt/token.go" ]; then
    mv utility/jwt/token.go utility/jwt/jwt.go
    echo "✓ 恢复 jwt.go"
fi

echo ""
echo "=== 测试编译 ==="
go mod tidy
go build -o temp_build

if [ $? -eq 0 ]; then
    echo "✓ 编译成功！"
    rm -f temp_build
else
    echo "✗ 编译失败，请检查错误"
fi
