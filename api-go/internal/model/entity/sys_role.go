// =================================================================================
// Code generated and target of the generated file. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysRole is the golang structure for table sys_role.
type SysRole struct {
	Id          string      `json:"id"          description:"主键id"`
	RoleName    string      `json:"roleName"    description:"角色名称"`
	RoleCode    string      `json:"roleCode"    description:"角色编码"`
	Description string      `json:"description" description:"描述"`
	CreateBy    string      `json:"createBy"    description:"创建人"`
	CreateTime  *gtime.Time `json:"createTime"  description:"创建时间"`
	UpdateBy    string      `json:"updateBy"    description:"更新人"`
	UpdateTime  *gtime.Time `json:"updateTime"  description:"更新时间"`
	RoleLevel   int         `json:"roleLevel"   description:"角色级别"`
}
