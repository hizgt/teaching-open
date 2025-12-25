package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SysDictDao 字典DAO
type SysDictDao struct {
	table   string
	group   string
	columns SysDictColumns
}

// SysDictColumns 字典表字段
type SysDictColumns struct {
	Id          string
	DictName    string
	DictCode    string
	Description string
	Type        string
	DelFlag     string
	CreateBy    string
	CreateTime  string
	UpdateBy    string
	UpdateTime  string
}

var sysDictColumns = SysDictColumns{
	Id:          "id",
	DictName:    "dict_name",
	DictCode:    "dict_code",
	Description: "description",
	Type:        "type",
	DelFlag:     "del_flag",
	CreateBy:    "create_by",
	CreateTime:  "create_time",
	UpdateBy:    "update_by",
	UpdateTime:  "update_time",
}

// NewSysDictDao 创建字典DAO
func NewSysDictDao() *SysDictDao {
	return &SysDictDao{
		table:   "sys_dict",
		group:   "default",
		columns: sysDictColumns,
	}
}

// DB 获取数据库连接
func (d *SysDictDao) DB() gdb.DB {
	return g.DB(d.group)
}

// Table 获取表名
func (d *SysDictDao) Table() string {
	return d.table
}

// Columns 获取字段
func (d *SysDictDao) Columns() SysDictColumns {
	return d.columns
}

// Ctx 获取上下文Model
func (d *SysDictDao) Ctx(ctx context.Context) *gdb.Model {
	return d.DB().Model(d.table).Safe().Ctx(ctx)
}

// SysDictItemDao 字典项DAO
type SysDictItemDao struct {
	table   string
	group   string
	columns SysDictItemColumns
}

// SysDictItemColumns 字典项表字段
type SysDictItemColumns struct {
	Id          string
	DictId      string
	ItemText    string
	ItemValue   string
	Description string
	SortOrder   string
	Status      string
	CreateBy    string
	CreateTime  string
	UpdateBy    string
	UpdateTime  string
}

var sysDictItemColumns = SysDictItemColumns{
	Id:          "id",
	DictId:      "dict_id",
	ItemText:    "item_text",
	ItemValue:   "item_value",
	Description: "description",
	SortOrder:   "sort_order",
	Status:      "status",
	CreateBy:    "create_by",
	CreateTime:  "create_time",
	UpdateBy:    "update_by",
	UpdateTime:  "update_time",
}

// NewSysDictItemDao 创建字典项DAO
func NewSysDictItemDao() *SysDictItemDao {
	return &SysDictItemDao{
		table:   "sys_dict_item",
		group:   "default",
		columns: sysDictItemColumns,
	}
}

// DB 获取数据库连接
func (d *SysDictItemDao) DB() gdb.DB {
	return g.DB(d.group)
}

// Table 获取表名
func (d *SysDictItemDao) Table() string {
	return d.table
}

// Columns 获取字段
func (d *SysDictItemDao) Columns() SysDictItemColumns {
	return d.columns
}

// Ctx 获取上下文Model
func (d *SysDictItemDao) Ctx(ctx context.Context) *gdb.Model {
	return d.DB().Model(d.table).Safe().Ctx(ctx)
}
