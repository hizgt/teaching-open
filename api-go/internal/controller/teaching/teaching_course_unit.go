package teaching

import (
	"context"
	"teaching-open/internal/model/vo"
	"teaching-open/internal/service"

	"github.com/gogf/gf/v2/frame/g"
)

// TeachingCourseUnitController 课程单元控制器
type TeachingCourseUnitController struct {
	unitService service.TeachingCourseUnitService
}

// NewTeachingCourseUnitController 创建课程单元控制器
func NewTeachingCourseUnitController() *TeachingCourseUnitController {
	return &TeachingCourseUnitController{
		unitService: service.NewTeachingCourseUnitService(),
	}
}

// List 课程单元列表
func (c *TeachingCourseUnitController) List(ctx context.Context, req *vo.CourseUnitListReq) (res *vo.CourseUnitListRes, err error) {
	res, err = c.unitService.List(ctx, req)
	return
}

// Create 创建课程单元
func (c *TeachingCourseUnitController) Create(ctx context.Context, req *vo.CourseUnitCreateReq) (res *vo.CourseUnitCreateRes, err error) {
	userId := g.RequestFromCtx(ctx).GetCtxVar("userId").String()
	res, err = c.unitService.Create(ctx, req, userId)
	return
}

// Update 更新课程单元
func (c *TeachingCourseUnitController) Update(ctx context.Context, req *vo.CourseUnitUpdateReq) (err error) {
	userId := g.RequestFromCtx(ctx).GetCtxVar("userId").String()
	err = c.unitService.Update(ctx, req, userId)
	return
}

// Delete 删除课程单元
func (c *TeachingCourseUnitController) Delete(ctx context.Context, req *vo.CourseUnitDeleteReq) (err error) {
	userId := g.RequestFromCtx(ctx).GetCtxVar("userId").String()
	err = c.unitService.Delete(ctx, req, userId)
	return
}

// Sort 排序课程单元
func (c *TeachingCourseUnitController) Sort(ctx context.Context, req *vo.CourseUnitSortReq) (err error) {
	userId := g.RequestFromCtx(ctx).GetCtxVar("userId").String()
	err = c.unitService.Sort(ctx, req, userId)
	return
}