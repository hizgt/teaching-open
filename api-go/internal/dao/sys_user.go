package dao

import (
	"teaching-open/internal/dao/internal"
)

var (
	// SysUser 用户DAO
	SysUser *internal.SysUserDao
)

func init() {
	SysUser = internal.NewSysUserDao()
}
