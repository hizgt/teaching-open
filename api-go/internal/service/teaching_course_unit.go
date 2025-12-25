package service

import (
	"context"
	"teaching-open/internal/dao"
	"teaching-open/internal/model/entity"
	"teaching-open/internal/model/vo"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/guid"
	"github.com/gogf/gf/v2/os/gtime"
)

// TeachingCourseUnitService 课程单元服务接口
type TeachingCourseUnitService interface {
	List(ctx context.Context, req *vo.CourseUnitListReq) (*vo.CourseUnitListRes, error)
	Create(ctx context.Context, req *vo.CourseUnitCreateReq, userId string) (*vo.CourseUnitCreateRes, error)
	Update(ctx context.Context, req *vo.CourseUnitUpdateReq, userId string) error
	Delete(ctx context.Context, req *vo.CourseUnitDeleteReq, userId string) error
	Sort(ctx context.Context, req *vo.CourseUnitSortReq, userId string) error
	QueryByCourseId(ctx context.Context, courseId string) ([]vo.UnitItem, error)
	QueryById(ctx context.Context, id string) (*vo.UnitItem, error)
}

// teachingCourseUnitServiceImpl 课程单元服务实现
type teachingCourseUnitServiceImpl struct{}

// NewTeachingCourseUnitService 创建课程单元服务实例
func NewTeachingCourseUnitService() TeachingCourseUnitService {
	return &teachingCourseUnitServiceImpl{}
}

// List 获取课程单元列表
func (s *teachingCourseUnitServiceImpl) List(ctx context.Context, req *vo.CourseUnitListReq) (*vo.CourseUnitListRes, error) {
	// 检查课程是否存在
	courseCount, err := dao.Course.Ctx(ctx).Where(dao.Course.Columns().Id, req.CourseId).Count()
	if err != nil {
		return nil, err
	}
	if courseCount == 0 {
		return nil, gerror.New("课程不存在")
	}

	var units []entity.TeachingCourseUnit
	err = dao.CourseUnit.Ctx(ctx).
		Where(dao.CourseUnit.Columns().CourseId, req.CourseId).
		OrderAsc(dao.CourseUnit.Columns().SortOrder).
		Scan(&units)
	if err != nil {
		return nil, err
	}

	// 转换为VO
	var records []vo.UnitItem
	for _, unit := range units {
		records = append(records, vo.UnitItem{
			Id:          unit.Id,
			Name:        unit.Name,
			Content:     unit.Content,
			SortOrder:   unit.SortOrder,
			ResourceUrl: unit.ResourceUrl,
		})
	}

	return &vo.CourseUnitListRes{Records: records}, nil
}

// Create 创建课程单元
func (s *teachingCourseUnitServiceImpl) Create(ctx context.Context, req *vo.CourseUnitCreateReq, userId string) (*vo.CourseUnitCreateRes, error) {
	// 检查课程是否存在且为草稿状态
	var course entity.TeachingCourse
	err := dao.Course.Ctx(ctx).
		Where(dao.Course.Columns().Id, req.CourseId).
		Scan(&course)
	if err != nil {
		return nil, err
	}
	if course.Id == "" {
		return nil, gerror.New("课程不存在")
	}
	if course.Status != "draft" {
		return nil, gerror.New("只能在草稿状态下添加单元")
	}

	// 获取最大排序号
	maxSort, err := dao.CourseUnit.Ctx(ctx).
		Where(dao.CourseUnit.Columns().CourseId, req.CourseId).
		Fields("MAX(sort_order) as max_sort").
		Value()
	if err != nil {
		return nil, err
	}

	sortOrder := 1
	if maxSort.Int() > 0 {
		sortOrder = maxSort.Int() + 1
	}

	unitId := guid.S()
	unit := entity.TeachingCourseUnit{
		Id:          unitId,
		CourseId:    req.CourseId,
		Name:        req.Name,
		Content:     req.Content,
		SortOrder:   sortOrder,
		ResourceUrl: req.ResourceUrl,
		CreateBy:    userId,
		CreateTime:  gtime.Now(),
		UpdateBy:    userId,
		UpdateTime:  gtime.Now(),
	}

	_, err = dao.CourseUnit.Ctx(ctx).Data(unit).Insert()
	if err != nil {
		return nil, err
	}

	return &vo.CourseUnitCreateRes{Id: unitId}, nil
}

// Update 更新课程单元
func (s *teachingCourseUnitServiceImpl) Update(ctx context.Context, req *vo.CourseUnitUpdateReq, userId string) error {
	// 检查单元是否存在
	count, err := dao.CourseUnit.Ctx(ctx).Where(dao.CourseUnit.Columns().Id, req.Id).Count()
	if err != nil {
		return err
	}
	if count == 0 {
		return gerror.New("课程单元不存在")
	}

	// 检查课程是否为草稿状态
	var unit entity.TeachingCourseUnit
	err = dao.CourseUnit.Ctx(ctx).Where(dao.CourseUnit.Columns().Id, req.Id).Scan(&unit)
	if err != nil {
		return err
	}

	var course entity.TeachingCourse
	err = dao.Course.Ctx(ctx).Where(dao.Course.Columns().Id, unit.CourseId).Scan(&course)
	if err != nil {
		return err
	}
	if course.Status != "draft" {
		return gerror.New("只能在草稿状态下编辑单元")
	}

	// 更新单元
	data := g.Map{
		dao.CourseUnit.Columns().Name:        req.Name,
		dao.CourseUnit.Columns().Content:     req.Content,
		dao.CourseUnit.Columns().SortOrder:   req.SortOrder,
		dao.CourseUnit.Columns().ResourceUrl: req.ResourceUrl,
		dao.CourseUnit.Columns().UpdateBy:    userId,
		dao.CourseUnit.Columns().UpdateTime:  gtime.Now(),
	}

	_, err = dao.CourseUnit.Ctx(ctx).Data(data).Where(dao.CourseUnit.Columns().Id, req.Id).Update()
	return err
}

