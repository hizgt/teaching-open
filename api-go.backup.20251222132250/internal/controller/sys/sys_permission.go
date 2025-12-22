package sys

import (
	"context"

	"github.com/gogf/gf/v2/errors/gerror"

	v1 "teaching-open/api/v1/sys"
	"teaching-open/internal/service"
)

// cSysPermission 权限控制器
type cSysPermission struct{}

var SysPermission = &cSysPermission{}

// GetList 获取权限列表
func (c *cSysPermission) GetList(ctx context.Context, req *v1.PermissionListReq) (res *v1.PermissionListRes, err error) {
	// 调用业务逻辑层查询权限列表
	list, err := service.SysPermission().GetList(ctx)
	if err != nil {
		return nil, gerror.Wrap(err, "查询权限列表失败")
	}

	return &v1.PermissionListRes{
		List: list,
	}, nil
}

// GetTree 获取权限树
func (c *cSysPermission) GetTree(ctx context.Context, req *v1.PermissionTreeReq) (res *v1.PermissionTreeRes, err error) {
	// 调用业务逻辑层查询权限树
	tree, err := service.SysPermission().GetTree(ctx)
	if err != nil {
		return nil, gerror.Wrap(err, "查询权限树失败")
	}

	return &v1.PermissionTreeRes{
		List: tree,
	}, nil
}

// Add 新增权限
func (c *cSysPermission) Add(ctx context.Context, req *v1.PermissionAddReq) (res *v1.PermissionInfo, err error) {
	// 调用业务逻辑层新增权限
	err = service.SysPermission().Add(ctx, req)
	if err != nil {
		return nil, gerror.Wrap(err, "新增权限失败")
	}

	return nil, nil
}

// Edit 编辑权限
func (c *cSysPermission) Edit(ctx context.Context, req *v1.PermissionEditReq) (res *v1.PermissionInfo, err error) {
	// 调用业务逻辑层编辑权限
	err = service.SysPermission().Edit(ctx, req)
	if err != nil {
		return nil, gerror.Wrap(err, "编辑权限失败")
	}

	return nil, nil
}

// Delete 删除权限
func (c *cSysPermission) Delete(ctx context.Context, req *v1.PermissionDeleteReq) (res interface{}, err error) {
	// 调用业务逻辑层删除权限
	err = service.SysPermission().Delete(ctx, req.Id)
	if err != nil {
		return nil, gerror.Wrap(err, "删除权限失败")
	}

	return nil, nil
}

// GetById 获取权限详情
func (c *cSysPermission) GetById(ctx context.Context, req *v1.PermissionGetReq) (res *v1.PermissionInfo, err error) {
	// 调用业务逻辑层查询权限详情
	permInfo, err := service.SysPermission().GetById(ctx, req.Id)
	if err != nil {
		return nil, gerror.Wrap(err, "查询权限详情失败")
	}

	return permInfo, nil
}

// GetRolePermissions 获取角色权限
func (c *cSysPermission) GetRolePermissions(ctx context.Context, req *v1.RolePermissionReq) (res *v1.RolePermissionRes, err error) {
	// 调用业务逻辑层查询角色权限
	permIds, err := service.SysPermission().GetRolePermissions(ctx, req.RoleId)
	if err != nil {
		return nil, gerror.Wrap(err, "查询角色权限失败")
	}

	return &v1.RolePermissionRes{
		PermissionIds: permIds,
	}, nil
}

// SaveRolePermissions 保存角色权限
func (c *cSysPermission) SaveRolePermissions(ctx context.Context, req *v1.SaveRolePermissionReq) (res interface{}, err error) {
	// 调用业务逻辑层保存角色权限
	err = service.SysPermission().SaveRolePermissions(ctx, req.RoleId, req.PermissionIds)
	if err != nil {
		return nil, gerror.Wrap(err, "保存角色权限失败")
	}

	return nil, nil
}

// GetUserPermissions 获取用户权限（Token方式）
func (c *cSysPermission) GetUserPermissions(ctx context.Context, req *v1.UserPermissionReq) (res *v1.UserPermissionRes, err error) {
	// 从context中获取当前用户ID
	userId := ctx.Value("userId")
	if userId == nil {
		return nil, gerror.New("用户未登录")
	}

	// 调用业务逻辑层查询用户权限
	menus, authList, err := service.SysPermission().GetUserPermissions(ctx, userId.(string))
	if err != nil {
		return nil, gerror.Wrap(err, "查询用户权限失败")
	}

	return &v1.UserPermissionRes{
		Menu: menus,
		Auth: authList,
	}, nil
}
