package sys

import (
	"context"
	"strings"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/guid"

	v1 "teaching-open/api/v1/sys"
	"teaching-open/internal/consts"
	"teaching-open/internal/dao"
	"teaching-open/internal/model/do"
	"teaching-open/internal/service"
)

type sTeachingWork struct{}

func init() {
	service.RegisterTeachingWork(NewTeachingWork())
}

func NewTeachingWork() *sTeachingWork {
	return &sTeachingWork{}
}

// List 作品列表
func (s *sTeachingWork) List(ctx context.Context, req *v1.WorkListReq) (list interface{}, total int, err error) {
	m := dao.TeachingWork.Ctx(ctx).Where("del_flag", 0)

	// 组织机构过滤
	orgCode := gconv.String(ctx.Value(consts.CtxKeyOrgCode))
	if orgCode != "" {
		m = m.WhereLike("sys_org_code", orgCode+"%")
	}

	// 条件查询
	if req.WorkName != "" {
		m = m.WhereLike("work_name", "%"+req.WorkName+"%")
	}
	if req.WorkType != "" {
		m = m.Where("work_type", req.WorkType)
	}
	if req.WorkStatus >= 0 {
		m = m.Where("work_status", req.WorkStatus)
	}
	if req.UserId != "" {
		m = m.Where("user_id", req.UserId)
	}
	if req.CourseId != "" {
		m = m.Where("course_id", req.CourseId)
	}
	if req.DepartId != "" {
		m = m.Where("depart_id", req.DepartId)
	}
	if req.WorkScene != "" {
		m = m.Where("work_scene", req.WorkScene)
	}

	total, err = m.Count()
	if err != nil {
		return nil, 0, gerror.Wrap(err, "查询作品总数失败")
	}

	var works []map[string]interface{}
	err = m.OrderDesc("create_time").
		Page(req.PageNo, req.PageSize).
		Scan(&works)
	if err != nil {
		return nil, 0, gerror.Wrap(err, "查询作品列表失败")
	}

	// 填充用户信息
	works = s.fillUserInfo(ctx, works)

	return works, total, nil
}

// Mine 我的作品
func (s *sTeachingWork) Mine(ctx context.Context, req *v1.WorkMineReq) (list interface{}, total int, err error) {
	userId := gconv.String(ctx.Value(consts.CtxKeyUserId))
	if userId == "" {
		return nil, 0, gerror.New("用户未登录")
	}

	m := dao.TeachingWork.Ctx(ctx).
		Where("user_id", userId).
		Where("del_flag", 0)

	// 条件查询
	if req.WorkName != "" {
		m = m.WhereLike("work_name", "%"+req.WorkName+"%")
	}
	if req.WorkType != "" {
		m = m.Where("work_type", req.WorkType)
	}
	if req.WorkStatus >= 0 {
		m = m.Where("work_status", req.WorkStatus)
	}
	if req.WorkScene != "" {
		m = m.Where("work_scene", req.WorkScene)
	}

	total, err = m.Count()
	if err != nil {
		return nil, 0, gerror.Wrap(err, "查询作品总数失败")
	}

	var works []map[string]interface{}
	err = m.OrderDesc("create_time").
		Page(req.PageNo, req.PageSize).
		Scan(&works)
	if err != nil {
		return nil, 0, gerror.Wrap(err, "查询作品列表失败")
	}

	return works, total, nil
}

// GreatWork 优秀作品
func (s *sTeachingWork) GreatWork(ctx context.Context, req *v1.WorkGreatReq) (list interface{}, total int, err error) {
	m := dao.TeachingWork.Ctx(ctx).
		Where("del_flag", 0).
		Where("work_status", 2) // 已批改/优秀

	if req.WorkType != "" {
		m = m.Where("work_type", req.WorkType)
	}

	total, err = m.Count()
	if err != nil {
		return nil, 0, gerror.Wrap(err, "查询优秀作品总数失败")
	}

	var works []map[string]interface{}
	err = m.OrderDesc("star_num").
		OrderDesc("create_time").
		Page(req.PageNo, req.PageSize).
		Scan(&works)
	if err != nil {
		return nil, 0, gerror.Wrap(err, "查询优秀作品失败")
	}

	works = s.fillUserInfo(ctx, works)

	return works, total, nil
}

