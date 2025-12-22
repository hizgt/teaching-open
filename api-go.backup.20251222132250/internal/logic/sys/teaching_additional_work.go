package sys

import (
	"context"
	"strings"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"

	v1 "teaching-open/api/v1/sys"
	"teaching-open/internal/dao"
	"teaching-open/internal/model/entity"
	"teaching-open/internal/service"
	"teaching-open/utility/jwt"
)

func init() {
	service.RegisterTeachingAdditionalWork(NewTeachingAdditionalWorkLogic())
}

type sTeachingAdditionalWork struct{}

func NewTeachingAdditionalWorkLogic() *sTeachingAdditionalWork {
	return &sTeachingAdditionalWork{}
}

// List 获取附加作业列表
func (s *sTeachingAdditionalWork) List(ctx context.Context, req *v1.AdditionalWorkListReq) (list interface{}, total int, err error) {
	m := dao.TeachingAdditionalWork.Ctx(ctx)

	// 作业名称查询
	if req.WorkName != "" {
		m = m.WhereLike("work_name", "%"+req.WorkName+"%")
	}
	// 代码类型查询
	if req.CodeType != "" {
		m = m.Where("code_type", req.CodeType)
	}
	// 状态查询
	if req.Status != nil {
		m = m.Where("status", req.Status)
	}
	// 班级查询
	if req.WorkDept != "" {
		m = m.WhereLike("work_dept", "%"+req.WorkDept+"%")
	}

	// 统计总数
	count, err := m.Count()
	if err != nil {
		return nil, 0, err
	}
	total = count

	// 分页查询
	pageNo := req.PageNo
	pageSize := req.PageSize
	if pageNo <= 0 {
		pageNo = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}

	var entities []*entity.TeachingAdditionalWork
	err = m.Order("create_time DESC").Page(pageNo, pageSize).Scan(&entities)
	if err != nil {
		return nil, 0, err
	}

	// 转换为响应格式
	var workInfos []*v1.AdditionalWorkInfo
	for _, e := range entities {
		workInfos = append(workInfos, s.entityToInfo(e))
	}

	return workInfos, total, nil
}

// ListByDept 按班级获取附加作业列表
func (s *sTeachingAdditionalWork) ListByDept(ctx context.Context, req *v1.AdditionalWorkByDeptReq) (list interface{}, total int, err error) {
	m := dao.TeachingAdditionalWork.Ctx(ctx).Where("status", 1) // 只查已发布

	// 班级查询
	if req.WorkDept != "" {
		m = m.WhereLike("work_dept", "%"+req.WorkDept+"%")
	}
	// 代码类型查询
	if req.CodeType != "" {
		m = m.Where("code_type", req.CodeType)
	}

	// 统计总数
	count, err := m.Count()
	if err != nil {
		return nil, 0, err
	}
	total = count

	// 分页查询
	pageNo := req.PageNo
	pageSize := req.PageSize
	if pageNo <= 0 {
		pageNo = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}

	var entities []*entity.TeachingAdditionalWork
	err = m.Order("create_time DESC").Page(pageNo, pageSize).Scan(&entities)
	if err != nil {
		return nil, 0, err
	}

	// 转换为响应格式
	var workInfos []*v1.AdditionalWorkInfo
	for _, e := range entities {
		workInfos = append(workInfos, s.entityToInfo(e))
	}

	return workInfos, total, nil
}

// Add 添加附加作业
func (s *sTeachingAdditionalWork) Add(ctx context.Context, req *v1.AdditionalWorkAddReq) (id string, err error) {
	username := jwt.GetUsername(ctx)
	newId := g.NewVar(nil).String()
	if newId == "" {
		newId = gconv.String(gtime.TimestampNano())
	}

	data := &entity.TeachingAdditionalWork{
		Id:              newId,
		CodeType:        req.CodeType,
		WorkName:        req.WorkName,
		WorkDesc:        req.WorkDesc,
		WorkCover:       req.WorkCover,
		WorkUrl:         req.WorkUrl,
		WorkDept:        req.WorkDept,
		WorkDocumentUrl: req.WorkDocumentUrl,
		Status:          0, // 默认草稿
		CreateBy:        username,
		CreateTime:      gtime.Now(),
		UpdateBy:        username,
		UpdateTime:      gtime.Now(),
	}

	_, err = dao.TeachingAdditionalWork.Ctx(ctx).Insert(data)
	if err != nil {
		return "", err
	}

	return newId, nil
}

