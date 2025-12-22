package sys

import (
	"github.com/gogf/gf/v2/frame/g"
)

// RoleInfo 角色信息
type RoleInfo struct {
	Id          string `json:"id" dc:"角色ID"`
	RoleName    string `json:"roleName" dc:"角色名称"`
	RoleCode    string `json:"roleCode" dc:"角色编码"`
	Description string `json:"description" dc:"描述"`
	RoleLevel   int    `json:"roleLevel" dc:"角色级别"`
	CreateTime  string `json:"createTime" dc:"创建时间"`
}

// RoleListReq 角色列表请求
type RoleListReq struct {
	g.Meta   `path:"/sys/role/list" method:"get" tags:"角色管理" summary:"角色列表"`
	Page     int    `json:"page" d:"1" v:"min:1#页码最小为1" dc:"页码"`
	PageSize int    `json:"pageSize" d:"10" v:"min:1|max:100#每页条数最小为1|每页条数最大为100" dc:"每页条数"`
	RoleName string `json:"roleName" dc:"角色名称(模糊查询)"`
	RoleCode string `json:"roleCode" dc:"角色编码(模糊查询)"`
}

// RoleListRes 角色列表响应
type RoleListRes struct {
	List     []*RoleInfo `json:"list" dc:"角色列表"`
	Total    int64       `json:"total" dc:"总数"`
	Page     int         `json:"page" dc:"当前页"`
	PageSize int         `json:"pageSize" dc:"每页条数"`
}

// RoleAddReq 新增角色请求
type RoleAddReq struct {
	g.Meta      `path:"/sys/role" method:"post" tags:"角色管理" summary:"新增角色"`
	RoleName    string `json:"roleName" v:"required#角色名称不能为空" dc:"角色名称"`
	RoleCode    string `json:"roleCode" v:"required#角色编码不能为空" dc:"角色编码"`
	Description string `json:"description" dc:"描述"`
	RoleLevel   int    `json:"roleLevel" d:"0" dc:"角色级别"`
}

// RoleEditReq 编辑角色请求
type RoleEditReq struct {
	g.Meta      `path:"/sys/role" method:"put" tags:"角色管理" summary:"编辑角色"`
	Id          string `json:"id" v:"required#角色ID不能为空" dc:"角色ID"`
	RoleName    string `json:"roleName" dc:"角色名称"`
	RoleCode    string `json:"roleCode" dc:"角色编码"`
	Description string `json:"description" dc:"描述"`
	RoleLevel   int    `json:"roleLevel" dc:"角色级别"`
}

// RoleDeleteReq 删除角色请求
type RoleDeleteReq struct {
	g.Meta `path:"/sys/role/:id" method:"delete" tags:"角色管理" summary:"删除角色"`
	Id     string `json:"id" v:"required#角色ID不能为空" dc:"角色ID"`
}

// RoleGetReq 获取角色详情请求
type RoleGetReq struct {
	g.Meta `path:"/sys/role/:id" method:"get" tags:"角色管理" summary:"角色详情"`
	Id     string `json:"id" v:"required#角色ID不能为空" dc:"角色ID"`
}

// RoleAllReq 获取所有角色请求（用于下拉选择）
type RoleAllReq struct {
	g.Meta `path:"/sys/role/all" method:"get" tags:"角色管理" summary:"所有角色"`
}

// RoleAllRes 所有角色响应
type RoleAllRes struct {
	List []*RoleInfo `json:"list" dc:"角色列表"`
}

// UserRoleReq 获取用户角色请求
type UserRoleReq struct {
	g.Meta `path:"/sys/user/role/:userId" method:"get" tags:"角色管理" summary:"获取用户角色"`
	UserId string `json:"userId" v:"required#用户ID不能为空" dc:"用户ID"`
}

// UserRoleRes 用户角色响应
type UserRoleRes struct {
	RoleIds []string    `json:"roleIds" dc:"角色ID列表"`
	Roles   []*RoleInfo `json:"roles" dc:"角色列表"`
}

// UserRoleSaveReq 保存用户角色请求
type UserRoleSaveReq struct {
	g.Meta  `path:"/sys/user/role" method:"post" tags:"角色管理" summary:"保存用户角色"`
	UserId  string   `json:"userId" v:"required#用户ID不能为空" dc:"用户ID"`
	RoleIds []string `json:"roleIds" dc:"角色ID列表"`
}
