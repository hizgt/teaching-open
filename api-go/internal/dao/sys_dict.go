package dao

import (
	"teaching-open/internal/dao/internal"
)

var (
	// SysDict 字典DAO
	SysDict *internal.SysDictDao
	// SysDictItem 字典项DAO
	SysDictItem *internal.SysDictItemDao
)

func init() {
	SysDict = internal.NewSysDictDao()
	SysDictItem = internal.NewSysDictItemDao()
}