// Edit 编辑附加作业
func (s *sTeachingAdditionalWork) Edit(ctx context.Context, req *v1.AdditionalWorkEditReq) error {
	username := jwt.GetUsername(ctx)

	data := g.Map{
		"code_type":         req.CodeType,
		"work_name":         req.WorkName,
		"work_desc":         req.WorkDesc,
		"work_cover":        req.WorkCover,
		"work_url":          req.WorkUrl,
		"work_dept":         req.WorkDept,
		"work_document_url": req.WorkDocumentUrl,
		"update_by":         username,
		"update_time":       gtime.Now(),
	}

	if req.Status != nil {
		data["status"] = req.Status
	}

	_, err := dao.TeachingAdditionalWork.Ctx(ctx).Where("id", req.Id).Data(data).Update()
	return err
}

// Delete 删除附加作业
func (s *sTeachingAdditionalWork) Delete(ctx context.Context, id string) error {
	_, err := dao.TeachingAdditionalWork.Ctx(ctx).Where("id", id).Delete()
	return err
}

// DeleteBatch 批量删除附加作业
func (s *sTeachingAdditionalWork) DeleteBatch(ctx context.Context, ids string) error {
	idList := strings.Split(ids, ",")
	_, err := dao.TeachingAdditionalWork.Ctx(ctx).WhereIn("id", idList).Delete()
	return err
}

// GetById 获取附加作业详情
func (s *sTeachingAdditionalWork) GetById(ctx context.Context, id string) (*v1.AdditionalWorkInfo, error) {
	var e *entity.TeachingAdditionalWork
	err := dao.TeachingAdditionalWork.Ctx(ctx).Where("id", id).Scan(&e)
	if err != nil {
		return nil, err
	}
	if e == nil {
		return nil, nil
	}

	return s.entityToInfo(e), nil
}

// Publish 发布附加作业
func (s *sTeachingAdditionalWork) Publish(ctx context.Context, id string) error {
	username := jwt.GetUsername(ctx)

	_, err := dao.TeachingAdditionalWork.Ctx(ctx).Where("id", id).Data(g.Map{
		"status":      1,
		"update_by":   username,
		"update_time": gtime.Now(),
	}).Update()
	return err
}

// Offline 下架附加作业
func (s *sTeachingAdditionalWork) Offline(ctx context.Context, id string) error {
	username := jwt.GetUsername(ctx)

	_, err := dao.TeachingAdditionalWork.Ctx(ctx).Where("id", id).Data(g.Map{
		"status":      0,
		"update_by":   username,
		"update_time": gtime.Now(),
	}).Update()
	return err
}

// entityToInfo 实体转换为响应格式
func (s *sTeachingAdditionalWork) entityToInfo(e *entity.TeachingAdditionalWork) *v1.AdditionalWorkInfo {
	return &v1.AdditionalWorkInfo{
		Id:              e.Id,
		CodeType:        e.CodeType,
		WorkName:        e.WorkName,
		WorkDesc:        e.WorkDesc,
		WorkCover:       e.WorkCover,
		WorkUrl:         e.WorkUrl,
		WorkDept:        e.WorkDept,
		WorkDocumentUrl: e.WorkDocumentUrl,
		Status:          e.Status,
		SysOrgCode:      e.SysOrgCode,
		CreateBy:        e.CreateBy,
		CreateTime:      e.CreateTime,
		UpdateBy:        e.UpdateBy,
		UpdateTime:      e.UpdateTime,
	}
}
