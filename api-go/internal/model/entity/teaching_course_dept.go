// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import "github.com/gogf/gf/v2/os/gtime"

// TeachingCourseDept is the golang structure for table teaching_course_dept.
type TeachingCourseDept struct {
	Id         string      `json:"id"         orm:"id"           description:"主键"`
	CreateBy   string      `json:"createBy"   orm:"create_by"    description:"创建人"`
	CreateTime *gtime.Time `json:"createTime" orm:"create_time"  description:"创建日期"`
	UpdateBy   string      `json:"updateBy"   orm:"update_by"    description:"更新人"`
	UpdateTime *gtime.Time `json:"updateTime" orm:"update_time"  description:"更新日期"`
	SysOrgCode string      `json:"sysOrgCode" orm:"sys_org_code" description:"所属部门"`
	DeptId     string      `json:"deptId"     orm:"dept_id"      description:"班级ID"`
	CourseId   string      `json:"courseId"   orm:"course_id"    description:"课程ID"`
	OpenTime   *gtime.Time `json:"openTime"   orm:"open_time"    description:"课程开课时间"`
}
