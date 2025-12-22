package service

import (
	"context"
	"teaching-open/internal/model/entity"
)

type ITeachingCourseUnit interface {
	GetList(ctx context.Context, courseId, unitName string, pageNo, pageSize int) (list []*entity.TeachingCourseUnit, total int, err error)
	GetByCourseId(ctx context.Context, courseId string) (list []*entity.TeachingCourseUnit, err error)
	Add(ctx context.Context, unit *entity.TeachingCourseUnit) (err error)
	Edit(ctx context.Context, unit *entity.TeachingCourseUnit) (err error)
	Delete(ctx context.Context, id string) (err error)
	DeleteBatch(ctx context.Context, ids []string) (err error)
	GetById(ctx context.Context, id string) (unit *entity.TeachingCourseUnit, err error)
	Sort(ctx context.Context, unitIds []string) (err error)
}

var localTeachingCourseUnit ITeachingCourseUnit

func TeachingCourseUnit() ITeachingCourseUnit {
	if localTeachingCourseUnit == nil {
		panic("implement not found for interface ITeachingCourseUnit, forgot register?")
	}
	return localTeachingCourseUnit
}

func RegisterTeachingCourseUnit(i ITeachingCourseUnit) {
	localTeachingCourseUnit = i
}