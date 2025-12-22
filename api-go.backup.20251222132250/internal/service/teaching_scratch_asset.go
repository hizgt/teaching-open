package service

import (
	"context"

	v1 "teaching-open/api/v1/sys"
)

// ITeachingScratchAsset Scratch素材服务接口
type ITeachingScratchAsset interface {
	List(ctx context.Context, req *v1.ScratchAssetListReq) (list interface{}, total int, err error)
	GetScratchAssets(ctx context.Context, assetType int) (list interface{}, err error)
	Add(ctx context.Context, req *v1.ScratchAssetAddReq) (id string, err error)
	Edit(ctx context.Context, req *v1.ScratchAssetEditReq) error
	Delete(ctx context.Context, id string) error
	DeleteBatch(ctx context.Context, ids string) error
	GetById(ctx context.Context, id string) (*v1.ScratchAssetInfo, error)
}

var localTeachingScratchAsset ITeachingScratchAsset

func TeachingScratchAsset() ITeachingScratchAsset {
	if localTeachingScratchAsset == nil {
		panic("ITeachingScratchAsset service not registered")
	}
	return localTeachingScratchAsset
}

func RegisterTeachingScratchAsset(s ITeachingScratchAsset) {
	localTeachingScratchAsset = s
}
