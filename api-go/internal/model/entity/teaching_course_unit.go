// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import "github.com/gogf/gf/v2/os/gtime"

// TeachingCourseUnit is the golang structure for table teaching_course_unit.
type TeachingCourseUnit struct {
	Id                string      `json:"id"                orm:"id"                  description:"主键"`
	CreateBy          string      `json:"createBy"          orm:"create_by"           description:"创建人"`
	CreateTime        *gtime.Time `json:"createTime"        orm:"create_time"         description:"创建日期"`
	UpdateBy          string      `json:"updateBy"          orm:"update_by"           description:"更新人"`
	UpdateTime        *gtime.Time `json:"updateTime"        orm:"update_time"         description:"更新日期"`
	SysOrgCode        string      `json:"sysOrgCode"        orm:"sys_org_code"        description:"所属部门"`
	DelFlag           int         `json:"delFlag"           orm:"del_flag"            description:"删除标识"`
	UnitName          string      `json:"unitName"          orm:"unit_name"           description:"单元名称"`
	UnitIntro         string      `json:"unitIntro"         orm:"unit_intro"          description:"单元简介"`
	UnitCover         string      `json:"unitCover"         orm:"unit_cover"          description:"课程封面"`
	CourseId          string      `json:"courseId"          orm:"course_id"           description:"课程外键ID"`
	CourseVideo       string      `json:"courseVideo"       orm:"course_video"        description:"课程视频"`
	CourseVideoSource int         `json:"courseVideoSource" orm:"course_video_source" description:"视频来源"`
	ShowCourseVideo   int         `json:"showCourseVideo"   orm:"show_course_video"   description:"对学生显示课程视频"`
	CourseCase        string      `json:"courseCase"        orm:"course_case"         description:"课程作业案例"`
	ShowCourseCase    int         `json:"showCourseCase"    orm:"show_course_case"    description:"对学生显示课程案例"`
	CoursePpt         string      `json:"coursePpt"         orm:"course_ppt"          description:"课件PPT"`
	ShowCoursePpt     int         `json:"showCoursePpt"     orm:"show_course_ppt"     description:"对学生显示课程资料"`
	CourseWorkType    int         `json:"courseWorkType"    orm:"course_work_type"    description:"作业类型"`
	CourseWork        string      `json:"courseWork"        orm:"course_work"         description:"课程作业"`
	CourseWorkAnswer  string      `json:"courseWorkAnswer"  orm:"course_work_answer"  description:"课程作业答案"`
	CoursePlan        string      `json:"coursePlan"        orm:"course_plan"         description:"教案"`
	ShowCoursePlan    int         `json:"showCoursePlan"    orm:"show_course_plan"    description:"对学生显示教案"`
	MapX              int         `json:"mapX"              orm:"map_x"               description:"地图X坐标"`
	MapY              int         `json:"mapY"              orm:"map_y"               description:"地图Y坐标"`
	MediaContent      string      `json:"mediaContent"      orm:"media_content"       description:"富文本课件"`
	OrderNum          int         `json:"orderNum"          orm:"order_num"           description:"排序"`
}