package sys

import (
	"context"
	"strings"

	"github.com/gogf/gf/v2/errors/gerror"

	v1 "teaching-open/api/v1/sys"
	"teaching-open/internal/model/entity"
	"teaching-open/internal/service"
)

type TeachingCourseUnitController struct{}

var TeachingCourseUnit = &TeachingCourseUnitController{}

func (c *TeachingCourseUnitController) GetList(ctx context.Context, req *v1.CourseUnitListReq) (res *v1.CourseUnitListRes, err error) {
	list, total, err := service.TeachingCourseUnit().GetList(ctx, req.CourseId, req.UnitName, req.PageNo, req.PageSize)
	if err != nil {
		return nil, gerror.Wrap(err, "查询课程单元列表失败")
	}
	return &v1.CourseUnitListRes{List: list, Total: total, PageNo: req.PageNo, PageSize: req.PageSize}, nil
}

func (c *TeachingCourseUnitController) GetByCourseId(ctx context.Context, req *v1.CourseUnitAllReq) (res *v1.CourseUnitAllRes, err error) {
	list, err := service.TeachingCourseUnit().GetByCourseId(ctx, req.CourseId)
	if err != nil {
		return nil, gerror.Wrap(err, "查询课程单元失败")
	}
	return &v1.CourseUnitAllRes{List: list}, nil
}

func (c *TeachingCourseUnitController) Add(ctx context.Context, req *v1.CourseUnitAddReq) (res *v1.CourseUnitAddRes, err error) {
	unit := &entity.TeachingCourseUnit{
		UnitName: req.UnitName, UnitIntro: req.UnitIntro, UnitCover: req.UnitCover,
		CourseId: req.CourseId, CourseVideo: req.CourseVideo, CourseVideoSource: req.CourseVideoSource,
		ShowCourseVideo: req.ShowCourseVideo, CourseCase: req.CourseCase, ShowCourseCase: req.ShowCourseCase,
		CoursePpt: req.CoursePpt, ShowCoursePpt: req.ShowCoursePpt, CourseWorkType: req.CourseWorkType,
		CourseWork: req.CourseWork, CourseWorkAnswer: req.CourseWorkAnswer, CoursePlan: req.CoursePlan,
		ShowCoursePlan: req.ShowCoursePlan, MapX: req.MapX, MapY: req.MapY,
		MediaContent: req.MediaContent, OrderNum: req.OrderNum,
	}
	err = service.TeachingCourseUnit().Add(ctx, unit)
	if err != nil {
		return nil, gerror.Wrap(err, "添加课程单元失败")
	}
	return &v1.CourseUnitAddRes{}, nil
}

func (c *TeachingCourseUnitController) Edit(ctx context.Context, req *v1.CourseUnitEditReq) (res *v1.CourseUnitEditRes, err error) {
	unit := &entity.TeachingCourseUnit{
		Id: req.Id, UnitName: req.UnitName, UnitIntro: req.UnitIntro, UnitCover: req.UnitCover,
		CourseId: req.CourseId, CourseVideo: req.CourseVideo, CourseVideoSource: req.CourseVideoSource,
		ShowCourseVideo: req.ShowCourseVideo, CourseCase: req.CourseCase, ShowCourseCase: req.ShowCourseCase,
		CoursePpt: req.CoursePpt, ShowCoursePpt: req.ShowCoursePpt, CourseWorkType: req.CourseWorkType,
		CourseWork: req.CourseWork, CourseWorkAnswer: req.CourseWorkAnswer, CoursePlan: req.CoursePlan,
		ShowCoursePlan: req.ShowCoursePlan, MapX: req.MapX, MapY: req.MapY,
		MediaContent: req.MediaContent, OrderNum: req.OrderNum,
	}
	err = service.TeachingCourseUnit().Edit(ctx, unit)
	if err != nil {
		return nil, gerror.Wrap(err, "编辑课程单元失败")
	}
	return &v1.CourseUnitEditRes{}, nil
}

func (c *TeachingCourseUnitController) Delete(ctx context.Context, req *v1.CourseUnitDeleteReq) (res *v1.CourseUnitDeleteRes, err error) {
	err = service.TeachingCourseUnit().Delete(ctx, req.Id)
	if err != nil {
		return nil, gerror.Wrap(err, "删除课程单元失败")
	}
	return &v1.CourseUnitDeleteRes{}, nil
}

func (c *TeachingCourseUnitController) DeleteBatch(ctx context.Context, req *v1.CourseUnitDeleteBatchReq) (res *v1.CourseUnitDeleteBatchRes, err error) {
	ids := strings.Split(req.Ids, ",")
	err = service.TeachingCourseUnit().DeleteBatch(ctx, ids)
	if err != nil {
		return nil, gerror.Wrap(err, "批量删除课程单元失败")
	}
	return &v1.CourseUnitDeleteBatchRes{}, nil
}

func (c *TeachingCourseUnitController) GetById(ctx context.Context, req *v1.CourseUnitGetByIdReq) (res *v1.CourseUnitGetByIdRes, err error) {
	unit, err := service.TeachingCourseUnit().GetById(ctx, req.Id)
	if err != nil {
		return nil, gerror.Wrap(err, "查询课程单元详情失败")
	}
	if unit == nil {
		return nil, gerror.New("课程单元不存在")
	}
	return &v1.CourseUnitGetByIdRes{TeachingCourseUnit: unit}, nil
}

func (c *TeachingCourseUnitController) Sort(ctx context.Context, req *v1.CourseUnitSortReq) (res *v1.CourseUnitSortRes, err error) {
	ids := strings.Split(req.UnitIds, ",")
	err = service.TeachingCourseUnit().Sort(ctx, ids)
	if err != nil {
		return nil, gerror.Wrap(err, "课程单元排序失败")
	}
	return &v1.CourseUnitSortRes{}, nil
}
