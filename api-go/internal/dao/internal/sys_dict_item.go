// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SysDictItemDao is the data access object for table sys_dict_item.
type SysDictItemDao struct {
	table   string             // table is the underlying table name of the DAO.
	group   string             // group is the database configuration group name of current DAO.
	columns SysDictItemColumns // columns contains all the column names of Table for convenient usage.
}

// SysDictItemColumns defines and stores column names for table sys_dict_item.
type SysDictItemColumns struct {
	Id          string
	DictId      string
	ItemText    string
	ItemValue   string
	Description string
	SortOrder   string
	Status      string
	CreateBy    string
	CreateTime  string
	UpdateBy    string
	UpdateTime  string
}

// sysDictItemColumns holds the columns for table sys_dict_item.
var sysDictItemColumns = SysDictItemColumns{
	Id:          "id",
	DictId:      "dict_id",
	ItemText:    "item_text",
	ItemValue:   "item_value",
	Description: "description",
	SortOrder:   "sort_order",
	Status:      "status",
	CreateBy:    "create_by",
	CreateTime:  "create_time",
	UpdateBy:    "update_by",
	UpdateTime:  "update_time",
}

// NewSysDictItemDao creates and returns a new DAO object for table data access.
func NewSysDictItemDao() *SysDictItemDao {
	return &SysDictItemDao{
		group:   "default",
		table:   "sys_dict_item",
		columns: sysDictItemColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *SysDictItemDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *SysDictItemDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *SysDictItemDao) Columns() SysDictItemColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *SysDictItemDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *SysDictItemDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
func (dao *SysDictItemDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
