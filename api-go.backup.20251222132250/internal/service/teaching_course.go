// Package service provides service interfaces for teaching management.
package service

import (
	"context"

	"teaching-open/internal/model/entity"
)

// ITeachingCourse 课程管理服务接口
type ITeachingCourse interface {
	// GetList 获取课程列表
	GetList(ctx context.Context, courseName, courseType, courseCategory string, isShared, showHome, pageNo, pageSize int) (list []*entity.TeachingCourse, total int, err error)

	// GetHomeCourse 获取首页展示课程
	GetHomeCourse(ctx context.Context) (list []*entity.TeachingCourse, err error)

	// Add 添加课程
	Add(ctx context.Context, course *entity.TeachingCourse) (err error)

	// Edit 编辑课程
	Edit(ctx context.Context, course *entity.TeachingCourse) (err error)

	// Delete 删除课程（软删除）
	Delete(ctx context.Context, id string) (err error)

	// DeleteBatch 批量删除课程（软删除）
	DeleteBatch(ctx context.Context, ids []string) (err error)

	// GetById 根据ID获取课程详情
	GetById(ctx context.Context, id string) (course *entity.TeachingCourse, err error)

	// Publish 发布/下架课程
	Publish(ctx context.Context, id string, showHome int) (err error)

	// SetShared 设置共享状态
	SetShared(ctx context.Context, id string, isShared int) (err error)

	// AuthorizeDept 授权部门
	AuthorizeDept(ctx context.Context, id string, departIds string) (err error)
}

var localTeachingCourse ITeachingCourse

// TeachingCourse 获取课程管理服务实例
func TeachingCourse() ITeachingCourse {
	if localTeachingCourse == nil {
		panic("implement not found for interface ITeachingCourse, forgot register?")
	}
	return localTeachingCourse
}

// RegisterTeachingCourse 注册课程管理服务实现
func RegisterTeachingCourse(i ITeachingCourse) {
	localTeachingCourse = i
}
