#!/bin/bash
# 修复teaching_depart_day_log.go文件，只保留前308行

set -e

echo "开始修复teaching_depart_day_log.go文件..."

cd /workspaces/teaching-open/api-go/internal/logic/sys

# 备份原文件
cp teaching_depart_day_log.go teaching_depart_day_log.go.bak

# 只保留前308行
head -308 teaching_depart_day_log.go.bak > teaching_depart_day_log.go

# 显示结果
echo "文件已修复！"
echo "当前行数: $(wc -l < teaching_depart_day_log.go)"
echo "原文件备份: teaching_depart_day_log.go.bak"

# 验证编译
echo ""
echo "验证编译..."
cd /workspaces/teaching-open/api-go
if go build -o /tmp/teaching-open-test 2>&1; then
    echo "✅ 编译成功!"
    rm -f /tmp/teaching-open-test
else
    echo "❌ 编译失败，请检查错误"
    exit 1
fi

echo ""
echo "修复完成! 现在可以提交代码了"
echo "运行: cd /workspaces/teaching-open && git add . && git commit -m 'fix: Phase 12 logic文件格式修复' && git push origin devgo"
