package dao

import (
	"teaching-open/internal/dao/internal"
)

var (
	// SysDepart 部门DAO
	SysDepart *internal.SysDepartDao
	// SysUserDepart 用户部门关联DAO
	SysUserDepart *internal.SysUserDepartDao
)

func init() {
	SysDepart = internal.NewSysDepartDao()
	SysUserDepart = internal.NewSysUserDepartDao()
}
