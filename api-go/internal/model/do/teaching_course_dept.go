// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import "github.com/gogf/gf/v2/frame/g"

// TeachingCourseDept is the golang structure for table teaching_course_dept.
type TeachingCourseDept struct {
	g.Meta     `orm:"table:teaching_course_dept, do:true"`
	Id         interface{} // 主键
	CreateBy   interface{} // 创建人
	CreateTime interface{} // 创建日期
	UpdateBy   interface{} // 更新人
	UpdateTime interface{} // 更新日期
	SysOrgCode interface{} // 所属部门
	DeptId     interface{} // 班级ID
	CourseId   interface{} // 课程ID
	OpenTime   interface{} // 课程开课时间
}
