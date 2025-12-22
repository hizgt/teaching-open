package sys

import (
	"context"

	"github.com/gogf/gf/v2/errors/gerror"

	v1 "teaching-open/api/v1/sys"
	"teaching-open/internal/service"
)

// cSysUser 用户控制器
type cSysUser struct{}

var SysUser = &cSysUser{}

// Login 用户登录
func (c *cSysUser) Login(ctx context.Context, req *v1.LoginReq) (res *v1.LoginRes, err error) {
	// 调用业务逻辑层处理登录
	token, userInfo, err := service.SysUser.Login(ctx, req.Username, req.Password)
	if err != nil {
		return nil, gerror.Wrap(err, "登录失败")
	}

	// 返回登录结果
	res = &v1.LoginRes{
		Token:    token,
		UserInfo: userInfo,
	}

	return res, nil
}

// GetList 获取用户列表
func (c *cSysUser) GetList(ctx context.Context, req *v1.UserListReq) (res *v1.UserListRes, err error) {
	// 调用业务逻辑层查询用户列表
	list, total, err := service.SysUser.GetList(ctx, req)
	if err != nil {
		return nil, gerror.Wrap(err, "查询用户列表失败")
	}

	// 返回结果
	res = &v1.UserListRes{
		List:     list,
		Total:    total,
		Page:     req.Page,
		PageSize: req.PageSize,
	}

	return res, nil
}

// Add 新增用户
func (c *cSysUser) Add(ctx context.Context, req *v1.UserAddReq) (res *v1.UserInfo, err error) {
	// 调用业务逻辑层新增用户
	err = service.SysUser.Add(ctx, req)
	if err != nil {
		return nil, gerror.Wrap(err, "新增用户失败")
	}

	// 返回空,表示成功
	return nil, nil
}

// Edit 编辑用户
func (c *cSysUser) Edit(ctx context.Context, req *v1.UserEditReq) (res *v1.UserInfo, err error) {
	// 调用业务逻辑层编辑用户
	err = service.SysUser.Edit(ctx, req)
	if err != nil {
		return nil, gerror.Wrap(err, "编辑用户失败")
	}

	// 返回空,表示成功
	return nil, nil
}

// Delete 删除用户
func (c *cSysUser) Delete(ctx context.Context, req *v1.UserDeleteReq) (res interface{}, err error) {
	// 调用业务逻辑层删除用户
	err = service.SysUser.Delete(ctx, req.Id)
	if err != nil {
		return nil, gerror.Wrap(err, "删除用户失败")
	}

	// 返回空,表示成功
	return nil, nil
}

// GetById 获取用户详情
func (c *cSysUser) GetById(ctx context.Context, req *v1.UserGetReq) (res *v1.UserInfo, err error) {
	// 调用业务逻辑层查询用户详情
	userInfo, err := service.SysUser.GetById(ctx, req.Id)
	if err != nil {
		return nil, gerror.Wrap(err, "查询用户详情失败")
	}

	return userInfo, nil
}