// Delete 删除课程单元
func (s *teachingCourseUnitServiceImpl) Delete(ctx context.Context, req *vo.CourseUnitDeleteReq, userId string) error {
	// 检查单元是否存在
	var unit entity.TeachingCourseUnit
	err := dao.CourseUnit.Ctx(ctx).Where(dao.CourseUnit.Columns().Id, req.Id).Scan(&unit)
	if err != nil {
		return err
	}
	if unit.Id == "" {
		return gerror.New("课程单元不存在")
	}

	// 检查课程是否为草稿状态
	var course entity.TeachingCourse
	err = dao.Course.Ctx(ctx).Where(dao.Course.Columns().Id, unit.CourseId).Scan(&course)
	if err != nil {
		return err
	}
	if course.Status != "draft" {
		return gerror.New("只能在草稿状态下删除单元")
	}

	// 删除单元
	_, err = dao.CourseUnit.Ctx(ctx).Where(dao.CourseUnit.Columns().Id, req.Id).Delete()
	if err != nil {
		return err
	}

	// 重新排序剩余单元
	return s.reorderUnits(ctx, unit.CourseId)
}

// Sort 排序课程单元
func (s *teachingCourseUnitServiceImpl) Sort(ctx context.Context, req *vo.CourseUnitSortReq, userId string) error {
	// 检查课程是否存在且为草稿状态
	var course entity.TeachingCourse
	err := dao.Course.Ctx(ctx).
		Where(dao.Course.Columns().Id, req.CourseId).
		Scan(&course)
	if err != nil {
		return err
	}
	if course.Id == "" {
		return gerror.New("课程不存在")
	}
	if course.Status != "draft" {
		return gerror.New("只能在草稿状态下排序单元")
	}

	// 批量更新排序
	for _, unitSort := range req.Units {
		data := g.Map{
			dao.CourseUnit.Columns().SortOrder: unitSort.SortOrder,
			dao.CourseUnit.Columns().UpdateBy:  userId,
			dao.CourseUnit.Columns().UpdateTime: gtime.Now(),
		}

		_, err = dao.CourseUnit.Ctx(ctx).
			Data(data).
			Where(dao.CourseUnit.Columns().Id, unitSort.Id).
			Where(dao.CourseUnit.Columns().CourseId, req.CourseId).
			Update()
		if err != nil {
			return err
		}
	}

	return nil
}

// reorderUnits 重新排序单元
func (s *teachingCourseUnitServiceImpl) reorderUnits(ctx context.Context, courseId string) error {
	var units []entity.TeachingCourseUnit
	err := dao.CourseUnit.Ctx(ctx).
		Where(dao.CourseUnit.Columns().CourseId, courseId).
		OrderAsc(dao.CourseUnit.Columns().SortOrder).
		Scan(&units)
	if err != nil {
		return err
	}

	// 重新分配排序号
	for i, unit := range units {
		data := g.Map{
			dao.CourseUnit.Columns().SortOrder: i + 1,
		}

		_, err = dao.CourseUnit.Ctx(ctx).
			Data(data).
			Where(dao.CourseUnit.Columns().Id, unit.Id).
			Update()
		if err != nil {
			return err
		}
	}

	return nil
}

// QueryByCourseId 根据课程ID查询单元
func (s *teachingCourseUnitServiceImpl) QueryByCourseId(ctx context.Context, courseId string) ([]vo.UnitItem, error) {
	var units []entity.TeachingCourseUnit
	err := dao.CourseUnit.Ctx(ctx).
		Where(dao.CourseUnit.Columns().CourseId, courseId).
		OrderAsc(dao.CourseUnit.Columns().SortOrder).
		Scan(&units)
	if err != nil {
		return nil, err
	}

	var records []vo.UnitItem
	for _, unit := range units {
		records = append(records, vo.UnitItem{
			Id:          unit.Id,
			Name:        unit.Name,
			Content:     unit.Content,
			SortOrder:   unit.SortOrder,
			ResourceUrl: unit.ResourceUrl,
		})
	}

	return records, nil
}

// QueryById 根据ID查询单元
func (s *teachingCourseUnitServiceImpl) QueryById(ctx context.Context, id string) (*vo.UnitItem, error) {
	var unit entity.TeachingCourseUnit
	err := dao.CourseUnit.Ctx(ctx).Where(dao.CourseUnit.Columns().Id, id).Scan(&unit)
	if err != nil {
		return nil, err
	}
	if unit.Id == "" {
		return nil, gerror.New("课程单元不存在")
	}

	return &vo.UnitItem{
		Id:          unit.Id,
		Name:        unit.Name,
		Content:     unit.Content,
		SortOrder:   unit.SortOrder,
		ResourceUrl: unit.ResourceUrl,
	}, nil
}
