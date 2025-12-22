package service

import (
	"context"

	v1 "teaching-open/api/v1/sys"
)

// ITeachingDepartDayLog 部门日志统计服务接口
type ITeachingDepartDayLog interface {
	GetReport(ctx context.Context, req *v1.DepartDayLogGetReportReq) (list []*v1.DepartDayLogInfo, total int64, err error)
	GetReportGroupByDepart(ctx context.Context, req *v1.DepartDayLogGroupByDepartReq) (list []*v1.DepartStatInfo, total int64, err error)
	GetReportGroupByMonth(ctx context.Context, req *v1.DepartDayLogGroupByMonthReq) (list []*v1.MonthStatInfo, err error)
	UnitViewLog(ctx context.Context, req *v1.DepartDayLogUnitViewReq) error
}

var localTeachingDepartDayLog ITeachingDepartDayLog

func TeachingDepartDayLog() ITeachingDepartDayLog {
	if localTeachingDepartDayLog == nil {
		panic("ITeachingDepartDayLog service not registered")
	}
	return localTeachingDepartDayLog
}

func RegisterTeachingDepartDayLog(s ITeachingDepartDayLog) {
	localTeachingDepartDayLog = s
}
