// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SysRoleDao is the data access object for table sys_role.
type SysRoleDao struct {
	table   string         // table is the underlying table name of the DAO.
	group   string         // group is the database configuration group name of current DAO.
	columns SysRoleColumns // columns contains all the column names of Table for convenient usage.
}

// SysRoleColumns defines and stores column names for table sys_role.
type SysRoleColumns struct {
	Id          string
	RoleName    string
	RoleCode    string
	Description string
	CreateBy    string
	CreateTime  string
	UpdateBy    string
	UpdateTime  string
	RoleLevel   string
}

// sysRoleColumns holds the columns for table sys_role.
var sysRoleColumns = SysRoleColumns{
	Id:          "id",
	RoleName:    "role_name",
	RoleCode:    "role_code",
	Description: "description",
	CreateBy:    "create_by",
	CreateTime:  "create_time",
	UpdateBy:    "update_by",
	UpdateTime:  "update_time",
	RoleLevel:   "role_level",
}

// NewSysRoleDao creates and returns a new DAO object for table data access.
func NewSysRoleDao() *SysRoleDao {
	return &SysRoleDao{
		group:   "default",
		table:   "sys_role",
		columns: sysRoleColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *SysRoleDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *SysRoleDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *SysRoleDao) Columns() SysRoleColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *SysRoleDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *SysRoleDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
func (dao *SysRoleDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
