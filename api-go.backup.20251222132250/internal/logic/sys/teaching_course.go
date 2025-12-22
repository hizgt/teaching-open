// Package sys provides logic implementations for teaching management.
package sys

import (
	"context"
	"strings"

	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/guid"

	"teaching-open/internal/consts"
	"teaching-open/internal/dao"
	"teaching-open/internal/model/entity"
	"teaching-open/internal/service"
)

func init() {
	service.RegisterTeachingCourse(NewTeachingCourseLogic())
}

// teachingCourseLogic 课程管理逻辑实现
type teachingCourseLogic struct{}

// NewTeachingCourseLogic 创建课程管理逻辑实例
func NewTeachingCourseLogic() *teachingCourseLogic {
	return &teachingCourseLogic{}
}

// GetList 获取课程列表
func (l *teachingCourseLogic) GetList(ctx context.Context, courseName, courseType, courseCategory string, isShared, showHome, pageNo, pageSize int) (list []*entity.TeachingCourse, total int, err error) {
	model := dao.TeachingCourse.Ctx(ctx).Where("del_flag", 0)

	// 课程名模糊查询
	if courseName != "" {
		model = model.WhereLike("course_name", "%"+courseName+"%")
	}

	// 课程类型筛选
	if courseType != "" {
		model = model.Where("course_type", courseType)
	}

	// 课程分类筛选
	if courseCategory != "" {
		model = model.Where("course_category", courseCategory)
	}

	// 共享状态筛选
	if isShared >= 0 {
		model = model.Where("is_shared", isShared)
	}

	// 首页展示筛选
	if showHome >= 0 {
		model = model.Where("show_home", showHome)
	}

	// 查询总数
	totalCount, err := model.Count()
	if err != nil {
		return nil, 0, err
	}
	total = totalCount

	// 分页查询
	offset := (pageNo - 1) * pageSize
	err = model.Order("order_num ASC, create_time DESC").Offset(offset).Limit(pageSize).Scan(&list)
	if err != nil {
		return nil, 0, err
	}

	return list, total, nil
}

// GetHomeCourse 获取首页展示课程
func (l *teachingCourseLogic) GetHomeCourse(ctx context.Context) (list []*entity.TeachingCourse, err error) {
	err = dao.TeachingCourse.Ctx(ctx).
		Where("del_flag", 0).
		Where("show_home", 1).
		Order("order_num ASC, create_time DESC").
		Scan(&list)
	return
}

// Add 添加课程
func (l *teachingCourseLogic) Add(ctx context.Context, course *entity.TeachingCourse) (err error) {
	// 生成ID
	course.Id = guid.S()

	// 设置创建信息
	course.CreateTime = gtime.Now()
	course.DelFlag = 0

	// 从上下文获取当前用户
	userId := ctx.Value(consts.CtxKeyUserId)
	if userId != nil {
		course.CreateBy = userId.(string)
	}
	orgCode := ctx.Value(consts.CtxKeyOrgCode)
	if orgCode != nil {
		course.SysOrgCode = orgCode.(string)
	}

	// 设置默认值
	if course.OrderNum == 0 {
		course.OrderNum = 1
	}

	_, err = dao.TeachingCourse.Ctx(ctx).Data(course).Insert()
	return
}

// Edit 编辑课程
func (l *teachingCourseLogic) Edit(ctx context.Context, course *entity.TeachingCourse) (err error) {
	// 设置更新信息
	course.UpdateTime = gtime.Now()

	// 从上下文获取当前用户
	userId := ctx.Value(consts.CtxKeyUserId)
	if userId != nil {
		course.UpdateBy = userId.(string)
	}

	// 更新
	_, err = dao.TeachingCourse.Ctx(ctx).
		Where("id", course.Id).
		Where("del_flag", 0).
		Data(course).
		Update()
	return
}

// Delete 删除课程（软删除）
func (l *teachingCourseLogic) Delete(ctx context.Context, id string) (err error) {
	updateBy := ""
	userId := ctx.Value(consts.CtxKeyUserId)
	if userId != nil {
		updateBy = userId.(string)
	}

	_, err = dao.TeachingCourse.Ctx(ctx).
		Where("id", id).
		Data(map[string]interface{}{
			"del_flag":    1,
			"update_by":   updateBy,
			"update_time": gtime.Now(),
		}).
		Update()
	return
}

// DeleteBatch 批量删除课程（软删除）
func (l *teachingCourseLogic) DeleteBatch(ctx context.Context, ids []string) (err error) {
	if len(ids) == 0 {
		return nil
	}

	updateBy := ""
	userId := ctx.Value(consts.CtxKeyUserId)
	if userId != nil {
		updateBy = userId.(string)
	}

	_, err = dao.TeachingCourse.Ctx(ctx).
		WhereIn("id", ids).
		Data(map[string]interface{}{
			"del_flag":    1,
			"update_by":   updateBy,
			"update_time": gtime.Now(),
		}).
		Update()
	return
}

// GetById 根据ID获取课程详情
func (l *teachingCourseLogic) GetById(ctx context.Context, id string) (course *entity.TeachingCourse, err error) {
	err = dao.TeachingCourse.Ctx(ctx).
		Where("id", id).
		Where("del_flag", 0).
		Scan(&course)
	return
}

// Publish 发布/下架课程
func (l *teachingCourseLogic) Publish(ctx context.Context, id string, showHome int) (err error) {
	updateBy := ""
	userId := ctx.Value(consts.CtxKeyUserId)
	if userId != nil {
		updateBy = userId.(string)
	}

	_, err = dao.TeachingCourse.Ctx(ctx).
		Where("id", id).
		Where("del_flag", 0).
		Data(map[string]interface{}{
			"show_home":   showHome,
			"update_by":   updateBy,
			"update_time": gtime.Now(),
		}).
		Update()
	return
}

// SetShared 设置共享状态
func (l *teachingCourseLogic) SetShared(ctx context.Context, id string, isShared int) (err error) {
	updateBy := ""
	userId := ctx.Value(consts.CtxKeyUserId)
	if userId != nil {
		updateBy = userId.(string)
	}

	_, err = dao.TeachingCourse.Ctx(ctx).
		Where("id", id).
		Where("del_flag", 0).
		Data(map[string]interface{}{
			"is_shared":   isShared,
			"update_by":   updateBy,
			"update_time": gtime.Now(),
		}).
		Update()
	return
}

// AuthorizeDept 授权部门
func (l *teachingCourseLogic) AuthorizeDept(ctx context.Context, id string, departIds string) (err error) {
	updateBy := ""
	userId := ctx.Value(consts.CtxKeyUserId)
	if userId != nil {
		updateBy = userId.(string)
	}

	// 处理部门ID，确保格式正确
	departIds = strings.TrimSpace(departIds)

	_, err = dao.TeachingCourse.Ctx(ctx).
		Where("id", id).
		Where("del_flag", 0).
		Data(map[string]interface{}{
			"depart_ids":  departIds,
			"update_by":   updateBy,
			"update_time": gtime.Now(),
		}).
		Update()
	return
}
