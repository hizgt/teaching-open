#!/bin/bash

# 清理和重建工具类文件

cd /root/teaching/api-go

echo "=== 清理损坏的文件 ==="

# 删除所有损坏的工具类文件
rm -f utility/response/response.go
rm -f utility/response/resp.go
rm -f utility/jwt/jwt.go
rm -f utility/jwt/token.go
rm -f api/middleware/auth.go
rm -f api/middleware/auth_middleware.go

echo "✓ 清理完成"
echo ""
echo "接下来需要重新创建以下文件:"
echo "- utility/response/response.go"
echo "- utility/jwt/jwt.go"
echo "- api/middleware/auth.go"
echo ""
echo "请使用 IDE 或编辑器重新创建这些文件"
