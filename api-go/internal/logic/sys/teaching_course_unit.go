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
	service.RegisterTeachingCourseUnit(NewTeachingCourseUnitLogic())
}

type teachingCourseUnitLogic struct{}

func NewTeachingCourseUnitLogic() *teachingCourseUnitLogic {
	return &teachingCourseUnitLogic{}
}

func (l *teachingCourseUnitLogic) GetList(ctx context.Context, courseId, unitName string, pageNo, pageSize int) (list []*entity.TeachingCourseUnit, total int, err error) {
	model := dao.TeachingCourseUnit.Ctx(ctx).Where("del_flag", 0)
	if courseId != "" {
		model = model.Where("course_id", courseId)
	}
	if unitName != "" {
		model = model.WhereLike("unit_name", "%"+unitName+"%")
	}
	totalCount, err := model.Count()
	if err != nil {
		return nil, 0, err
	}
	total = totalCount
	offset := (pageNo - 1) * pageSize
	err = model.Order("order_num ASC, create_time DESC").Offset(offset).Limit(pageSize).Scan(&list)
	return
}

func (l *teachingCourseUnitLogic) GetByCourseId(ctx context.Context, courseId string) (list []*entity.TeachingCourseUnit, err error) {
	err = dao.TeachingCourseUnit.Ctx(ctx).Where("course_id", courseId).Where("del_flag", 0).Order("order_num ASC").Scan(&list)
	return
}

func (l *teachingCourseUnitLogic) Add(ctx context.Context, unit *entity.TeachingCourseUnit) (err error) {
	unit.Id = guid.S()
	unit.CreateTime = gtime.Now()
	unit.DelFlag = 0
	userId := ctx.Value(consts.CtxKeyUserId)
	if userId != nil {
		unit.CreateBy = userId.(string)
	}
	orgCode := ctx.Value(consts.CtxKeyOrgCode)
	if orgCode != nil {
		unit.SysOrgCode = orgCode.(string)
	}
	if unit.OrderNum == 0 {
		unit.OrderNum = 1
	}
	if unit.ShowCourseVideo == 0 {
		unit.ShowCourseVideo = 1
	}
	if unit.ShowCourseCase == 0 {
		unit.ShowCourseCase = 1
	}
	if unit.CourseVideoSource == 0 {
		unit.CourseVideoSource = 1
	}
	_, err = dao.TeachingCourseUnit.Ctx(ctx).Data(unit).Insert()
	return
}

func (l *teachingCourseUnitLogic) Edit(ctx context.Context, unit *entity.TeachingCourseUnit) (err error) {
	unit.UpdateTime = gtime.Now()
	userId := ctx.Value(consts.CtxKeyUserId)
	if userId != nil {
		unit.UpdateBy = userId.(string)
	}
	_, err = dao.TeachingCourseUnit.Ctx(ctx).Where("id", unit.Id).Where("del_flag", 0).Data(unit).Update()
	return
}

func (l *teachingCourseUnitLogic) Delete(ctx context.Context, id string) (err error) {
	updateBy := ""
	userId := ctx.Value(consts.CtxKeyUserId)
	if userId != nil {
		updateBy = userId.(string)
	}
	_, err = dao.TeachingCourseUnit.Ctx(ctx).Where("id", id).Data(map[string]interface{}{
		"del_flag": 1, "update_by": updateBy, "update_time": gtime.Now(),
	}).Update()
	return
}

func (l *teachingCourseUnitLogic) DeleteBatch(ctx context.Context, ids []string) (err error) {
	if len(ids) == 0 {
		return nil
	}
	updateBy := ""
	userId := ctx.Value(consts.CtxKeyUserId)
	if userId != nil {
		updateBy = userId.(string)
	}
	_, err = dao.TeachingCourseUnit.Ctx(ctx).WhereIn("id", ids).Data(map[string]interface{}{
		"del_flag": 1, "update_by": updateBy, "update_time": gtime.Now(),
	}).Update()
	return
}

func (l *teachingCourseUnitLogic) GetById(ctx context.Context, id string) (unit *entity.TeachingCourseUnit, err error) {
	err = dao.TeachingCourseUnit.Ctx(ctx).Where("id", id).Where("del_flag", 0).Scan(&unit)
	return
}

func (l *teachingCourseUnitLogic) Sort(ctx context.Context, unitIds []string) (err error) {
	updateBy := ""
	userId := ctx.Value(consts.CtxKeyUserId)
	if userId != nil {
		updateBy = userId.(string)
	}
	for i, id := range unitIds {
		_, err = dao.TeachingCourseUnit.Ctx(ctx).Where("id", id).Data(map[string]interface{}{
			"order_num": i + 1, "update_by": updateBy, "update_time": gtime.Now(),
		}).Update()
		if err != nil {
			return err
		}
	}
	return nil
}