// StarWork 收藏作品（点赞数高的作品）
func (s *sTeachingWork) StarWork(ctx context.Context, req *v1.WorkStarReq) (list interface{}, total int, err error) {
	m := dao.TeachingWork.Ctx(ctx).
		Where("del_flag", 0).
		Where("star_num >", 0)

	if req.WorkType != "" {
		m = m.Where("work_type", req.WorkType)
	}

	total, err = m.Count()
	if err != nil {
		return nil, 0, gerror.Wrap(err, "查询收藏作品总数失败")
	}

	var works []map[string]interface{}
	err = m.OrderDesc("star_num").
		OrderDesc("create_time").
		Page(req.PageNo, req.PageSize).
		Scan(&works)
	if err != nil {
		return nil, 0, gerror.Wrap(err, "查询收藏作品失败")
	}

	works = s.fillUserInfo(ctx, works)

	return works, total, nil
}

// Leaderboard 作品排行榜
func (s *sTeachingWork) Leaderboard(ctx context.Context, req *v1.WorkLeaderboardReq) (list interface{}, total int, err error) {
	m := dao.TeachingWork.Ctx(ctx).Where("del_flag", 0)

	if req.WorkType != "" {
		m = m.Where("work_type", req.WorkType)
	}

	total, err = m.Count()
	if err != nil {
		return nil, 0, gerror.Wrap(err, "查询排行榜总数失败")
	}

	// 排序方式
	orderField := "star_num"
	if req.SortBy == "view_num" {
		orderField = "view_num"
	} else if req.SortBy == "collect_num" {
		orderField = "collect_num"
	}

	var works []map[string]interface{}
	err = m.OrderDesc(orderField).
		OrderDesc("create_time").
		Page(req.PageNo, req.PageSize).
		Scan(&works)
	if err != nil {
		return nil, 0, gerror.Wrap(err, "查询排行榜失败")
	}

	works = s.fillUserInfo(ctx, works)

	return works, total, nil
}

// Add 添加作品
func (s *sTeachingWork) Add(ctx context.Context, req *v1.WorkAddReq) (id string, err error) {
	userId := gconv.String(ctx.Value(consts.CtxKeyUserId))
	username := gconv.String(ctx.Value(consts.CtxKeyUsername))
	orgCode := gconv.String(ctx.Value(consts.CtxKeyOrgCode))

	id = guid.S()
	_, err = dao.TeachingWork.Ctx(ctx).Insert(do.TeachingWork{
		Id:           id,
		CreateBy:     username,
		CreateTime:   gtime.Now(),
		SysOrgCode:   orgCode,
		UserId:       userId,
		DepartId:     req.DepartId,
		CourseId:     req.CourseId,
		WorkName:     req.WorkName,
		WorkType:     req.WorkType,
		WorkFile:     req.WorkFile,
		WorkCover:    req.WorkCover,
		WorkStatus:   0, // 草稿
		StarNum:      0,
		CollectNum:   0,
		DelFlag:      0,
		ViewNum:      0,
		AdditionalId: req.AdditionalId,
		WorkScene:    req.WorkScene,
		HasCloudData: req.HasCloudData,
	})
	if err != nil {
		return "", gerror.Wrap(err, "添加作品失败")
	}

	return id, nil
}

// Edit 编辑作品
func (s *sTeachingWork) Edit(ctx context.Context, req *v1.WorkEditReq) error {
	username := gconv.String(ctx.Value(consts.CtxKeyUsername))

	updateData := g.Map{
		"update_by":   username,
		"update_time": gtime.Now(),
	}

	if req.WorkName != "" {
		updateData["work_name"] = req.WorkName
	}
	if req.WorkType != "" {
		updateData["work_type"] = req.WorkType
	}
	if req.WorkFile != "" {
		updateData["work_file"] = req.WorkFile
	}
	if req.WorkCover != "" {
		updateData["work_cover"] = req.WorkCover
	}
	if req.WorkStatus >= 0 {
		updateData["work_status"] = req.WorkStatus
	}
	updateData["has_cloud_data"] = req.HasCloudData

	_, err := dao.TeachingWork.Ctx(ctx).
		Where("id", req.Id).
		Data(updateData).
		Update()
	if err != nil {
		return gerror.Wrap(err, "编辑作品失败")
	}

	return nil
}

