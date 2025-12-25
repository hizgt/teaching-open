package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysPermission 权限表
type SysPermission struct {
	Id                 string      `json:"id"                 orm:"id,primary"            description:"权限ID"`
	ParentId           string      `json:"parentId"           orm:"parent_id"             description:"父ID"`
	Name               string      `json:"name"               orm:"name"                  description:"菜单名称"`
	Url                string      `json:"url"                orm:"url"                   description:"路径"`
	Component          string      `json:"component"          orm:"component"             description:"组件"`
	ComponentName      string      `json:"componentName"      orm:"component_name"        description:"组件名称"`
	Redirect           string      `json:"redirect"           orm:"redirect"              description:"重定向"`
	MenuType           int         `json:"menuType"           orm:"menu_type"             description:"菜单类型(0菜单1子菜单2按钮)"`
	Perms              string      `json:"perms"              orm:"perms"                 description:"权限标识"`
	PermsType          string      `json:"permsType"          orm:"perms_type"            description:"权限类型"`
	SortNo             float64     `json:"sortNo"             orm:"sort_no"               description:"排序"`
	AlwaysShow         bool        `json:"alwaysShow"         orm:"always_show"           description:"始终显示"`
	Icon               string      `json:"icon"               orm:"icon"                  description:"图标"`
	IsRoute            bool        `json:"isRoute"            orm:"is_route"              description:"是否路由"`
	IsLeaf             bool        `json:"isLeaf"             orm:"is_leaf"               description:"是否叶子节点"`
	KeepAlive          bool        `json:"keepAlive"          orm:"keep_alive"            description:"缓存页面"`
	Hidden             int         `json:"hidden"             orm:"hidden"                description:"隐藏"`
	Description        string      `json:"description"        orm:"description"           description:"描述"`
	DelFlag            int         `json:"delFlag"            orm:"del_flag"              description:"删除标记"`
	RuleFlag           int         `json:"ruleFlag"           orm:"rule_flag"             description:"规则标记"`
	Status             string      `json:"status"             orm:"status"                description:"状态"`
	InternalOrExternal bool        `json:"internalOrExternal" orm:"internal_or_external"  description:"内部或外部"`
	CreateBy           string      `json:"createBy"           orm:"create_by"             description:"创建人"`
	CreateTime         *gtime.Time `json:"createTime"         orm:"create_time"           description:"创建时间"`
	UpdateBy           string      `json:"updateBy"           orm:"update_by"             description:"更新人"`
	UpdateTime         *gtime.Time `json:"updateTime"         orm:"update_time"           description:"更新时间"`
}

// TableName 表名
func (e *SysPermission) TableName() string {
	return "sys_permission"
}

// SysRolePermission 角色权限关联表
type SysRolePermission struct {
	Id           string      `json:"id"           orm:"id,primary"      description:"ID"`
	RoleId       string      `json:"roleId"       orm:"role_id"         description:"角色ID"`
	PermissionId string      `json:"permissionId" orm:"permission_id"   description:"权限ID"`
	DataRuleIds  string      `json:"dataRuleIds"  orm:"data_rule_ids"   description:"数据规则ID"`
	OperateDate  *gtime.Time `json:"operateDate"  orm:"operate_date"    description:"操作时间"`
	OperateIp    string      `json:"operateIp"    orm:"operate_ip"      description:"操作IP"`
}

// TableName 表名
func (e *SysRolePermission) TableName() string {
	return "sys_role_permission"
}
