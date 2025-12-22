// =================================================================================
// Controller for sys dict module
// =================================================================================

package sys

import (
	"context"

	"github.com/gogf/gf/v2/errors/gerror"

	v1 "teaching-open/api/v1/sys"
	"teaching-open/internal/service"
)

type cDict struct{}

var Dict = &cDict{}

// List 获取字典列表
func (c *cDict) List(ctx context.Context, req *v1.DictListReq) (res *v1.DictListRes, err error) {
	// 设置默认值
	if req.Page <= 0 {
		req.Page = 1
	}
	if req.PageSize <= 0 {
		req.PageSize = 10
	}

	list, total, err := service.SysDict().GetList(ctx, req)
	if err != nil {
		return nil, gerror.Wrap(err, "查询字典列表失败")
	}

	return &v1.DictListRes{
		List:     list,
		Total:    total,
		Page:     req.Page,
		PageSize: req.PageSize,
	}, nil
}

// Add 添加字典
func (c *cDict) Add(ctx context.Context, req *v1.DictAddReq) (res *v1.DictAddRes, err error) {
	err = service.SysDict().Add(ctx, req)
	if err != nil {
		return nil, gerror.Wrap(err, "添加字典失败")
	}

	return &v1.DictAddRes{}, nil
}

// Edit 编辑字典
func (c *cDict) Edit(ctx context.Context, req *v1.DictEditReq) (res *v1.DictEditRes, err error) {
	err = service.SysDict().Edit(ctx, req)
	if err != nil {
		return nil, gerror.Wrap(err, "编辑字典失败")
	}

	return &v1.DictEditRes{}, nil
}

// Delete 删除字典
func (c *cDict) Delete(ctx context.Context, req *v1.DictDeleteReq) (res *v1.DictDeleteRes, err error) {
	err = service.SysDict().Delete(ctx, req.Id)
	if err != nil {
		return nil, gerror.Wrap(err, "删除字典失败")
	}

	return &v1.DictDeleteRes{}, nil
}

// GetById 根据ID获取字典详情
func (c *cDict) GetById(ctx context.Context, req *v1.DictGetByIdReq) (res *v1.DictGetByIdRes, err error) {
	info, err := service.SysDict().GetById(ctx, req.Id)
	if err != nil {
		return nil, gerror.Wrap(err, "查询字典详情失败")
	}

	return &v1.DictGetByIdRes{
		DictInfo: info,
	}, nil
}

// ItemList 根据字典ID获取字典项列表
func (c *cDict) ItemList(ctx context.Context, req *v1.DictItemListReq) (res *v1.DictItemListRes, err error) {
	list, err := service.SysDict().GetItemsByDictId(ctx, req.DictId)
	if err != nil {
		return nil, gerror.Wrap(err, "查询字典项列表失败")
	}

	return &v1.DictItemListRes{
		List: list,
	}, nil
}

// GetItemsByCode 根据字典编码获取字典项列表
func (c *cDict) GetItemsByCode(ctx context.Context, req *v1.DictItemsByCodeReq) (res *v1.DictItemsByCodeRes, err error) {
	list, err := service.SysDict().GetItemsByDictCode(ctx, req.DictCode)
	if err != nil {
		return nil, gerror.Wrap(err, "查询字典项失败")
	}

	return &v1.DictItemsByCodeRes{
		List: list,
	}, nil
}

// AddItem 添加字典项
func (c *cDict) AddItem(ctx context.Context, req *v1.DictItemAddReq) (res *v1.DictItemAddRes, err error) {
	err = service.SysDict().AddItem(ctx, req)
	if err != nil {
		return nil, gerror.Wrap(err, "添加字典项失败")
	}

	return &v1.DictItemAddRes{}, nil
}

// EditItem 编辑字典项
func (c *cDict) EditItem(ctx context.Context, req *v1.DictItemEditReq) (res *v1.DictItemEditRes, err error) {
	err = service.SysDict().EditItem(ctx, req)
	if err != nil {
		return nil, gerror.Wrap(err, "编辑字典项失败")
	}

	return &v1.DictItemEditRes{}, nil
}

// DeleteItem 删除字典项
func (c *cDict) DeleteItem(ctx context.Context, req *v1.DictItemDeleteReq) (res *v1.DictItemDeleteRes, err error) {
	err = service.SysDict().DeleteItem(ctx, req.Id)
	if err != nil {
		return nil, gerror.Wrap(err, "删除字典项失败")
	}

	return &v1.DictItemDeleteRes{}, nil
}
