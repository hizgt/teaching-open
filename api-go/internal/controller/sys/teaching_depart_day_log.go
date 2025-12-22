package sys

import (
	"context"

	v1 "teaching-open/api/v1/sys"
	"teaching-open/internal/service"
)

var TeachingDepartDayLog = cTeachingDepartDayLog{}

type cTeachingDepartDayLog struct{}

// GetReport 统计报表
func (c *cTeachingDepartDayLog) GetReport(ctx context.Context, req *v1.DepartDayLogGetReportReq) (res *v1.DepartDayLogGetReportRes, err error) {
	list, total, err := service.TeachingDepartDayLog().GetReport(ctx, req)
	if err != nil {
		return nil, err
	}

	return &v1.DepartDayLogGetReportRes{
		List:  list,
		Total: total,
	}, nil
}

// GetReportGroupByDepart 按部门统计
func (c *cTeachingDepartDayLog) GetReportGroupByDepart(ctx context.Context, req *v1.DepartDayLogGroupByDepartReq) (res *v1.DepartDayLogGroupByDepartRes, err error) {
	list, total, err := service.TeachingDepartDayLog().GetReportGroupByDepart(ctx, req)
	if err != nil {
		return nil, err
	}

	return &v1.DepartDayLogGroupByDepartRes{
		List:  list,
		Total: total,
	}, nil
}

// GetReportGroupByMonth 按月份统计
func (c *cTeachingDepartDayLog) GetReportGroupByMonth(ctx context.Context, req *v1.DepartDayLogGroupByMonthReq) (res *v1.DepartDayLogGroupByMonthRes, err error) {
	list, err := service.TeachingDepartDayLog().GetReportGroupByMonth(ctx, req)
	if err != nil {
		return nil, err
	}

	return &v1.DepartDayLogGroupByMonthRes{
		List: list,
	}, nil
}

// UnitViewLog 记录单元查看日志
func (c *cTeachingDepartDayLog) UnitViewLog(ctx context.Context, req *v1.DepartDayLogUnitViewReq) (res *v1.DepartDayLogUnitViewRes, err error) {
	err = service.TeachingDepartDayLog().UnitViewLog(ctx, req)
	if err != nil {
		return nil, err
	}

	return &v1.DepartDayLogUnitViewRes{}, nil
}
