package teaching

import (
	"github.com/gogf/gf/v2/net/ghttp"

	"teaching-open/internal/model/vo"
	"teaching-open/internal/service"
	"teaching-open/utility/response"
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
func (c *TeachingCourseDeptController) List(r *ghttp.Request) {
	var req vo.CourseDeptListReq
	if err := r.Parse(&req); err != nil {
		response.Error(r, err.Error())
		return
	}

	res, err := c.deptService.List(r.Context(), &req)
	if err != nil {
		response.Error(r, err.Error())
		return
	}

	response.PageSuccess(r, res.Records, res.Total, res.Page, res.PageSize)
}

// QueryByDeptId 根据部门ID查询课程
func (c *TeachingCourseDeptController) QueryByDeptId(r *ghttp.Request) {
	deptId := r.Get("deptId").String()
	if deptId == "" {
		response.Error(r, "部门ID不能为空")
		return
	}

	res, err := c.deptService.QueryByDeptId(r.Context(), deptId)
	if err != nil {
		response.Error(r, err.Error())
		return
	}

	response.Success(r, res)
}

// QueryByCourseId 根据课程ID查询部门
func (c *TeachingCourseDeptController) QueryByCourseId(r *ghttp.Request) {
	courseId := r.Get("courseId").String()
	if courseId == "" {
		response.Error(r, "课程ID不能为空")
		return
	}

	res, err := c.deptService.QueryByCourseId(r.Context(), courseId)
	if err != nil {
		response.Error(r, err.Error())
		return
	}

	response.Success(r, res)
}

// AddOrUpdate 添加或更新课程部门关联
func (c *TeachingCourseDeptController) AddOrUpdate(r *ghttp.Request) {
	var req vo.CourseDeptAssignReq
	if err := r.Parse(&req); err != nil {
		response.Error(r, err.Error())
		return
	}

	userId := r.Get("userId").String()
	err := c.deptService.Assign(r.Context(), &req, userId)
	if err != nil {
		response.Error(r, err.Error())
		return
	}

	response.SuccessMsg(r, "操作成功")
}

// Delete 删除课程部门关联
func (c *TeachingCourseDeptController) Delete(r *ghttp.Request) {
	id := r.Get("id").String()
	if id == "" {
		response.Error(r, "关联ID不能为空")
		return
	}

	userId := r.Get("userId").String()
	req := &vo.CourseDeptRemoveReq{Id: id}
	err := c.deptService.Remove(r.Context(), req, userId)
	if err != nil {
		response.Error(r, err.Error())
		return
	}

	response.SuccessMsg(r, "删除成功")
}

// BatchAdd 批量添加课程部门关联
func (c *TeachingCourseDeptController) BatchAdd(r *ghttp.Request) {
	var req vo.CourseDeptBatchAddReq
	if err := r.Parse(&req); err != nil {
		response.Error(r, err.Error())
		return
	}

	userId := r.Get("userId").String()
	err := c.deptService.BatchAdd(r.Context(), &req, userId)
	if err != nil {
		response.Error(r, err.Error())
		return
	}

	response.SuccessMsg(r, "批量添加成功")
}
