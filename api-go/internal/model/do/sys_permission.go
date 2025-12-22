// =================================================================================
// Code generated and target of the generated file. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysPermission is the golang structure of table sys_permission for DAO operations like Where/Data.
type SysPermission struct {
	g.Meta             `orm:"table:sys_permission, do:true"`
	Id                 interface{} // 主键id
	ParentId           interface{} // 父id
	Name               interface{} // 菜单标题
	Url                interface{} // 路径
	Component          interface{} // 组件
	ComponentName      interface{} // 组件名字
	Redirect           interface{} // 一级菜单跳转地址
	MenuType           interface{} // 菜单类型(0:一级菜单; 1:子菜单:2:按钮权限)
	Perms              interface{} // 菜单权限编码
	PermsType          interface{} // 权限策略1显示2禁用
	SortNo             interface{} // 菜单排序
	AlwaysShow         interface{} // 聚合子路由: 1是0否
	Icon               interface{} // 菜单图标
	IsRoute            interface{} // 是否路由菜单: 0:不是  1:是（默认值1）
	IsLeaf             interface{} // 是否叶子节点:    1:是   0:不是
	KeepAlive          interface{} // 是否缓存该页面:    1:是   0:不是
	Hidden             interface{} // 是否隐藏路由: 0否,1是
	Description        interface{} // 描述
	CreateBy           interface{} // 创建人
	CreateTime         *gtime.Time // 创建时间
	UpdateBy           interface{} // 更新人
	UpdateTime         *gtime.Time // 更新时间
	DelFlag            interface{} // 删除状态 0正常 1已删除
	RuleFlag           interface{} // 是否添加数据权限1是0否
	Status             interface{} // 按钮权限状态(0无效1有效)
	InternalOrExternal interface{} // 外链菜单打开方式 0/内部打开 1/外部打开
}
