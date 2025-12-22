// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// TeachingDepartDayLog is the golang structure for table teaching_depart_day_log.
type TeachingDepartDayLog struct {
	Id                         string      `json:"id"                         description:"主键ID"`
	DepartId                   string      `json:"departId"                   description:"班级ID"`
	DepartName                 string      `json:"departName"                 description:"班级名"`
	UnitOpenCount              int         `json:"unitOpenCount"              description:"开课次数"`
	CourseWorkAssignCount      int         `json:"courseWorkAssignCount"      description:"课程作业布置次数"`
	AdditionalWorkAssignCount  int         `json:"additionalWorkAssignCount"  description:"附加作业布置次数"`
	CourseWorkCorrectCount     int         `json:"courseWorkCorrectCount"     description:"课程作业批改次数"`
	AdditionalWorkCorrectCount int         `json:"additionalWorkCorrectCount" description:"附加作业批改次数"`
	CourseWorkSubmitCount      int         `json:"courseWorkSubmitCount"      description:"课程作业提交次数"`
	AdditionalWorkSubmitCount  int         `json:"additionalWorkSubmitCount"  description:"附加作业提交次数"`
	CreateTime                 *gtime.Time `json:"createTime"                 description:"日期"`
}
