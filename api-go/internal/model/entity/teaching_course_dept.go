package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// TeachingCourseDept 课程部门关联表
type TeachingCourseDept struct {
	Id       string      `json:"id"       description:"关联ID"`
	CourseId string      `json:"courseId" description:"课程ID"`
	DeptId   string      `json:"deptId"   description:"部门ID"`
	CreateBy string      `json:"createBy" description:"创建人"`
	CreateTime *gtime.Time `json:"createTime" description:"创建时间"`
}