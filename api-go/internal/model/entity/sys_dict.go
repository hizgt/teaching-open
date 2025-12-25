package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysDict 字典表
type SysDict struct {
	Id          string      `json:"id"          orm:"id,primary"    description:"字典ID"`
	DictName    string      `json:"dictName"    orm:"dict_name"     description:"字典名称"`
	DictCode    string      `json:"dictCode"    orm:"dict_code"     description:"字典编码"`
	Description string      `json:"description" orm:"description"   description:"描述"`
	Type        int         `json:"type"        orm:"type"          description:"类型(0字符1数字)"`
	DelFlag     int         `json:"delFlag"     orm:"del_flag"      description:"删除标记"`
	CreateBy    string      `json:"createBy"    orm:"create_by"     description:"创建人"`
	CreateTime  *gtime.Time `json:"createTime"  orm:"create_time"   description:"创建时间"`
	UpdateBy    string      `json:"updateBy"    orm:"update_by"     description:"更新人"`
	UpdateTime  *gtime.Time `json:"updateTime"  orm:"update_time"   description:"更新时间"`
}

// TableName 表名
func (e *SysDict) TableName() string {
	return "sys_dict"
}

// SysDictItem 字典项表
type SysDictItem struct {
	Id          string      `json:"id"          orm:"id,primary"    description:"字典项ID"`
	DictId      string      `json:"dictId"      orm:"dict_id"       description:"字典ID"`
	ItemText    string      `json:"itemText"    orm:"item_text"     description:"字典项文本"`
	ItemValue   string      `json:"itemValue"   orm:"item_value"    description:"字典项值"`
	Description string      `json:"description" orm:"description"   description:"描述"`
	SortOrder   int         `json:"sortOrder"   orm:"sort_order"    description:"排序"`
	Status      int         `json:"status"      orm:"status"        description:"状态(1启用0禁用)"`
	CreateBy    string      `json:"createBy"    orm:"create_by"     description:"创建人"`
	CreateTime  *gtime.Time `json:"createTime"  orm:"create_time"   description:"创建时间"`
	UpdateBy    string      `json:"updateBy"    orm:"update_by"     description:"更新人"`
	UpdateTime  *gtime.Time `json:"updateTime"  orm:"update_time"   description:"更新时间"`
}

// TableName 表名
func (e *SysDictItem) TableName() string {
	return "sys_dict_item"
}
