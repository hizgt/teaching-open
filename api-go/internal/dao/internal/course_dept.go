// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// CourseDeptDao is the data access object for table teaching_course_dept.
type CourseDeptDao struct {
	table   string             // table is the underlying table name of the DAO.
	group   string             // group is the database configuration group name of current DAO.
	columns CourseDeptColumns  // columns contains all the column names of Table for convenient usage.
}

// CourseDeptColumns defines and stores column names for table teaching_course_dept.
type CourseDeptColumns struct {
	Id        string //
	CourseId  string //
	DeptId    string //
	CreateBy  string //
	CreateTime string //
}

// courseDeptColumns holds the columns for table teaching_course_dept.
var courseDeptColumns = CourseDeptColumns{
	Id:         "id",
	CourseId:   "course_id",
	DeptId:     "dept_id",
	CreateBy:   "create_by",
	CreateTime: "create_time",
}

// NewCourseDeptDao creates and returns a new DAO object for table data access.
func NewCourseDeptDao() *CourseDeptDao {
	return &CourseDeptDao{
		group:   "default",
		table:   "teaching_course_dept",
		columns: courseDeptColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *CourseDeptDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *CourseDeptDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *CourseDeptDao) Columns() CourseDeptColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *CourseDeptDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *CourseDeptDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic.
func (dao *CourseDeptDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}