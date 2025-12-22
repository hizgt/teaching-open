package service

import (
	"context"

	v1 "teaching-open/api/v1/sys"
)

// ISysRole 角色服务接口
type ISysRole interface {
	// GetList 获取角色列表
	GetList(ctx context.Context, req *v1.RoleListReq) (list []*v1.RoleInfo, total int64, err error)
	// Add 新增角色
	Add(ctx context.Context, req *v1.RoleAddReq) error
	// Edit 编辑角色
	Edit(ctx context.Context, req *v1.RoleEditReq) error
	// Delete 删除角色
	Delete(ctx context.Context, id string) error
	// GetById 根据ID获取角色
	GetById(ctx context.Context, id string) (*v1.RoleInfo, error)
	// GetAll 获取所有角色
	GetAll(ctx context.Context) ([]*v1.RoleInfo, error)
	// GetUserRoles 获取用户角色
	GetUserRoles(ctx context.Context, userId string) ([]string, []*v1.RoleInfo, error)
	// SaveUserRoles 保存用户角色
	SaveUserRoles(ctx context.Context, userId string, roleIds []string) error
}

var SysRole ISysRole
