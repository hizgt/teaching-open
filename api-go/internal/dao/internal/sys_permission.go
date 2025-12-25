package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SysPermissionDao 权限DAO
type SysPermissionDao struct {
	table   string
	group   string
	columns SysPermissionColumns
}

// SysPermissionColumns 权限表字段
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
	DelFlag            string
	RuleFlag           string
	Status             string
	InternalOrExternal string
	CreateBy           string
	CreateTime         string
	UpdateBy           string
	UpdateTime         string
}

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
	DelFlag:            "del_flag",
	RuleFlag:           "rule_flag",
	Status:             "status",
	InternalOrExternal: "internal_or_external",
	CreateBy:           "create_by",
	CreateTime:         "create_time",
	UpdateBy:           "update_by",
	UpdateTime:         "update_time",
}

// NewSysPermissionDao 创建权限DAO
func NewSysPermissionDao() *SysPermissionDao {
	return &SysPermissionDao{
		table:   "sys_permission",
		group:   "default",
		columns: sysPermissionColumns,
	}
}

// DB 获取数据库连接
func (d *SysPermissionDao) DB() gdb.DB {
	return g.DB(d.group)
}

// Table 获取表名
func (d *SysPermissionDao) Table() string {
	return d.table
}

// Columns 获取字段
func (d *SysPermissionDao) Columns() SysPermissionColumns {
	return d.columns
}

// Ctx 获取上下文Model
func (d *SysPermissionDao) Ctx(ctx context.Context) *gdb.Model {
	return d.DB().Model(d.table).Safe().Ctx(ctx)
}

// SysRolePermissionDao 角色权限关联DAO
type SysRolePermissionDao struct {
	table   string
	group   string
	columns SysRolePermissionColumns
}

// SysRolePermissionColumns 角色权限关联表字段
type SysRolePermissionColumns struct {
	Id           string
	RoleId       string
	PermissionId string
	DataRuleIds  string
	OperateDate  string
	OperateIp    string
}

var sysRolePermissionColumns = SysRolePermissionColumns{
	Id:           "id",
	RoleId:       "role_id",
	PermissionId: "permission_id",
	DataRuleIds:  "data_rule_ids",
	OperateDate:  "operate_date",
	OperateIp:    "operate_ip",
}

// NewSysRolePermissionDao 创建角色权限关联DAO
func NewSysRolePermissionDao() *SysRolePermissionDao {
	return &SysRolePermissionDao{
		table:   "sys_role_permission",
		group:   "default",
		columns: sysRolePermissionColumns,
	}
}

// DB 获取数据库连接
func (d *SysRolePermissionDao) DB() gdb.DB {
	return g.DB(d.group)
}

// Table 获取表名
func (d *SysRolePermissionDao) Table() string {
	return d.table
}

// Columns 获取字段
func (d *SysRolePermissionDao) Columns() SysRolePermissionColumns {
	return d.columns
}

// Ctx 获取上下文Model
func (d *SysRolePermissionDao) Ctx(ctx context.Context) *gdb.Model {
	return d.DB().Model(d.table).Safe().Ctx(ctx)
}
