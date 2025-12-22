// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import "github.com/gogf/gf/v2/os/gtime"

// TeachingCourse is the golang structure for table teaching_course.
type TeachingCourse struct {
	Id             string      `json:"id"             orm:"id"              description:"主键"`
	CreateBy       string      `json:"createBy"       orm:"create_by"       description:"创建人"`
	CreateTime     *gtime.Time `json:"createTime"     orm:"create_time"     description:"创建日期"`
	UpdateBy       string      `json:"updateBy"       orm:"update_by"       description:"更新人"`
	UpdateTime     *gtime.Time `json:"updateTime"     orm:"update_time"     description:"更新日期"`
	SysOrgCode     string      `json:"sysOrgCode"     orm:"sys_org_code"    description:"所属部门"`
	DelFlag        int         `json:"delFlag"        orm:"del_flag"        description:"删除标志"`
	CourseName     string      `json:"courseName"     orm:"course_name"     description:"科目名"`
	CourseDesc     string      `json:"courseDesc"     orm:"course_desc"     description:"科目介绍"`
	CourseIcon     string      `json:"courseIcon"     orm:"course_icon"     description:"科目图标"`
	CourseCover    string      `json:"courseCover"    orm:"course_cover"    description:"科目封面"`
	ShowType       int         `json:"showType"       orm:"show_type"       description:"展示类型"`
	CourseMap      string      `json:"courseMap"      orm:"course_map"      description:"课程地图"`
	IsShared       int         `json:"isShared"       orm:"is_shared"       description:"是否共享课程"`
	ShowHome       int         `json:"showHome"       orm:"show_home"       description:"首页展示"`
	OrderNum       int         `json:"orderNum"       orm:"order_num"       description:"排序"`
	DepartIds      string      `json:"departIds"      orm:"depart_ids"      description:"授权部门"`
	CourseType     string      `json:"courseType"     orm:"course_type"     description:"课程类型"`
	CourseCategory string      `json:"courseCategory" orm:"course_category" description:"课程分类"`
}
