package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysRole 角色表
type SysRole struct {
	Id          string      `json:"id"          orm:"id,primary"    description:"角色ID"`
	RoleName    string      `json:"roleName"    orm:"role_name"     description:"角色名称"`
	RoleCode    string      `json:"roleCode"    orm:"role_code"     description:"角色编码"`
	RoleLevel   int         `json:"roleLevel"   orm:"role_level"    description:"角色级别"`
	Description string      `json:"description" orm:"description"   description:"描述"`
	CreateBy    string      `json:"createBy"    orm:"create_by"     description:"创建人"`
	CreateTime  *gtime.Time `json:"createTime"  orm:"create_time"   description:"创建时间"`
	UpdateBy    string      `json:"updateBy"    orm:"update_by"     description:"更新人"`
	UpdateTime  *gtime.Time `json:"updateTime"  orm:"update_time"   description:"更新时间"`
}

// TableName 表名
func (e *SysRole) TableName() string {
	return "sys_role"
}

// SysUserRole 用户角色关联表
type SysUserRole struct {
	Id     string `json:"id"     orm:"id,primary" description:"ID"`
	UserId string `json:"userId" orm:"user_id"    description:"用户ID"`
	RoleId string `json:"roleId" orm:"role_id"    description:"角色ID"`
}

// TableName 表名
func (e *SysUserRole) TableName() string {
	return "sys_user_role"
}
