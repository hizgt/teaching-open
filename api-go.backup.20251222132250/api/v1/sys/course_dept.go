package sys

import (
	"github.com/gogf/gf/v2/frame/g"
	"teaching-open/internal/model/entity"
)

// =============================================================================
// Course Dept API - 班级课程关联接口
// =============================================================================

// CourseDeptListReq 班级课程列表请求
type CourseDeptListReq struct {
	g.Meta   `path:"/teaching/teachingCourseDept/list" method:"get" tags:"班级课程管理" summary:"班级课程列表"`
	DeptId   string `json:"deptId" dc:"班级ID"`
	CourseId string `json:"courseId" dc:"课程ID"`
	PageNo   int    `json:"pageNo" d:"1"`
	PageSize int    `json:"pageSize" d:"10"`
}

// CourseDeptListRes 班级课程列表响应
type CourseDeptListRes struct {
	g.Meta   `mime:"application/json"`
	List     []*entity.TeachingCourseDept `json:"records"`
	Total    int                          `json:"total"`
	PageNo   int                          `json:"current"`
	PageSize int                          `json:"size"`
}

// CourseDeptByDeptReq 获取班级的课程列表请求
type CourseDeptByDeptReq struct {
	g.Meta `path:"/teaching/teachingCourseDept/queryByDeptId" method:"get" tags:"班级课程管理" summary:"获取班级的课程列表"`
	DeptId string `json:"deptId" v:"required#班级ID不能为空" dc:"班级ID"`
}

// CourseDeptByDeptRes 获取班级的课程列表响应
type CourseDeptByDeptRes struct {
	g.Meta `mime:"application/json"`
	List   []*entity.TeachingCourseDept `json:"list"`
}

// CourseDeptByCourseReq 获取课程授权的班级列表请求
type CourseDeptByCourseReq struct {
	g.Meta   `path:"/teaching/teachingCourseDept/queryByCourseId" method:"get" tags:"班级课程管理" summary:"获取课程授权的班级列表"`
	CourseId string `json:"courseId" v:"required#课程ID不能为空" dc:"课程ID"`
}

// CourseDeptByCourseRes 获取课程授权的班级列表响应
type CourseDeptByCourseRes struct {
	g.Meta `mime:"application/json"`
	List   []*entity.TeachingCourseDept `json:"list"`
}

// CourseDeptAddOrUpdateReq 添加或更新班级课程请求
type CourseDeptAddOrUpdateReq struct {
	g.Meta   `path:"/teaching/teachingCourseDept/addOrUpdate" method:"post" tags:"班级课程管理" summary:"添加或更新班级课程"`
	DeptId   string `json:"deptId" v:"required#班级ID不能为空" dc:"班级ID"`
	CourseId string `json:"courseId" v:"required#课程ID不能为空" dc:"课程ID"`
	OpenTime string `json:"openTime" dc:"开课时间"`
}

// CourseDeptAddOrUpdateRes 添加或更新班级课程响应
type CourseDeptAddOrUpdateRes struct {
	g.Meta `mime:"application/json"`
}

// CourseDeptDeleteReq 删除班级课程请求
type CourseDeptDeleteReq struct {
	g.Meta `path:"/teaching/teachingCourseDept/delete" method:"delete" tags:"班级课程管理" summary:"删除班级课程"`
	Id     string `json:"id" v:"required#ID不能为空" dc:"班级课程关联ID"`
}

// CourseDeptDeleteRes 删除班级课程响应
type CourseDeptDeleteRes struct {
	g.Meta `mime:"application/json"`
}

// CourseDeptBatchAddReq 批量添加班级课程请求
type CourseDeptBatchAddReq struct {
	g.Meta   `path:"/teaching/teachingCourseDept/batchAdd" method:"post" tags:"班级课程管理" summary:"批量添加班级课程"`
	DeptIds  string `json:"deptIds" v:"required#班级ID不能为空" dc:"班级ID，多个用逗号分隔"`
	CourseId string `json:"courseId" v:"required#课程ID不能为空" dc:"课程ID"`
}

// CourseDeptBatchAddRes 批量添加班级课程响应
type CourseDeptBatchAddRes struct {
	g.Meta `mime:"application/json"`
}
