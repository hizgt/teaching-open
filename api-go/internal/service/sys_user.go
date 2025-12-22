package service

import (
	"context"

	v1 "teaching-open/api/v1/sys"
)

// ISysUser 用户服务接口
type ISysUser interface {
	// Login 用户登录
	Login(ctx context.Context, username, password string) (token string, userInfo *v1.UserInfo, err error)
	// GetList 获取用户列表
	GetList(ctx context.Context, req *v1.UserListReq) (list []*v1.UserInfo, total int64, err error)
	// Add 新增用户
	Add(ctx context.Context, req *v1.UserAddReq) error
	// Edit 编辑用户
	Edit(ctx context.Context, req *v1.UserEditReq) error
	// Delete 删除用户
	Delete(ctx context.Context, id string) error
	// GetById 根据ID获取用户
	GetById(ctx context.Context, id string) (*v1.UserInfo, error)
}

var SysUser ISysUser
