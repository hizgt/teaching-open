package teaching

import (
	"github.com/gogf/gf/v2/net/ghttp"

	"teaching-open/internal/model/vo"
	"teaching-open/internal/service"
	"teaching-open/utility/response"
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
func (c *TeachingCourseController) List(r *ghttp.Request) {
	var req vo.CourseListReq
	if err := r.Parse(&req); err != nil {
		response.Error(r, err.Error())
		return
	}

	res, err := c.courseService.List(r.Context(), &req)
	if err != nil {
		response.Error(r, err.Error())
		return
	}

	response.PageSuccess(r, res.Records, res.Total, res.Page, res.PageSize)
}

// Add 创建课程
func (c *TeachingCourseController) Add(r *ghttp.Request) {
	var req vo.CourseCreateReq
	if err := r.Parse(&req); err != nil {
		response.Error(r, err.Error())
		return
	}

	userId := r.Get("userId").String()
	res, err := c.courseService.Create(r.Context(), &req, userId)
	if err != nil {
		response.Error(r, err.Error())
		return
	}

	response.Success(r, res)
}

// Edit 更新课程
func (c *TeachingCourseController) Edit(r *ghttp.Request) {
	var req vo.CourseUpdateReq
	if err := r.Parse(&req); err != nil {
		response.Error(r, err.Error())
		return
	}

	userId := r.Get("userId").String()
	err := c.courseService.Update(r.Context(), &req, userId)
	if err != nil {
		response.Error(r, err.Error())
		return
	}

	response.SuccessMsg(r, "修改成功")
}

// Delete 删除课程
func (c *TeachingCourseController) Delete(r *ghttp.Request) {
	id := r.Get("id").String()
	if id == "" {
		response.Error(r, "课程ID不能为空")
		return
	}

	userId := r.Get("userId").String()
	req := &vo.CourseDeleteReq{Id: id}
	err := c.courseService.Delete(r.Context(), req, userId)
	if err != nil {
		response.Error(r, err.Error())
		return
	}

	response.SuccessMsg(r, "删除成功")
}

// QueryById 课程详情
func (c *TeachingCourseController) QueryById(r *ghttp.Request) {
	id := r.Get("id").String()
	if id == "" {
		response.Error(r, "课程ID不能为空")
		return
	}

	res, err := c.courseService.Detail(r.Context(), id)
	if err != nil {
		response.Error(r, err.Error())
		return
	}

	response.Success(r, res)
}

// Publish 发布/下架课程
func (c *TeachingCourseController) Publish(r *ghttp.Request) {
	var req vo.CoursePublishReq
	if err := r.Parse(&req); err != nil {
		response.Error(r, err.Error())
		return
	}

	userId := r.Get("userId").String()
	err := c.courseService.Publish(r.Context(), &req, userId)
	if err != nil {
		response.Error(r, err.Error())
		return
	}

	response.SuccessMsg(r, "操作成功")
}

// GetHomeCourse 获取首页课程
func (c *TeachingCourseController) GetHomeCourse(r *ghttp.Request) {
	res, err := c.courseService.GetHomeCourse(r.Context())
	if err != nil {
		response.Error(r, err.Error())
		return
	}

	response.Success(r, res)
}
