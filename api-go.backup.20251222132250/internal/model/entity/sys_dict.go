// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysDict is the golang structure for table sys_dict.
type SysDict struct {
	Id          string      `json:"id"          description:"ID"`
	DictName    string      `json:"dictName"    description:"字典名称"`
	DictCode    string      `json:"dictCode"    description:"字典编码"`
	Description string      `json:"description" description:"描述"`
	DelFlag     int         `json:"delFlag"     description:"删除状态"`
	CreateBy    string      `json:"createBy"    description:"创建人"`
	CreateTime  *gtime.Time `json:"createTime"  description:"创建时间"`
	UpdateBy    string      `json:"updateBy"    description:"更新人"`
	UpdateTime  *gtime.Time `json:"updateTime"  description:"更新时间"`
	Type        int         `json:"type"        description:"字典类型0为string,1为number"`
}
