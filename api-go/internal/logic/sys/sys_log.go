// =================================================================================
// Logic implementation for sys log module
// =================================================================================

package sys

import (
	"context"
	"errors"
	"strings"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/guid"

	v1 "teaching-open/api/v1/sys"
	"teaching-open/internal/dao"
	"teaching-open/internal/model/do"
	"teaching-open/internal/model/entity"
	"teaching-open/internal/service"
)

type sSysLog struct{}

func init() {
	service.RegisterSysLog(&sSysLog{})
}

// GetList 获取系统日志列表
func (s *sSysLog) GetList(ctx context.Context, req *v1.LogListReq) (list []*v1.LogInfo, total int64, err error) {
	// 构建查询
	m := dao.SysLog.Ctx(ctx)

	// 日志类型筛选
	if req.LogType > 0 {
		m = m.Where(dao.SysLog.Columns().LogType, req.LogType)
	}

	// 操作用户名称模糊查询
	if req.Username != "" {
		m = m.WhereLike(dao.SysLog.Columns().Username, "%"+req.Username+"%")
	}

	// IP模糊查询
	if req.Ip != "" {
		m = m.WhereLike(dao.SysLog.Columns().Ip, "%"+req.Ip+"%")
	}

	// 日志内容模糊查询
	if req.LogContent != "" {
		m = m.WhereLike(dao.SysLog.Columns().LogContent, "%"+req.LogContent+"%")
	}

	// 操作类型筛选
	if req.OperateType > 0 {
		m = m.Where(dao.SysLog.Columns().OperateType, req.OperateType)
	}

	// 时间范围筛选
	if req.StartTime != "" {
		m = m.WhereGTE(dao.SysLog.Columns().CreateTime, req.StartTime)
	}
	if req.EndTime != "" {
		m = m.WhereLTE(dao.SysLog.Columns().CreateTime, req.EndTime+" 23:59:59")
	}

	// 分页查询
	var logs []*entity.SysLog
	var totalInt int
	err = m.Page(req.Page, req.PageSize).
		OrderDesc(dao.SysLog.Columns().CreateTime).
		ScanAndCount(&logs, &totalInt, false)
	if err != nil {
		g.Log().Error(ctx, "查询系统日志列表失败:", err)
		return nil, 0, errors.New("查询系统日志列表失败")
	}
	total = int64(totalInt)

	// 转换为LogInfo
	list = make([]*v1.LogInfo, 0, len(logs))
	for _, log := range logs {
		createTime := ""
		if log.CreateTime != nil {
			createTime = log.CreateTime.String()
		}
		list = append(list, &v1.LogInfo{
			Id:           log.Id,
			LogType:      log.LogType,
			LogContent:   log.LogContent,
			OperateType:  log.OperateType,
			Userid:       log.Userid,
			Username:     log.Username,
			Ip:           log.Ip,
			Method:       log.Method,
			RequestUrl:   log.RequestUrl,
			RequestParam: log.RequestParam,
			RequestType:  log.RequestType,
			CostTime:     log.CostTime,
			CreateTime:   createTime,
		})
	}

	return list, total, nil
}

// Delete 删除系统日志
func (s *sSysLog) Delete(ctx context.Context, id string) error {
	// 检查日志是否存在
	count, err := dao.SysLog.Ctx(ctx).
		Where(dao.SysLog.Columns().Id, id).
		Count()
	if err != nil {
		g.Log().Error(ctx, "查询系统日志失败:", err)
		return errors.New("查询系统日志失败")
	}
	if count == 0 {
		return errors.New("日志不存在")
	}

	// 删除日志
	_, err = dao.SysLog.Ctx(ctx).
		Where(dao.SysLog.Columns().Id, id).
		Delete()
	if err != nil {
		g.Log().Error(ctx, "删除系统日志失败:", err)
		return errors.New("删除系统日志失败")
	}

	g.Log().Infof(ctx, "删除系统日志成功: id=%s", id)
	return nil
}

