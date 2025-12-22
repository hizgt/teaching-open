// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SysDictDao is the data access object for table sys_dict.
type SysDictDao struct {
	table   string         // table is the underlying table name of the DAO.
	group   string         // group is the database configuration group name of current DAO.
	columns SysDictColumns // columns contains all the column names of Table for convenient usage.
}

// SysDictColumns defines and stores column names for table sys_dict.
type SysDictColumns struct {
	Id          string
	DictName    string
	DictCode    string
	Description string
	DelFlag     string
	CreateBy    string
	CreateTime  string
	UpdateBy    string
	UpdateTime  string
	Type        string
}

// sysDictColumns holds the columns for table sys_dict.
var sysDictColumns = SysDictColumns{
	Id:          "id",
	DictName:    "dict_name",
	DictCode:    "dict_code",
	Description: "description",
	DelFlag:     "del_flag",
	CreateBy:    "create_by",
	CreateTime:  "create_time",
	UpdateBy:    "update_by",
	UpdateTime:  "update_time",
	Type:        "type",
}

// NewSysDictDao creates and returns a new DAO object for table data access.
func NewSysDictDao() *SysDictDao {
	return &SysDictDao{
		group:   "default",
		table:   "sys_dict",
		columns: sysDictColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *SysDictDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *SysDictDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *SysDictDao) Columns() SysDictColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *SysDictDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *SysDictDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
func (dao *SysDictDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
