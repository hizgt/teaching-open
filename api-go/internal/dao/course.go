// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package dao

import (
	"teaching-open/internal/dao/internal"
)

var (
	// Course is the DAO for table teaching_course.
	Course *internal.CourseDao
	// CourseUnit is the DAO for table teaching_course_unit.
	CourseUnit *internal.CourseUnitDao
	// CourseDept is the DAO for table teaching_course_dept.
	CourseDept *internal.CourseDeptDao
)

func init() {
	Course = internal.NewCourseDao()
	CourseUnit = internal.NewCourseUnitDao()
	CourseDept = internal.NewCourseDeptDao()
}