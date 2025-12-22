package teaching

import (
	"github.com/gogf/gf/v2/frame/g"
)

// 课程部门关联列表
type CourseDeptListReq struct {
	g.Meta   `path:"/teaching/courseDept/list" method:"get" tags:"教学管理" summary:"课程部门关联列表"`
	CourseId string `json:"courseId" dc:"课程ID"`
	DeptId   string `json:"deptId" dc:"部门ID"`
}

type CourseDeptListRes struct {
	Records []CourseDeptItem `json:"records"`
}

type CourseDeptItem struct {
	Id         string `json:"id"`
	CourseId   string `json:"courseId"`
	CourseName string `json:"courseName"`
	DeptId     string `json:"deptId"`
	DeptName   string `json:"deptName"`
	CreateBy   string `json:"createBy"`
	CreateTime string `json:"createTime"`
}

// 批量分配课程到部门
type CourseDeptAssignReq struct {
	g.Meta   `path:"/teaching/courseDept/assign" method:"post" tags:"教学管理" summary:"批量分配课程到部门"`
	CourseId string   `json:"courseId" v:"required" dc:"课程ID"`
	DeptIds  []string `json:"deptIds" v:"required" dc:"部门ID列表"`
}

// 移除课程部门关联
type CourseDeptRemoveReq struct {
	g.Meta `path:"/teaching/courseDept/remove" method:"delete" tags:"教学管理" summary:"移除课程部门关联"`
	Id     string `json:"id" v:"required" dc:"关联ID"`
}