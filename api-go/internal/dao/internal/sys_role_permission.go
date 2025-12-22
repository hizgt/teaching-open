// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SysRolePermissionDao is the data access object for table sys_role_permission.
type SysRolePermissionDao struct {
	table   string                   // table is the underlying table name of the DAO.
	group   string                   // group is the database configuration group name of current DAO.
	columns SysRolePermissionColumns // columns contains all the column names of Table for convenient usage.
}

// SysRolePermissionColumns defines and stores column names for table sys_role_permission.
type SysRolePermissionColumns struct {
	Id           string
	RoleId       string
	PermissionId string
	DataRuleIds  string
}

// sysRolePermissionColumns holds the columns for table sys_role_permission.
var sysRolePermissionColumns = SysRolePermissionColumns{
	Id:           "id",
	RoleId:       "role_id",
	PermissionId: "permission_id",
	DataRuleIds:  "data_rule_ids",
}

// NewSysRolePermissionDao creates and returns a new DAO object for table data access.
func NewSysRolePermissionDao() *SysRolePermissionDao {
	return &SysRolePermissionDao{
		group:   "default",
		table:   "sys_role_permission",
		columns: sysRolePermissionColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *SysRolePermissionDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *SysRolePermissionDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *SysRolePermissionDao) Columns() SysRolePermissionColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *SysRolePermissionDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *SysRolePermissionDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
func (dao *SysRolePermissionDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
