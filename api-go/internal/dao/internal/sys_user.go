package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SysUserDao 用户DAO
type SysUserDao struct {
	table   string
	group   string
	columns SysUserColumns
}

// SysUserColumns 用户表字段
type SysUserColumns struct {
	Id           string
	Username     string
	Realname     string
	Password     string
	Salt         string
	Avatar       string
	Birthday     string
	Sex          string
	Email        string
	Phone        string
	OrgCode      string
	Status       string
	DelFlag      string
	WorkNo       string
	Post         string
	School       string
	Telephone    string
	DepartIds    string
	ThirdId      string
	ThirdType    string
	UserIdentity string
	ActivitiSync string
	CreateBy     string
	CreateTime   string
	UpdateBy     string
	UpdateTime   string
}

var sysUserColumns = SysUserColumns{
	Id:           "id",
	Username:     "username",
	Realname:     "realname",
	Password:     "password",
	Salt:         "salt",
	Avatar:       "avatar",
	Birthday:     "birthday",
	Sex:          "sex",
	Email:        "email",
	Phone:        "phone",
	OrgCode:      "org_code",
	Status:       "status",
	DelFlag:      "del_flag",
	WorkNo:       "work_no",
	Post:         "post",
	School:       "school",
	Telephone:    "telephone",
	DepartIds:    "depart_ids",
	ThirdId:      "third_id",
	ThirdType:    "third_type",
	UserIdentity: "user_identity",
	ActivitiSync: "activiti_sync",
	CreateBy:     "create_by",
	CreateTime:   "create_time",
	UpdateBy:     "update_by",
	UpdateTime:   "update_time",
}

// NewSysUserDao 创建用户DAO
func NewSysUserDao() *SysUserDao {
	return &SysUserDao{
		table:   "sys_user",
		group:   "default",
		columns: sysUserColumns,
	}
}

// DB 获取数据库连接
func (d *SysUserDao) DB() gdb.DB {
	return g.DB(d.group)
}

// Table 获取表名
func (d *SysUserDao) Table() string {
	return d.table
}

// Columns 获取字段
func (d *SysUserDao) Columns() SysUserColumns {
	return d.columns
}

// Group 获取分组
func (d *SysUserDao) Group() string {
	return d.group
}

// Ctx 获取上下文Model
func (d *SysUserDao) Ctx(ctx context.Context) *gdb.Model {
	return d.DB().Model(d.table).Safe().Ctx(ctx)
}

// Transaction 事务
func (d *SysUserDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) error {
	return d.DB().Transaction(ctx, f)
}
