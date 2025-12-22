package teaching

import (
	"context"
	"teaching-open/internal/model/vo"
	"teaching-open/internal/service"

	"github.com/gogf/gf/v2/frame/g"
)

// TeachingCourseDeptController 课程部门关联控制器
type TeachingCourseDeptController struct {
	deptService service.TeachingCourseDeptService
}

// NewTeachingCourseDeptController 创建课程部门关联控制器
func NewTeachingCourseDeptController() *TeachingCourseDeptController {
	return &TeachingCourseDeptController{
		deptService: service.NewTeachingCourseDeptService(),
	}
}

// List 课程部门关联列表
func (c *TeachingCourseDeptController) List(ctx context.Context, req *vo.CourseDeptListReq) (res *vo.CourseDeptListRes, err error) {
	res, err = c.deptService.List(ctx, req)
	return
}

// Assign 批量分配课程到部门
func (c *TeachingCourseDeptController) Assign(ctx context.Context, req *vo.CourseDeptAssignReq) (err error) {
	userId := g.RequestFromCtx(ctx).GetCtxVar("userId").String()
	err = c.deptService.Assign(ctx, req, userId)
	return
}

// Remove 移除课程部门关联
func (c *TeachingCourseDeptController) Remove(ctx context.Context, req *vo.CourseDeptRemoveReq) (err error) {
	userId := g.RequestFromCtx(ctx).GetCtxVar("userId").String()
	err = c.deptService.Remove(ctx, req, userId)
	return
}