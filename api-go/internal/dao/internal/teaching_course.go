// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// TeachingCourseDao is the data access object for table teaching_course.
type TeachingCourseDao struct {
	table   string                // table is the underlying table name of the DAO.
	group   string                // group is the database configuration group name of current DAO.
	columns TeachingCourseColumns // columns contains all the column names of Table for convenient usage.
}

// TeachingCourseColumns defines and stores column names for table teaching_course.
type TeachingCourseColumns struct {
	Id             string
	CreateBy       string
	CreateTime     string
	UpdateBy       string
	UpdateTime     string
	SysOrgCode     string
	DelFlag        string
	CourseName     string
	CourseDesc     string
	CourseIcon     string
	CourseCover    string
	ShowType       string
	CourseMap      string
	IsShared       string
	ShowHome       string
	OrderNum       string
	DepartIds      string
	CourseType     string
	CourseCategory string
}

// teachingCourseColumns holds the columns for table teaching_course.
var teachingCourseColumns = TeachingCourseColumns{
	Id:             "id",
	CreateBy:       "create_by",
	CreateTime:     "create_time",
	UpdateBy:       "update_by",
	UpdateTime:     "update_time",
	SysOrgCode:     "sys_org_code",
	DelFlag:        "del_flag",
	CourseName:     "course_name",
	CourseDesc:     "course_desc",
	CourseIcon:     "course_icon",
	CourseCover:    "course_cover",
	ShowType:       "show_type",
	CourseMap:      "course_map",
	IsShared:       "is_shared",
	ShowHome:       "show_home",
	OrderNum:       "order_num",
	DepartIds:      "depart_ids",
	CourseType:     "course_type",
	CourseCategory: "course_category",
}

// NewTeachingCourseDao creates and returns a new DAO object for table data access.
func NewTeachingCourseDao() *TeachingCourseDao {
	return &TeachingCourseDao{
		group:   "default",
		table:   "teaching_course",
		columns: teachingCourseColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *TeachingCourseDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *TeachingCourseDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *TeachingCourseDao) Columns() TeachingCourseColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *TeachingCourseDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *TeachingCourseDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
func (dao *TeachingCourseDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
