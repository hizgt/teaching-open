// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// TeachingCourseUnitDao is the data access object for table teaching_course_unit.
type TeachingCourseUnitDao struct {
	table   string
	group   string
	columns TeachingCourseUnitColumns
}

// TeachingCourseUnitColumns defines and stores column names for table teaching_course_unit.
type TeachingCourseUnitColumns struct {
	Id                string
	CreateBy          string
	CreateTime        string
	UpdateBy          string
	UpdateTime        string
	SysOrgCode        string
	DelFlag           string
	UnitName          string
	UnitIntro         string
	UnitCover         string
	CourseId          string
	CourseVideo       string
	CourseVideoSource string
	ShowCourseVideo   string
	CourseCase        string
	ShowCourseCase    string
	CoursePpt         string
	ShowCoursePpt     string
	CourseWorkType    string
	CourseWork        string
	CourseWorkAnswer  string
	CoursePlan        string
	ShowCoursePlan    string
	MapX              string
	MapY              string
	MediaContent      string
	OrderNum          string
}

var teachingCourseUnitColumns = TeachingCourseUnitColumns{
	Id:                "id",
	CreateBy:          "create_by",
	CreateTime:        "create_time",
	UpdateBy:          "update_by",
	UpdateTime:        "update_time",
	SysOrgCode:        "sys_org_code",
	DelFlag:           "del_flag",
	UnitName:          "unit_name",
	UnitIntro:         "unit_intro",
	UnitCover:         "unit_cover",
	CourseId:          "course_id",
	CourseVideo:       "course_video",
	CourseVideoSource: "course_video_source",
	ShowCourseVideo:   "show_course_video",
	CourseCase:        "course_case",
	ShowCourseCase:    "show_course_case",
	CoursePpt:         "course_ppt",
	ShowCoursePpt:     "show_course_ppt",
	CourseWorkType:    "course_work_type",
	CourseWork:        "course_work",
	CourseWorkAnswer:  "course_work_answer",
	CoursePlan:        "course_plan",
	ShowCoursePlan:    "show_course_plan",
	MapX:              "map_x",
	MapY:              "map_y",
	MediaContent:      "media_content",
	OrderNum:          "order_num",
}

func NewTeachingCourseUnitDao() *TeachingCourseUnitDao {
	return &TeachingCourseUnitDao{
		group:   "default",
		table:   "teaching_course_unit",
		columns: teachingCourseUnitColumns,
	}
}

func (dao *TeachingCourseUnitDao) DB() gdb.DB {
	return g.DB(dao.group)
}

func (dao *TeachingCourseUnitDao) Table() string {
	return dao.table
}

func (dao *TeachingCourseUnitDao) Columns() TeachingCourseUnitColumns {
	return dao.columns
}

func (dao *TeachingCourseUnitDao) Group() string {
	return dao.group
}

func (dao *TeachingCourseUnitDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

func (dao *TeachingCourseUnitDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}