// Delete 删除作品（软删除）
func (s *sTeachingWork) Delete(ctx context.Context, id string) error {
	_, err := dao.TeachingWork.Ctx(ctx).
		Where("id", id).
		Data(g.Map{"del_flag": 1}).
		Update()
	if err != nil {
		return gerror.Wrap(err, "删除作品失败")
	}

	return nil
}

// DeleteBatch 批量删除作品
func (s *sTeachingWork) DeleteBatch(ctx context.Context, ids string) error {
	idList := strings.Split(ids, ",")
	_, err := dao.TeachingWork.Ctx(ctx).
		WhereIn("id", idList).
		Data(g.Map{"del_flag": 1}).
		Update()
	if err != nil {
		return gerror.Wrap(err, "批量删除作品失败")
	}

	return nil
}

// GetById 作品详情
func (s *sTeachingWork) GetById(ctx context.Context, id string) (*v1.WorkInfo, error) {
	var work *v1.WorkInfo
	err := dao.TeachingWork.Ctx(ctx).
		Where("id", id).
		Where("del_flag", 0).
		Scan(&work)
	if err != nil {
		return nil, gerror.Wrap(err, "查询作品失败")
	}
	if work == nil {
		return nil, gerror.New("作品不存在")
	}

	// 增加查看次数
	_, _ = dao.TeachingWork.Ctx(ctx).
		Where("id", id).
		Increment("view_num", 1)

	// 填充用户信息
	s.fillSingleWorkInfo(ctx, work)

	return work, nil
}

// Submit 提交作品
func (s *sTeachingWork) Submit(ctx context.Context, id string) error {
	username := gconv.String(ctx.Value(consts.CtxKeyUsername))

	_, err := dao.TeachingWork.Ctx(ctx).
		Where("id", id).
		Data(g.Map{
			"work_status": 1, // 已提交
			"update_by":   username,
			"update_time": gtime.Now(),
		}).
		Update()
	if err != nil {
		return gerror.Wrap(err, "提交作品失败")
	}

	return nil
}

// StudentWorkInfo 学生作品信息
func (s *sTeachingWork) StudentWorkInfo(ctx context.Context, req *v1.WorkStudentInfoReq) (workCount int, works interface{}, err error) {
	m := dao.TeachingWork.Ctx(ctx).Where("del_flag", 0)

	if req.UserId != "" {
		m = m.Where("user_id", req.UserId)
	}
	if req.CourseId != "" {
		m = m.Where("course_id", req.CourseId)
	}

	workCount, err = m.Count()
	if err != nil {
		return 0, nil, gerror.Wrap(err, "查询学生作品数量失败")
	}

	var workList []map[string]interface{}
	err = m.OrderDesc("create_time").Limit(10).Scan(&workList)
	if err != nil {
		return 0, nil, gerror.Wrap(err, "查询学生作品失败")
	}

	return workCount, workList, nil
}

// SendWork 发送作品给其他用户
func (s *sTeachingWork) SendWork(ctx context.Context, req *v1.WorkSendReq) error {
	// 查询原作品
	var work map[string]interface{}
	err := dao.TeachingWork.Ctx(ctx).Where("id", req.WorkId).Scan(&work)
	if err != nil {
		return gerror.Wrap(err, "查询作品失败")
	}
	if work == nil {
		return gerror.New("作品不存在")
	}

	username := gconv.String(ctx.Value(consts.CtxKeyUsername))

	// 复制作品给目标用户
	newId := guid.S()
	_, err = dao.TeachingWork.Ctx(ctx).Insert(do.TeachingWork{
		Id:           newId,
		CreateBy:     username,
		CreateTime:   gtime.Now(),
		SysOrgCode:   gconv.String(work["sys_org_code"]),
		UserId:       req.UserId,
		DepartId:     gconv.String(work["depart_id"]),
		CourseId:     gconv.String(work["course_id"]),
		WorkName:     gconv.String(work["work_name"]),
		WorkType:     gconv.String(work["work_type"]),
		WorkFile:     gconv.String(work["work_file"]),
		WorkCover:    gconv.String(work["work_cover"]),
		WorkStatus:   0,
		StarNum:      0,
		CollectNum:   0,
		DelFlag:      0,
		ViewNum:      0,
		AdditionalId: gconv.String(work["additional_id"]),
		WorkScene:    gconv.String(work["work_scene"]),
		HasCloudData: gconv.Int(work["has_cloud_data"]),
	})
	if err != nil {
		return gerror.Wrap(err, "发送作品失败")
	}

	return nil
}

