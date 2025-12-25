package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SysDepartDao 部门DAO
type SysDepartDao struct {
	table   string
	group   string
	columns SysDepartColumns
}

// SysDepartColumns 部门表字段
type SysDepartColumns struct {
	Id             string
	ParentId       string
	DepartName     string
	DepartNameEn   string
	DepartNameAbbr string
	DepartOrder    string
	Description    string
	OrgCategory    string
	OrgType        string
	OrgCode        string
	Mobile         string
	Fax            string
	Address        string
	Memo           string
	Status         string
	DelFlag        string
	CreateBy       string
	CreateTime     string
	UpdateBy       string
	UpdateTime     string
}

var sysDepartColumns = SysDepartColumns{
	Id:             "id",
	ParentId:       "parent_id",
	DepartName:     "depart_name",
	DepartNameEn:   "depart_name_en",
	DepartNameAbbr: "depart_name_abbr",
	DepartOrder:    "depart_order",
	Description:    "description",
	OrgCategory:    "org_category",
	OrgType:        "org_type",
	OrgCode:        "org_code",
	Mobile:         "mobile",
	Fax:            "fax",
	Address:        "address",
	Memo:           "memo",
	Status:         "status",
	DelFlag:        "del_flag",
	CreateBy:       "create_by",
	CreateTime:     "create_time",
	UpdateBy:       "update_by",
	UpdateTime:     "update_time",
}

// NewSysDepartDao 创建部门DAO
func NewSysDepartDao() *SysDepartDao {
	return &SysDepartDao{
		table:   "sys_depart",
		group:   "default",
		columns: sysDepartColumns,
	}
}

// DB 获取数据库连接
func (d *SysDepartDao) DB() gdb.DB {
	return g.DB(d.group)
}

// Table 获取表名
func (d *SysDepartDao) Table() string {
	return d.table
}

// Columns 获取字段
func (d *SysDepartDao) Columns() SysDepartColumns {
	return d.columns
}

// Ctx 获取上下文Model
func (d *SysDepartDao) Ctx(ctx context.Context) *gdb.Model {
	return d.DB().Model(d.table).Safe().Ctx(ctx)
}

// SysUserDepartDao 用户部门关联DAO
type SysUserDepartDao struct {
	table   string
	group   string
	columns SysUserDepartColumns
}

// SysUserDepartColumns 用户部门关联表字段
type SysUserDepartColumns struct {
	Id     string
	UserId string
	DepId  string
}

var sysUserDepartColumns = SysUserDepartColumns{
	Id:     "id",
	UserId: "user_id",
	DepId:  "dep_id",
}

// NewSysUserDepartDao 创建用户部门关联DAO
func NewSysUserDepartDao() *SysUserDepartDao {
	return &SysUserDepartDao{
		table:   "sys_user_depart",
		group:   "default",
		columns: sysUserDepartColumns,
	}
}

// DB 获取数据库连接
func (d *SysUserDepartDao) DB() gdb.DB {
	return g.DB(d.group)
}

// Table 获取表名
func (d *SysUserDepartDao) Table() string {
	return d.table
}

// Columns 获取字段
func (d *SysUserDepartDao) Columns() SysUserDepartColumns {
	return d.columns
}

// Ctx 获取上下文Model
func (d *SysUserDepartDao) Ctx(ctx context.Context) *gdb.Model {
	return d.DB().Model(d.table).Safe().Ctx(ctx)
}
