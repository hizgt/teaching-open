// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// TeachingDepartDayLog is the golang structure for table teaching_depart_day_log.
type TeachingDepartDayLog struct {
	g.Meta                     `orm:"table:teaching_depart_day_log, do:true"`
	Id                         interface{} // 主键ID
	DepartId                   interface{} // 班级ID
	DepartName                 interface{} // 班级名
	UnitOpenCount              interface{} // 开课次数
	CourseWorkAssignCount      interface{} // 课程作业布置次数
	AdditionalWorkAssignCount  interface{} // 附加作业布置次数
	CourseWorkCorrectCount     interface{} // 课程作业批改次数
	AdditionalWorkCorrectCount interface{} // 附加作业批改次数
	CourseWorkSubmitCount      interface{} // 课程作业提交次数
	AdditionalWorkSubmitCount  interface{} // 附加作业提交次数
	CreateTime                 interface{} // 日期
}
