package sys

import (
	"context"
	"errors"

	v1 "teaching-open/api/v1/sys"
	"teaching-open/internal/service"
)

var TeachingNews = cTeachingNews{}

type cTeachingNews struct{}

// List 新闻列表
func (c *cTeachingNews) List(ctx context.Context, req *v1.NewsListReq) (res *v1.NewsListRes, err error) {
	list, total, err := service.TeachingNews().List(ctx, req)
	if err != nil {
		return nil, err
	}

	return &v1.NewsListRes{
		List:  list,
		Total: total,
	}, nil
}

// PublicList 公开新闻列表
func (c *cTeachingNews) PublicList(ctx context.Context, req *v1.NewsPublicListReq) (res *v1.NewsPublicListRes, err error) {
	list, total, err := service.TeachingNews().PublicList(ctx, req)
	if err != nil {
		return nil, err
	}

	return &v1.NewsPublicListRes{
		List:  list,
		Total: total,
	}, nil
}

// Add 添加新闻
func (c *cTeachingNews) Add(ctx context.Context, req *v1.NewsAddReq) (res *v1.NewsAddRes, err error) {
	id, err := service.TeachingNews().Add(ctx, req)
	if err != nil {
		return nil, err
	}

	return &v1.NewsAddRes{
		Id: id,
	}, nil
}

// Edit 编辑新闻
func (c *cTeachingNews) Edit(ctx context.Context, req *v1.NewsEditReq) (res *v1.NewsEditRes, err error) {
	err = service.TeachingNews().Edit(ctx, req)
	if err != nil {
		return nil, err
	}

	return &v1.NewsEditRes{}, nil
}

// Delete 删除新闻
func (c *cTeachingNews) Delete(ctx context.Context, req *v1.NewsDeleteReq) (res *v1.NewsDeleteRes, err error) {
	err = service.TeachingNews().Delete(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return &v1.NewsDeleteRes{}, nil
}

// DeleteBatch 批量删除新闻
func (c *cTeachingNews) DeleteBatch(ctx context.Context, req *v1.NewsDeleteBatchReq) (res *v1.NewsDeleteBatchRes, err error) {
	err = service.TeachingNews().DeleteBatch(ctx, req.Ids)
	if err != nil {
		return nil, err
	}

	return &v1.NewsDeleteBatchRes{}, nil
}

// QueryById 获取新闻详情
func (c *cTeachingNews) QueryById(ctx context.Context, req *v1.NewsGetByIdReq) (res *v1.NewsGetByIdRes, err error) {
	info, err := service.TeachingNews().GetById(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	if info == nil {
		return nil, errors.New("新闻不存在")
	}

	return &v1.NewsGetByIdRes{
		NewsInfo: info,
	}, nil
}

// Publish 发布新闻
func (c *cTeachingNews) Publish(ctx context.Context, req *v1.NewsPublishReq) (res *v1.NewsPublishRes, err error) {
	err = service.TeachingNews().Publish(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return &v1.NewsPublishRes{}, nil
}

// Offline 下架新闻
func (c *cTeachingNews) Offline(ctx context.Context, req *v1.NewsOfflineReq) (res *v1.NewsOfflineRes, err error) {
	err = service.TeachingNews().Offline(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return &v1.NewsOfflineRes{}, nil
}
