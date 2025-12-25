package dao

import (
	"teaching-open/internal/dao/internal"
)

var (
	// SysPermission 权限DAO
	SysPermission *internal.SysPermissionDao
	// SysRolePermission 角色权限关联DAO
	SysRolePermission *internal.SysRolePermissionDao
)

func init() {
	SysPermission = internal.NewSysPermissionDao()
	SysRolePermission = internal.NewSysRolePermissionDao()
}
