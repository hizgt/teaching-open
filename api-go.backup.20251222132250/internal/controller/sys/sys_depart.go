package sys

import (
	"context"

	"github.com/gogf/gf/v2/errors/gerror"

	v1 "teaching-open/api/v1/sys"
	"teaching-open/internal/service"
)

// cSysDepart 部门控制器
type cSysDepart struct{}

var SysDepart = &cSysDepart{}

// GetTree 获取部门树
func (c *cSysDepart) GetTree(ctx context.Context, req *v1.DepartTreeReq) (res *v1.DepartTreeRes, err error) {
	tree, err := service.SysDepart().GetTree(ctx)
	if err != nil {
		return nil, gerror.Wrap(err, "查询部门树失败")
	}

	return &v1.DepartTreeRes{
		List: tree,
	}, nil
}

// Add 新增部门
func (c *cSysDepart) Add(ctx context.Context, req *v1.DepartAddReq) (res *v1.DepartAddRes, err error) {
	err = service.SysDepart().Add(ctx, req)
	if err != nil {
		return nil, gerror.Wrap(err, "新增部门失败")
	}

	return &v1.DepartAddRes{}, nil
}

// Edit 编辑部门
func (c *cSysDepart) Edit(ctx context.Context, req *v1.DepartEditReq) (res *v1.DepartEditRes, err error) {
	err = service.SysDepart().Edit(ctx, req)
	if err != nil {
		return nil, gerror.Wrap(err, "编辑部门失败")
	}

	return &v1.DepartEditRes{}, nil
}

// Delete 删除部门
func (c *cSysDepart) Delete(ctx context.Context, req *v1.DepartDeleteReq) (res *v1.DepartDeleteRes, err error) {
	err = service.SysDepart().Delete(ctx, req.Id)
	if err != nil {
		return nil, gerror.Wrap(err, "删除部门失败")
	}

	return &v1.DepartDeleteRes{}, nil
}

// GetById 获取部门详情
func (c *cSysDepart) GetById(ctx context.Context, req *v1.DepartGetReq) (res *v1.DepartGetRes, err error) {
	info, err := service.SysDepart().GetById(ctx, req.Id)
	if err != nil {
		return nil, gerror.Wrap(err, "查询部门详情失败")
	}

	return &v1.DepartGetRes{
		DepartInfo: info,
	}, nil
}

// GetIdTree 获取部门ID树
func (c *cSysDepart) GetIdTree(ctx context.Context, req *v1.DepartIdTreeReq) (res *v1.DepartIdTreeRes, err error) {
	tree, err := service.SysDepart().GetIdTree(ctx)
	if err != nil {
		return nil, gerror.Wrap(err, "查询部门ID树失败")
	}

	return &v1.DepartIdTreeRes{
		List: tree,
	}, nil
}

// SearchBy 关键字搜索部门
func (c *cSysDepart) SearchBy(ctx context.Context, req *v1.DepartSearchReq) (res *v1.DepartSearchRes, err error) {
	list, err := service.SysDepart().SearchByKeyword(ctx, req.Keyword)
	if err != nil {
		return nil, gerror.Wrap(err, "搜索部门失败")
	}

	return &v1.DepartSearchRes{
		List: list,
	}, nil
}

// GetUserDeparts 获取用户部门
func (c *cSysDepart) GetUserDeparts(ctx context.Context, req *v1.UserDepartReq) (res *v1.UserDepartRes, err error) {
	depIds, departs, err := service.SysDepart().GetUserDeparts(ctx, req.UserId)
	if err != nil {
		return nil, gerror.Wrap(err, "查询用户部门失败")
	}

	return &v1.UserDepartRes{
		DepartIds: depIds,
		Departs:   departs,
	}, nil
}

// SaveUserDepart 保存用户部门
func (c *cSysDepart) SaveUserDepart(ctx context.Context, req *v1.SaveUserDepartReq) (res *v1.SaveUserDepartRes, err error) {
	err = service.SysDepart().SaveUserDeparts(ctx, req.UserId, req.DepartIds)
	if err != nil {
		return nil, gerror.Wrap(err, "保存用户部门失败")
	}

	return &v1.SaveUserDepartRes{}, nil
}
