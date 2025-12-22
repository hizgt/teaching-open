// Package sys provides controllers for teaching management.
package sys

import (
	"context"
	"strings"

	"github.com/gogf/gf/v2/errors/gerror"

	v1 "teaching-open/api/v1/sys"
	"teaching-open/internal/model/entity"
	"teaching-open/internal/service"
)

// TeachingCourseController 课程管理控制器
type TeachingCourseController struct{}

var TeachingCourse = &TeachingCourseController{}

// GetList 获取课程列表
func (c *TeachingCourseController) GetList(ctx context.Context, req *v1.CourseListReq) (res *v1.CourseListRes, err error) {
	list, total, err := service.TeachingCourse().GetList(
		ctx,
		req.CourseName,
		req.CourseType,
		req.CourseCategory,
		req.IsShared,
		req.ShowHome,
		req.PageNo,
		req.PageSize,
	)
	if err != nil {
		return nil, gerror.Wrap(err, "查询课程列表失败")
	}

	return &v1.CourseListRes{
		List:     list,
		Total:    total,
		PageNo:   req.PageNo,
		PageSize: req.PageSize,
	}, nil
}

// GetHomeCourse 获取首页课程
func (c *TeachingCourseController) GetHomeCourse(ctx context.Context, req *v1.CourseHomeCourseReq) (res *v1.CourseHomeCourseRes, err error) {
	list, err := service.TeachingCourse().GetHomeCourse(ctx)
	if err != nil {
		return nil, gerror.Wrap(err, "查询首页课程失败")
	}

	return &v1.CourseHomeCourseRes{
		List: list,
	}, nil
}

// Add 添加课程
func (c *TeachingCourseController) Add(ctx context.Context, req *v1.CourseAddReq) (res *v1.CourseAddRes, err error) {
	course := &entity.TeachingCourse{
		CourseName:     req.CourseName,
		CourseDesc:     req.CourseDesc,
		CourseIcon:     req.CourseIcon,
		CourseCover:    req.CourseCover,
		ShowType:       req.ShowType,
		CourseMap:      req.CourseMap,
		IsShared:       req.IsShared,
		ShowHome:       req.ShowHome,
		OrderNum:       req.OrderNum,
		DepartIds:      req.DepartIds,
		CourseType:     req.CourseType,
		CourseCategory: req.CourseCategory,
	}

	err = service.TeachingCourse().Add(ctx, course)
	if err != nil {
		return nil, gerror.Wrap(err, "添加课程失败")
	}

	return &v1.CourseAddRes{}, nil
}

// Edit 编辑课程
func (c *TeachingCourseController) Edit(ctx context.Context, req *v1.CourseEditReq) (res *v1.CourseEditRes, err error) {
	course := &entity.TeachingCourse{
		Id:             req.Id,
		CourseName:     req.CourseName,
		CourseDesc:     req.CourseDesc,
		CourseIcon:     req.CourseIcon,
		CourseCover:    req.CourseCover,
		ShowType:       req.ShowType,
		CourseMap:      req.CourseMap,
		IsShared:       req.IsShared,
		ShowHome:       req.ShowHome,
		OrderNum:       req.OrderNum,
		DepartIds:      req.DepartIds,
		CourseType:     req.CourseType,
		CourseCategory: req.CourseCategory,
	}

	err = service.TeachingCourse().Edit(ctx, course)
	if err != nil {
		return nil, gerror.Wrap(err, "编辑课程失败")
	}

	return &v1.CourseEditRes{}, nil
}

// Delete 删除课程
func (c *TeachingCourseController) Delete(ctx context.Context, req *v1.CourseDeleteReq) (res *v1.CourseDeleteRes, err error) {
	err = service.TeachingCourse().Delete(ctx, req.Id)
	if err != nil {
		return nil, gerror.Wrap(err, "删除课程失败")
	}

	return &v1.CourseDeleteRes{}, nil
}

// DeleteBatch 批量删除课程
func (c *TeachingCourseController) DeleteBatch(ctx context.Context, req *v1.CourseDeleteBatchReq) (res *v1.CourseDeleteBatchRes, err error) {
	ids := strings.Split(req.Ids, ",")
	err = service.TeachingCourse().DeleteBatch(ctx, ids)
	if err != nil {
		return nil, gerror.Wrap(err, "批量删除课程失败")
	}

	return &v1.CourseDeleteBatchRes{}, nil
}

// GetById 获取课程详情
func (c *TeachingCourseController) GetById(ctx context.Context, req *v1.CourseGetByIdReq) (res *v1.CourseGetByIdRes, err error) {
	course, err := service.TeachingCourse().GetById(ctx, req.Id)
	if err != nil {
		return nil, gerror.Wrap(err, "查询课程详情失败")
	}
	if course == nil {
		return nil, gerror.New("课程不存在")
	}

	return &v1.CourseGetByIdRes{
		TeachingCourse: course,
	}, nil
}

// Publish 发布/下架课程
func (c *TeachingCourseController) Publish(ctx context.Context, req *v1.CoursePublishReq) (res *v1.CoursePublishRes, err error) {
	err = service.TeachingCourse().Publish(ctx, req.Id, req.ShowHome)
	if err != nil {
		return nil, gerror.Wrap(err, "操作失败")
	}

	return &v1.CoursePublishRes{}, nil
}

// SetShared 设置共享
func (c *TeachingCourseController) SetShared(ctx context.Context, req *v1.CourseSetSharedReq) (res *v1.CourseSetSharedRes, err error) {
	err = service.TeachingCourse().SetShared(ctx, req.Id, req.IsShared)
	if err != nil {
		return nil, gerror.Wrap(err, "设置共享状态失败")
	}

	return &v1.CourseSetSharedRes{}, nil
}

// AuthorizeDept 授权部门
func (c *TeachingCourseController) AuthorizeDept(ctx context.Context, req *v1.CourseAuthorizeDeptReq) (res *v1.CourseAuthorizeDeptRes, err error) {
	err = service.TeachingCourse().AuthorizeDept(ctx, req.Id, req.DepartIds)
	if err != nil {
		return nil, gerror.Wrap(err, "授权部门失败")
	}

	return &v1.CourseAuthorizeDeptRes{}, nil
}
