#!/usr/bin/env python3
# -*- coding: utf-8 -*-

import os
import subprocess
import sys
from datetime import datetime

def run_command(cmd, description=""):
    """执行shell命令"""
    print(f"\n{'='*60}")
    if description:
        print(f"▶️  {description}")
    print(f"⏳ 执行: {cmd}")
    print(f"{'='*60}")
    
    try:
        result = subprocess.run(cmd, shell=True, capture_output=True, text=True, cwd="/root/teaching")
        
        if result.stdout:
            print(result.stdout)
        
        if result.returncode != 0:
            print(f"❌ 错误: {result.stderr}")
            return False
        
        print(f"✅ 成功")
        return True
    
    except Exception as e:
        print(f"❌ 异常: {str(e)}")
        return False

def main():
    os.chdir("/root/teaching")
    
    print("""
╔══════════════════════════════════════════════════════╗
║        🚀 Git 自动提交脚本                           ║
║        Phase 12-13: 后端完成 + 前端规划              ║
╚══════════════════════════════════════════════════════╝
    """)
    
    # 1. 检查分支
    print("\n【步骤1】检查当前分支...")
    result = subprocess.run("git rev-parse --abbrev-ref HEAD", shell=True, capture_output=True, text=True)
    current_branch = result.stdout.strip()
    print(f"当前分支: {current_branch}")
    
    if current_branch != "devgo":
        print(f"⚠️  警告: 当前分支是 {current_branch}，预期是 devgo")
        response = input("是否继续? (y/n): ")
        if response.lower() != 'y':
            print("❌ 已取消")
            return False
    
    # 2. 查看待提交文件
    print("\n【步骤2】查看待提交文件...")
    result = subprocess.run("git status --short | wc -l", shell=True, capture_output=True, text=True)
    file_count = result.stdout.strip()
    print(f"待提交文件数: {file_count}")
    
    result = subprocess.run("git status --short | head -30", shell=True, capture_output=True, text=True)
    print(result.stdout)
    
    # 3. 显示变更统计
    print("\n【步骤3】变更统计...")
    result = subprocess.run("git diff --stat | tail -1", shell=True, capture_output=True, text=True)
    print(result.stdout)
    
    # 4. 确认提交
    print("\n【步骤4】确认提交...")
    response = input("\n确认要提交这些变更吗? (y/n): ")
    if response.lower() != 'y':
        print("❌ 已取消")
        return False
    
    # 5. 暂存所有文件
    if not run_command("git add -A", "暂存所有文件"):
        return False
    
    # 6. 验证暂存
    print("\n【步骤5】验证暂存...")
    result = subprocess.run("git status --short | head -10", shell=True, capture_output=True, text=True)
    print(result.stdout)
    
    # 7. 提交
    commit_message = """Phase 12-13: 后端完成 + 前端规划 + 完整DAO层生成

【后端完成】
- Phase 12: 部门日志统计 (4 APIs)
  * GetReport: 分页统计报表
  * GetReportGroupByDepart: 按部门聚合统计
  * GetReportGroupByMonth: 按月份时间序列统计
  * UnitViewLog: 单元浏览日志记录
- 累计134个APIs (89% 完成度)

【DAO层生成】
- 系统管理表: 13张 (13 Entity + 13 DO + 13 DAO + 11 Service)
- 教学管理表: 10张 (10 Entity + 10 DO + 10 DAO + 10 Service)
- 共生成 ~12500 行自动代码

【中间件和工具】
- CORS/Logger/Error/Auth 4个中间件
- JWT认证工具
- 统一响应格式
- 错误码常量定义

【前端规划】
- Phase 13: 4周详细开发计划
- 快速启动指南 (100KB)
- Vite/Axios/Pinia/Router 配置示例
- 初始化脚本

【文档完善】
- 更新 changelog.md
- 新增接口规范、开发计划、工作报告
- 新增快速启动指南和完成报告

【编译验证】
✅ go build 编译成功"""
    
    cmd = f'git commit -m "{commit_message}"'
    if not run_command(cmd, "提交变更"):
        return False
    
    # 8. 推送到远程
    print("\n【步骤6】推送到远程...")
    response = input("\n是否要推送到远程? (y/n): ")
    if response.lower() == 'y':
        if not run_command("git push origin devgo", "推送到远程"):
            print("⚠️  推送失败，但本地提交已完成")
            return False
    
    # 9. 验证提交
    print("\n【步骤7】验证提交...")
    result = subprocess.run("git log --oneline -1", shell=True, capture_output=True, text=True)
    print(f"最新提交: {result.stdout}")
    
    result = subprocess.run("git log --oneline -5", shell=True, capture_output=True, text=True)
    print(f"\n最近5次提交:\n{result.stdout}")
    
    print("\n" + "="*60)
    print("✅ 操作完成！")
    print("="*60)
    print("\n下一步:")
    print("1. 验证远程仓库: git branch -v")
    print("2. 查看提交: git log --oneline")
    print("3. 开始Phase 13: bash init-phase13-vue3.sh")
    
    return True

if __name__ == "__main__":
    try:
        success = main()
        sys.exit(0 if success else 1)
    except KeyboardInterrupt:
        print("\n\n❌ 用户取消操作")
        sys.exit(1)
    except Exception as e:
        print(f"\n❌ 发生错误: {str(e)}")
        sys.exit(1)
