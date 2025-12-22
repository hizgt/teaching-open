// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// TeachingDepartDayLogDao is the data access object for table teaching_depart_day_log.
type TeachingDepartDayLogDao struct {
	table   string                      // table is the underlying table name of the DAO.
	group   string                      // group is the database configuration group name of current DAO.
	columns TeachingDepartDayLogColumns // columns contains all the column names of Table for convenient usage.
}

// TeachingDepartDayLogColumns defines and stores column names for table teaching_depart_day_log.
type TeachingDepartDayLogColumns struct {
	Id                         string // 主键ID
	DepartId                   string // 班级ID
	DepartName                 string // 班级名
	UnitOpenCount              string // 开课次数
	CourseWorkAssignCount      string // 课程作业布置次数
	AdditionalWorkAssignCount  string // 附加作业布置次数
	CourseWorkCorrectCount     string // 课程作业批改次数
	AdditionalWorkCorrectCount string // 附加作业批改次数
	CourseWorkSubmitCount      string // 课程作业提交次数
	AdditionalWorkSubmitCount  string // 附加作业提交次数
	CreateTime                 string // 日期
}

// teachingDepartDayLogColumns holds the columns for table teaching_depart_day_log.
var teachingDepartDayLogColumns = TeachingDepartDayLogColumns{
	Id:                         "id",
	DepartId:                   "depart_id",
	DepartName:                 "depart_name",
	UnitOpenCount:              "unit_open_count",
	CourseWorkAssignCount:      "course_work_assign_count",
	AdditionalWorkAssignCount:  "additional_work_assign_count",
	CourseWorkCorrectCount:     "course_work_correct_count",
	AdditionalWorkCorrectCount: "additional_work_correct_count",
	CourseWorkSubmitCount:      "course_work_submit_count",
	AdditionalWorkSubmitCount:  "additional_work_submit_count",
	CreateTime:                 "create_time",
}

// NewTeachingDepartDayLogDao creates and returns a new DAO object for table data access.
func NewTeachingDepartDayLogDao() *TeachingDepartDayLogDao {
	return &TeachingDepartDayLogDao{
		group:   "default",
		table:   "teaching_depart_day_log",
		columns: teachingDepartDayLogColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *TeachingDepartDayLogDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *TeachingDepartDayLogDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *TeachingDepartDayLogDao) Columns() TeachingDepartDayLogColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *TeachingDepartDayLogDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *TeachingDepartDayLogDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
func (dao *TeachingDepartDayLogDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
