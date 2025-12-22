// =================================================================================
// Code generated and target of the generated file. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysRole is the golang structure of table sys_role for DAO operations like Where/Data.
type SysRole struct {
	g.Meta      `orm:"table:sys_role, do:true"`
	Id          interface{} // 主键id
	RoleName    interface{} // 角色名称
	RoleCode    interface{} // 角色编码
	Description interface{} // 描述
	CreateBy    interface{} // 创建人
	CreateTime  *gtime.Time // 创建时间
	UpdateBy    interface{} // 更新人
	UpdateTime  *gtime.Time // 更新时间
	RoleLevel   interface{} // 角色级别
}
