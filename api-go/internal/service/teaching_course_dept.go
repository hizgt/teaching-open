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

// TeachingCourseDeptService 课程部门关联服务接口
type TeachingCourseDeptService interface {
	List(ctx context.Context, req *vo.CourseDeptListReq) (*vo.CourseDeptListRes, error)
	Assign(ctx context.Context, req *vo.CourseDeptAssignReq, userId string) error
	Remove(ctx context.Context, req *vo.CourseDeptRemoveReq, userId string) error
	QueryByDeptId(ctx context.Context, deptId string) ([]vo.CourseItem, error)
	QueryByCourseId(ctx context.Context, courseId string) ([]vo.DeptItem, error)
	BatchAdd(ctx context.Context, req *vo.CourseDeptBatchAddReq, userId string) error
}

// teachingCourseDeptServiceImpl 课程部门关联服务实现
type teachingCourseDeptServiceImpl struct{}

// NewTeachingCourseDeptService 创建课程部门关联服务实例
func NewTeachingCourseDeptService() TeachingCourseDeptService {
	return &teachingCourseDeptServiceImpl{}
}

// List 获取课程部门关联列表
func (s *teachingCourseDeptServiceImpl) List(ctx context.Context, req *vo.CourseDeptListReq) (*vo.CourseDeptListRes, error) {
	model := dao.CourseDept.Ctx(ctx)

	// 构建查询条件
	if req.CourseId != "" {
		model = model.Where(dao.CourseDept.Columns().CourseId, req.CourseId)
	}
	if req.DeptId != "" {
		model = model.Where(dao.CourseDept.Columns().DeptId, req.DeptId)
	}

	var relations []entity.TeachingCourseDept
	err := model.Scan(&relations)
	if err != nil {
		return nil, err
	}

	// 转换为VO，包含课程和部门信息
	var records []vo.CourseDeptItem
	for _, relation := range relations {
		// 获取课程信息
		var course entity.TeachingCourse
		err = dao.Course.Ctx(ctx).Where(dao.Course.Columns().Id, relation.CourseId).Scan(&course)
		if err != nil {
			continue
		}

		// 获取部门信息（这里假设有sys_depart表）
		deptName := "未知部门"
		// 这里应该查询部门表获取部门名称
		// deptResult, err := dao.SysDepart.Ctx(ctx).Where("id", relation.DeptId).Value("name")
		// if err == nil {
		//     deptName = deptResult.String()
		// }

		records = append(records, vo.CourseDeptItem{
			Id:         relation.Id,
			CourseId:   relation.CourseId,
			CourseName: course.Name,
			DeptId:     relation.DeptId,
			DeptName:   deptName,
			CreateBy:   relation.CreateBy,
			CreateTime: relation.CreateTime.Time.Format("2006-01-02 15:04:05"),
		})
	}

	return &vo.CourseDeptListRes{Records: records}, nil
}

// Assign 批量分配课程到部门
func (s *teachingCourseDeptServiceImpl) Assign(ctx context.Context, req *vo.CourseDeptAssignReq, userId string) error {
	// 检查课程是否存在
	courseCount, err := dao.Course.Ctx(ctx).Where(dao.Course.Columns().Id, req.CourseId).Count()
	if err != nil {
		return err
	}
	if courseCount == 0 {
		return gerror.New("课程不存在")
	}

	// 检查课程是否已发布
	var course entity.TeachingCourse
	err = dao.Course.Ctx(ctx).Where(dao.Course.Columns().Id, req.CourseId).Scan(&course)
	if err != nil {
		return err
	}
	if course.Status != "published" {
		return gerror.New("只能为已发布的课程分配部门")
	}

	// 删除原有关联
	_, err = dao.CourseDept.Ctx(ctx).Where(dao.CourseDept.Columns().CourseId, req.CourseId).Delete()
	if err != nil {
		return err
	}

	// 添加新关联
	if len(req.DeptIds) > 0 {
		var relations []g.Map
		for _, deptId := range req.DeptIds {
			relations = append(relations, g.Map{
				dao.CourseDept.Columns().Id:         guid.S(),
				dao.CourseDept.Columns().CourseId:   req.CourseId,
				dao.CourseDept.Columns().DeptId:     deptId,
				dao.CourseDept.Columns().CreateBy:   userId,
				dao.CourseDept.Columns().CreateTime: gtime.Now(),
			})
		}

		_, err = dao.CourseDept.Ctx(ctx).Data(relations).Insert()
		return err
	}

	return nil
}

