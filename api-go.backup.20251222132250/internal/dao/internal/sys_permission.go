// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SysPermissionDao is the data access object for table sys_permission.
type SysPermissionDao struct {
	table   string               // table is the underlying table name of the DAO.
	group   string               // group is the database configuration group name of current DAO.
	columns SysPermissionColumns // columns contains all the column names of Table for convenient usage.
}

// SysPermissionColumns defines and stores column names for table sys_permission.
type SysPermissionColumns struct {
	Id                 string
	ParentId           string
	Name               string
	Url                string
	Component          string
	ComponentName      string
	Redirect           string
	MenuType           string
	Perms              string
	PermsType          string
	SortNo             string
	AlwaysShow         string
	Icon               string
	IsRoute            string
	IsLeaf             string
	KeepAlive          string
	Hidden             string
	Description        string
	CreateBy           string
	CreateTime         string
	UpdateBy           string
	UpdateTime         string
	DelFlag            string
	RuleFlag           string
	Status             string
	InternalOrExternal string
}

// sysPermissionColumns holds the columns for table sys_permission.
var sysPermissionColumns = SysPermissionColumns{
	Id:                 "id",
	ParentId:           "parent_id",
	Name:               "name",
	Url:                "url",
	Component:          "component",
	ComponentName:      "component_name",
	Redirect:           "redirect",
	MenuType:           "menu_type",
	Perms:              "perms",
	PermsType:          "perms_type",
	SortNo:             "sort_no",
	AlwaysShow:         "always_show",
	Icon:               "icon",
	IsRoute:            "is_route",
	IsLeaf:             "is_leaf",
	KeepAlive:          "keep_alive",
	Hidden:             "hidden",
	Description:        "description",
	CreateBy:           "create_by",
	CreateTime:         "create_time",
	UpdateBy:           "update_by",
	UpdateTime:         "update_time",
	DelFlag:            "del_flag",
	RuleFlag:           "rule_flag",
	Status:             "status",
	InternalOrExternal: "internal_or_external",
}

// NewSysPermissionDao creates and returns a new DAO object for table data access.
func NewSysPermissionDao() *SysPermissionDao {
	return &SysPermissionDao{
		group:   "default",
		table:   "sys_permission",
		columns: sysPermissionColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *SysPermissionDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *SysPermissionDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *SysPermissionDao) Columns() SysPermissionColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *SysPermissionDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *SysPermissionDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
func (dao *SysPermissionDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
