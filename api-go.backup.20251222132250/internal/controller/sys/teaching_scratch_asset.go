package sys

import (
	"context"
	"errors"

	v1 "teaching-open/api/v1/sys"
	"teaching-open/internal/service"
)

var TeachingScratchAsset = cTeachingScratchAsset{}

type cTeachingScratchAsset struct{}

// List 素材列表
func (c *cTeachingScratchAsset) List(ctx context.Context, req *v1.ScratchAssetListReq) (res *v1.ScratchAssetListRes, err error) {
	list, total, err := service.TeachingScratchAsset().List(ctx, req)
	if err != nil {
		return nil, err
	}

	return &v1.ScratchAssetListRes{
		List:  list,
		Total: total,
	}, nil
}

// GetScratchAssets 获取Scratch素材（编辑器用）
func (c *cTeachingScratchAsset) GetScratchAssets(ctx context.Context, req *v1.ScratchAssetGetReq) (res *v1.ScratchAssetGetRes, err error) {
	list, err := service.TeachingScratchAsset().GetScratchAssets(ctx, req.AssetType)
	if err != nil {
		return nil, err
	}

	return &v1.ScratchAssetGetRes{
		List: list,
	}, nil
}

// Add 添加素材
func (c *cTeachingScratchAsset) Add(ctx context.Context, req *v1.ScratchAssetAddReq) (res *v1.ScratchAssetAddRes, err error) {
	id, err := service.TeachingScratchAsset().Add(ctx, req)
	if err != nil {
		return nil, err
	}

	return &v1.ScratchAssetAddRes{
		Id: id,
	}, nil
}

// Edit 编辑素材
func (c *cTeachingScratchAsset) Edit(ctx context.Context, req *v1.ScratchAssetEditReq) (res *v1.ScratchAssetEditRes, err error) {
	err = service.TeachingScratchAsset().Edit(ctx, req)
	if err != nil {
		return nil, err
	}

	return &v1.ScratchAssetEditRes{}, nil
}

// Delete 删除素材
func (c *cTeachingScratchAsset) Delete(ctx context.Context, req *v1.ScratchAssetDeleteReq) (res *v1.ScratchAssetDeleteRes, err error) {
	err = service.TeachingScratchAsset().Delete(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return &v1.ScratchAssetDeleteRes{}, nil
}

// DeleteBatch 批量删除素材
func (c *cTeachingScratchAsset) DeleteBatch(ctx context.Context, req *v1.ScratchAssetDeleteBatchReq) (res *v1.ScratchAssetDeleteBatchRes, err error) {
	err = service.TeachingScratchAsset().DeleteBatch(ctx, req.Ids)
	if err != nil {
		return nil, err
	}

	return &v1.ScratchAssetDeleteBatchRes{}, nil
}

// QueryById 获取素材详情
func (c *cTeachingScratchAsset) QueryById(ctx context.Context, req *v1.ScratchAssetGetByIdReq) (res *v1.ScratchAssetGetByIdRes, err error) {
	info, err := service.TeachingScratchAsset().GetById(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	if info == nil {
		return nil, errors.New("素材不存在")
	}

	return &v1.ScratchAssetGetByIdRes{
		ScratchAssetInfo: info,
	}, nil
}
