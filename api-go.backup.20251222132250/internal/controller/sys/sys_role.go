package sys

import (
	"context"

	"github.com/gogf/gf/v2/errors/gerror"

	v1 "teaching-open/api/v1/sys"
	"teaching-open/internal/service"
)

// cSysRole 角色控制器
type cSysRole struct{}

var SysRole = &cSysRole{}

// GetList 获取角色列表
func (c *cSysRole) GetList(ctx context.Context, req *v1.RoleListReq) (res *v1.RoleListRes, err error) {
	// 调用业务逻辑层查询角色列表
	list, total, err := service.SysRole.GetList(ctx, req)
	if err != nil {
		return nil, gerror.Wrap(err, "查询角色列表失败")
	}

	// 返回结果
	res = &v1.RoleListRes{
		List:     list,
		Total:    total,
		Page:     req.Page,
		PageSize: req.PageSize,
	}

	return res, nil
}

// Add 新增角色
func (c *cSysRole) Add(ctx context.Context, req *v1.RoleAddReq) (res *v1.RoleInfo, err error) {
	// 调用业务逻辑层新增角色
	err = service.SysRole.Add(ctx, req)
	if err != nil {
		return nil, gerror.Wrap(err, "新增角色失败")
	}

	// 返回空,表示成功
	return nil, nil
}

// Edit 编辑角色
func (c *cSysRole) Edit(ctx context.Context, req *v1.RoleEditReq) (res *v1.RoleInfo, err error) {
	// 调用业务逻辑层编辑角色
	err = service.SysRole.Edit(ctx, req)
	if err != nil {
		return nil, gerror.Wrap(err, "编辑角色失败")
	}

	// 返回空,表示成功
	return nil, nil
}

// Delete 删除角色
func (c *cSysRole) Delete(ctx context.Context, req *v1.RoleDeleteReq) (res interface{}, err error) {
	// 调用业务逻辑层删除角色
	err = service.SysRole.Delete(ctx, req.Id)
	if err != nil {
		return nil, gerror.Wrap(err, "删除角色失败")
	}

	// 返回空,表示成功
	return nil, nil
}

// GetById 获取角色详情
func (c *cSysRole) GetById(ctx context.Context, req *v1.RoleGetReq) (res *v1.RoleInfo, err error) {
	// 调用业务逻辑层查询角色详情
	roleInfo, err := service.SysRole.GetById(ctx, req.Id)
	if err != nil {
		return nil, gerror.Wrap(err, "查询角色详情失败")
	}

	return roleInfo, nil
}

// GetAll 获取所有角色
func (c *cSysRole) GetAll(ctx context.Context, req *v1.RoleAllReq) (res *v1.RoleAllRes, err error) {
	// 调用业务逻辑层查询所有角色
	list, err := service.SysRole.GetAll(ctx)
	if err != nil {
		return nil, gerror.Wrap(err, "查询所有角色失败")
	}

	return &v1.RoleAllRes{
		List: list,
	}, nil
}

// GetUserRoles 获取用户角色
func (c *cSysRole) GetUserRoles(ctx context.Context, req *v1.UserRoleReq) (res *v1.UserRoleRes, err error) {
	// 调用业务逻辑层查询用户角色
	roleIds, roles, err := service.SysRole.GetUserRoles(ctx, req.UserId)
	if err != nil {
		return nil, gerror.Wrap(err, "查询用户角色失败")
	}

	return &v1.UserRoleRes{
		RoleIds: roleIds,
		Roles:   roles,
	}, nil
}

// SaveUserRoles 保存用户角色
func (c *cSysRole) SaveUserRoles(ctx context.Context, req *v1.UserRoleSaveReq) (res interface{}, err error) {
	// 调用业务逻辑层保存用户角色
	err = service.SysRole.SaveUserRoles(ctx, req.UserId, req.RoleIds)
	if err != nil {
		return nil, gerror.Wrap(err, "保存用户角色失败")
	}

	// 返回空,表示成功
	return nil, nil
}
