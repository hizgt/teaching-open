// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysDictItem is the golang structure of table sys_dict_item for DAO operations like Where/Data.
type SysDictItem struct {
	g.Meta      `orm:"table:sys_dict_item, do:true"`
	Id          interface{}
	DictId      interface{}
	ItemText    interface{}
	ItemValue   interface{}
	Description interface{}
	SortOrder   interface{}
	Status      interface{}
	CreateBy    interface{}
	CreateTime  *gtime.Time
	UpdateBy    interface{}
	UpdateTime  *gtime.Time
}
