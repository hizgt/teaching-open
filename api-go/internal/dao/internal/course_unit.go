// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// CourseUnitDao is the data access object for table teaching_course_unit.
type CourseUnitDao struct {
	table   string             // table is the underlying table name of the DAO.
	group   string             // group is the database configuration group name of current DAO.
	columns CourseUnitColumns  // columns contains all the column names of Table for convenient usage.
}

// CourseUnitColumns defines and stores column names for table teaching_course_unit.
type CourseUnitColumns struct {
	Id          string //
	CourseId    string //
	Name        string //
	Content     string //
	SortOrder   string //
	ResourceUrl string //
	CreateBy    string //
	CreateTime  string //
	UpdateBy    string //
	UpdateTime  string //
}

// courseUnitColumns holds the columns for table teaching_course_unit.
var courseUnitColumns = CourseUnitColumns{
	Id:          "id",
	CourseId:    "course_id",
	Name:        "name",
	Content:     "content",
	SortOrder:   "sort_order",
	ResourceUrl: "resource_url",
	CreateBy:    "create_by",
	CreateTime:  "create_time",
	UpdateBy:    "update_by",
	UpdateTime:  "update_time",
}

// NewCourseUnitDao creates and returns a new DAO object for table data access.
func NewCourseUnitDao() *CourseUnitDao {
	return &CourseUnitDao{
		group:   "default",
		table:   "teaching_course_unit",
		columns: courseUnitColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *CourseUnitDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *CourseUnitDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *CourseUnitDao) Columns() CourseUnitColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *CourseUnitDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *CourseUnitDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic.
func (dao *CourseUnitDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}