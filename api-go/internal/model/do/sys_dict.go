// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysDict is the golang structure of table sys_dict for DAO operations like Where/Data.
type SysDict struct {
	g.Meta      `orm:"table:sys_dict, do:true"`
	Id          interface{}
	DictName    interface{}
	DictCode    interface{}
	Description interface{}
	DelFlag     interface{}
	CreateBy    interface{}
	CreateTime  *gtime.Time
	UpdateBy    interface{}
	UpdateTime  *gtime.Time
	Type        interface{}
}