// MineAdditional 我的附加作业作品
func (s *sTeachingWork) MineAdditional(ctx context.Context, req *v1.WorkMineAdditionalReq) (list interface{}, total int, err error) {
	userId := gconv.String(ctx.Value(consts.CtxKeyUserId))
	if userId == "" {
		return nil, 0, gerror.New("用户未登录")
	}

	m := dao.TeachingWork.Ctx(ctx).
		Where("user_id", userId).
		Where("del_flag", 0).
		Where("work_scene", "additional")

	if req.AdditionalId != "" {
		m = m.Where("additional_id", req.AdditionalId)
	}

	total, err = m.Count()
	if err != nil {
		return nil, 0, gerror.Wrap(err, "查询附加作业作品总数失败")
	}

	var works []map[string]interface{}
	err = m.OrderDesc("create_time").
		Page(req.PageNo, req.PageSize).
		Scan(&works)
	if err != nil {
		return nil, 0, gerror.Wrap(err, "查询附加作业作品失败")
	}

	return works, total, nil
}

// StarToggle 点赞/取消点赞
func (s *sTeachingWork) StarToggle(ctx context.Context, workId string) (isStared bool, err error) {
	// 简化实现：直接增加点赞数
	_, err = dao.TeachingWork.Ctx(ctx).
		Where("id", workId).
		Increment("star_num", 1)
	if err != nil {
		return false, gerror.Wrap(err, "点赞失败")
	}

	return true, nil
}

// CollectToggle 收藏/取消收藏
func (s *sTeachingWork) CollectToggle(ctx context.Context, workId string) (isCollected bool, err error) {
	// 简化实现：直接增加收藏数
	_, err = dao.TeachingWork.Ctx(ctx).
		Where("id", workId).
		Increment("collect_num", 1)
	if err != nil {
		return false, gerror.Wrap(err, "收藏失败")
	}

	return true, nil
}

// CorrectList 批改记录列表
func (s *sTeachingWork) CorrectList(ctx context.Context, workId string) (list interface{}, err error) {
	var corrects []map[string]interface{}
	err = dao.TeachingWorkCorrect.Ctx(ctx).
		Where("work_id", workId).
		OrderDesc("create_time").
		Scan(&corrects)
	if err != nil {
		return nil, gerror.Wrap(err, "查询批改记录失败")
	}

	return corrects, nil
}

// CorrectAdd 添加批改记录
func (s *sTeachingWork) CorrectAdd(ctx context.Context, req *v1.WorkCorrectAddReq) error {
	username := gconv.String(ctx.Value(consts.CtxKeyUsername))
	orgCode := gconv.String(ctx.Value(consts.CtxKeyOrgCode))

	// 检查是否已有批改记录
	count, err := dao.TeachingWorkCorrect.Ctx(ctx).
		Where("work_id", req.WorkId).
		Count()
	if err != nil {
		return gerror.Wrap(err, "查询批改记录失败")
	}

	if count > 0 {
		// 更新现有记录
		_, err = dao.TeachingWorkCorrect.Ctx(ctx).
			Where("work_id", req.WorkId).
			Data(g.Map{
				"score":       req.Score,
				"comment":     req.Comment,
				"update_by":   username,
				"update_time": gtime.Now(),
			}).
			Update()
	} else {
		// 新增记录
		_, err = dao.TeachingWorkCorrect.Ctx(ctx).Insert(do.TeachingWorkCorrect{
			Id:         guid.S(),
			CreateBy:   username,
			CreateTime: gtime.Now(),
			SysOrgCode: orgCode,
			WorkId:     req.WorkId,
			Score:      req.Score,
			Comment:    req.Comment,
		})
	}
	if err != nil {
		return gerror.Wrap(err, "批改作品失败")
	}

	// 更新作品状态为已批改
	_, _ = dao.TeachingWork.Ctx(ctx).
		Where("id", req.WorkId).
		Data(g.Map{
			"work_status": 2,
			"update_by":   username,
			"update_time": gtime.Now(),
		}).
		Update()

	return nil
}

