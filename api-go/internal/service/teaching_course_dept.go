package service

import (
	"context"
	"teaching-open/internal/model/entity"
)

// ITeachingCourseDept 班级课程管理服务接口
type ITeachingCourseDept interface {
	// GetList 获取班级课程列表（分页）
	GetList(ctx context.Context, deptId, courseId string, pageNo, pageSize int) (list []*entity.TeachingCourseDept, total int, err error)
	// GetByDeptId 获取班级的课程列表
	GetByDeptId(ctx context.Context, deptId string) (list []*entity.TeachingCourseDept, err error)
	// GetByCourseId 获取课程授权的班级列表
	GetByCourseId(ctx context.Context, courseId string) (list []*entity.TeachingCourseDept, err error)
	// AddOrUpdate 添加或更新班级课程
	AddOrUpdate(ctx context.Context, courseDept *entity.TeachingCourseDept) (err error)
	// Delete 删除班级课程
	Delete(ctx context.Context, id string) (err error)
	// BatchAdd 批量添加班级课程
	BatchAdd(ctx context.Context, deptIds []string, courseId string) (err error)
}

var localTeachingCourseDept ITeachingCourseDept

func TeachingCourseDept() ITeachingCourseDept {
	if localTeachingCourseDept == nil {
		panic("implement not found for interface ITeachingCourseDept, forgot register?")
	}
	return localTeachingCourseDept
}

func RegisterTeachingCourseDept(i ITeachingCourseDept) {
	localTeachingCourseDept = i
}
