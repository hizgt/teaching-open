package service

import (
	"context"

	v1 "teaching-open/api/v1/sys"
)

// ITeachingAdditionalWork 附加作业服务接口
type ITeachingAdditionalWork interface {
	List(ctx context.Context, req *v1.AdditionalWorkListReq) (list interface{}, total int, err error)
	ListByDept(ctx context.Context, req *v1.AdditionalWorkByDeptReq) (list interface{}, total int, err error)
	Add(ctx context.Context, req *v1.AdditionalWorkAddReq) (id string, err error)
	Edit(ctx context.Context, req *v1.AdditionalWorkEditReq) error
	Delete(ctx context.Context, id string) error
	DeleteBatch(ctx context.Context, ids string) error
	GetById(ctx context.Context, id string) (*v1.AdditionalWorkInfo, error)
	Publish(ctx context.Context, id string) error
	Offline(ctx context.Context, id string) error
}

var localTeachingAdditionalWork ITeachingAdditionalWork

func TeachingAdditionalWork() ITeachingAdditionalWork {
	if localTeachingAdditionalWork == nil {
		panic("ITeachingAdditionalWork service not registered")
	}
	return localTeachingAdditionalWork
}

func RegisterTeachingAdditionalWork(s ITeachingAdditionalWork) {
	localTeachingAdditionalWork = s
}
