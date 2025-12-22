// =================================================================================
// Service interface for sys log module
// =================================================================================

package service

import (
	"context"

	v1 "teaching-open/api/v1/sys"
)

// ISysLog 系统日志服务接口
type ISysLog interface {
	// GetList 获取系统日志列表
	GetList(ctx context.Context, req *v1.LogListReq) (list []*v1.LogInfo, total int64, err error)
	// Delete 删除系统日志
	Delete(ctx context.Context, id string) error
	// DeleteBatch 批量删除系统日志
	DeleteBatch(ctx context.Context, ids string) error
	// Clear 清空系统日志
	Clear(ctx context.Context, logType int) (count int64, err error)
	// Add 添加系统日志（用于中间件记录日志）
	Add(ctx context.Context, logInfo *v1.LogInfo) error
	// GetDataLogList 获取数据日志列表
	GetDataLogList(ctx context.Context, req *v1.DataLogListReq) (list []*v1.DataLogInfo, total int64, err error)
	// GetDataLogById 根据ID获取数据日志详情
	GetDataLogById(ctx context.Context, id string) (*v1.DataLogInfo, error)
	// GetDataLogHistory 获取数据变更历史
	GetDataLogHistory(ctx context.Context, dataTable string, dataId string) ([]*v1.DataLogInfo, error)
	// AddDataLog 添加数据日志（用于记录数据变更）
	AddDataLog(ctx context.Context, dataTable string, dataId string, dataContent string) error
}

var localSysLog ISysLog

// SysLog 获取系统日志服务实例
func SysLog() ISysLog {
	if localSysLog == nil {
		panic("implement not found for interface ISysLog, forgot register?")
	}
	return localSysLog
}

// RegisterSysLog 注册系统日志服务实现
func RegisterSysLog(i ISysLog) {
	localSysLog = i
}