// CommentList 评论列表
func (s *sTeachingWork) CommentList(ctx context.Context, workId string) (list interface{}, err error) {
	var comments []map[string]interface{}
	err = dao.TeachingWorkComment.Ctx(ctx).
		Where("work_id", workId).
		OrderDesc("create_time").
		Scan(&comments)
	if err != nil {
		return nil, gerror.Wrap(err, "查询评论列表失败")
	}

	// 填充用户信息
	for i := range comments {
		userId := gconv.String(comments[i]["user_id"])
		if userId != "" {
			var user map[string]interface{}
			_ = dao.SysUser.Ctx(ctx).
				Fields("username", "realname", "avatar").
				Where("id", userId).
				Scan(&user)
			if user != nil {
				comments[i]["userName"] = user["username"]
				comments[i]["realname"] = user["realname"]
				comments[i]["avatar"] = user["avatar"]
			}
		}
	}

	return comments, nil
}

// CommentAdd 添加评论
func (s *sTeachingWork) CommentAdd(ctx context.Context, req *v1.WorkCommentAddReq) error {
	userId := gconv.String(ctx.Value(consts.CtxKeyUserId))
	username := gconv.String(ctx.Value(consts.CtxKeyUsername))
	orgCode := gconv.String(ctx.Value(consts.CtxKeyOrgCode))

	_, err := dao.TeachingWorkComment.Ctx(ctx).Insert(do.TeachingWorkComment{
		Id:         guid.S(),
		CreateBy:   username,
		CreateTime: gtime.Now(),
		SysOrgCode: orgCode,
		WorkId:     req.WorkId,
		Comment:    req.Comment,
		UserId:     userId,
	})
	if err != nil {
		return gerror.Wrap(err, "添加评论失败")
	}

	return nil
}

// CommentDelete 删除评论
func (s *sTeachingWork) CommentDelete(ctx context.Context, id string) error {
	_, err := dao.TeachingWorkComment.Ctx(ctx).
		Where("id", id).
		Delete()
	if err != nil {
		return gerror.Wrap(err, "删除评论失败")
	}

	return nil
}

// TagGet 获取作品标签
func (s *sTeachingWork) TagGet(ctx context.Context, workId string) (tags []string, err error) {
	// TODO: 如果需要专门的标签表，可以后续实现
	// 这里返回空列表
	return []string{}, nil
}

// TagSet 设置作品标签
func (s *sTeachingWork) TagSet(ctx context.Context, workId, tag string) error {
	// TODO: 如果需要专门的标签表，可以后续实现
	return nil
}

// TagDelete 删除作品标签
func (s *sTeachingWork) TagDelete(ctx context.Context, workId, tag string) error {
	// TODO: 如果需要专门的标签表，可以后续实现
	return nil
}

// fillUserInfo 填充用户信息
func (s *sTeachingWork) fillUserInfo(ctx context.Context, works []map[string]interface{}) []map[string]interface{} {
	for i := range works {
		userId := gconv.String(works[i]["user_id"])
		if userId != "" {
			var user map[string]interface{}
			_ = dao.SysUser.Ctx(ctx).
				Fields("username", "realname", "avatar").
				Where("id", userId).
				Scan(&user)
			if user != nil {
				works[i]["userName"] = user["username"]
				works[i]["realname"] = user["realname"]
				works[i]["avatar"] = user["avatar"]
			}
		}
	}
	return works
}

// fillSingleWorkInfo 填充单个作品的扩展信息
func (s *sTeachingWork) fillSingleWorkInfo(ctx context.Context, work *v1.WorkInfo) {
	if work == nil {
		return
	}

	// 填充用户信息
	if work.UserId != "" {
		var user map[string]interface{}
		_ = dao.SysUser.Ctx(ctx).
			Fields("username", "realname").
			Where("id", work.UserId).
			Scan(&user)
		if user != nil {
			work.UserName = gconv.String(user["username"])
			work.RealName = gconv.String(user["realname"])
		}
	}

	// 填充课程信息
	if work.CourseId != "" {
		var course map[string]interface{}
		_ = dao.TeachingCourse.Ctx(ctx).
			Fields("course_name").
			Where("id", work.CourseId).
			Scan(&course)
		if course != nil {
			work.CourseName = gconv.String(course["course_name"])
		}
	}

	// 填充部门信息
	if work.DepartId != "" {
		var depart map[string]interface{}
		_ = dao.SysDepart.Ctx(ctx).
			Fields("depart_name").
			Where("id", work.DepartId).
			Scan(&depart)
		if depart != nil {
			work.DepartName = gconv.String(depart["depart_name"])
		}
	}
}