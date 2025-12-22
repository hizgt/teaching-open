// =================================================================================
// API definitions for sys permission module
// =================================================================================

package sys

import (
	"github.com/gogf/gf/v2/frame/g"
)

// PermissionInfo 权限信息
type PermissionInfo struct {
	Id                 string            `json:"id"`
	ParentId           string            `json:"parentId"`
	Name               string            `json:"name"`
	Url                string            `json:"url"`
	Component          string            `json:"component"`
	ComponentName      string            `json:"componentName"`
	Redirect           string            `json:"redirect"`
	MenuType           int               `json:"menuType"`
	Perms              string            `json:"perms"`
	PermsType          string            `json:"permsType"`
	SortNo             float64           `json:"sortNo"`
	AlwaysShow         int               `json:"alwaysShow"`
	Icon               string            `json:"icon"`
	IsRoute            int               `json:"isRoute"`
	IsLeaf             int               `json:"isLeaf"`
	KeepAlive          int               `json:"keepAlive"`
	Hidden             int               `json:"hidden"`
	Description        string            `json:"description"`
	Status             string            `json:"status"`
	InternalOrExternal int               `json:"internalOrExternal"`
	CreateTime         string            `json:"createTime"`
	Children           []*PermissionInfo `json:"children,omitempty"`
}

// PermissionListReq 权限列表请求
type PermissionListReq struct {
	g.Meta `path:"/sys/permission/list" method:"get" tags:"权限管理" summary:"获取权限列表"`
}

// PermissionListRes 权限列表响应
type PermissionListRes struct {
	List []*PermissionInfo `json:"list"`
}

// PermissionTreeReq 权限树请求
type PermissionTreeReq struct {
	g.Meta `path:"/sys/permission/queryTreeList" method:"get" tags:"权限管理" summary:"获取权限树"`
}

// PermissionTreeRes 权限树响应
type PermissionTreeRes struct {
	List []*PermissionInfo `json:"list"`
}

// PermissionAddReq 添加权限请求
type PermissionAddReq struct {
	g.Meta             `path:"/sys/permission/add" method:"post" tags:"权限管理" summary:"添加权限"`
	ParentId           string  `json:"parentId"`
	Name               string  `json:"name" v:"required#菜单名称不能为空"`
	Url                string  `json:"url"`
	Component          string  `json:"component"`
	ComponentName      string  `json:"componentName"`
	Redirect           string  `json:"redirect"`
	MenuType           int     `json:"menuType" v:"required|in:0,1,2#菜单类型不能为空|菜单类型值无效"`
	Perms              string  `json:"perms"`
	PermsType          string  `json:"permsType"`
	SortNo             float64 `json:"sortNo"`
	AlwaysShow         int     `json:"alwaysShow"`
	Icon               string  `json:"icon"`
	IsRoute            int     `json:"isRoute"`
	IsLeaf             int     `json:"isLeaf"`
	KeepAlive          int     `json:"keepAlive"`
	Hidden             int     `json:"hidden"`
	Description        string  `json:"description"`
	Status             string  `json:"status"`
	InternalOrExternal int     `json:"internalOrExternal"`
}

// PermissionAddRes 添加权限响应
type PermissionAddRes struct{}

// PermissionEditReq 编辑权限请求
type PermissionEditReq struct {
	g.Meta             `path:"/sys/permission/edit" method:"put" tags:"权限管理" summary:"编辑权限"`
	Id                 string  `json:"id" v:"required#权限ID不能为空"`
	ParentId           string  `json:"parentId"`
	Name               string  `json:"name" v:"required#菜单名称不能为空"`
	Url                string  `json:"url"`
	Component          string  `json:"component"`
	ComponentName      string  `json:"componentName"`
	Redirect           string  `json:"redirect"`
	MenuType           int     `json:"menuType" v:"required|in:0,1,2#菜单类型不能为空|菜单类型值无效"`
	Perms              string  `json:"perms"`
	PermsType          string  `json:"permsType"`
	SortNo             float64 `json:"sortNo"`
	AlwaysShow         int     `json:"alwaysShow"`
	Icon               string  `json:"icon"`
	IsRoute            int     `json:"isRoute"`
	IsLeaf             int     `json:"isLeaf"`
	KeepAlive          int     `json:"keepAlive"`
	Hidden             int     `json:"hidden"`
	Description        string  `json:"description"`
	Status             string  `json:"status"`
	InternalOrExternal int     `json:"internalOrExternal"`
}

// PermissionEditRes 编辑权限响应
type PermissionEditRes struct{}

// PermissionDeleteReq 删除权限请求
type PermissionDeleteReq struct {
	g.Meta `path:"/sys/permission/delete" method:"delete" tags:"权限管理" summary:"删除权限"`
	Id     string `json:"id" v:"required#权限ID不能为空"`
}

// PermissionDeleteRes 删除权限响应
type PermissionDeleteRes struct{}

// PermissionGetReq 获取权限详情请求
type PermissionGetReq struct {
	g.Meta `path:"/sys/permission/getById" method:"get" tags:"权限管理" summary:"获取权限详情"`
	Id     string `json:"id" v:"required#权限ID不能为空"`
}

// PermissionGetRes 获取权限详情响应
type PermissionGetRes struct {
	*PermissionInfo
}

// RolePermissionReq 获取角色权限请求
type RolePermissionReq struct {
	g.Meta `path:"/sys/permission/queryRolePermission" method:"get" tags:"权限管理" summary:"获取角色权限"`
	RoleId string `json:"roleId" v:"required#角色ID不能为空"`
}

// RolePermissionRes 获取角色权限响应
type RolePermissionRes struct {
	PermissionIds []string `json:"permissionIds"`
}

// SaveRolePermissionReq 保存角色权限请求
type SaveRolePermissionReq struct {
	g.Meta        `path:"/sys/permission/saveRolePermission" method:"post" tags:"权限管理" summary:"保存角色权限"`
	RoleId        string   `json:"roleId" v:"required#角色ID不能为空"`
	PermissionIds []string `json:"permissionIds"`
}

// SaveRolePermissionRes 保存角色权限响应
type SaveRolePermissionRes struct{}

// UserPermissionReq 获取用户权限请求
type UserPermissionReq struct {
	g.Meta `path:"/sys/permission/getUserPermissionByToken" method:"get" tags:"权限管理" summary:"获取当前用户权限"`
}

// UserPermissionRes 获取用户权限响应
type UserPermissionRes struct {
	Menu []*PermissionInfo `json:"menu"`
	Auth []string          `json:"auth"`
}
