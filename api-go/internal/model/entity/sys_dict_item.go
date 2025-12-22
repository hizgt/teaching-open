// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysDictItem is the golang structure for table sys_dict_item.
type SysDictItem struct {
	Id          string      `json:"id"          description:"ID"`
	DictId      string      `json:"dictId"      description:"字典id"`
	ItemText    string      `json:"itemText"    description:"字典项文本"`
	ItemValue   string      `json:"itemValue"   description:"字典项值"`
	Description string      `json:"description" description:"描述"`
	SortOrder   int         `json:"sortOrder"   description:"排序"`
	Status      int         `json:"status"      description:"状态（1启用 0不启用）"`
	CreateBy    string      `json:"createBy"    description:"创建人"`
	CreateTime  *gtime.Time `json:"createTime"  description:"创建时间"`
	UpdateBy    string      `json:"updateBy"    description:"更新人"`
	UpdateTime  *gtime.Time `json:"updateTime"  description:"更新时间"`
}
