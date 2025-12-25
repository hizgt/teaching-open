package teaching

import (
	"context"
	"teaching-open/internal/model/vo"
	"teaching-open/internal/service"

	"github.com/gogf/gf/v2/frame/g"
)

// TeachingCourseController 课程控制器
type TeachingCourseController struct {
	courseService service.TeachingCourseService
}

// NewTeachingCourseController 创建课程控制器
func NewTeachingCourseController() *TeachingCourseController {
	return &TeachingCourseController{
		courseService: service.NewTeachingCourseService(),
	}
}

// List 课程列表
func (c *TeachingCourseController) List(ctx context.Context, req *vo.CourseListReq) (res *vo.CourseListRes, err error) {
	res, err = c.courseService.List(ctx, req)
	return
}

// Create 创建课程
func (c *TeachingCourseController) Create(ctx context.Context, req *vo.CourseCreateReq) (res *vo.CourseCreateRes, err error) {
	userId := g.RequestFromCtx(ctx).GetCtxVar("userId").String()
	res, err = c.courseService.Create(ctx, req, userId)
	return
}

// Update 更新课程
func (c *TeachingCourseController) Update(ctx context.Context, req *vo.CourseUpdateReq) (err error) {
	userId := g.RequestFromCtx(ctx).GetCtxVar("userId").String()
	err = c.courseService.Update(ctx, req, userId)
	return
}

// Delete 删除课程
func (c *TeachingCourseController) Delete(ctx context.Context, req *vo.CourseDeleteReq) (err error) {
	userId := g.RequestFromCtx(ctx).GetCtxVar("userId").String()
	err = c.courseService.Delete(ctx, req, userId)
	return
}

// Detail 课程详情
func (c *TeachingCourseController) Detail(ctx context.Context, req *vo.CourseDetailReq) (res *vo.CourseDetailRes, err error) {
	res, err = c.courseService.Detail(ctx, req.Id)
	return
}

// Publish 发布/下架课程
func (c *TeachingCourseController) Publish(ctx context.Context, req *vo.CoursePublishReq) (err error) {
	userId := g.RequestFromCtx(ctx).GetCtxVar("userId").String()
	err = c.courseService.Publish(ctx, req, userId)
	return
}