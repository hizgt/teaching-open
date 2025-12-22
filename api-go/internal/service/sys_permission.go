// =================================================================================
// Service interface for sys permission module
// =================================================================================

package service

import (
	"context"

	v1 "teaching-open/api/v1/sys"
)

// ISysPermission 权限服务接口
type ISysPermission interface {
	// GetList 获取权限列表（平铺）
	GetList(ctx context.Context) ([]*v1.PermissionInfo, error)

	// GetTree 获取权限树
	GetTree(ctx context.Context) ([]*v1.PermissionInfo, error)

	// Add 添加权限
	Add(ctx context.Context, req *v1.PermissionAddReq) error

	// Edit 编辑权限
	Edit(ctx context.Context, req *v1.PermissionEditReq) error

	// Delete 删除权限
	Delete(ctx context.Context, id string) error

	// GetById 根据ID获取权限详情
	GetById(ctx context.Context, id string) (*v1.PermissionInfo, error)

	// GetRolePermissions 获取角色的权限ID列表
	GetRolePermissions(ctx context.Context, roleId string) ([]string, error)

	// SaveRolePermissions 保存角色权限
	SaveRolePermissions(ctx context.Context, roleId string, permissionIds []string) error

	// GetUserPermissions 获取用户权限（菜单和权限标识）
	GetUserPermissions(ctx context.Context, userId string) ([]*v1.PermissionInfo, []string, error)
}

var localSysPermission ISysPermission

// SysPermission 获取权限服务实例
func SysPermission() ISysPermission {
	if localSysPermission == nil {
		panic("implement not found for interface ISysPermission, forgot register?")
	}
	return localSysPermission
}

// RegisterSysPermission 注册权限服务实现
func RegisterSysPermission(i ISysPermission) {
	localSysPermission = i
}
