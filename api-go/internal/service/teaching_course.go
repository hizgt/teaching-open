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

// TeachingCourseService 课程服务接口
type TeachingCourseService interface {
	List(ctx context.Context, req *vo.CourseListReq) (*vo.CourseListRes, error)
	Create(ctx context.Context, req *vo.CourseCreateReq, userId string) (*vo.CourseCreateRes, error)
	Update(ctx context.Context, req *vo.CourseUpdateReq, userId string) error
	Delete(ctx context.Context, req *vo.CourseDeleteReq, userId string) error
	Detail(ctx context.Context, id string) (*vo.CourseDetailRes, error)
	Publish(ctx context.Context, req *vo.CoursePublishReq, userId string) error
}

// teachingCourseServiceImpl 课程服务实现
type teachingCourseServiceImpl struct{}

// NewTeachingCourseService 创建课程服务实例
func NewTeachingCourseService() TeachingCourseService {
	return &teachingCourseServiceImpl{}
}

// List 获取课程列表
func (s *teachingCourseServiceImpl) List(ctx context.Context, req *vo.CourseListReq) (*vo.CourseListRes, error) {
	model := dao.Course.Ctx(ctx)

	// 构建查询条件
	if req.Name != "" {
		model = model.WhereLike(dao.Course.Columns().Name, "%"+req.Name+"%")
	}
	if req.Type != "" {
		model = model.Where(dao.Course.Columns().Type, req.Type)
	}
	if req.Status != "" {
		model = model.Where(dao.Course.Columns().Status, req.Status)
	}
	if req.CreateBy != "" {
		model = model.Where(dao.Course.Columns().CreateBy, req.CreateBy)
	}

	// 分页查询
	total, err := model.Count()
	if err != nil {
		return nil, err
	}

	var courses []entity.TeachingCourse
	err = model.Page(req.Page, req.PageSize).OrderDesc(dao.Course.Columns().CreateTime).Scan(&courses)
	if err != nil {
		return nil, err
	}

	// 转换为VO
	var records []vo.CourseItem
	for _, course := range courses {
		// 获取单元数量
		unitCount, _ := dao.CourseUnit.Ctx(ctx).Where(dao.CourseUnit.Columns().CourseId, course.Id).Count()

		records = append(records, vo.CourseItem{
			Id:          course.Id,
			Name:        course.Name,
			Type:        course.Type,
			Description: course.Description,
			CoverImage:  course.CoverImage,
			Status:      course.Status,
			CreateBy:    course.CreateBy,
			CreateTime:  course.CreateTime.Time.Format("2006-01-02 15:04:05"),
			UnitCount:   int(unitCount),
		})
	}

	return &vo.CourseListRes{
		Records:   records,
		Total:     total,
		Page:      req.Page,
		PageSize:  req.PageSize,
	}, nil
}

// Create 创建课程
func (s *teachingCourseServiceImpl) Create(ctx context.Context, req *vo.CourseCreateReq, userId string) (*vo.CourseCreateRes, error) {
	courseId := guid.S()

	// 创建课程
	course := entity.TeachingCourse{
		Id:          courseId,
		Name:        req.Name,
		Type:        req.Type,
		Description: req.Description,
		CoverImage:  req.CoverImage,
		Status:      "draft",
		CreateBy:    userId,
		CreateTime:  gtime.Now(),
		UpdateBy:    userId,
		UpdateTime:  gtime.Now(),
	}

	_, err := dao.Course.Ctx(ctx).Data(course).Insert()
	if err != nil {
		return nil, err
	}

	// 关联部门
	if len(req.DeptIds) > 0 {
		err = s.assignCourseToDepts(ctx, courseId, req.DeptIds, userId)
		if err != nil {
			return nil, err
		}
	}

	return &vo.CourseCreateRes{Id: courseId}, nil
}

// Update 更新课程
func (s *teachingCourseServiceImpl) Update(ctx context.Context, req *vo.CourseUpdateReq, userId string) error {
	// 检查课程是否存在
	count, err := dao.Course.Ctx(ctx).Where(dao.Course.Columns().Id, req.Id).Count()
	if err != nil {
		return err
	}
	if count == 0 {
		return gerror.New("课程不存在")
	}

	// 更新课程
	data := g.Map{
		dao.Course.Columns().Name:        req.Name,
		dao.Course.Columns().Type:        req.Type,
		dao.Course.Columns().Description: req.Description,
		dao.Course.Columns().CoverImage:  req.CoverImage,
		dao.Course.Columns().UpdateBy:    userId,
		dao.Course.Columns().UpdateTime:  gtime.Now(),
	}

	_, err = dao.Course.Ctx(ctx).Data(data).Where(dao.Course.Columns().Id, req.Id).Update()
	if err != nil {
		return err
	}

	// 更新部门关联
	if len(req.DeptIds) > 0 {
		// 先删除原有关联
		_, err = dao.CourseDept.Ctx(ctx).Where(dao.CourseDept.Columns().CourseId, req.Id).Delete()
		if err != nil {
			return err
		}

		// 添加新关联
		err = s.assignCourseToDepts(ctx, req.Id, req.DeptIds, userId)
		if err != nil {
			return err
		}
	}

	return nil
}

