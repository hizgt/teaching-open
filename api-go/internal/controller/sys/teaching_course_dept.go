package sys

import (
	"context"
	"strings"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gtime"

	v1 "teaching-open/api/v1/sys"
	"teaching-open/internal/model/entity"
	"teaching-open/internal/service"
)

type TeachingCourseDeptController struct{}

var TeachingCourseDept = &TeachingCourseDeptController{}

func (c *TeachingCourseDeptController) GetList(ctx context.Context, req *v1.CourseDeptListReq) (res *v1.CourseDeptListRes, err error) {
	list, total, err := service.TeachingCourseDept().GetList(ctx, req.DeptId, req.CourseId, req.PageNo, req.PageSize)
	if err != nil {
		return nil, gerror.Wrap(err, "查询班级课程列表失败")
	}
	return &v1.CourseDeptListRes{List: list, Total: total, PageNo: req.PageNo, PageSize: req.PageSize}, nil
}

func (c *TeachingCourseDeptController) GetByDeptId(ctx context.Context, req *v1.CourseDeptByDeptReq) (res *v1.CourseDeptByDeptRes, err error) {
	list, err := service.TeachingCourseDept().GetByDeptId(ctx, req.DeptId)
	if err != nil {
		return nil, gerror.Wrap(err, "查询班级课程失败")
	}
	return &v1.CourseDeptByDeptRes{List: list}, nil
}

func (c *TeachingCourseDeptController) GetByCourseId(ctx context.Context, req *v1.CourseDeptByCourseReq) (res *v1.CourseDeptByCourseRes, err error) {
	list, err := service.TeachingCourseDept().GetByCourseId(ctx, req.CourseId)
	if err != nil {
		return nil, gerror.Wrap(err, "查询课程授权班级失败")
	}
	return &v1.CourseDeptByCourseRes{List: list}, nil
}

func (c *TeachingCourseDeptController) AddOrUpdate(ctx context.Context, req *v1.CourseDeptAddOrUpdateReq) (res *v1.CourseDeptAddOrUpdateRes, err error) {
	var openTime *gtime.Time
	if req.OpenTime != "" {
		openTime = gtime.NewFromStr(req.OpenTime)
	}
	courseDept := &entity.TeachingCourseDept{
		DeptId:   req.DeptId,
		CourseId: req.CourseId,
		OpenTime: openTime,
	}
	err = service.TeachingCourseDept().AddOrUpdate(ctx, courseDept)
	if err != nil {
		return nil, gerror.Wrap(err, "操作失败")
	}
	return &v1.CourseDeptAddOrUpdateRes{}, nil
}

func (c *TeachingCourseDeptController) Delete(ctx context.Context, req *v1.CourseDeptDeleteReq) (res *v1.CourseDeptDeleteRes, err error) {
	err = service.TeachingCourseDept().Delete(ctx, req.Id)
	if err != nil {
		return nil, gerror.Wrap(err, "删除班级课程失败")
	}
	return &v1.CourseDeptDeleteRes{}, nil
}

func (c *TeachingCourseDeptController) BatchAdd(ctx context.Context, req *v1.CourseDeptBatchAddReq) (res *v1.CourseDeptBatchAddRes, err error) {
	deptIds := strings.Split(req.DeptIds, ",")
	err = service.TeachingCourseDept().BatchAdd(ctx, deptIds, req.CourseId)
	if err != nil {
		return nil, gerror.Wrap(err, "批量添加班级课程失败")
	}
	return &v1.CourseDeptBatchAddRes{}, nil
}
