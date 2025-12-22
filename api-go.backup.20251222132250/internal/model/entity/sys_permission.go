// =================================================================================
// Code generated and target of the generated file. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysPermission is the golang structure for table sys_permission.
type SysPermission struct {
	Id                 string      `json:"id"                 description:"主键id"`
	ParentId           string      `json:"parentId"           description:"父id"`
	Name               string      `json:"name"               description:"菜单标题"`
	Url                string      `json:"url"                description:"路径"`
	Component          string      `json:"component"          description:"组件"`
	ComponentName      string      `json:"componentName"      description:"组件名字"`
	Redirect           string      `json:"redirect"           description:"一级菜单跳转地址"`
	MenuType           int         `json:"menuType"           description:"菜单类型(0:一级菜单; 1:子菜单:2:按钮权限)"`
	Perms              string      `json:"perms"              description:"菜单权限编码"`
	PermsType          string      `json:"permsType"          description:"权限策略1显示2禁用"`
	SortNo             float64     `json:"sortNo"             description:"菜单排序"`
	AlwaysShow         int         `json:"alwaysShow"         description:"聚合子路由: 1是0否"`
	Icon               string      `json:"icon"               description:"菜单图标"`
	IsRoute            int         `json:"isRoute"            description:"是否路由菜单: 0:不是  1:是（默认值1）"`
	IsLeaf             int         `json:"isLeaf"             description:"是否叶子节点:    1:是   0:不是"`
	KeepAlive          int         `json:"keepAlive"          description:"是否缓存该页面:    1:是   0:不是"`
	Hidden             int         `json:"hidden"             description:"是否隐藏路由: 0否,1是"`
	Description        string      `json:"description"        description:"描述"`
	CreateBy           string      `json:"createBy"           description:"创建人"`
	CreateTime         *gtime.Time `json:"createTime"         description:"创建时间"`
	UpdateBy           string      `json:"updateBy"           description:"更新人"`
	UpdateTime         *gtime.Time `json:"updateTime"         description:"更新时间"`
	DelFlag            int         `json:"delFlag"            description:"删除状态 0正常 1已删除"`
	RuleFlag           int         `json:"ruleFlag"           description:"是否添加数据权限1是0否"`
	Status             string      `json:"status"             description:"按钮权限状态(0无效1有效)"`
	InternalOrExternal int         `json:"internalOrExternal" description:"外链菜单打开方式 0/内部打开 1/外部打开"`
}
