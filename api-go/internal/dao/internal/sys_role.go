package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SysRoleDao 角色DAO
type SysRoleDao struct {
	table   string
	group   string
	columns SysRoleColumns
}

// SysRoleColumns 角色表字段
type SysRoleColumns struct {
	Id          string
	RoleName    string
	RoleCode    string
	RoleLevel   string
	Description string
	CreateBy    string
	CreateTime  string
	UpdateBy    string
	UpdateTime  string
}

var sysRoleColumns = SysRoleColumns{
	Id:          "id",
	RoleName:    "role_name",
	RoleCode:    "role_code",
	RoleLevel:   "role_level",
	Description: "description",
	CreateBy:    "create_by",
	CreateTime:  "create_time",
	UpdateBy:    "update_by",
	UpdateTime:  "update_time",
}

// NewSysRoleDao 创建角色DAO
func NewSysRoleDao() *SysRoleDao {
	return &SysRoleDao{
		table:   "sys_role",
		group:   "default",
		columns: sysRoleColumns,
	}
}

// DB 获取数据库连接
func (d *SysRoleDao) DB() gdb.DB {
	return g.DB(d.group)
}

// Table 获取表名
func (d *SysRoleDao) Table() string {
	return d.table
}

// Columns 获取字段
func (d *SysRoleDao) Columns() SysRoleColumns {
	return d.columns
}

// Ctx 获取上下文Model
func (d *SysRoleDao) Ctx(ctx context.Context) *gdb.Model {
	return d.DB().Model(d.table).Safe().Ctx(ctx)
}

// SysUserRoleDao 用户角色关联DAO
type SysUserRoleDao struct {
	table   string
	group   string
	columns SysUserRoleColumns
}

// SysUserRoleColumns 用户角色关联表字段
type SysUserRoleColumns struct {
	Id     string
	UserId string
	RoleId string
}

var sysUserRoleColumns = SysUserRoleColumns{
	Id:     "id",
	UserId: "user_id",
	RoleId: "role_id",
}

// NewSysUserRoleDao 创建用户角色关联DAO
func NewSysUserRoleDao() *SysUserRoleDao {
	return &SysUserRoleDao{
		table:   "sys_user_role",
		group:   "default",
		columns: sysUserRoleColumns,
	}
}

// DB 获取数据库连接
func (d *SysUserRoleDao) DB() gdb.DB {
	return g.DB(d.group)
}

// Table 获取表名
func (d *SysUserRoleDao) Table() string {
	return d.table
}

// Columns 获取字段
func (d *SysUserRoleDao) Columns() SysUserRoleColumns {
	return d.columns
}

// Ctx 获取上下文Model
func (d *SysUserRoleDao) Ctx(ctx context.Context) *gdb.Model {
	return d.DB().Model(d.table).Safe().Ctx(ctx)
}