// DeleteBatch 批量删除系统日志
func (s *sSysLog) DeleteBatch(ctx context.Context, ids string) error {
	if ids == "" {
		return errors.New("日志ID不能为空")
	}

	idList := strings.Split(ids, ",")
	if len(idList) == 0 {
		return errors.New("日志ID不能为空")
	}

	// 批量删除
	_, err := dao.SysLog.Ctx(ctx).
		WhereIn(dao.SysLog.Columns().Id, idList).
		Delete()
	if err != nil {
		g.Log().Error(ctx, "批量删除系统日志失败:", err)
		return errors.New("批量删除系统日志失败")
	}

	g.Log().Infof(ctx, "批量删除系统日志成功: ids=%s", ids)
	return nil
}

// Clear 清空系统日志
func (s *sSysLog) Clear(ctx context.Context, logType int) (count int64, err error) {
	m := dao.SysLog.Ctx(ctx)

	// 如果指定了日志类型，只清空该类型的日志
	if logType > 0 {
		m = m.Where(dao.SysLog.Columns().LogType, logType)
	}

	// 先统计数量
	countInt, err := m.Count()
	if err != nil {
		g.Log().Error(ctx, "统计系统日志数量失败:", err)
		return 0, errors.New("统计系统日志数量失败")
	}
	count = int64(countInt)

	// 执行删除
	if logType > 0 {
		_, err = dao.SysLog.Ctx(ctx).
			Where(dao.SysLog.Columns().LogType, logType).
			Delete()
	} else {
		_, err = dao.SysLog.Ctx(ctx).Delete()
	}
	if err != nil {
		g.Log().Error(ctx, "清空系统日志失败:", err)
		return 0, errors.New("清空系统日志失败")
	}

	g.Log().Infof(ctx, "清空系统日志成功: logType=%d, count=%d", logType, count)
	return count, nil
}

// Add 添加系统日志
func (s *sSysLog) Add(ctx context.Context, logInfo *v1.LogInfo) error {
	// 插入日志
	_, err := dao.SysLog.Ctx(ctx).Insert(do.SysLog{
		Id:           guid.S(),
		LogType:      logInfo.LogType,
		LogContent:   logInfo.LogContent,
		OperateType:  logInfo.OperateType,
		Userid:       logInfo.Userid,
		Username:     logInfo.Username,
		Ip:           logInfo.Ip,
		Method:       logInfo.Method,
		RequestUrl:   logInfo.RequestUrl,
		RequestParam: logInfo.RequestParam,
		RequestType:  logInfo.RequestType,
		CostTime:     logInfo.CostTime,
		CreateTime:   gtime.Now(),
	})
	if err != nil {
		g.Log().Error(ctx, "添加系统日志失败:", err)
		return errors.New("添加系统日志失败")
	}

	return nil
}

// GetDataLogList 获取数据日志列表
func (s *sSysLog) GetDataLogList(ctx context.Context, req *v1.DataLogListReq) (list []*v1.DataLogInfo, total int64, err error) {
	// 构建查询
	m := dao.SysDataLog.Ctx(ctx)

	// 表名筛选
	if req.DataTable != "" {
		m = m.Where(dao.SysDataLog.Columns().DataTable, req.DataTable)
	}

	// 数据ID筛选
	if req.DataId != "" {
		m = m.Where(dao.SysDataLog.Columns().DataId, req.DataId)
	}

	// 创建人模糊查询
	if req.CreateBy != "" {
		m = m.WhereLike(dao.SysDataLog.Columns().CreateBy, "%"+req.CreateBy+"%")
	}

	// 时间范围筛选
	if req.StartTime != "" {
		m = m.WhereGTE(dao.SysDataLog.Columns().CreateTime, req.StartTime)
	}
	if req.EndTime != "" {
		m = m.WhereLTE(dao.SysDataLog.Columns().CreateTime, req.EndTime+" 23:59:59")
	}

	// 分页查询
	var logs []*entity.SysDataLog
	var totalInt int
	err = m.Page(req.Page, req.PageSize).
		OrderDesc(dao.SysDataLog.Columns().CreateTime).
		ScanAndCount(&logs, &totalInt, false)
	if err != nil {
		g.Log().Error(ctx, "查询数据日志列表失败:", err)
		return nil, 0, errors.New("查询数据日志列表失败")
	}
	total = int64(totalInt)

	// 转换为DataLogInfo
	list = make([]*v1.DataLogInfo, 0, len(logs))
	for _, log := range logs {
		createTime := ""
		if log.CreateTime != nil {
			createTime = log.CreateTime.String()
		}
		list = append(list, &v1.DataLogInfo{
			Id:          log.Id,
			CreateBy:    log.CreateBy,
			CreateTime:  createTime,
			DataTable:   log.DataTable,
			DataId:      log.DataId,
			DataContent: log.DataContent,
			DataVersion: log.DataVersion,
		})
	}

	return list, total, nil
}