// Remove 移除课程部门关联
func (s *teachingCourseDeptServiceImpl) Remove(ctx context.Context, req *vo.CourseDeptRemoveReq, userId string) error {
	// 检查关联是否存在
	count, err := dao.CourseDept.Ctx(ctx).Where(dao.CourseDept.Columns().Id, req.Id).Count()
	if err != nil {
		return err
	}
	if count == 0 {
		return gerror.New("课程部门关联不存在")
	}

	// 删除关联
	_, err = dao.CourseDept.Ctx(ctx).Where(dao.CourseDept.Columns().Id, req.Id).Delete()
	return err
}

// QueryByDeptId 根据部门ID查询课程
func (s *teachingCourseDeptServiceImpl) QueryByDeptId(ctx context.Context, deptId string) ([]vo.CourseItem, error) {
	var relations []entity.TeachingCourseDept
	err := dao.CourseDept.Ctx(ctx).Where(dao.CourseDept.Columns().DeptId, deptId).Scan(&relations)
	if err != nil {
		return nil, err
	}

	var records []vo.CourseItem
	for _, relation := range relations {
		var course entity.TeachingCourse
		err = dao.Course.Ctx(ctx).Where(dao.Course.Columns().Id, relation.CourseId).Scan(&course)
		if err != nil || course.Id == "" {
			continue
		}

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

	return records, nil
}

// QueryByCourseId 根据课程ID查询部门
func (s *teachingCourseDeptServiceImpl) QueryByCourseId(ctx context.Context, courseId string) ([]vo.DeptItem, error) {
	var depts []vo.DeptItem
	err := g.DB().Ctx(ctx).Raw(`
		SELECT d.id, d.depart_name as name 
		FROM teaching_course_dept cd 
		LEFT JOIN sys_depart d ON cd.dept_id = d.id 
		WHERE cd.course_id = ? AND d.del_flag = '0'
	`, courseId).Scan(&depts)
	if err != nil {
		return nil, err
	}

	return depts, nil
}

// BatchAdd 批量添加课程部门关联
func (s *teachingCourseDeptServiceImpl) BatchAdd(ctx context.Context, req *vo.CourseDeptBatchAddReq, userId string) error {
	if len(req.CourseIds) == 0 || len(req.DeptIds) == 0 {
		return gerror.New("课程ID和部门ID不能为空")
	}

	var relations []g.Map
	for _, courseId := range req.CourseIds {
		for _, deptId := range req.DeptIds {
			// 检查是否已存在
			count, err := dao.CourseDept.Ctx(ctx).
				Where(dao.CourseDept.Columns().CourseId, courseId).
				Where(dao.CourseDept.Columns().DeptId, deptId).
				Count()
			if err != nil {
				return err
			}
			if count > 0 {
				continue
			}

			relations = append(relations, g.Map{
				dao.CourseDept.Columns().Id:         guid.S(),
				dao.CourseDept.Columns().CourseId:   courseId,
				dao.CourseDept.Columns().DeptId:     deptId,
				dao.CourseDept.Columns().CreateBy:   userId,
				dao.CourseDept.Columns().CreateTime: gtime.Now(),
			})
		}
	}

	if len(relations) > 0 {
		_, err := dao.CourseDept.Ctx(ctx).Data(relations).Insert()
		return err
	}

	return nil
}
