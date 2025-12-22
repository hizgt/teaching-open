// =================================================================================
// Code generated and target of the generated file. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// SysRolePermission is the golang structure of table sys_role_permission for DAO operations like Where/Data.
type SysRolePermission struct {
	g.Meta       `orm:"table:sys_role_permission, do:true"`
	Id           interface{} // 主键id
	RoleId       interface{} // 角色id
	PermissionId interface{} // 权限id
	DataRuleIds  interface{} // 数据规则ID
}
