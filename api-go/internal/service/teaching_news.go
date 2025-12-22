package service

import (
	"context"

	v1 "teaching-open/api/v1/sys"
)

// ITeachingNews 新闻公告服务接口
type ITeachingNews interface {
	List(ctx context.Context, req *v1.NewsListReq) (list interface{}, total int, err error)
	PublicList(ctx context.Context, req *v1.NewsPublicListReq) (list interface{}, total int, err error)
	Add(ctx context.Context, req *v1.NewsAddReq) (id string, err error)
	Edit(ctx context.Context, req *v1.NewsEditReq) error
	Delete(ctx context.Context, id string) error
	DeleteBatch(ctx context.Context, ids string) error
	GetById(ctx context.Context, id string) (*v1.NewsInfo, error)
	Publish(ctx context.Context, id string) error
	Offline(ctx context.Context, id string) error
}

var localTeachingNews ITeachingNews

func TeachingNews() ITeachingNews {
	if localTeachingNews == nil {
		panic("ITeachingNews service not registered")
	}
	return localTeachingNews
}

func RegisterTeachingNews(s ITeachingNews) {
	localTeachingNews = s
}
