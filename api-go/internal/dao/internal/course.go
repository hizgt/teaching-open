// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// CourseDao is the data access object for table teaching_course.
type CourseDao struct {
	table   string          // table is the underlying table name of the DAO.
	group   string          // group is the database configuration group name of current DAO.
	columns CourseColumns   // columns contains all the column names of Table for convenient usage.
}

// CourseColumns defines and stores column names for table teaching_course.
type CourseColumns struct {
	Id          string //
	Name        string //
	Type        string //
	Description string //
	CoverImage  string //
	Status      string //
	CreateBy    string //
	CreateTime  string //
	UpdateBy    string //
	UpdateTime  string //
}

// courseColumns holds the columns for table teaching_course.
var courseColumns = CourseColumns{
	Id:          "id",
	Name:        "name",
	Type:        "type",
	Description: "description",
	CoverImage:  "cover_image",
	Status:      "status",
	CreateBy:    "create_by",
	CreateTime:  "create_time",
	UpdateBy:    "update_by",
	UpdateTime:  "update_time",
}

// NewCourseDao creates and returns a new DAO object for table data access.
func NewCourseDao() *CourseDao {
	return &CourseDao{
		group:   "default",
		table:   "teaching_course",
		columns: courseColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *CourseDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *CourseDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *CourseDao) Columns() CourseColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *CourseDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *CourseDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic.
func (dao *CourseDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}