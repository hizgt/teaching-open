package dao

import (
	"teaching-open/internal/dao/internal"
)

var (
	// SysRole 角色DAO
	SysRole *internal.SysRoleDao
	// SysUserRole 用户角色关联DAO
	SysUserRole *internal.SysUserRoleDao
)

func init() {
	SysRole = internal.NewSysRoleDao()
	SysUserRole = internal.NewSysUserRoleDao()
}