// Delete 删除课程
func (s *teachingCourseServiceImpl) Delete(ctx context.Context, req *vo.CourseDeleteReq, userId string) error {
	// 检查课程是否存在
	count, err := dao.Course.Ctx(ctx).Where(dao.Course.Columns().Id, req.Id).Count()
	if err != nil {
		return err
	}
	if count == 0 {
		return gerror.New("课程不存在")
	}

	// 检查是否有单元
	unitCount, err := dao.CourseUnit.Ctx(ctx).Where(dao.CourseUnit.Columns().CourseId, req.Id).Count()
	if err != nil {
		return err
	}
	if unitCount > 0 {
		return gerror.New("课程下有单元，不能删除")
	}

	// 删除部门关联
	_, err = dao.CourseDept.Ctx(ctx).Where(dao.CourseDept.Columns().CourseId, req.Id).Delete()
	if err != nil {
		return err
	}

	// 删除课程
	_, err = dao.Course.Ctx(ctx).Where(dao.Course.Columns().Id, req.Id).Delete()
	return err
}

// Detail 获取课程详情
func (s *teachingCourseServiceImpl) Detail(ctx context.Context, id string) (*vo.CourseDetailRes, error) {
	var course entity.TeachingCourse
	err := dao.Course.Ctx(ctx).Where(dao.Course.Columns().Id, id).Scan(&course)
	if err != nil {
		return nil, err
	}
	if course.Id == "" {
		return nil, gerror.New("课程不存在")
	}

	// 获取关联部门
	var depts []vo.DeptItem
	err = dao.CourseDept.Ctx(ctx).
		Fields("d.id, d.name").
		As("cd").
		LeftJoin("sys_depart d", "cd.dept_id = d.id").
		Where("cd.course_id", id).
		Scan(&depts)
	if err != nil {
		g.Log().Warning(ctx, "获取课程部门关联失败", err)
	}

	// 获取课程单元
	var units []vo.UnitItem
	err = dao.CourseUnit.Ctx(ctx).
		Where(dao.CourseUnit.Columns().CourseId, id).
		OrderAsc(dao.CourseUnit.Columns().SortOrder).
		Scan(&units)
	if err != nil {
		g.Log().Warning(ctx, "获取课程单元失败", err)
	}

	return &vo.CourseDetailRes{
		Id:          course.Id,
		Name:        course.Name,
		Type:        course.Type,
		Description: course.Description,
		CoverImage:  course.CoverImage,
		Status:      course.Status,
		CreateBy:    course.CreateBy,
		CreateTime:  course.CreateTime.Time.Format("2006-01-02 15:04:05"),
		UpdateBy:    course.UpdateBy,
		UpdateTime:  course.UpdateTime.Time.Format("2006-01-02 15:04:05"),
		Depts:       depts,
		Units:       units,
	}, nil
}

// Publish 发布/下架课程
func (s *teachingCourseServiceImpl) Publish(ctx context.Context, req *vo.CoursePublishReq, userId string) error {
	// 检查课程是否存在
	count, err := dao.Course.Ctx(ctx).Where(dao.Course.Columns().Id, req.Id).Count()
	if err != nil {
		return err
	}
	if count == 0 {
		return gerror.New("课程不存在")
	}

	// 更新状态
	data := g.Map{
		dao.Course.Columns().Status:     req.Status,
		dao.Course.Columns().UpdateBy:   userId,
		dao.Course.Columns().UpdateTime: gtime.Now(),
	}

	_, err = dao.Course.Ctx(ctx).Data(data).Where(dao.Course.Columns().Id, req.Id).Update()
	return err
}

// assignCourseToDepts 分配课程到部门
func (s *teachingCourseServiceImpl) assignCourseToDepts(ctx context.Context, courseId string, deptIds []string, userId string) error {
	var deptRelations []g.Map
	for _, deptId := range deptIds {
		deptRelations = append(deptRelations, g.Map{
			dao.CourseDept.Columns().Id:        guid.S(),
			dao.CourseDept.Columns().CourseId:  courseId,
			dao.CourseDept.Columns().DeptId:    deptId,
			dao.CourseDept.Columns().CreateBy:  userId,
			dao.CourseDept.Columns().CreateTime: gtime.Now(),
		})
	}

	if len(deptRelations) > 0 {
		_, err := dao.CourseDept.Ctx(ctx).Data(deptRelations).Insert()
		return err
	}

	return nil
}