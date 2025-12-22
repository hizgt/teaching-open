// =================================================================================
// Code generated and target of the generated file. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// SysUserRole is the golang structure of table sys_user_role for DAO operations like Where/Data.
type SysUserRole struct {
	g.Meta `orm:"table:sys_user_role, do:true"`
	Id     interface{} // 主键id
	UserId interface{} // 用户id
	RoleId interface{} // 角色id
}
