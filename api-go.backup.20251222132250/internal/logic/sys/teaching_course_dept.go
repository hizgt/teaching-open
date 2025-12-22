package sys

import (
	"context"

	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/guid"

	"teaching-open/internal/consts"
	"teaching-open/internal/dao"
	"teaching-open/internal/model/entity"
	"teaching-open/internal/service"
)

func init() {
	service.RegisterTeachingCourseDept(NewTeachingCourseDeptLogic())
}

type teachingCourseDeptLogic struct{}

func NewTeachingCourseDeptLogic() *teachingCourseDeptLogic {
	return &teachingCourseDeptLogic{}
}

func (l *teachingCourseDeptLogic) GetList(ctx context.Context, deptId, courseId string, pageNo, pageSize int) (list []*entity.TeachingCourseDept, total int, err error) {
	model := dao.TeachingCourseDept.Ctx(ctx)
	if deptId != "" {
		model = model.Where("dept_id", deptId)
	}
	if courseId != "" {
		model = model.Where("course_id", courseId)
	}
	totalCount, err := model.Count()
	if err != nil {
		return nil, 0, err
	}
	total = totalCount
	offset := (pageNo - 1) * pageSize
	err = model.Order("create_time DESC").Offset(offset).Limit(pageSize).Scan(&list)
	return
}

func (l *teachingCourseDeptLogic) GetByDeptId(ctx context.Context, deptId string) (list []*entity.TeachingCourseDept, err error) {
	err = dao.TeachingCourseDept.Ctx(ctx).Where("dept_id", deptId).Order("create_time DESC").Scan(&list)
	return
}

func (l *teachingCourseDeptLogic) GetByCourseId(ctx context.Context, courseId string) (list []*entity.TeachingCourseDept, err error) {
	err = dao.TeachingCourseDept.Ctx(ctx).Where("course_id", courseId).Order("create_time DESC").Scan(&list)
	return
}

func (l *teachingCourseDeptLogic) AddOrUpdate(ctx context.Context, courseDept *entity.TeachingCourseDept) (err error) {
	// 检查是否已存在
	var existing *entity.TeachingCourseDept
	err = dao.TeachingCourseDept.Ctx(ctx).
		Where("dept_id", courseDept.DeptId).
		Where("course_id", courseDept.CourseId).
		Scan(&existing)
	if err != nil {
		return err
	}

	userId := ctx.Value(consts.CtxKeyUserId)
	userIdStr := ""
	if userId != nil {
		userIdStr = userId.(string)
	}
	orgCode := ctx.Value(consts.CtxKeyOrgCode)
	orgCodeStr := ""
	if orgCode != nil {
		orgCodeStr = orgCode.(string)
	}

	if existing != nil {
		// 更新
		existing.OpenTime = courseDept.OpenTime
		existing.UpdateBy = userIdStr
		existing.UpdateTime = gtime.Now()
		_, err = dao.TeachingCourseDept.Ctx(ctx).Where("id", existing.Id).Data(existing).Update()
	} else {
		// 新增
		courseDept.Id = guid.S()
		courseDept.CreateBy = userIdStr
		courseDept.CreateTime = gtime.Now()
		courseDept.SysOrgCode = orgCodeStr
		_, err = dao.TeachingCourseDept.Ctx(ctx).Data(courseDept).Insert()
	}
	return
}

func (l *teachingCourseDeptLogic) Delete(ctx context.Context, id string) (err error) {
	_, err = dao.TeachingCourseDept.Ctx(ctx).Where("id", id).Delete()
	return
}

func (l *teachingCourseDeptLogic) BatchAdd(ctx context.Context, deptIds []string, courseId string) (err error) {
	userId := ctx.Value(consts.CtxKeyUserId)
	userIdStr := ""
	if userId != nil {
		userIdStr = userId.(string)
	}
	orgCode := ctx.Value(consts.CtxKeyOrgCode)
	orgCodeStr := ""
	if orgCode != nil {
		orgCodeStr = orgCode.(string)
	}

	for _, deptId := range deptIds {
		// 检查是否已存在
		count, err := dao.TeachingCourseDept.Ctx(ctx).
			Where("dept_id", deptId).
			Where("course_id", courseId).
			Count()
		if err != nil {
			return err
		}
		if count > 0 {
			continue // 已存在，跳过
		}

		courseDept := &entity.TeachingCourseDept{
			Id:         guid.S(),
			DeptId:     deptId,
			CourseId:   courseId,
			CreateBy:   userIdStr,
			CreateTime: gtime.Now(),
			SysOrgCode: orgCodeStr,
		}
		_, err = dao.TeachingCourseDept.Ctx(ctx).Data(courseDept).Insert()
		if err != nil {
			return err
		}
	}
	return nil
}
