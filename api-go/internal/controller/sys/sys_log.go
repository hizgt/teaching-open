// =================================================================================
// Controller for sys log module
// =================================================================================

package sys

import (
	"context"

	"github.com/gogf/gf/v2/errors/gerror"

	v1 "teaching-open/api/v1/sys"
	"teaching-open/internal/service"
)

type cLog struct{}

var SysLog = &cLog{}

// GetList 获取系统日志列表
func (c *cLog) GetList(ctx context.Context, req *v1.LogListReq) (res *v1.LogListRes, err error) {
	// 设置默认值
	if req.Page <= 0 {
		req.Page = 1
	}
	if req.PageSize <= 0 {
		req.PageSize = 10
	}

	list, total, err := service.SysLog().GetList(ctx, req)
	if err != nil {
		return nil, gerror.Wrap(err, "查询系统日志列表失败")
	}

	return &v1.LogListRes{
		List:     list,
		Total:    total,
		Page:     req.Page,
		PageSize: req.PageSize,
	}, nil
}

// Delete 删除系统日志
func (c *cLog) Delete(ctx context.Context, req *v1.LogDeleteReq) (res *v1.LogDeleteRes, err error) {
	err = service.SysLog().Delete(ctx, req.Id)
	if err != nil {
		return nil, gerror.Wrap(err, "删除系统日志失败")
	}

	return &v1.LogDeleteRes{}, nil
}

// DeleteBatch 批量删除系统日志
func (c *cLog) DeleteBatch(ctx context.Context, req *v1.LogDeleteBatchReq) (res *v1.LogDeleteBatchRes, err error) {
	err = service.SysLog().DeleteBatch(ctx, req.Ids)
	if err != nil {
		return nil, gerror.Wrap(err, "批量删除系统日志失败")
	}

	return &v1.LogDeleteBatchRes{}, nil
}

// Clear 清空系统日志
func (c *cLog) Clear(ctx context.Context, req *v1.LogClearReq) (res *v1.LogClearRes, err error) {
	count, err := service.SysLog().Clear(ctx, req.LogType)
	if err != nil {
		return nil, gerror.Wrap(err, "清空系统日志失败")
	}

	return &v1.LogClearRes{
		Count: count,
	}, nil
}

// GetDataLogList 获取数据日志列表
func (c *cLog) GetDataLogList(ctx context.Context, req *v1.DataLogListReq) (res *v1.DataLogListRes, err error) {
	// 设置默认值
	if req.Page <= 0 {
		req.Page = 1
	}
	if req.PageSize <= 0 {
		req.PageSize = 10
	}

	list, total, err := service.SysLog().GetDataLogList(ctx, req)
	if err != nil {
		return nil, gerror.Wrap(err, "查询数据日志列表失败")
	}

	return &v1.DataLogListRes{
		List:     list,
		Total:    total,
		Page:     req.Page,
		PageSize: req.PageSize,
	}, nil
}

// GetDataLogById 查询数据日志详情
func (c *cLog) GetDataLogById(ctx context.Context, req *v1.DataLogGetByIdReq) (res *v1.DataLogGetByIdRes, err error) {
	info, err := service.SysLog().GetDataLogById(ctx, req.Id)
	if err != nil {
		return nil, gerror.Wrap(err, "查询数据日志详情失败")
	}

	return &v1.DataLogGetByIdRes{
		DataLogInfo: info,
	}, nil
}

// GetDataLogHistory 获取数据变更历史
func (c *cLog) GetDataLogHistory(ctx context.Context, req *v1.DataLogHistoryReq) (res *v1.DataLogHistoryRes, err error) {
	list, err := service.SysLog().GetDataLogHistory(ctx, req.DataTable, req.DataId)
	if err != nil {
		return nil, gerror.Wrap(err, "查询数据变更历史失败")
	}

	return &v1.DataLogHistoryRes{
		List: list,
	}, nil
}
