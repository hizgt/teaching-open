// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// TeachingWorkCorrectDao is the data access object for table teaching_work_correct.
type TeachingWorkCorrectDao struct {
	table   string                     // table is the underlying table name of the DAO.
	group   string                     // group is the database configuration group name of current DAO.
	columns TeachingWorkCorrectColumns // columns contains all the column names of Table for convenient usage.
}

// TeachingWorkCorrectColumns defines and stores column names for table teaching_work_correct.
type TeachingWorkCorrectColumns struct {
	Id         string // 主键
	CreateBy   string // 创建人
	CreateTime string // 创建日期
	UpdateBy   string // 更新人
	UpdateTime string // 更新日期
	SysOrgCode string // 所属部门
	WorkId     string // 作业ID
	Score      string // 评分
	Comment    string // 评语
}

// teachingWorkCorrectColumns holds the columns for table teaching_work_correct.
var teachingWorkCorrectColumns = TeachingWorkCorrectColumns{
	Id:         "id",
	CreateBy:   "create_by",
	CreateTime: "create_time",
	UpdateBy:   "update_by",
	UpdateTime: "update_time",
	SysOrgCode: "sys_org_code",
	WorkId:     "work_id",
	Score:      "score",
	Comment:    "comment",
}

// NewTeachingWorkCorrectDao creates and returns a new DAO object for table data access.
func NewTeachingWorkCorrectDao() *TeachingWorkCorrectDao {
	return &TeachingWorkCorrectDao{
		group:   "default",
		table:   "teaching_work_correct",
		columns: teachingWorkCorrectColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *TeachingWorkCorrectDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current DAO.
func (dao *TeachingWorkCorrectDao) Table() string {
	return dao.table
}

// Columns returns all column names of current DAO.
func (dao *TeachingWorkCorrectDao) Columns() TeachingWorkCorrectColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current DAO.
func (dao *TeachingWorkCorrectDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for current DAO. It automatically sets the context for current operation.
func (dao *TeachingWorkCorrectDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
func (dao *TeachingWorkCorrectDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}