// GetDataLogById 根据ID获取数据日志详情
func (s *sSysLog) GetDataLogById(ctx context.Context, id string) (*v1.DataLogInfo, error) {
	var log entity.SysDataLog
	err := dao.SysDataLog.Ctx(ctx).
		Where(dao.SysDataLog.Columns().Id, id).
		Scan(&log)
	if err != nil {
		g.Log().Error(ctx, "查询数据日志失败:", err)
		return nil, errors.New("查询数据日志失败")
	}
	if log.Id == "" {
		return nil, errors.New("数据日志不存在")
	}

	createTime := ""
	if log.CreateTime != nil {
		createTime = log.CreateTime.String()
	}

	return &v1.DataLogInfo{
		Id:          log.Id,
		CreateBy:    log.CreateBy,
		CreateTime:  createTime,
		DataTable:   log.DataTable,
		DataId:      log.DataId,
		DataContent: log.DataContent,
		DataVersion: log.DataVersion,
	}, nil
}

// GetDataLogHistory 获取数据变更历史
func (s *sSysLog) GetDataLogHistory(ctx context.Context, dataTable string, dataId string) ([]*v1.DataLogInfo, error) {
	var logs []*entity.SysDataLog
	err := dao.SysDataLog.Ctx(ctx).
		Where(dao.SysDataLog.Columns().DataTable, dataTable).
		Where(dao.SysDataLog.Columns().DataId, dataId).
		OrderDesc(dao.SysDataLog.Columns().DataVersion).
		Scan(&logs)
	if err != nil {
		g.Log().Error(ctx, "查询数据变更历史失败:", err)
		return nil, errors.New("查询数据变更历史失败")
	}

	list := make([]*v1.DataLogInfo, 0, len(logs))
	for _, log := range logs {
		createTime := ""
		if log.CreateTime != nil {
			createTime = log.CreateTime.String()
		}
		list = append(list, &v1.DataLogInfo{
			Id:          log.Id,
			CreateBy:    log.CreateBy,
			CreateTime:  createTime,
			DataTable:   log.DataTable,
			DataId:      log.DataId,
			DataContent: log.DataContent,
			DataVersion: log.DataVersion,
		})
	}

	return list, nil
}

// AddDataLog 添加数据日志
func (s *sSysLog) AddDataLog(ctx context.Context, dataTable string, dataId string, dataContent string) error {
	// 获取当前最大版本号
	var maxVersion int
	value, err := dao.SysDataLog.Ctx(ctx).
		Where(dao.SysDataLog.Columns().DataTable, dataTable).
		Where(dao.SysDataLog.Columns().DataId, dataId).
		Max(dao.SysDataLog.Columns().DataVersion)
	if err != nil {
		g.Log().Warning(ctx, "获取数据日志最大版本号失败:", err)
		maxVersion = 0
	} else {
		maxVersion = int(value)
	}

	// 插入数据日志
	_, err = dao.SysDataLog.Ctx(ctx).Insert(do.SysDataLog{
		Id:          guid.S(),
		DataTable:   dataTable,
		DataId:      dataId,
		DataContent: dataContent,
		DataVersion: maxVersion + 1,
		CreateTime:  gtime.Now(),
	})
	if err != nil {
		g.Log().Error(ctx, "添加数据日志失败:", err)
		return errors.New("添加数据日志失败")
	}

	return nil
}
