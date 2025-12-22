package sys

import (
	"context"
	"errors"

	v1 "teaching-open/api/v1/sys"
	"teaching-open/internal/service"
)

var TeachingAdditionalWork = cTeachingAdditionalWork{}

type cTeachingAdditionalWork struct{}

// List 附加作业列表
func (c *cTeachingAdditionalWork) List(ctx context.Context, req *v1.AdditionalWorkListReq) (res *v1.AdditionalWorkListRes, err error) {
	list, total, err := service.TeachingAdditionalWork().List(ctx, req)
	if err != nil {
		return nil, err
	}

	return &v1.AdditionalWorkListRes{
		List:  list,
		Total: total,
	}, nil
}

// ListByDept 按班级获取附加作业列表
func (c *cTeachingAdditionalWork) ListByDept(ctx context.Context, req *v1.AdditionalWorkByDeptReq) (res *v1.AdditionalWorkByDeptRes, err error) {
	list, total, err := service.TeachingAdditionalWork().ListByDept(ctx, req)
	if err != nil {
		return nil, err
	}

	return &v1.AdditionalWorkByDeptRes{
		List:  list,
		Total: total,
	}, nil
}

// Add 添加附加作业
func (c *cTeachingAdditionalWork) Add(ctx context.Context, req *v1.AdditionalWorkAddReq) (res *v1.AdditionalWorkAddRes, err error) {
	id, err := service.TeachingAdditionalWork().Add(ctx, req)
	if err != nil {
		return nil, err
	}

	return &v1.AdditionalWorkAddRes{
		Id: id,
	}, nil
}

// Edit 编辑附加作业
func (c *cTeachingAdditionalWork) Edit(ctx context.Context, req *v1.AdditionalWorkEditReq) (res *v1.AdditionalWorkEditRes, err error) {
	err = service.TeachingAdditionalWork().Edit(ctx, req)
	if err != nil {
		return nil, err
	}

	return &v1.AdditionalWorkEditRes{}, nil
}

// Delete 删除附加作业
func (c *cTeachingAdditionalWork) Delete(ctx context.Context, req *v1.AdditionalWorkDeleteReq) (res *v1.AdditionalWorkDeleteRes, err error) {
	err = service.TeachingAdditionalWork().Delete(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return &v1.AdditionalWorkDeleteRes{}, nil
}

// DeleteBatch 批量删除附加作业
func (c *cTeachingAdditionalWork) DeleteBatch(ctx context.Context, req *v1.AdditionalWorkDeleteBatchReq) (res *v1.AdditionalWorkDeleteBatchRes, err error) {
	err = service.TeachingAdditionalWork().DeleteBatch(ctx, req.Ids)
	if err != nil {
		return nil, err
	}

	return &v1.AdditionalWorkDeleteBatchRes{}, nil
}

// QueryById 获取附加作业详情
func (c *cTeachingAdditionalWork) QueryById(ctx context.Context, req *v1.AdditionalWorkGetByIdReq) (res *v1.AdditionalWorkGetByIdRes, err error) {
	info, err := service.TeachingAdditionalWork().GetById(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	if info == nil {
		return nil, errors.New("附加作业不存在")
	}

	return &v1.AdditionalWorkGetByIdRes{
		Info: info,
	}, nil
}

// Publish 发布附加作业
func (c *cTeachingAdditionalWork) Publish(ctx context.Context, req *v1.AdditionalWorkPublishReq) (res *v1.AdditionalWorkPublishRes, err error) {
	err = service.TeachingAdditionalWork().Publish(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return &v1.AdditionalWorkPublishRes{}, nil
}

// Offline 下架附加作业
func (c *cTeachingAdditionalWork) Offline(ctx context.Context, req *v1.AdditionalWorkOfflineReq) (res *v1.AdditionalWorkOfflineRes, err error) {
	err = service.TeachingAdditionalWork().Offline(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return &v1.AdditionalWorkOfflineRes{}, nil
}
