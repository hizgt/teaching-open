package teaching

import (
	"github.com/gogf/gf/v2/net/ghttp"

	"teaching-open/internal/model/vo"
	"teaching-open/internal/service"
	"teaching-open/utility/response"
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
func (c *TeachingCourseUnitController) List(r *ghttp.Request) {
	var req vo.CourseUnitListReq
	if err := r.Parse(&req); err != nil {
		response.Error(r, err.Error())
		return
	}

	res, err := c.unitService.List(r.Context(), &req)
	if err != nil {
		response.Error(r, err.Error())
		return
	}

	response.PageSuccess(r, res.Records, res.Total, res.Page, res.PageSize)
}

// QueryByCourseId 根据课程ID查询单元
func (c *TeachingCourseUnitController) QueryByCourseId(r *ghttp.Request) {
	courseId := r.Get("courseId").String()
	if courseId == "" {
		response.Error(r, "课程ID不能为空")
		return
	}

	res, err := c.unitService.QueryByCourseId(r.Context(), courseId)
	if err != nil {
		response.Error(r, err.Error())
		return
	}

	response.Success(r, res)
}

// Add 创建课程单元
func (c *TeachingCourseUnitController) Add(r *ghttp.Request) {
	var req vo.CourseUnitCreateReq
	if err := r.Parse(&req); err != nil {
		response.Error(r, err.Error())
		return
	}

	userId := r.Get("userId").String()
	res, err := c.unitService.Create(r.Context(), &req, userId)
	if err != nil {
		response.Error(r, err.Error())
		return
	}

	response.Success(r, res)
}

// Edit 更新课程单元
func (c *TeachingCourseUnitController) Edit(r *ghttp.Request) {
	var req vo.CourseUnitUpdateReq
	if err := r.Parse(&req); err != nil {
		response.Error(r, err.Error())
		return
	}

	userId := r.Get("userId").String()
	err := c.unitService.Update(r.Context(), &req, userId)
	if err != nil {
		response.Error(r, err.Error())
		return
	}

	response.SuccessMsg(r, "修改成功")
}

// Delete 删除课程单元
func (c *TeachingCourseUnitController) Delete(r *ghttp.Request) {
	id := r.Get("id").String()
	if id == "" {
		response.Error(r, "单元ID不能为空")
		return
	}

	userId := r.Get("userId").String()
	req := &vo.CourseUnitDeleteReq{Id: id}
	err := c.unitService.Delete(r.Context(), req, userId)
	if err != nil {
		response.Error(r, err.Error())
		return
	}

	response.SuccessMsg(r, "删除成功")
}

// QueryById 单元详情
func (c *TeachingCourseUnitController) QueryById(r *ghttp.Request) {
	id := r.Get("id").String()
	if id == "" {
		response.Error(r, "单元ID不能为空")
		return
	}

	res, err := c.unitService.QueryById(r.Context(), id)
	if err != nil {
		response.Error(r, err.Error())
		return
	}

	response.Success(r, res)
}

// Sort 排序课程单元
func (c *TeachingCourseUnitController) Sort(r *ghttp.Request) {
	var req vo.CourseUnitSortReq
	if err := r.Parse(&req); err != nil {
		response.Error(r, err.Error())
		return
	}

	userId := r.Get("userId").String()
	err := c.unitService.Sort(r.Context(), &req, userId)
	if err != nil {
		response.Error(r, err.Error())
		return
	}

	response.SuccessMsg(r, "排序成功")
}
